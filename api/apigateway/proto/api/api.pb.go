// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api/api.proto

package go_micro_srv_album

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Album struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentAlbum          int64    `protobuf:"varint,2,opt,name=parentAlbum,proto3" json:"parentAlbum,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Album) Reset()         { *m = Album{} }
func (m *Album) String() string { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()    {}
func (*Album) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{0}
}

func (m *Album) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Album.Unmarshal(m, b)
}
func (m *Album) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Album.Marshal(b, m, deterministic)
}
func (m *Album) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Album.Merge(m, src)
}
func (m *Album) XXX_Size() int {
	return xxx_messageInfo_Album.Size(m)
}
func (m *Album) XXX_DiscardUnknown() {
	xxx_messageInfo_Album.DiscardUnknown(m)
}

var xxx_messageInfo_Album proto.InternalMessageInfo

func (m *Album) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Album) GetParentAlbum() int64 {
	if m != nil {
		return m.ParentAlbum
	}
	return 0
}

func (m *Album) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Album) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Album) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Album) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type UploadImageRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_Chunk
	Data                 isUploadImageRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadImageRequest) Reset()         { *m = UploadImageRequest{} }
func (m *UploadImageRequest) String() string { return proto.CompactTextString(m) }
func (*UploadImageRequest) ProtoMessage()    {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{1}
}

func (m *UploadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRequest.Unmarshal(m, b)
}
func (m *UploadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRequest.Marshal(b, m, deterministic)
}
func (m *UploadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRequest.Merge(m, src)
}
func (m *UploadImageRequest) XXX_Size() int {
	return xxx_messageInfo_UploadImageRequest.Size(m)
}
func (m *UploadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRequest proto.InternalMessageInfo

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_Chunk) isUploadImageRequest_Data() {}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadImageRequest) GetChunk() []byte {
	if x, ok := m.GetData().(*UploadImageRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadImageRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_Chunk)(nil),
	}
}

type ImageInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AlbumId              int64    `protobuf:"varint,2,opt,name=albumId,proto3" json:"albumId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{2}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageInfo) GetAlbumId() int64 {
	if m != nil {
		return m.AlbumId
	}
	return 0
}

type Albums struct {
	List                 []*Album `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Albums) Reset()         { *m = Albums{} }
func (m *Albums) String() string { return proto.CompactTextString(m) }
func (*Albums) ProtoMessage()    {}
func (*Albums) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{3}
}

func (m *Albums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Albums.Unmarshal(m, b)
}
func (m *Albums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Albums.Marshal(b, m, deterministic)
}
func (m *Albums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Albums.Merge(m, src)
}
func (m *Albums) XXX_Size() int {
	return xxx_messageInfo_Albums.Size(m)
}
func (m *Albums) XXX_DiscardUnknown() {
	xxx_messageInfo_Albums.DiscardUnknown(m)
}

var xxx_messageInfo_Albums proto.InternalMessageInfo

func (m *Albums) GetList() []*Album {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Albums) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type AlbumResponse struct {
	Album                *Album   `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlbumResponse) Reset()         { *m = AlbumResponse{} }
func (m *AlbumResponse) String() string { return proto.CompactTextString(m) }
func (*AlbumResponse) ProtoMessage()    {}
func (*AlbumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{4}
}

func (m *AlbumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlbumResponse.Unmarshal(m, b)
}
func (m *AlbumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlbumResponse.Marshal(b, m, deterministic)
}
func (m *AlbumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlbumResponse.Merge(m, src)
}
func (m *AlbumResponse) XXX_Size() int {
	return xxx_messageInfo_AlbumResponse.Size(m)
}
func (m *AlbumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlbumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlbumResponse proto.InternalMessageInfo

func (m *AlbumResponse) GetAlbum() *Album {
	if m != nil {
		return m.Album
	}
	return nil
}

type ImageResponse struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageResponse) Reset()         { *m = ImageResponse{} }
func (m *ImageResponse) String() string { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()    {}
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{5}
}

