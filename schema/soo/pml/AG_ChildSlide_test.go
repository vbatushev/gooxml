// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/vbatushev/gooxml/schema/soo/pml"
)

func TestAG_ChildSlideConstructor(t *testing.T) {
	v := pml.NewAG_ChildSlide()
	if v == nil {
		t.Errorf("pml.NewAG_ChildSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.AG_ChildSlide should validate: %s", err)
	}
}

func TestAG_ChildSlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewAG_ChildSlide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewAG_ChildSlide()
	xml.Unmarshal(buf, v2)
}
