// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/protocol/lorawan/lorawan.proto
// DO NOT EDIT!

/*
	Package lorawan is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/protocol/lorawan/lorawan.proto

	It has these top-level messages:
		Metadata
		TxConfiguration
		ActivationMetadata
*/
package lorawan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Modulation int32

const (
	Modulation_LORA Modulation = 0
	Modulation_FSK  Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}
var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}
func (Modulation) EnumDescriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{0} }

type Region int32

const (
	Region_EU_863_870 Region = 0
	Region_US_902_928 Region = 1
	Region_CN_779_787 Region = 2
	Region_EU_433     Region = 3
	Region_AU_915_928 Region = 4
	Region_CN_470_510 Region = 5
	Region_AS_923     Region = 6
	Region_SK_920_923 Region = 7
)

var Region_name = map[int32]string{
	0: "EU_863_870",
	1: "US_902_928",
	2: "CN_779_787",
	3: "EU_433",
	4: "AU_915_928",
	5: "CN_470_510",
	6: "AS_923",
	7: "SK_920_923",
}
var Region_value = map[string]int32{
	"EU_863_870": 0,
	"US_902_928": 1,
	"CN_779_787": 2,
	"EU_433":     3,
	"AU_915_928": 4,
	"CN_470_510": 5,
	"AS_923":     6,
	"SK_920_923": 7,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}
func (Region) EnumDescriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{1} }

type Metadata struct {
	Modulation Modulation `protobuf:"varint,11,opt,name=modulation,proto3,enum=lorawan.Modulation" json:"modulation,omitempty"`
	DataRate   string     `protobuf:"bytes,12,opt,name=data_rate,json=dataRate,proto3" json:"data_rate,omitempty"`
	BitRate    uint32     `protobuf:"varint,13,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate,omitempty"`
	CodingRate string     `protobuf:"bytes,14,opt,name=coding_rate,json=codingRate,proto3" json:"coding_rate,omitempty"`
	FCnt       uint32     `protobuf:"varint,15,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{0} }

type TxConfiguration struct {
	Modulation Modulation `protobuf:"varint,11,opt,name=modulation,proto3,enum=lorawan.Modulation" json:"modulation,omitempty"`
	DataRate   string     `protobuf:"bytes,12,opt,name=data_rate,json=dataRate,proto3" json:"data_rate,omitempty"`
	BitRate    uint32     `protobuf:"varint,13,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate,omitempty"`
	CodingRate string     `protobuf:"bytes,14,opt,name=coding_rate,json=codingRate,proto3" json:"coding_rate,omitempty"`
	FCnt       uint32     `protobuf:"varint,15,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (m *TxConfiguration) Reset()                    { *m = TxConfiguration{} }
func (m *TxConfiguration) String() string            { return proto.CompactTextString(m) }
func (*TxConfiguration) ProtoMessage()               {}
func (*TxConfiguration) Descriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{1} }

type ActivationMetadata struct {
	AppEui      *github_com_TheThingsNetwork_ttn_core_types.AppEUI  `protobuf:"bytes,1,opt,name=app_eui,json=appEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.AppEUI" json:"app_eui,omitempty"`
	DevEui      *github_com_TheThingsNetwork_ttn_core_types.DevEUI  `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevEUI" json:"dev_eui,omitempty"`
	DevAddr     *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,3,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
	NwkSKey     *github_com_TheThingsNetwork_ttn_core_types.NwkSKey `protobuf:"bytes,4,opt,name=nwk_s_key,json=nwkSKey,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.NwkSKey" json:"nwk_s_key,omitempty"`
	Rx1DrOffset uint32                                              `protobuf:"varint,11,opt,name=rx1_dr_offset,json=rx1DrOffset,proto3" json:"rx1_dr_offset,omitempty"`
	Rx2Dr       uint32                                              `protobuf:"varint,12,opt,name=rx2_dr,json=rx2Dr,proto3" json:"rx2_dr,omitempty"`
	RxDelay     uint32                                              `protobuf:"varint,13,opt,name=rx_delay,json=rxDelay,proto3" json:"rx_delay,omitempty"`
	CfList      []uint64                                            `protobuf:"varint,14,rep,packed,name=cf_list,json=cfList" json:"cf_list,omitempty"`
}

