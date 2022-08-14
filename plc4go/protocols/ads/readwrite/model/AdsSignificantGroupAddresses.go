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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsSignificantGroupAddresses is an enum
type AdsSignificantGroupAddresses uint32

type IAdsSignificantGroupAddresses interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	AdsSignificantGroupAddresses_SYMBOL_TABLE    AdsSignificantGroupAddresses = 0x0000F00B
	AdsSignificantGroupAddresses_DATA_TYPE_TABLE AdsSignificantGroupAddresses = 0x0000F00E
	AdsSignificantGroupAddresses_TABLE_SIZES     AdsSignificantGroupAddresses = 0x0000F00F
)

var AdsSignificantGroupAddressesValues []AdsSignificantGroupAddresses

func init() {
	_ = errors.New
	AdsSignificantGroupAddressesValues = []AdsSignificantGroupAddresses{
		AdsSignificantGroupAddresses_SYMBOL_TABLE,
		AdsSignificantGroupAddresses_DATA_TYPE_TABLE,
		AdsSignificantGroupAddresses_TABLE_SIZES,
	}
}

func AdsSignificantGroupAddressesByValue(value uint32) (enum AdsSignificantGroupAddresses, ok bool) {
	switch value {
	case 0x0000F00B:
		return AdsSignificantGroupAddresses_SYMBOL_TABLE, true
	case 0x0000F00E:
		return AdsSignificantGroupAddresses_DATA_TYPE_TABLE, true
	case 0x0000F00F:
		return AdsSignificantGroupAddresses_TABLE_SIZES, true
	}
	return 0, false
}

func AdsSignificantGroupAddressesByName(value string) (enum AdsSignificantGroupAddresses, ok bool) {
	switch value {
	case "SYMBOL_TABLE":
		return AdsSignificantGroupAddresses_SYMBOL_TABLE, true
	case "DATA_TYPE_TABLE":
		return AdsSignificantGroupAddresses_DATA_TYPE_TABLE, true
	case "TABLE_SIZES":
		return AdsSignificantGroupAddresses_TABLE_SIZES, true
	}
	return 0, false
}

func AdsSignificantGroupAddressesKnows(value uint32) bool {
	for _, typeValue := range AdsSignificantGroupAddressesValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAdsSignificantGroupAddresses(structType interface{}) AdsSignificantGroupAddresses {
	castFunc := func(typ interface{}) AdsSignificantGroupAddresses {
		if sAdsSignificantGroupAddresses, ok := typ.(AdsSignificantGroupAddresses); ok {
			return sAdsSignificantGroupAddresses
		}
		return 0
	}
	return castFunc(structType)
}

func (m AdsSignificantGroupAddresses) GetLengthInBits() uint16 {
	return 32
}

func (m AdsSignificantGroupAddresses) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsSignificantGroupAddressesParse(readBuffer utils.ReadBuffer) (AdsSignificantGroupAddresses, error) {
	val, err := readBuffer.ReadUint32("AdsSignificantGroupAddresses", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AdsSignificantGroupAddresses")
	}
	if enum, ok := AdsSignificantGroupAddressesByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return AdsSignificantGroupAddresses(val), nil
	} else {
		return enum, nil
	}
}

func (e AdsSignificantGroupAddresses) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint32("AdsSignificantGroupAddresses", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AdsSignificantGroupAddresses) PLC4XEnumName() string {
	switch e {
	case AdsSignificantGroupAddresses_SYMBOL_TABLE:
		return "SYMBOL_TABLE"
	case AdsSignificantGroupAddresses_DATA_TYPE_TABLE:
		return "DATA_TYPE_TABLE"
	case AdsSignificantGroupAddresses_TABLE_SIZES:
		return "TABLE_SIZES"
	}
	return ""
}

func (e AdsSignificantGroupAddresses) String() string {
	return e.PLC4XEnumName()
}
