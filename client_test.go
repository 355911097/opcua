package opcua

import (
	"github.com/pascaldekloe/goe/verify"
	"scada_da/pkg/scada/OPCUA/opcua/ua"
	"testing"
)

func TestClient_Send_DoesNotPanicWhenDisconnected(t *testing.T) {
	c := NewClient("opc.tcp://example.com:4840")
	err := c.Send(&ua.ReadRequest{}, func(i interface{}) error {
		return nil
	})
	verify.Values(t, "", err, ua.StatusBadServerNotConnected)
}
