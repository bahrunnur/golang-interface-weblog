package happyconsumer_test

import (
	"testing"

	happyconsumer "whatev.com/interface/happy-consumer"
)

type mockClient struct{}

func (c mockClient) AddToken(name string)          {}
func (c mockClient) GetTokenCount(name string) int { return 0 }

func Test_TokenService(t *testing.T) {
	ts := happyconsumer.NewTokenService(mockClient{})
	if ts == nil {
		t.Fail()
	}
}
