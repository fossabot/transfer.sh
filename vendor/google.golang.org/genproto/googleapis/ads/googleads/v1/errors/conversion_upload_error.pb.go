// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/conversion_upload_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible conversion upload errors.
type ConversionUploadErrorEnum_ConversionUploadError int32

const (
	// Enum unspecified.
	ConversionUploadErrorEnum_UNSPECIFIED ConversionUploadErrorEnum_ConversionUploadError = 0
	// The received error code is not known in this version.
	ConversionUploadErrorEnum_UNKNOWN ConversionUploadErrorEnum_ConversionUploadError = 1
	// The request contained more than 2000 conversions.
	ConversionUploadErrorEnum_TOO_MANY_CONVERSIONS_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 2
	// The specified gclid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_GCLID ConversionUploadErrorEnum_ConversionUploadError = 3
	// The specified conversion_date_time is before the event time
	// associated with the given gclid.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_GCLID ConversionUploadErrorEnum_ConversionUploadError = 4
	// The click associated with the given gclid is either too old to be
	// imported or occurred outside of the click through lookback window for the
	// specified conversion action.
	ConversionUploadErrorEnum_EXPIRED_GCLID ConversionUploadErrorEnum_ConversionUploadError = 5
	// The click associated with the given gclid occurred too recently. Please
	// try uploading again after 24 hours have passed since the click occurred.
	ConversionUploadErrorEnum_TOO_RECENT_GCLID ConversionUploadErrorEnum_ConversionUploadError = 6
	// The click associated with the given gclid could not be found in the
	// system. This can happen if Google Click IDs are collected for non Google
	// Ads clicks.
	ConversionUploadErrorEnum_GCLID_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 7
	// The click associated with the given gclid is owned by a customer
	// account that the uploading customer does not manage.
	ConversionUploadErrorEnum_UNAUTHORIZED_CUSTOMER ConversionUploadErrorEnum_ConversionUploadError = 8
	// No upload eligible conversion action that matches the provided
	// information can be found for the customer.
	ConversionUploadErrorEnum_INVALID_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 9
	// The specified conversion action was created too recently.
	// Please try the upload again after 4-6 hours have passed since the
	// conversion action was created.
	ConversionUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 10
	// The click associated with the given gclid does not contain conversion
	// tracking information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME ConversionUploadErrorEnum_ConversionUploadError = 11
	// The specified conversion action does not use an external attribution
	// model, but external_attribution_data was set.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 12
	// The specified conversion action uses an external attribution model, but
	// external_attribution_data or one of its contained fields was not set.
	// Both external_attribution_credit and external_attribution_model must be
	// set for externally attributed conversion actions.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 13
	// Order IDs are not supported for conversion actions which use an external
	// attribution model.
	ConversionUploadErrorEnum_ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 14
	// A conversion with the same order id and conversion action combination
	// already exists in our system.
	ConversionUploadErrorEnum_ORDER_ID_ALREADY_IN_USE ConversionUploadErrorEnum_ConversionUploadError = 15
	// The request contained two or more conversions with the same order id and
	// conversion action combination.
	ConversionUploadErrorEnum_DUPLICATE_ORDER_ID ConversionUploadErrorEnum_ConversionUploadError = 16
)

var ConversionUploadErrorEnum_ConversionUploadError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "TOO_MANY_CONVERSIONS_IN_REQUEST",
	3:  "UNPARSEABLE_GCLID",
	4:  "CONVERSION_PRECEDES_GCLID",
	5:  "EXPIRED_GCLID",
	6:  "TOO_RECENT_GCLID",
	7:  "GCLID_NOT_FOUND",
	8:  "UNAUTHORIZED_CUSTOMER",
	9:  "INVALID_CONVERSION_ACTION",
	10: "TOO_RECENT_CONVERSION_ACTION",
	11: "CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME",
	12: "EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	13: "EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	14: "ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
	15: "ORDER_ID_ALREADY_IN_USE",
	16: "DUPLICATE_ORDER_ID",
}
var ConversionUploadErrorEnum_ConversionUploadError_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"TOO_MANY_CONVERSIONS_IN_REQUEST": 2,
	"UNPARSEABLE_GCLID":               3,
	"CONVERSION_PRECEDES_GCLID":       4,
	"EXPIRED_GCLID":                   5,
	"TOO_RECENT_GCLID":                6,
	"GCLID_NOT_FOUND":                 7,
	"UNAUTHORIZED_CUSTOMER":           8,
	"INVALID_CONVERSION_ACTION":       9,
	"TOO_RECENT_CONVERSION_ACTION":    10,
	"CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME":                            11,
	"EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 12,
	"EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 13,
	"ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION":            14,
	"ORDER_ID_ALREADY_IN_USE": 15,
	"DUPLICATE_ORDER_ID":      16,
}

