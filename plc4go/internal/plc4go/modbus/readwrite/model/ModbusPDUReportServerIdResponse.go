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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUReportServerIdResponse struct {
	Value  []int8
	Parent *ModbusPDU
}

// The corresponding interface
type IModbusPDUReportServerIdResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReportServerIdResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUReportServerIdResponse) FunctionFlag() uint8 {
	return 0x11
}

func (m *ModbusPDUReportServerIdResponse) Response() bool {
	return true
}

func (m *ModbusPDUReportServerIdResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReportServerIdResponse(value []int8) *ModbusPDU {
	child := &ModbusPDUReportServerIdResponse{
		Value:  value,
		Parent: NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUReportServerIdResponse(structType interface{}) *ModbusPDUReportServerIdResponse {
	castFunc := func(typ interface{}) *ModbusPDUReportServerIdResponse {
		if casted, ok := typ.(ModbusPDUReportServerIdResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReportServerIdResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReportServerIdResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReportServerIdResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReportServerIdResponse) GetTypeName() string {
	return "ModbusPDUReportServerIdResponse"
}

func (m *ModbusPDUReportServerIdResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReportServerIdResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *ModbusPDUReportServerIdResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReportServerIdResponseParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := io.ReadUint8(8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (value)
	// Count array
	value := make([]int8, byteCount)
	for curItem := uint16(0); curItem < uint16(byteCount); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'value' field")
		}
		value[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReportServerIdResponse{
		Value:  value,
		Parent: &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUReportServerIdResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.Value)))
		_byteCountErr := io.WriteUint8(8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		if m.Value != nil {
			for _, _element := range m.Value {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'value' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReportServerIdResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "value":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Value = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *ModbusPDUReportServerIdResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_encodedValue := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Value))
	_encodedValue = strings.ToUpper(_encodedValue)
	if err := e.EncodeElement(_encodedValue, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUReportServerIdResponse) String() string {
	return string(m.Box("", 120))
}

func (m ModbusPDUReportServerIdResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "ModbusPDUReportServerIdResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Implicit Field (byteCount)
		byteCount := uint8(uint8(len(m.Value)))
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ByteCount", byteCount, -1))
		// Array Field (value)
		if m.Value != nil {
			// Simple array base type
			boxes = append(boxes, utils.BoxedDumpAnything("Value", m.Value))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