func (m *ActivationMetadata) Reset()                    { *m = ActivationMetadata{} }
func (m *ActivationMetadata) String() string            { return proto.CompactTextString(m) }
func (*ActivationMetadata) ProtoMessage()               {}
func (*ActivationMetadata) Descriptor() ([]byte, []int) { return fileDescriptorLorawan, []int{2} }

func init() {
	proto.RegisterType((*Metadata)(nil), "lorawan.Metadata")
	proto.RegisterType((*TxConfiguration)(nil), "lorawan.TxConfiguration")
	proto.RegisterType((*ActivationMetadata)(nil), "lorawan.ActivationMetadata")
	proto.RegisterEnum("lorawan.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("lorawan.Region", Region_name, Region_value)
}
func (m *Metadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Metadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Modulation != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Modulation))
	}
	if len(m.DataRate) > 0 {
		data[i] = 0x62
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.DataRate)))
		i += copy(data[i:], m.DataRate)
	}
	if m.BitRate != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintLorawan(data, i, uint64(m.BitRate))
	}
	if len(m.CodingRate) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.CodingRate)))
		i += copy(data[i:], m.CodingRate)
	}
	if m.FCnt != 0 {
		data[i] = 0x78
		i++
		i = encodeVarintLorawan(data, i, uint64(m.FCnt))
	}
	return i, nil
}

func (m *TxConfiguration) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TxConfiguration) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Modulation != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Modulation))
	}
	if len(m.DataRate) > 0 {
		data[i] = 0x62
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.DataRate)))
		i += copy(data[i:], m.DataRate)
	}
	if m.BitRate != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintLorawan(data, i, uint64(m.BitRate))
	}
	if len(m.CodingRate) > 0 {
		data[i] = 0x72
		i++
		i = encodeVarintLorawan(data, i, uint64(len(m.CodingRate)))
		i += copy(data[i:], m.CodingRate)
	}
	if m.FCnt != 0 {
		data[i] = 0x78
		i++
		i = encodeVarintLorawan(data, i, uint64(m.FCnt))
	}
	return i, nil
}

func (m *ActivationMetadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ActivationMetadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppEui != nil {
		data[i] = 0xa
		i++
		i = encodeVarintLorawan(data, i, uint64(m.AppEui.Size()))
		n1, err := m.AppEui.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DevEui != nil {
		data[i] = 0x12
		i++
		i = encodeVarintLorawan(data, i, uint64(m.DevEui.Size()))
		n2, err := m.DevEui.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.DevAddr != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintLorawan(data, i, uint64(m.DevAddr.Size()))
		n3, err := m.DevAddr.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.NwkSKey != nil {
		data[i] = 0x22
		i++
		i = encodeVarintLorawan(data, i, uint64(m.NwkSKey.Size()))
		n4, err := m.NwkSKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Rx1DrOffset != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Rx1DrOffset))
	}
	if m.Rx2Dr != 0 {
		data[i] = 0x60
		i++
		i = encodeVarintLorawan(data, i, uint64(m.Rx2Dr))
	}
	if m.RxDelay != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintLorawan(data, i, uint64(m.RxDelay))
	}
	if len(m.CfList) > 0 {
		data6 := make([]byte, len(m.CfList)*10)
		var j5 int
		for _, num := range m.CfList {
			for num >= 1<<7 {
				data6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			data6[j5] = uint8(num)
			j5++
		}
		data[i] = 0x72
		i++
		i = encodeVarintLorawan(data, i, uint64(j5))
		i += copy(data[i:], data6[:j5])
	}
	return i, nil
}