func (x ConversionUploadErrorEnum_ConversionUploadError) String() string {
	return proto.EnumName(ConversionUploadErrorEnum_ConversionUploadError_name, int32(x))
}
func (ConversionUploadErrorEnum_ConversionUploadError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_error_f0cddf865a2edd2e, []int{0, 0}
}

// Container for enum describing possible conversion upload errors.
type ConversionUploadErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionUploadErrorEnum) Reset()         { *m = ConversionUploadErrorEnum{} }
func (m *ConversionUploadErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionUploadErrorEnum) ProtoMessage()    {}
func (*ConversionUploadErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_upload_error_f0cddf865a2edd2e, []int{0}
}
func (m *ConversionUploadErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionUploadErrorEnum.Unmarshal(m, b)
}
func (m *ConversionUploadErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionUploadErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ConversionUploadErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionUploadErrorEnum.Merge(dst, src)
}
func (m *ConversionUploadErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionUploadErrorEnum.Size(m)
}
func (m *ConversionUploadErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionUploadErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionUploadErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionUploadErrorEnum)(nil), "google.ads.googleads.v1.errors.ConversionUploadErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.ConversionUploadErrorEnum_ConversionUploadError", ConversionUploadErrorEnum_ConversionUploadError_name, ConversionUploadErrorEnum_ConversionUploadError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/conversion_upload_error.proto", fileDescriptor_conversion_upload_error_f0cddf865a2edd2e)
}

