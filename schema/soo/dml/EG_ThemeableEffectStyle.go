// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/vbatushev/gooxml"
)

type EG_ThemeableEffectStyle struct {
	Effect    *CT_EffectProperties
	EffectRef *CT_StyleMatrixReference
}

func NewEG_ThemeableEffectStyle() *EG_ThemeableEffectStyle {
	ret := &EG_ThemeableEffectStyle{}
	return ret
}

func (m *EG_ThemeableEffectStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "a:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.EffectRef != nil {
		seeffectRef := xml.StartElement{Name: xml.Name{Local: "a:effectRef"}}
		e.EncodeElement(m.EffectRef, seeffectRef)
	}
	return nil
}

func (m *EG_ThemeableEffectStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ThemeableEffectStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effect"}:
				m.Effect = NewCT_EffectProperties()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectRef"}:
				m.EffectRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.EffectRef, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_ThemeableEffectStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ThemeableEffectStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ThemeableEffectStyle and its children
func (m *EG_ThemeableEffectStyle) Validate() error {
	return m.ValidateWithPath("EG_ThemeableEffectStyle")
}

// ValidateWithPath validates the EG_ThemeableEffectStyle and its children, prefixing error messages with path
func (m *EG_ThemeableEffectStyle) ValidateWithPath(path string) error {
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.EffectRef != nil {
		if err := m.EffectRef.ValidateWithPath(path + "/EffectRef"); err != nil {
			return err
		}
	}
	return nil
}
