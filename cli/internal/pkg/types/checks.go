package types

import (
	"encoding/json"
)

type CheckResult uint16

const (
	Failed        CheckResult = iota
	Passed                    = iota
	Errored                   = iota
	NotApplicable             = iota
)

func (c CheckResult) String() string {
	switch c {
	case Failed:
		return "Failed"
	case Passed:
		return "Passed"
	case Errored:
		return "Errored"
	case NotApplicable:
		return "Not Applicable"
	default:
		return "Unknown"
	}
}

func (c CheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

type CheckError struct {
	Err        error             `json:"-"`
	Check      CheckType         `json:"check"`
	Violations map[string]string `json:"violations"`
}

type CheckType string

const (
	GoCGuardrails = "GoCGuardrails"
)

type CheckReport struct {
	EntityType 	string						`json:"entity_type"`
	EntityId   	string						`json:"entity_id"`
	Timestamp	string						`json:"rfc3339_timestamp"`
	Results    	map[CheckType]CheckResult	`json:"results"`
	Errors     	[]CheckError				`json:"errors"`
}

type ICheckable interface {
	Check(checkTypes []CheckType) CheckReport
}
