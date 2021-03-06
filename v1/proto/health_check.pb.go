// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: health_check.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthStatus_Result int32

const (
	HealthStatus_UNKNOWN  HealthStatus_Result = 0
	HealthStatus_HEALTHY  HealthStatus_Result = 1
	HealthStatus_UNHEALTH HealthStatus_Result = 2
)

var HealthStatus_Result_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTH",
}
var HealthStatus_Result_value = map[string]int32{
	"UNKNOWN":  0,
	"HEALTHY":  1,
	"UNHEALTH": 2,
}

func (x HealthStatus_Result) String() string {
	return proto1.EnumName(HealthStatus_Result_name, int32(x))
}
func (HealthStatus_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorHealthCheck, []int{2, 0}
}

type HeathRequest struct {
	// Types that are valid to be assigned to HealthRequest:
	//	*HeathRequest_Initialization
	//	*HeathRequest_Status
	HealthRequest isHeathRequest_HealthRequest `protobuf_oneof:"health_request"`
}

func (m *HeathRequest) Reset()                    { *m = HeathRequest{} }
func (m *HeathRequest) String() string            { return proto1.CompactTextString(m) }
func (*HeathRequest) ProtoMessage()               {}
func (*HeathRequest) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{0} }

type isHeathRequest_HealthRequest interface {
	isHeathRequest_HealthRequest()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HeathRequest_Initialization struct {
	Initialization *HeathInitialization `protobuf:"bytes,1,opt,name=initialization,oneof"`
}
type HeathRequest_Status struct {
	Status *HealthStatus `protobuf:"bytes,2,opt,name=status,oneof"`
}

func (*HeathRequest_Initialization) isHeathRequest_HealthRequest() {}
func (*HeathRequest_Status) isHeathRequest_HealthRequest()         {}

func (m *HeathRequest) GetHealthRequest() isHeathRequest_HealthRequest {
	if m != nil {
		return m.HealthRequest
	}
	return nil
}

func (m *HeathRequest) GetInitialization() *HeathInitialization {
	if x, ok := m.GetHealthRequest().(*HeathRequest_Initialization); ok {
		return x.Initialization
	}
	return nil
}

func (m *HeathRequest) GetStatus() *HealthStatus {
	if x, ok := m.GetHealthRequest().(*HeathRequest_Status); ok {
		return x.Status
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HeathRequest) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _HeathRequest_OneofMarshaler, _HeathRequest_OneofUnmarshaler, _HeathRequest_OneofSizer, []interface{}{
		(*HeathRequest_Initialization)(nil),
		(*HeathRequest_Status)(nil),
	}
}

func _HeathRequest_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*HeathRequest)
	// health_request
	switch x := m.HealthRequest.(type) {
	case *HeathRequest_Initialization:
		_ = b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Initialization); err != nil {
			return err
		}
	case *HeathRequest_Status:
		_ = b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Status); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HeathRequest.HealthRequest has unexpected type %T", x)
	}
	return nil
}

func _HeathRequest_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*HeathRequest)
	switch tag {
	case 1: // health_request.initialization
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(HeathInitialization)
		err := b.DecodeMessage(msg)
		m.HealthRequest = &HeathRequest_Initialization{msg}
		return true, err
	case 2: // health_request.status
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(HealthStatus)
		err := b.DecodeMessage(msg)
		m.HealthRequest = &HeathRequest_Status{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HeathRequest_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*HeathRequest)
	// health_request
	switch x := m.HealthRequest.(type) {
	case *HeathRequest_Initialization:
		s := proto1.Size(x.Initialization)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *HeathRequest_Status:
		s := proto1.Size(x.Status)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HeathInitialization struct {
	NodeID string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (m *HeathInitialization) Reset()                    { *m = HeathInitialization{} }
func (m *HeathInitialization) String() string            { return proto1.CompactTextString(m) }
func (*HeathInitialization) ProtoMessage()               {}
func (*HeathInitialization) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{1} }

type HealthStatus struct {
	NodeID  string              `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Service string              `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Status  HealthStatus_Result `protobuf:"varint,3,opt,name=status,proto3,enum=prizem.api.v1.HealthStatus_Result" json:"status,omitempty"`
}

func (m *HealthStatus) Reset()                    { *m = HealthStatus{} }
func (m *HealthStatus) String() string            { return proto1.CompactTextString(m) }
func (*HealthStatus) ProtoMessage()               {}
func (*HealthStatus) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{2} }

type HealthAssignments struct {
	Assignments []HealthAssignment `protobuf:"bytes,1,rep,name=assignments" json:"assignments"`
}

func (m *HealthAssignments) Reset()                    { *m = HealthAssignments{} }
func (m *HealthAssignments) String() string            { return proto1.CompactTextString(m) }
func (*HealthAssignments) ProtoMessage()               {}
func (*HealthAssignments) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{3} }

