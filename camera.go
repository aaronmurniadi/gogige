package gogige

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aaronmurniadi/gogige/genapi"
	"github.com/aaronmurniadi/gogige/gvcp"
)

// Camera is a connected GigE Vision device with GenICam feature access.
type Camera struct {
	IP     string
	gvcp   *gvcp.GVCP
	nodes  *genapi.NodeMap
	log    Logger
	mu     sync.Mutex
	closed bool
}

// New builds a Camera from an already-connected GVCP client and NodeMap.
func New(ip string, g *gvcp.GVCP, nodes *genapi.NodeMap, log Logger) *Camera {
	if log == nil {
		log = NopLogger{}
	}
	return &Camera{IP: ip, gvcp: g, nodes: nodes, log: log}
}

// Connect opens GVCP to the camera at ip, takes control, and loads GenICam XML.
func Connect(ip string) (*Camera, error) {
	if ip == "" {
		return nil, errors.New("gige: ip address must not be empty")
	}
	g, err := gvcp.DialGVCP(ip, 2*time.Second)
	if err != nil {
		return nil, err
	}
	if err := g.TakeControl(); err != nil {
		_ = g.Close()
		return nil, fmt.Errorf("gige: take control: %w", err)
	}
	xmlData, err := genapi.FetchXML(g)
	if err != nil {
		_ = g.LeaveControl()
		_ = g.Close()
		return nil, fmt.Errorf("gige: fetch XML: %w", err)
	}
	nm, err := genapi.ParseNodeMap(xmlData, g)
	if err != nil {
		_ = g.LeaveControl()
		_ = g.Close()
		return nil, err
	}
	// Refresh CCP after potentially long XML fetch (heartbeat window).
	if err := g.TakeControl(); err != nil {
		_ = g.Close()
		return nil, fmt.Errorf("gige: re-take control: %w", err)
	}
	return New(ip, g, nm, NopLogger{}), nil
}

// Close releases control and the GVCP socket.
func (c *Camera) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return
	}
	c.closed = true
	if c.gvcp == nil {
		return
	}
	_ = c.gvcp.LeaveControl()
	_ = c.gvcp.Close()
	c.gvcp = nil
}

// GVCP returns the underlying control client.
func (c *Camera) GVCP() *gvcp.GVCP {
	if c == nil {
		return nil
	}
	return c.gvcp
}

// Logger returns the camera logger (never nil).
func (c *Camera) Logger() Logger {
	if c == nil || c.log == nil {
		return NopLogger{}
	}
	return c.log
}

// NodeMap returns the loaded GenICam map.
func (c *Camera) NodeMap() *genapi.NodeMap { return c.nodes }

// Has reports whether a GenICam feature exists (gvcp.Commander).
func (c *Camera) Has(name string) bool {
	return c != nil && c.nodes != nil && c.nodes.Has(name)
}

// SetInteger sets an integer GenICam feature.
func (c *Camera) SetInteger(name string, v int64) error {
	return c.nodes.SetInteger(name, v)
}

// Integer reads the current value of an integer GenICam feature.
func (c *Camera) Integer(name string) (int64, error) {
	return c.nodes.ReadInteger(name)
}

// SetBoolean sets a boolean GenICam feature.
func (c *Camera) SetBoolean(name string, v bool) error {
	return c.nodes.SetBoolean(name, v)
}

// Boolean reads the current value of a boolean GenICam feature.
func (c *Camera) Boolean(name string) (bool, error) {
	return c.nodes.ReadBoolean(name)
}

// SetFloat sets a float GenICam feature.
func (c *Camera) SetFloat(name string, v float64) error {
	return c.nodes.SetFloat(name, v)
}

// Float reads the current value of a float GenICam feature.
func (c *Camera) Float(name string) (float64, error) {
	if c == nil || c.nodes == nil {
		return 0, errors.New("gige: camera not connected")
	}
	return c.nodes.ReadFloat(name)
}

// SetString sets a string or enumeration GenICam feature.
func (c *Camera) SetString(name, v string) error {
	return c.nodes.SetString(name, v)
}

// SetEnum sets an enumeration GenICam feature.
func (c *Camera) SetEnum(name, value string) error {
	return c.nodes.SetString(name, value)
}

// Enum reads the current enumeration symbol of an enumeration GenICam feature.
func (c *Camera) Enum(name string) (string, error) {
	if c == nil || c.nodes == nil {
		return "", errors.New("gige: camera not connected")
	}
	return c.nodes.CurrentEnum(name)
}

// String reads the current string value of a string GenICam feature.
func (c *Camera) String(name string) (string, error) {
	if c == nil || c.nodes == nil {
		return "", errors.New("gige: camera not connected")
	}
	return c.nodes.ReadString(name)
}

// ExecuteCommand executes a GenICam command node.
func (c *Camera) ExecuteCommand(name string) error {
	return c.nodes.Execute(name)
}

