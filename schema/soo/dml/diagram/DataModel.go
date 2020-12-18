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

	"github.com/carmel/gooxml"
	"github.com/vbatushev/gooxml/schema/soo/dml"
)

type DataModel struct {
	CT_DataModel
}

func NewDataModel() *DataModel {
	ret := &DataModel{}
	ret.CT_DataModel = *NewCT_DataModel()
	return ret
}

func (m *DataModel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "dataModel"
	return m.CT_DataModel.MarshalXML(e, start)
}

func (m *DataModel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_DataModel = *NewCT_DataModel()
lDataModel:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "ptLst"}:
				if err := d.DecodeElement(m.PtLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "cxnLst"}:
				m.CxnLst = NewCT_CxnList()
				if err := d.DecodeElement(m.CxnLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "bg"}:
				m.Bg = dml.NewCT_BackgroundFormatting()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "whole"}:
				m.Whole = dml.NewCT_WholeE2oFormatting()
				if err := d.DecodeElement(m.Whole, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on DataModel %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lDataModel
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the DataModel and its children
func (m *DataModel) Validate() error {
	return m.ValidateWithPath("DataModel")
}

// ValidateWithPath validates the DataModel and its children, prefixing error messages with path
func (m *DataModel) ValidateWithPath(path string) error {
	if err := m.CT_DataModel.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
