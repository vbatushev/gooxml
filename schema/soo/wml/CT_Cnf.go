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

	"github.com/vbatushev/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Cnf struct {
	// Conditional Formatting Bit Mask
	ValAttr *string
	// First Row
	FirstRowAttr *sharedTypes.ST_OnOff
	// Last Row
	LastRowAttr *sharedTypes.ST_OnOff
	// First Column
	FirstColumnAttr *sharedTypes.ST_OnOff
	// Last Column
	LastColumnAttr *sharedTypes.ST_OnOff
	// Odd Numbered Vertical Band
	OddVBandAttr *sharedTypes.ST_OnOff
	// Even Numbered Vertical Band
	EvenVBandAttr *sharedTypes.ST_OnOff
	// Odd Numbered Horizontal Band
	OddHBandAttr *sharedTypes.ST_OnOff
	// Even Numbered Horizontal Band
	EvenHBandAttr *sharedTypes.ST_OnOff
	// First Row and First Column
	FirstRowFirstColumnAttr *sharedTypes.ST_OnOff
	// First Row and Last Column
	FirstRowLastColumnAttr *sharedTypes.ST_OnOff
	// Last Row and First Column
	LastRowFirstColumnAttr *sharedTypes.ST_OnOff
	// Last Row and Last Column
	LastRowLastColumnAttr *sharedTypes.ST_OnOff
}

func NewCT_Cnf() *CT_Cnf {
	ret := &CT_Cnf{}
	return ret
}

func (m *CT_Cnf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.FirstRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstRow"},
			Value: fmt.Sprintf("%v", *m.FirstRowAttr)})
	}
	if m.LastRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastRow"},
			Value: fmt.Sprintf("%v", *m.LastRowAttr)})
	}
	if m.FirstColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstColumn"},
			Value: fmt.Sprintf("%v", *m.FirstColumnAttr)})
	}
	if m.LastColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastColumn"},
			Value: fmt.Sprintf("%v", *m.LastColumnAttr)})
	}
	if m.OddVBandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:oddVBand"},
			Value: fmt.Sprintf("%v", *m.OddVBandAttr)})
	}
	if m.EvenVBandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:evenVBand"},
			Value: fmt.Sprintf("%v", *m.EvenVBandAttr)})
	}
	if m.OddHBandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:oddHBand"},
			Value: fmt.Sprintf("%v", *m.OddHBandAttr)})
	}
	if m.EvenHBandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:evenHBand"},
			Value: fmt.Sprintf("%v", *m.EvenHBandAttr)})
	}
	if m.FirstRowFirstColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstRowFirstColumn"},
			Value: fmt.Sprintf("%v", *m.FirstRowFirstColumnAttr)})
	}
	if m.FirstRowLastColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstRowLastColumn"},
			Value: fmt.Sprintf("%v", *m.FirstRowLastColumnAttr)})
	}
	if m.LastRowFirstColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastRowFirstColumn"},
			Value: fmt.Sprintf("%v", *m.LastRowFirstColumnAttr)})
	}
	if m.LastRowLastColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastRowLastColumn"},
			Value: fmt.Sprintf("%v", *m.LastRowLastColumnAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Cnf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "evenVBand" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.EvenVBandAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstRow" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FirstRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastRow" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LastRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FirstColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LastColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "oddVBand" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.OddVBandAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
		if attr.Name.Local == "oddHBand" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.OddHBandAttr = &parsed
			continue
		}
		if attr.Name.Local == "evenHBand" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.EvenHBandAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstRowFirstColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FirstRowFirstColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstRowLastColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FirstRowLastColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastRowFirstColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LastRowFirstColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastRowLastColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LastRowLastColumnAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Cnf: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Cnf and its children
func (m *CT_Cnf) Validate() error {
	return m.ValidateWithPath("CT_Cnf")
}

// ValidateWithPath validates the CT_Cnf and its children, prefixing error messages with path
func (m *CT_Cnf) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if !ST_CnfPatternRe.MatchString(*m.ValAttr) {
			return fmt.Errorf(`%s/m.ValAttr must match '%s' (have %v)`, path, ST_CnfPatternRe, *m.ValAttr)
		}
	}
	if m.FirstRowAttr != nil {
		if err := m.FirstRowAttr.ValidateWithPath(path + "/FirstRowAttr"); err != nil {
			return err
		}
	}
	if m.LastRowAttr != nil {
		if err := m.LastRowAttr.ValidateWithPath(path + "/LastRowAttr"); err != nil {
			return err
		}
	}
	if m.FirstColumnAttr != nil {
		if err := m.FirstColumnAttr.ValidateWithPath(path + "/FirstColumnAttr"); err != nil {
			return err
		}
	}
	if m.LastColumnAttr != nil {
		if err := m.LastColumnAttr.ValidateWithPath(path + "/LastColumnAttr"); err != nil {
			return err
		}
	}
	if m.OddVBandAttr != nil {
		if err := m.OddVBandAttr.ValidateWithPath(path + "/OddVBandAttr"); err != nil {
			return err
		}
	}
	if m.EvenVBandAttr != nil {
		if err := m.EvenVBandAttr.ValidateWithPath(path + "/EvenVBandAttr"); err != nil {
			return err
		}
	}
	if m.OddHBandAttr != nil {
		if err := m.OddHBandAttr.ValidateWithPath(path + "/OddHBandAttr"); err != nil {
			return err
		}
	}
	if m.EvenHBandAttr != nil {
		if err := m.EvenHBandAttr.ValidateWithPath(path + "/EvenHBandAttr"); err != nil {
			return err
		}
	}
	if m.FirstRowFirstColumnAttr != nil {
		if err := m.FirstRowFirstColumnAttr.ValidateWithPath(path + "/FirstRowFirstColumnAttr"); err != nil {
			return err
		}
	}
	if m.FirstRowLastColumnAttr != nil {
		if err := m.FirstRowLastColumnAttr.ValidateWithPath(path + "/FirstRowLastColumnAttr"); err != nil {
			return err
		}
	}
	if m.LastRowFirstColumnAttr != nil {
		if err := m.LastRowFirstColumnAttr.ValidateWithPath(path + "/LastRowFirstColumnAttr"); err != nil {
			return err
		}
	}
	if m.LastRowLastColumnAttr != nil {
		if err := m.LastRowLastColumnAttr.ValidateWithPath(path + "/LastRowLastColumnAttr"); err != nil {
			return err
		}
	}
	return nil
}
