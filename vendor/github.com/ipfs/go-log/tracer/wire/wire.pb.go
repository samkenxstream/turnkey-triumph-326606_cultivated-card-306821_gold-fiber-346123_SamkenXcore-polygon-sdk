// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wire.proto

package wire

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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

type TracerState struct {
	TraceId      uint64            `protobuf:"fixed64,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       uint64            `protobuf:"fixed64,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled,proto3" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems,proto3" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TracerState) Reset()         { *m = TracerState{} }
func (m *TracerState) String() string { return proto.CompactTextString(m) }
func (*TracerState) ProtoMessage()    {}
func (*TracerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2dcdddcdf68d8e0, []int{0}
}
func (m *TracerState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TracerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TracerState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TracerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracerState.Merge(m, src)
}
func (m *TracerState) XXX_Size() int {
	return m.Size()
}
func (m *TracerState) XXX_DiscardUnknown() {
	xxx_messageInfo_TracerState.DiscardUnknown(m)
}

var xxx_messageInfo_TracerState proto.InternalMessageInfo

func (m *TracerState) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *TracerState) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *TracerState) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func (m *TracerState) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*TracerState)(nil), "loggabletracer.wire.TracerState")
	proto.RegisterMapType((map[string]string)(nil), "loggabletracer.wire.TracerState.BaggageItemsEntry")
}

func init() { proto.RegisterFile("wire.proto", fileDescriptor_f2dcdddcdf68d8e0) }

var fileDescriptor_f2dcdddcdf68d8e0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xcf, 0x2c, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0xc9, 0x4f, 0x4f, 0x4f, 0x4c, 0xca, 0x49,
	0x2d, 0x29, 0x4a, 0x4c, 0x4e, 0x2d, 0xd2, 0x03, 0x49, 0x29, 0x7d, 0x65, 0xe4, 0xe2, 0x0e, 0x01,
	0xf3, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0x24, 0xb9, 0x38, 0xc0, 0xd2, 0xf1, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x6c, 0x41, 0xec, 0x60, 0xbe, 0x67, 0x8a, 0x90, 0x38, 0x17, 0x7b, 0x71,
	0x41, 0x62, 0x1e, 0x48, 0x86, 0x09, 0x2c, 0xc3, 0x06, 0xe2, 0x7a, 0xa6, 0x08, 0x49, 0x70, 0xb1,
	0x17, 0x27, 0xe6, 0x16, 0xe4, 0xa4, 0xa6, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xb8,
	0x42, 0xe1, 0x5c, 0xbc, 0x49, 0x89, 0xe9, 0xe9, 0x89, 0xe9, 0xa9, 0xf1, 0x99, 0x25, 0xa9, 0xb9,
	0xc5, 0x12, 0x2c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x46, 0x7a, 0x58, 0x9c, 0xa2, 0x87, 0xe4, 0x0c,
	0x3d, 0x27, 0x88, 0x2e, 0x4f, 0x90, 0x26, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0x9e, 0x24, 0x24,
	0x21, 0x29, 0x7b, 0x2e, 0x41, 0x0c, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x60, 0x67,
	0x73, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x60, 0x07, 0x73,
	0x06, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4e, 0x72, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xc5, 0x02, 0x72, 0x4c, 0x12, 0x1b, 0x38, 0xcc, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0x48, 0xf4, 0xf8, 0x41, 0x01, 0x00, 0x00,
}

func (m *TracerState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TracerState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TraceId != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.TraceId))
		i += 8
	}
	if m.SpanId != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.SpanId))
		i += 8
	}
	if m.Sampled {
		dAtA[i] = 0x18
		i++
		if m.Sampled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.BaggageItems) > 0 {
		for k, _ := range m.BaggageItems {
			dAtA[i] = 0x22
			i++
			v := m.BaggageItems[k]
			mapSize := 1 + len(k) + sovWire(uint64(len(k))) + 1 + len(v) + sovWire(uint64(len(v)))
			i = encodeVarintWire(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintWire(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintWire(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintWire(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TracerState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TraceId != 0 {
		n += 9
	}
	if m.SpanId != 0 {
		n += 9
	}
	if m.Sampled {
		n += 2
	}
	if len(m.BaggageItems) > 0 {
		for k, v := range m.BaggageItems {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWire(uint64(len(k))) + 1 + len(v) + sovWire(uint64(len(v)))
			n += mapEntrySize + 1 + sovWire(uint64(mapEntrySize))
		}
	}
	return n
}

func sovWire(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWire(x uint64) (n int) {
	return sovWire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TracerState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
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
			return fmt.Errorf("proto: TracerState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TracerState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			m.TraceId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.SpanId = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sampled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sampled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaggageItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaggageItems == nil {
				m.BaggageItems = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWire
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWire
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWire
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWire
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWire
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWire
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWire
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWire(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWire
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.BaggageItems[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWire
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
func skipWire(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWire
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
					return 0, ErrIntOverflowWire
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
					return 0, ErrIntOverflowWire
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
				return 0, ErrInvalidLengthWire
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWire
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWire
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
				next, err := skipWire(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWire
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
	ErrInvalidLengthWire = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWire   = fmt.Errorf("proto: integer overflow")
)