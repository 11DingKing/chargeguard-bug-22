package charging

type EvidenceStore struct{ photos [][]byte }

// Save stores a defensive copy of the submitted photos so that evidence already
// persisted cannot change when the caller later reuses the same client slice for
// subsequent uploads. Both the outer slice and each photo's bytes are copied,
// isolating the saved result from any further mutation of the caller's data.
func (s *EvidenceStore) Save(photos [][]byte) {
	copied := make([][]byte, len(photos))
	for i, p := range photos {
		copied[i] = append([]byte(nil), p...)
	}
	s.photos = copied
}
func (s *EvidenceStore) First() string {
	if len(s.photos) == 0 {
		return ""
	}
	return string(s.photos[0])
}
