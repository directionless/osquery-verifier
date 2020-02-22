package echoserver

import (
	"context"
	"fmt"

	"github.com/directionless/osquery-verifier/pkg/osquery"
	"github.com/kolide/kit/ulid"
)

type server struct{}

func New() *server {
	s := &server{}
	return s
}

func (s *server) Enroll(_ context.Context, req osquery.EnrollmentRequest) (*osquery.EnrollmentResponse, error) {
	resp := &osquery.EnrollmentResponse{NodeKey: ulid.New()}
	fmt.Printf("enrollment from %s. Returned key %s\n", req.Secret, resp.NodeKey)

	return resp, nil
}

func (s *server) Config(_ context.Context, req osquery.ConfigRequest) (*osquery.ConfigResponse, error) {
	queries := map[string]osquery.Query{
		"time": osquery.Query{
			SQL:      "select timestamp from time",
			Interval: 2,
		},
	}
	resp := &osquery.ConfigResponse{
		Schedule: queries,
	}
	fmt.Printf("config from %s. Returned...\n", req.NodeKey)

	return resp, nil
}

func (s *server) Log(_ context.Context, req osquery.LogRequest) (*osquery.LogResponse, error) {
	fmt.Printf("logs from %s. type: %s. length: %d\n", req.NodeKey, req.LogType, len(req.Data))

	resp := &osquery.LogResponse{}
	return resp, nil
}

func (s *server) DistributedRead(_ context.Context, req osquery.DistributedReadRequest) (*osquery.DistributedReadResponse, error) {
	fmt.Printf("DistributedRead from %s. Returned...\n", req.NodeKey)

	resp := &osquery.DistributedReadResponse{}
	return resp, nil

}

func (s *server) DistributedWrite(_ context.Context, req osquery.DistributedWriteRequest) (*osquery.DistributedWriteResponse, error) {
	fmt.Printf("DistributedWrite from %s. Returned...\n", req.NodeKey)

	resp := &osquery.DistributedWriteResponse{}
	return resp, nil
}

func (s *server) CarveInit(_ context.Context, req osquery.CarveInitRequest) (*osquery.CarveInitResponse, error) {
	return nil, nil
}

func (s *server) CarveBlock(_ context.Context, req osquery.CarveBlockRequest) (*osquery.CarveBlockResponse, error) {
	return nil, nil
}
