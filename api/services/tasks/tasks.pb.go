// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tasks.proto

package tasks

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type QueueTaskRequest struct {
}

func (m *QueueTaskRequest) Reset()      { *m = QueueTaskRequest{} }
func (*QueueTaskRequest) ProtoMessage() {}
func (*QueueTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{0}
}
func (m *QueueTaskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueTaskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueTaskRequest.Merge(m, src)
}
func (m *QueueTaskRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueueTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueTaskRequest proto.InternalMessageInfo

func (*QueueTaskRequest) XXX_MessageName() string {
	return "tasks.QueueTaskRequest"
}

type QueueTaskResponse struct {
}

func (m *QueueTaskResponse) Reset()      { *m = QueueTaskResponse{} }
func (*QueueTaskResponse) ProtoMessage() {}
func (*QueueTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{1}
}
func (m *QueueTaskResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueTaskResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueTaskResponse.Merge(m, src)
}
func (m *QueueTaskResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueueTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueueTaskResponse proto.InternalMessageInfo

func (*QueueTaskResponse) XXX_MessageName() string {
	return "tasks.QueueTaskResponse"
}
func init() {
	proto.RegisterType((*QueueTaskRequest)(nil), "tasks.QueueTaskRequest")
	golang_proto.RegisterType((*QueueTaskRequest)(nil), "tasks.QueueTaskRequest")
	proto.RegisterType((*QueueTaskResponse)(nil), "tasks.QueueTaskResponse")
	golang_proto.RegisterType((*QueueTaskResponse)(nil), "tasks.QueueTaskResponse")
}

func init() { proto.RegisterFile("tasks.proto", fileDescriptor_b3834c8ef8464a3f) }
func init() { golang_proto.RegisterFile("tasks.proto", fileDescriptor_b3834c8ef8464a3f) }

var fileDescriptor_b3834c8ef8464a3f = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x49, 0x2c, 0xce,
	0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x22, 0xfa, 0x20, 0x16, 0x44, 0x52, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55,
	0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0xaa,
	0x55, 0x49, 0x88, 0x4b, 0x20, 0xb0, 0x34, 0xb5, 0x34, 0x35, 0x24, 0xb1, 0x38, 0x3b, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x49, 0x98, 0x4b, 0x10, 0x49, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0xd5, 0x28, 0x9d, 0x8b, 0x15, 0xc4, 0x2f, 0x16, 0x8a, 0xe3, 0xe2, 0x84, 0xcb, 0x0a, 0x89, 0xeb,
	0x41, 0xdc, 0x81, 0x6e, 0x86, 0x94, 0x04, 0xa6, 0x04, 0xc4, 0x20, 0x25, 0xb9, 0xa6, 0xcb, 0x4f,
	0x26, 0x33, 0x49, 0x28, 0x09, 0x83, 0x5d, 0x54, 0x66, 0xa8, 0x0f, 0x56, 0xa8, 0x5f, 0x08, 0x52,
	0x68, 0xc5, 0xa8, 0xe5, 0xe4, 0x72, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e,
	0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d,
	0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc0, 0x63, 0x39,
	0xc6, 0x13, 0x8f, 0xe5, 0x18, 0x2f, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d,
	0xec, 0x3d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x49, 0xe3, 0x23, 0x28, 0x01, 0x00,
	0x00,
}

func (this *QueueTaskRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueTaskRequest)
	if !ok {
		that2, ok := that.(QueueTaskRequest)
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
	return true
}
func (this *QueueTaskResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueueTaskResponse)
	if !ok {
		that2, ok := that.(QueueTaskResponse)
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
	return true
}
func (this *QueueTaskRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&tasks.QueueTaskRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueueTaskResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&tasks.QueueTaskResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTasks(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TasksClient interface {
	QueueTask(ctx context.Context, in *QueueTaskRequest, opts ...grpc.CallOption) (*QueueTaskResponse, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) QueueTask(ctx context.Context, in *QueueTaskRequest, opts ...grpc.CallOption) (*QueueTaskResponse, error) {
	out := new(QueueTaskResponse)
	err := c.cc.Invoke(ctx, "/tasks.Tasks/QueueTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
type TasksServer interface {
	QueueTask(context.Context, *QueueTaskRequest) (*QueueTaskResponse, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_QueueTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).QueueTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.Tasks/QueueTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).QueueTask(ctx, req.(*QueueTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tasks.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueueTask",
			Handler:    _Tasks_QueueTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks.proto",
}

func (m *QueueTaskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueTaskRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *QueueTaskResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueTaskResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintTasks(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedQueueTaskRequest(r randyTasks, easy bool) *QueueTaskRequest {
	this := &QueueTaskRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedQueueTaskResponse(r randyTasks, easy bool) *QueueTaskResponse {
	this := &QueueTaskResponse{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTasks interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTasks(r randyTasks) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTasks(r randyTasks) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTasks(r)
	}
	return string(tmps)
}
func randUnrecognizedTasks(r randyTasks, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTasks(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTasks(dAtA []byte, r randyTasks, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTasks(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *QueueTaskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueueTaskResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTasks(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTasks(x uint64) (n int) {
	return sovTasks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueueTaskRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueueTaskRequest{`,
		`}`,
	}, "")
	return s
}
func (this *QueueTaskResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueueTaskResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringTasks(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueueTaskRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: QueueTaskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueTaskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTasks
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
func (m *QueueTaskResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: QueueTaskResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueTaskResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTasks
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
func skipTasks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTasks
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
					return 0, ErrIntOverflowTasks
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
					return 0, ErrIntOverflowTasks
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
				return 0, ErrInvalidLengthTasks
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTasks
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTasks
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
				next, err := skipTasks(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTasks
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
	ErrInvalidLengthTasks = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTasks   = fmt.Errorf("proto: integer overflow")
)
