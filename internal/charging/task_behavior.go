package charging

type EvidenceStore struct{ photos [][]byte }

func (s *EvidenceStore) Save(photos [][]byte) { s.photos = photos }
func (s *EvidenceStore) First() string {
	if len(s.photos) == 0 {
		return ""
	}
	return string(s.photos[0])
}
