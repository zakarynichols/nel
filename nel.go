package nel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// The regex expression used in this example is checking that the value is composed
// of a list of comma-separated strings that only contain alphanumeric characters
// and dash. If the format is not correct, it will return an error.
const CommaSeparatedAlphanumericWithDashRegex = "^([a-zA-Z0-9-]+)(,[a-zA-Z0-9-]+)*$"

// NEL struct holds the fields for the NEL header
type NEL struct {
	ReportTo          string  `json:"report_to,omitempty"`
	MaxAge            int     `json:"max_age"`
	FailureFraction   float64 `json:"failure_fraction,omitempty"`
	SuccessFraction   float64 `json:"success_fraction,omitempty"`
	Expires           string  `json:"expires,omitempty"`
	IncludeSubdomains bool    `json:"include_subdomains,omitempty"`
	RequestHeaders    string  `json:"request_headers,omitempty"`
	ResponseHeaders   string  `json:"response_headers,omitempty"`
}

// SetNEL adds the NEL header to a response
func SetNEL(w http.ResponseWriter, n *NEL) error {
	if err := n.validate(); err != nil {
		return err
	}
	w.Header().Set("NEL", n.toJSON())
	return nil
}

// RemoveNEL removes the policy its reporting to
func RemoveNEL(w http.ResponseWriter, n *NEL) {
	n.Remove()
	w.Header().Set("NEL", n.toJSON())
}

func (n *NEL) validReportTo() error {
	if n.ReportTo == "" {
		return fmt.Errorf("report_to field is required")
	}
	return nil
}

func (n *NEL) validateMaxAge() error {
	if n.MaxAge == 0 {
		return fmt.Errorf("NEL policy removed")
	}
	if n.MaxAge < 0 {
		return fmt.Errorf("max_age must be non-negative")
	}
	return nil
}

func (n *NEL) validateSuccessFraction() error {
	if n.SuccessFraction != 0.0 && (n.SuccessFraction < 0.0 || n.SuccessFraction > 1.0) {
		return fmt.Errorf("success_fraction must be a number between 0.0 and 1.0")
	}
	return nil
}

func (n *NEL) validateFailureFraction() error {
	if n.SuccessFraction != 0.0 && (n.SuccessFraction < 0.0 || n.SuccessFraction > 1.0) {
		return fmt.Errorf("success_fraction must be a number between 0.0 and 1.0")
	}
	return nil
}

func (n *NEL) validateIncludeSubdomains() error {
	if !n.IncludeSubdomains {
		return fmt.Errorf("include_subdomains must be a boolean value")
	}
	return nil
}

func (n *NEL) validateRequestHeaders() error {
	if n.RequestHeaders != "" {
		// check if the request headers match a specific format
		if !regexp.MustCompile(CommaSeparatedAlphanumericWithDashRegex).MatchString(n.RequestHeaders) {
			return fmt.Errorf("request_headers must be a list of comma-separated header field names")
		}
	}
	return nil
}

func (n *NEL) validateResponseHeaders() error {
	if n.ResponseHeaders != "" {
		// check if the response headers match a specific format
		if !regexp.MustCompile(CommaSeparatedAlphanumericWithDashRegex).MatchString(n.ResponseHeaders) {
			return fmt.Errorf("response_headers must be a list of comma-separated header field names")
		}
	}
	return nil
}

func (n *NEL) validate() error {
	var err error
	err = n.validReportTo()
	if err != nil {
		return err
	}
	err = n.validateMaxAge()
	if err != nil {
		return err
	}
	err = n.validateSuccessFraction()
	if err != nil {
		return err
	}
	err = n.validateFailureFraction()
	if err != nil {
		return err
	}
	err = n.validateIncludeSubdomains()
	if err != nil {
		return err
	}
	err = n.validateRequestHeaders()
	if err != nil {
		return err
	}
	err = n.validateResponseHeaders()
	if err != nil {
		return err
	}
	return nil
}

func (n *NEL) Remove() {
	n.MaxAge = 0
}

func (n *NEL) toJSON() string {
	b, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(b)
}
