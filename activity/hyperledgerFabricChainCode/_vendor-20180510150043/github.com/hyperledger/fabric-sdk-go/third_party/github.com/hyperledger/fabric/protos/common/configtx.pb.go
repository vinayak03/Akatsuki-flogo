/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/configtx.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ConfigType is an enumeration of possible types for the config.  The type field in the config is
// an int32 for extensibility, but this enum type should generally be used to populate it
type ConfigType int32

const (
	ConfigType_CHANNEL  ConfigType = 0
	ConfigType_RESOURCE ConfigType = 1
)

var ConfigType_name = map[int32]string{
	0: "CHANNEL",
	1: "RESOURCE",
}
var ConfigType_value = map[string]int32{
	"CHANNEL":  0,
	"RESOURCE": 1,
}

func (x ConfigType) String() string {
	return proto.EnumName(ConfigType_name, int32(x))
}
func (ConfigType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// ConfigEnvelope is designed to contain _all_ configuration for a chain with no dependency
// on previous configuration transactions.
//
// It is generated with the following scheme:
//   1. Retrieve the existing configuration
//   2. Note the config properties (ConfigValue, ConfigPolicy, ConfigGroup) to be modified
//   3. Add any intermediate ConfigGroups to the ConfigUpdate.read_set (sparsely)
//   4. Add any additional desired dependencies to ConfigUpdate.read_set (sparsely)
//   5. Modify the config properties, incrementing each version by 1, set them in the ConfigUpdate.write_set
//      Note: any element not modified but specified should already be in the read_set, so may be specified sparsely
//   6. Create ConfigUpdate message and marshal it into ConfigUpdateEnvelope.update and encode the required signatures
//     a) Each signature is of type ConfigSignature
//     b) The ConfigSignature signature is over the concatenation of signature_header and the ConfigUpdate bytes (which includes a ChainHeader)
//   5. Submit new Config for ordering in Envelope signed by submitter
//     a) The Envelope Payload has data set to the marshaled ConfigEnvelope
//     b) The Envelope Payload has a header of type Header.Type.CONFIG_UPDATE
//
// The configuration manager will verify:
//   1. All items in the read_set exist at the read versions
//   2. All items in the write_set at a different version than, or not in, the read_set have been appropriately signed according to their mod_policy
//   3. The new configuration satisfies the ConfigSchema
type ConfigEnvelope struct {
	Config     *Config   `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	LastUpdate *Envelope `protobuf:"bytes,2,opt,name=last_update,json=lastUpdate" json:"last_update,omitempty"`
}

func (m *ConfigEnvelope) Reset()                    { *m = ConfigEnvelope{} }
func (m *ConfigEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ConfigEnvelope) ProtoMessage()               {}
func (*ConfigEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ConfigEnvelope) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigEnvelope) GetLastUpdate() *Envelope {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

type ConfigGroupSchema struct {
	Groups   map[string]*ConfigGroupSchema  `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values   map[string]*ConfigValueSchema  `protobuf:"bytes,2,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies map[string]*ConfigPolicySchema `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ConfigGroupSchema) Reset()                    { *m = ConfigGroupSchema{} }
func (m *ConfigGroupSchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigGroupSchema) ProtoMessage()               {}
func (*ConfigGroupSchema) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ConfigGroupSchema) GetGroups() map[string]*ConfigGroupSchema {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigGroupSchema) GetValues() map[string]*ConfigValueSchema {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ConfigGroupSchema) GetPolicies() map[string]*ConfigPolicySchema {
	if m != nil {
		return m.Policies
	}
	return nil
}

type ConfigValueSchema struct {
}

func (m *ConfigValueSchema) Reset()                    { *m = ConfigValueSchema{} }
func (m *ConfigValueSchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigValueSchema) ProtoMessage()               {}
func (*ConfigValueSchema) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type ConfigPolicySchema struct {
}

func (m *ConfigPolicySchema) Reset()                    { *m = ConfigPolicySchema{} }
func (m *ConfigPolicySchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigPolicySchema) ProtoMessage()               {}
func (*ConfigPolicySchema) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// Config represents the config for a particular channel
type Config struct {
	Sequence     uint64       `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	ChannelGroup *ConfigGroup `protobuf:"bytes,2,opt,name=channel_group,json=channelGroup" json:"channel_group,omitempty"`
	Type         int32        `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Config) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Config) GetChannelGroup() *ConfigGroup {
	if m != nil {
		return m.ChannelGroup
	}
	return nil
}

