package models

type MockScheme struct {
	Port  string
	Paths []*Path
}

func (ms *MockScheme) IsValid() error {
	for _, p := range ms.Paths {
		if err := p.Endpoints.IsValid(); err != nil {
			return err
		}
	}
	return nil
}