type HealthAssignment struct {
	NodeID  string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *HealthAssignment) Reset()                    { *m = HealthAssignment{} }
func (m *HealthAssignment) String() string            { return proto1.CompactTextString(m) }
func (*HealthAssignment) ProtoMessage()               {}
func (*HealthAssignment) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{4} }

func init() {
	proto1.RegisterType((*HeathRequest)(nil), "prizem.api.v1.HeathRequest")
	proto1.RegisterType((*HeathInitialization)(nil), "prizem.api.v1.HeathInitialization")
	proto1.RegisterType((*HealthStatus)(nil), "prizem.api.v1.HealthStatus")
	proto1.RegisterType((*HealthAssignments)(nil), "prizem.api.v1.HealthAssignments")
	proto1.RegisterType((*HealthAssignment)(nil), "prizem.api.v1.HealthAssignment")
	proto1.RegisterEnum("prizem.api.v1.HealthStatus_Result", HealthStatus_Result_name, HealthStatus_Result_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HealthCheckReporter service

type HealthCheckReporterClient interface {
	Report(ctx context.Context, opts ...grpc.CallOption) (HealthCheckReporter_ReportClient, error)
}

type healthCheckReporterClient struct {
	cc *grpc.ClientConn
}

func NewHealthCheckReporterClient(cc *grpc.ClientConn) HealthCheckReporterClient {
	return &healthCheckReporterClient{cc}
}

func (c *healthCheckReporterClient) Report(ctx context.Context, opts ...grpc.CallOption) (HealthCheckReporter_ReportClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HealthCheckReporter_serviceDesc.Streams[0], c.cc, "/prizem.api.v1.HealthCheckReporter/Report", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthCheckReporterReportClient{stream}
	return x, nil
}

type HealthCheckReporter_ReportClient interface {
	Send(*HeathRequest) error
	Recv() (*HealthAssignments, error)
	grpc.ClientStream
}

type healthCheckReporterReportClient struct {
	grpc.ClientStream
}

func (x *healthCheckReporterReportClient) Send(m *HeathRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthCheckReporterReportClient) Recv() (*HealthAssignments, error) {
	m := new(HealthAssignments)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for HealthCheckReporter service

type HealthCheckReporterServer interface {
	Report(HealthCheckReporter_ReportServer) error
}

func RegisterHealthCheckReporterServer(s *grpc.Server, srv HealthCheckReporterServer) {
	s.RegisterService(&_HealthCheckReporter_serviceDesc, srv)
}

func _HealthCheckReporter_Report_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthCheckReporterServer).Report(&healthCheckReporterReportServer{stream})
}

type HealthCheckReporter_ReportServer interface {
	Send(*HealthAssignments) error
	Recv() (*HeathRequest, error)
	grpc.ServerStream
}

type healthCheckReporterReportServer struct {
	grpc.ServerStream
}

func (x *healthCheckReporterReportServer) Send(m *HealthAssignments) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthCheckReporterReportServer) Recv() (*HeathRequest, error) {
	m := new(HeathRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HealthCheckReporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prizem.api.v1.HealthCheckReporter",
	HandlerType: (*HealthCheckReporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Report",
			Handler:       _HealthCheckReporter_Report_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "health_check.proto",
}

func (m *HeathRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeathRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HealthRequest != nil {
		nn1, err := m.HealthRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *HeathRequest_Initialization) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Initialization != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(m.Initialization.Size()))
		n2, err := m.Initialization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *HeathRequest_Status) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Status != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(m.Status.Size()))
		n3, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *HeathInitialization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeathInitialization) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NodeID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.NodeID)))
		i += copy(dAtA[i:], m.NodeID)
	}
	return i, nil
}

