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

// MessageDirectionEnum The direction of the message. One of in out.
type MessageDirectionEnum string

// List of messageDirectionEnum
const (
	IN MessageDirectionEnum = "in"
	OUT MessageDirectionEnum = "out"
)

// All allowed values of MessageDirectionEnum enum
var AllowedMessageDirectionEnumEnumValues = []MessageDirectionEnum{
	"in",
	"out",
}

func (v *MessageDirectionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageDirectionEnum(value)
	for _, existing := range AllowedMessageDirectionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageDirectionEnum", value)
}

// NewMessageDirectionEnumFromValue returns a pointer to a valid MessageDirectionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageDirectionEnumFromValue(v string) (*MessageDirectionEnum, error) {
	ev := MessageDirectionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageDirectionEnum: valid values are %v", v, AllowedMessageDirectionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageDirectionEnum) IsValid() bool {
	for _, existing := range AllowedMessageDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to messageDirectionEnum value
func (v MessageDirectionEnum) Ptr() *MessageDirectionEnum {
	return &v
}

type NullableMessageDirectionEnum struct {
	value *MessageDirectionEnum
	isSet bool
}

func (v NullableMessageDirectionEnum) Get() *MessageDirectionEnum {
	return v.value
}

func (v *NullableMessageDirectionEnum) Set(val *MessageDirectionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDirectionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDirectionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDirectionEnum(val *MessageDirectionEnum) *NullableMessageDirectionEnum {
	return &NullableMessageDirectionEnum{value: val, isSet: true}
}

func (v NullableMessageDirectionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDirectionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

