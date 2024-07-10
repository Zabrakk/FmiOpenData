package http

import (
	"io"
	nhttp "net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFromUrl(t *testing.T) {
	server := httptest.NewServer(nhttp.HandlerFunc(func(hrw nhttp.ResponseWriter, req *nhttp.Request) {
		hrw.Write([]byte("Synop"))
	}))
	defer server.Close()

	rc, err := GetFromUrl(server.URL)
	if err != nil {
		t.Fatalf("%s", err)
	}
	b, err := io.ReadAll(rc)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if string(b[:]) != "Synop" {
		t.Fatalf("Get returned incorrect data")
	}
}
