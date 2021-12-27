// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthAction int32

const (
	AuthAction_ONLY_ERAD  AuthAction = 0
	AuthAction_READ_WRITE AuthAction = 1
)

var AuthAction_name = map[int32]string{
	0: "ONLY_ERAD",
	1: "READ_WRITE",
}
var AuthAction_value = map[string]int32{
	"ONLY_ERAD":  0,
	"READ_WRITE": 1,
}

func (x AuthAction) String() string {
	return proto.EnumName(AuthAction_name, int32(x))
}
func (AuthAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{0}
}

type ResourceType int32

const (
	ResourceType_Namespaces   ResourceType = 0
	ResourceType_Services     ResourceType = 1
	ResourceType_ConfigGroups ResourceType = 2
)

var ResourceType_name = map[int32]string{
	0: "Namespaces",
	1: "Services",
	2: "ConfigGroups",
}
var ResourceType_value = map[string]int32{
	"Namespaces":   0,
	"Services":     1,
	"ConfigGroups": 2,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{1}
}

type User struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             *wrappers.StringValue `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Source               *wrappers.StringValue `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	AuthToken            *wrappers.StringValue `protobuf:"bytes,6,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,8,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,9,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *User) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *User) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *User) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *User) GetSource() *wrappers.StringValue {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *User) GetAuthToken() *wrappers.StringValue {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

func (m *User) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *User) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *User) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type UserToken struct {
	UserId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthToken            *wrappers.StringValue `protobuf:"bytes,2,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserToken) Reset()         { *m = UserToken{} }
func (m *UserToken) String() string { return proto.CompactTextString(m) }
func (*UserToken) ProtoMessage()    {}
func (*UserToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{1}
}
func (m *UserToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToken.Unmarshal(m, b)
}
func (m *UserToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToken.Marshal(b, m, deterministic)
}
func (dst *UserToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToken.Merge(dst, src)
}
func (m *UserToken) XXX_Size() int {
	return xxx_messageInfo_UserToken.Size(m)
}
func (m *UserToken) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToken.DiscardUnknown(m)
}

var xxx_messageInfo_UserToken proto.InternalMessageInfo

func (m *UserToken) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *UserToken) GetAuthToken() *wrappers.StringValue {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

type UserGroup struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	AuthToken            *wrappers.StringValue `protobuf:"bytes,4,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,6,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,7,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserGroup) Reset()         { *m = UserGroup{} }
func (m *UserGroup) String() string { return proto.CompactTextString(m) }
func (*UserGroup) ProtoMessage()    {}
func (*UserGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{2}
}
func (m *UserGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroup.Unmarshal(m, b)
}
func (m *UserGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroup.Marshal(b, m, deterministic)
}
func (dst *UserGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroup.Merge(dst, src)
}
func (m *UserGroup) XXX_Size() int {
	return xxx_messageInfo_UserGroup.Size(m)
}
func (m *UserGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroup proto.InternalMessageInfo

func (m *UserGroup) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserGroup) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *UserGroup) GetAuthToken() *wrappers.StringValue {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

func (m *UserGroup) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *UserGroup) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *UserGroup) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type UserGroupRelation struct {
	GroupId              *wrappers.StringValue   `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserIds              []*wrappers.StringValue `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	Join                 *wrappers.BoolValue     `protobuf:"bytes,3,opt,name=join,proto3" json:"join,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UserGroupRelation) Reset()         { *m = UserGroupRelation{} }
func (m *UserGroupRelation) String() string { return proto.CompactTextString(m) }
func (*UserGroupRelation) ProtoMessage()    {}
func (*UserGroupRelation) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{3}
}
func (m *UserGroupRelation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroupRelation.Unmarshal(m, b)
}
func (m *UserGroupRelation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroupRelation.Marshal(b, m, deterministic)
}
func (dst *UserGroupRelation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupRelation.Merge(dst, src)
}
func (m *UserGroupRelation) XXX_Size() int {
	return xxx_messageInfo_UserGroupRelation.Size(m)
}
func (m *UserGroupRelation) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupRelation.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupRelation proto.InternalMessageInfo

