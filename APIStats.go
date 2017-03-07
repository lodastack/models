package models

import (
	"regexp"
)

// HTTPResponse struct
type HTTPResponse struct {
	Name                string `json:"name"`
	Interval            string `json:"interval"`
	MeasurementType     string `json:"measurement_type"`
	Comment             string `json:"comment"`
	Address             string `json:"address"`
	Body                string `json:"body"`
	Method              string `json:"method"`
	ResponseTimeout     string `json:"responseTimeout"`
	ResponseStringMatch string `json:"responseStringMatch"`
	CompiledStringMatch *regexp.Regexp

	InsecureSkipVerify bool
}
