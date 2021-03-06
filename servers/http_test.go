package servers_test

import (
	"github.com/jamieabc/go-simple-http-for-test/servers"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	url := ":5566"
	s := servers.NewHttp(url)
	defer s.Stop()

	m := http.NewServeMux()
	h := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "written\n")
	}
	m.HandleFunc("/hello", h)
	srv := httptest.NewServer(m)
	defer srv.Close()

	s.Run(m)

	resp, err := http.Get("http://localhost:5566/hello")
	defer resp.Body.Close()

	assert.Nil(t, err, "wrong request")

	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err, "wrong read")
	assert.Equal(t, []byte("written\n"), data, "wrong content")
}

func TestStop(t *testing.T) {
	url := ":5566"
	s := servers.NewHttp(url)

	m := http.NewServeMux()
	h := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "hello world\n")
	}
	m.HandleFunc("/", h)
	srv := httptest.NewServer(m)
	defer srv.Close()

	s.Run(m)
	s.Stop()

	_, err := http.Get("http://localhost:5566/")
	assert.NotNil(t, err, "wrong Stop")
}
