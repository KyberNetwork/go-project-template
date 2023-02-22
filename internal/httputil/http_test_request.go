package httputil

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// AssertFn is the function used to asserts the test response.
type AssertFn func(t *testing.T, resp *httptest.ResponseRecorder)

// HTTPTestCase struct for http test case.
type HTTPTestCase struct {
	Msg      string
	Endpoint string
	Method   string
	Params   map[string]string
	Body     []byte
	Assert   AssertFn
}

// RunHTTPTestCase run http request test case.
func RunHTTPTestCase(t *testing.T, tc HTTPTestCase, handler http.Handler) {
	t.Helper()
	req, err := http.NewRequestWithContext(context.Background(), tc.Method, tc.Endpoint, bytes.NewBuffer(tc.Body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for k, v := range tc.Params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	tc.Assert(t, resp)
}

// AssertCode asserts that the response matched the expected error code.
func AssertCode(code int) AssertFn {
	return func(t *testing.T, resp *httptest.ResponseRecorder) {
		t.Helper()
		if resp.Code != code {
			t.Fatalf("wrong return code, expected: %d, got %d", code, resp.Code)
		}
	}
}
