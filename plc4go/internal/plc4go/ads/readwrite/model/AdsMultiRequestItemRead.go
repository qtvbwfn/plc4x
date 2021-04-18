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
type AdsMultiRequestItemRead struct {
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemReadLength  uint32
	Parent          *AdsMultiRequestItem
}

// The corresponding interface
type IAdsMultiRequestItemRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsMultiRequestItemRead) IndexGroup() uint32 {
	return 61568
}

func (m *AdsMultiRequestItemRead) InitializeParent(parent *AdsMultiRequestItem) {
}

func NewAdsMultiRequestItemRead(itemIndexGroup uint32, itemIndexOffset uint32, itemReadLength uint32) *AdsMultiRequestItem {
	child := &AdsMultiRequestItemRead{
		ItemIndexGroup:  itemIndexGroup,
		ItemIndexOffset: itemIndexOffset,
		ItemReadLength:  itemReadLength,
		Parent:          NewAdsMultiRequestItem(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsMultiRequestItemRead(structType interface{}) *AdsMultiRequestItemRead {
	castFunc := func(typ interface{}) *AdsMultiRequestItemRead {
		if casted, ok := typ.(AdsMultiRequestItemRead); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsMultiRequestItemRead); ok {
			return casted
		}
		if casted, ok := typ.(AdsMultiRequestItem); ok {
			return CastAdsMultiRequestItemRead(casted.Child)
		}
		if casted, ok := typ.(*AdsMultiRequestItem); ok {
			return CastAdsMultiRequestItemRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsMultiRequestItemRead) GetTypeName() string {
	return "AdsMultiRequestItemRead"
}

func (m *AdsMultiRequestItemRead) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsMultiRequestItemRead) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemReadLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsMultiRequestItemRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsMultiRequestItemReadParse(io *utils.ReadBuffer) (*AdsMultiRequestItem, error) {

	// Simple Field (itemIndexGroup)
	itemIndexGroup, _itemIndexGroupErr := io.ReadUint32(32)
	if _itemIndexGroupErr != nil {
		return nil, errors.Wrap(_itemIndexGroupErr, "Error parsing 'itemIndexGroup' field")
	}

	// Simple Field (itemIndexOffset)
	itemIndexOffset, _itemIndexOffsetErr := io.ReadUint32(32)
	if _itemIndexOffsetErr != nil {
		return nil, errors.Wrap(_itemIndexOffsetErr, "Error parsing 'itemIndexOffset' field")
	}

	// Simple Field (itemReadLength)
	itemReadLength, _itemReadLengthErr := io.ReadUint32(32)
	if _itemReadLengthErr != nil {
		return nil, errors.Wrap(_itemReadLengthErr, "Error parsing 'itemReadLength' field")
	}

	// Create a partially initialized instance
	_child := &AdsMultiRequestItemRead{
		ItemIndexGroup:  itemIndexGroup,
		ItemIndexOffset: itemIndexOffset,
		ItemReadLength:  itemReadLength,
		Parent:          &AdsMultiRequestItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsMultiRequestItemRead) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (itemIndexGroup)
		itemIndexGroup := uint32(m.ItemIndexGroup)
		_itemIndexGroupErr := io.WriteUint32(32, (itemIndexGroup))
		if _itemIndexGroupErr != nil {
			return errors.Wrap(_itemIndexGroupErr, "Error serializing 'itemIndexGroup' field")
		}

		// Simple Field (itemIndexOffset)
		itemIndexOffset := uint32(m.ItemIndexOffset)
		_itemIndexOffsetErr := io.WriteUint32(32, (itemIndexOffset))
		if _itemIndexOffsetErr != nil {
			return errors.Wrap(_itemIndexOffsetErr, "Error serializing 'itemIndexOffset' field")
		}

		// Simple Field (itemReadLength)
		itemReadLength := uint32(m.ItemReadLength)
		_itemReadLengthErr := io.WriteUint32(32, (itemReadLength))
		if _itemReadLengthErr != nil {
			return errors.Wrap(_itemReadLengthErr, "Error serializing 'itemReadLength' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsMultiRequestItemRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemIndexGroup":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ItemIndexGroup = data
			case "itemIndexOffset":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ItemIndexOffset = data
			case "itemReadLength":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ItemReadLength = data
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

func (m *AdsMultiRequestItemRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.ItemIndexGroup, xml.StartElement{Name: xml.Name{Local: "itemIndexGroup"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ItemIndexOffset, xml.StartElement{Name: xml.Name{Local: "itemIndexOffset"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ItemReadLength, xml.StartElement{Name: xml.Name{Local: "itemReadLength"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsMultiRequestItemRead) String() string {
	return string(m.Box("", 120))
}

func (m AdsMultiRequestItemRead) Box(name string, width int) utils.AsciiBox {
	boxName := "AdsMultiRequestItemRead"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ItemIndexGroup", m.ItemIndexGroup, -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ItemIndexOffset", m.ItemIndexOffset, -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ItemReadLength", m.ItemReadLength, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
