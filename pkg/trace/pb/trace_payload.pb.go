// Code generated by protoc-gen-gogo.
// source: trace_payload.proto
// DO NOT EDIT!

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TracePayload struct {
	HostName     string      `protobuf:"bytes,1,opt,name=hostName,proto3" json:"hostName,omitempty"`
	Env          string      `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	Traces       []*APITrace `protobuf:"bytes,3,rep,name=traces" json:"traces,omitempty"`
	Transactions []*Span     `protobuf:"bytes,4,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *TracePayload) Reset()                    { *m = TracePayload{} }
func (m *TracePayload) String() string            { return proto.CompactTextString(m) }
func (*TracePayload) ProtoMessage()               {}
func (*TracePayload) Descriptor() ([]byte, []int) { return fileDescriptorTracePayload, []int{0} }

func (m *TracePayload) GetTraces() []*APITrace {
	if m != nil {
		return m.Traces
	}
	return nil
}

func (m *TracePayload) GetTransactions() []*Span {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*TracePayload)(nil), "model.TracePayload")
}
func (m *TracePayload) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TracePayload) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HostName) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintTracePayload(data, i, uint64(len(m.HostName)))
		i += copy(data[i:], m.HostName)
	}
	if len(m.Env) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintTracePayload(data, i, uint64(len(m.Env)))
		i += copy(data[i:], m.Env)
	}
	if len(m.Traces) > 0 {
		for _, msg := range m.Traces {
			data[i] = 0x1a
			i++
			i = encodeVarintTracePayload(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Transactions) > 0 {
		for _, msg := range m.Transactions {
			data[i] = 0x22
			i++
			i = encodeVarintTracePayload(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64TracePayload(data []byte, offset int, v uint64) int {
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
func encodeFixed32TracePayload(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTracePayload(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *TracePayload) Size() (n int) {
	var l int
	_ = l
	l = len(m.HostName)
	if l > 0 {
		n += 1 + l + sovTracePayload(uint64(l))
	}
	l = len(m.Env)
	if l > 0 {
		n += 1 + l + sovTracePayload(uint64(l))
	}
	if len(m.Traces) > 0 {
		for _, e := range m.Traces {
			l = e.Size()
			n += 1 + l + sovTracePayload(uint64(l))
		}
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovTracePayload(uint64(l))
		}
	}
	return n
}

func sovTracePayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTracePayload(x uint64) (n int) {
	return sovTracePayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TracePayload) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracePayload
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
			return fmt.Errorf("proto: TracePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TracePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracePayload
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
				return ErrInvalidLengthTracePayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracePayload
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
				return ErrInvalidLengthTracePayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Traces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracePayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracePayload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Traces = append(m.Traces, &APITrace{})
			if err := m.Traces[len(m.Traces)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracePayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracePayload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &Span{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracePayload(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracePayload
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
func skipTracePayload(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracePayload
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
					return 0, ErrIntOverflowTracePayload
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
					return 0, ErrIntOverflowTracePayload
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
				return 0, ErrInvalidLengthTracePayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTracePayload
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
				next, err := skipTracePayload(data[start:])
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
	ErrInvalidLengthTracePayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracePayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("trace_payload.proto", fileDescriptorTracePayload) }

var fileDescriptorTracePayload = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x29, 0x4a, 0x4c,
	0x4e, 0x8d, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x91, 0xe2, 0x06, 0xcb, 0x41, 0xc4, 0xa4, 0xb8, 0x8a, 0x0b,
	0x12, 0xf3, 0x20, 0x6c, 0xa5, 0x69, 0x8c, 0x5c, 0x3c, 0x21, 0x20, 0xb9, 0x00, 0x88, 0x36, 0x21,
	0x29, 0x2e, 0x8e, 0x8c, 0xfc, 0xe2, 0x12, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x38, 0x5f, 0x48, 0x80, 0x8b, 0x39, 0x35, 0xaf, 0x4c, 0x82, 0x09, 0x2c, 0x0c, 0x62,
	0x0a, 0xa9, 0x73, 0xb1, 0x81, 0x4d, 0x2e, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd7,
	0x03, 0xdb, 0xa7, 0xe7, 0x18, 0xe0, 0x09, 0x36, 0x35, 0x08, 0x2a, 0x2d, 0xa4, 0xcf, 0xc5, 0x53,
	0x52, 0x94, 0x98, 0x57, 0x9c, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x57, 0x2c, 0xc1, 0x02, 0x56, 0xce,
	0x0d, 0x55, 0x1e, 0x5c, 0x90, 0x98, 0x17, 0x84, 0xa2, 0xc0, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec,
	0x62, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xc9, 0x04, 0x9e, 0xe8, 0x00, 0x00, 0x00,
}
