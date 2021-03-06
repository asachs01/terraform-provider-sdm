// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account_grants.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// AccountGrantCreateRequest specifies what kind of AccountGrants should be registered in
// the organizations fleet.
type AccountGrantCreateRequest struct {
	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new AccountGrant.
	AccountGrant         *AccountGrant `protobuf:"bytes,2,opt,name=account_grant,json=accountGrant,proto3" json:"account_grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AccountGrantCreateRequest) Reset()         { *m = AccountGrantCreateRequest{} }
func (m *AccountGrantCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantCreateRequest) ProtoMessage()    {}
func (*AccountGrantCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{0}
}

func (m *AccountGrantCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantCreateRequest.Unmarshal(m, b)
}
func (m *AccountGrantCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantCreateRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantCreateRequest.Merge(m, src)
}
func (m *AccountGrantCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantCreateRequest.Size(m)
}
func (m *AccountGrantCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantCreateRequest proto.InternalMessageInfo

func (m *AccountGrantCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantCreateRequest) GetAccountGrant() *AccountGrant {
	if m != nil {
		return m.AccountGrant
	}
	return nil
}

// AccountGrantCreateResponse reports how the AccountGrants were created in the system.
type AccountGrantCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created AccountGrant.
	AccountGrant *AccountGrant `protobuf:"bytes,2,opt,name=account_grant,json=accountGrant,proto3" json:"account_grant,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantCreateResponse) Reset()         { *m = AccountGrantCreateResponse{} }
func (m *AccountGrantCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantCreateResponse) ProtoMessage()    {}
func (*AccountGrantCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{1}
}

func (m *AccountGrantCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantCreateResponse.Unmarshal(m, b)
}
func (m *AccountGrantCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantCreateResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantCreateResponse.Merge(m, src)
}
func (m *AccountGrantCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantCreateResponse.Size(m)
}
func (m *AccountGrantCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantCreateResponse proto.InternalMessageInfo

func (m *AccountGrantCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantCreateResponse) GetAccountGrant() *AccountGrant {
	if m != nil {
		return m.AccountGrant
	}
	return nil
}

func (m *AccountGrantCreateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrantGetRequest specifies which AccountGrant to retrieve.
type AccountGrantGetRequest struct {
	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the AccountGrant to retrieve.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGrantGetRequest) Reset()         { *m = AccountGrantGetRequest{} }
func (m *AccountGrantGetRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantGetRequest) ProtoMessage()    {}
func (*AccountGrantGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{2}
}

func (m *AccountGrantGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantGetRequest.Unmarshal(m, b)
}
func (m *AccountGrantGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantGetRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantGetRequest.Merge(m, src)
}
func (m *AccountGrantGetRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantGetRequest.Size(m)
}
func (m *AccountGrantGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantGetRequest proto.InternalMessageInfo

func (m *AccountGrantGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AccountGrantGetResponse returns a requested AccountGrant.
type AccountGrantGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested AccountGrant.
	AccountGrant *AccountGrant `protobuf:"bytes,2,opt,name=account_grant,json=accountGrant,proto3" json:"account_grant,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantGetResponse) Reset()         { *m = AccountGrantGetResponse{} }
func (m *AccountGrantGetResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantGetResponse) ProtoMessage()    {}
func (*AccountGrantGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{3}
}

func (m *AccountGrantGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantGetResponse.Unmarshal(m, b)
}
func (m *AccountGrantGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantGetResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantGetResponse.Merge(m, src)
}
func (m *AccountGrantGetResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantGetResponse.Size(m)
}
func (m *AccountGrantGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantGetResponse proto.InternalMessageInfo

func (m *AccountGrantGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantGetResponse) GetAccountGrant() *AccountGrant {
	if m != nil {
		return m.AccountGrant
	}
	return nil
}

func (m *AccountGrantGetResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrantDeleteRequest identifies a AccountGrant by ID to delete.
type AccountGrantDeleteRequest struct {
	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the AccountGrant to delete.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGrantDeleteRequest) Reset()         { *m = AccountGrantDeleteRequest{} }
func (m *AccountGrantDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantDeleteRequest) ProtoMessage()    {}
func (*AccountGrantDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{4}
}

func (m *AccountGrantDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantDeleteRequest.Unmarshal(m, b)
}
func (m *AccountGrantDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantDeleteRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantDeleteRequest.Merge(m, src)
}
func (m *AccountGrantDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantDeleteRequest.Size(m)
}
func (m *AccountGrantDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantDeleteRequest proto.InternalMessageInfo

func (m *AccountGrantDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AccountGrantDeleteResponse returns information about a AccountGrant that was deleted.
type AccountGrantDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantDeleteResponse) Reset()         { *m = AccountGrantDeleteResponse{} }
func (m *AccountGrantDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantDeleteResponse) ProtoMessage()    {}
func (*AccountGrantDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{5}
}

func (m *AccountGrantDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantDeleteResponse.Unmarshal(m, b)
}
func (m *AccountGrantDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantDeleteResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantDeleteResponse.Merge(m, src)
}
func (m *AccountGrantDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantDeleteResponse.Size(m)
}
func (m *AccountGrantDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantDeleteResponse proto.InternalMessageInfo

func (m *AccountGrantDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrantListRequest specifies criteria for retrieving a list of AccountGrants.
type AccountGrantListRequest struct {
	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGrantListRequest) Reset()         { *m = AccountGrantListRequest{} }
func (m *AccountGrantListRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantListRequest) ProtoMessage()    {}
func (*AccountGrantListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{6}
}

func (m *AccountGrantListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantListRequest.Unmarshal(m, b)
}
func (m *AccountGrantListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantListRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantListRequest.Merge(m, src)
}
func (m *AccountGrantListRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantListRequest.Size(m)
}
func (m *AccountGrantListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantListRequest proto.InternalMessageInfo

func (m *AccountGrantListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// AccountGrantListResponse returns a list of AccountGrants that meet the criteria of a
// AccountGrantListRequest.
type AccountGrantListResponse struct {
	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	AccountGrants []*AccountGrant `protobuf:"bytes,2,rep,name=account_grants,json=accountGrants,proto3" json:"account_grants,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantListResponse) Reset()         { *m = AccountGrantListResponse{} }
func (m *AccountGrantListResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantListResponse) ProtoMessage()    {}
func (*AccountGrantListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{7}
}

func (m *AccountGrantListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantListResponse.Unmarshal(m, b)
}
func (m *AccountGrantListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantListResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantListResponse.Merge(m, src)
}
func (m *AccountGrantListResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantListResponse.Size(m)
}
func (m *AccountGrantListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantListResponse proto.InternalMessageInfo

func (m *AccountGrantListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantListResponse) GetAccountGrants() []*AccountGrant {
	if m != nil {
		return m.AccountGrants
	}
	return nil
}

func (m *AccountGrantListResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
type AccountGrant struct {
	// Unique identifier of the AccountGrant.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the composite role of this AccountGrant.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The id of the attached role of this AccountGrant.
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The timestamp when the resource will be granted. Optional. Both start_at
	// and end_at must be defined together, or not defined at all.
	StartFrom *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty"`
	// The timestamp when the resource grant will expire. Optional. Both
	// start_at and end_at must be defined together, or not defined at all.
	ValidUntil           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AccountGrant) Reset()         { *m = AccountGrant{} }
func (m *AccountGrant) String() string { return proto.CompactTextString(m) }
func (*AccountGrant) ProtoMessage()    {}
func (*AccountGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{8}
}

func (m *AccountGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrant.Unmarshal(m, b)
}
func (m *AccountGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrant.Marshal(b, m, deterministic)
}
func (m *AccountGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrant.Merge(m, src)
}
func (m *AccountGrant) XXX_Size() int {
	return xxx_messageInfo_AccountGrant.Size(m)
}
func (m *AccountGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrant.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrant proto.InternalMessageInfo

func (m *AccountGrant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountGrant) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AccountGrant) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AccountGrant) GetStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.StartFrom
	}
	return nil
}

func (m *AccountGrant) GetValidUntil() *timestamp.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountGrantCreateRequest)(nil), "v1.AccountGrantCreateRequest")
	proto.RegisterType((*AccountGrantCreateResponse)(nil), "v1.AccountGrantCreateResponse")
	proto.RegisterType((*AccountGrantGetRequest)(nil), "v1.AccountGrantGetRequest")
	proto.RegisterType((*AccountGrantGetResponse)(nil), "v1.AccountGrantGetResponse")
	proto.RegisterType((*AccountGrantDeleteRequest)(nil), "v1.AccountGrantDeleteRequest")
	proto.RegisterType((*AccountGrantDeleteResponse)(nil), "v1.AccountGrantDeleteResponse")
	proto.RegisterType((*AccountGrantListRequest)(nil), "v1.AccountGrantListRequest")
	proto.RegisterType((*AccountGrantListResponse)(nil), "v1.AccountGrantListResponse")
	proto.RegisterType((*AccountGrant)(nil), "v1.AccountGrant")
}

func init() { proto.RegisterFile("account_grants.proto", fileDescriptor_5a322caa55cfb217) }

var fileDescriptor_5a322caa55cfb217 = []byte{
	// 969 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0x3a, 0xae, 0x2d, 0x4f, 0x12, 0x04, 0x43, 0x9a, 0x38, 0xdb, 0xd0, 0xae, 0xa6, 0x12,
	0x82, 0x10, 0x7b, 0x71, 0x30, 0x42, 0xf2, 0x85, 0xda, 0x8d, 0x88, 0x22, 0x05, 0x29, 0x72, 0x69,
	0x25, 0x0e, 0x68, 0x99, 0xec, 0x8e, 0x37, 0x03, 0xde, 0x9d, 0x65, 0x66, 0x6c, 0x57, 0xaa, 0x7a,
	0x41, 0xe2, 0xc8, 0xa5, 0xfc, 0x07, 0x45, 0x42, 0x88, 0x13, 0x92, 0x2f, 0x9c, 0x10, 0xaa, 0x38,
	0xf5, 0xd2, 0x03, 0x7f, 0x02, 0x5c, 0xb8, 0xe2, 0x0b, 0xe5, 0x84, 0x66, 0x7f, 0xd8, 0x3b, 0xb6,
	0x13, 0x84, 0xc2, 0x85, 0x53, 0xb2, 0xf3, 0xde, 0xfb, 0xbe, 0x37, 0xdf, 0x7c, 0xf3, 0xc6, 0x60,
	0x03, 0xbb, 0x2e, 0x1b, 0x84, 0xd2, 0xf1, 0x39, 0x0e, 0xa5, 0xa8, 0x47, 0x9c, 0x49, 0x06, 0x0b,
	0xc3, 0x86, 0xb9, 0xe3, 0x33, 0xe6, 0xf7, 0x89, 0x8d, 0x23, 0x6a, 0xe3, 0x30, 0x64, 0x12, 0x4b,
	0xca, 0xc2, 0x34, 0xc3, 0xdc, 0x8b, 0xff, 0xb8, 0x35, 0x9f, 0x84, 0x35, 0x31, 0xc2, 0xbe, 0x4f,
	0xb8, 0xcd, 0xa2, 0x38, 0x63, 0x49, 0xf6, 0x8d, 0x14, 0x2b, 0xfe, 0x3a, 0x1d, 0xf4, 0x6c, 0x49,
	0x03, 0x22, 0x24, 0x0e, 0xa2, 0x34, 0x61, 0x3d, 0xad, 0x4d, 0x3f, 0x81, 0x88, 0x88, 0x9b, 0xfc,
	0x8f, 0xbe, 0x34, 0xc0, 0x76, 0x3b, 0x69, 0xf2, 0x50, 0xf5, 0x78, 0x9b, 0x13, 0x2c, 0x49, 0x97,
	0x7c, 0x36, 0x20, 0x42, 0xc2, 0x1a, 0x28, 0x06, 0x44, 0xe2, 0xaa, 0x61, 0x19, 0xaf, 0xad, 0xee,
	0x6f, 0xd7, 0x87, 0x8d, 0xba, 0x96, 0xf0, 0x3e, 0x91, 0xd8, 0xc3, 0x12, 0x77, 0xe3, 0x34, 0xd8,
	0x06, 0xeb, 0xda, 0x86, 0xab, 0x85, 0xb8, 0xee, 0x45, 0x55, 0x97, 0x27, 0xe9, 0x80, 0x3f, 0xfe,
	0x1c, 0x97, 0xaf, 0x7c, 0x3f, 0x19, 0x97, 0x8d, 0xee, 0x1a, 0xce, 0x45, 0xd0, 0xef, 0x06, 0x30,
	0x97, 0xf5, 0x23, 0x22, 0x16, 0x0a, 0x02, 0x5b, 0x5a, 0x43, 0x66, 0xbe, 0xa1, 0x24, 0x23, 0xeb,
	0x48, 0xa3, 0xf8, 0xaf, 0xba, 0x83, 0xb7, 0x00, 0xe0, 0x58, 0x12, 0xa7, 0x4f, 0x03, 0x2a, 0xab,
	0x2b, 0x71, 0xfd, 0x55, 0x55, 0xdf, 0xc5, 0x92, 0x1c, 0xab, 0xc5, 0xa5, 0xfc, 0x15, 0x9e, 0x85,
	0x5b, 0xe0, 0x2f, 0xb5, 0xfc, 0xad, 0x5a, 0x46, 0x1f, 0x83, 0xcd, 0x3c, 0xef, 0x21, 0x91, 0x99,
	0xee, 0xbb, 0xda, 0x36, 0x37, 0x15, 0xc3, 0x2c, 0x3a, 0x27, 0xba, 0x09, 0x0a, 0xd4, 0x8b, 0xf7,
	0x52, 0xd1, 0x48, 0x0b, 0xd4, 0x43, 0xbf, 0x1a, 0x60, 0x6b, 0x81, 0x22, 0x95, 0xf2, 0x1d, 0x8d,
	0x63, 0x6b, 0xca, 0xf1, 0xff, 0xd3, 0xb1, 0xa7, 0x5b, 0xf8, 0x80, 0xf4, 0xc9, 0x85, 0x16, 0xd6,
	0x12, 0xfe, 0x85, 0x9a, 0xdf, 0xcc, 0x79, 0x33, 0xc3, 0x39, 0xdf, 0x9b, 0x7a, 0xc6, 0x05, 0x9a,
	0xea, 0x82, 0x14, 0x2e, 0x29, 0xc8, 0x27, 0xfa, 0xa9, 0x1f, 0x53, 0x31, 0x75, 0xd6, 0x1b, 0x8b,
	0xa7, 0x9e, 0x0b, 0xcf, 0x89, 0x81, 0x40, 0xa9, 0x47, 0xfb, 0x92, 0xf0, 0x25, 0x82, 0xa4, 0x11,
	0xf4, 0xcc, 0x00, 0xd5, 0x45, 0xb2, 0x54, 0x92, 0x3d, 0x8d, 0xad, 0x3a, 0x63, 0xd3, 0x05, 0x49,
	0xe9, 0x6e, 0x83, 0x17, 0xf4, 0x79, 0x59, 0x2d, 0x58, 0x2b, 0x17, 0x38, 0xeb, 0x87, 0xb8, 0x91,
	0xf5, 0xbc, 0xb3, 0xc4, 0xe5, 0xad, 0x85, 0x7e, 0x2e, 0x82, 0xb5, 0x3c, 0x1b, 0x7c, 0x33, 0xf6,
	0x84, 0x11, 0x4b, 0x60, 0xa9, 0x9a, 0x6b, 0x8f, 0x27, 0xe3, 0x72, 0xe1, 0xe8, 0x20, 0x2e, 0x7d,
	0x3a, 0x19, 0x97, 0xc1, 0x09, 0xe1, 0x01, 0x15, 0x82, 0xb2, 0x50, 0x39, 0x05, 0x9e, 0x80, 0x55,
	0x4e, 0x04, 0x1b, 0x70, 0x97, 0x38, 0x53, 0x3b, 0xd9, 0xaa, 0x74, 0x57, 0x95, 0xae, 0x1d, 0x60,
	0x89, 0x93, 0x70, 0x0a, 0xf2, 0xd3, 0x14, 0x69, 0x16, 0xea, 0x82, 0x0c, 0xe3, 0xc8, 0x83, 0x1d,
	0x00, 0x32, 0x6d, 0xa8, 0x17, 0x6f, 0xab, 0xd2, 0xb9, 0xa9, 0x00, 0xaf, 0x2b, 0xc0, 0xd2, 0x5d,
	0x41, 0xf8, 0x3c, 0x54, 0x51, 0x2d, 0x76, 0x2b, 0x69, 0xd9, 0x91, 0x07, 0xef, 0x00, 0x20, 0x24,
	0xe6, 0xd2, 0xe9, 0x71, 0x16, 0x54, 0x8b, 0xa9, 0x4d, 0x93, 0xc7, 0xa3, 0x9e, 0x3d, 0x1e, 0xf5,
	0x0f, 0xb2, 0xc7, 0xa3, 0x53, 0x55, 0xf8, 0x2f, 0x2b, 0xfc, 0xca, 0x1d, 0x55, 0xf7, 0x1e, 0x67,
	0x41, 0xaa, 0x96, 0xc8, 0xbe, 0xe1, 0x3d, 0xb0, 0x3a, 0xc4, 0x7d, 0xea, 0x39, 0x83, 0x50, 0xd2,
	0x7e, 0xf5, 0xca, 0x3f, 0xa2, 0x6e, 0x2b, 0xd4, 0x0d, 0x85, 0x0a, 0xee, 0xa9, 0xc2, 0xbb, 0xaa,
	0x2e, 0x81, 0x05, 0xc3, 0xe9, 0x42, 0xeb, 0x47, 0xe3, 0x51, 0xfb, 0x60, 0xbf, 0x03, 0x6f, 0x3d,
	0xb0, 0x10, 0xf5, 0x50, 0xcb, 0x42, 0xd8, 0xaf, 0xed, 0x37, 0x9b, 0x68, 0xcf, 0x42, 0x39, 0x6d,
	0x55, 0x80, 0x8b, 0x5a, 0xa3, 0xd1, 0x50, 0x81, 0x99, 0x44, 0xa8, 0x85, 0x70, 0xad, 0xd9, 0x6c,
	0xa2, 0x87, 0xea, 0x5e, 0xf0, 0xc7, 0xfa, 0x21, 0xc5, 0xd7, 0xe4, 0xc9, 0x64, 0x5c, 0xfe, 0x50,
	0x05, 0x5e, 0x95, 0x3d, 0x87, 0xdc, 0xc7, 0x41, 0xd4, 0x27, 0xc2, 0xd6, 0x1c, 0xe8, 0x64, 0x4c,
	0x75, 0x79, 0x5f, 0x7e, 0x37, 0x19, 0x97, 0x5f, 0x3f, 0x3f, 0x55, 0x59, 0xca, 0x99, 0x65, 0xef,
	0x7f, 0x5d, 0x04, 0xeb, 0x6d, 0xcd, 0x9a, 0xcf, 0x0c, 0x50, 0x4a, 0x5e, 0x2b, 0xf8, 0xca, 0xbc,
	0xa5, 0xb5, 0x67, 0xd5, 0xbc, 0x7e, 0x5e, 0x38, 0xb9, 0x37, 0xe8, 0x0b, 0xe3, 0x51, 0x9b, 0x20,
	0x17, 0x58, 0xc7, 0x04, 0xf3, 0xd0, 0x3a, 0x63, 0x23, 0x4b, 0x32, 0x2b, 0xc0, 0x9f, 0x12, 0x0b,
	0x5b, 0x9a, 0x77, 0xdf, 0x3d, 0x93, 0x32, 0x12, 0x2d, 0xdb, 0x1e, 0x8d, 0x46, 0x75, 0x21, 0x39,
	0x0b, 0x7d, 0x2f, 0xa8, 0xbb, 0x2c, 0xb0, 0x3d, 0xe6, 0x8a, 0xf8, 0xe7, 0x87, 0x20, 0x7c, 0x48,
	0x5d, 0x22, 0x6c, 0xad, 0xdb, 0x9b, 0x09, 0xe5, 0xe7, 0xbf, 0xfc, 0xf6, 0x55, 0x61, 0x0b, 0x41,
	0x7b, 0xd8, 0xd0, 0xf7, 0x2b, 0x5a, 0xc6, 0x2e, 0x74, 0xc0, 0xca, 0x21, 0x91, 0xd0, 0x9c, 0x6f,
	0x77, 0xf6, 0x58, 0x99, 0xd7, 0x96, 0xc6, 0xd2, 0x7d, 0xdc, 0x88, 0xe1, 0xb7, 0xe1, 0xd6, 0x22,
	0xbc, 0xfd, 0x80, 0x7a, 0x0f, 0xe1, 0x19, 0x28, 0x25, 0x33, 0x74, 0x51, 0x31, 0x6d, 0x8a, 0x2f,
	0x2a, 0xa6, 0x8f, 0xde, 0x8c, 0x69, 0xf7, 0x5c, 0xa6, 0x8f, 0x40, 0x51, 0x8d, 0x26, 0xb8, 0xd0,
	0x6f, 0x6e, 0x3c, 0x9a, 0x3b, 0xcb, 0x83, 0x29, 0x87, 0x19, 0x73, 0x6c, 0xc0, 0x25, 0x62, 0x99,
	0x9b, 0x4f, 0x9f, 0x8f, 0xcb, 0x2f, 0x3d, 0x79, 0x3e, 0x2e, 0x6b, 0xb3, 0xa5, 0xf3, 0x36, 0xd8,
	0x71, 0x59, 0x30, 0x3b, 0x19, 0x1c, 0x51, 0xc5, 0x11, 0xf5, 0x07, 0xc1, 0x29, 0x0d, 0xfd, 0xce,
	0x55, 0xed, 0x50, 0x4e, 0xd2, 0xe5, 0xd3, 0x52, 0x7c, 0xb1, 0xde, 0xfa, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x17, 0xdd, 0x30, 0x61, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountGrantsClient is the client API for AccountGrants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountGrantsClient interface {
	// Create registers a new AccountGrant.
	Create(ctx context.Context, in *AccountGrantCreateRequest, opts ...grpc.CallOption) (*AccountGrantCreateResponse, error)
	// Get reads one AccountGrant by ID.
	Get(ctx context.Context, in *AccountGrantGetRequest, opts ...grpc.CallOption) (*AccountGrantGetResponse, error)
	// Delete removes a AccountGrant by ID.
	Delete(ctx context.Context, in *AccountGrantDeleteRequest, opts ...grpc.CallOption) (*AccountGrantDeleteResponse, error)
	// List gets a list of AccountGrants matching a given set of criteria.
	List(ctx context.Context, in *AccountGrantListRequest, opts ...grpc.CallOption) (*AccountGrantListResponse, error)
}

type accountGrantsClient struct {
	cc *grpc.ClientConn
}

func NewAccountGrantsClient(cc *grpc.ClientConn) AccountGrantsClient {
	return &accountGrantsClient{cc}
}

func (c *accountGrantsClient) Create(ctx context.Context, in *AccountGrantCreateRequest, opts ...grpc.CallOption) (*AccountGrantCreateResponse, error) {
	out := new(AccountGrantCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGrantsClient) Get(ctx context.Context, in *AccountGrantGetRequest, opts ...grpc.CallOption) (*AccountGrantGetResponse, error) {
	out := new(AccountGrantGetResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGrantsClient) Delete(ctx context.Context, in *AccountGrantDeleteRequest, opts ...grpc.CallOption) (*AccountGrantDeleteResponse, error) {
	out := new(AccountGrantDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGrantsClient) List(ctx context.Context, in *AccountGrantListRequest, opts ...grpc.CallOption) (*AccountGrantListResponse, error) {
	out := new(AccountGrantListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountGrantsServer is the server API for AccountGrants service.
type AccountGrantsServer interface {
	// Create registers a new AccountGrant.
	Create(context.Context, *AccountGrantCreateRequest) (*AccountGrantCreateResponse, error)
	// Get reads one AccountGrant by ID.
	Get(context.Context, *AccountGrantGetRequest) (*AccountGrantGetResponse, error)
	// Delete removes a AccountGrant by ID.
	Delete(context.Context, *AccountGrantDeleteRequest) (*AccountGrantDeleteResponse, error)
	// List gets a list of AccountGrants matching a given set of criteria.
	List(context.Context, *AccountGrantListRequest) (*AccountGrantListResponse, error)
}

// UnimplementedAccountGrantsServer can be embedded to have forward compatible implementations.
type UnimplementedAccountGrantsServer struct {
}

func (*UnimplementedAccountGrantsServer) Create(ctx context.Context, req *AccountGrantCreateRequest) (*AccountGrantCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccountGrantsServer) Get(ctx context.Context, req *AccountGrantGetRequest) (*AccountGrantGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccountGrantsServer) Delete(ctx context.Context, req *AccountGrantDeleteRequest) (*AccountGrantDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedAccountGrantsServer) List(ctx context.Context, req *AccountGrantListRequest) (*AccountGrantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterAccountGrantsServer(s *grpc.Server, srv AccountGrantsServer) {
	s.RegisterService(&_AccountGrants_serviceDesc, srv)
}

func _AccountGrants_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).Create(ctx, req.(*AccountGrantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGrants_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).Get(ctx, req.(*AccountGrantGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGrants_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).Delete(ctx, req.(*AccountGrantDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGrants_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).List(ctx, req.(*AccountGrantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountGrants_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountGrants",
	HandlerType: (*AccountGrantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccountGrants_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountGrants_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountGrants_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccountGrants_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_grants.proto",
}
