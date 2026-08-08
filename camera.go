package gogige

import (
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

// SetBooleanFeature sets a boolean GenICam feature.
func (c *Camera) SetBooleanFeature(name string, v bool) error {
	return c.nodes.SetBoolean(name, v)
}

// BooleanFeature reads the current value of a boolean GenICam feature.
func (c *Camera) BooleanFeature(name string) (bool, error) {
	return c.nodes.ReadBoolean(name)
}

// SetIntFeature sets an integer GenICam feature.
func (c *Camera) SetIntFeature(name string, v int64) error {
	return c.nodes.SetInteger(name, v)
}

// SetFloatFeature sets a float GenICam feature.
func (c *Camera) SetFloatFeature(name string, v float64) error {
	return c.nodes.SetFloat(name, v)
}

// SetStringFeature sets a string or enumeration GenICam feature.
func (c *Camera) SetStringFeature(name, v string) error {
	return c.nodes.SetString(name, v)
}

// ExecuteCommand executes a GenICam command node.
func (c *Camera) ExecuteCommand(name string) error {
	return c.nodes.Execute(name)
}

// Execute implements gvcp.Commander.
func (c *Camera) Execute(name string) error { return c.ExecuteCommand(name) }

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
