package handlers_test

import (
	"github.com/jamieabc/go-simple-http-for-test/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	h := handlers.New()
	s := httptest.NewServer(h)
	defer s.Close()

	res, err := http.Get(s.URL + "/hello")
	assert.Nil(t, err, "wrong Get")

	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "wrong read")
	assert.Equal(t, []byte("hello, world!\n"), data, "wrong content")
}
