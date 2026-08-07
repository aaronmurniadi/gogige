package gige

import (
	"github.com/aaronmurniadi/gogige/genapi"
	"github.com/aaronmurniadi/gogige/gvcp"
	"github.com/aaronmurniadi/gogige/gvsp"
	"github.com/aaronmurniadi/gogige/internal/color"
)

// Type aliases keep the root consumer API ergonomic for protocol types.
type (
	Port      = gvcp.Port
	GVCP      = gvcp.GVCP
	NodeMap   = genapi.NodeMap
	Stream    = gvsp.Stream
	Frame     = gvsp.Frame
	Sample    = gvsp.Sample
	ImageKind = gvsp.ImageKind
)

// BSCF image kinds (Huaray Frame.h ImageType).
const (
	ImageUnknown = gvsp.ImageUnknown
	ImageMono    = gvsp.ImageMono
	ImageDepth   = gvsp.ImageDepth
	ImageColor   = gvsp.ImageColor
)

// Re-exports of common constructors / parsers.
var (
	DialGVCP           = gvcp.DialGVCP
	ParseNodeMap       = genapi.ParseNodeMap
	FetchXML           = genapi.FetchXML
	ListenStream       = gvsp.ListenStream
	SampleFromBSCF     = gvsp.SampleFromBSCF
	SampleFromBSCFKind = gvsp.SampleFromBSCFKind
	SampleAllFromBSCF  = gvsp.SampleAllFromBSCF
	ParseBSCF          = gvsp.ParseBSCF
	ParseImageKind     = gvsp.ParseImageKind
	IsBSCF             = gvsp.IsBSCF
	EncodeJPEG         = color.EncodeJPEG
	StartAcquisition   = gvcp.StartAcquisition
	StopAcquisition    = gvcp.StopAcquisition
)
