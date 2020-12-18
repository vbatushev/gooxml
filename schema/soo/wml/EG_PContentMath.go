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

	"github.com/carmel/gooxml"
	"github.com/vbatushev/gooxml/schema/soo/ofc/math"
)

type EG_PContentMath struct {
	EG_PContentBase          []*EG_PContentBase
	EG_ContentRunContentBase []*EG_ContentRunContentBase
}

func NewEG_PContentMath() *EG_PContentMath {
	ret := &EG_PContentMath{}
	return ret
}

func (m *EG_PContentMath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:EG_PContentMath"
	if m.EG_PContentBase != nil {
		for _, c := range m.EG_PContentBase {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	if m.EG_ContentRunContentBase != nil {
		for _, c := range m.EG_ContentRunContentBase {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	return nil
}

func (m *EG_PContentMath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_PContentMath:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"}:
				tmppcontentbase := NewEG_PContentBase()
				tmppcontentbase.CustomXml = NewCT_CustomXmlRun()
				if err := d.DecodeElement(tmppcontentbase.CustomXml, &el); err != nil {
					return err
				}
				m.EG_PContentBase = append(m.EG_PContentBase, tmppcontentbase)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fldSimple"}:
				tmppcontentbase := NewEG_PContentBase()
				tmp := NewCT_SimpleField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmppcontentbase.FldSimple = append(tmppcontentbase.FldSimple, tmp)
				m.EG_PContentBase = append(m.EG_PContentBase, tmppcontentbase)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hyperlink"}:
				tmppcontentbase := NewEG_PContentBase()
				tmppcontentbase.Hyperlink = NewCT_Hyperlink()
				if err := d.DecodeElement(tmppcontentbase.Hyperlink, &el); err != nil {
					return err
				}
				m.EG_PContentBase = append(m.EG_PContentBase, tmppcontentbase)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "smartTag"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmpcontentruncontentbase.SmartTag = NewCT_SmartTagRun()
				if err := d.DecodeElement(tmpcontentruncontentbase.SmartTag, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdt"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmpcontentruncontentbase.Sdt = NewCT_SdtRun()
				if err := d.DecodeElement(tmpcontentruncontentbase.Sdt, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofErr"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMathPara"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMath"}:
				tmpcontentruncontentbase := NewEG_ContentRunContentBase()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_ContentRunContentBase = append(m.EG_ContentRunContentBase, tmpcontentruncontentbase)
				tmpcontentruncontentbase.EG_RunLevelElts = append(tmpcontentruncontentbase.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				gooxml.Log("skipping unsupported element on EG_PContentMath %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_PContentMath
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_PContentMath and its children
func (m *EG_PContentMath) Validate() error {
	return m.ValidateWithPath("EG_PContentMath")
}

// ValidateWithPath validates the EG_PContentMath and its children, prefixing error messages with path
func (m *EG_PContentMath) ValidateWithPath(path string) error {
	for i, v := range m.EG_PContentBase {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_PContentBase[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.EG_ContentRunContentBase {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ContentRunContentBase[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
