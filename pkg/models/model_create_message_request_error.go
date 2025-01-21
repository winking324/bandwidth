/*
Bandwidth

Bandwidth's Communication APIs

API version: 1.0.0
Contact: letstalk@bandwidth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// CreateMessageRequestError struct for CreateMessageRequestError
type CreateMessageRequestError struct {
	Type string `json:"type"`
	Description string `json:"description"`
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
}

type _CreateMessageRequestError CreateMessageRequestError

// NewCreateMessageRequestError instantiates a new CreateMessageRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageRequestError(type_ string, description string) *CreateMessageRequestError {
	this := CreateMessageRequestError{}
	this.Type = type_
	this.Description = description
	return &this
}

// NewCreateMessageRequestErrorWithDefaults instantiates a new CreateMessageRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageRequestErrorWithDefaults() *CreateMessageRequestError {
	this := CreateMessageRequestError{}
	return &this
}

// GetType returns the Type field value
func (o *CreateMessageRequestError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequestError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateMessageRequestError) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *CreateMessageRequestError) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequestError) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateMessageRequestError) SetDescription(v string) {
	o.Description = v
}

// GetFieldErrors returns the FieldErrors field value if set, zero value otherwise.
func (o *CreateMessageRequestError) GetFieldErrors() []FieldError {
	if o == nil || utils.IsNil(o.FieldErrors) {
		var ret []FieldError
		return ret
	}
	return o.FieldErrors
}

// GetFieldErrorsOk returns a tuple with the FieldErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMessageRequestError) GetFieldErrorsOk() ([]FieldError, bool) {
	if o == nil || utils.IsNil(o.FieldErrors) {
		return nil, false
	}
	return o.FieldErrors, true
}

// HasFieldErrors returns a boolean if a field has been set.
func (o *CreateMessageRequestError) HasFieldErrors() bool {
	if o != nil && !utils.IsNil(o.FieldErrors) {
		return true
	}

	return false
}

// SetFieldErrors gets a reference to the given []FieldError and assigns it to the FieldErrors field.
func (o *CreateMessageRequestError) SetFieldErrors(v []FieldError) {
	o.FieldErrors = v
}

func (o CreateMessageRequestError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageRequestError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["description"] = o.Description
	if !utils.IsNil(o.FieldErrors) {
		toSerialize["fieldErrors"] = o.FieldErrors
	}
	return toSerialize, nil
}

func (o *CreateMessageRequestError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"description",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateMessageRequestError := _CreateMessageRequestError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMessageRequestError)

	if err != nil {
		return err
	}

	*o = CreateMessageRequestError(varCreateMessageRequestError)

	return err
}

type NullableCreateMessageRequestError struct {
	value *CreateMessageRequestError
	isSet bool
}

func (v NullableCreateMessageRequestError) Get() *CreateMessageRequestError {
	return v.value
}

func (v *NullableCreateMessageRequestError) Set(val *CreateMessageRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageRequestError(val *CreateMessageRequestError) *NullableCreateMessageRequestError {
	return &NullableCreateMessageRequestError{value: val, isSet: true}
}

func (v NullableCreateMessageRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