var fileDescriptor_conversion_upload_error_f0cddf865a2edd2e = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xd1, 0x4e, 0x13, 0x41,
	0x14, 0x86, 0xa5, 0x08, 0xe8, 0x20, 0x32, 0x8c, 0xa2, 0x01, 0x11, 0x4d, 0xbd, 0xdf, 0xa6, 0x9a,
	0x78, 0xb1, 0x7a, 0x33, 0xdd, 0x39, 0xd4, 0x09, 0xed, 0xcc, 0x32, 0x3b, 0x53, 0x81, 0x34, 0x99,
	0x54, 0xda, 0x34, 0x4d, 0x60, 0xa7, 0xe9, 0x02, 0x6f, 0xe0, 0x8b, 0xe8, 0x9d, 0x8f, 0xe2, 0xa3,
	0x78, 0xef, 0xbd, 0x99, 0xdd, 0xee, 0x4a, 0x62, 0x25, 0x7a, 0xd5, 0x93, 0x33, 0xdf, 0xff, 0xff,
	0xa7, 0x9b, 0x73, 0xd0, 0xfb, 0xb1, 0x73, 0xe3, 0xf3, 0x51, 0x63, 0x30, 0xcc, 0x1a, 0x45, 0xe9,
	0xab, 0xeb, 0x66, 0x63, 0x34, 0x9b, 0xb9, 0x59, 0xd6, 0x38, 0x73, 0xe9, 0xf5, 0x68, 0x96, 0x4d,
	0x5c, 0x6a, 0xaf, 0xa6, 0xe7, 0x6e, 0x30, 0xb4, 0xf9, 0x43, 0x30, 0x9d, 0xb9, 0x4b, 0x47, 0xf6,
	0x0b, 0x49, 0x30, 0x18, 0x66, 0x41, 0xa5, 0x0e, 0xae, 0x9b, 0x41, 0xa1, 0xde, 0xdd, 0x2b, 0xdd,
	0xa7, 0x93, 0xc6, 0x20, 0x4d, 0xdd, 0xe5, 0xe0, 0x72, 0xe2, 0xd2, 0xac, 0x50, 0xd7, 0xbf, 0xae,
	0xa0, 0x9d, 0xa8, 0xf2, 0x37, 0xb9, 0x3d, 0x78, 0x21, 0xa4, 0x57, 0x17, 0xf5, 0xcf, 0x2b, 0x68,
	0x7b, 0xe1, 0x2b, 0xd9, 0x44, 0xeb, 0x46, 0x24, 0x31, 0x44, 0xfc, 0x80, 0x03, 0xc3, 0x77, 0xc8,
	0x3a, 0x5a, 0x33, 0xe2, 0x50, 0xc8, 0x8f, 0x02, 0x2f, 0x91, 0x57, 0xe8, 0x85, 0x96, 0xd2, 0x76,
	0xa9, 0x38, 0xb1, 0x91, 0x14, 0x3d, 0x50, 0x09, 0x97, 0x22, 0xb1, 0x5c, 0x58, 0x05, 0x47, 0x06,
	0x12, 0x8d, 0x6b, 0x64, 0x1b, 0x6d, 0x19, 0x11, 0x53, 0x95, 0x00, 0x6d, 0x75, 0xc0, 0xb6, 0xa3,
	0x0e, 0x67, 0x78, 0x99, 0x3c, 0x47, 0x3b, 0xbf, 0x25, 0x36, 0x56, 0x10, 0x01, 0x83, 0x64, 0xfe,
	0x7c, 0x97, 0x6c, 0xa1, 0x0d, 0x38, 0x8e, 0xb9, 0x02, 0x36, 0x6f, 0xad, 0x90, 0xc7, 0x08, 0xfb,
	0x34, 0x4f, 0x0a, 0x3d, 0xef, 0xae, 0x92, 0x47, 0x68, 0x33, 0x2f, 0xad, 0x90, 0xda, 0x1e, 0x48,
	0x23, 0x18, 0x5e, 0x23, 0x3b, 0x68, 0xdb, 0x08, 0x6a, 0xf4, 0x07, 0xa9, 0xf8, 0x29, 0x30, 0x1b,
	0x99, 0x44, 0xcb, 0x2e, 0x28, 0x7c, 0xcf, 0xe7, 0x72, 0xd1, 0xa3, 0x5e, 0x71, 0x23, 0x9f, 0x46,
	0x9a, 0x4b, 0x81, 0xef, 0x93, 0x97, 0x68, 0xef, 0x46, 0xc8, 0x9f, 0x04, 0x22, 0x6f, 0xd1, 0xeb,
	0x1b, 0x6d, 0xad, 0x68, 0x74, 0xc8, 0x45, 0x3b, 0x8f, 0x07, 0xe1, 0xff, 0x22, 0xb3, 0x54, 0x5b,
	0xde, 0x8d, 0x15, 0x24, 0x05, 0xc2, 0xbb, 0x80, 0xd7, 0xc9, 0x11, 0xea, 0xc2, 0xb1, 0x06, 0x25,
	0x68, 0xc7, 0x52, 0xad, 0x15, 0x6f, 0x19, 0xef, 0x68, 0x19, 0xd5, 0xd4, 0x26, 0xe0, 0x87, 0x57,
	0x56, 0x48, 0x61, 0x4b, 0xaa, 0x73, 0x52, 0x71, 0xb0, 0x68, 0xd8, 0x07, 0xb7, 0x5b, 0xfa, 0x81,
	0x4a, 0xdb, 0x7f, 0xb5, 0xdc, 0x20, 0x07, 0xa8, 0x25, 0x15, 0x03, 0x65, 0xe7, 0x5f, 0x34, 0x06,
	0xd5, 0xe5, 0xda, 0xd3, 0xff, 0xe3, 0xf3, 0x90, 0x3c, 0x43, 0x4f, 0x2b, 0x1f, 0xda, 0x51, 0x40,
	0xd9, 0x89, 0x5f, 0x0b, 0x93, 0x00, 0xde, 0x24, 0x4f, 0x10, 0x61, 0x26, 0xee, 0xf0, 0x88, 0x6a,
	0xb0, 0x25, 0x86, 0x71, 0xeb, 0xe7, 0x12, 0xaa, 0x9f, 0xb9, 0x8b, 0xe0, 0xf6, 0x55, 0x6f, 0xed,
	0x2e, 0xdc, 0xd5, 0xd8, 0x2f, 0x7a, 0xbc, 0x74, 0xca, 0xe6, 0xea, 0xb1, 0x3b, 0x1f, 0xa4, 0xe3,
	0xc0, 0xcd, 0xc6, 0x8d, 0xf1, 0x28, 0xcd, 0xcf, 0xa0, 0x3c, 0xbb, 0xe9, 0x24, 0xfb, 0xdb, 0x15,
	0xbe, 0x2b, 0x7e, 0xbe, 0xd4, 0x96, 0xdb, 0x94, 0x7e, 0xab, 0xed, 0xb7, 0x0b, 0x33, 0x3a, 0xcc,
	0x82, 0xa2, 0xf4, 0x55, 0xaf, 0x19, 0xe4, 0x91, 0xd9, 0xf7, 0x12, 0xe8, 0xd3, 0x61, 0xd6, 0xaf,
	0x80, 0x7e, 0xaf, 0xd9, 0x2f, 0x80, 0x1f, 0xb5, 0x7a, 0xd1, 0x0d, 0x43, 0x3a, 0xcc, 0xc2, 0xb0,
	0x42, 0xc2, 0xb0, 0xd7, 0x0c, 0xc3, 0x02, 0xfa, 0xb4, 0x9a, 0x4f, 0xf7, 0xe6, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0xde, 0x1b, 0xca, 0x22, 0x04, 0x00, 0x00,
}