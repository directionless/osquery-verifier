package osquery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

type tlsServer struct {
	remote      OsqueryRemote
	errReporter func(error)
}

func NewTLSServer(osqRemote OsqueryRemote) *tlsServer {
	s := &tlsServer{
		remote:      osqRemote,
		errReporter: func(err error) { fmt.Println(err) },
	}

	return s
}

func (s *tlsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	upath := path.Clean(r.URL.Path)

	switch upath {
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
		s.handleError(w, fmt.Errorf("404: %s", upath))
	}

}

func (s *tlsServer) handleError(w http.ResponseWriter, err error) {
	if err != nil {
		s.errReporter(err)
	}

	http.Error(
		w,
		"{}",
		http.StatusInternalServerError,
	)
}

// commonEncodeAndSend is a common bit of error checking and json response encoding
func (s *tlsServer) commonEncodeAndSend(name string, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		s.handleError(w, fmt.Errorf("handling %s: %w", name, err))
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		s.handleError(w, fmt.Errorf("json encode %s: %w", name, err))
		return

	}

}

func (s *tlsServer) handleEnroll(w http.ResponseWriter, r *http.Request) {
	req := EnrollmentRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode enrollment: %w", err))
		return
	}

	resp, err := s.remote.Enroll(r.Context(), req)
	s.commonEncodeAndSend("enroll", w, resp, err)
}

func (s *tlsServer) handleConfig(w http.ResponseWriter, r *http.Request) {
	req := ConfigRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode config: %w", err))
		return
	}

	resp, err := s.remote.Config(r.Context(), req)
	s.commonEncodeAndSend("enroll", w, resp, err)
}

func (s *tlsServer) handleLog(w http.ResponseWriter, r *http.Request) {
	req := LogRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode log: %w", err))
		return

	}
	resp, err := s.remote.Log(r.Context(), req)
	s.commonEncodeAndSend("log", w, resp, err)
}

func (s *tlsServer) handleDistributedRead(w http.ResponseWriter, r *http.Request) {
	req := DistributedReadRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode distributed read: %w", err))
		return

	}
	resp, err := s.remote.DistributedRead(r.Context(), req)
	s.commonEncodeAndSend("distributed read", w, resp, err)

}

func (s *tlsServer) handleDistributedWrite(w http.ResponseWriter, r *http.Request) {
	req := DistributedWriteRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode distributed write: %w", err))
		return

	}
	resp, err := s.remote.DistributedWrite(r.Context(), req)
	s.commonEncodeAndSend("distributed write", w, resp, err)
}

func (s *tlsServer) handleCarveInit(w http.ResponseWriter, r *http.Request) {
	req := CarveInitRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode carve init: %w", err))
		return

	}
	resp, err := s.remote.CarveInit(r.Context(), req)
	s.commonEncodeAndSend("carve init", w, resp, err)
}

func (s *tlsServer) handleCarveBlock(w http.ResponseWriter, r *http.Request) {
	req := CarveBlockRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handleError(w, fmt.Errorf("json decode carve block: %w", err))
		return

	}
	resp, err := s.remote.CarveBlock(r.Context(), req)
	s.commonEncodeAndSend("carve block", w, resp, err)
}
