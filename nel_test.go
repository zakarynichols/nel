package nel

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetNEL(t *testing.T) {
	n := &NEL{
		ReportTo:          "group",
		MaxAge:            3600,
		FailureFraction:   0.1,
		SuccessFraction:   0.5,
		Expires:           "2022-01-01",
		IncludeSubdomains: true,
		RequestHeaders:    "header1,header2",
		ResponseHeaders:   "header3,header4",
	}

	_, _ = http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	if err := SetNEL(w, n); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if w.Header().Get("NEL") != "{\"report_to\":\"group\",\"max_age\":3600,\"failure_fraction\":0.1,\"success_fraction\":0.5,\"expires\":\"2022-01-01\",\"include_subdomains\":true,\"request_headers\":\"header1,header2\",\"response_headers\":\"header3,header4\"}" {
		t.Errorf("Unexpected NEL header value: %s", w.Header().Get("NEL"))
	}
}

func TestRemoveNEL(t *testing.T) {
	// Create a new NEL policy
	n := NEL{
		ReportTo: "default",
		MaxAge:   3600,
	}
	// Create a new HTTP response writer
	w := httptest.NewRecorder()
	// Call the RemoveNEL function
	RemoveNEL(w, &n)
	// Get the NEL header from the response
	nelHeader := w.Header().Get("NEL")
	// Check if the NEL header is set to the expected value
	expected := `{"report_to":"default","max_age":0}`
	if nelHeader != expected {
		t.Errorf("RemoveNEL() did not set the correct NEL header, expected: %s, got: %s", expected, nelHeader)
	}
	// Check if the max_age is set to 0
	if n.MaxAge != 0 {
		t.Errorf("RemoveNEL() did not set max_age to 0, got: %d", n.MaxAge)
	}
}
