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
	Component = gvsp.Component
)

// SFNC-style imaging components (BSCF ImageType wire values).
const (
	ComponentUnknown = gvsp.ComponentUnknown
	ComponentMono    = gvsp.ComponentMono
	ComponentDepth   = gvsp.ComponentDepth
	ComponentColor   = gvsp.ComponentColor
)

// Re-exports of common constructors / parsers.
var (
	DialGVCP                = gvcp.DialGVCP
	ParseNodeMap            = genapi.ParseNodeMap
	FetchXML                = genapi.FetchXML
	ListenStream            = gvsp.ListenStream
	SampleFromBSCF          = gvsp.SampleFromBSCF
	SampleFromBSCFComponent = gvsp.SampleFromBSCFComponent
	SampleAllFromBSCF       = gvsp.SampleAllFromBSCF
	ParseBSCF               = gvsp.ParseBSCF
	ParseComponent          = gvsp.ParseComponent
	IsBSCF                  = gvsp.IsBSCF
	EncodeJPEG              = color.EncodeJPEG
	StartAcquisition        = gvcp.StartAcquisition
	StopAcquisition         = gvcp.StopAcquisition
)
