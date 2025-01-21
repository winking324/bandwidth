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




// LookupRequest Create phone number lookup request.
type LookupRequest struct {
	Tns []string `json:"tns"`
}

type _LookupRequest LookupRequest

// NewLookupRequest instantiates a new LookupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupRequest(tns []string) *LookupRequest {
	this := LookupRequest{}
	this.Tns = tns
	return &this
}

// NewLookupRequestWithDefaults instantiates a new LookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupRequestWithDefaults() *LookupRequest {
	this := LookupRequest{}
	return &this
}

// GetTns returns the Tns field value
func (o *LookupRequest) GetTns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tns
}

// GetTnsOk returns a tuple with the Tns field value
// and a boolean to check if the value has been set.
func (o *LookupRequest) GetTnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tns, true
}

// SetTns sets field value
func (o *LookupRequest) SetTns(v []string) {
	o.Tns = v
}

func (o LookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tns"] = o.Tns
	return toSerialize, nil
}

func (o *LookupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tns",
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

	varLookupRequest := _LookupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLookupRequest)

	if err != nil {
		return err
	}

	*o = LookupRequest(varLookupRequest)

	return err
}

type NullableLookupRequest struct {
	value *LookupRequest
	isSet bool
}

func (v NullableLookupRequest) Get() *LookupRequest {
	return v.value
}

func (v *NullableLookupRequest) Set(val *LookupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupRequest(val *LookupRequest) *NullableLookupRequest {
	return &NullableLookupRequest{value: val, isSet: true}
}

func (v NullableLookupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


