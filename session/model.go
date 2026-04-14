package session

type session struct {
	store map[string]string
}

func NewSession() *session {

	return &session{store: make(map[string]string)}
}

// one of the things is uid--->email
func (s *session) Get(key string) (string, bool) {

	val, ok := (s.store)[key]

	return val, ok
}

func (s *session) Put(key, val string) {

	s.store[key] = val
}