func (m *Config) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ConfigUpdateEnvelope struct {
	ConfigUpdate []byte             `protobuf:"bytes,1,opt,name=config_update,json=configUpdate,proto3" json:"config_update,omitempty"`
	Signatures   []*ConfigSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *ConfigUpdateEnvelope) Reset()                    { *m = ConfigUpdateEnvelope{} }
func (m *ConfigUpdateEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ConfigUpdateEnvelope) ProtoMessage()               {}
func (*ConfigUpdateEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ConfigUpdateEnvelope) GetConfigUpdate() []byte {
	if m != nil {
		return m.ConfigUpdate
	}
	return nil
}

func (m *ConfigUpdateEnvelope) GetSignatures() []*ConfigSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// ConfigUpdate is used to submit a subset of config and to have the orderer apply to Config
// it is always submitted inside a ConfigUpdateEnvelope which allows the addition of signatures
// resulting in a new total configuration.  The update is applied as follows:
// 1. The versions from all of the elements in the read_set is verified against the versions in the existing config.
//    If there is a mismatch in the read versions, then the config update fails and is rejected.
// 2. Any elements in the write_set with the same version as the read_set are ignored.
// 3. The corresponding mod_policy for every remaining element in the write_set is collected.
// 4. Each policy is checked against the signatures from the ConfigUpdateEnvelope, any failing to verify are rejected
// 5. The write_set is applied to the Config and the ConfigGroupSchema verifies that the updates were legal
type ConfigUpdate struct {
	ChannelId    string            `protobuf:"bytes,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	ReadSet      *ConfigGroup      `protobuf:"bytes,2,opt,name=read_set,json=readSet" json:"read_set,omitempty"`
	WriteSet     *ConfigGroup      `protobuf:"bytes,3,opt,name=write_set,json=writeSet" json:"write_set,omitempty"`
	Type         int32             `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	IsolatedData map[string][]byte `protobuf:"bytes,5,rep,name=isolated_data,json=isolatedData" json:"isolated_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ConfigUpdate) Reset()                    { *m = ConfigUpdate{} }
func (m *ConfigUpdate) String() string            { return proto.CompactTextString(m) }
func (*ConfigUpdate) ProtoMessage()               {}
func (*ConfigUpdate) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ConfigUpdate) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ConfigUpdate) GetReadSet() *ConfigGroup {
	if m != nil {
		return m.ReadSet
	}
	return nil
}

func (m *ConfigUpdate) GetWriteSet() *ConfigGroup {
	if m != nil {
		return m.WriteSet
	}
	return nil
}

func (m *ConfigUpdate) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ConfigUpdate) GetIsolatedData() map[string][]byte {
	if m != nil {
		return m.IsolatedData
	}
	return nil
}

// ConfigGroup is the hierarchical data structure for holding config
type ConfigGroup struct {
	Version   uint64                   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Groups    map[string]*ConfigGroup  `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values    map[string]*ConfigValue  `protobuf:"bytes,3,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies  map[string]*ConfigPolicy `protobuf:"bytes,4,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ModPolicy string                   `protobuf:"bytes,5,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigGroup) Reset()                    { *m = ConfigGroup{} }
func (m *ConfigGroup) String() string            { return proto.CompactTextString(m) }
func (*ConfigGroup) ProtoMessage()               {}
func (*ConfigGroup) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ConfigGroup) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ConfigGroup) GetGroups() map[string]*ConfigGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigGroup) GetValues() map[string]*ConfigValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ConfigGroup) GetPolicies() map[string]*ConfigPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *ConfigGroup) GetModPolicy() string {
	if m != nil {
		return m.ModPolicy
	}
	return ""
}

// ConfigValue represents an individual piece of config data
type ConfigValue struct {
	Version   uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ModPolicy string `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigValue) Reset()                    { *m = ConfigValue{} }
func (m *ConfigValue) String() string            { return proto.CompactTextString(m) }
func (*ConfigValue) ProtoMessage()               {}
func (*ConfigValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ConfigValue) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ConfigValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ConfigValue) GetModPolicy() string {
	if m != nil {
		return m.ModPolicy
	}
	return ""
}

type ConfigPolicy struct {
	Version   uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Policy    *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
	ModPolicy string  `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigPolicy) Reset()                    { *m = ConfigPolicy{} }
