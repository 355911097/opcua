package opcua

import (
	"github.com/pascaldekloe/goe/verify"
	"github.com/355911097/opcua/ua"
	"testing"
)

func TestClient_Send_DoesNotPanicWhenDisconnected(t *testing.T) {
	c := NewClient("opc.tcp://example.com:4840")
	err := c.Send(&ua.ReadRequest{}, func(i interface{}) error {
		return nil
	})
	verify.Values(t, "", err, ua.StatusBadServerNotConnected)
}
