// Code generated by protoc-gen-go. DO NOT EDIT.
// source: asset.proto

/*
Package asset is a generated protocol buffer package.

It is generated from these files:
	asset.proto

It has these top-level messages:
	GetFileUploadURLRequest
	GetFileUploadURLResponse
	RegisterFileRequest
	RegisterFileResponse
	QueryUploadedDataRequest
	QueryUploadedDataResponse
	QueryUploadedData
	QueryUploadedRow
	RegisterRequest
	RegisterResponse
	QueryRequest
	QueryResponse
	QueryData
	QueryRow
	QueryAllAssetRequest
	QueryPara
	QueryAllAssetResponse
	ModifyRequest
	ModifyResponse
	GetFileUploadStatRequest
	GetFileUploadStatResponse
	GetDownLoadURLRequest
	GetDownLoadURLResponse
	QueryByIDRequest
	GetUserPurchaseAssetListRequest
	GetUserPurchaseAssetListResponse
	QueryPurchaseData
	QueryPurchaseRow
*/
package asset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetFileUploadURLRequest struct {
	//    string username = 1;
	//    int64 fileSize = 2;
	//    string fileName = 3;
	//    int32 filePolicy = 4;
	//    int32 fileNumber = 5;
	//    string fileash = 6;
	//    string signature = 7;
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *GetFileUploadURLRequest) Reset()                    { *m = GetFileUploadURLRequest{} }
func (m *GetFileUploadURLRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFileUploadURLRequest) ProtoMessage()               {}
func (*GetFileUploadURLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetFileUploadURLRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type GetFileUploadURLResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *GetFileUploadURLResponse) Reset()                    { *m = GetFileUploadURLResponse{} }
func (m *GetFileUploadURLResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFileUploadURLResponse) ProtoMessage()               {}
func (*GetFileUploadURLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetFileUploadURLResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetFileUploadURLResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetFileUploadURLResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type RegisterFileRequest struct {
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *RegisterFileRequest) Reset()                    { *m = RegisterFileRequest{} }
func (m *RegisterFileRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterFileRequest) ProtoMessage()               {}
func (*RegisterFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterFileRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type RegisterFileResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *RegisterFileResponse) Reset()                    { *m = RegisterFileResponse{} }
func (m *RegisterFileResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterFileResponse) ProtoMessage()               {}
func (*RegisterFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterFileResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterFileResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegisterFileResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type QueryUploadedDataRequest struct {
	PageSize int32  `protobuf:"varint,1,opt,name=pageSize" json:"pageSize"`
	PageNum  int32  `protobuf:"varint,2,opt,name=pageNum" json:"pageNum"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username"`
}

func (m *QueryUploadedDataRequest) Reset()                    { *m = QueryUploadedDataRequest{} }
func (m *QueryUploadedDataRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryUploadedDataRequest) ProtoMessage()               {}
func (*QueryUploadedDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryUploadedDataRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryUploadedDataRequest) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryUploadedDataRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type QueryUploadedDataResponse struct {
	Code int32              `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *QueryUploadedData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string             `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryUploadedDataResponse) Reset()                    { *m = QueryUploadedDataResponse{} }
func (m *QueryUploadedDataResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryUploadedDataResponse) ProtoMessage()               {}
func (*QueryUploadedDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryUploadedDataResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryUploadedDataResponse) GetData() *QueryUploadedData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryUploadedDataResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryUploadedData struct {
	PageNum  uint64              `protobuf:"varint,1,opt,name=pageNum" json:"pageNum"`
	RowCount uint64              `protobuf:"varint,2,opt,name=rowCount" json:"rowCount"`
	Row      []*QueryUploadedRow `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *QueryUploadedData) Reset()                    { *m = QueryUploadedData{} }
func (m *QueryUploadedData) String() string            { return proto.CompactTextString(m) }
func (*QueryUploadedData) ProtoMessage()               {}
func (*QueryUploadedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryUploadedData) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryUploadedData) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *QueryUploadedData) GetRow() []*QueryUploadedRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type QueryUploadedRow struct {
	Username   string `protobuf:"bytes,1,opt,name=username" json:"username"`
	FileHash   string `protobuf:"bytes,2,opt,name=file_hash,json=fileHash" json:"file_hash"`
	FileName   string `protobuf:"bytes,3,opt,name=file_name,json=fileName" json:"file_name"`
	FileSize   uint64 `protobuf:"varint,4,opt,name=file_size,json=fileSize" json:"file_size"`
	FilePolicy string `protobuf:"bytes,5,opt,name=file_policy,json=filePolicy" json:"file_policy"`
	FileNumber uint64 `protobuf:"varint,6,opt,name=file_number,json=fileNumber" json:"file_number"`
	AuthPath   string `protobuf:"bytes,7,opt,name=auth_path,json=authPath" json:"auth_path"`
	CreateTime string `protobuf:"bytes,8,opt,name=create_time,json=createTime" json:"create_time"`
}

func (m *QueryUploadedRow) Reset()                    { *m = QueryUploadedRow{} }
func (m *QueryUploadedRow) String() string            { return proto.CompactTextString(m) }
func (*QueryUploadedRow) ProtoMessage()               {}
func (*QueryUploadedRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryUploadedRow) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryUploadedRow) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

func (m *QueryUploadedRow) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *QueryUploadedRow) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *QueryUploadedRow) GetFilePolicy() string {
	if m != nil {
		return m.FilePolicy
	}
	return ""
}

func (m *QueryUploadedRow) GetFileNumber() uint64 {
	if m != nil {
		return m.FileNumber
	}
	return 0
}

func (m *QueryUploadedRow) GetAuthPath() string {
	if m != nil {
		return m.AuthPath
	}
	return ""
}

func (m *QueryUploadedRow) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type RegisterRequest struct {
	// string username = 1;
	// string asset_id = 2;
	// string session_id = 3;
	// string asset_name = 4;
	// int32 feature_tag = 5;
	// string asset_sample_path = 6;
	// string asset_sample_hash = 7;
	// string asset_storage_path = 8;
	// string asset_storage_hash = 9;
	// string expire_time = 10;
	// int64 price = 11;
	// string description = 12;
	// string upload_date = 13;
	// string signature = 14;
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegisterRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type RegisterResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RegisterResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegisterResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type QueryRequest struct {
	PageSize   int32  `protobuf:"varint,1,opt,name=pageSize" json:"pageSize"`
	PageNum    int32  `protobuf:"varint,2,opt,name=pageNum" json:"pageNum"`
	Username   string `protobuf:"bytes,3,opt,name=username" json:"username"`
	FeatureTag uint64 `protobuf:"varint,4,opt,name=featureTag" json:"featureTag"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *QueryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryRequest) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryRequest) GetFeatureTag() uint64 {
	if m != nil {
		return m.FeatureTag
	}
	return 0
}

type QueryResponse struct {
	Code int32      `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *QueryData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string     `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *QueryResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryResponse) GetData() *QueryData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryData struct {
	PageNum  uint64      `protobuf:"varint,1,opt,name=pageNum" json:"pageNum"`
	RowCount uint64      `protobuf:"varint,2,opt,name=rowCount" json:"rowCount"`
	Row      []*QueryRow `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *QueryData) Reset()                    { *m = QueryData{} }
func (m *QueryData) String() string            { return proto.CompactTextString(m) }
func (*QueryData) ProtoMessage()               {}
func (*QueryData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *QueryData) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryData) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *QueryData) GetRow() []*QueryRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type QueryRow struct {
	AssetId     string `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id"`
	Username    string `protobuf:"bytes,2,opt,name=username" json:"username"`
	AssetName   string `protobuf:"bytes,3,opt,name=asset_name,json=assetName" json:"asset_name"`
	AssetType   string `protobuf:"bytes,4,opt,name=asset_type,json=assetType" json:"asset_type"`
	FeatureTag1 string `protobuf:"bytes,5,opt,name=feature_tag1,json=featureTag1" json:"feature_tag1"`
	FeatureTag2 string `protobuf:"bytes,6,opt,name=feature_tag2,json=featureTag2" json:"feature_tag2"`
	FeatureTag3 string `protobuf:"bytes,7,opt,name=feature_tag3,json=featureTag3" json:"feature_tag3"`
	SamplePath  string `protobuf:"bytes,8,opt,name=sample_path,json=samplePath" json:"sample_path"`
	SampleHash  string `protobuf:"bytes,9,opt,name=sample_hash,json=sampleHash" json:"sample_hash"`
	StoragePath string `protobuf:"bytes,10,opt,name=storage_path,json=storagePath" json:"storage_path"`
	StorageHash string `protobuf:"bytes,11,opt,name=storage_hash,json=storageHash" json:"storage_hash"`
	ExpireTime  uint32 `protobuf:"varint,12,opt,name=expire_time,json=expireTime" json:"expire_time"`
	Price       uint64 `protobuf:"varint,13,opt,name=price" json:"price"`
	Description string `protobuf:"bytes,14,opt,name=description" json:"description"`
	UploadDate  uint32 `protobuf:"varint,15,opt,name=upload_date,json=uploadDate" json:"upload_date"`
	CreateTime  string `protobuf:"bytes,16,opt,name=create_time,json=createTime" json:"create_time"`
}

func (m *QueryRow) Reset()                    { *m = QueryRow{} }
func (m *QueryRow) String() string            { return proto.CompactTextString(m) }
func (*QueryRow) ProtoMessage()               {}
func (*QueryRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *QueryRow) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *QueryRow) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryRow) GetAssetName() string {
	if m != nil {
		return m.AssetName
	}
	return ""
}

