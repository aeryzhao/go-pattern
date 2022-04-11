package main

type iShot interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shot struct {
	logo string
	size int
}

func (s *shot) setLogo(logo string) {
	s.logo = logo
}

func (s *shot) setSize(size int) {
	s.size = size
}

func (s *shot) getLogo() string {
	return s.logo
}

func (s *shot) getSize() int {
	return s.size
}