func encodeFixed64Lorawan(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Lorawan(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLorawan(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if m.Modulation != 0 {
		n += 1 + sovLorawan(uint64(m.Modulation))
	}
	l = len(m.DataRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.BitRate != 0 {
		n += 1 + sovLorawan(uint64(m.BitRate))
	}
	l = len(m.CodingRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.FCnt != 0 {
		n += 1 + sovLorawan(uint64(m.FCnt))
	}
	return n
}

func (m *TxConfiguration) Size() (n int) {
	var l int
	_ = l
	if m.Modulation != 0 {
		n += 1 + sovLorawan(uint64(m.Modulation))
	}
	l = len(m.DataRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.BitRate != 0 {
		n += 1 + sovLorawan(uint64(m.BitRate))
	}
	l = len(m.CodingRate)
	if l > 0 {
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.FCnt != 0 {
		n += 1 + sovLorawan(uint64(m.FCnt))
	}
	return n
}

func (m *ActivationMetadata) Size() (n int) {
	var l int
	_ = l
	if m.AppEui != nil {
		l = m.AppEui.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.DevEui != nil {
		l = m.DevEui.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.NwkSKey != nil {
		l = m.NwkSKey.Size()
		n += 1 + l + sovLorawan(uint64(l))
	}
	if m.Rx1DrOffset != 0 {
		n += 1 + sovLorawan(uint64(m.Rx1DrOffset))
	}
	if m.Rx2Dr != 0 {
		n += 1 + sovLorawan(uint64(m.Rx2Dr))
	}
	if m.RxDelay != 0 {
		n += 1 + sovLorawan(uint64(m.RxDelay))
	}
	if len(m.CfList) > 0 {
		l = 0
		for _, e := range m.CfList {
			l += sovLorawan(uint64(e))
		}
		n += 1 + sovLorawan(uint64(l)) + l
	}
	return n
}

func sovLorawan(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLorawan(x uint64) (n int) {
	return sovLorawan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLorawan
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modulation", wireType)
			}
			m.Modulation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Modulation |= (Modulation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitRate", wireType)
			}
			m.BitRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.BitRate |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodingRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLorawan(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLorawan
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxConfiguration) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLorawan
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TxConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modulation", wireType)
			}
			m.Modulation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Modulation |= (Modulation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitRate", wireType)
			}
			m.BitRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.BitRate |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodingRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLorawan(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLorawan
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ActivationMetadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLorawan
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivationMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivationMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.AppEUI
			m.AppEui = &v
			if err := m.AppEui.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevEUI
			m.DevEui = &v
			if err := m.DevEui.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLorawan
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.NwkSKey
			m.NwkSKey = &v
			if err := m.NwkSKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rx1DrOffset", wireType)
			}
			m.Rx1DrOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Rx1DrOffset |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rx2Dr", wireType)
			}
			m.Rx2Dr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Rx2Dr |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxDelay", wireType)
			}
			m.RxDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RxDelay |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLorawan
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLorawan
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLorawan
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := data[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CfList = append(m.CfList, v)
				}
			} else if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLorawan
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CfList = append(m.CfList, v)
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CfList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLorawan(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLorawan
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLorawan(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLorawan
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLorawan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLorawan
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLorawan
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLorawan(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLorawan = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLorawan   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/protocol/lorawan/lorawan.proto", fileDescriptorLorawan)
}

var fileDescriptorLorawan = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x94, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc7, 0x31, 0x09, 0x76, 0x18, 0x48, 0xb0, 0x8c, 0x7e, 0xfa, 0xb9, 0xad, 0x94, 0x44, 0x9c,
	0x22, 0xa4, 0xe6, 0x8f, 0x03, 0x24, 0x39, 0x06, 0x12, 0xa4, 0x0a, 0x08, 0xaa, 0x43, 0x2e, 0xbd,
	0xac, 0x36, 0xf6, 0xda, 0xac, 0x12, 0xbc, 0xd6, 0x66, 0x93, 0x38, 0xb7, 0x3e, 0x46, 0x5f, 0xa2,
	0xb7, 0x3e, 0x44, 0xd5, 0x53, 0xcf, 0x1c, 0x50, 0x45, 0x5f, 0xa4, 0xda, 0x35, 0x50, 0x6e, 0x55,
	0xb9, 0xf5, 0x94, 0xf9, 0xce, 0x77, 0xe6, 0xb3, 0xa3, 0xcc, 0xc8, 0x70, 0x1c, 0x52, 0x71, 0x3d,
	0x1f, 0x57, 0x3d, 0x76, 0x53, 0xbb, 0xba, 0x26, 0x57, 0xd7, 0x34, 0x0a, 0x67, 0x03, 0x22, 0x96,
	0x8c, 0x4f, 0x6a, 0x42, 0x44, 0x35, 0x1c, 0xd3, 0x5a, 0xcc, 0x99, 0x60, 0x1e, 0x9b, 0xd6, 0xa6,
	0x8c, 0xe3, 0x25, 0x8e, 0x1e, 0x7f, 0xab, 0xca, 0xb0, 0x8c, 0x07, 0xf9, 0xfa, 0xed, 0x33, 0x58,
	0xc8, 0x42, 0x96, 0x36, 0x8e, 0xe7, 0x81, 0x52, 0x4a, 0xa8, 0x28, 0xed, 0xdb, 0xfb, 0xac, 0x41,
	0xee, 0x82, 0x08, 0xec, 0x63, 0x81, 0xad, 0x26, 0xc0, 0x0d, 0xf3, 0xe7, 0x53, 0x2c, 0x28, 0x8b,
	0xec, 0xad, 0xb2, 0x56, 0x29, 0x38, 0xbb, 0xd5, 0xc7, 0x87, 0x2e, 0x9e, 0x2c, 0xf7, 0x59, 0x99,
	0xf5, 0x06, 0x36, 0x65, 0x33, 0xe2, 0x58, 0x10, 0x7b, 0xbb, 0xac, 0x55, 0x36, 0xdd, 0x9c, 0x4c,
	0xb8, 0x58, 0x10, 0xeb, 0x15, 0xe4, 0xc6, 0x54, 0xa4, 0x5e, 0xbe, 0xac, 0x55, 0xf2, 0xae, 0x31,
	0xa6, 0x42, 0x59, 0x25, 0xd8, 0xf2, 0x98, 0x4f, 0xa3, 0x30, 0x75, 0x0b, 0xaa, 0x13, 0xd2, 0x94,
	0x2a, 0xd8, 0x85, 0x8d, 0x00, 0x79, 0x91, 0xb0, 0x77, 0x54, 0x63, 0x36, 0x38, 0x89, 0xc4, 0xde,
	0x17, 0x0d, 0x76, 0xae, 0x92, 0x13, 0x16, 0x05, 0x34, 0x9c, 0xf3, 0x74, 0x82, 0x7f, 0x60, 0xec,
	0x6f, 0x19, 0xb0, 0xba, 0x9e, 0xa0, 0x0b, 0xf5, 0xf8, 0xd3, 0x1f, 0x3e, 0x00, 0x03, 0xc7, 0x31,
	0x22, 0x73, 0x6a, 0x6b, 0x65, 0xad, 0xb2, 0x7d, 0x7c, 0x78, 0x7b, 0x57, 0x6a, 0xfc, 0xe9, 0x1c,
	0x3c, 0xc6, 0x49, 0x4d, 0xac, 0x62, 0x32, 0xab, 0x76, 0xe3, 0xb8, 0x3f, 0x7a, 0xe7, 0xea, 0x38,
	0x8e, 0xfb, 0x73, 0x2a, 0x79, 0x3e, 0x59, 0x28, 0xde, 0xfa, 0x8b, 0x78, 0x3d, 0xb2, 0x50, 0x3c,
	0x9f, 0x2c, 0x24, 0xef, 0x3d, 0xe4, 0x24, 0x0f, 0xfb, 0x3e, 0xb7, 0x33, 0x0a, 0x78, 0x74, 0x7b,
	0x57, 0x72, 0xfe, 0x0e, 0xd8, 0xf5, 0x7d, 0xee, 0xca, 0xb9, 0x64, 0x60, 0xb9, 0xb0, 0x19, 0x2d,
	0x27, 0x68, 0x86, 0x26, 0x64, 0x65, 0x67, 0x5f, 0xc4, 0x1c, 0x2c, 0x27, 0xc3, 0x33, 0xb2, 0x72,
	0x8d, 0x28, 0x0d, 0xac, 0x3d, 0xc8, 0xf3, 0xa4, 0x81, 0x7c, 0x8e, 0x58, 0x10, 0xcc, 0x88, 0x50,
	0x37, 0x90, 0x77, 0xb7, 0x78, 0xd2, 0xe8, 0xf1, 0x4b, 0x95, 0xb2, 0xfe, 0x03, 0x9d, 0x27, 0x0e,
	0xf2, 0xb9, 0x5a, 0x76, 0xde, 0xdd, 0xe0, 0x89, 0xd3, 0xe3, 0x72, 0xd3, 0x3c, 0x41, 0x3e, 0x99,
	0xe2, 0xd5, 0xe3, 0xa6, 0x79, 0xd2, 0x93, 0xd2, 0xfa, 0x1f, 0x0c, 0x2f, 0x40, 0x53, 0x3a, 0x13,
	0x76, 0xa1, 0x9c, 0xa9, 0x64, 0x5d, 0xdd, 0x0b, 0xce, 0xe9, 0x4c, 0xec, 0x97, 0x00, 0x7e, 0x1f,
	0x95, 0x95, 0x83, 0xec, 0xf9, 0xa5, 0xdb, 0x35, 0xd7, 0x2c, 0x03, 0x32, 0xa7, 0xc3, 0x33, 0x53,
	0xdb, 0xff, 0xa8, 0x81, 0xee, 0x92, 0x50, 0xba, 0x05, 0x80, 0xfe, 0x08, 0xb5, 0x8f, 0x9a, 0xa8,
	0xdd, 0xaa, 0x9b, 0x6b, 0x52, 0x8f, 0x86, 0xa8, 0x53, 0x77, 0x50, 0xc7, 0x69, 0x9b, 0x9a, 0xd4,
	0x27, 0x03, 0xd4, 0x6a, 0x75, 0x50, 0xab, 0xdd, 0x32, 0xd7, 0x2d, 0x00, 0xbd, 0x3f, 0x42, 0x07,
	0xcd, 0xa6, 0x99, 0x91, 0x5e, 0x77, 0x84, 0x3a, 0x8d, 0x43, 0x55, 0x9b, 0x7d, 0xa8, 0x3d, 0x68,
	0xd5, 0xd1, 0x61, 0xa3, 0x6e, 0x6e, 0xc8, 0xda, 0xee, 0x10, 0x75, 0x9c, 0xa6, 0xa9, 0x4b, 0x6f,
	0x78, 0x86, 0x3a, 0x4e, 0x5d, 0x69, 0xe3, 0xf8, 0xf4, 0xeb, 0x7d, 0x51, 0xfb, 0x7e, 0x5f, 0xd4,
	0x7e, 0xdc, 0x17, 0xb5, 0x4f, 0x3f, 0x8b, 0x6b, 0x1f, 0x0e, 0x5e, 0xf2, 0x95, 0x19, 0xeb, 0x2a,
	0xd3, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x84, 0xe2, 0x37, 0xdc, 0xa4, 0x04, 0x00, 0x00,
}
