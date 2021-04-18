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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

type BACnetNotifyType uint8

type IBACnetNotifyType interface {
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

const (
	BACnetNotifyType_ALARM            BACnetNotifyType = 0x0
	BACnetNotifyType_EVENT            BACnetNotifyType = 0x1
	BACnetNotifyType_ACK_NOTIFICATION BACnetNotifyType = 0x2
)

var BACnetNotifyTypeValues []BACnetNotifyType

func init() {
	BACnetNotifyTypeValues = []BACnetNotifyType{
		BACnetNotifyType_ALARM,
		BACnetNotifyType_EVENT,
		BACnetNotifyType_ACK_NOTIFICATION,
	}
}

func BACnetNotifyTypeByValue(value uint8) BACnetNotifyType {
	switch value {
	case 0x0:
		return BACnetNotifyType_ALARM
	case 0x1:
		return BACnetNotifyType_EVENT
	case 0x2:
		return BACnetNotifyType_ACK_NOTIFICATION
	}
	return 0
}

func BACnetNotifyTypeByName(value string) BACnetNotifyType {
	switch value {
	case "ALARM":
		return BACnetNotifyType_ALARM
	case "EVENT":
		return BACnetNotifyType_EVENT
	case "ACK_NOTIFICATION":
		return BACnetNotifyType_ACK_NOTIFICATION
	}
	return 0
}

func CastBACnetNotifyType(structType interface{}) BACnetNotifyType {
	castFunc := func(typ interface{}) BACnetNotifyType {
		if sBACnetNotifyType, ok := typ.(BACnetNotifyType); ok {
			return sBACnetNotifyType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNotifyType) LengthInBits() uint16 {
	return 4
}

func (m BACnetNotifyType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetNotifyTypeParse(io *utils.ReadBuffer) (BACnetNotifyType, error) {
	val, err := io.ReadUint8(4)
	if err != nil {
		return 0, nil
	}
	return BACnetNotifyTypeByValue(val), nil
}

func (e BACnetNotifyType) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint8(4, uint8(e))
	return err
}

func (m *BACnetNotifyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			tok := token.(xml.CharData)
			*m = BACnetNotifyTypeByName(string(tok))
		}
	}
}

func (m BACnetNotifyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e BACnetNotifyType) name() string {
	switch e {
	case BACnetNotifyType_ALARM:
		return "ALARM"
	case BACnetNotifyType_EVENT:
		return "EVENT"
	case BACnetNotifyType_ACK_NOTIFICATION:
		return "ACK_NOTIFICATION"
	}
	return ""
}

func (e BACnetNotifyType) String() string {
	return e.name()
}

func (m BACnetNotifyType) Box(s string, i int) utils.AsciiBox {
	boxName := "BACnetNotifyType"
	if s != "" {
		boxName += "/" + s
	}
	return utils.BoxString(boxName, fmt.Sprintf("%x %s", uint8(m), m.name()), -1)
}
