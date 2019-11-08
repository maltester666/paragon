// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agent.proto

package events

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	codec "github.com/kcarretto/paragon/api/codec"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AgentCheckin struct {
	PublicIP string               `protobuf:"bytes,1,opt,name=publicIP,proto3" json:"publicIP,omitempty"`
	SeenTime int64                `protobuf:"varint,2,opt,name=SeenTime,proto3" json:"SeenTime,omitempty"`
	Agent    *codec.AgentMetadata `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (m *AgentCheckin) Reset()      { *m = AgentCheckin{} }
func (*AgentCheckin) ProtoMessage() {}
func (*AgentCheckin) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}
func (m *AgentCheckin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AgentCheckin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AgentCheckin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgentCheckin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentCheckin.Merge(m, src)
}
func (m *AgentCheckin) XXX_Size() int {
	return m.Size()
}
func (m *AgentCheckin) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentCheckin.DiscardUnknown(m)
}

var xxx_messageInfo_AgentCheckin proto.InternalMessageInfo

func (m *AgentCheckin) GetPublicIP() string {
	if m != nil {
		return m.PublicIP
	}
	return ""
}

func (m *AgentCheckin) GetSeenTime() int64 {
	if m != nil {
		return m.SeenTime
	}
	return 0
}

func (m *AgentCheckin) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentCheckin)(nil), "events.AgentCheckin")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x96,
	0x12, 0x4d, 0x2c, 0xc8, 0xd4, 0x4f, 0xce, 0x4f, 0x49, 0x4d, 0xd6, 0x47, 0x92, 0x56, 0x2a, 0xe2,
	0xe2, 0x71, 0x04, 0x71, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x33, 0xf3, 0x84, 0xa4, 0xb8, 0x38, 0x0a,
	0x4a, 0x93, 0x72, 0x32, 0x93, 0x3d, 0x03, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c,
	0x90, 0x5c, 0x70, 0x6a, 0x6a, 0x5e, 0x48, 0x66, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x73,
	0x10, 0x9c, 0x2f, 0xa4, 0xc5, 0xc5, 0x0a, 0x36, 0x56, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x44, 0x0f, 0x6c, 0x95, 0x1e, 0xd8, 0x6c, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x20,
	0x88, 0x12, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e,
	0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x07, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x62, 0x75, 0x27, 0xde, 0x00, 0x00, 0x00,
}

func (this *AgentCheckin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgentCheckin)
	if !ok {
		that2, ok := that.(AgentCheckin)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PublicIP != that1.PublicIP {
		return false
	}
	if this.SeenTime != that1.SeenTime {
		return false
	}
	if !this.Agent.Equal(that1.Agent) {
		return false
	}
	return true
}
func (this *AgentCheckin) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&events.AgentCheckin{")
	s = append(s, "PublicIP: "+fmt.Sprintf("%#v", this.PublicIP)+",\n")
	s = append(s, "SeenTime: "+fmt.Sprintf("%#v", this.SeenTime)+",\n")
	if this.Agent != nil {
		s = append(s, "Agent: "+fmt.Sprintf("%#v", this.Agent)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAgent(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AgentCheckin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentCheckin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PublicIP) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.PublicIP)))
		i += copy(dAtA[i:], m.PublicIP)
	}
	if m.SeenTime != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.SeenTime))
	}
	if m.Agent != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.Agent.Size()))
		n1, err := m.Agent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AgentCheckin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicIP)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.SeenTime != 0 {
		n += 1 + sovAgent(uint64(m.SeenTime))
	}
	if m.Agent != nil {
		l = m.Agent.Size()
		n += 1 + l + sovAgent(uint64(l))
	}
	return n
}

func sovAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AgentCheckin) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentCheckin{`,
		`PublicIP:` + fmt.Sprintf("%v", this.PublicIP) + `,`,
		`SeenTime:` + fmt.Sprintf("%v", this.SeenTime) + `,`,
		`Agent:` + strings.Replace(fmt.Sprintf("%v", this.Agent), "AgentMetadata", "codec.AgentMetadata", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAgent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AgentCheckin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AgentCheckin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentCheckin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeenTime", wireType)
			}
			m.SeenTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeenTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Agent == nil {
				m.Agent = &codec.AgentMetadata{}
			}
			if err := m.Agent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
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
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAgent
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAgent
				}
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
	ErrInvalidLengthAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent   = fmt.Errorf("proto: integer overflow")
)
