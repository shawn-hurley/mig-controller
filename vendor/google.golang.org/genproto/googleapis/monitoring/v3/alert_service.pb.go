// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/alert_service.proto

package monitoring

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// The protocol for the `CreateAlertPolicy` request.
type CreateAlertPolicyRequest struct {
	// The project in which to create the alerting policy. The format is
	// `projects/[PROJECT_ID]`.
	//
	// Note that this field names the parent container in which the alerting
	// policy will be written, not the name of the created policy. The alerting
	// policy that is returned will have a name that contains a normalized
	// representation of this name as a prefix but adds a suffix of the form
	// `/alertPolicies/[POLICY_ID]`, identifying the policy in the container.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The requested alerting policy. You should omit the `name` field in this
	// policy. The name will be returned in the new policy, including
	// a new [ALERT_POLICY_ID] value.
	AlertPolicy          *AlertPolicy `protobuf:"bytes,2,opt,name=alert_policy,json=alertPolicy,proto3" json:"alert_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateAlertPolicyRequest) Reset()         { *m = CreateAlertPolicyRequest{} }
func (m *CreateAlertPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAlertPolicyRequest) ProtoMessage()    {}
func (*CreateAlertPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45362b2a456d1bf, []int{0}
}

func (m *CreateAlertPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAlertPolicyRequest.Unmarshal(m, b)
}
func (m *CreateAlertPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAlertPolicyRequest.Marshal(b, m, deterministic)
}
func (m *CreateAlertPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAlertPolicyRequest.Merge(m, src)
}
func (m *CreateAlertPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAlertPolicyRequest.Size(m)
}
func (m *CreateAlertPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAlertPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAlertPolicyRequest proto.InternalMessageInfo

func (m *CreateAlertPolicyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAlertPolicyRequest) GetAlertPolicy() *AlertPolicy {
	if m != nil {
		return m.AlertPolicy
	}
	return nil
}

// The protocol for the `GetAlertPolicy` request.
type GetAlertPolicyRequest struct {
	// The alerting policy to retrieve. The format is
	//
	//     projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAlertPolicyRequest) Reset()         { *m = GetAlertPolicyRequest{} }
func (m *GetAlertPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*GetAlertPolicyRequest) ProtoMessage()    {}
func (*GetAlertPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45362b2a456d1bf, []int{1}
}

func (m *GetAlertPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlertPolicyRequest.Unmarshal(m, b)
}
func (m *GetAlertPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlertPolicyRequest.Marshal(b, m, deterministic)
}
func (m *GetAlertPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlertPolicyRequest.Merge(m, src)
}
func (m *GetAlertPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_GetAlertPolicyRequest.Size(m)
}
func (m *GetAlertPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlertPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlertPolicyRequest proto.InternalMessageInfo

func (m *GetAlertPolicyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The protocol for the `ListAlertPolicies` request.
type ListAlertPoliciesRequest struct {
	// The project whose alert policies are to be listed. The format is
	//
	//     projects/[PROJECT_ID]
	//
	// Note that this field names the parent container in which the alerting
	// policies to be listed are stored. To retrieve a single alerting policy
	// by name, use the
	// [GetAlertPolicy][google.monitoring.v3.AlertPolicyService.GetAlertPolicy]
	// operation, instead.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// If provided, this field specifies the criteria that must be met by
	// alert policies to be included in the response.
	//
	// For more details, see [sorting and
	// filtering](/monitoring/api/v3/sorting-and-filtering).
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// A comma-separated list of fields by which to sort the result. Supports
	// the same set of field references as the `filter` field. Entries can be
	// prefixed with a minus sign to sort by the field in descending order.
	//
	// For more details, see [sorting and
	// filtering](/monitoring/api/v3/sorting-and-filtering).
	OrderBy string `protobuf:"bytes,6,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// The maximum number of results to return in a single response.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return more results from the previous method call.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAlertPoliciesRequest) Reset()         { *m = ListAlertPoliciesRequest{} }
func (m *ListAlertPoliciesRequest) String() string { return proto.CompactTextString(m) }
func (*ListAlertPoliciesRequest) ProtoMessage()    {}
func (*ListAlertPoliciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45362b2a456d1bf, []int{2}
}

func (m *ListAlertPoliciesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAlertPoliciesRequest.Unmarshal(m, b)
}
func (m *ListAlertPoliciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAlertPoliciesRequest.Marshal(b, m, deterministic)
}
func (m *ListAlertPoliciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAlertPoliciesRequest.Merge(m, src)
}
func (m *ListAlertPoliciesRequest) XXX_Size() int {
	return xxx_messageInfo_ListAlertPoliciesRequest.Size(m)
}
func (m *ListAlertPoliciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAlertPoliciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAlertPoliciesRequest proto.InternalMessageInfo

func (m *ListAlertPoliciesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListAlertPoliciesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListAlertPoliciesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListAlertPoliciesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAlertPoliciesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The protocol for the `ListAlertPolicies` response.
type ListAlertPoliciesResponse struct {
	// The returned alert policies.
	AlertPolicies []*AlertPolicy `protobuf:"bytes,3,rep,name=alert_policies,json=alertPolicies,proto3" json:"alert_policies,omitempty"`
	// If there might be more results than were returned, then this field is set
	// to a non-empty value. To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAlertPoliciesResponse) Reset()         { *m = ListAlertPoliciesResponse{} }
func (m *ListAlertPoliciesResponse) String() string { return proto.CompactTextString(m) }
func (*ListAlertPoliciesResponse) ProtoMessage()    {}
func (*ListAlertPoliciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45362b2a456d1bf, []int{3}
}

func (m *ListAlertPoliciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAlertPoliciesResponse.Unmarshal(m, b)
}
func (m *ListAlertPoliciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAlertPoliciesResponse.Marshal(b, m, deterministic)
}
func (m *ListAlertPoliciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAlertPoliciesResponse.Merge(m, src)
}
func (m *ListAlertPoliciesResponse) XXX_Size() int {
	return xxx_messageInfo_ListAlertPoliciesResponse.Size(m)
}
func (m *ListAlertPoliciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAlertPoliciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAlertPoliciesResponse proto.InternalMessageInfo

func (m *ListAlertPoliciesResponse) GetAlertPolicies() []*AlertPolicy {
	if m != nil {
		return m.AlertPolicies
	}
	return nil
}

func (m *ListAlertPoliciesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The protocol for the `UpdateAlertPolicy` request.
type UpdateAlertPolicyRequest struct {
	// Optional. A list of alerting policy field names. If this field is not
	// empty, each listed field in the existing alerting policy is set to the
	// value of the corresponding field in the supplied policy (`alert_policy`),
	// or to the field's default value if the field is not in the supplied
	// alerting policy.  Fields not listed retain their previous value.
	//
	// Examples of valid field masks include `display_name`, `documentation`,
	// `documentation.content`, `documentation.mime_type`, `user_labels`,
	// `user_label.nameofkey`, `enabled`, `conditions`, `combiner`, etc.
	//
	// If this field is empty, then the supplied alerting policy replaces the
	// existing policy. It is the same as deleting the existing policy and
	// adding the supplied policy, except for the following:
	//
	// +   The new policy will have the same `[ALERT_POLICY_ID]` as the former
	//     policy. This gives you continuity with the former policy in your
	//     notifications and incidents.
	// +   Conditions in the new policy will keep their former `[CONDITION_ID]` if
	//     the supplied condition includes the `name` field with that
	//     `[CONDITION_ID]`. If the supplied condition omits the `name` field,
	//     then a new `[CONDITION_ID]` is created.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Required. The updated alerting policy or the updated values for the
	// fields listed in `update_mask`.
	// If `update_mask` is not empty, any fields in this policy that are
	// not in `update_mask` are ignored.
	AlertPolicy          *AlertPolicy `protobuf:"bytes,3,opt,name=alert_policy,json=alertPolicy,proto3" json:"alert_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateAlertPolicyRequest) Reset()         { *m = UpdateAlertPolicyRequest{} }
func (m *UpdateAlertPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAlertPolicyRequest) ProtoMessage()    {}
func (*UpdateAlertPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45362b2a456d1bf, []int{4}
}

func (m *UpdateAlertPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAlertPolicyRequest.Unmarshal(m, b)
}
func (m *UpdateAlertPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAlertPolicyRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAlertPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAlertPolicyRequest.Merge(m, src)
}
func (m *UpdateAlertPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAlertPolicyRequest.Size(m)
}
func (m *UpdateAlertPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAlertPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAlertPolicyRequest proto.InternalMessageInfo

func (m *UpdateAlertPolicyRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateAlertPolicyRequest) GetAlertPolicy() *AlertPolicy {
	if m != nil {
		return m.AlertPolicy
	}
	return nil
}

// The protocol for the `DeleteAlertPolicy` request.
type DeleteAlertPolicyRequest struct {
	// The alerting policy to delete. The format is:
	//
	//     projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	//
	// For more information, see [AlertPolicy][google.monitoring.v3.AlertPolicy].
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAlertPolicyRequest) Reset()         { *m = DeleteAlertPolicyRequest{} }
func (m *DeleteAlertPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAlertPolicyRequest) ProtoMessage()    {}
func (*DeleteAlertPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c45362b2a456d1bf, []int{5}
}

func (m *DeleteAlertPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAlertPolicyRequest.Unmarshal(m, b)
}
func (m *DeleteAlertPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAlertPolicyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAlertPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAlertPolicyRequest.Merge(m, src)
}
func (m *DeleteAlertPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAlertPolicyRequest.Size(m)
}
func (m *DeleteAlertPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAlertPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAlertPolicyRequest proto.InternalMessageInfo

func (m *DeleteAlertPolicyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAlertPolicyRequest)(nil), "google.monitoring.v3.CreateAlertPolicyRequest")
	proto.RegisterType((*GetAlertPolicyRequest)(nil), "google.monitoring.v3.GetAlertPolicyRequest")
	proto.RegisterType((*ListAlertPoliciesRequest)(nil), "google.monitoring.v3.ListAlertPoliciesRequest")
	proto.RegisterType((*ListAlertPoliciesResponse)(nil), "google.monitoring.v3.ListAlertPoliciesResponse")
	proto.RegisterType((*UpdateAlertPolicyRequest)(nil), "google.monitoring.v3.UpdateAlertPolicyRequest")
	proto.RegisterType((*DeleteAlertPolicyRequest)(nil), "google.monitoring.v3.DeleteAlertPolicyRequest")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/alert_service.proto", fileDescriptor_c45362b2a456d1bf)
}

var fileDescriptor_c45362b2a456d1bf = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x4f, 0x13, 0x4d,
	0x18, 0xce, 0xb6, 0xfc, 0x1c, 0x3e, 0xf8, 0xd2, 0xc9, 0xf7, 0xe1, 0xb2, 0x68, 0x52, 0x6b, 0x50,
	0x02, 0xb2, 0x93, 0xb4, 0x27, 0x21, 0x9a, 0x50, 0x50, 0x38, 0x48, 0xd2, 0x14, 0xe5, 0x60, 0x48,
	0x9a, 0xa1, 0x7d, 0xbb, 0x8c, 0x6c, 0x77, 0xd6, 0x9d, 0x29, 0x58, 0x0c, 0x17, 0x6f, 0x9a, 0x18,
	0x0f, 0x9e, 0x3d, 0x78, 0xc4, 0x93, 0x7f, 0x07, 0x47, 0xbd, 0x79, 0xf6, 0x0f, 0x31, 0x3b, 0xbb,
	0x65, 0xb7, 0x74, 0x37, 0xac, 0xde, 0x3a, 0xf3, 0x3e, 0x33, 0xcf, 0x33, 0xcf, 0xfb, 0xbc, 0x5b,
	0xb4, 0x68, 0x71, 0x6e, 0xd9, 0x40, 0x3a, 0xdc, 0x61, 0x92, 0x7b, 0xcc, 0xb1, 0xc8, 0x71, 0x85,
	0x50, 0x1b, 0x3c, 0xd9, 0x10, 0xe0, 0x1d, 0xb3, 0x26, 0x98, 0xae, 0xc7, 0x25, 0xc7, 0xff, 0x05,
	0x48, 0x33, 0x42, 0x9a, 0xc7, 0x15, 0xe3, 0x66, 0x78, 0x9e, 0xba, 0x8c, 0x50, 0xc7, 0xe1, 0x92,
	0x4a, 0xc6, 0x1d, 0x11, 0x9c, 0x31, 0x8a, 0xe9, 0xb7, 0x87, 0x88, 0xf9, 0x10, 0xa1, 0x56, 0x07,
	0xdd, 0x36, 0x81, 0x8e, 0x2b, 0x7b, 0x57, 0x8e, 0x5f, 0x16, 0xdb, 0x0c, 0xec, 0x56, 0xa3, 0x43,
	0xc5, 0x51, 0x88, 0xb8, 0x11, 0xa3, 0x6f, 0xda, 0x0c, 0x9c, 0xf0, 0xde, 0x92, 0x44, 0xfa, 0x86,
	0x07, 0x54, 0xc2, 0xba, 0x4f, 0x56, 0xe3, 0x36, 0x6b, 0xf6, 0xea, 0xf0, 0xaa, 0x0b, 0x42, 0x62,
	0x8c, 0x46, 0x1c, 0xda, 0x01, 0x3d, 0x5f, 0xd4, 0x16, 0x27, 0xeb, 0xea, 0x37, 0xde, 0x44, 0xff,
	0x04, 0x8f, 0x76, 0x15, 0x54, 0xcf, 0x15, 0xb5, 0xc5, 0xa9, 0xf2, 0x6d, 0x33, 0xe9, 0xd1, 0x66,
	0xfc, 0xce, 0x29, 0x1a, 0x2d, 0x4a, 0xcb, 0xe8, 0xff, 0x2d, 0x90, 0xd9, 0x28, 0x4b, 0x9f, 0x35,
	0xa4, 0x3f, 0x65, 0x22, 0x06, 0x67, 0x20, 0xae, 0x1e, 0x18, 0x89, 0x69, 0x9c, 0x45, 0x63, 0x6d,
	0x66, 0x4b, 0xf0, 0xf4, 0x51, 0xb5, 0x1b, 0xae, 0xf0, 0x1c, 0x9a, 0xe0, 0x5e, 0x0b, 0xbc, 0xc6,
	0x41, 0x4f, 0x1f, 0x53, 0x95, 0x71, 0xb5, 0xae, 0xf6, 0xf0, 0x3c, 0x9a, 0x74, 0xa9, 0x05, 0x0d,
	0xc1, 0x4e, 0x41, 0xbd, 0x69, 0xb4, 0x3e, 0xe1, 0x6f, 0xec, 0xb2, 0x53, 0xc0, 0xb7, 0x10, 0x52,
	0x45, 0xc9, 0x8f, 0xc0, 0x09, 0xa5, 0x29, 0xf8, 0x33, 0x7f, 0xa3, 0xf4, 0x41, 0x43, 0x73, 0x09,
	0xfa, 0x84, 0xcb, 0x1d, 0x01, 0x78, 0x1b, 0xcd, 0xc4, 0x0c, 0x63, 0x20, 0xf4, 0x7c, 0x31, 0x9f,
	0xcd, 0xb2, 0x69, 0x1a, 0xbf, 0x11, 0xdf, 0x45, 0xff, 0x3a, 0xf0, 0x5a, 0x36, 0x62, 0x5a, 0x72,
	0x4a, 0xcb, 0xb4, 0xbf, 0x5d, 0xbb, 0xd4, 0xe3, 0xfb, 0xf5, 0xdc, 0x6d, 0x25, 0xf7, 0x74, 0x0d,
	0x4d, 0x75, 0x55, 0x4d, 0xa5, 0x23, 0x6c, 0x9f, 0xd1, 0xd7, 0xd2, 0x0f, 0x90, 0xf9, 0xc4, 0x0f,
	0xd0, 0x0e, 0x15, 0x47, 0x75, 0x14, 0xc0, 0xfd, 0xdf, 0x43, 0xcd, 0xcf, 0xff, 0x55, 0xf3, 0x4d,
	0xa4, 0x6f, 0x82, 0x0d, 0x59, 0x23, 0x57, 0xfe, 0x39, 0x8e, 0x70, 0x0c, 0xba, 0x1b, 0x4c, 0x1b,
	0xfe, 0xa2, 0xa1, 0xc2, 0x90, 0xed, 0xd8, 0x4c, 0x16, 0x93, 0x96, 0x1f, 0x83, 0x64, 0xc6, 0x07,
	0xfd, 0x2c, 0x2d, 0xbf, 0xfd, 0xf1, 0xeb, 0x53, 0x6e, 0x01, 0xdf, 0xf1, 0x27, 0xf4, 0x8d, 0x2f,
	0xf0, 0xa1, 0xeb, 0xf1, 0x97, 0xd0, 0x94, 0x82, 0x2c, 0x9d, 0x91, 0xc1, 0x96, 0x7d, 0xd4, 0xd0,
	0xcc, 0x60, 0xd0, 0xf1, 0x72, 0x32, 0x61, 0xe2, 0x38, 0x18, 0xd7, 0x5b, 0x5b, 0x5a, 0x51, 0x7a,
	0xee, 0xe1, 0x85, 0x24, 0x3d, 0x83, 0x72, 0xc8, 0xd2, 0x99, 0x72, 0x6d, 0x68, 0xe0, 0xd3, 0x5c,
	0x4b, 0xfb, 0x32, 0x64, 0xd1, 0xf5, 0x40, 0xe9, 0xaa, 0x94, 0xb2, 0xf8, 0xb4, 0x3a, 0x10, 0x2b,
	0xfc, 0x5e, 0x43, 0x85, 0xa1, 0x84, 0xa4, 0x69, 0x4c, 0x8b, 0x92, 0x31, 0x3b, 0x14, 0xea, 0xc7,
	0xfe, 0x27, 0xb3, 0x6f, 0xd8, 0x52, 0x46, 0xc3, 0xbe, 0x69, 0xa8, 0x30, 0x34, 0x4d, 0x69, 0x62,
	0xd2, 0xc6, 0x2e, 0x8b, 0x61, 0xdb, 0x4a, 0x57, 0xb5, 0x5c, 0x56, 0xba, 0xe2, 0x86, 0x98, 0xd7,
	0x89, 0x1c, 0xf4, 0xcf, 0x38, 0xd7, 0x2e, 0xd6, 0xe7, 0x62, 0x44, 0x01, 0x35, 0x75, 0x99, 0x30,
	0x9b, 0xbc, 0xf3, 0x7d, 0xfd, 0x9d, 0x76, 0x28, 0xa5, 0x2b, 0x56, 0x09, 0x39, 0x39, 0x39, 0xb9,
	0x52, 0x25, 0xb4, 0x2b, 0x0f, 0x49, 0xd3, 0xe6, 0xdd, 0xd6, 0x8a, 0x6b, 0x53, 0xd9, 0xe6, 0x5e,
	0xe7, 0xfe, 0x75, 0xf0, 0x88, 0xeb, 0x0f, 0xa0, 0xa6, 0x07, 0xb4, 0x55, 0x3d, 0xd7, 0x90, 0xde,
	0xe4, 0x9d, 0x44, 0x7b, 0xaa, 0x05, 0xe5, 0x4f, 0x38, 0xf0, 0x35, 0xbf, 0x8d, 0x35, 0xed, 0xc5,
	0xa3, 0x10, 0x6a, 0x71, 0x9b, 0x3a, 0x96, 0xc9, 0x3d, 0x8b, 0x58, 0xe0, 0xa8, 0x26, 0x93, 0x88,
	0x71, 0xf0, 0xaf, 0x74, 0x2d, 0x5a, 0x7d, 0xcd, 0x19, 0x5b, 0xc1, 0x05, 0x1b, 0xfe, 0x23, 0xcd,
	0x9d, 0x88, 0x71, 0xaf, 0x72, 0xd1, 0x2f, 0xee, 0xab, 0xe2, 0x7e, 0x54, 0xdc, 0xdf, 0xab, 0x1c,
	0x8c, 0x29, 0x92, 0xca, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0xce, 0x66, 0xda, 0x0c, 0x08,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertPolicyServiceClient is the client API for AlertPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertPolicyServiceClient interface {
	// Lists the existing alerting policies for the project.
	ListAlertPolicies(ctx context.Context, in *ListAlertPoliciesRequest, opts ...grpc.CallOption) (*ListAlertPoliciesResponse, error)
	// Gets a single alerting policy.
	GetAlertPolicy(ctx context.Context, in *GetAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error)
	// Creates a new alerting policy.
	CreateAlertPolicy(ctx context.Context, in *CreateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error)
	// Deletes an alerting policy.
	DeleteAlertPolicy(ctx context.Context, in *DeleteAlertPolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates an alerting policy. You can either replace the entire policy with
	// a new one or replace only certain fields in the current alerting policy by
	// specifying the fields to be updated via `updateMask`. Returns the
	// updated alerting policy.
	UpdateAlertPolicy(ctx context.Context, in *UpdateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error)
}

type alertPolicyServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlertPolicyServiceClient(cc *grpc.ClientConn) AlertPolicyServiceClient {
	return &alertPolicyServiceClient{cc}
}

func (c *alertPolicyServiceClient) ListAlertPolicies(ctx context.Context, in *ListAlertPoliciesRequest, opts ...grpc.CallOption) (*ListAlertPoliciesResponse, error) {
	out := new(ListAlertPoliciesResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.AlertPolicyService/ListAlertPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) GetAlertPolicy(ctx context.Context, in *GetAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error) {
	out := new(AlertPolicy)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.AlertPolicyService/GetAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) CreateAlertPolicy(ctx context.Context, in *CreateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error) {
	out := new(AlertPolicy)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.AlertPolicyService/CreateAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) DeleteAlertPolicy(ctx context.Context, in *DeleteAlertPolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.AlertPolicyService/DeleteAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) UpdateAlertPolicy(ctx context.Context, in *UpdateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error) {
	out := new(AlertPolicy)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.AlertPolicyService/UpdateAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertPolicyServiceServer is the server API for AlertPolicyService service.
type AlertPolicyServiceServer interface {
	// Lists the existing alerting policies for the project.
	ListAlertPolicies(context.Context, *ListAlertPoliciesRequest) (*ListAlertPoliciesResponse, error)
	// Gets a single alerting policy.
	GetAlertPolicy(context.Context, *GetAlertPolicyRequest) (*AlertPolicy, error)
	// Creates a new alerting policy.
	CreateAlertPolicy(context.Context, *CreateAlertPolicyRequest) (*AlertPolicy, error)
	// Deletes an alerting policy.
	DeleteAlertPolicy(context.Context, *DeleteAlertPolicyRequest) (*empty.Empty, error)
	// Updates an alerting policy. You can either replace the entire policy with
	// a new one or replace only certain fields in the current alerting policy by
	// specifying the fields to be updated via `updateMask`. Returns the
	// updated alerting policy.
	UpdateAlertPolicy(context.Context, *UpdateAlertPolicyRequest) (*AlertPolicy, error)
}

func RegisterAlertPolicyServiceServer(s *grpc.Server, srv AlertPolicyServiceServer) {
	s.RegisterService(&_AlertPolicyService_serviceDesc, srv)
}

func _AlertPolicyService_ListAlertPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).ListAlertPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.AlertPolicyService/ListAlertPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).ListAlertPolicies(ctx, req.(*ListAlertPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_GetAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).GetAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.AlertPolicyService/GetAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).GetAlertPolicy(ctx, req.(*GetAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_CreateAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).CreateAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.AlertPolicyService/CreateAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).CreateAlertPolicy(ctx, req.(*CreateAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_DeleteAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).DeleteAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.AlertPolicyService/DeleteAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).DeleteAlertPolicy(ctx, req.(*DeleteAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_UpdateAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).UpdateAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.AlertPolicyService/UpdateAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).UpdateAlertPolicy(ctx, req.(*UpdateAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertPolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.AlertPolicyService",
	HandlerType: (*AlertPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlertPolicies",
			Handler:    _AlertPolicyService_ListAlertPolicies_Handler,
		},
		{
			MethodName: "GetAlertPolicy",
			Handler:    _AlertPolicyService_GetAlertPolicy_Handler,
		},
		{
			MethodName: "CreateAlertPolicy",
			Handler:    _AlertPolicyService_CreateAlertPolicy_Handler,
		},
		{
			MethodName: "DeleteAlertPolicy",
			Handler:    _AlertPolicyService_DeleteAlertPolicy_Handler,
		},
		{
			MethodName: "UpdateAlertPolicy",
			Handler:    _AlertPolicyService_UpdateAlertPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/monitoring/v3/alert_service.proto",
}