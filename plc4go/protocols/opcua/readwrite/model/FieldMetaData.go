/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// FieldMetaData is the corresponding interface of FieldMetaData
type FieldMetaData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetFieldFlags returns FieldFlags (property field)
	GetFieldFlags() DataSetFieldFlags
	// GetBuiltInType returns BuiltInType (property field)
	GetBuiltInType() uint8
	// GetDataType returns DataType (property field)
	GetDataType() NodeId
	// GetValueRank returns ValueRank (property field)
	GetValueRank() int32
	// GetNoOfArrayDimensions returns NoOfArrayDimensions (property field)
	GetNoOfArrayDimensions() int32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() []uint32
	// GetMaxStringLength returns MaxStringLength (property field)
	GetMaxStringLength() uint32
	// GetDataSetFieldId returns DataSetFieldId (property field)
	GetDataSetFieldId() GuidValue
	// GetNoOfProperties returns NoOfProperties (property field)
	GetNoOfProperties() int32
	// GetProperties returns Properties (property field)
	GetProperties() []ExtensionObjectDefinition
}

// FieldMetaDataExactly can be used when we want exactly this type and not a type which fulfills FieldMetaData.
// This is useful for switch cases.
type FieldMetaDataExactly interface {
	FieldMetaData
	isFieldMetaData() bool
}

// _FieldMetaData is the data-structure of this message
type _FieldMetaData struct {
	*_ExtensionObjectDefinition
	Name                PascalString
	Description         LocalizedText
	FieldFlags          DataSetFieldFlags
	BuiltInType         uint8
	DataType            NodeId
	ValueRank           int32
	NoOfArrayDimensions int32
	ArrayDimensions     []uint32
	MaxStringLength     uint32
	DataSetFieldId      GuidValue
	NoOfProperties      int32
	Properties          []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FieldMetaData) GetIdentifier() string {
	return "14526"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FieldMetaData) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_FieldMetaData) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FieldMetaData) GetName() PascalString {
	return m.Name
}

func (m *_FieldMetaData) GetDescription() LocalizedText {
	return m.Description
}

func (m *_FieldMetaData) GetFieldFlags() DataSetFieldFlags {
	return m.FieldFlags
}

func (m *_FieldMetaData) GetBuiltInType() uint8 {
	return m.BuiltInType
}

func (m *_FieldMetaData) GetDataType() NodeId {
	return m.DataType
}

func (m *_FieldMetaData) GetValueRank() int32 {
	return m.ValueRank
}

func (m *_FieldMetaData) GetNoOfArrayDimensions() int32 {
	return m.NoOfArrayDimensions
}

func (m *_FieldMetaData) GetArrayDimensions() []uint32 {
	return m.ArrayDimensions
}

func (m *_FieldMetaData) GetMaxStringLength() uint32 {
	return m.MaxStringLength
}

func (m *_FieldMetaData) GetDataSetFieldId() GuidValue {
	return m.DataSetFieldId
}

func (m *_FieldMetaData) GetNoOfProperties() int32 {
	return m.NoOfProperties
}

