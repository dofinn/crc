// automatically generated by stateify.

// +build go1.9
// +build !go1.17

package tcpip

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *StdClock) StateTypeName() string {
	return "pkg/tcpip.StdClock"
}

func (s *StdClock) StateFields() []string {
	return []string{}
}

func (s *StdClock) beforeSave() {}

func (s *StdClock) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
}

func (s *StdClock) afterLoad() {}

func (s *StdClock) StateLoad(stateSourceObject state.Source) {
}

func init() {
	state.Register((*StdClock)(nil))
}
