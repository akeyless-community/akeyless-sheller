package sheller

import "os"

type OSWrapper interface {
	Stat(name string) (os.FileInfo, error)
}

type RealOS struct{}

func (RealOS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
