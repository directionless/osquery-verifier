package osquery

import (
	"context"

	"github.com/kolide/osquery-go/plugin/distributed"
)

// OsqueryRemote provices a remote interface to osquery
type OsqueryRemote interface {
	Enroll(context.Context, EnrollmentRequest) (*EnrollmentResponse, error)
	Config(context.Context, ConfigRequest) (*ConfigResponse, error)
	Log(context.Context, LogRequest) (*LogResponse, error)
	DistributedRead(context.Context, DistributedReadRequest) (*DistributedReadResponse, error)
	DistributedWrite(context.Context, DistributedWriteRequest) (*DistributedWriteResponse, error)
	CarveInit(context.Context, CarveInitRequest) (*CarveInitResponse, error)
	CarveBlock(context.Context, CarveBlockRequest) (*CarveBlockResponse, error)
}

type EnrollmentRequest struct {
	Secret         string `json:"enroll_secret"`
	HostIdentifier string `json:"host_identifier"`
	Details        map[string]string
}

type EnrollmentResponse struct {
	NodeKey     string `json:"node_key"`
	NodeInvalid bool   `json:"node_invalid"`
	ErrorCode   string `json:"error_code,omitempty"`
}

type ConfigRequest struct {
	NodeKey string `json:"node_key"`
}

// ConfigResponse is a json config. No wrapping.
type ConfigResponse struct {
	Schedule    map[string]Query `json:"schedule"`
	ErrorCode   string           `json:"error_code"`
	NodeInvalid bool             `json:"node_invalid"`
}

type LogRequest struct {
	NodeKey string `json:"node_key"`
	LogType string `json:"log_type"`
	//Data    json.RawMessage `json:"data"`
	Data []map[string]interface{} `json:"data"`
}
type LogResponse struct {
	ErrorCode   string `json:"error_code,omitempty"`
	NodeInvalid bool   `json:"node_invalid"`
}

type DistributedReadRequest struct {
	NodeKey string `json:"node_key"`
}

type DistributedReadResponse struct {
	Queries     distributed.GetQueriesResult
	Accelerate  uint   `json:"accelerate,omitempty"`
	NodeInvalid bool   `json:"node_invalid"`
	ErrorCode   string `json:"error_code,omitempty"`
}

type DistributedWriteRequest struct {
	NodeKey string `json:"node_key"`
	//Results  kolide.OsqueryDistributedQueryResults `json:"queries"`
	//Statuses map[string]kolide.OsqueryStatus `json:"statuses"`
}

type DistributedWriteResponse struct {
	ErrorCode string `json:"error_code,omitempty"`
}

type CarveInitRequest struct{}
type CarveInitResponse struct{}

type CarveBlockRequest struct{}
type CarveBlockResponse struct{}

type Query struct {
	SQL      string `json:"query"`
	Interval int    `json:"interval"`
}
