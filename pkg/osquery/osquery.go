package osquery

import "context"

// OsqueryRemote provices a remote interface to osquery
type OsqueryRemote interface {
	Enroll(context.Context, EnrollmentRequest) (EnrollmentResponse, error)
	Config(context.Context, ConfigRequest) (ConfigResponse, error)
	Log(context.Context, LogRequest) (LogResponse, error)
	DistributedRead(context.Context, DistributedReadRequest) (DistributedReadResponse, error)
	DistributedWrite(context.Context, DistributedWriteRequest) (DistributedWriteResponse, error)
	CarveInit(context.Context, CarveInitRequest) (CarveInitResponse, error)
	CarveBlock(context.Context, CarveBlockRequest) (CarveBlockResponse, error)
}

type EnrollmentRequest struct {
	EnrollSecret      string `json:"enroll_secret"`
	HostIdentifier    string `json:"host_identifier"`
	EnrollmentDetails map[string]string
}

type EnrollmentResponse struct{}

type ConfigRequest struct{}
type ConfigResponse struct{}

type LogRequest struct{}
type LogResponse struct{}

type DistributedReadRequest struct{}
type DistributedReadResponse struct{}

type DistributedWriteRequest struct{}
type DistributedWriteResponse struct{}

type CarveInitRequest struct{}
type CarveInitResponse struct{}

type CarveBlockRequest struct{}
type CarveBlockResponse struct{}
