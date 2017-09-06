package pipeline

import (
	"strings"

	"github.com/pkg/errors"

	"golang.org/x/sync/syncmap"
)

var steps syncmap.Map

func FromName(s string) (Step, error) {
	s = strings.ToLower(s)
	val, ok := steps.Load(s)
	if !ok {
		log.WithField("Step", s).
			Warn("cannot find Step")
		return nil, errors.Errorf("cannot find Step %v", s)
	}
	Step, ok := val.(Step)
	if !ok {
		log.WithField("Step", s).
			Warn("invalid Step")
		return nil, errors.Errorf("invalid Step %v", s)
	}
	return Step, nil
}

func Register(s Step) {
	steps.Store(strings.ToLower(s.Info()), s)
}

func Steps() []string {
	names := []string{}
	steps.Range(func(key, _ interface{}) bool {
		if name, ok := key.(string); ok {
			names = append(names, name)
		}
		return true
	})
	return names
}
