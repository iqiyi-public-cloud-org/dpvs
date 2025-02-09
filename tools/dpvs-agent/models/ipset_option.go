// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpsetOption IpsetOption defines common options for ipset operations.
//
// swagger:model IpsetOption
type IpsetOption struct {

	// When add members to ipset with Force set, the already existing members are replaced sliently instead of emitting an EDPVS_EXIST error; When delete non-existent memebers from ipset, the DPSVS_NOTEXIST error is ignored.
	//
	Force *bool `json:"Force,omitempty"`

	// Nomatch excludes a small element range from an ipset, which is mainly used by network-cidr based ipset.
	//
	NoMatch *bool `json:"NoMatch,omitempty"`
}

// Validate validates this ipset option
func (m *IpsetOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipset option based on context it is used
func (m *IpsetOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpsetOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpsetOption) UnmarshalBinary(b []byte) error {
	var res IpsetOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
