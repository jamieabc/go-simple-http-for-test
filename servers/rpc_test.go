package servers_test

import (
	"github.com/jamieabc/go-simple-http-for-test/servers"
	"github.com/stretchr/testify/assert"
	"net/rpc"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	port := ":9988"
	s := servers.NewRpc(port)
	s.Run(nil)

	c, err := rpc.DialHTTP("tcp", port)
	assert.Nil(t, err, "wrong Dial")

	arg := servers.RpcArg{
		A: 5,
		B: 6,
	}

	var reply int
	err = c.Call("Calculator.Add", &arg, &reply)

	assert.Nil(t, err, "wrong Call")
	assert.Equal(t, arg.A+arg.B, reply, "wrong result")
}
