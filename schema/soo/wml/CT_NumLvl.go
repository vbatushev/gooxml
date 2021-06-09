// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/vbatushev/gooxml"
)

type CT_NumLvl struct {
	// Numbering Level ID
	IlvlAttr int64
	// Numbering Level Starting Value Override
	StartOverride *CT_DecimalNumber
	// Numbering Level Override Definition
	Lvl *CT_Lvl
}

func NewCT_NumLvl() *CT_NumLvl {
	ret := &CT_NumLvl{}
	return ret
}

func (m *CT_NumLvl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:ilvl"},
		Value: fmt.Sprintf("%v", m.IlvlAttr)})
	e.EncodeToken(start)
	if m.StartOverride != nil {
		sestartOverride := xml.StartElement{Name: xml.Name{Local: "w:startOverride"}}
		e.EncodeElement(m.StartOverride, sestartOverride)
	}
	if m.Lvl != nil {
		selvl := xml.StartElement{Name: xml.Name{Local: "w:lvl"}}
		e.EncodeElement(m.Lvl, selvl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumLvl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ilvl" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IlvlAttr = parsed
			continue
		}
	}
lCT_NumLvl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "startOverride"}:
				m.StartOverride = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.StartOverride, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lvl"}:
				m.Lvl = NewCT_Lvl()
				if err := d.DecodeElement(m.Lvl, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_NumLvl %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumLvl
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumLvl and its children
func (m *CT_NumLvl) Validate() error {
	return m.ValidateWithPath("CT_NumLvl")
}

// ValidateWithPath validates the CT_NumLvl and its children, prefixing error messages with path
func (m *CT_NumLvl) ValidateWithPath(path string) error {
	if m.StartOverride != nil {
		if err := m.StartOverride.ValidateWithPath(path + "/StartOverride"); err != nil {
			return err
		}
	}
	if m.Lvl != nil {
		if err := m.Lvl.ValidateWithPath(path + "/Lvl"); err != nil {
			return err
		}
	}
	return nil
}