func (m *QueryRow) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *QueryRow) GetFeatureTag1() string {
	if m != nil {
		return m.FeatureTag1
	}
	return ""
}

func (m *QueryRow) GetFeatureTag2() string {
	if m != nil {
		return m.FeatureTag2
	}
	return ""
}

func (m *QueryRow) GetFeatureTag3() string {
	if m != nil {
		return m.FeatureTag3
	}
	return ""
}

func (m *QueryRow) GetSamplePath() string {
	if m != nil {
		return m.SamplePath
	}
	return ""
}

func (m *QueryRow) GetSampleHash() string {
	if m != nil {
		return m.SampleHash
	}
	return ""
}

func (m *QueryRow) GetStoragePath() string {
	if m != nil {
		return m.StoragePath
	}
	return ""
}

func (m *QueryRow) GetStorageHash() string {
	if m != nil {
		return m.StorageHash
	}
	return ""
}

func (m *QueryRow) GetExpireTime() uint32 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *QueryRow) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *QueryRow) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *QueryRow) GetUploadDate() uint32 {
	if m != nil {
		return m.UploadDate
	}
	return 0
}

func (m *QueryRow) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type QueryAllAssetRequest struct {
	Username  string     `protobuf:"bytes,1,opt,name=username" json:"username"`
	SessionId string     `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id"`
	QueryPara *QueryPara `protobuf:"bytes,3,opt,name=query_para,json=queryPara" json:"query_para"`
}