func (m *ImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageResponse.Unmarshal(m, b)
}
func (m *ImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageResponse.Marshal(b, m, deterministic)
}
func (m *ImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageResponse.Merge(m, src)
}
func (m *ImageResponse) XXX_Size() int {
	return xxx_messageInfo_ImageResponse.Size(m)
}
func (m *ImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageResponse proto.InternalMessageInfo

func (m *ImageResponse) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type Images struct {
	List                 []*Image `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Images) Reset()         { *m = Images{} }
func (m *Images) String() string { return proto.CompactTextString(m) }
func (*Images) ProtoMessage()    {}
func (*Images) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{6}
}

func (m *Images) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Images.Unmarshal(m, b)
}
func (m *Images) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Images.Marshal(b, m, deterministic)
}
func (m *Images) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Images.Merge(m, src)
}
func (m *Images) XXX_Size() int {
	return xxx_messageInfo_Images.Size(m)
}
func (m *Images) XXX_DiscardUnknown() {
	xxx_messageInfo_Images.DiscardUnknown(m)
}

var xxx_messageInfo_Images proto.InternalMessageInfo

func (m *Images) GetList() []*Image {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Images) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Image struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AlbumId              int64    `protobuf:"varint,2,opt,name=albumId,proto3" json:"albumId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt            int64    `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{7}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Image) GetAlbumId() int64 {
	if m != nil {
		return m.AlbumId
	}
	return 0
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Image) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{9}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	AlbumId              int64    `protobuf:"varint,4,opt,name=albumId,proto3" json:"albumId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{10}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetAlbumId() int64 {
	if m != nil {
		return m.AlbumId
	}
	return 0
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca6d5bbc959d58c0, []int{11}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Album)(nil), "go.micro.srv.album.Album")
	proto.RegisterType((*UploadImageRequest)(nil), "go.micro.srv.album.UploadImageRequest")
	proto.RegisterType((*ImageInfo)(nil), "go.micro.srv.album.ImageInfo")
	proto.RegisterType((*Albums)(nil), "go.micro.srv.album.Albums")
	proto.RegisterType((*AlbumResponse)(nil), "go.micro.srv.album.AlbumResponse")
	proto.RegisterType((*ImageResponse)(nil), "go.micro.srv.album.ImageResponse")
	proto.RegisterType((*Images)(nil), "go.micro.srv.album.Images")
	proto.RegisterType((*Image)(nil), "go.micro.srv.album.Image")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.album.DeleteRequest")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.album.GetRequest")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.album.SearchRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.album.Response")
}

func init() { proto.RegisterFile("proto/api/api.proto", fileDescriptor_ca6d5bbc959d58c0) }

var fileDescriptor_ca6d5bbc959d58c0 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6b, 0xdb, 0x4e,
	0x10, 0xb5, 0xac, 0x8f, 0x5f, 0x3c, 0x8e, 0x7f, 0x84, 0x6d, 0x08, 0xaa, 0x71, 0x5b, 0x47, 0x87,
	0xe2, 0x4b, 0x15, 0x70, 0x4e, 0xbd, 0x35, 0x6d, 0x21, 0x36, 0xa4, 0xb4, 0x28, 0x94, 0x9e, 0x37,
	0xd2, 0xda, 0x5e, 0xaa, 0xaf, 0x48, 0xab, 0x40, 0xff, 0x8e, 0x9c, 0xfb, 0xbf, 0x96, 0x9d, 0x95,
	0xe5, 0x2f, 0x49, 0x09, 0xa5, 0x07, 0xc3, 0xcc, 0xce, 0x9b, 0xa7, 0x99, 0xa7, 0xb7, 0x32, 0xbc,
	0x48, 0xb3, 0x44, 0x24, 0x17, 0x34, 0xe5, 0xf2, 0xe7, 0x62, 0x46, 0xc8, 0x32, 0x71, 0x23, 0xee,
	0x67, 0x89, 0x9b, 0x67, 0x0f, 0x2e, 0x0d, 0xef, 0x8a, 0xc8, 0xf9, 0xad, 0x81, 0x79, 0x25, 0x23,
	0xf2, 0x3f, 0x74, 0x79, 0x60, 0x6b, 0x63, 0x6d, 0xa2, 0x7b, 0x5d, 0x1e, 0x90, 0x31, 0xf4, 0x53,
	0x9a, 0xb1, 0x58, 0x60, 0xd9, 0xee, 0x62, 0x61, 0xfb, 0x88, 0x10, 0x30, 0x62, 0x1a, 0x31, 0x5b,
	0x1f, 0x6b, 0x93, 0x9e, 0x87, 0x31, 0x19, 0x41, 0xcf, 0xcf, 0x18, 0x15, 0x2c, 0xb8, 0x12, 0xb6,
	0x81, 0x3d, 0x9b, 0x03, 0x59, 0x2d, 0xd2, 0xa0, 0xac, 0x9a, 0xaa, 0x5a, 0x1d, 0x48, 0xbe, 0x94,
	0x8a, 0x95, 0x6d, 0x29, 0x3e, 0x19, 0x3b, 0x1c, 0xc8, 0xf7, 0x34, 0x4c, 0x68, 0x30, 0x8f, 0xe8,
	0x92, 0x79, 0xec, 0xbe, 0x60, 0xb9, 0x20, 0x97, 0x60, 0xf0, 0x78, 0x91, 0xe0, 0xb4, 0xfd, 0xe9,
	0x2b, 0xf7, 0x70, 0x31, 0x17, 0xf1, 0xf3, 0x78, 0x91, 0xcc, 0x3a, 0x1e, 0x82, 0xc9, 0x19, 0x98,
	0xfe, 0xaa, 0x88, 0x7f, 0xe2, 0x2a, 0xc7, 0xb3, 0x8e, 0xa7, 0xd2, 0x8f, 0x16, 0x18, 0x01, 0x15,
	0xd4, 0x79, 0x0f, 0xbd, 0xaa, 0xa9, 0xda, 0x4d, 0xdb, 0xda, 0xcd, 0x86, 0xff, 0x90, 0x7b, 0x1e,
	0x94, 0x6a, 0xac, 0x53, 0xe7, 0x0b, 0x58, 0x28, 0x49, 0x4e, 0xde, 0x81, 0x11, 0xf2, 0x5c, 0xd8,
	0xda, 0x58, 0x9f, 0xf4, 0xa7, 0x2f, 0xeb, 0x26, 0x43, 0xa4, 0x87, 0x30, 0x72, 0x0a, 0xa6, 0x9f,
	0x14, 0xb1, 0x28, 0x09, 0x55, 0xe2, 0x7c, 0x80, 0x81, 0x02, 0xb1, 0x3c, 0x4d, 0xe2, 0x9c, 0x91,
	0x0b, 0x30, 0xb1, 0xb7, 0x5c, 0xb8, 0x85, 0x56, 0xe1, 0x24, 0x43, 0x29, 0xd8, 0x86, 0x81, 0xcb,
	0x83, 0x36, 0x06, 0xd5, 0xa1, 0x70, 0x72, 0x25, 0xcc, 0x9f, 0xb5, 0x92, 0xea, 0x6c, 0x5b, 0xe9,
	0x51, 0x03, 0x13, 0x51, 0x07, 0x3e, 0x6b, 0x54, 0xb5, 0xd6, 0x5f, 0x27, 0xa0, 0x17, 0x59, 0x88,
	0xce, 0xea, 0x79, 0x32, 0xdc, 0x75, 0x9c, 0xd9, 0xea, 0x38, 0x6b, 0xcf, 0x71, 0xce, 0x1b, 0x18,
	0x7c, 0x66, 0x21, 0x13, 0x95, 0xb1, 0xf6, 0x86, 0x73, 0x46, 0x00, 0xd7, 0x4c, 0x34, 0x55, 0x23,
	0x18, 0xdc, 0x32, 0x9a, 0xf9, 0xab, 0x35, 0xe0, 0x14, 0xcc, 0xfb, 0x82, 0x65, 0xbf, 0x4a, 0xdb,
	0xa8, 0x84, 0x9c, 0x81, 0x95, 0x2c, 0x16, 0x39, 0x5b, 0x4b, 0x52, 0x66, 0x12, 0x1d, 0xf2, 0x88,
	0x0b, 0x5c, 0x50, 0xf7, 0x54, 0xb2, 0xad, 0x87, 0xb1, 0xeb, 0xb2, 0x11, 0x1c, 0x55, 0xef, 0xf3,
	0x04, 0xf4, 0x28, 0x5f, 0x96, 0xcf, 0x91, 0xe1, 0xf4, 0xd1, 0x84, 0x63, 0xf4, 0xc0, 0x2d, 0xcb,
	0x1e, 0xb8, 0xcf, 0xc8, 0x0c, 0xfa, 0x9f, 0x50, 0x07, 0x75, 0x5b, 0x9b, 0x4d, 0x33, 0x1c, 0xd5,
	0x95, 0xd6, 0x8f, 0x72, 0x3a, 0xe4, 0xc7, 0x9a, 0x49, 0xbd, 0xc1, 0xb7, 0x75, 0xf0, 0xc3, 0x5b,
	0xfa, 0x14, 0xed, 0x44, 0x23, 0xdf, 0xa0, 0xaf, 0xf4, 0x57, 0x23, 0x9e, 0xd7, 0x35, 0xec, 0xbc,
	0xa0, 0x27, 0x47, 0xad, 0x18, 0xd5, 0xa8, 0xff, 0x80, 0xf1, 0x2b, 0x1c, 0x5d, 0xb3, 0xf2, 0x8b,
	0xf7, 0xba, 0x0e, 0xbb, 0x31, 0xc8, 0xf0, 0xbc, 0xf9, 0x62, 0xee, 0x13, 0xaa, 0xf9, 0xfe, 0x8a,
	0x70, 0xe7, 0x66, 0x3b, 0x1d, 0x72, 0x03, 0xbd, 0x1b, 0x9e, 0x8b, 0x16, 0x0d, 0x77, 0x5c, 0x3a,
	0x1c, 0x36, 0x4e, 0x99, 0x6f, 0xd8, 0x5a, 0xf4, 0x7b, 0x06, 0x9b, 0xfa, 0x74, 0x38, 0x9d, 0x3b,
	0x0b, 0xff, 0x7a, 0x2e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xaf, 0xa0, 0xaa, 0x91, 0x06,
	0x00, 0x00,
}
