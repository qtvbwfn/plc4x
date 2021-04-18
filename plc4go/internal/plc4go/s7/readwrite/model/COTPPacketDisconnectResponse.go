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
type COTPPacketDisconnectResponse struct {
	DestinationReference uint16
	SourceReference      uint16
	Parent               *COTPPacket
}

// The corresponding interface
type ICOTPPacketDisconnectResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketDisconnectResponse) TpduCode() uint8 {
	return 0xC0
}

func (m *COTPPacketDisconnectResponse) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.Parent.Parameters = parameters
	m.Parent.Payload = payload
}

func NewCOTPPacketDisconnectResponse(destinationReference uint16, sourceReference uint16, parameters []*COTPParameter, payload *S7Message) *COTPPacket {
	child := &COTPPacketDisconnectResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		Parent:               NewCOTPPacket(parameters, payload),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPPacketDisconnectResponse(structType interface{}) *COTPPacketDisconnectResponse {
	castFunc := func(typ interface{}) *COTPPacketDisconnectResponse {
		if casted, ok := typ.(COTPPacketDisconnectResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacketDisconnectResponse); ok {
			return casted
		}
		if casted, ok := typ.(COTPPacket); ok {
			return CastCOTPPacketDisconnectResponse(casted.Child)
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return CastCOTPPacketDisconnectResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacketDisconnectResponse) GetTypeName() string {
	return "COTPPacketDisconnectResponse"
}

func (m *COTPPacketDisconnectResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *COTPPacketDisconnectResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	return lengthInBits
}

func (m *COTPPacketDisconnectResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPPacketDisconnectResponseParse(io *utils.ReadBuffer) (*COTPPacket, error) {

	// Simple Field (destinationReference)
	destinationReference, _destinationReferenceErr := io.ReadUint16(16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field")
	}

	// Simple Field (sourceReference)
	sourceReference, _sourceReferenceErr := io.ReadUint16(16)
	if _sourceReferenceErr != nil {
		return nil, errors.Wrap(_sourceReferenceErr, "Error parsing 'sourceReference' field")
	}

	// Create a partially initialized instance
	_child := &COTPPacketDisconnectResponse{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		Parent:               &COTPPacket{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPPacketDisconnectResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (destinationReference)
		destinationReference := uint16(m.DestinationReference)
		_destinationReferenceErr := io.WriteUint16(16, (destinationReference))
		if _destinationReferenceErr != nil {
			return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
		}

		// Simple Field (sourceReference)
		sourceReference := uint16(m.SourceReference)
		_sourceReferenceErr := io.WriteUint16(16, (sourceReference))
		if _sourceReferenceErr != nil {
			return errors.Wrap(_sourceReferenceErr, "Error serializing 'sourceReference' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPPacketDisconnectResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "destinationReference":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DestinationReference = data
			case "sourceReference":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceReference = data
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

func (m *COTPPacketDisconnectResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DestinationReference, xml.StartElement{Name: xml.Name{Local: "destinationReference"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceReference, xml.StartElement{Name: xml.Name{Local: "sourceReference"}}); err != nil {
		return err
	}
	return nil
}

func (m COTPPacketDisconnectResponse) String() string {
	return string(m.Box("", 120))
}

func (m COTPPacketDisconnectResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "COTPPacketDisconnectResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DestinationReference", m.DestinationReference, -1))
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("SourceReference", m.SourceReference, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