// Execute implements gvcp.Commander.
func (c *Camera) Execute(name string) error { return c.ExecuteCommand(name) }

// Features returns a Features view backed by this Camera, so callers that
// already hold a *Camera can use the same feature accessors as a Device.
func (c *Camera) Features() Features { return cameraFeatures{c: c} }

// GrabSample opens a transient GVSP stream on the camera, picks one frame, and
// returns it as a Sample with a JPEG for the requested component. comp of
// ComponentUnknown selects ComponentColor. Closes the stream before returning.
// Prefer StartStream for continuous capture; this is for one-off grabs from a
// Camera you already hold.
func (c *Camera) GrabSample(ctx context.Context, comp Component) (Sample, error) {
	if c == nil {
		return Sample{PackCount: -1}, errors.New("gige: nil camera")
	}
	s := NewFromCamera(c)
	if comp != ComponentUnknown {
		s.component = comp
	}
	if err := s.Open(c.IP); err != nil {
		return Sample{PackCount: -1}, err
	}
	defer s.Close()
	return s.Grab(ctx)
}

// GrabAllSamples opens a transient GVSP stream on the camera, picks one frame,
// and returns a Sample for every BSCF component (color, depth, mono, …).
func (c *Camera) GrabAllSamples(ctx context.Context) ([]Sample, error) {
	if c == nil {
		return nil, errors.New("gige: nil camera")
	}
	s := NewFromCamera(c)
	if err := s.Open(c.IP); err != nil {
		return nil, err
	}
	defer s.Close()
	return s.GrabAll(ctx)
}

// GrabComponents opens a transient GVSP stream on the camera and returns one
// frame's BSCF components without JPEG encoding (raw Data, dimensions, pixel
// format).
func (c *Camera) GrabComponents(ctx context.Context) ([]Sample, error) {
	if c == nil {
		return nil, errors.New("gige: nil camera")
	}
	s := NewFromCamera(c)
	if err := s.Open(c.IP); err != nil {
		return nil, err
	}
	defer s.Close()
	return s.GrabComponents(ctx)
}

// GrabJPEG opens a transient GVSP stream on the camera, picks one frame, and
// returns the JPEG bytes for the requested component (default color).
func (c *Camera) GrabJPEG(ctx context.Context, comp Component) ([]byte, error) {
	sample, err := c.GrabSample(ctx, comp)
	if err != nil {
		return nil, err
	}
	if len(sample.JPEG) == 0 {
		return nil, errors.New("gige: grab: empty frame")
	}
	return sample.JPEG, nil
}

// The following *Feature getters/setters are the pre-Phase-4 names, kept as
// thin aliases so existing code keeps compiling. Prefer the short forms
// (SetInteger, Integer, SetEnum, Enum, SetBool, Bool, SetFloat, Float,
// SetString, String) documented on Camera.

// SetBooleanFeature sets a boolean GenICam feature (alias for SetBoolean).
func (c *Camera) SetBooleanFeature(name string, v bool) error {
	return c.SetBoolean(name, v)
}

// BooleanFeature reads a boolean GenICam feature (alias for Boolean).
func (c *Camera) BooleanFeature(name string) (bool, error) {
	return c.Boolean(name)
}

// SetIntFeature sets an integer GenICam feature (alias for SetInteger).
func (c *Camera) SetIntFeature(name string, v int64) error {
	return c.SetInteger(name, v)
}

// SetFloatFeature sets a float GenICam feature (alias for SetFloat).
func (c *Camera) SetFloatFeature(name string, v float64) error {
	return c.SetFloat(name, v)
}

// SetStringFeature sets a string or enumeration GenICam feature (alias for
// SetString/SetEnum).
func (c *Camera) SetStringFeature(name, v string) error {
	return c.SetString(name, v)
}

// ApplyControlPair sets one GenICam feature from "Name=value" string.
func ApplyControlPair(c *Camera, pair string) error {
	eq := strings.IndexByte(pair, '=')
	if eq <= 0 || eq == len(pair)-1 {
		return fmt.Errorf("gige: bad control pair %q", pair)
	}
	feature := pair[:eq]
	val := pair[eq+1:]
	var err error
	switch strings.ToLower(val) {
	case "true":
		err = c.SetBooleanFeature(feature, true)
	case "false":
		err = c.SetBooleanFeature(feature, false)
	default:
		if i, perr := strconv.ParseInt(val, 10, 64); perr == nil && strconv.FormatInt(i, 10) == val {
			err = c.SetIntFeature(feature, i)
		} else if f, perr := strconv.ParseFloat(val, 64); perr == nil && strings.ContainsAny(val, ".eE") {
			err = c.SetFloatFeature(feature, f)
		} else {
			err = c.SetStringFeature(feature, val)
		}
	}
	if err != nil {
		return fmt.Errorf("gige: control %s: %w", pair, err)
	}
	return nil
}

var _ gvcp.Commander = (*Camera)(nil)
