package mocks

import "testing"

/////////////////////////////////////////////////////////////
// Testing With Mocks
/////////////////////////////////////////////////////////////

type mockSender struct {
	SendCallCount int
	SentTo string
}

func (m *mockSender) Send(addr string, message string) error {
	m.SendCallCount++
	m.SentTo = addr
	return nil
}

func TestWelcome(t *testing.T) {
	s := &mockSender{}
	Welcome("chris@hello.com", s)

	if s.SendCallCount != 1 {
		t.Error("Send should be called once")
	}
	if s.SentTo != "chris@hello.com" {
		t.Error("Sent to wrong address")
	}
}
