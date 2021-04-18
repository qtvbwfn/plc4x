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

type DeviceDescriptor uint16

type IDeviceDescriptor interface {
	FirmwareType() FirmwareType
	MediumType() DeviceDescriptorMediumType
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

const (
	DeviceDescriptor_TP1_BCU_1_SYSTEM_1_0       DeviceDescriptor = 0x0010
	DeviceDescriptor_TP1_BCU_1_SYSTEM_1_1       DeviceDescriptor = 0x0011
	DeviceDescriptor_TP1_BCU_1_SYSTEM_1_2       DeviceDescriptor = 0x0012
	DeviceDescriptor_TP1_BCU_1_SYSTEM_1_3       DeviceDescriptor = 0x0013
	DeviceDescriptor_TP1_BCU_2_SYSTEM_2_0       DeviceDescriptor = 0x0020
	DeviceDescriptor_TP1_BCU_2_SYSTEM_2_1       DeviceDescriptor = 0x0021
	DeviceDescriptor_TP1_BCU_2_SYSTEM_2_5       DeviceDescriptor = 0x0025
	DeviceDescriptor_TP1_SYSTEM_300             DeviceDescriptor = 0x0300
	DeviceDescriptor_TP1_BIM_M112_0             DeviceDescriptor = 0x0700
	DeviceDescriptor_TP1_BIM_M112_1             DeviceDescriptor = 0x0701
	DeviceDescriptor_TP1_BIM_M112_5             DeviceDescriptor = 0x0705
	DeviceDescriptor_TP1_SYSTEM_B               DeviceDescriptor = 0x07B0
	DeviceDescriptor_TP1_IR_DECODER_0           DeviceDescriptor = 0x0810
	DeviceDescriptor_TP1_IR_DECODER_1           DeviceDescriptor = 0x0811
	DeviceDescriptor_TP1_COUPLER_0              DeviceDescriptor = 0x0910
	DeviceDescriptor_TP1_COUPLER_1              DeviceDescriptor = 0x0911
	DeviceDescriptor_TP1_COUPLER_2              DeviceDescriptor = 0x0912
	DeviceDescriptor_TP1_KNXNETIP_ROUTER        DeviceDescriptor = 0x091A
	DeviceDescriptor_TP1_NONE_D                 DeviceDescriptor = 0x0AFD
	DeviceDescriptor_TP1_NONE_E                 DeviceDescriptor = 0x0AFE
	DeviceDescriptor_PL110_BCU_1_2              DeviceDescriptor = 0x1012
	DeviceDescriptor_PL110_BCU_1_3              DeviceDescriptor = 0x1013
	DeviceDescriptor_PL110_SYSTEM_B             DeviceDescriptor = 0x17B0
	DeviceDescriptor_PL110_MEDIA_COUPLER_PL_TP  DeviceDescriptor = 0x1900
	DeviceDescriptor_RF_BI_DIRECTIONAL_DEVICES  DeviceDescriptor = 0x2010
	DeviceDescriptor_RF_UNI_DIRECTIONAL_DEVICES DeviceDescriptor = 0x2110
	DeviceDescriptor_TP0_BCU_1                  DeviceDescriptor = 0x3012
	DeviceDescriptor_PL132_BCU_1                DeviceDescriptor = 0x4012
	DeviceDescriptor_KNX_IP_SYSTEM7             DeviceDescriptor = 0x5705
)

var DeviceDescriptorValues []DeviceDescriptor

func init() {
	DeviceDescriptorValues = []DeviceDescriptor{
		DeviceDescriptor_TP1_BCU_1_SYSTEM_1_0,
		DeviceDescriptor_TP1_BCU_1_SYSTEM_1_1,
		DeviceDescriptor_TP1_BCU_1_SYSTEM_1_2,
		DeviceDescriptor_TP1_BCU_1_SYSTEM_1_3,
		DeviceDescriptor_TP1_BCU_2_SYSTEM_2_0,
		DeviceDescriptor_TP1_BCU_2_SYSTEM_2_1,
		DeviceDescriptor_TP1_BCU_2_SYSTEM_2_5,
		DeviceDescriptor_TP1_SYSTEM_300,
		DeviceDescriptor_TP1_BIM_M112_0,
		DeviceDescriptor_TP1_BIM_M112_1,
		DeviceDescriptor_TP1_BIM_M112_5,
		DeviceDescriptor_TP1_SYSTEM_B,
		DeviceDescriptor_TP1_IR_DECODER_0,
		DeviceDescriptor_TP1_IR_DECODER_1,
		DeviceDescriptor_TP1_COUPLER_0,
		DeviceDescriptor_TP1_COUPLER_1,
		DeviceDescriptor_TP1_COUPLER_2,
		DeviceDescriptor_TP1_KNXNETIP_ROUTER,
		DeviceDescriptor_TP1_NONE_D,
		DeviceDescriptor_TP1_NONE_E,
		DeviceDescriptor_PL110_BCU_1_2,
		DeviceDescriptor_PL110_BCU_1_3,
		DeviceDescriptor_PL110_SYSTEM_B,
		DeviceDescriptor_PL110_MEDIA_COUPLER_PL_TP,
		DeviceDescriptor_RF_BI_DIRECTIONAL_DEVICES,
		DeviceDescriptor_RF_UNI_DIRECTIONAL_DEVICES,
		DeviceDescriptor_TP0_BCU_1,
		DeviceDescriptor_PL132_BCU_1,
		DeviceDescriptor_KNX_IP_SYSTEM7,
	}
}

func (e DeviceDescriptor) FirmwareType() FirmwareType {
	switch e {
	case 0x0010:
		{ /* '0x0010' */
			return FirmwareType_SYSTEM_1
		}
	case 0x0011:
		{ /* '0x0011' */
			return FirmwareType_SYSTEM_1
		}
	case 0x0012:
		{ /* '0x0012' */
			return FirmwareType_SYSTEM_1
		}
	case 0x0013:
		{ /* '0x0013' */
			return FirmwareType_SYSTEM_1
		}
	case 0x0020:
		{ /* '0x0020' */
			return FirmwareType_SYSTEM_2
		}
	case 0x0021:
		{ /* '0x0021' */
			return FirmwareType_SYSTEM_2
		}
	case 0x0025:
		{ /* '0x0025' */
			return FirmwareType_SYSTEM_2
		}
	case 0x0300:
		{ /* '0x0300' */
			return FirmwareType_SYSTEM_300
		}
	case 0x0700:
		{ /* '0x0700' */
			return FirmwareType_SYSTEM_7
		}
	case 0x0701:
		{ /* '0x0701' */
			return FirmwareType_SYSTEM_7
		}
	case 0x0705:
		{ /* '0x0705' */
			return FirmwareType_SYSTEM_7
		}
	case 0x07B0:
		{ /* '0x07B0' */
			return FirmwareType_SYSTEM_B
		}
	case 0x0810:
		{ /* '0x0810' */
			return FirmwareType_IR_DECODER
		}
	case 0x0811:
		{ /* '0x0811' */
			return FirmwareType_IR_DECODER
		}
	case 0x0910:
		{ /* '0x0910' */
			return FirmwareType_COUPLER
		}
	case 0x0911:
		{ /* '0x0911' */
			return FirmwareType_COUPLER
		}
	case 0x0912:
		{ /* '0x0912' */
			return FirmwareType_COUPLER
		}
	case 0x091A:
		{ /* '0x091A' */
			return FirmwareType_COUPLER
		}
	case 0x0AFD:
		{ /* '0x0AFD' */
			return FirmwareType_NONE
		}
	case 0x0AFE:
		{ /* '0x0AFE' */
			return FirmwareType_NONE
		}
	case 0x1012:
		{ /* '0x1012' */
			return FirmwareType_SYSTEM_1
		}
	case 0x1013:
		{ /* '0x1013' */
			return FirmwareType_SYSTEM_1
		}
	case 0x17B0:
		{ /* '0x17B0' */
			return FirmwareType_SYSTEM_B
		}
	case 0x1900:
		{ /* '0x1900' */
			return FirmwareType_MEDIA_COUPLER_PL_TP
		}
	case 0x2010:
		{ /* '0x2010' */
			return FirmwareType_RF_BI_DIRECTIONAL_DEVICES
		}
	case 0x2110:
		{ /* '0x2110' */
			return FirmwareType_RF_UNI_DIRECTIONAL_DEVICES
		}
	case 0x3012:
		{ /* '0x3012' */
			return FirmwareType_SYSTEM_1
		}
	case 0x4012:
		{ /* '0x4012' */
			return FirmwareType_SYSTEM_1
		}
	case 0x5705:
		{ /* '0x5705' */
			return FirmwareType_SYSTEM_7
		}
	default:
		{
			return 0
		}
	}
}

func (e DeviceDescriptor) MediumType() DeviceDescriptorMediumType {
	switch e {
	case 0x0010:
		{ /* '0x0010' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0011:
		{ /* '0x0011' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0012:
		{ /* '0x0012' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0013:
		{ /* '0x0013' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0020:
		{ /* '0x0020' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0021:
		{ /* '0x0021' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0025:
		{ /* '0x0025' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0300:
		{ /* '0x0300' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0700:
		{ /* '0x0700' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0701:
		{ /* '0x0701' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0705:
		{ /* '0x0705' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x07B0:
		{ /* '0x07B0' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0810:
		{ /* '0x0810' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0811:
		{ /* '0x0811' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0910:
		{ /* '0x0910' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0911:
		{ /* '0x0911' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0912:
		{ /* '0x0912' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x091A:
		{ /* '0x091A' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0AFD:
		{ /* '0x0AFD' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x0AFE:
		{ /* '0x0AFE' */
			return DeviceDescriptorMediumType_TP1
		}
	case 0x1012:
		{ /* '0x1012' */
			return DeviceDescriptorMediumType_PL110
		}
	case 0x1013:
		{ /* '0x1013' */
			return DeviceDescriptorMediumType_PL110
		}
	case 0x17B0:
		{ /* '0x17B0' */
			return DeviceDescriptorMediumType_PL110
		}
	case 0x1900:
		{ /* '0x1900' */
			return DeviceDescriptorMediumType_PL110
		}
	case 0x2010:
		{ /* '0x2010' */
			return DeviceDescriptorMediumType_RF
		}
	case 0x2110:
		{ /* '0x2110' */
			return DeviceDescriptorMediumType_RF
		}
	case 0x3012:
		{ /* '0x3012' */
			return DeviceDescriptorMediumType_TP0
		}
	case 0x4012:
		{ /* '0x4012' */
			return DeviceDescriptorMediumType_PL132
		}
	case 0x5705:
		{ /* '0x5705' */
			return DeviceDescriptorMediumType_KNX_IP
		}
	default:
		{
			return 0
		}
	}
}
func DeviceDescriptorByValue(value uint16) DeviceDescriptor {
	switch value {
	case 0x0010:
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_0
	case 0x0011:
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_1
	case 0x0012:
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_2
	case 0x0013:
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_3
	case 0x0020:
		return DeviceDescriptor_TP1_BCU_2_SYSTEM_2_0
	case 0x0021:
		return DeviceDescriptor_TP1_BCU_2_SYSTEM_2_1
	case 0x0025:
		return DeviceDescriptor_TP1_BCU_2_SYSTEM_2_5
	case 0x0300:
		return DeviceDescriptor_TP1_SYSTEM_300
	case 0x0700:
		return DeviceDescriptor_TP1_BIM_M112_0
	case 0x0701:
		return DeviceDescriptor_TP1_BIM_M112_1
	case 0x0705:
		return DeviceDescriptor_TP1_BIM_M112_5
	case 0x07B0:
		return DeviceDescriptor_TP1_SYSTEM_B
	case 0x0810:
		return DeviceDescriptor_TP1_IR_DECODER_0
	case 0x0811:
		return DeviceDescriptor_TP1_IR_DECODER_1
	case 0x0910:
		return DeviceDescriptor_TP1_COUPLER_0
	case 0x0911:
		return DeviceDescriptor_TP1_COUPLER_1
	case 0x0912:
		return DeviceDescriptor_TP1_COUPLER_2
	case 0x091A:
		return DeviceDescriptor_TP1_KNXNETIP_ROUTER
	case 0x0AFD:
		return DeviceDescriptor_TP1_NONE_D
	case 0x0AFE:
		return DeviceDescriptor_TP1_NONE_E
	case 0x1012:
		return DeviceDescriptor_PL110_BCU_1_2
	case 0x1013:
		return DeviceDescriptor_PL110_BCU_1_3
	case 0x17B0:
		return DeviceDescriptor_PL110_SYSTEM_B
	case 0x1900:
		return DeviceDescriptor_PL110_MEDIA_COUPLER_PL_TP
	case 0x2010:
		return DeviceDescriptor_RF_BI_DIRECTIONAL_DEVICES
	case 0x2110:
		return DeviceDescriptor_RF_UNI_DIRECTIONAL_DEVICES
	case 0x3012:
		return DeviceDescriptor_TP0_BCU_1
	case 0x4012:
		return DeviceDescriptor_PL132_BCU_1
	case 0x5705:
		return DeviceDescriptor_KNX_IP_SYSTEM7
	}
	return 0
}

func DeviceDescriptorByName(value string) DeviceDescriptor {
	switch value {
	case "TP1_BCU_1_SYSTEM_1_0":
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_0
	case "TP1_BCU_1_SYSTEM_1_1":
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_1
	case "TP1_BCU_1_SYSTEM_1_2":
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_2
	case "TP1_BCU_1_SYSTEM_1_3":
		return DeviceDescriptor_TP1_BCU_1_SYSTEM_1_3
	case "TP1_BCU_2_SYSTEM_2_0":
		return DeviceDescriptor_TP1_BCU_2_SYSTEM_2_0
	case "TP1_BCU_2_SYSTEM_2_1":
		return DeviceDescriptor_TP1_BCU_2_SYSTEM_2_1
	case "TP1_BCU_2_SYSTEM_2_5":
		return DeviceDescriptor_TP1_BCU_2_SYSTEM_2_5
	case "TP1_SYSTEM_300":
		return DeviceDescriptor_TP1_SYSTEM_300
	case "TP1_BIM_M112_0":
		return DeviceDescriptor_TP1_BIM_M112_0
	case "TP1_BIM_M112_1":
		return DeviceDescriptor_TP1_BIM_M112_1
	case "TP1_BIM_M112_5":
		return DeviceDescriptor_TP1_BIM_M112_5
	case "TP1_SYSTEM_B":
		return DeviceDescriptor_TP1_SYSTEM_B
	case "TP1_IR_DECODER_0":
		return DeviceDescriptor_TP1_IR_DECODER_0
	case "TP1_IR_DECODER_1":
		return DeviceDescriptor_TP1_IR_DECODER_1
	case "TP1_COUPLER_0":
		return DeviceDescriptor_TP1_COUPLER_0
	case "TP1_COUPLER_1":
		return DeviceDescriptor_TP1_COUPLER_1
	case "TP1_COUPLER_2":
		return DeviceDescriptor_TP1_COUPLER_2
	case "TP1_KNXNETIP_ROUTER":
		return DeviceDescriptor_TP1_KNXNETIP_ROUTER
	case "TP1_NONE_D":
		return DeviceDescriptor_TP1_NONE_D
	case "TP1_NONE_E":
		return DeviceDescriptor_TP1_NONE_E
	case "PL110_BCU_1_2":
		return DeviceDescriptor_PL110_BCU_1_2
	case "PL110_BCU_1_3":
		return DeviceDescriptor_PL110_BCU_1_3
	case "PL110_SYSTEM_B":
		return DeviceDescriptor_PL110_SYSTEM_B
	case "PL110_MEDIA_COUPLER_PL_TP":
		return DeviceDescriptor_PL110_MEDIA_COUPLER_PL_TP
	case "RF_BI_DIRECTIONAL_DEVICES":
		return DeviceDescriptor_RF_BI_DIRECTIONAL_DEVICES
	case "RF_UNI_DIRECTIONAL_DEVICES":
		return DeviceDescriptor_RF_UNI_DIRECTIONAL_DEVICES
	case "TP0_BCU_1":
		return DeviceDescriptor_TP0_BCU_1
	case "PL132_BCU_1":
		return DeviceDescriptor_PL132_BCU_1
	case "KNX_IP_SYSTEM7":
		return DeviceDescriptor_KNX_IP_SYSTEM7
	}
	return 0
}

func CastDeviceDescriptor(structType interface{}) DeviceDescriptor {
	castFunc := func(typ interface{}) DeviceDescriptor {
		if sDeviceDescriptor, ok := typ.(DeviceDescriptor); ok {
			return sDeviceDescriptor
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeviceDescriptor) LengthInBits() uint16 {
	return 16
}

func (m DeviceDescriptor) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceDescriptorParse(io *utils.ReadBuffer) (DeviceDescriptor, error) {
	val, err := io.ReadUint16(16)
	if err != nil {
		return 0, nil
	}
	return DeviceDescriptorByValue(val), nil
}

func (e DeviceDescriptor) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint16(16, uint16(e))
	return err
}

func (m *DeviceDescriptor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			*m = DeviceDescriptorByName(string(tok))
		}
	}
}

