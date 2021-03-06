// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recaptchaenterprise/v1beta1/recaptchaenterprise.proto

package recaptchaenterprise

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum that reprensents the types of annotations.
type AnnotateAssessmentRequest_Annotation int32

const (
	// Default unspecified type.
	AnnotateAssessmentRequest_ANNOTATION_UNSPECIFIED AnnotateAssessmentRequest_Annotation = 0
	// Provides information that the event turned out to be legitimate.
	AnnotateAssessmentRequest_LEGITIMATE AnnotateAssessmentRequest_Annotation = 1
	// Provides information that the event turned out to be fraudulent.
	AnnotateAssessmentRequest_FRAUDULENT AnnotateAssessmentRequest_Annotation = 2
)

var AnnotateAssessmentRequest_Annotation_name = map[int32]string{
	0: "ANNOTATION_UNSPECIFIED",
	1: "LEGITIMATE",
	2: "FRAUDULENT",
}

var AnnotateAssessmentRequest_Annotation_value = map[string]int32{
	"ANNOTATION_UNSPECIFIED": 0,
	"LEGITIMATE":             1,
	"FRAUDULENT":             2,
}

func (x AnnotateAssessmentRequest_Annotation) String() string {
	return proto.EnumName(AnnotateAssessmentRequest_Annotation_name, int32(x))
}

func (AnnotateAssessmentRequest_Annotation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{1, 0}
}

// LINT.IfChange(classification_reason)
// Reasons contributing to the risk analysis verdict.
type Assessment_ClassificationReason int32

const (
	// Default unspecified type.
	Assessment_CLASSIFICATION_REASON_UNSPECIFIED Assessment_ClassificationReason = 0
	// Interactions matched the behavior of an automated agent.
	Assessment_AUTOMATION Assessment_ClassificationReason = 1
	// The event originated from an illegitimate environment.
	Assessment_UNEXPECTED_ENVIRONMENT Assessment_ClassificationReason = 2
	// Traffic volume from the event source is higher than normal.
	Assessment_TOO_MUCH_TRAFFIC Assessment_ClassificationReason = 3
	// Interactions with the site were significantly different than expected
	// patterns.
	Assessment_UNEXPECTED_USAGE_PATTERNS Assessment_ClassificationReason = 4
	// Too little traffic has been received from this site thus far to generate
	// quality risk analysis.
	Assessment_LOW_CONFIDENCE_SCORE Assessment_ClassificationReason = 5
)

var Assessment_ClassificationReason_name = map[int32]string{
	0: "CLASSIFICATION_REASON_UNSPECIFIED",
	1: "AUTOMATION",
	2: "UNEXPECTED_ENVIRONMENT",
	3: "TOO_MUCH_TRAFFIC",
	4: "UNEXPECTED_USAGE_PATTERNS",
	5: "LOW_CONFIDENCE_SCORE",
}

var Assessment_ClassificationReason_value = map[string]int32{
	"CLASSIFICATION_REASON_UNSPECIFIED": 0,
	"AUTOMATION":                        1,
	"UNEXPECTED_ENVIRONMENT":            2,
	"TOO_MUCH_TRAFFIC":                  3,
	"UNEXPECTED_USAGE_PATTERNS":         4,
	"LOW_CONFIDENCE_SCORE":              5,
}

func (x Assessment_ClassificationReason) String() string {
	return proto.EnumName(Assessment_ClassificationReason_name, int32(x))
}

func (Assessment_ClassificationReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{3, 0}
}

// LINT.IfChange
// Enum that represents the types of invalid token reasons.
type TokenProperties_InvalidReason int32

const (
	// Default unspecified type.
	TokenProperties_INVALID_REASON_UNSPECIFIED TokenProperties_InvalidReason = 0
	// If the failure reason was not accounted for.
	TokenProperties_UNKNOWN_INVALID_REASON TokenProperties_InvalidReason = 1
	// The provided user verification token was malformed.
	TokenProperties_MALFORMED TokenProperties_InvalidReason = 2
	// The user verification token had expired.
	TokenProperties_EXPIRED TokenProperties_InvalidReason = 3
	// The user verification had already been seen.
	TokenProperties_DUPE TokenProperties_InvalidReason = 4
	// The user verification token did not match the provided site key.
	// This may be a configuration error (e.g. development keys used in
	// production) or end users trying to use verification tokens from other
	// sites.
	TokenProperties_SITE_MISMATCH TokenProperties_InvalidReason = 5
	// The user verification token was not present.  It is a required input.
	TokenProperties_MISSING TokenProperties_InvalidReason = 6
)

