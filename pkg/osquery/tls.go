package osquery

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
)

type tlsServer struct {
	remote OsqueryRemote
}

func NewTLSServer(osqRemote OsqueryRemote) *tlsServer {
	s := &tlsServer{
		remote: osqRemote,
	}

	return s
}

func (s *tlsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch path.Clean(r.URL.Path) {
	case "/enroll":
		s.handleEnroll(w, r)
	case "/config":
		s.handleConfig(w, r)
	case "/log":
		s.handleLog(w, r)
	case "/distributed_read":
		s.handleDistributedRead(w, r)
	case "/distributed_write":
		s.handleDistributedWrite(w, r)
	case "/carve_init":
		s.handleCarveInit(w, r)
	case "/carve_block":
		s.handleCarveBlock(w, r)
	default:
	}
}

func (s *tlsServer) handleEnroll(w http.ResponseWriter, r *http.Request) {
	req := EnrollmentRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.Enroll(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func (s *tlsServer) handleConfig(w http.ResponseWriter, r *http.Request) {
	req := ConfigRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.Config(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func (s *tlsServer) handleLog(w http.ResponseWriter, r *http.Request) {
	req := LogRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.Log(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func (s *tlsServer) handleDistributedRead(w http.ResponseWriter, r *http.Request) {
	req := DistributedReadRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.DistributedRead(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func (s *tlsServer) handleDistributedWrite(w http.ResponseWriter, r *http.Request) {
	req := DistributedWriteRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.DistributedWrite(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func (s *tlsServer) handleCarveInit(w http.ResponseWriter, r *http.Request) {
	req := CarveInitRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.CarveInit(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func (s *tlsServer) handleCarveBlock(w http.ResponseWriter, r *http.Request) {
	req := CarveBlockRequest{}
	if err := decodeJSON(r, req); err != nil {
		// FIXME
	}
	resp, err := s.remote.CarveBlock(context.TODO(), req)
	if err != nil {
		// FIXME
	}
	if err := encodeJSON(w, resp); err != nil {
		// FIXME
	}
}

func encodeJSON(w http.ResponseWriter, request interface{}) error {
	return json.NewEncoder(w).Encode(request)
}

func decodeJSON(req *http.Request, request interface{}) error {
	return json.NewDecoder(req.Body).Decode(&request)
}