func (m *HealthStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NodeID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.NodeID)))
		i += copy(dAtA[i:], m.NodeID)
	}
	if len(m.Service) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.Service)))
		i += copy(dAtA[i:], m.Service)
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *HealthAssignments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthAssignments) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Assignments) > 0 {
		for _, msg := range m.Assignments {
			dAtA[i] = 0xa
			i++
			i = encodeVarintHealthCheck(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HealthAssignment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthAssignment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NodeID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.NodeID)))
		i += copy(dAtA[i:], m.NodeID)
	}
	if len(m.Service) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.Service)))
		i += copy(dAtA[i:], m.Service)
	}
	return i, nil
}

func encodeVarintHealthCheck(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HeathRequest) Size() (n int) {
	var l int
	_ = l
	if m.HealthRequest != nil {
		n += m.HealthRequest.Size()
	}
	return n
}

func (m *HeathRequest_Initialization) Size() (n int) {
	var l int
	_ = l
	if m.Initialization != nil {
		l = m.Initialization.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	return n
}
func (m *HeathRequest_Status) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	return n
}
func (m *HeathInitialization) Size() (n int) {
	var l int
	_ = l
	l = len(m.NodeID)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	return n
}

func (m *HealthStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.NodeID)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovHealthCheck(uint64(m.Status))
	}
	return n
}

func (m *HealthAssignments) Size() (n int) {
	var l int
	_ = l
	if len(m.Assignments) > 0 {
		for _, e := range m.Assignments {
			l = e.Size()
			n += 1 + l + sovHealthCheck(uint64(l))
		}
	}
	return n
}

func (m *HealthAssignment) Size() (n int) {
	var l int
	_ = l
	l = len(m.NodeID)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	return n
}

