package models

import "errors"

type Path struct {
	Path      string `yaml:"path"`
	Endpoints Endpoints
}

func (p *Path) IsValid() error {
	if p.Path == "" {
		return errors.New("one of path is empty")
	}
	if err := p.Endpoints.IsValid(); err != nil {
		return err
	}
	return nil
}

type Paths []*Path

func (p Paths) IsValid() error {
	if len(p) == 0 {
		return errors.New("zero paths")
	}
	for _, p := range p {
		if err := p.IsValid(); err != nil {
			return err
		}
	}
	return nil
}
