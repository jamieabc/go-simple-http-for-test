package server_test

import (
	"github.com/jamieabc/go-simple-http-for-test/server"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRun(t *testing.T) {
	url := ":5566"
	s := server.New(url)
	s.Run()

	resp, err := http.Get("http://localhost:5566/hello")
	assert.Nil(t, err, "wrong request")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err, "wrong read")

	assert.Equal(t, []byte("hello, world!\n"), body, "wrong response content")
}
