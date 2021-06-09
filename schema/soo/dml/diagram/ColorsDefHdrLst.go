// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"

	"github.com/vbatushev/gooxml"
)

type ColorsDefHdrLst struct {
	CT_ColorTransformHeaderLst
}

func NewColorsDefHdrLst() *ColorsDefHdrLst {
	ret := &ColorsDefHdrLst{}
	ret.CT_ColorTransformHeaderLst = *NewCT_ColorTransformHeaderLst()
	return ret
}

func (m *ColorsDefHdrLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "colorsDefHdrLst"
	return m.CT_ColorTransformHeaderLst.MarshalXML(e, start)
}

func (m *ColorsDefHdrLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ColorTransformHeaderLst = *NewCT_ColorTransformHeaderLst()
lColorsDefHdrLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "colorsDefHdr"}:
				tmp := NewCT_ColorTransformHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ColorsDefHdr = append(m.ColorsDefHdr, tmp)
			default:
				gooxml.Log("skipping unsupported element on ColorsDefHdrLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lColorsDefHdrLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ColorsDefHdrLst and its children
func (m *ColorsDefHdrLst) Validate() error {
	return m.ValidateWithPath("ColorsDefHdrLst")
}

// ValidateWithPath validates the ColorsDefHdrLst and its children, prefixing error messages with path
func (m *ColorsDefHdrLst) ValidateWithPath(path string) error {
	if err := m.CT_ColorTransformHeaderLst.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
