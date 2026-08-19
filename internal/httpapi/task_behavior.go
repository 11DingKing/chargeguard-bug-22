package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

var evidenceStore charging.EvidenceStore

func ResetTaskHTTPState()              { evidenceStore = charging.EvidenceStore{} }
func SaveTaskEvidence(photos [][]byte) { evidenceStore.Save(photos) }
func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"first": evidenceStore.First()})
}
