/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetContextTagEventType is the data-structure of this message
type BACnetContextTagEventType struct {
	*BACnetContextTag
	EventType        BACnetEventType
	ProprietaryValue uint32

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
	ActualLength             uint32
}

// IBACnetContextTagEventType is the corresponding interface of BACnetContextTagEventType
type IBACnetContextTagEventType interface {
	IBACnetContextTag
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetContextTagEventType) GetDataType() BACnetDataType {
	return BACnetDataType_EVENT_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagEventType) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagEventType) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetContextTagEventType) GetEventType() BACnetEventType {
	return m.EventType
}

func (m *BACnetContextTagEventType) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetContextTagEventType) GetIsProprietary() bool {
	return bool(bool((m.GetEventType()) == (BACnetEventType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagEventType factory function for BACnetContextTagEventType
func NewBACnetContextTagEventType(eventType BACnetEventType, proprietaryValue uint32, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool, actualLength uint32) *BACnetContextTagEventType {
	_result := &BACnetContextTagEventType{
		EventType:        eventType,
		ProprietaryValue: proprietaryValue,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetContextTagEventType(structType interface{}) *BACnetContextTagEventType {
	if casted, ok := structType.(BACnetContextTagEventType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagEventType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagEventType(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagEventType(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagEventType) GetTypeName() string {
	return "BACnetContextTagEventType"
}

func (m *BACnetContextTagEventType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagEventType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Manual Field (eventType)
	lengthInBits += uint16(int32(m.GetActualLength()) * int32(int32(8)))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagEventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagEventTypeParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, actualLength uint32) (*BACnetContextTagEventType, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagEventType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseValidationError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Manual Field (eventType)
	eventType, _eventTypeErr := ReadEventType(readBuffer, actualLength)
	if _eventTypeErr != nil {
		return nil, errors.Wrap(_eventTypeErr, "Error parsing 'eventType' field")
	}

	// Manual Field (proprietaryValue)
	proprietaryValue, _proprietaryValueErr := ReadProprietaryEventType(readBuffer, eventType, actualLength)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}

	// Virtual field
	_isProprietary := bool((eventType) == (BACnetEventType_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	if closeErr := readBuffer.CloseContext("BACnetContextTagEventType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagEventType{
		EventType:        eventType,
		ProprietaryValue: proprietaryValue,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child, nil
}

func (m *BACnetContextTagEventType) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEventType"); pushErr != nil {
			return pushErr
		}

		// Manual Field (eventType)
		_eventTypeErr := WriteEventType(writeBuffer, m.GetEventType())
		if _eventTypeErr != nil {
			return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
		}

		// Manual Field (proprietaryValue)
		_proprietaryValueErr := WriteProprietaryEventType(writeBuffer, m.GetEventType(), m.GetProprietaryValue())
		if _proprietaryValueErr != nil {
			return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
		}
		// Virtual field
		if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
			return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEventType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
