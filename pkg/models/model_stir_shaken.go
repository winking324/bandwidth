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
)




// StirShaken struct for StirShaken
type StirShaken struct {
	// (optional) The verification status indicating whether the verification was successful or not. Possible values are TN-Verification-Passed and TN-Verification-Failed.
	Verstat *string `json:"verstat,omitempty"`
	// (optional) The attestation level verified by Bandwidth. Possible values are A (full), B (partial) or C (gateway).
	AttestationIndicator *string `json:"attestationIndicator,omitempty"`
	// (optional) A unique origination identifier.
	OriginatingId *string `json:"originatingId,omitempty"`
}

// NewStirShaken instantiates a new StirShaken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStirShaken() *StirShaken {
	this := StirShaken{}
	return &this
}

// NewStirShakenWithDefaults instantiates a new StirShaken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStirShakenWithDefaults() *StirShaken {
	this := StirShaken{}
	return &this
}

// GetVerstat returns the Verstat field value if set, zero value otherwise.
func (o *StirShaken) GetVerstat() string {
	if o == nil || utils.IsNil(o.Verstat) {
		var ret string
		return ret
	}
	return *o.Verstat
}

// GetVerstatOk returns a tuple with the Verstat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StirShaken) GetVerstatOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Verstat) {
		return nil, false
	}
	return o.Verstat, true
}

// HasVerstat returns a boolean if a field has been set.
func (o *StirShaken) HasVerstat() bool {
	if o != nil && !utils.IsNil(o.Verstat) {
		return true
	}

	return false
}

// SetVerstat gets a reference to the given string and assigns it to the Verstat field.
func (o *StirShaken) SetVerstat(v string) {
	o.Verstat = &v
}

// GetAttestationIndicator returns the AttestationIndicator field value if set, zero value otherwise.
func (o *StirShaken) GetAttestationIndicator() string {
	if o == nil || utils.IsNil(o.AttestationIndicator) {
		var ret string
		return ret
	}
	return *o.AttestationIndicator
}

// GetAttestationIndicatorOk returns a tuple with the AttestationIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StirShaken) GetAttestationIndicatorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AttestationIndicator) {
		return nil, false
	}
	return o.AttestationIndicator, true
}

// HasAttestationIndicator returns a boolean if a field has been set.
func (o *StirShaken) HasAttestationIndicator() bool {
	if o != nil && !utils.IsNil(o.AttestationIndicator) {
		return true
	}

	return false
}

// SetAttestationIndicator gets a reference to the given string and assigns it to the AttestationIndicator field.
func (o *StirShaken) SetAttestationIndicator(v string) {
	o.AttestationIndicator = &v
}

// GetOriginatingId returns the OriginatingId field value if set, zero value otherwise.
func (o *StirShaken) GetOriginatingId() string {
	if o == nil || utils.IsNil(o.OriginatingId) {
		var ret string
		return ret
	}
	return *o.OriginatingId
}

// GetOriginatingIdOk returns a tuple with the OriginatingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StirShaken) GetOriginatingIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginatingId) {
		return nil, false
	}
	return o.OriginatingId, true
}

// HasOriginatingId returns a boolean if a field has been set.
func (o *StirShaken) HasOriginatingId() bool {
	if o != nil && !utils.IsNil(o.OriginatingId) {
		return true
	}

	return false
}

// SetOriginatingId gets a reference to the given string and assigns it to the OriginatingId field.
func (o *StirShaken) SetOriginatingId(v string) {
	o.OriginatingId = &v
}

func (o StirShaken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StirShaken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Verstat) {
		toSerialize["verstat"] = o.Verstat
	}
	if !utils.IsNil(o.AttestationIndicator) {
		toSerialize["attestationIndicator"] = o.AttestationIndicator
	}
	if !utils.IsNil(o.OriginatingId) {
		toSerialize["originatingId"] = o.OriginatingId
	}
	return toSerialize, nil
}

type NullableStirShaken struct {
	value *StirShaken
	isSet bool
}

func (v NullableStirShaken) Get() *StirShaken {
	return v.value
}

func (v *NullableStirShaken) Set(val *StirShaken) {
	v.value = val
	v.isSet = true
}

func (v NullableStirShaken) IsSet() bool {
	return v.isSet
}

func (v *NullableStirShaken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStirShaken(val *StirShaken) *NullableStirShaken {
	return &NullableStirShaken{value: val, isSet: true}
}

func (v NullableStirShaken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStirShaken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