var TokenProperties_InvalidReason_name = map[int32]string{
	0: "INVALID_REASON_UNSPECIFIED",
	1: "UNKNOWN_INVALID_REASON",
	2: "MALFORMED",
	3: "EXPIRED",
	4: "DUPE",
	5: "SITE_MISMATCH",
	6: "MISSING",
}

var TokenProperties_InvalidReason_value = map[string]int32{
	"INVALID_REASON_UNSPECIFIED": 0,
	"UNKNOWN_INVALID_REASON":     1,
	"MALFORMED":                  2,
	"EXPIRED":                    3,
	"DUPE":                       4,
	"SITE_MISMATCH":              5,
	"MISSING":                    6,
}

func (x TokenProperties_InvalidReason) String() string {
	return proto.EnumName(TokenProperties_InvalidReason_name, int32(x))
}

func (TokenProperties_InvalidReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{5, 0}
}

// The create assessment request message.
type CreateAssessmentRequest struct {
	// Required. The name of the project in which the assessment will be created,
	// in the format "projects/{project_number}".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The assessment details.
	Assessment           *Assessment `protobuf:"bytes,2,opt,name=assessment,proto3" json:"assessment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateAssessmentRequest) Reset()         { *m = CreateAssessmentRequest{} }
func (m *CreateAssessmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAssessmentRequest) ProtoMessage()    {}
func (*CreateAssessmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{0}
}

func (m *CreateAssessmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAssessmentRequest.Unmarshal(m, b)
}
func (m *CreateAssessmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAssessmentRequest.Marshal(b, m, deterministic)
}
func (m *CreateAssessmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAssessmentRequest.Merge(m, src)
}
func (m *CreateAssessmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAssessmentRequest.Size(m)
}
func (m *CreateAssessmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAssessmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAssessmentRequest proto.InternalMessageInfo

func (m *CreateAssessmentRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateAssessmentRequest) GetAssessment() *Assessment {
	if m != nil {
		return m.Assessment
	}
	return nil
}

// The request message to annotate an Assessment.
type AnnotateAssessmentRequest struct {
	// Required. The resource name of the Assessment, in the format
	// "projects/{project_number}/assessments/{assessment_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The annotation that will be assigned to the Event.
	Annotation           AnnotateAssessmentRequest_Annotation `protobuf:"varint,2,opt,name=annotation,proto3,enum=google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentRequest_Annotation" json:"annotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AnnotateAssessmentRequest) Reset()         { *m = AnnotateAssessmentRequest{} }
func (m *AnnotateAssessmentRequest) String() string { return proto.CompactTextString(m) }
func (*AnnotateAssessmentRequest) ProtoMessage()    {}
func (*AnnotateAssessmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{1}
}

func (m *AnnotateAssessmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotateAssessmentRequest.Unmarshal(m, b)
}
func (m *AnnotateAssessmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotateAssessmentRequest.Marshal(b, m, deterministic)
}
func (m *AnnotateAssessmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotateAssessmentRequest.Merge(m, src)
}
func (m *AnnotateAssessmentRequest) XXX_Size() int {
	return xxx_messageInfo_AnnotateAssessmentRequest.Size(m)
}
func (m *AnnotateAssessmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotateAssessmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotateAssessmentRequest proto.InternalMessageInfo

func (m *AnnotateAssessmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnnotateAssessmentRequest) GetAnnotation() AnnotateAssessmentRequest_Annotation {
	if m != nil {
		return m.Annotation
	}
	return AnnotateAssessmentRequest_ANNOTATION_UNSPECIFIED
}

// Empty response for AnnotateAssessment.
type AnnotateAssessmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotateAssessmentResponse) Reset()         { *m = AnnotateAssessmentResponse{} }
func (m *AnnotateAssessmentResponse) String() string { return proto.CompactTextString(m) }
func (*AnnotateAssessmentResponse) ProtoMessage()    {}
func (*AnnotateAssessmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{2}
}

func (m *AnnotateAssessmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotateAssessmentResponse.Unmarshal(m, b)
}
func (m *AnnotateAssessmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotateAssessmentResponse.Marshal(b, m, deterministic)
}
func (m *AnnotateAssessmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotateAssessmentResponse.Merge(m, src)
}
func (m *AnnotateAssessmentResponse) XXX_Size() int {
	return xxx_messageInfo_AnnotateAssessmentResponse.Size(m)
}
func (m *AnnotateAssessmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotateAssessmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotateAssessmentResponse proto.InternalMessageInfo

// A recaptcha assessment resource.
type Assessment struct {
	// Output only. The resource name for the Assessment in the format
	// "projects/{project_number}/assessments/{assessment_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The event being assessed.
	Event *Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	// Output only. Legitimate event score from 0.0 to 1.0.
	// (1.0 means very likely legitimate traffic while 0.0 means very likely
	// non-legitimate traffic).
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	// Output only. Properties of the provided event token.
	TokenProperties *TokenProperties `protobuf:"bytes,4,opt,name=token_properties,json=tokenProperties,proto3" json:"token_properties,omitempty"`
	// Output only. Reasons contributing to the risk analysis verdict.
	Reasons              []Assessment_ClassificationReason `protobuf:"varint,5,rep,packed,name=reasons,proto3,enum=google.cloud.recaptchaenterprise.v1beta1.Assessment_ClassificationReason" json:"reasons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Assessment) Reset()         { *m = Assessment{} }
func (m *Assessment) String() string { return proto.CompactTextString(m) }
func (*Assessment) ProtoMessage()    {}
func (*Assessment) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{3}
}

func (m *Assessment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assessment.Unmarshal(m, b)
}
func (m *Assessment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assessment.Marshal(b, m, deterministic)
}
func (m *Assessment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assessment.Merge(m, src)
}
func (m *Assessment) XXX_Size() int {
	return xxx_messageInfo_Assessment.Size(m)
}
func (m *Assessment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assessment.DiscardUnknown(m)
}

var xxx_messageInfo_Assessment proto.InternalMessageInfo

func (m *Assessment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Assessment) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Assessment) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Assessment) GetTokenProperties() *TokenProperties {
	if m != nil {
		return m.TokenProperties
	}
	return nil
}

func (m *Assessment) GetReasons() []Assessment_ClassificationReason {
	if m != nil {
		return m.Reasons
	}
	return nil
}

type Event struct {
	// Required. The user response token provided by the reCAPTCHA client-side integration
	// on your site.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Required. The site key that was used to invoke reCAPTCHA on your site and generate
	// the token.
	SiteKey              string   `protobuf:"bytes,2,opt,name=site_key,json=siteKey,proto3" json:"site_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{4}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Event) GetSiteKey() string {
	if m != nil {
		return m.SiteKey
	}
	return ""
}

type TokenProperties struct {
	// Whether the provided user response token is valid.
	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	// Reason associated with the response when valid = false.
	InvalidReason TokenProperties_InvalidReason `protobuf:"varint,2,opt,name=invalid_reason,json=invalidReason,proto3,enum=google.cloud.recaptchaenterprise.v1beta1.TokenProperties_InvalidReason" json:"invalid_reason,omitempty"`
	// The timestamp corresponding to the generation of the token.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The hostname of the page on which the token was generated.
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Action name provided at token generation.
	Action               string   `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenProperties) Reset()         { *m = TokenProperties{} }
func (m *TokenProperties) String() string { return proto.CompactTextString(m) }
func (*TokenProperties) ProtoMessage()    {}
func (*TokenProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b3da91aae1398d, []int{5}
}

func (m *TokenProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenProperties.Unmarshal(m, b)
}
func (m *TokenProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenProperties.Marshal(b, m, deterministic)
}
func (m *TokenProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenProperties.Merge(m, src)
}
func (m *TokenProperties) XXX_Size() int {
	return xxx_messageInfo_TokenProperties.Size(m)
}
func (m *TokenProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenProperties.DiscardUnknown(m)
}

var xxx_messageInfo_TokenProperties proto.InternalMessageInfo

func (m *TokenProperties) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *TokenProperties) GetInvalidReason() TokenProperties_InvalidReason {
	if m != nil {
		return m.InvalidReason
	}
	return TokenProperties_INVALID_REASON_UNSPECIFIED
}

func (m *TokenProperties) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *TokenProperties) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *TokenProperties) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentRequest_Annotation", AnnotateAssessmentRequest_Annotation_name, AnnotateAssessmentRequest_Annotation_value)
	proto.RegisterEnum("google.cloud.recaptchaenterprise.v1beta1.Assessment_ClassificationReason", Assessment_ClassificationReason_name, Assessment_ClassificationReason_value)
	proto.RegisterEnum("google.cloud.recaptchaenterprise.v1beta1.TokenProperties_InvalidReason", TokenProperties_InvalidReason_name, TokenProperties_InvalidReason_value)
	proto.RegisterType((*CreateAssessmentRequest)(nil), "google.cloud.recaptchaenterprise.v1beta1.CreateAssessmentRequest")
	proto.RegisterType((*AnnotateAssessmentRequest)(nil), "google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentRequest")
	proto.RegisterType((*AnnotateAssessmentResponse)(nil), "google.cloud.recaptchaenterprise.v1beta1.AnnotateAssessmentResponse")
	proto.RegisterType((*Assessment)(nil), "google.cloud.recaptchaenterprise.v1beta1.Assessment")
	proto.RegisterType((*Event)(nil), "google.cloud.recaptchaenterprise.v1beta1.Event")
	proto.RegisterType((*TokenProperties)(nil), "google.cloud.recaptchaenterprise.v1beta1.TokenProperties")
}

func init() {
	proto.RegisterFile("google/cloud/recaptchaenterprise/v1beta1/recaptchaenterprise.proto", fileDescriptor_72b3da91aae1398d)
}

var fileDescriptor_72b3da91aae1398d = []byte{
	// 1146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xc6, 0xf9, 0xd1, 0x1f, 0x53, 0xb5, 0xf5, 0x8e, 0xaa, 0x6d, 0x1a, 0x2d, 0x4b, 0xd7, 0x12,
	0x28, 0xea, 0x6e, 0x6d, 0xb5, 0x0b, 0x48, 0x74, 0xc5, 0x61, 0xe2, 0x4c, 0x52, 0xab, 0x89, 0x13,
	0xd9, 0x4e, 0x77, 0x81, 0x4a, 0x96, 0xeb, 0x4e, 0x13, 0xb3, 0x89, 0x6d, 0x3c, 0x6e, 0xab, 0x55,
	0xd5, 0x0b, 0xff, 0x01, 0xe2, 0xca, 0x09, 0x89, 0x0b, 0x7f, 0x01, 0x57, 0xae, 0x2b, 0x71, 0x81,
	0xdb, 0x4a, 0x48, 0x3d, 0x70, 0xe2, 0xc2, 0x15, 0x71, 0x40, 0xc8, 0x63, 0x27, 0x4e, 0x8a, 0x97,
	0xed, 0x96, 0xbd, 0x79, 0xe6, 0xbd, 0xf9, 0xde, 0x7c, 0xf3, 0xbe, 0xf7, 0x9e, 0x41, 0xb5, 0xe7,
	0x79, 0xbd, 0x01, 0x91, 0xec, 0x81, 0x77, 0x72, 0x24, 0x05, 0xc4, 0xb6, 0xfc, 0xd0, 0xee, 0x5b,
	0xc4, 0x0d, 0x49, 0xe0, 0x07, 0x0e, 0x25, 0xd2, 0xe9, 0xd6, 0x21, 0x09, 0xad, 0xad, 0x2c, 0x9b,
	0xe8, 0x07, 0x5e, 0xe8, 0xc1, 0x4a, 0x8c, 0x21, 0x32, 0x0c, 0x31, 0xcb, 0x2f, 0xc1, 0x28, 0xdf,
	0x49, 0xa2, 0x59, 0xbe, 0x23, 0x59, 0xae, 0xeb, 0x85, 0x56, 0xe8, 0x78, 0x2e, 0x8d, 0x71, 0xca,
	0xab, 0x13, 0x56, 0x7b, 0xe0, 0x10, 0x37, 0x4c, 0x0c, 0xef, 0x4c, 0x18, 0x8e, 0x1d, 0x32, 0x38,
	0x32, 0x0f, 0x49, 0xdf, 0x3a, 0x75, 0xbc, 0x20, 0x71, 0x58, 0x9b, 0x70, 0x08, 0x08, 0xf5, 0x4e,
	0x02, 0x9b, 0x5c, 0x39, 0xcb, 0x56, 0x87, 0x27, 0xc7, 0x52, 0xe8, 0x0c, 0x09, 0x0d, 0xad, 0xa1,
	0x1f, 0x3b, 0x08, 0x3f, 0x72, 0x60, 0x55, 0x0e, 0x88, 0x15, 0x12, 0x44, 0x29, 0xa1, 0x74, 0x48,
	0xdc, 0x50, 0x23, 0x5f, 0x9c, 0x10, 0x1a, 0xc2, 0x3d, 0x30, 0xe3, 0x5b, 0x01, 0x71, 0xc3, 0x12,
	0xb7, 0xce, 0x55, 0xe6, 0xab, 0x0f, 0x2f, 0x51, 0xee, 0x2f, 0xb4, 0x09, 0xee, 0x33, 0xaa, 0xa3,
	0x48, 0x43, 0xcb, 0xb5, 0x7a, 0x24, 0x10, 0xe3, 0x48, 0x96, 0xef, 0x50, 0xd1, 0xf6, 0x86, 0x52,
	0x27, 0xf0, 0x3e, 0x27, 0x76, 0xa8, 0x25, 0x10, 0xf0, 0x13, 0x00, 0xac, 0x71, 0x84, 0x52, 0x6e,
	0x9d, 0xab, 0x2c, 0x6c, 0xbf, 0x2f, 0x5e, 0xf7, 0xed, 0xc4, 0xf4, 0x76, 0xd5, 0xfc, 0x25, 0xca,
	0x69, 0x13, 0x60, 0xc2, 0x77, 0x39, 0xb0, 0x86, 0xe2, 0xf7, 0xcc, 0x60, 0xa1, 0x80, 0x82, 0x6b,
	0x0d, 0x49, 0xc2, 0xe1, 0x03, 0xc6, 0x41, 0x02, 0x9b, 0x59, 0xc1, 0xae, 0x50, 0x98, 0xc0, 0x62,
	0x10, 0x90, 0x02, 0x90, 0xe6, 0x8d, 0x71, 0x58, 0xda, 0x56, 0x5f, 0x83, 0xc3, 0xcb, 0xee, 0x38,
	0xb2, 0x38, 0x9e, 0x3b, 0x62, 0x37, 0xde, 0x10, 0x76, 0x01, 0x48, 0xcd, 0xb0, 0x0c, 0x6e, 0x23,
	0x55, 0x6d, 0x1b, 0xc8, 0x50, 0xda, 0xaa, 0xd9, 0x55, 0xf5, 0x0e, 0x96, 0x95, 0xba, 0x82, 0x6b,
	0xfc, 0x5b, 0x70, 0x09, 0x80, 0x26, 0x6e, 0x28, 0x86, 0xd2, 0x42, 0x06, 0xe6, 0xb9, 0x68, 0x5d,
	0xd7, 0x50, 0xb7, 0xd6, 0x6d, 0x62, 0xd5, 0xe0, 0x73, 0xc2, 0x1d, 0x50, 0xce, 0xba, 0x02, 0xf5,
	0x3d, 0x97, 0x12, 0xe1, 0xcf, 0x02, 0x00, 0xe9, 0x36, 0x5c, 0x9d, 0x7a, 0xb6, 0xfc, 0x25, 0xca,
	0x27, 0x8f, 0x80, 0x41, 0x91, 0x9c, 0xa6, 0x39, 0x94, 0xae, 0xcf, 0x1f, 0x47, 0xc7, 0xb4, 0xf8,
	0x34, 0x5c, 0x03, 0x45, 0x6a, 0x7b, 0x01, 0x29, 0xe5, 0xd7, 0xb9, 0x4a, 0x2e, 0x0e, 0x10, 0xef,
	0x40, 0x07, 0xf0, 0xa1, 0xf7, 0x94, 0xb8, 0xa6, 0x1f, 0x78, 0x3e, 0x09, 0x42, 0x87, 0xd0, 0x52,
	0x81, 0x05, 0xfb, 0xe8, 0xfa, 0xc1, 0x8c, 0x08, 0xa1, 0x33, 0x06, 0x88, 0x03, 0x2c, 0x87, 0xd3,
	0xbb, 0xb0, 0x0f, 0x66, 0x03, 0x62, 0x51, 0xcf, 0xa5, 0xa5, 0xe2, 0x7a, 0xbe, 0xb2, 0xb4, 0xad,
	0xdc, 0x44, 0x92, 0xa2, 0x3c, 0xb0, 0x28, 0x75, 0x8e, 0x1d, 0x9b, 0x25, 0x49, 0x63, 0x88, 0x71,
	0xc4, 0x11, 0xbc, 0xf0, 0x03, 0x07, 0x56, 0xb2, 0xdc, 0xe0, 0xbb, 0xe0, 0x9e, 0xdc, 0x44, 0xba,
	0xae, 0xd4, 0x15, 0x39, 0xce, 0xaa, 0x86, 0x91, 0x9e, 0x95, 0x5c, 0xd4, 0x35, 0xda, 0x2d, 0xe6,
	0xc2, 0x73, 0x91, 0x10, 0xba, 0x2a, 0x7e, 0xd2, 0xc1, 0xb2, 0x81, 0x6b, 0x26, 0x56, 0xf7, 0x15,
	0xad, 0xad, 0xb6, 0x58, 0xa2, 0xe1, 0x0a, 0xe0, 0x8d, 0x76, 0xdb, 0x6c, 0x75, 0xe5, 0x5d, 0xd3,
	0xd0, 0x50, 0xbd, 0xae, 0xc8, 0x7c, 0x1e, 0xbe, 0x0d, 0xd6, 0x26, 0x4e, 0x74, 0x75, 0xd4, 0xc0,
	0x66, 0x07, 0x19, 0x06, 0xd6, 0x54, 0x9d, 0x2f, 0xc0, 0x12, 0x58, 0x69, 0xb6, 0x1f, 0x9b, 0x72,
	0x5b, 0xad, 0x2b, 0x35, 0xac, 0xca, 0xd8, 0xd4, 0xe5, 0xb6, 0x86, 0xf9, 0xe2, 0x8e, 0xf9, 0x3b,
	0x3a, 0x78, 0xcd, 0x82, 0x81, 0xf7, 0xfd, 0xb8, 0xfe, 0xa9, 0x74, 0x9e, 0x7c, 0x5d, 0x48, 0x69,
	0xc9, 0x52, 0xe9, 0x3c, 0x5d, 0x5c, 0x08, 0x55, 0x50, 0xc4, 0x23, 0x51, 0xb0, 0x0c, 0xa5, 0xaa,
	0xcb, 0x69, 0xf1, 0x0e, 0xbc, 0x0b, 0xe6, 0xa8, 0x13, 0x12, 0xf3, 0x29, 0x79, 0xc6, 0x94, 0x97,
	0x58, 0x67, 0xa3, 0xcd, 0x3d, 0xf2, 0x4c, 0xf8, 0x26, 0x0f, 0x96, 0xaf, 0xe4, 0x1c, 0xae, 0x80,
	0xe2, 0xa9, 0x35, 0x70, 0x8e, 0x18, 0xdc, 0x9c, 0x16, 0x2f, 0xa0, 0x0b, 0x96, 0x1c, 0x97, 0x7d,
	0x9a, 0x71, 0x72, 0x92, 0x4a, 0x6e, 0xdc, 0x58, 0x5c, 0xa2, 0x12, 0xe3, 0xc5, 0x19, 0xd5, 0x16,
	0x9d, 0xc9, 0x25, 0x7c, 0x04, 0x16, 0x6c, 0xd6, 0x61, 0xcd, 0xa8, 0xf9, 0x32, 0xbd, 0x2f, 0x6c,
	0x97, 0x47, 0xc1, 0x46, 0x9d, 0x59, 0x34, 0x46, 0x9d, 0x59, 0x03, 0xb1, 0x7b, 0xb4, 0x01, 0xcb,
	0x60, 0xae, 0xef, 0xd1, 0x90, 0x95, 0x62, 0x54, 0x03, 0xf3, 0xda, 0x78, 0x0d, 0x6f, 0x83, 0x19,
	0xcb, 0x66, 0xad, 0xa8, 0xc8, 0x2c, 0xc9, 0x4a, 0xf8, 0x8a, 0x03, 0x8b, 0x53, 0x37, 0x82, 0x77,
	0x41, 0x59, 0x51, 0xf7, 0x51, 0x53, 0xa9, 0x65, 0x8b, 0x8b, 0x89, 0x69, 0x4f, 0x6d, 0x3f, 0x56,
	0xcd, 0x69, 0x3f, 0x9e, 0x83, 0x8b, 0x60, 0xbe, 0x85, 0x9a, 0xf5, 0xb6, 0xd6, 0xc2, 0x35, 0x3e,
	0x07, 0x17, 0xc0, 0x2c, 0x7e, 0xd2, 0x51, 0x34, 0x5c, 0xe3, 0xf3, 0x70, 0x0e, 0x14, 0x6a, 0xdd,
	0x0e, 0xe6, 0x0b, 0xf0, 0x16, 0x58, 0xd4, 0x15, 0x03, 0x9b, 0x2d, 0x45, 0x6f, 0x21, 0x43, 0xde,
	0xe5, 0x8b, 0x91, 0x67, 0x4b, 0xd1, 0x75, 0x45, 0x6d, 0xf0, 0x33, 0xdb, 0x3f, 0x15, 0xc0, 0x3d,
	0x6d, 0xf4, 0xa2, 0x78, 0xfc, 0xa2, 0x3a, 0x09, 0x4e, 0x1d, 0x9b, 0xec, 0x6f, 0x55, 0xa3, 0x77,
	0x85, 0xbf, 0x72, 0x80, 0xbf, 0x3a, 0x8d, 0x20, 0xba, 0x7e, 0x5e, 0x5e, 0x32, 0xc9, 0xca, 0x37,
	0x1a, 0x34, 0x42, 0xe7, 0x05, 0xba, 0x15, 0x4f, 0xaf, 0x07, 0xa9, 0x5a, 0xbf, 0xfc, 0xe5, 0xb7,
	0xaf, 0x73, 0x1f, 0x0a, 0x95, 0xf1, 0x9f, 0xc1, 0x79, 0xec, 0xf2, 0xf1, 0x58, 0xf0, 0x1b, 0x53,
	0x4a, 0xdf, 0x99, 0x98, 0x54, 0xf0, 0x0f, 0x0e, 0xc0, 0x7f, 0xb7, 0x60, 0x28, 0xbf, 0x81, 0x19,
	0x52, 0xae, 0xfd, 0x3f, 0x90, 0x64, 0x0a, 0xa8, 0x2f, 0xd0, 0x72, 0x24, 0xae, 0x07, 0xe9, 0x04,
	0x4a, 0x18, 0x6f, 0xa5, 0x8c, 0x23, 0x87, 0x09, 0xbe, 0x53, 0x85, 0xbd, 0x71, 0xb1, 0x93, 0x1c,
	0x25, 0x3b, 0xdc, 0x46, 0x79, 0xff, 0x39, 0x12, 0x5e, 0xdd, 0x39, 0x7e, 0x46, 0x62, 0x3f, 0x0c,
	0x7d, 0xba, 0x23, 0x49, 0x67, 0x67, 0x67, 0x57, 0xdb, 0x8a, 0x75, 0x12, 0xf6, 0xe3, 0x5f, 0xb4,
	0x4d, 0x7f, 0x60, 0x85, 0xc7, 0x5e, 0x30, 0xac, 0xfe, 0xcd, 0x81, 0xf7, 0x6c, 0x6f, 0x38, 0xe2,
	0xfc, 0x1f, 0x6c, 0xab, 0xa5, 0x0c, 0xdd, 0x75, 0xa2, 0xaa, 0xeb, 0x70, 0x9f, 0x7e, 0x96, 0x9c,
	0xef, 0x79, 0x03, 0xcb, 0xed, 0x89, 0x5e, 0xd0, 0x93, 0x7a, 0xc4, 0x65, 0x35, 0x29, 0xa5, 0xb7,
	0x78, 0xf5, 0xff, 0xe1, 0xa3, 0x0c, 0xdb, 0xb7, 0xb9, 0x42, 0x43, 0xd6, 0xf0, 0xf7, 0xb9, 0x4a,
	0x23, 0x0e, 0x22, 0xb3, 0xc4, 0x64, 0xdc, 0x46, 0x4c, 0xf4, 0xff, 0x7c, 0xe4, 0x7a, 0xc0, 0x5c,
	0x0f, 0x32, 0x5c, 0x0f, 0xf6, 0xe3, 0x80, 0x87, 0x33, 0xec, 0x92, 0x0f, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0xc1, 0x01, 0xca, 0xc3, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecaptchaEnterpriseServiceV1Beta1Client is the client API for RecaptchaEnterpriseServiceV1Beta1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecaptchaEnterpriseServiceV1Beta1Client interface {
	// Creates an Assessment of the likelihood an event is legitimate.
	CreateAssessment(ctx context.Context, in *CreateAssessmentRequest, opts ...grpc.CallOption) (*Assessment, error)
	// Annotates a previously created Assessment to provide additional information
	// on whether the event turned out to be authentic or fradulent.
	AnnotateAssessment(ctx context.Context, in *AnnotateAssessmentRequest, opts ...grpc.CallOption) (*AnnotateAssessmentResponse, error)
}

type recaptchaEnterpriseServiceV1Beta1Client struct {
	cc *grpc.ClientConn
}

func NewRecaptchaEnterpriseServiceV1Beta1Client(cc *grpc.ClientConn) RecaptchaEnterpriseServiceV1Beta1Client {
	return &recaptchaEnterpriseServiceV1Beta1Client{cc}
}

func (c *recaptchaEnterpriseServiceV1Beta1Client) CreateAssessment(ctx context.Context, in *CreateAssessmentRequest, opts ...grpc.CallOption) (*Assessment, error) {
	out := new(Assessment)
	err := c.cc.Invoke(ctx, "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/CreateAssessment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recaptchaEnterpriseServiceV1Beta1Client) AnnotateAssessment(ctx context.Context, in *AnnotateAssessmentRequest, opts ...grpc.CallOption) (*AnnotateAssessmentResponse, error) {
	out := new(AnnotateAssessmentResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/AnnotateAssessment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecaptchaEnterpriseServiceV1Beta1Server is the server API for RecaptchaEnterpriseServiceV1Beta1 service.
type RecaptchaEnterpriseServiceV1Beta1Server interface {
	// Creates an Assessment of the likelihood an event is legitimate.
	CreateAssessment(context.Context, *CreateAssessmentRequest) (*Assessment, error)
	// Annotates a previously created Assessment to provide additional information
	// on whether the event turned out to be authentic or fradulent.
	AnnotateAssessment(context.Context, *AnnotateAssessmentRequest) (*AnnotateAssessmentResponse, error)
}

func RegisterRecaptchaEnterpriseServiceV1Beta1Server(s *grpc.Server, srv RecaptchaEnterpriseServiceV1Beta1Server) {
	s.RegisterService(&_RecaptchaEnterpriseServiceV1Beta1_serviceDesc, srv)
}

func _RecaptchaEnterpriseServiceV1Beta1_CreateAssessment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssessmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).CreateAssessment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/CreateAssessment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).CreateAssessment(ctx, req.(*CreateAssessmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecaptchaEnterpriseServiceV1Beta1_AnnotateAssessment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnotateAssessmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).AnnotateAssessment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1/AnnotateAssessment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecaptchaEnterpriseServiceV1Beta1Server).AnnotateAssessment(ctx, req.(*AnnotateAssessmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecaptchaEnterpriseServiceV1Beta1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recaptchaenterprise.v1beta1.RecaptchaEnterpriseServiceV1Beta1",
	HandlerType: (*RecaptchaEnterpriseServiceV1Beta1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAssessment",
			Handler:    _RecaptchaEnterpriseServiceV1Beta1_CreateAssessment_Handler,
		},
		{
			MethodName: "AnnotateAssessment",
			Handler:    _RecaptchaEnterpriseServiceV1Beta1_AnnotateAssessment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recaptchaenterprise/v1beta1/recaptchaenterprise.proto",
}