func (m *_FieldMetaData) GetProperties() []ExtensionObjectDefinition {
	return m.Properties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFieldMetaData factory function for _FieldMetaData
func NewFieldMetaData(name PascalString, description LocalizedText, fieldFlags DataSetFieldFlags, builtInType uint8, dataType NodeId, valueRank int32, noOfArrayDimensions int32, arrayDimensions []uint32, maxStringLength uint32, dataSetFieldId GuidValue, noOfProperties int32, properties []ExtensionObjectDefinition) *_FieldMetaData {
	_result := &_FieldMetaData{
		Name:                       name,
		Description:                description,
		FieldFlags:                 fieldFlags,
		BuiltInType:                builtInType,
		DataType:                   dataType,
		ValueRank:                  valueRank,
		NoOfArrayDimensions:        noOfArrayDimensions,
		ArrayDimensions:            arrayDimensions,
		MaxStringLength:            maxStringLength,
		DataSetFieldId:             dataSetFieldId,
		NoOfProperties:             noOfProperties,
		Properties:                 properties,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFieldMetaData(structType any) FieldMetaData {
	if casted, ok := structType.(FieldMetaData); ok {
		return casted
	}
	if casted, ok := structType.(*FieldMetaData); ok {
		return *casted
	}
	return nil
}

func (m *_FieldMetaData) GetTypeName() string {
	return "FieldMetaData"
}

func (m *_FieldMetaData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (fieldFlags)
	lengthInBits += 16

	// Simple field (builtInType)
	lengthInBits += 8

	// Simple field (dataType)
	lengthInBits += m.DataType.GetLengthInBits(ctx)

	// Simple field (valueRank)
	lengthInBits += 32

	// Simple field (noOfArrayDimensions)
	lengthInBits += 32

	// Array field
	if len(m.ArrayDimensions) > 0 {
		lengthInBits += 32 * uint16(len(m.ArrayDimensions))
	}

	// Simple field (maxStringLength)
	lengthInBits += 32

	// Simple field (dataSetFieldId)
	lengthInBits += m.DataSetFieldId.GetLengthInBits(ctx)

	// Simple field (noOfProperties)
	lengthInBits += 32

	// Array field
	if len(m.Properties) > 0 {
		for _curItem, element := range m.Properties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Properties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FieldMetaData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FieldMetaDataParse(ctx context.Context, theBytes []byte, identifier string) (FieldMetaData, error) {
	return FieldMetaDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func FieldMetaDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (FieldMetaData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("FieldMetaData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FieldMetaData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of FieldMetaData")
	}
	name := _name.(PascalString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	// Simple Field (description)
	if pullErr := readBuffer.PullContext("description"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for description")
	}
	_description, _descriptionErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _descriptionErr != nil {
		return nil, errors.Wrap(_descriptionErr, "Error parsing 'description' field of FieldMetaData")
	}
	description := _description.(LocalizedText)
	if closeErr := readBuffer.CloseContext("description"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for description")
	}

	// Simple Field (fieldFlags)
	if pullErr := readBuffer.PullContext("fieldFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fieldFlags")
	}
	_fieldFlags, _fieldFlagsErr := DataSetFieldFlagsParseWithBuffer(ctx, readBuffer)
	if _fieldFlagsErr != nil {
		return nil, errors.Wrap(_fieldFlagsErr, "Error parsing 'fieldFlags' field of FieldMetaData")
	}
	fieldFlags := _fieldFlags
	if closeErr := readBuffer.CloseContext("fieldFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fieldFlags")
	}

	// Simple Field (builtInType)
	_builtInType, _builtInTypeErr := readBuffer.ReadUint8("builtInType", 8)
	if _builtInTypeErr != nil {
		return nil, errors.Wrap(_builtInTypeErr, "Error parsing 'builtInType' field of FieldMetaData")
	}
	builtInType := _builtInType

	// Simple Field (dataType)
	if pullErr := readBuffer.PullContext("dataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataType")
	}
	_dataType, _dataTypeErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field of FieldMetaData")
	}
	dataType := _dataType.(NodeId)
	if closeErr := readBuffer.CloseContext("dataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataType")
	}

	// Simple Field (valueRank)
	_valueRank, _valueRankErr := readBuffer.ReadInt32("valueRank", 32)
	if _valueRankErr != nil {
		return nil, errors.Wrap(_valueRankErr, "Error parsing 'valueRank' field of FieldMetaData")
	}
	valueRank := _valueRank

	// Simple Field (noOfArrayDimensions)
	_noOfArrayDimensions, _noOfArrayDimensionsErr := readBuffer.ReadInt32("noOfArrayDimensions", 32)
	if _noOfArrayDimensionsErr != nil {
		return nil, errors.Wrap(_noOfArrayDimensionsErr, "Error parsing 'noOfArrayDimensions' field of FieldMetaData")
	}
	noOfArrayDimensions := _noOfArrayDimensions

	// Array field (arrayDimensions)
	if pullErr := readBuffer.PullContext("arrayDimensions", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for arrayDimensions")
	}
	// Count array
	arrayDimensions := make([]uint32, max(noOfArrayDimensions, 0))
	// This happens when the size is set conditional to 0
	if len(arrayDimensions) == 0 {
		arrayDimensions = nil
	}
	{
		_numItems := uint16(max(noOfArrayDimensions, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint32("", 32)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'arrayDimensions' field of FieldMetaData")
			}
			arrayDimensions[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("arrayDimensions", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for arrayDimensions")
	}

	// Simple Field (maxStringLength)
	_maxStringLength, _maxStringLengthErr := readBuffer.ReadUint32("maxStringLength", 32)
	if _maxStringLengthErr != nil {
		return nil, errors.Wrap(_maxStringLengthErr, "Error parsing 'maxStringLength' field of FieldMetaData")
	}
	maxStringLength := _maxStringLength

	// Simple Field (dataSetFieldId)
	if pullErr := readBuffer.PullContext("dataSetFieldId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataSetFieldId")
	}
	_dataSetFieldId, _dataSetFieldIdErr := GuidValueParseWithBuffer(ctx, readBuffer)
	if _dataSetFieldIdErr != nil {
		return nil, errors.Wrap(_dataSetFieldIdErr, "Error parsing 'dataSetFieldId' field of FieldMetaData")
	}
	dataSetFieldId := _dataSetFieldId.(GuidValue)
	if closeErr := readBuffer.CloseContext("dataSetFieldId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataSetFieldId")
	}

	// Simple Field (noOfProperties)
	_noOfProperties, _noOfPropertiesErr := readBuffer.ReadInt32("noOfProperties", 32)
	if _noOfPropertiesErr != nil {
		return nil, errors.Wrap(_noOfPropertiesErr, "Error parsing 'noOfProperties' field of FieldMetaData")
	}
	noOfProperties := _noOfProperties

	// Array field (properties)
	if pullErr := readBuffer.PullContext("properties", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for properties")
	}
	// Count array
	properties := make([]ExtensionObjectDefinition, max(noOfProperties, 0))
	// This happens when the size is set conditional to 0
	if len(properties) == 0 {
		properties = nil
	}
	{
		_numItems := uint16(max(noOfProperties, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "14535")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'properties' field of FieldMetaData")
			}
			properties[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("properties", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for properties")
	}

	if closeErr := readBuffer.CloseContext("FieldMetaData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FieldMetaData")
	}

	// Create a partially initialized instance
	_child := &_FieldMetaData{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Name:                       name,
		Description:                description,
		FieldFlags:                 fieldFlags,
		BuiltInType:                builtInType,
		DataType:                   dataType,
		ValueRank:                  valueRank,
		NoOfArrayDimensions:        noOfArrayDimensions,
		ArrayDimensions:            arrayDimensions,
		MaxStringLength:            maxStringLength,
		DataSetFieldId:             dataSetFieldId,
		NoOfProperties:             noOfProperties,
		Properties:                 properties,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_FieldMetaData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FieldMetaData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FieldMetaData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FieldMetaData")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		// Simple Field (description)
		if pushErr := writeBuffer.PushContext("description"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for description")
		}
		_descriptionErr := writeBuffer.WriteSerializable(ctx, m.GetDescription())
		if popErr := writeBuffer.PopContext("description"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for description")
		}
		if _descriptionErr != nil {
			return errors.Wrap(_descriptionErr, "Error serializing 'description' field")
		}

		// Simple Field (fieldFlags)
		if pushErr := writeBuffer.PushContext("fieldFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fieldFlags")
		}
		_fieldFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetFieldFlags())
		if popErr := writeBuffer.PopContext("fieldFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fieldFlags")
		}
		if _fieldFlagsErr != nil {
			return errors.Wrap(_fieldFlagsErr, "Error serializing 'fieldFlags' field")
		}

		// Simple Field (builtInType)
		builtInType := uint8(m.GetBuiltInType())
		_builtInTypeErr := writeBuffer.WriteUint8("builtInType", 8, uint8((builtInType)))
		if _builtInTypeErr != nil {
			return errors.Wrap(_builtInTypeErr, "Error serializing 'builtInType' field")
		}

		// Simple Field (dataType)
		if pushErr := writeBuffer.PushContext("dataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataType")
		}
		_dataTypeErr := writeBuffer.WriteSerializable(ctx, m.GetDataType())
		if popErr := writeBuffer.PopContext("dataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataType")
		}
		if _dataTypeErr != nil {
			return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
		}

		// Simple Field (valueRank)
		valueRank := int32(m.GetValueRank())
		_valueRankErr := writeBuffer.WriteInt32("valueRank", 32, int32((valueRank)))
		if _valueRankErr != nil {
			return errors.Wrap(_valueRankErr, "Error serializing 'valueRank' field")
		}

		// Simple Field (noOfArrayDimensions)
		noOfArrayDimensions := int32(m.GetNoOfArrayDimensions())
		_noOfArrayDimensionsErr := writeBuffer.WriteInt32("noOfArrayDimensions", 32, int32((noOfArrayDimensions)))
		if _noOfArrayDimensionsErr != nil {
			return errors.Wrap(_noOfArrayDimensionsErr, "Error serializing 'noOfArrayDimensions' field")
		}

		// Array Field (arrayDimensions)
		if pushErr := writeBuffer.PushContext("arrayDimensions", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayDimensions")
		}
		for _curItem, _element := range m.GetArrayDimensions() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint32("", 32, uint32(_element))
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'arrayDimensions' field")
			}
		}
		if popErr := writeBuffer.PopContext("arrayDimensions", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayDimensions")
		}

		// Simple Field (maxStringLength)
		maxStringLength := uint32(m.GetMaxStringLength())
		_maxStringLengthErr := writeBuffer.WriteUint32("maxStringLength", 32, uint32((maxStringLength)))
		if _maxStringLengthErr != nil {
			return errors.Wrap(_maxStringLengthErr, "Error serializing 'maxStringLength' field")
		}

		// Simple Field (dataSetFieldId)
		if pushErr := writeBuffer.PushContext("dataSetFieldId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataSetFieldId")
		}
		_dataSetFieldIdErr := writeBuffer.WriteSerializable(ctx, m.GetDataSetFieldId())
		if popErr := writeBuffer.PopContext("dataSetFieldId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataSetFieldId")
		}
		if _dataSetFieldIdErr != nil {
			return errors.Wrap(_dataSetFieldIdErr, "Error serializing 'dataSetFieldId' field")
		}

		// Simple Field (noOfProperties)
		noOfProperties := int32(m.GetNoOfProperties())
		_noOfPropertiesErr := writeBuffer.WriteInt32("noOfProperties", 32, int32((noOfProperties)))
		if _noOfPropertiesErr != nil {
			return errors.Wrap(_noOfPropertiesErr, "Error serializing 'noOfProperties' field")
		}

		// Array Field (properties)
		if pushErr := writeBuffer.PushContext("properties", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for properties")
		}
		for _curItem, _element := range m.GetProperties() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetProperties()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'properties' field")
			}
		}
		if popErr := writeBuffer.PopContext("properties", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for properties")
		}

		if popErr := writeBuffer.PopContext("FieldMetaData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FieldMetaData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FieldMetaData) isFieldMetaData() bool {
	return true
}

func (m *_FieldMetaData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
