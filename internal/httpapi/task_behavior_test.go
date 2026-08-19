package httpapi

import (
	"net/http/httptest"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	ResetTaskHTTPState()
	photos := [][]byte{[]byte("extinguisher"), []byte("barrier"), []byte("sign")}
	SaveTaskEvidence(photos)
	copy(photos[0], []byte("overwritten!!"))
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("GET", "/task", nil))
	if rr.Body.String() != "{\"first\":\"extinguisher\"}\n" {
		t.Fatalf("body=%s", rr.Body.String())
	}
}
