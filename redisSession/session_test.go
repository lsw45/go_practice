package redisSession

import (
	"bytes"
	"encoding/base64"
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ResponseRecorder is an implementation of http.ResponseWriter that
// records its mutations for later inspection in tests.
type ResponseRecorder struct {
	Code      int           // the HTTP response code from WriteHeader
	HeaderMap http.Header   // the HTTP response headers
	Body      *bytes.Buffer // if non-nil, the bytes.Buffer to append written data to
	Flushed   bool
}

// Header returns the response headers.
func (rw *ResponseRecorder) Header() http.Header {
	return rw.HeaderMap
}

// Write always succeeds and writes to rw.Body, if not nil.
func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
	if rw.Body != nil {
		rw.Body.Write(buf)
	}
	if rw.Code == 0 {
		rw.Code = http.StatusOK
	}
	return len(buf), nil
}

// WriteHeader sets rw.Code.
func (rw *ResponseRecorder) WriteHeader(code int) {
	rw.Code = code
}

// RedisStore
func TestRedisStore(t *testing.T) {
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer store.Close()

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	rsp := &ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
	}
	// Get a session.
	var session *sessions.Session
	if session, err = store.Get(req, "session-key"); err != nil {
		t.Fatalf("Error getting session: %v", err)
	}
	// Get a flash.
	flashes := session.Flashes()
	if len(flashes) != 0 {
		t.Errorf("Expected empty flashes; Got %v", flashes)
	}
	// Add some flashes.
	session.AddFlash("foo")
	session.AddFlash("bar")
	// Custom key.
	session.AddFlash("baz", "custom_key")
	// Save.
	if err = sessions.Save(req, rsp); err != nil {
		t.Fatalf("Error saving session: %v", err)
	}
	hdr := rsp.HeaderMap
	cookies, ok := hdr["Set-Cookie"]
	if !ok || len(cookies) != 1 {
		t.Fatalf("No cookies. Header: %v", hdr)
	}
}

// RediStore change MaxLength of session
func TestSession(t *testing.T) {
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		t.Fatal(err.Error())
	}
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	writeRecord := httptest.NewRecorder()

	session, err := store.New(req, "my session")
	if err != nil {
		t.Fatal("failed to create session", err)
	}
	session.Values["big"] = make([]byte, base64.StdEncoding.DecodedLen(4096*2))
	err = session.Save(req, writeRecord)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	store.SetMaxLength(4096 * 3) // A bit more than the value size to account for encoding overhead.
	err = session.Save(req, writeRecord)
	if err != nil {
		t.Fatal("failed to Save:", err)
	}
}