func (m *QueryAllAssetRequest) Reset()                    { *m = QueryAllAssetRequest{} }
func (m *QueryAllAssetRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryAllAssetRequest) ProtoMessage()               {}
func (*QueryAllAssetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *QueryAllAssetRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryAllAssetRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *QueryAllAssetRequest) GetQueryPara() *QueryPara {
	if m != nil {
		return m.QueryPara
	}
	return nil
}

type QueryPara struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id"`
}

func (m *QueryPara) Reset()                    { *m = QueryPara{} }
func (m *QueryPara) String() string            { return proto.CompactTextString(m) }
func (*QueryPara) ProtoMessage()               {}
func (*QueryPara) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *QueryPara) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryPara) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type QueryAllAssetResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *QueryAllAssetResponse) Reset()                    { *m = QueryAllAssetResponse{} }
func (m *QueryAllAssetResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryAllAssetResponse) ProtoMessage()               {}
func (*QueryAllAssetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *QueryAllAssetResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryAllAssetResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryAllAssetResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ModifyRequest struct {
	//    string username = 1;
	//    string signature = 2;
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *ModifyRequest) Reset()                    { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()               {}
func (*ModifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ModifyRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type ModifyResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *ModifyResponse) Reset()                    { *m = ModifyResponse{} }
func (m *ModifyResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyResponse) ProtoMessage()               {}
func (*ModifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *ModifyResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ModifyResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ModifyResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetFileUploadStatRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username"`
	FileName string `protobuf:"bytes,2,opt,name=fileName" json:"fileName"`
}

func (m *GetFileUploadStatRequest) Reset()                    { *m = GetFileUploadStatRequest{} }
func (m *GetFileUploadStatRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFileUploadStatRequest) ProtoMessage()               {}
func (*GetFileUploadStatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *GetFileUploadStatRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetFileUploadStatRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type GetFileUploadStatResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *GetFileUploadStatResponse) Reset()                    { *m = GetFileUploadStatResponse{} }
func (m *GetFileUploadStatResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFileUploadStatResponse) ProtoMessage()               {}
func (*GetFileUploadStatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *GetFileUploadStatResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetFileUploadStatResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetFileUploadStatResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetDownLoadURLRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username"`
	FileName string `protobuf:"bytes,2,opt,name=fileName" json:"fileName"`
}

func (m *GetDownLoadURLRequest) Reset()                    { *m = GetDownLoadURLRequest{} }
func (m *GetDownLoadURLRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDownLoadURLRequest) ProtoMessage()               {}
func (*GetDownLoadURLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *GetDownLoadURLRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetDownLoadURLRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type GetDownLoadURLResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data string `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *GetDownLoadURLResponse) Reset()                    { *m = GetDownLoadURLResponse{} }
func (m *GetDownLoadURLResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDownLoadURLResponse) ProtoMessage()               {}
func (*GetDownLoadURLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *GetDownLoadURLResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetDownLoadURLResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetDownLoadURLResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type QueryByIDRequest struct {
	PageSize int32  `protobuf:"varint,1,opt,name=pageSize" json:"pageSize"`
	PageNum  int32  `protobuf:"varint,2,opt,name=pageNum" json:"pageNum"`
	AssetID  string `protobuf:"bytes,3,opt,name=assetID" json:"assetID"`
}

func (m *QueryByIDRequest) Reset()                    { *m = QueryByIDRequest{} }
func (m *QueryByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryByIDRequest) ProtoMessage()               {}
func (*QueryByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *QueryByIDRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryByIDRequest) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryByIDRequest) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

type GetUserPurchaseAssetListRequest struct {
	PageSize int32  `protobuf:"varint,1,opt,name=pageSize" json:"pageSize"`
	PageNum  int32  `protobuf:"varint,2,opt,name=pageNum" json:"pageNum"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username"`
	AssetId  string `protobuf:"bytes,4,opt,name=asset_id,json=assetId" json:"asset_id"`
}

func (m *GetUserPurchaseAssetListRequest) Reset()         { *m = GetUserPurchaseAssetListRequest{} }
func (m *GetUserPurchaseAssetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserPurchaseAssetListRequest) ProtoMessage()    {}
func (*GetUserPurchaseAssetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{24}
}

func (m *GetUserPurchaseAssetListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetUserPurchaseAssetListRequest) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetUserPurchaseAssetListRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserPurchaseAssetListRequest) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

type GetUserPurchaseAssetListResponse struct {
	Code int32              `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *QueryPurchaseData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string             `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetUserPurchaseAssetListResponse) Reset()         { *m = GetUserPurchaseAssetListResponse{} }
func (m *GetUserPurchaseAssetListResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserPurchaseAssetListResponse) ProtoMessage()    {}
func (*GetUserPurchaseAssetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{25}
}

func (m *GetUserPurchaseAssetListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetUserPurchaseAssetListResponse) GetData() *QueryPurchaseData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetUserPurchaseAssetListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryPurchaseData struct {
	PageNum  uint64              `protobuf:"varint,1,opt,name=pageNum" json:"pageNum"`
	RowCount uint64              `protobuf:"varint,2,opt,name=rowCount" json:"rowCount"`
	Row      []*QueryPurchaseRow `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *QueryPurchaseData) Reset()                    { *m = QueryPurchaseData{} }
func (m *QueryPurchaseData) String() string            { return proto.CompactTextString(m) }
func (*QueryPurchaseData) ProtoMessage()               {}
func (*QueryPurchaseData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

func (m *QueryPurchaseData) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryPurchaseData) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *QueryPurchaseData) GetRow() []*QueryPurchaseRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type QueryPurchaseRow struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id"`
	AssetId   string `protobuf:"bytes,3,opt,name=asset_id,json=assetId" json:"asset_id"`
	DataReqId string `protobuf:"bytes,4,opt,name=data_req_id,json=dataReqId" json:"data_req_id"`
	Consumer  string `protobuf:"bytes,5,opt,name=consumer" json:"consumer"`
}

func (m *QueryPurchaseRow) Reset()                    { *m = QueryPurchaseRow{} }
func (m *QueryPurchaseRow) String() string            { return proto.CompactTextString(m) }
func (*QueryPurchaseRow) ProtoMessage()               {}
func (*QueryPurchaseRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

func (m *QueryPurchaseRow) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryPurchaseRow) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *QueryPurchaseRow) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *QueryPurchaseRow) GetDataReqId() string {
	if m != nil {
		return m.DataReqId
	}
	return ""
}

