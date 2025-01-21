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




// LookupStatus If requestId exists, the result for that request is returned. See the Examples for details on the various responses that you can receive. Generally, if you see a Response Code of 0 in a result for a TN, information will be available for it.  Any other Response Code will indicate no information was available for the TN.
type LookupStatus struct {
	// The requestId.
	RequestId *string `json:"requestId,omitempty"`
	Status *LookupStatusEnum `json:"status,omitempty"`
	// The carrier information results for the specified telephone number.
	Result []LookupResult `json:"result,omitempty"`
	// The telephone numbers whose lookup failed.
	FailedTelephoneNumbers []string `json:"failedTelephoneNumbers,omitempty"`
}

// NewLookupStatus instantiates a new LookupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupStatus() *LookupStatus {
	this := LookupStatus{}
	return &this
}

// NewLookupStatusWithDefaults instantiates a new LookupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupStatusWithDefaults() *LookupStatus {
	this := LookupStatus{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *LookupStatus) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupStatus) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *LookupStatus) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *LookupStatus) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LookupStatus) GetStatus() LookupStatusEnum {
	if o == nil || utils.IsNil(o.Status) {
		var ret LookupStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupStatus) GetStatusOk() (*LookupStatusEnum, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LookupStatus) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given LookupStatusEnum and assigns it to the Status field.
func (o *LookupStatus) SetStatus(v LookupStatusEnum) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *LookupStatus) GetResult() []LookupResult {
	if o == nil || utils.IsNil(o.Result) {
		var ret []LookupResult
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupStatus) GetResultOk() ([]LookupResult, bool) {
	if o == nil || utils.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *LookupStatus) HasResult() bool {
	if o != nil && !utils.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []LookupResult and assigns it to the Result field.
func (o *LookupStatus) SetResult(v []LookupResult) {
	o.Result = v
}

// GetFailedTelephoneNumbers returns the FailedTelephoneNumbers field value if set, zero value otherwise.
func (o *LookupStatus) GetFailedTelephoneNumbers() []string {
	if o == nil || utils.IsNil(o.FailedTelephoneNumbers) {
		var ret []string
		return ret
	}
	return o.FailedTelephoneNumbers
}

// GetFailedTelephoneNumbersOk returns a tuple with the FailedTelephoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupStatus) GetFailedTelephoneNumbersOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.FailedTelephoneNumbers) {
		return nil, false
	}
	return o.FailedTelephoneNumbers, true
}

// HasFailedTelephoneNumbers returns a boolean if a field has been set.
func (o *LookupStatus) HasFailedTelephoneNumbers() bool {
	if o != nil && !utils.IsNil(o.FailedTelephoneNumbers) {
		return true
	}

	return false
}

// SetFailedTelephoneNumbers gets a reference to the given []string and assigns it to the FailedTelephoneNumbers field.
func (o *LookupStatus) SetFailedTelephoneNumbers(v []string) {
	o.FailedTelephoneNumbers = v
}

func (o LookupStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !utils.IsNil(o.FailedTelephoneNumbers) {
		toSerialize["failedTelephoneNumbers"] = o.FailedTelephoneNumbers
	}
	return toSerialize, nil
}

type NullableLookupStatus struct {
	value *LookupStatus
	isSet bool
}

func (v NullableLookupStatus) Get() *LookupStatus {
	return v.value
}

func (v *NullableLookupStatus) Set(val *LookupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupStatus(val *LookupStatus) *NullableLookupStatus {
	return &NullableLookupStatus{value: val, isSet: true}
}

func (v NullableLookupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


