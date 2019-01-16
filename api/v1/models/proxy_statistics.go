// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProxyStatistics Statistics of a set of proxy redirects for an endpoint
// swagger:model ProxyStatistics
// +k8s:deepcopy-gen=true
type ProxyStatistics struct {

	// The port the proxy is listening on
	AllocatedProxyPort int64 `json:"allocated-proxy-port,omitempty"`

	// Location of where the redirect is installed
	Location string `json:"location,omitempty"`

	// The port subject to the redirect
	Port int64 `json:"port,omitempty"`

	// Name of the L7 protocol
	Protocol string `json:"protocol,omitempty"`

	// Statistics of this set of proxy redirect
	Statistics *RequestResponseStatistics `json:"statistics,omitempty"`
}

/* polymorph ProxyStatistics allocated-proxy-port false */

/* polymorph ProxyStatistics location false */

/* polymorph ProxyStatistics port false */

/* polymorph ProxyStatistics protocol false */

/* polymorph ProxyStatistics statistics false */

// Validate validates this proxy statistics
func (m *ProxyStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var proxyStatisticsTypeLocationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ingress","egress"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		proxyStatisticsTypeLocationPropEnum = append(proxyStatisticsTypeLocationPropEnum, v)
	}
}

const (
	// ProxyStatisticsLocationIngress captures enum value "ingress"
	ProxyStatisticsLocationIngress string = "ingress"
	// ProxyStatisticsLocationEgress captures enum value "egress"
	ProxyStatisticsLocationEgress string = "egress"
)

// prop value enum
func (m *ProxyStatistics) validateLocationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, proxyStatisticsTypeLocationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProxyStatistics) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocationEnum("location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

func (m *ProxyStatistics) validateStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {

		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyStatistics) UnmarshalBinary(b []byte) error {
	var res ProxyStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
