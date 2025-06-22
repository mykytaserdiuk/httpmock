package models

import "errors"

type MockScheme struct {
	Port  string
	Paths Paths
}

func (ms *MockScheme) IsValid() error {
	if ms.Paths == nil || len(ms.Paths) < 1 {
		return errors.New("paths is empty")
	}
	if ms.Port == "" {
		return errors.New("port is empty")
	}
	if err := ms.Paths.IsValid(); err != nil {
		return err
	}
	return nil
}
