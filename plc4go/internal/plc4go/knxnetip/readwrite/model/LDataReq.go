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
type LDataReq struct {
	AdditionalInformationLength uint8
	AdditionalInformation       []*CEMIAdditionalInformation
	DataFrame                   *LDataFrame
	Parent                      *CEMI
}

// The corresponding interface
type ILDataReq interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataReq) MessageCode() uint8 {
	return 0x11
}

func (m *LDataReq) InitializeParent(parent *CEMI) {
}

func NewLDataReq(additionalInformationLength uint8, additionalInformation []*CEMIAdditionalInformation, dataFrame *LDataFrame) *CEMI {
	child := &LDataReq{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Parent:                      NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLDataReq(structType interface{}) *LDataReq {
	castFunc := func(typ interface{}) *LDataReq {
		if casted, ok := typ.(LDataReq); ok {
			return &casted
		}
		if casted, ok := typ.(*LDataReq); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLDataReq(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLDataReq(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LDataReq) GetTypeName() string {
	return "LDataReq"
}

func (m *LDataReq) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LDataReq) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.LengthInBits()

	return lengthInBits
}

func (m *LDataReq) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LDataReqParse(io *utils.ReadBuffer) (*CEMI, error) {

	// Simple Field (additionalInformationLength)
	additionalInformationLength, _additionalInformationLengthErr := io.ReadUint8(8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field")
	}

	// Array field (additionalInformation)
	// Length array
	additionalInformation := make([]*CEMIAdditionalInformation, 0)
	_additionalInformationLength := additionalInformationLength
	_additionalInformationEndPos := io.GetPos() + uint16(_additionalInformationLength)
	for io.GetPos() < _additionalInformationEndPos {
		_item, _err := CEMIAdditionalInformationParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field")
		}
		additionalInformation = append(additionalInformation, _item)
	}

	// Simple Field (dataFrame)
	dataFrame, _dataFrameErr := LDataFrameParse(io)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field")
	}

	// Create a partially initialized instance
	_child := &LDataReq{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Parent:                      &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LDataReq) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.AdditionalInformationLength)
		_additionalInformationLengthErr := io.WriteUint8(8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			for _, _element := range m.AdditionalInformation {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
				}
			}
		}

		// Simple Field (dataFrame)
		_dataFrameErr := m.DataFrame.Serialize(io)
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *LDataReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "additionalInformationLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdditionalInformationLength = data
			case "additionalInformation":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *CEMIAdditionalInformation
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.AdditionalInformation = append(m.AdditionalInformation, dt)
					default:
						break arrayLoop
					}
				}
			case "dataFrame":
				var dt *LDataFrame
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.DataFrame = dt
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

func (m *LDataReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.AdditionalInformationLength, xml.StartElement{Name: xml.Name{Local: "additionalInformationLength"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.AdditionalInformation {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DataFrame, xml.StartElement{Name: xml.Name{Local: "dataFrame"}}); err != nil {
		return err
	}
	return nil
}

func (m LDataReq) String() string {
	return string(m.Box("", 120))
}

func (m LDataReq) Box(name string, width int) utils.AsciiBox {
	boxName := "LDataReq"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("AdditionalInformationLength", m.AdditionalInformationLength, -1))
		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.AdditionalInformation {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("AdditionalInformation", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Simple field (case simple)
		// TODO  waaaa org.apache.plc4x.plugins.codegenerator.types.references.DefaultComplexTypeReference@3348c987
		boxes = append(boxes, m.DataFrame.Box("dataFrame", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
