package vim25

import "encoding/xml"

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.html?path=7_0_0_2_6_17_8#retrievePropertiesEx
type RetrievePropertiesEx struct {
	XMLName xml.Name              `xml:"urn:vim25 RetrievePropertiesEx"`
	This    *PropertyCollector    `xml:"_this"`
	SpecSet []*PropertyFilterSpec `xml:"specSet"`
	Options RetrieveOptions       `xml:"options"`
}

type RetrievePropertiesExResponse struct {
	XMLName        xml.Name       `xml:"urn:vim25 RetrievePropertiesExResponse"`
	RetrieveResult RetrieveResult `xml:"returnval"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.ObjectSpec.html
type ObjectSpec struct {
	XsiType   string                  `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr,ommitempty"`
	Obj       *ManagedObjectReference `xml:"obj"`
	Skip      bool                    `xml:"skip"`
	SelectSet []interface{}           `xml:"selectSet"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.SelectionSpec.html
type SelectionSpec struct {
	Name string `xml:"name"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.TraversalSpec.html
type TraversalSpec struct {
	SelectionSpec
	XsiType string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Type    string `xml:"type"`
	Path    string `xml:"path"`
	Skip    bool   `xml:"skip"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.PropertySpec.html
type PropertySpec struct {
	Type    string   `xml:"type"`
	PathSet []string `xml:"pathSet"`
	All     bool     `xml:"all,omitempty"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.FilterSpec.html
type PropertyFilterSpec struct {
	PropSet                       []*PropertySpec `xml:"propSet"`
	ObjectSet                     []*ObjectSpec   `xml:"objectSet"`
	ReportMissingObjectsInResults bool            `xml:"reportMissingObjectsInResults"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvmodl.query.PropertyCollector.RetrieveOptions.html
type RetrieveOptions struct {
	MaxObjects int `xml:"maxObjects,omitempty"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.RetrieveResult.html
type RetrieveResult struct {
	Objects []ObjectContent `xml:"objects"`
	Token   string          `xml:"token"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/vmodl.query.PropertyCollector.ObjectContent.html
type ObjectContent struct {
	MissingSet []MissingProperty       `xml:"missingSet"`
	Obj        *ManagedObjectReference `xml:"obj"`
	PropSet    []DynamicProperty       `xml:"propSet"`
}

// http://pubs.vmware.com/vsphere-55/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvmodl.query.PropertyCollector.MissingProperty.html
type MissingProperty struct {
	// Fault LocalizedMethodFault `xml:"fault"` // TODO(igm)
	Path string `xml:"path"`
}
