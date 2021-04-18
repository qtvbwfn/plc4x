//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type COTPParameterCalledTsap struct {
	TsapId uint16
	Parent *COTPParameter
}

// The corresponding interface
type ICOTPParameterCalledTsap interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPParameterCalledTsap) ParameterType() uint8 {
	return 0xC2
}

func (m *COTPParameterCalledTsap) InitializeParent(parent *COTPParameter) {
}

func NewCOTPParameterCalledTsap(tsapId uint16) *COTPParameter {
	child := &COTPParameterCalledTsap{
		TsapId: tsapId,
		Parent: NewCOTPParameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPParameterCalledTsap(structType interface{}) *COTPParameterCalledTsap {
	castFunc := func(typ interface{}) *COTPParameterCalledTsap {
		if casted, ok := typ.(COTPParameterCalledTsap); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPParameterCalledTsap); ok {
			return casted
		}
		if casted, ok := typ.(COTPParameter); ok {
			return CastCOTPParameterCalledTsap(casted.Child)
		}
		if casted, ok := typ.(*COTPParameter); ok {
			return CastCOTPParameterCalledTsap(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPParameterCalledTsap) GetTypeName() string {
	return "COTPParameterCalledTsap"
}

func (m *COTPParameterCalledTsap) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *COTPParameterCalledTsap) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (tsapId)
	lengthInBits += 16

	return lengthInBits
}

func (m *COTPParameterCalledTsap) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPParameterCalledTsapParse(io *utils.ReadBuffer) (*COTPParameter, error) {

	// Simple Field (tsapId)
	tsapId, _tsapIdErr := io.ReadUint16(16)
	if _tsapIdErr != nil {
		return nil, errors.Wrap(_tsapIdErr, "Error parsing 'tsapId' field")
	}

	// Create a partially initialized instance
	_child := &COTPParameterCalledTsap{
		TsapId: tsapId,
		Parent: &COTPParameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPParameterCalledTsap) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (tsapId)
		tsapId := uint16(m.TsapId)
		_tsapIdErr := io.WriteUint16(16, (tsapId))
		if _tsapIdErr != nil {
			return errors.Wrap(_tsapIdErr, "Error serializing 'tsapId' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPParameterCalledTsap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "tsapId":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TsapId = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *COTPParameterCalledTsap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.TsapId, xml.StartElement{Name: xml.Name{Local: "tsapId"}}); err != nil {
		return err
	}
	return nil
}

func (m COTPParameterCalledTsap) String() string {
	return string(m.Box("", 120))
}

func (m COTPParameterCalledTsap) Box(name string, width int) utils.AsciiBox {
	boxName := "COTPParameterCalledTsap"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("TsapId", m.TsapId, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