func sovHealthCheck(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealthCheck(x uint64) (n int) {
	return sovHealthCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeathRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeathRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeathRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Initialization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HeathInitialization{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.HealthRequest = &HeathRequest_Initialization{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HealthStatus{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.HealthRequest = &HeathRequest_Status{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func (m *HeathInitialization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeathInitialization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeathInitialization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func (m *HealthStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (HealthStatus_Result(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func (m *HealthAssignments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthAssignments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthAssignments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assignments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assignments = append(m.Assignments, HealthAssignment{})
			if err := m.Assignments[len(m.Assignments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func (m *HealthAssignment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthAssignment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthAssignment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func skipHealthCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealthCheck
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
					return 0, ErrIntOverflowHealthCheck
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
					return 0, ErrIntOverflowHealthCheck
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHealthCheck
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealthCheck
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
				next, err := skipHealthCheck(dAtA[start:])
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
	ErrInvalidLengthHealthCheck = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealthCheck   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("health_check.proto", fileDescriptorHealthCheck) }

var fileDescriptorHealthCheck = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0xb5, 0x5b, 0xc8, 0x50, 0x4f, 0x19, 0x05, 0xb3, 0x89, 0x5a, 0x29, 0x13, 0xb9, 0x9b, 0x6c,
	0x48, 0xcb, 0x20, 0x36, 0xdd, 0x35, 0x80, 0xc8, 0x88, 0x12, 0x84, 0xa1, 0x42, 0x20, 0xa4, 0x2a,
	0x4d, 0x4c, 0x62, 0x91, 0x89, 0x43, 0xec, 0x74, 0xd1, 0x93, 0x70, 0x01, 0x24, 0xee, 0xc0, 0x05,
	0x66, 0xc9, 0x09, 0x2a, 0x1a, 0x2e, 0xc0, 0x11, 0x50, 0x9c, 0xe9, 0x94, 0x96, 0x48, 0x20, 0x75,
	0x15, 0xbf, 0xe4, 0xbd, 0xff, 0xdf, 0x73, 0x1e, 0xc2, 0x19, 0x8b, 0x72, 0x95, 0x1d, 0xc6, 0x19,
	0x8b, 0x3f, 0x7a, 0x65, 0x25, 0x94, 0xc0, 0xb7, 0xcb, 0x8a, 0x9f, 0xb0, 0x99, 0x17, 0x95, 0xdc,
	0x3b, 0xbe, 0xbf, 0x71, 0x2f, 0xe5, 0x2a, 0xab, 0x8f, 0xbc, 0x58, 0xcc, 0xb6, 0x53, 0x91, 0x8a,
	0x6d, 0xcd, 0x3a, 0xaa, 0x3f, 0x68, 0xa4, 0x81, 0x3e, 0x75, 0x6a, 0xf2, 0x05, 0xa2, 0xf5, 0x80,
	0x45, 0x2a, 0xa3, 0xec, 0x53, 0xcd, 0xa4, 0xc2, 0xfb, 0x68, 0xc4, 0x0b, 0xae, 0x78, 0x94, 0xf3,
	0x93, 0x48, 0x71, 0x51, 0x58, 0xd0, 0x81, 0xee, 0x70, 0x42, 0xbc, 0x4b, 0x7b, 0x3c, 0x2d, 0x9a,
	0x5e, 0x62, 0x06, 0x80, 0x5e, 0xd1, 0xe2, 0x87, 0xc8, 0x90, 0x2a, 0x52, 0xb5, 0xb4, 0x56, 0xf4,
	0x94, 0xcd, 0xbf, 0xa7, 0xe4, 0x2a, 0x7b, 0xa5, 0x29, 0x01, 0xa0, 0x0b, 0xb2, 0x6f, 0xa2, 0xd1,
	0x22, 0x69, 0xd5, 0xd9, 0x22, 0xbb, 0xe8, 0x6e, 0xcf, 0x46, 0xbc, 0x85, 0x06, 0x85, 0x48, 0xd8,
	0x21, 0x4f, 0xb4, 0xcd, 0x35, 0x1f, 0x35, 0xa7, 0x63, 0x23, 0x14, 0x09, 0x9b, 0x3e, 0xa6, 0x46,
	0xfb, 0x69, 0x9a, 0x90, 0x6f, 0x5d, 0xc6, 0xe5, 0xa2, 0xff, 0x52, 0x61, 0x0b, 0x0d, 0x24, 0xab,
	0x8e, 0x79, 0xcc, 0xb4, 0xf7, 0x35, 0x7a, 0x0e, 0xf1, 0xee, 0x32, 0xd4, 0xaa, 0x03, 0xdd, 0x51,
	0xdf, 0xd5, 0x2c, 0x77, 0x79, 0x94, 0xc9, 0x3a, 0x57, 0xe7, 0xc9, 0xc8, 0x0e, 0x32, 0xba, 0x37,
	0x78, 0x88, 0x06, 0x07, 0xe1, 0xb3, 0xf0, 0xc5, 0x9b, 0xd0, 0x04, 0x2d, 0x08, 0x9e, 0xec, 0xed,
	0xbf, 0x0e, 0xde, 0x9a, 0x10, 0xaf, 0xa3, 0x5b, 0x07, 0x61, 0x07, 0xcd, 0x15, 0xf2, 0x1e, 0xdd,
	0xe9, 0x06, 0xee, 0x49, 0xc9, 0xd3, 0x62, 0xc6, 0x0a, 0x25, 0xf1, 0x53, 0x34, 0x8c, 0x2e, 0xa0,
	0x05, 0x9d, 0x55, 0x77, 0x38, 0x19, 0xf7, 0xfa, 0xb8, 0x90, 0xf9, 0x37, 0xe6, 0xa7, 0x63, 0x40,
	0xff, 0x54, 0x92, 0x97, 0xc8, 0xbc, 0x4a, 0xbb, 0xe6, 0xf5, 0x4c, 0x12, 0xfd, 0xab, 0x72, 0x95,
	0x3d, 0x6a, 0x5b, 0x4a, 0x59, 0x29, 0x2a, 0xc5, 0x2a, 0xfc, 0xbc, 0x4d, 0xde, 0x9e, 0xf1, 0x66,
	0x5f, 0x95, 0x16, 0xfd, 0xdb, 0x70, 0xfe, 0x11, 0x42, 0x12, 0xe0, 0xc2, 0x1d, 0xe8, 0x6f, 0xcd,
	0xcf, 0x6c, 0xf0, 0xeb, 0xcc, 0x06, 0x5f, 0x1b, 0x1b, 0xcc, 0x1b, 0x1b, 0x7e, 0x6f, 0x6c, 0xf8,
	0xa3, 0xb1, 0xe1, 0xe7, 0x9f, 0x36, 0x78, 0x77, 0xb3, 0x6b, 0xbd, 0xa1, 0x1f, 0x0f, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x21, 0xb8, 0x26, 0xe3, 0x38, 0x03, 0x00, 0x00,
}