func (m *ConfigPolicy) String() string            { return proto.CompactTextString(m) }
func (*ConfigPolicy) ProtoMessage()               {}
func (*ConfigPolicy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *ConfigPolicy) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ConfigPolicy) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *ConfigPolicy) GetModPolicy() string {
	if m != nil {
		return m.ModPolicy
	}
	return ""
}

type ConfigSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfigSignature) Reset()                    { *m = ConfigSignature{} }
func (m *ConfigSignature) String() string            { return proto.CompactTextString(m) }
func (*ConfigSignature) ProtoMessage()               {}
func (*ConfigSignature) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *ConfigSignature) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *ConfigSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigEnvelope)(nil), "sdk.common.ConfigEnvelope")
	proto.RegisterType((*ConfigGroupSchema)(nil), "sdk.common.ConfigGroupSchema")
	proto.RegisterType((*ConfigValueSchema)(nil), "sdk.common.ConfigValueSchema")
	proto.RegisterType((*ConfigPolicySchema)(nil), "sdk.common.ConfigPolicySchema")
	proto.RegisterType((*Config)(nil), "sdk.common.Config")
	proto.RegisterType((*ConfigUpdateEnvelope)(nil), "sdk.common.ConfigUpdateEnvelope")
	proto.RegisterType((*ConfigUpdate)(nil), "sdk.common.ConfigUpdate")
	proto.RegisterType((*ConfigGroup)(nil), "sdk.common.ConfigGroup")
	proto.RegisterType((*ConfigValue)(nil), "sdk.common.ConfigValue")
	proto.RegisterType((*ConfigPolicy)(nil), "sdk.common.ConfigPolicy")
	proto.RegisterType((*ConfigSignature)(nil), "sdk.common.ConfigSignature")
	proto.RegisterEnum("sdk.common.ConfigType", ConfigType_name, ConfigType_value)
}