func (m DeviceDescriptor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e DeviceDescriptor) name() string {
	switch e {
	case DeviceDescriptor_TP1_BCU_1_SYSTEM_1_0:
		return "TP1_BCU_1_SYSTEM_1_0"
	case DeviceDescriptor_TP1_BCU_1_SYSTEM_1_1:
		return "TP1_BCU_1_SYSTEM_1_1"
	case DeviceDescriptor_TP1_BCU_1_SYSTEM_1_2:
		return "TP1_BCU_1_SYSTEM_1_2"
	case DeviceDescriptor_TP1_BCU_1_SYSTEM_1_3:
		return "TP1_BCU_1_SYSTEM_1_3"
	case DeviceDescriptor_TP1_BCU_2_SYSTEM_2_0:
		return "TP1_BCU_2_SYSTEM_2_0"
	case DeviceDescriptor_TP1_BCU_2_SYSTEM_2_1:
		return "TP1_BCU_2_SYSTEM_2_1"
	case DeviceDescriptor_TP1_BCU_2_SYSTEM_2_5:
		return "TP1_BCU_2_SYSTEM_2_5"
	case DeviceDescriptor_TP1_SYSTEM_300:
		return "TP1_SYSTEM_300"
	case DeviceDescriptor_TP1_BIM_M112_0:
		return "TP1_BIM_M112_0"
	case DeviceDescriptor_TP1_BIM_M112_1:
		return "TP1_BIM_M112_1"
	case DeviceDescriptor_TP1_BIM_M112_5:
		return "TP1_BIM_M112_5"
	case DeviceDescriptor_TP1_SYSTEM_B:
		return "TP1_SYSTEM_B"
	case DeviceDescriptor_TP1_IR_DECODER_0:
		return "TP1_IR_DECODER_0"
	case DeviceDescriptor_TP1_IR_DECODER_1:
		return "TP1_IR_DECODER_1"
	case DeviceDescriptor_TP1_COUPLER_0:
		return "TP1_COUPLER_0"
	case DeviceDescriptor_TP1_COUPLER_1:
		return "TP1_COUPLER_1"
	case DeviceDescriptor_TP1_COUPLER_2:
		return "TP1_COUPLER_2"
	case DeviceDescriptor_TP1_KNXNETIP_ROUTER:
		return "TP1_KNXNETIP_ROUTER"
	case DeviceDescriptor_TP1_NONE_D:
		return "TP1_NONE_D"
	case DeviceDescriptor_TP1_NONE_E:
		return "TP1_NONE_E"
	case DeviceDescriptor_PL110_BCU_1_2:
		return "PL110_BCU_1_2"
	case DeviceDescriptor_PL110_BCU_1_3:
		return "PL110_BCU_1_3"
	case DeviceDescriptor_PL110_SYSTEM_B:
		return "PL110_SYSTEM_B"
	case DeviceDescriptor_PL110_MEDIA_COUPLER_PL_TP:
		return "PL110_MEDIA_COUPLER_PL_TP"
	case DeviceDescriptor_RF_BI_DIRECTIONAL_DEVICES:
		return "RF_BI_DIRECTIONAL_DEVICES"
	case DeviceDescriptor_RF_UNI_DIRECTIONAL_DEVICES:
		return "RF_UNI_DIRECTIONAL_DEVICES"
	case DeviceDescriptor_TP0_BCU_1:
		return "TP0_BCU_1"
	case DeviceDescriptor_PL132_BCU_1:
		return "PL132_BCU_1"
	case DeviceDescriptor_KNX_IP_SYSTEM7:
		return "KNX_IP_SYSTEM7"
	}
	return ""
}

func (e DeviceDescriptor) String() string {
	return e.name()
}

func (m DeviceDescriptor) Box(s string, i int) utils.AsciiBox {
	boxName := "DeviceDescriptor"
	if s != "" {
		boxName += "/" + s
	}
	return utils.BoxString(boxName, fmt.Sprintf("%x %s", uint16(m), m.name()), -1)
}
