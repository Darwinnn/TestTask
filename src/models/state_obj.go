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

// StateObj state obj
// swagger:model stateObj
type StateObj struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// state
	// Required: true
	// Enum: [win lost]
	State *string `json:"state"`

	// transaction Id
	// Required: true
	// Format: uuid
	TransactionID *strfmt.UUID `json:"transactionId"`
}

// Validate validates this state obj
func (m *StateObj) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateObj) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

var stateObjTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["win","lost"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stateObjTypeStatePropEnum = append(stateObjTypeStatePropEnum, v)
	}
}

const (

	// StateObjStateWin captures enum value "win"
	StateObjStateWin string = "win"

	// StateObjStateLost captures enum value "lost"
	StateObjStateLost string = "lost"
)

// prop value enum
func (m *StateObj) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stateObjTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StateObj) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *StateObj) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transactionId", "body", m.TransactionID); err != nil {
		return err
	}

	if err := validate.FormatOf("transactionId", "body", "uuid", m.TransactionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateObj) UnmarshalBinary(b []byte) error {
	var res StateObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}