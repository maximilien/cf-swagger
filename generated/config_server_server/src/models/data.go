package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Data the value for the data

swagger:model Data
*/
type Data struct {

	/* the string name of the type for the value
	 */
	Type *string `json:"type,omitempty"`
}

// Validate validates this data
func (m *Data) Validate(formats strfmt.Registry) error {
	return nil
}