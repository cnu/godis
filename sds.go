package godis

import "sync"

// SDS is Simple Dynamic strings. like redis' sds.
type SDS struct {
	value string
	sync.RWMutex
}

// NewSDS returns a reference to a new SDS
func NewSDS(v string) *SDS {
	return &SDS{value: v}
}

func (s *SDS) String() string {
	return s.get()
}

func (s *SDS) get() string {
	s.RLock()
	defer s.RUnlock()
	return s.value
}
