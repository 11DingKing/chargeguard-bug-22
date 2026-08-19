package charging

type EvidenceStore struct{ photos [][]byte }

func (s *EvidenceStore) Save(photos [][]byte) {
	s.photos = make([][]byte, len(photos))
	for i, p := range photos {
		s.photos[i] = append([]byte(nil), p...)
	}
}
func (s *EvidenceStore) First() string {
	if len(s.photos) == 0 {
		return ""
	}
	return string(s.photos[0])
}
