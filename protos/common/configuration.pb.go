// Code generated by protoc-gen-go.
// source: common/configuration.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HashingAlgorithm is encoded into the configuration transaction as  a configuration item of type Chain
// with a Key of "HashingAlgorithm" and a Value of  HashingAlgorithm as marshaled protobuf bytes
type HashingAlgorithm struct {
	// Currently supported algorithms are: SHAKE256
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HashingAlgorithm) Reset()                    { *m = HashingAlgorithm{} }
func (m *HashingAlgorithm) String() string            { return proto.CompactTextString(m) }
func (*HashingAlgorithm) ProtoMessage()               {}
func (*HashingAlgorithm) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// BlockDataHashingStructure is encoded into the configuration transaction as a configuration item of
// type Chain with a Key of "BlockDataHashingStructure" and a Value of HashingAlgorithm as marshaled protobuf bytes
type BlockDataHashingStructure struct {
	// width specifies the width of the Merkle tree to use when computing the BlockDataHash
	// in order to replicate flat hashing, set this width to MAX_UINT32
	Width uint32 `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
}

func (m *BlockDataHashingStructure) Reset()                    { *m = BlockDataHashingStructure{} }
func (m *BlockDataHashingStructure) String() string            { return proto.CompactTextString(m) }
func (*BlockDataHashingStructure) ProtoMessage()               {}
func (*BlockDataHashingStructure) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// OrdererAddresses is encoded into the configuration transaction as a configuration item of type Chain
// with a Key of "OrdererAddresses" and a Value of OrdererAddresses as marshaled protobuf bytes
type OrdererAddresses struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *OrdererAddresses) Reset()                    { *m = OrdererAddresses{} }
func (m *OrdererAddresses) String() string            { return proto.CompactTextString(m) }
func (*OrdererAddresses) ProtoMessage()               {}
func (*OrdererAddresses) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func init() {
	proto.RegisterType((*HashingAlgorithm)(nil), "common.HashingAlgorithm")
	proto.RegisterType((*BlockDataHashingStructure)(nil), "common.BlockDataHashingStructure")
	proto.RegisterType((*OrdererAddresses)(nil), "common.OrdererAddresses")
}

func init() { proto.RegisterFile("common/configuration.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0xd4, 0xc2, 0x0e, 0x08, 0x25, 0x78, 0xa8, 0xe2, 0xa1, 0x2c, 0x22, 0x05, 0x61,
	0xa3, 0xf8, 0x0b, 0x5a, 0x3c, 0x78, 0x13, 0xb6, 0x37, 0x6f, 0xd9, 0x64, 0x9a, 0x04, 0x77, 0x33,
	0x65, 0x32, 0x8b, 0xf8, 0xef, 0xc5, 0x8d, 0x62, 0x6f, 0xf3, 0xcd, 0xfb, 0x1e, 0x3c, 0xb8, 0xb1,
	0x34, 0x8e, 0x94, 0xb4, 0xa5, 0x74, 0x88, 0x7e, 0x62, 0x23, 0x91, 0x52, 0x7b, 0x64, 0x12, 0x52,
	0x8b, 0x92, 0x35, 0xf7, 0xb0, 0x7c, 0x35, 0x39, 0xc4, 0xe4, 0xb7, 0x83, 0x27, 0x8e, 0x12, 0x46,
	0xa5, 0xe0, 0x3c, 0x99, 0x11, 0x57, 0xd5, 0xba, 0xda, 0xd4, 0xdd, 0x7c, 0x37, 0x4f, 0x70, 0xbd,
	0x1b, 0xc8, 0x7e, 0xbc, 0x18, 0x31, 0xbf, 0x85, 0xbd, 0xf0, 0x64, 0x65, 0x62, 0x54, 0x57, 0x70,
	0xf1, 0x19, 0x9d, 0x84, 0xb9, 0x71, 0xd9, 0x15, 0x68, 0x1e, 0x61, 0xf9, 0xc6, 0x0e, 0x19, 0x79,
	0xeb, 0x1c, 0x63, 0xce, 0x98, 0xd5, 0x2d, 0xd4, 0xe6, 0x0f, 0x56, 0xd5, 0xfa, 0x6c, 0x53, 0x77,
	0xff, 0x8f, 0xdd, 0x1e, 0xee, 0x88, 0x7d, 0x1b, 0xbe, 0x8e, 0xc8, 0x03, 0x3a, 0x8f, 0xdc, 0x1e,
	0x4c, 0xcf, 0xd1, 0x96, 0xd1, 0xb9, 0x2d, 0xa3, 0xdf, 0x1f, 0x7c, 0x94, 0x30, 0xf5, 0x3f, 0xa8,
	0x4f, 0x64, 0x5d, 0x64, 0x5d, 0x64, 0x5d, 0xe4, 0x7e, 0x31, 0xe3, 0xf3, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0x7e, 0xef, 0xa1, 0x0e, 0x01, 0x00, 0x00,
}