package slightlyunhappyconsumer_test

import (
	"testing"

	slightlyunhappyconsumer "whatev.com/interface/slightly-unhappy-consumer"
)

type mockClient struct{}

func (c mockClient) AddToken(name string)                    {}
func (c mockClient) GetTokenCount(name string) int           { return 0 }
func (c mockClient) UpdateTokenName(old, new string)         {}
func (c mockClient) UpdateTokenCount(name string, count int) {}
func (c mockClient) RemoveToken(name string)                 {}
func (c mockClient) ResetToken(name string)                  {}
func (c mockClient) Size() int                               { return 0 }

func Test_TokenService(t *testing.T) {
	ts := slightlyunhappyconsumer.NewTokenService(mockClient{})
	if ts == nil {
		t.Fail()
	}
}
