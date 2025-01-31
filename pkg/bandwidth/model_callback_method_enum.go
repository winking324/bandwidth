/*
Bandwidth

Bandwidth's Communication APIs

API version: 1.0.0
Contact: letstalk@bandwidth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bandwidth

import (
	"encoding/json"
	"fmt"
)

// CallbackMethodEnum The HTTP method to use to deliver the callback. GET or POST. Default value is POST.
type CallbackMethodEnum string

// List of callbackMethodEnum
const (
	CallbackMethodGet CallbackMethodEnum = "GET"
	CallbackMethodPost CallbackMethodEnum = "POST"
)

// All allowed values of CallbackMethodEnum enum
var AllowedCallbackMethodEnumEnumValues = []CallbackMethodEnum{
	"GET",
	"POST",
}

func (v *CallbackMethodEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CallbackMethodEnum(value)
	for _, existing := range AllowedCallbackMethodEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CallbackMethodEnum", value)
}

// NewCallbackMethodEnumFromValue returns a pointer to a valid CallbackMethodEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCallbackMethodEnumFromValue(v string) (*CallbackMethodEnum, error) {
	ev := CallbackMethodEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CallbackMethodEnum: valid values are %v", v, AllowedCallbackMethodEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CallbackMethodEnum) IsValid() bool {
	for _, existing := range AllowedCallbackMethodEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to callbackMethodEnum value
func (v CallbackMethodEnum) Ptr() *CallbackMethodEnum {
	return &v
}

type NullableCallbackMethodEnum struct {
	value *CallbackMethodEnum
	isSet bool
}

func (v NullableCallbackMethodEnum) Get() *CallbackMethodEnum {
	return v.value
}

func (v *NullableCallbackMethodEnum) Set(val *CallbackMethodEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCallbackMethodEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCallbackMethodEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallbackMethodEnum(val *CallbackMethodEnum) *NullableCallbackMethodEnum {
	return &NullableCallbackMethodEnum{value: val, isSet: true}
}

func (v NullableCallbackMethodEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallbackMethodEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