func (m *QueryPurchaseRow) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFileUploadURLRequest)(nil), "GetFileUploadURLRequest")
	proto.RegisterType((*GetFileUploadURLResponse)(nil), "GetFileUploadURLResponse")
	proto.RegisterType((*RegisterFileRequest)(nil), "RegisterFileRequest")
	proto.RegisterType((*RegisterFileResponse)(nil), "RegisterFileResponse")
	proto.RegisterType((*QueryUploadedDataRequest)(nil), "QueryUploadedDataRequest")
	proto.RegisterType((*QueryUploadedDataResponse)(nil), "QueryUploadedDataResponse")
	proto.RegisterType((*QueryUploadedData)(nil), "QueryUploadedData")
	proto.RegisterType((*QueryUploadedRow)(nil), "QueryUploadedRow")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*QueryRequest)(nil), "QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "QueryResponse")
	proto.RegisterType((*QueryData)(nil), "QueryData")
	proto.RegisterType((*QueryRow)(nil), "QueryRow")
	proto.RegisterType((*QueryAllAssetRequest)(nil), "QueryAllAssetRequest")
	proto.RegisterType((*QueryPara)(nil), "QueryPara")
	proto.RegisterType((*QueryAllAssetResponse)(nil), "QueryAllAssetResponse")
	proto.RegisterType((*ModifyRequest)(nil), "ModifyRequest")
	proto.RegisterType((*ModifyResponse)(nil), "ModifyResponse")
	proto.RegisterType((*GetFileUploadStatRequest)(nil), "GetFileUploadStatRequest")
	proto.RegisterType((*GetFileUploadStatResponse)(nil), "GetFileUploadStatResponse")
	proto.RegisterType((*GetDownLoadURLRequest)(nil), "GetDownLoadURLRequest")
	proto.RegisterType((*GetDownLoadURLResponse)(nil), "GetDownLoadURLResponse")
	proto.RegisterType((*QueryByIDRequest)(nil), "QueryByIDRequest")
	proto.RegisterType((*GetUserPurchaseAssetListRequest)(nil), "GetUserPurchaseAssetListRequest")
	proto.RegisterType((*GetUserPurchaseAssetListResponse)(nil), "GetUserPurchaseAssetListResponse")
	proto.RegisterType((*QueryPurchaseData)(nil), "QueryPurchaseData")
	proto.RegisterType((*QueryPurchaseRow)(nil), "QueryPurchaseRow")
}

