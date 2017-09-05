package flow

import (
	"strings"

	"github.com/pkg/errors"

	"golang.org/x/sync/syncmap"
)

var processes syncmap.Map

func FromName(s string) (Process, error) {
	s = strings.ToLower(s)
	val, ok := processes.Load(s)
	if !ok {
		log.WithField("process", s).
			Warn("cannot find process")
		return nil, errors.Errorf("cannot find process %v", s)
	}
	process, ok := val.(Process)
	if !ok {
		log.WithField("process", s).
			Warn("invalid process")
		return nil, errors.Errorf("invalid process %v", s)
	}
	return process, nil
}

func Register(name string, s Process) {
	processes.Store(strings.ToLower(name), s)
}

func Processes() []string {
	names := []string{}
	processes.Range(func(key, _ interface{}) bool {
		if name, ok := key.(string); ok {
			names = append(names, name)
		}
		return true
	})
	return names
}
