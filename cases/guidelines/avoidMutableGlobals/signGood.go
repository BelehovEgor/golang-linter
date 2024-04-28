package avoidMutableGlobals

import "time"

type signer struct {
	now func() time.Time
}

func newSigner() *signer {
	return &signer{
		now: time.Now,
	}
}

func (s *signer) Sign(msg string) string {
	now := s.now()
	return signWithTime(msg, now)
}

func signWithTime(msg string, now time.Time) string {
	return msg + now.GoString()
}
