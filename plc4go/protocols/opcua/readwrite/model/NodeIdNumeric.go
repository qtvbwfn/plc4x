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

// NodeIdNumeric is the corresponding interface of NodeIdNumeric
type NodeIdNumeric interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() uint32
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// NodeIdNumericExactly can be used when we want exactly this type and not a type which fulfills NodeIdNumeric.
// This is useful for switch cases.
type NodeIdNumericExactly interface {
	NodeIdNumeric
	isNodeIdNumeric() bool
}

// _NodeIdNumeric is the data-structure of this message
type _NodeIdNumeric struct {
	*_NodeIdTypeDefinition
	NamespaceIndex uint16
	Id             uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdNumeric) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeNumeric
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdNumeric) InitializeParent(parent NodeIdTypeDefinition) {}

func (m *_NodeIdNumeric) GetParent() NodeIdTypeDefinition {
	return m._NodeIdTypeDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdNumeric) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdNumeric) GetId() uint32 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdNumeric) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdNumeric factory function for _NodeIdNumeric
func NewNodeIdNumeric(namespaceIndex uint16, id uint32) *_NodeIdNumeric {
	_result := &_NodeIdNumeric{
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
		_NodeIdTypeDefinition: NewNodeIdTypeDefinition(),
	}
	_result._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeIdNumeric(structType any) NodeIdNumeric {
	if casted, ok := structType.(NodeIdNumeric); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdNumeric); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdNumeric) GetTypeName() string {
	return "NodeIdNumeric"
}

func (m *_NodeIdNumeric) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (id)
	lengthInBits += 32

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdNumeric) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdNumericParse(ctx context.Context, theBytes []byte) (NodeIdNumeric, error) {
	return NodeIdNumericParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdNumericParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdNumeric, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeIdNumeric"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdNumeric")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint16("namespaceIndex", 16)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of NodeIdNumeric")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (id)
	_id, _idErr := readBuffer.ReadUint32("id", 32)
	if _idErr != nil {
		return nil, errors.Wrap(_idErr, "Error parsing 'id' field of NodeIdNumeric")
	}
	id := _id

	// Virtual field
	_identifier := id
	identifier := fmt.Sprintf("%v", _identifier)
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdNumeric"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdNumeric")
	}

	// Create a partially initialized instance
	_child := &_NodeIdNumeric{
		_NodeIdTypeDefinition: &_NodeIdTypeDefinition{},
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
	}
	_child._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeIdNumeric) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdNumeric) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdNumeric"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdNumeric")
		}

		// Simple Field (namespaceIndex)
		namespaceIndex := uint16(m.GetNamespaceIndex())
		_namespaceIndexErr := writeBuffer.WriteUint16("namespaceIndex", 16, uint16((namespaceIndex)))
		if _namespaceIndexErr != nil {
			return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
		}

		// Simple Field (id)
		id := uint32(m.GetId())
		_idErr := writeBuffer.WriteUint32("id", 32, uint32((id)))
		if _idErr != nil {
			return errors.Wrap(_idErr, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdNumeric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdNumeric")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdNumeric) isNodeIdNumeric() bool {
	return true
}

func (m *_NodeIdNumeric) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
