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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataFileType is the data-structure of this message
type BACnetConstructedDataFileType struct {
	*BACnetConstructedData
	FileType *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataFileType is the corresponding interface of BACnetConstructedDataFileType
type IBACnetConstructedDataFileType interface {
	IBACnetConstructedData
	// GetFileType returns FileType (property field)
	GetFileType() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataFileType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataFileType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FILE_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFileType) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFileType) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFileType) GetFileType() *BACnetApplicationTagCharacterString {
	return m.FileType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFileType factory function for BACnetConstructedDataFileType
func NewBACnetConstructedDataFileType(fileType *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataFileType {
	_result := &BACnetConstructedDataFileType{
		FileType:              fileType,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFileType(structType interface{}) *BACnetConstructedDataFileType {
	if casted, ok := structType.(BACnetConstructedDataFileType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFileType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFileType(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFileType(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFileType) GetTypeName() string {
	return "BACnetConstructedDataFileType"
}

func (m *BACnetConstructedDataFileType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFileType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileType)
	lengthInBits += m.FileType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataFileType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFileTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataFileType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFileType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileType)
	if pullErr := readBuffer.PullContext("fileType"); pullErr != nil {
		return nil, pullErr
	}
	_fileType, _fileTypeErr := BACnetApplicationTagParse(readBuffer)
	if _fileTypeErr != nil {
		return nil, errors.Wrap(_fileTypeErr, "Error parsing 'fileType' field")
	}
	fileType := CastBACnetApplicationTagCharacterString(_fileType)
	if closeErr := readBuffer.CloseContext("fileType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFileType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFileType{
		FileType:              CastBACnetApplicationTagCharacterString(fileType),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFileType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFileType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (fileType)
		if pushErr := writeBuffer.PushContext("fileType"); pushErr != nil {
			return pushErr
		}
		_fileTypeErr := m.FileType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileType"); popErr != nil {
			return popErr
		}
		if _fileTypeErr != nil {
			return errors.Wrap(_fileTypeErr, "Error serializing 'fileType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFileType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFileType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
