package mocks

import "fmt"

/////////////////////////////////////////////////////////////
// Mocks
/////////////////////////////////////////////////////////////

type Sender interface {
	Send(addr string, message string) error
}

func Welcome(addr string, s Sender) error {
	if err := s.Send(addr, "Welcome"); err != nil {
		return fmt.Errorf("Send welcome failed %w", err)
	}

	return nil
}
