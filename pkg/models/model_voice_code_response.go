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




// VoiceCodeResponse struct for VoiceCodeResponse
type VoiceCodeResponse struct {
	// Programmable Voice API Call ID.
	CallId *string `json:"callId,omitempty"`
}

// NewVoiceCodeResponse instantiates a new VoiceCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceCodeResponse() *VoiceCodeResponse {
	this := VoiceCodeResponse{}
	return &this
}

// NewVoiceCodeResponseWithDefaults instantiates a new VoiceCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceCodeResponseWithDefaults() *VoiceCodeResponse {
	this := VoiceCodeResponse{}
	return &this
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *VoiceCodeResponse) GetCallId() string {
	if o == nil || utils.IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceCodeResponse) GetCallIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *VoiceCodeResponse) HasCallId() bool {
	if o != nil && !utils.IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *VoiceCodeResponse) SetCallId(v string) {
	o.CallId = &v
}

func (o VoiceCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CallId) {
		toSerialize["callId"] = o.CallId
	}
	return toSerialize, nil
}

type NullableVoiceCodeResponse struct {
	value *VoiceCodeResponse
	isSet bool
}

func (v NullableVoiceCodeResponse) Get() *VoiceCodeResponse {
	return v.value
}

func (v *NullableVoiceCodeResponse) Set(val *VoiceCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceCodeResponse(val *VoiceCodeResponse) *NullableVoiceCodeResponse {
	return &NullableVoiceCodeResponse{value: val, isSet: true}
}

func (v NullableVoiceCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