func init() { proto.RegisterFile("common/configtx.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xff, 0x6e, 0x12, 0x4b,
	0x14, 0xbe, 0xb0, 0x40, 0xe1, 0x00, 0x2d, 0x9d, 0x72, 0x73, 0xf7, 0x12, 0x8d, 0x75, 0xd5, 0xfe,
	0x32, 0x81, 0x5a, 0xff, 0x68, 0x63, 0xd2, 0x18, 0x45, 0x62, 0x1b, 0x0d, 0xea, 0xd2, 0x6a, 0xd2,
	0x98, 0x90, 0xe9, 0xee, 0x14, 0x36, 0x85, 0x9d, 0x75, 0x77, 0xa8, 0xf2, 0x58, 0x3e, 0x8f, 0x6f,
	0xe0, 0x53, 0x98, 0x9d, 0x99, 0xdd, 0xce, 0x96, 0x05, 0xe2, 0x5f, 0xdd, 0x73, 0xe6, 0xfb, 0xbe,
	0x39, 0xbf, 0xe6, 0x50, 0xf8, 0xd7, 0xa2, 0xe3, 0x31, 0x75, 0x5b, 0x16, 0x75, 0xaf, 0x9c, 0x01,
	0xfb, 0xd1, 0xf4, 0x7c, 0xca, 0x28, 0x2a, 0x08, 0x77, 0x63, 0x23, 0x3e, 0x0e, 0xff, 0x88, 0xc3,
	0x46, 0xc4, 0xf1, 0xe8, 0xc8, 0xb1, 0x1c, 0x12, 0x08, 0xb7, 0x71, 0x0d, 0xab, 0x6d, 0xae, 0xd2,
	0x71, 0x6f, 0xc8, 0x88, 0x7a, 0x04, 0x6d, 0x41, 0x41, 0xe8, 0xea, 0x99, 0xcd, 0xcc, 0x4e, 0xf9,
	0x60, 0xb5, 0x29, 0x75, 0x04, 0xce, 0x94, 0xa7, 0xe8, 0x19, 0x94, 0x47, 0x38, 0x60, 0xfd, 0x89,
	0x67, 0x63, 0x46, 0xf4, 0x2c, 0x07, 0xd7, 0x22, 0x70, 0x24, 0x67, 0x42, 0x08, 0x3a, 0xe7, 0x18,
	0xe3, 0x97, 0x06, 0xeb, 0x42, 0xe5, 0xad, 0x4f, 0x27, 0x5e, 0xcf, 0x1a, 0x92, 0x31, 0x46, 0xc7,
	0x50, 0x18, 0x84, 0x66, 0xa0, 0x67, 0x36, 0xb5, 0x9d, 0xf2, 0xc1, 0x93, 0xe4, 0x85, 0x0a, 0xb4,
	0xc9, 0xbf, 0x83, 0x8e, 0xcb, 0xfc, 0xa9, 0x29, 0x49, 0x21, 0xfd, 0x06, 0x8f, 0x26, 0x24, 0xd0,
	0xb3, 0xcb, 0xe8, 0x9f, 0x39, 0x4e, 0xd2, 0x05, 0x09, 0xb5, 0xa1, 0x18, 0x95, 0x44, 0xd7, 0xb8,
	0xc0, 0xf6, 0x7c, 0x81, 0x8f, 0x12, 0x29, 0x24, 0x62, 0x62, 0xe3, 0x0c, 0xca, 0x4a, 0x68, 0xa8,
	0x06, 0xda, 0x35, 0x99, 0xf2, 0xfa, 0x95, 0xcc, 0xf0, 0x13, 0xb5, 0x20, 0xcf, 0xef, 0x93, 0x65,
	0xfa, 0x7f, 0xee, 0x15, 0xa6, 0xc0, 0xbd, 0xc8, 0x1e, 0x65, 0x42, 0x55, 0x25, 0xe2, 0xbf, 0x56,
	0xe5, 0xdc, 0x59, 0xd5, 0x2f, 0x50, 0x4d, 0xa4, 0x91, 0xa2, 0xbb, 0x9f, 0xd4, 0x6d, 0x24, 0x75,
	0x39, 0x7b, 0x3a, 0x23, 0x6c, 0x6c, 0x44, 0xcd, 0x55, 0x2e, 0x36, 0xea, 0x80, 0x66, 0x59, 0x86,
	0x0f, 0x05, 0xe1, 0x45, 0x0d, 0x28, 0x06, 0xe4, 0xdb, 0x84, 0xb8, 0x16, 0xe1, 0x11, 0xe4, 0xcc,
	0xd8, 0x46, 0x47, 0x50, 0xb5, 0x86, 0xd8, 0x75, 0xc9, 0xa8, 0xcf, 0x7b, 0x2d, 0xc3, 0xd9, 0x48,
	0x29, 0x9e, 0x59, 0x91, 0x48, 0x6e, 0x21, 0x04, 0x39, 0x36, 0xf5, 0x88, 0xae, 0x6d, 0x66, 0x76,
	0xf2, 0x26, 0xff, 0x36, 0x18, 0xd4, 0x05, 0x41, 0x0c, 0x63, 0x3c, 0xef, 0x8f, 0xa0, 0x2a, 0x26,
	0x3a, 0x9a, 0xe4, 0x30, 0x8c, 0x8a, 0x59, 0xb1, 0x14, 0x30, 0x3a, 0x04, 0x08, 0x9c, 0x81, 0x8b,
	0xd9, 0xc4, 0x8f, 0x07, 0xed, 0xbf, 0x64, 0x1c, 0xbd, 0xe8, 0xdc, 0x54, 0xa0, 0xc6, 0xcf, 0x2c,
	0x54, 0xd4, 0x6b, 0xd1, 0x7d, 0x80, 0x28, 0x29, 0xc7, 0x96, 0x45, 0x2f, 0x49, 0xcf, 0xa9, 0x8d,
	0x9a, 0x50, 0xf4, 0x09, 0xb6, 0xfb, 0x01, 0x61, 0x8b, 0xd2, 0x5d, 0x09, 0x41, 0x3d, 0xc2, 0xd0,
	0x3e, 0x94, 0xbe, 0xfb, 0x0e, 0x23, 0x9c, 0xa0, 0xcd, 0x27, 0x14, 0x39, 0x2a, 0x64, 0x44, 0xb5,
	0xc9, 0xdd, 0xd6, 0x06, 0xbd, 0x83, 0xaa, 0x13, 0xd0, 0x11, 0x66, 0xc4, 0xee, 0xdb, 0x98, 0x61,
	0x3d, 0xcf, 0x33, 0xdc, 0x4a, 0x2a, 0x89, 0x0c, 0x9a, 0xa7, 0x12, 0xf9, 0x06, 0x33, 0x2c, 0x1e,
	0x42, 0xc5, 0x51, 0x5c, 0x8d, 0x97, 0xb0, 0x3e, 0x03, 0x49, 0x19, 0xb2, 0xba, 0x3a, 0x64, 0x15,
	0x75, 0x90, 0x7e, 0x6b, 0x50, 0x56, 0x62, 0x47, 0x3a, 0xac, 0xdc, 0x10, 0x3f, 0x70, 0xa8, 0x2b,
	0x47, 0x24, 0x32, 0xd1, 0x61, 0xbc, 0x3a, 0x44, 0x4b, 0x1e, 0xa4, 0xa4, 0x9e, 0xba, 0x34, 0x0e,
	0xe3, 0xa5, 0xa1, 0xcd, 0x27, 0xa6, 0xad, 0x8b, 0x63, 0x65, 0x5d, 0xe4, 0x38, 0xf5, 0x61, 0x1a,
	0x75, 0xce, 0xa2, 0x08, 0xbb, 0x3f, 0xa6, 0x76, 0x9f, 0xdb, 0x53, 0x3d, 0x2f, 0xba, 0x3f, 0xa6,
	0xb6, 0x78, 0x1d, 0x8d, 0xee, 0xb2, 0x3d, 0xb2, 0x9b, 0x7c, 0x99, 0xa9, 0xad, 0x56, 0xde, 0x7a,
	0x77, 0xd9, 0x06, 0x59, 0xac, 0xc7, 0xb9, 0xaa, 0xde, 0xa7, 0xe5, 0xbb, 0x63, 0x2f, 0xa9, 0x58,
	0x4f, 0xdb, 0x1d, 0x6a, 0xb3, 0xbf, 0x46, 0xbd, 0xe6, 0x97, 0x2d, 0xe8, 0x75, 0xea, 0xbc, 0xdc,
	0x29, 0xa8, 0x76, 0xa7, 0xa0, 0x06, 0x8d, 0x5e, 0x9f, 0xb0, 0x17, 0xc8, 0x6f, 0x41, 0x41, 0x8a,
	0x64, 0x93, 0x3f, 0x7b, 0x32, 0x64, 0x79, 0xba, 0xec, 0xc2, 0x0b, 0x58, 0xbb, 0xb3, 0x0e, 0xd0,
	0x2e, 0xd4, 0xe2, 0x85, 0xd0, 0x1f, 0x12, 0x6c, 0x13, 0x5f, 0xee, 0x98, 0xb5, 0xd8, 0x7f, 0xc2,
	0xdd, 0xe8, 0x1e, 0x94, 0x62, 0x97, 0xcc, 0xf3, 0xd6, 0xb1, 0xb7, 0x0d, 0x20, 0xb4, 0xcf, 0xc2,
	0x37, 0x5b, 0x86, 0x95, 0xf6, 0xc9, 0xab, 0x6e, 0xb7, 0xf3, 0xbe, 0xf6, 0x0f, 0xaa, 0x40, 0xd1,
	0xec, 0xf4, 0x3e, 0x9c, 0x9b, 0xed, 0x4e, 0x2d, 0xf3, 0xba, 0x07, 0x8f, 0xa9, 0x3f, 0x68, 0x0e,
	0xa7, 0x1e, 0xf1, 0x47, 0xc4, 0x1e, 0x10, 0xbf, 0x79, 0x85, 0x2f, 0x7d, 0xc7, 0x12, 0x3f, 0xfa,
	0x81, 0x4c, 0xed, 0xe2, 0xe9, 0xc0, 0x61, 0xc3, 0xc9, 0x65, 0x68, 0xb6, 0x14, 0x70, 0x4b, 0x80,
	0x5b, 0x02, 0x2c, 0xff, 0x8d, 0xb8, 0x2c, 0x70, 0xf3, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x08, 0xe9, 0x1f, 0x76, 0x7d, 0x08, 0x00, 0x00,
}