func init() { proto.RegisterFile("asset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x5b, 0x4f, 0x1b, 0x47,
	0x14, 0xb6, 0xb1, 0x0d, 0xf6, 0x31, 0x17, 0x33, 0x05, 0xb2, 0x6c, 0x14, 0x20, 0x53, 0xa9, 0xa2,
	0x8a, 0x3a, 0x2a, 0xa0, 0x3e, 0x56, 0x6a, 0x12, 0x14, 0x4a, 0x45, 0x29, 0xd9, 0xc0, 0x6b, 0xad,
	0xc1, 0x3b, 0xd8, 0x2b, 0xd9, 0xde, 0x65, 0x67, 0x56, 0xd4, 0x55, 0x1f, 0xfa, 0xd2, 0xa7, 0x4a,
	0xfd, 0x15, 0xfd, 0x21, 0xfd, 0x69, 0xd5, 0x5c, 0x76, 0x99, 0x5d, 0xaf, 0x53, 0x37, 0x4e, 0xde,
	0x66, 0xce, 0x9c, 0xdb, 0x9c, 0xf9, 0xce, 0x65, 0xa0, 0x4d, 0x39, 0x67, 0x82, 0x44, 0x71, 0x28,
	0x42, 0xfc, 0x0d, 0x3c, 0x39, 0x63, 0xe2, 0x4d, 0x30, 0x64, 0x37, 0xd1, 0x30, 0xa4, 0xfe, 0x8d,
	0x77, 0xe1, 0xb1, 0xfb, 0x84, 0x71, 0x81, 0x5c, 0x68, 0x46, 0x21, 0x17, 0xaf, 0x42, 0x7f, 0xe2,
	0x54, 0x0f, 0xaa, 0x87, 0x2d, 0x2f, 0xdb, 0xe3, 0x6b, 0x70, 0xa6, 0xc5, 0x78, 0x14, 0x8e, 0x39,
	0x43, 0x08, 0xea, 0xbd, 0xd0, 0x67, 0x4a, 0x66, 0xcd, 0x53, 0x6b, 0xd4, 0x81, 0xda, 0x88, 0xf7,
	0x9d, 0x25, 0xa5, 0x46, 0x2e, 0x25, 0x97, 0x4f, 0x05, 0x75, 0x6a, 0x8a, 0xa4, 0xd6, 0xf8, 0x08,
	0x3e, 0xf3, 0x58, 0x3f, 0xe0, 0x82, 0xc5, 0x52, 0xf5, 0x3c, 0x8e, 0x5c, 0xc1, 0x56, 0x5e, 0x64,
	0x61, 0x27, 0x86, 0xe0, 0xbc, 0x4d, 0x58, 0x3c, 0xd1, 0x17, 0x63, 0xfe, 0x29, 0x15, 0xd4, 0xf6,
	0x84, 0xf6, 0xd9, 0xbb, 0xe0, 0x57, 0xad, 0xb9, 0xe1, 0x65, 0x7b, 0xe4, 0xc0, 0x8a, 0x5c, 0x5f,
	0x26, 0x23, 0x65, 0xa1, 0xe1, 0xa5, 0x5b, 0x29, 0x95, 0x70, 0x16, 0x8f, 0xe9, 0x88, 0x19, 0x4b,
	0xd9, 0x1e, 0x07, 0xb0, 0x5b, 0x62, 0xad, 0xe4, 0x12, 0x0d, 0x73, 0x89, 0x2f, 0x8c, 0xcb, 0xd2,
	0x46, 0xfb, 0x18, 0x91, 0x69, 0x69, 0x75, 0x9e, 0x5e, 0xb6, 0x96, 0x5d, 0x16, 0x8f, 0x61, 0x73,
	0x8a, 0xd9, 0xf6, 0x5a, 0x5a, 0xa9, 0xe7, 0xbc, 0x8e, 0xc3, 0x87, 0xd7, 0x61, 0x32, 0x16, 0xca,
	0x58, 0xdd, 0xcb, 0xf6, 0xe8, 0x73, 0xa8, 0xc5, 0xe1, 0x83, 0x53, 0x3b, 0xa8, 0x1d, 0xb6, 0x8f,
	0x37, 0xf3, 0x3e, 0x78, 0xe1, 0x83, 0x27, 0x4f, 0xf1, 0x1f, 0x4b, 0xd0, 0x29, 0x9e, 0xe4, 0x62,
	0x51, 0xcd, 0xc7, 0x02, 0x3d, 0x85, 0xd6, 0x5d, 0x30, 0x64, 0xdd, 0x01, 0xe5, 0x03, 0xf3, 0x4a,
	0x4d, 0x49, 0xf8, 0x9e, 0xf2, 0x41, 0x76, 0x68, 0x47, 0x51, 0x12, 0x2e, 0x6d, 0x49, 0x2e, 0x1f,
	0xa6, 0xae, 0x9d, 0x95, 0x04, 0xf5, 0x30, 0xfb, 0xd0, 0x56, 0x87, 0x51, 0x38, 0x0c, 0x7a, 0x13,
	0xa7, 0xa1, 0x64, 0x41, 0x92, 0xae, 0x14, 0x25, 0x63, 0x18, 0x27, 0xa3, 0x5b, 0x16, 0x3b, 0xcb,
	0x4a, 0x5e, 0x31, 0x5c, 0x2a, 0x8a, 0x54, 0x4f, 0x13, 0x31, 0xe8, 0x46, 0x54, 0x0c, 0x9c, 0x15,
	0x6d, 0x5b, 0x12, 0xae, 0xa8, 0x18, 0x48, 0xe9, 0x5e, 0xcc, 0xa8, 0x60, 0x5d, 0x11, 0x8c, 0x98,
	0xd3, 0xd4, 0xea, 0x35, 0xe9, 0x3a, 0x18, 0x31, 0xfc, 0x15, 0x6c, 0xa4, 0x10, 0x9d, 0x07, 0xd1,
	0x17, 0xd0, 0x79, 0x64, 0x5f, 0x18, 0xcd, 0xbf, 0x57, 0x61, 0x55, 0x3d, 0xc2, 0x27, 0x83, 0x30,
	0xda, 0x03, 0xb8, 0x63, 0x54, 0x24, 0x31, 0xbb, 0xa6, 0x7d, 0x13, 0x7d, 0x8b, 0x82, 0x6f, 0x60,
	0xcd, 0x78, 0xf0, 0x1e, 0x58, 0xef, 0xe5, 0x60, 0x0d, 0x1a, 0x52, 0xef, 0x85, 0xf3, 0xcf, 0xd0,
	0xca, 0x98, 0x3e, 0x10, 0xc6, 0x4f, 0x6d, 0x18, 0xb7, 0xb4, 0xcd, 0x0c, 0xbe, 0x7f, 0xd6, 0xa1,
	0x99, 0x52, 0xd0, 0x2e, 0x34, 0x55, 0xd5, 0xec, 0x06, 0xbe, 0x79, 0xb0, 0x15, 0xb5, 0x3f, 0xf7,
	0x73, 0xa1, 0x59, 0x2a, 0x84, 0xe6, 0x19, 0x80, 0x16, 0xb3, 0x02, 0xd7, 0x52, 0x94, 0xcb, 0xdc,
	0xb1, 0x98, 0x44, 0x1a, 0xb7, 0xe9, 0xf1, 0xf5, 0x24, 0x62, 0xe8, 0x39, 0xac, 0x9a, 0x30, 0x76,
	0x05, 0xed, 0x1f, 0x19, 0xe4, 0xb6, 0x1f, 0x43, 0x7b, 0x54, 0x60, 0x39, 0x56, 0xd8, 0xcd, 0xb1,
	0x1c, 0x17, 0x58, 0x4e, 0x0c, 0x7e, 0x2d, 0x96, 0x13, 0x09, 0x61, 0x4e, 0x47, 0x91, 0xcc, 0x11,
	0x89, 0x70, 0x03, 0x61, 0x4d, 0x4a, 0x31, 0x6e, 0x18, 0x54, 0x6e, 0xb6, 0x6c, 0x06, 0x95, 0x9d,
	0xcf, 0x61, 0x95, 0x8b, 0x30, 0xa6, 0x7d, 0xa3, 0x02, 0xb4, 0x11, 0x43, 0x53, 0x3a, 0x2c, 0x16,
	0xa5, 0xa4, 0x9d, 0x63, 0x51, 0x5a, 0xf6, 0xa1, 0xcd, 0x7e, 0x89, 0x82, 0xd8, 0xa4, 0xd2, 0xaa,
	0x42, 0x3b, 0x68, 0x92, 0x4c, 0x25, 0xb4, 0x05, 0x8d, 0x28, 0x0e, 0x7a, 0xcc, 0x59, 0x53, 0x2f,
	0xa9, 0x37, 0xe8, 0x00, 0xda, 0x3e, 0xe3, 0xbd, 0x38, 0x88, 0x44, 0x10, 0x8e, 0x9d, 0x75, 0xad,
	0xd8, 0x22, 0x49, 0xc5, 0x89, 0x2a, 0x42, 0x5d, 0x9f, 0x0a, 0xe6, 0x6c, 0x68, 0xc5, 0x9a, 0x74,
	0x4a, 0x05, 0x2b, 0x26, 0x71, 0x67, 0x2a, 0x89, 0x7f, 0x83, 0x2d, 0x05, 0x86, 0x97, 0xc3, 0xe1,
	0x4b, 0xf9, 0x40, 0x56, 0x3a, 0xcd, 0xac, 0x67, 0xcf, 0x00, 0x38, 0xe3, 0x3c, 0x08, 0xc7, 0x12,
	0x36, 0x1a, 0x1b, 0x2d, 0x43, 0x39, 0xf7, 0xd1, 0x97, 0x00, 0xf7, 0x52, 0x65, 0x37, 0xa2, 0xb1,
	0x4e, 0xda, 0x0c, 0xf8, 0x57, 0x34, 0xa6, 0x5e, 0xeb, 0x3e, 0x5d, 0xe2, 0x37, 0x06, 0xeb, 0x72,
	0xb3, 0x80, 0x49, 0xfc, 0x16, 0xb6, 0x0b, 0xb7, 0x58, 0xb8, 0xc0, 0xbc, 0x80, 0xb5, 0x1f, 0x43,
	0x3f, 0xb8, 0x9b, 0xcc, 0x53, 0xdb, 0x7e, 0x80, 0xf5, 0x94, 0x79, 0x61, 0xc3, 0x5e, 0x61, 0x04,
	0x79, 0x27, 0xe8, 0x5c, 0xaf, 0xe2, 0x42, 0xd6, 0x37, 0xec, 0x26, 0x23, 0xf7, 0xf8, 0x06, 0x76,
	0x4b, 0x74, 0x2e, 0xec, 0xea, 0x4f, 0xb0, 0x7d, 0xc6, 0xc4, 0x69, 0xf8, 0x30, 0xbe, 0x98, 0x1a,
	0xb1, 0x3e, 0xc8, 0x4f, 0x0f, 0x76, 0x8a, 0x0a, 0x17, 0x76, 0xf2, 0xd6, 0x74, 0xeb, 0x57, 0x93,
	0xf3, 0xd3, 0xc5, 0x9a, 0x85, 0x03, 0xa6, 0x38, 0x9e, 0x1a, 0x03, 0xe9, 0x16, 0xff, 0x55, 0x85,
	0xfd, 0x33, 0x26, 0x6e, 0x38, 0x8b, 0xaf, 0x92, 0xb8, 0x37, 0xa0, 0x9c, 0x29, 0x1c, 0x5e, 0x04,
	0x5c, 0x7c, 0xba, 0x06, 0x65, 0x17, 0xef, 0x7a, 0xae, 0x78, 0xe3, 0x08, 0x0e, 0x66, 0xfb, 0xf3,
	0x7f, 0xa7, 0xb0, 0x54, 0xc5, 0x5c, 0x53, 0x98, 0xcd, 0xfc, 0x11, 0xa7, 0xb0, 0x54, 0x6d, 0xd6,
	0xc6, 0xfe, 0xae, 0x9a, 0x77, 0xb5, 0x4e, 0x16, 0xa9, 0x5a, 0x76, 0x30, 0x6b, 0xf9, 0x4e, 0xb8,
	0x07, 0x6d, 0x79, 0xe9, 0x6e, 0xcc, 0xee, 0x1f, 0x43, 0xdd, 0xf2, 0xf5, 0xfc, 0xac, 0x3b, 0x65,
	0x2f, 0x1c, 0xf3, 0x64, 0xc4, 0x62, 0xd3, 0xcb, 0xb2, 0xfd, 0xf1, 0x3f, 0x0d, 0x68, 0xa8, 0xd0,
	0xa3, 0x73, 0xe8, 0x14, 0xbf, 0x16, 0xc8, 0x21, 0x33, 0x3e, 0x29, 0xee, 0x2e, 0x99, 0xf5, 0x0f,
	0xc1, 0x15, 0x74, 0x01, 0x9b, 0x53, 0xe9, 0x8c, 0x0a, 0x12, 0x56, 0xd9, 0x70, 0x5d, 0x32, 0x33,
	0xfb, 0x71, 0x05, 0x7d, 0x0b, 0xab, 0xf6, 0x57, 0x03, 0x6d, 0x91, 0x92, 0xcf, 0x8a, 0xbb, 0x4d,
	0xca, 0xfe, 0x23, 0xda, 0x99, 0xe9, 0xf1, 0x7b, 0x97, 0xcc, 0xfa, 0x6b, 0xb8, 0x2e, 0x99, 0xf9,
	0x31, 0xc0, 0x15, 0xf4, 0x1a, 0xd6, 0xf3, 0x15, 0x00, 0xed, 0x90, 0xd2, 0x1a, 0xe3, 0x3e, 0x21,
	0xe5, 0xa5, 0x02, 0x57, 0xd0, 0x11, 0x34, 0x53, 0x67, 0x51, 0x87, 0x14, 0x86, 0x54, 0x77, 0x93,
	0x14, 0xe7, 0x50, 0x5c, 0x41, 0x87, 0xd0, 0x50, 0x6e, 0xa1, 0x35, 0x62, 0x8f, 0x95, 0xee, 0x3a,
	0xc9, 0xcd, 0x78, 0xb8, 0x82, 0xbe, 0x33, 0x63, 0x5f, 0xda, 0x6b, 0xd0, 0x36, 0x29, 0xeb, 0xa0,
	0xee, 0x0e, 0x29, 0x6d, 0x49, 0xb8, 0x82, 0x5e, 0xc0, 0xb2, 0xee, 0x16, 0x68, 0x9d, 0xe4, 0x7a,
	0x8c, 0xbb, 0x41, 0xf2, 0x6d, 0x04, 0x57, 0xd0, 0xd7, 0xa6, 0x45, 0xca, 0xf2, 0x85, 0x4c, 0x32,
	0x58, 0xa5, 0xac, 0xc4, 0xc1, 0x4b, 0xd5, 0x40, 0x4a, 0x73, 0x1f, 0x1d, 0x90, 0xff, 0x28, 0x53,
	0xd3, 0xfa, 0x6e, 0x97, 0xd5, 0x8f, 0xfa, 0xe4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xc3,
	0x1e, 0x18, 0x60, 0x0f, 0x00, 0x00,
}
