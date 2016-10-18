package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChartResourceAttributes chart resource attributes

swagger:model chartResourceAttributes
*/
type ChartResourceAttributes struct {

	/* created

	Required: true
	Min Length: 1
	*/
	Created *string `json:"created"`

	/* description

	Required: true
	Min Length: 1
	*/
	Description *string `json:"description"`

	/* digest

	Required: true
	Min Length: 1
	*/
	Digest *string `json:"digest"`

	/* home

	Required: true
	Min Length: 1
	*/
	Home *string `json:"home"`

	/* name

	Required: true
	Min Length: 1
	*/
	Name *string `json:"name"`

	/* repo

	Required: true
	*/
	Repo *string `json:"repo"`

	/* sources

	Required: true
	*/
	Sources []string `json:"sources"`

	/* urls

	Required: true
	*/
	Urls []string `json:"urls"`
}

// Validate validates this chart resource attributes
func (m *ChartResourceAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDigest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHome(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUrls(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartResourceAttributes) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.MinLength("created", "body", string(*m.Created), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateDigest(formats strfmt.Registry) error {

	if err := validate.Required("digest", "body", m.Digest); err != nil {
		return err
	}

	if err := validate.MinLength("digest", "body", string(*m.Digest), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateHome(formats strfmt.Registry) error {

	if err := validate.Required("home", "body", m.Home); err != nil {
		return err
	}

	if err := validate.MinLength("home", "body", string(*m.Home), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateRepo(formats strfmt.Registry) error {

	if err := validate.Required("repo", "body", m.Repo); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	return nil
}

func (m *ChartResourceAttributes) validateUrls(formats strfmt.Registry) error {

	if err := validate.Required("urls", "body", m.Urls); err != nil {
		return err
	}

	return nil
}
