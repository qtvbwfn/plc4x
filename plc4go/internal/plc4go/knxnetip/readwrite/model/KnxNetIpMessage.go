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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const KnxNetIpMessage_PROTOCOLVERSION uint8 = 0x10

// The data-structure of this message
type KnxNetIpMessage struct {
	Child IKnxNetIpMessageChild
}

// The corresponding interface
type IKnxNetIpMessage interface {
	MsgType() uint16
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IKnxNetIpMessageParent interface {
	SerializeParent(io utils.WriteBuffer, child IKnxNetIpMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type IKnxNetIpMessageChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *KnxNetIpMessage)
	GetTypeName() string
	IKnxNetIpMessage
	utils.AsciiBoxer
}

func NewKnxNetIpMessage() *KnxNetIpMessage {
	return &KnxNetIpMessage{}
}

func CastKnxNetIpMessage(structType interface{}) *KnxNetIpMessage {
	castFunc := func(typ interface{}) *KnxNetIpMessage {
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxNetIpMessage) GetTypeName() string {
	return "KnxNetIpMessage"
}

func (m *KnxNetIpMessage) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxNetIpMessage) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *KnxNetIpMessage) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8

	// Const Field (protocolVersion)
	lengthInBits += 8
	// Discriminator Field (msgType)
	lengthInBits += 16

	// Implicit Field (totalLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *KnxNetIpMessage) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetIpMessageParse(io *utils.ReadBuffer) (*KnxNetIpMessage, error) {

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength, _headerLengthErr := io.ReadUint8(8)
	_ = headerLength
	if _headerLengthErr != nil {
		return nil, errors.Wrap(_headerLengthErr, "Error parsing 'headerLength' field")
	}

	// Const Field (protocolVersion)
	protocolVersion, _protocolVersionErr := io.ReadUint8(8)
	if _protocolVersionErr != nil {
		return nil, errors.Wrap(_protocolVersionErr, "Error parsing 'protocolVersion' field")
	}
	if protocolVersion != KnxNetIpMessage_PROTOCOLVERSION {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", KnxNetIpMessage_PROTOCOLVERSION) + " but got " + fmt.Sprintf("%d", protocolVersion))
	}

	// Discriminator Field (msgType) (Used as input to a switch field)
	msgType, _msgTypeErr := io.ReadUint16(16)
	if _msgTypeErr != nil {
		return nil, errors.Wrap(_msgTypeErr, "Error parsing 'msgType' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	totalLength, _totalLengthErr := io.ReadUint16(16)
	_ = totalLength
	if _totalLengthErr != nil {
		return nil, errors.Wrap(_totalLengthErr, "Error parsing 'totalLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *KnxNetIpMessage
	var typeSwitchError error
	switch {
	case msgType == 0x0201: // SearchRequest
		_parent, typeSwitchError = SearchRequestParse(io)
	case msgType == 0x0202: // SearchResponse
		_parent, typeSwitchError = SearchResponseParse(io)
	case msgType == 0x0203: // DescriptionRequest
		_parent, typeSwitchError = DescriptionRequestParse(io)
	case msgType == 0x0204: // DescriptionResponse
		_parent, typeSwitchError = DescriptionResponseParse(io)
	case msgType == 0x0205: // ConnectionRequest
		_parent, typeSwitchError = ConnectionRequestParse(io)
	case msgType == 0x0206: // ConnectionResponse
		_parent, typeSwitchError = ConnectionResponseParse(io)
	case msgType == 0x0207: // ConnectionStateRequest
		_parent, typeSwitchError = ConnectionStateRequestParse(io)
	case msgType == 0x0208: // ConnectionStateResponse
		_parent, typeSwitchError = ConnectionStateResponseParse(io)
	case msgType == 0x0209: // DisconnectRequest
		_parent, typeSwitchError = DisconnectRequestParse(io)
	case msgType == 0x020A: // DisconnectResponse
		_parent, typeSwitchError = DisconnectResponseParse(io)
	case msgType == 0x020B: // UnknownMessage
		_parent, typeSwitchError = UnknownMessageParse(io, totalLength)
	case msgType == 0x0310: // DeviceConfigurationRequest
		_parent, typeSwitchError = DeviceConfigurationRequestParse(io, totalLength)
	case msgType == 0x0311: // DeviceConfigurationAck
		_parent, typeSwitchError = DeviceConfigurationAckParse(io)
	case msgType == 0x0420: // TunnelingRequest
		_parent, typeSwitchError = TunnelingRequestParse(io, totalLength)
	case msgType == 0x0421: // TunnelingResponse
		_parent, typeSwitchError = TunnelingResponseParse(io)
	case msgType == 0x0530: // RoutingIndication
		_parent, typeSwitchError = RoutingIndicationParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *KnxNetIpMessage) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *KnxNetIpMessage) SerializeParent(io utils.WriteBuffer, child IKnxNetIpMessage, serializeChildFunction func() error) error {

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(6))
	_headerLengthErr := io.WriteUint8(8, (headerLength))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	// Const Field (protocolVersion)
	_protocolVersionErr := io.WriteUint8(8, 0x10)
	if _protocolVersionErr != nil {
		return errors.Wrap(_protocolVersionErr, "Error serializing 'protocolVersion' field")
	}

	// Discriminator Field (msgType) (Used as input to a switch field)
	msgType := uint16(child.MsgType())
	_msgTypeErr := io.WriteUint16(16, (msgType))

	if _msgTypeErr != nil {
		return errors.Wrap(_msgTypeErr, "Error serializing 'msgType' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	totalLength := uint16(uint16(m.LengthInBytes()))
	_totalLengthErr := io.WriteUint16(16, (totalLength))
	if _totalLengthErr != nil {
		return errors.Wrap(_totalLengthErr, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	return nil
}

func (m *KnxNetIpMessage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// RoutingIndication needs special treatment as it has no fields
		case "org.apache.plc4x.java.knxnetip.readwrite.RoutingIndication":
			if m.Child == nil {
				m.Child = &RoutingIndication{
					Parent: m,
				}
			}
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of KnxNetIpMessage")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.knxnetip.readwrite.SearchRequest":
					var dt *SearchRequest
					if m.Child != nil {
						dt = m.Child.(*SearchRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.SearchResponse":
					var dt *SearchResponse
					if m.Child != nil {
						dt = m.Child.(*SearchResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.DescriptionRequest":
					var dt *DescriptionRequest
					if m.Child != nil {
						dt = m.Child.(*DescriptionRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.DescriptionResponse":
					var dt *DescriptionResponse
					if m.Child != nil {
						dt = m.Child.(*DescriptionResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionRequest":
					var dt *ConnectionRequest
					if m.Child != nil {
						dt = m.Child.(*ConnectionRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionResponse":
					var dt *ConnectionResponse
					if m.Child != nil {
						dt = m.Child.(*ConnectionResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionStateRequest":
					var dt *ConnectionStateRequest
					if m.Child != nil {
						dt = m.Child.(*ConnectionStateRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionStateResponse":
					var dt *ConnectionStateResponse
					if m.Child != nil {
						dt = m.Child.(*ConnectionStateResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.DisconnectRequest":
					var dt *DisconnectRequest
					if m.Child != nil {
						dt = m.Child.(*DisconnectRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.DisconnectResponse":
					var dt *DisconnectResponse
					if m.Child != nil {
						dt = m.Child.(*DisconnectResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.UnknownMessage":
					var dt *UnknownMessage
					if m.Child != nil {
						dt = m.Child.(*UnknownMessage)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.DeviceConfigurationRequest":
					var dt *DeviceConfigurationRequest
					if m.Child != nil {
						dt = m.Child.(*DeviceConfigurationRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.DeviceConfigurationAck":
					var dt *DeviceConfigurationAck
					if m.Child != nil {
						dt = m.Child.(*DeviceConfigurationAck)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.TunnelingRequest":
					var dt *TunnelingRequest
					if m.Child != nil {
						dt = m.Child.(*TunnelingRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.TunnelingResponse":
					var dt *TunnelingResponse
					if m.Child != nil {
						dt = m.Child.(*TunnelingResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.RoutingIndication":
					var dt *RoutingIndication
					if m.Child != nil {
						dt = m.Child.(*RoutingIndication)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *KnxNetIpMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m KnxNetIpMessage) String() string {
	return string(m.Box("", 120))
}

func (m *KnxNetIpMessage) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *KnxNetIpMessage) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "KnxNetIpMessage"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Implicit Field (headerLength)
	headerLength := uint8(uint8(6))
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("HeaderLength", headerLength, -1))
	// Const Field (protocolVersion)
	boxes = append(boxes, utils.BoxAnything("ProtocolVersion", 0x10, -1))
	// Discriminator Field (msgType) (Used as input to a switch field)
	// msgType := uint16(child.MsgType())
	// uint16 can be boxed as anything with the least amount of space
	// boxes = append(boxes, utils.BoxAnything("MsgType", msgType, -1))
	// Implicit Field (totalLength)
	totalLength := uint16(uint16(m.LengthInBytes()))
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TotalLength", totalLength, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
