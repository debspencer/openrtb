package openrtb

import "errors"

// The Metric object is associated with an impression as an array of metrics.
// These metrics can offer insigt into the impression to assist with decisioning
// such as average recent viewability, click-through rate, etc.  Each metric is
// identified by its type, reports the value of the metric, and optionally
// identifies the source or vendor measuring the value (openrtb 2.5 only feature)

type Metric struct {
	Type     string        `json:"type"`              // Type
	Value    float64       `json:"value"`             // Value
	Vendor   string        `json:"vendor,omitempty"`  //Vendor, recommended.
	Ext      Extension     `json:"ext,omitempty"`
}

// Validation errors
var (
	ErrInvalidMetric = errors.New("openrtb: metric needs a type and value")
)

// Validate required attributes
func (mt *Metric) Validate() error {
	if len(mt.Type) == 0 {
		return ErrInvalidMetric
	}

	return nil
}
