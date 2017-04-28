package env

import (
	"github.com/golang/glog"
	//"os"
	"os"
	"sync"
)

type Env interface {
	Get(name string) string
	Print()
}

type envOnce struct {
	kind string
	envM map[string]string
	sync.Once
}

func NewEnv(kind string, envs ...string) Env {
	m := make(map[string]string)

	e := &envOnce{
		kind: kind,
		envM: m,
	}

	e.Do(func() {
		for _, env := range envs {
			e.envM[env] = os.Getenv(env)
		}
	})
	return e
}

func (e *envOnce) Get(name string) string {
	return e.envM[name]
}

func (e *envOnce) Print() {
	for k, v := range e.envM {
		glog.Infof("[Env] %s:{%s: %s}\n", e.kind, k, v)
	}
}