func (m *UserGroupRelation) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *UserGroupRelation) GetUserIds() []*wrappers.StringValue {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *UserGroupRelation) GetJoin() *wrappers.BoolValue {
	if m != nil {
		return m.Join
	}
	return nil
}

type Principal struct {
	Users                []*wrappers.StringValue `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Groups               []*wrappers.StringValue `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{4}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (dst *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(dst, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetUsers() []*wrappers.StringValue {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Principal) GetGroups() []*wrappers.StringValue {
	if m != nil {
		return m.Groups
	}
	return nil
}

type Resource struct {
	Namespaces           []*wrappers.StringValue `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Services             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	ConfigGroups         []*wrappers.StringValue `protobuf:"bytes,3,rep,name=config_groups,json=configGroups,proto3" json:"config_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{5}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetNamespaces() []*wrappers.StringValue {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *Resource) GetServices() []*wrappers.StringValue {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Resource) GetConfigGroups() []*wrappers.StringValue {
	if m != nil {
		return m.ConfigGroups
	}
	return nil
}

type AuthStrategy struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Principal            *Principal            `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`
	Resource             *Resource             `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	Action               AuthAction            `protobuf:"varint,5,opt,name=action,proto3,enum=v1.AuthAction" json:"action,omitempty"`
	Comment              *wrappers.StringValue `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	Ctime                *wrappers.StringValue `protobuf:"bytes,8,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrappers.StringValue `protobuf:"bytes,9,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AuthStrategy) Reset()         { *m = AuthStrategy{} }
func (m *AuthStrategy) String() string { return proto.CompactTextString(m) }
func (*AuthStrategy) ProtoMessage()    {}
func (*AuthStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_09bd8f67b01777f7, []int{6}
}
func (m *AuthStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStrategy.Unmarshal(m, b)
}
func (m *AuthStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStrategy.Marshal(b, m, deterministic)
}
func (dst *AuthStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStrategy.Merge(dst, src)
}
func (m *AuthStrategy) XXX_Size() int {
	return xxx_messageInfo_AuthStrategy.Size(m)
}
func (m *AuthStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStrategy proto.InternalMessageInfo

func (m *AuthStrategy) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AuthStrategy) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AuthStrategy) GetPrincipal() *Principal {
	if m != nil {
		return m.Principal
	}
	return nil
}

func (m *AuthStrategy) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *AuthStrategy) GetAction() AuthAction {
	if m != nil {
		return m.Action
	}
	return AuthAction_ONLY_ERAD
}

func (m *AuthStrategy) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *AuthStrategy) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AuthStrategy) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *AuthStrategy) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "v1.User")
	proto.RegisterType((*UserToken)(nil), "v1.UserToken")
	proto.RegisterType((*UserGroup)(nil), "v1.UserGroup")
	proto.RegisterType((*UserGroupRelation)(nil), "v1.UserGroupRelation")
	proto.RegisterType((*Principal)(nil), "v1.Principal")
	proto.RegisterType((*Resource)(nil), "v1.Resource")
	proto.RegisterType((*AuthStrategy)(nil), "v1.AuthStrategy")
	proto.RegisterEnum("v1.AuthAction", AuthAction_name, AuthAction_value)
	proto.RegisterEnum("v1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_09bd8f67b01777f7) }

var fileDescriptor_auth_09bd8f67b01777f7 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x6b, 0x3b, 0xf1, 0xcf, 0xc5, 0x89, 0xcc, 0xac, 0xac, 0x0a, 0xa1, 0x2a, 0x0b, 0x54,
	0xb5, 0xc8, 0xa5, 0xe1, 0x57, 0x02, 0x21, 0x05, 0x1a, 0x55, 0x91, 0x50, 0x41, 0x4e, 0x00, 0xb1,
	0x8a, 0x5c, 0x67, 0xea, 0x1a, 0x12, 0x8f, 0x35, 0x63, 0x27, 0xea, 0x8a, 0xa7, 0xe2, 0x01, 0x58,
	0xf0, 0x2c, 0xbc, 0x04, 0x0b, 0x34, 0x9e, 0xb1, 0x1b, 0x60, 0xc1, 0x34, 0x52, 0xbb, 0x9d, 0xf9,
	0x8e, 0x7d, 0xef, 0x3d, 0x67, 0x2e, 0x40, 0x54, 0x16, 0xe7, 0x41, 0x4e, 0x49, 0x41, 0x90, 0xbe,
	0x3c, 0xdc, 0xbe, 0x9b, 0x10, 0x92, 0xcc, 0xf1, 0x41, 0x75, 0x72, 0x5a, 0x9e, 0x1d, 0xac, 0x68,
	0x94, 0xe7, 0x98, 0x32, 0xc1, 0xf4, 0x7e, 0x1a, 0xd0, 0x7a, 0xcf, 0x30, 0x45, 0xf7, 0x41, 0x4f,
	0x67, 0xbe, 0xb6, 0xa3, 0xed, 0xde, 0xea, 0xdf, 0x09, 0x84, 0x2a, 0xa8, 0x55, 0xc1, 0xb8, 0xa0,
	0x69, 0x96, 0x7c, 0x88, 0xe6, 0x25, 0x0e, 0xf5, 0x74, 0x86, 0x1e, 0x40, 0x2b, 0x8b, 0x16, 0xd8,
	0xd7, 0x15, 0xf8, 0x8a, 0x44, 0xcf, 0xc0, 0xce, 0x23, 0xc6, 0x56, 0x84, 0xce, 0x7c, 0x43, 0x41,
	0xd5, 0xd0, 0xa8, 0x0f, 0x6d, 0xb2, 0xca, 0x30, 0xf5, 0x5b, 0x0a, 0x32, 0x81, 0xa2, 0x47, 0x60,
	0x32, 0x52, 0xd2, 0x18, 0xfb, 0x6d, 0x05, 0x91, 0x64, 0xd1, 0x73, 0x31, 0xbe, 0x69, 0x41, 0xbe,
	0xe0, 0xcc, 0x37, 0x15, 0x94, 0x0e, 0xe7, 0x27, 0x1c, 0x47, 0x4f, 0xc0, 0x8a, 0xc9, 0x62, 0x81,
	0xb3, 0xc2, 0xb7, 0x14, 0x94, 0x35, 0xcc, 0xdb, 0x8b, 0x8b, 0x74, 0x81, 0x7d, 0x5b, 0xa5, 0xbd,
	0x0a, 0xe5, 0x9a, 0x45, 0xa5, 0x71, 0x54, 0x34, 0x15, 0xda, 0xfb, 0x0a, 0x0e, 0x37, 0x5a, 0x14,
	0xfb, 0x18, 0xac, 0x92, 0x61, 0x3a, 0x55, 0xb4, 0xdc, 0xe4, 0xf0, 0x68, 0xf6, 0xd7, 0x80, 0xf4,
	0x2b, 0x0d, 0xa8, 0xf7, 0x4b, 0x17, 0x15, 0x1c, 0x53, 0x52, 0xe6, 0xd7, 0x9e, 0xb7, 0x26, 0x35,
	0x86, 0x7a, 0x6a, 0xfe, 0x6c, 0xaf, 0xb5, 0xb1, 0xff, 0xed, 0x8d, 0xfc, 0x37, 0x37, 0xf0, 0xdf,
	0x52, 0xf7, 0xff, 0x9b, 0x06, 0xb7, 0x9b, 0xf1, 0x87, 0x78, 0x1e, 0x15, 0x29, 0xc9, 0xd0, 0x53,
	0xb0, 0x13, 0x7e, 0xa0, 0x9a, 0x04, 0xab, 0xa2, 0x47, 0x33, 0x2e, 0x94, 0x09, 0x62, 0xbe, 0xbe,
	0x63, 0xfc, 0x5f, 0x28, 0x22, 0xc4, 0x50, 0x00, 0xad, 0xcf, 0x24, 0xcd, 0xa4, 0x2f, 0xdb, 0xff,
	0x88, 0x5e, 0x11, 0x32, 0x97, 0x46, 0x72, 0xae, 0x57, 0x82, 0xf3, 0x8e, 0xa6, 0x59, 0x9c, 0xe6,
	0xd1, 0x9c, 0x37, 0xce, 0xbf, 0xc3, 0x7c, 0x4d, 0xe1, 0x97, 0x02, 0xe5, 0xbb, 0xa0, 0x2a, 0x5a,
	0xad, 0x4e, 0xc9, 0xf6, 0x7e, 0x68, 0x60, 0x87, 0x58, 0x2e, 0x86, 0x17, 0x00, 0x3c, 0x54, 0x2c,
	0x8f, 0x62, 0xac, 0xf6, 0xef, 0x35, 0x9e, 0xaf, 0x3e, 0x86, 0xe9, 0x32, 0xe5, 0x5a, 0x95, 0x12,
	0x1a, 0x1a, 0x0d, 0xa0, 0x13, 0x93, 0xec, 0x2c, 0x4d, 0xa6, 0xb2, 0x03, 0x43, 0x41, 0xee, 0x0a,
	0xc9, 0xb1, 0xe8, 0xe3, 0xbb, 0x01, 0xee, 0xa0, 0x2c, 0xce, 0xc7, 0x05, 0x8d, 0x0a, 0x9c, 0x5c,
	0x5c, 0xfb, 0xc3, 0xdb, 0x07, 0x27, 0xaf, 0xfd, 0x92, 0x26, 0x77, 0x82, 0xe5, 0x61, 0xd0, 0x98,
	0x18, 0x5e, 0xde, 0xa3, 0x5d, 0xb0, 0xa9, 0x1c, 0xb2, 0x7c, 0x6f, 0x2e, 0x67, 0xeb, 0xc1, 0x87,
	0xcd, 0x2d, 0xba, 0x07, 0x66, 0x14, 0xf3, 0xc8, 0x56, 0xaf, 0xab, 0xdb, 0xef, 0x72, 0x8e, 0x37,
	0x36, 0xa8, 0x4e, 0x43, 0x79, 0xbb, 0xfe, 0x0c, 0xcd, 0x2b, 0x3e, 0x43, 0xb1, 0x2f, 0x2c, 0xf5,
	0x7d, 0x71, 0x43, 0xab, 0x7b, 0x6f, 0x1f, 0xe0, 0xb2, 0x53, 0xd4, 0x01, 0xe7, 0xed, 0xc9, 0x9b,
	0x4f, 0xd3, 0x61, 0x38, 0x38, 0xf2, 0xb6, 0x50, 0x17, 0x20, 0x1c, 0x0e, 0x8e, 0xa6, 0x1f, 0xc3,
	0xd1, 0x64, 0xe8, 0x69, 0x7b, 0x2f, 0xc1, 0xad, 0xc7, 0x37, 0xb9, 0xc8, 0x31, 0xbf, 0x3f, 0x69,
	0xb2, 0xe8, 0x6d, 0x21, 0x17, 0xec, 0xb1, 0xcc, 0x97, 0xa7, 0x21, 0x0f, 0xdc, 0xd7, 0x6b, 0x71,
	0xf1, 0xf4, 0x53, 0xb3, 0xaa, 0xe4, 0xe1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xd1, 0x37,
	0x38, 0x4a, 0x08, 0x00, 0x00,
}
