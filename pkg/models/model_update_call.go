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




// UpdateCall struct for UpdateCall
type UpdateCall struct {
	State NullableCallStateEnum `json:"state,omitempty"`
	// The URL to send the [Redirect](/docs/voice/bxml/redirect) event to which will provide new BXML.  Required if `state` is `active`.  Not allowed if `state` is `completed`.
	RedirectUrl utils.NullableString `json:"redirectUrl,omitempty"`
	RedirectMethod NullableRedirectMethodEnum `json:"redirectMethod,omitempty"`
	// Basic auth username.
	Username utils.NullableString `json:"username,omitempty"`
	// Basic auth password.
	Password utils.NullableString `json:"password,omitempty"`
	// A fallback url which, if provided, will be used to retry the redirect callback delivery in case `redirectUrl` fails to respond.
	RedirectFallbackUrl utils.NullableString `json:"redirectFallbackUrl,omitempty"`
	RedirectFallbackMethod NullableRedirectMethodEnum `json:"redirectFallbackMethod,omitempty"`
	// Basic auth username.
	FallbackUsername utils.NullableString `json:"fallbackUsername,omitempty"`
	// Basic auth password.
	FallbackPassword utils.NullableString `json:"fallbackPassword,omitempty"`
	// A custom string that will be sent with this and all future callbacks unless overwritten by a future `tag` attribute or [`<Tag>`](/docs/voice/bxml/tag) verb, or cleared.  May be cleared by setting `tag=\"\"`.  Max length 256 characters.  Not allowed if `state` is `completed`.
	Tag utils.NullableString `json:"tag,omitempty"`
}

// NewUpdateCall instantiates a new UpdateCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCall() *UpdateCall {
	this := UpdateCall{}
	var state CallStateEnum = ACTIVE
	this.State = *NewNullableCallStateEnum(&state)
	var redirectMethod RedirectMethodEnum = POST
	this.RedirectMethod = *NewNullableRedirectMethodEnum(&redirectMethod)
	var redirectFallbackMethod RedirectMethodEnum = POST
	this.RedirectFallbackMethod = *NewNullableRedirectMethodEnum(&redirectFallbackMethod)
	return &this
}

// NewUpdateCallWithDefaults instantiates a new UpdateCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCallWithDefaults() *UpdateCall {
	this := UpdateCall{}
	var state CallStateEnum = ACTIVE
	this.State = *NewNullableCallStateEnum(&state)
	var redirectMethod RedirectMethodEnum = POST
	this.RedirectMethod = *NewNullableRedirectMethodEnum(&redirectMethod)
	var redirectFallbackMethod RedirectMethodEnum = POST
	this.RedirectFallbackMethod = *NewNullableRedirectMethodEnum(&redirectFallbackMethod)
	return &this
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetState() CallStateEnum {
	if o == nil || utils.IsNil(o.State.Get()) {
		var ret CallStateEnum
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetStateOk() (*CallStateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *UpdateCall) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableCallStateEnum and assigns it to the State field.
func (o *UpdateCall) SetState(v CallStateEnum) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *UpdateCall) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *UpdateCall) UnsetState() {
	o.State.Unset()
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetRedirectUrl() string {
	if o == nil || utils.IsNil(o.RedirectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UpdateCall) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given utils.NullableString and assigns it to the RedirectUrl field.
func (o *UpdateCall) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *UpdateCall) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *UpdateCall) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetRedirectMethod returns the RedirectMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetRedirectMethod() RedirectMethodEnum {
	if o == nil || utils.IsNil(o.RedirectMethod.Get()) {
		var ret RedirectMethodEnum
		return ret
	}
	return *o.RedirectMethod.Get()
}

// GetRedirectMethodOk returns a tuple with the RedirectMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetRedirectMethodOk() (*RedirectMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectMethod.Get(), o.RedirectMethod.IsSet()
}

// HasRedirectMethod returns a boolean if a field has been set.
func (o *UpdateCall) HasRedirectMethod() bool {
	if o != nil && o.RedirectMethod.IsSet() {
		return true
	}

	return false
}

// SetRedirectMethod gets a reference to the given NullableRedirectMethodEnum and assigns it to the RedirectMethod field.
func (o *UpdateCall) SetRedirectMethod(v RedirectMethodEnum) {
	o.RedirectMethod.Set(&v)
}
// SetRedirectMethodNil sets the value for RedirectMethod to be an explicit nil
func (o *UpdateCall) SetRedirectMethodNil() {
	o.RedirectMethod.Set(nil)
}

// UnsetRedirectMethod ensures that no value is present for RedirectMethod, not even an explicit nil
func (o *UpdateCall) UnsetRedirectMethod() {
	o.RedirectMethod.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetUsername() string {
	if o == nil || utils.IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateCall) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given utils.NullableString and assigns it to the Username field.
func (o *UpdateCall) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UpdateCall) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UpdateCall) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetPassword() string {
	if o == nil || utils.IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateCall) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given utils.NullableString and assigns it to the Password field.
func (o *UpdateCall) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UpdateCall) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UpdateCall) UnsetPassword() {
	o.Password.Unset()
}

// GetRedirectFallbackUrl returns the RedirectFallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetRedirectFallbackUrl() string {
	if o == nil || utils.IsNil(o.RedirectFallbackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectFallbackUrl.Get()
}

// GetRedirectFallbackUrlOk returns a tuple with the RedirectFallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetRedirectFallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectFallbackUrl.Get(), o.RedirectFallbackUrl.IsSet()
}

// HasRedirectFallbackUrl returns a boolean if a field has been set.
func (o *UpdateCall) HasRedirectFallbackUrl() bool {
	if o != nil && o.RedirectFallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectFallbackUrl gets a reference to the given utils.NullableString and assigns it to the RedirectFallbackUrl field.
func (o *UpdateCall) SetRedirectFallbackUrl(v string) {
	o.RedirectFallbackUrl.Set(&v)
}
// SetRedirectFallbackUrlNil sets the value for RedirectFallbackUrl to be an explicit nil
func (o *UpdateCall) SetRedirectFallbackUrlNil() {
	o.RedirectFallbackUrl.Set(nil)
}

// UnsetRedirectFallbackUrl ensures that no value is present for RedirectFallbackUrl, not even an explicit nil
func (o *UpdateCall) UnsetRedirectFallbackUrl() {
	o.RedirectFallbackUrl.Unset()
}

// GetRedirectFallbackMethod returns the RedirectFallbackMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetRedirectFallbackMethod() RedirectMethodEnum {
	if o == nil || utils.IsNil(o.RedirectFallbackMethod.Get()) {
		var ret RedirectMethodEnum
		return ret
	}
	return *o.RedirectFallbackMethod.Get()
}

// GetRedirectFallbackMethodOk returns a tuple with the RedirectFallbackMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetRedirectFallbackMethodOk() (*RedirectMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectFallbackMethod.Get(), o.RedirectFallbackMethod.IsSet()
}

// HasRedirectFallbackMethod returns a boolean if a field has been set.
func (o *UpdateCall) HasRedirectFallbackMethod() bool {
	if o != nil && o.RedirectFallbackMethod.IsSet() {
		return true
	}

	return false
}

// SetRedirectFallbackMethod gets a reference to the given NullableRedirectMethodEnum and assigns it to the RedirectFallbackMethod field.
func (o *UpdateCall) SetRedirectFallbackMethod(v RedirectMethodEnum) {
	o.RedirectFallbackMethod.Set(&v)
}
// SetRedirectFallbackMethodNil sets the value for RedirectFallbackMethod to be an explicit nil
func (o *UpdateCall) SetRedirectFallbackMethodNil() {
	o.RedirectFallbackMethod.Set(nil)
}

// UnsetRedirectFallbackMethod ensures that no value is present for RedirectFallbackMethod, not even an explicit nil
func (o *UpdateCall) UnsetRedirectFallbackMethod() {
	o.RedirectFallbackMethod.Unset()
}

// GetFallbackUsername returns the FallbackUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetFallbackUsername() string {
	if o == nil || utils.IsNil(o.FallbackUsername.Get()) {
		var ret string
		return ret
	}
	return *o.FallbackUsername.Get()
}

// GetFallbackUsernameOk returns a tuple with the FallbackUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetFallbackUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackUsername.Get(), o.FallbackUsername.IsSet()
}

// HasFallbackUsername returns a boolean if a field has been set.
func (o *UpdateCall) HasFallbackUsername() bool {
	if o != nil && o.FallbackUsername.IsSet() {
		return true
	}

	return false
}

// SetFallbackUsername gets a reference to the given utils.NullableString and assigns it to the FallbackUsername field.
func (o *UpdateCall) SetFallbackUsername(v string) {
	o.FallbackUsername.Set(&v)
}
// SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil
func (o *UpdateCall) SetFallbackUsernameNil() {
	o.FallbackUsername.Set(nil)
}

// UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
func (o *UpdateCall) UnsetFallbackUsername() {
	o.FallbackUsername.Unset()
}

// GetFallbackPassword returns the FallbackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetFallbackPassword() string {
	if o == nil || utils.IsNil(o.FallbackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.FallbackPassword.Get()
}

// GetFallbackPasswordOk returns a tuple with the FallbackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetFallbackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackPassword.Get(), o.FallbackPassword.IsSet()
}

// HasFallbackPassword returns a boolean if a field has been set.
func (o *UpdateCall) HasFallbackPassword() bool {
	if o != nil && o.FallbackPassword.IsSet() {
		return true
	}

	return false
}

// SetFallbackPassword gets a reference to the given utils.NullableString and assigns it to the FallbackPassword field.
func (o *UpdateCall) SetFallbackPassword(v string) {
	o.FallbackPassword.Set(&v)
}
// SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil
func (o *UpdateCall) SetFallbackPasswordNil() {
	o.FallbackPassword.Set(nil)
}

// UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil
func (o *UpdateCall) UnsetFallbackPassword() {
	o.FallbackPassword.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCall) GetTag() string {
	if o == nil || utils.IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCall) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *UpdateCall) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given utils.NullableString and assigns it to the Tag field.
func (o *UpdateCall) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *UpdateCall) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *UpdateCall) UnsetTag() {
	o.Tag.Unset()
}

func (o UpdateCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirectUrl"] = o.RedirectUrl.Get()
	}
	if o.RedirectMethod.IsSet() {
		toSerialize["redirectMethod"] = o.RedirectMethod.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.RedirectFallbackUrl.IsSet() {
		toSerialize["redirectFallbackUrl"] = o.RedirectFallbackUrl.Get()
	}
	if o.RedirectFallbackMethod.IsSet() {
		toSerialize["redirectFallbackMethod"] = o.RedirectFallbackMethod.Get()
	}
	if o.FallbackUsername.IsSet() {
		toSerialize["fallbackUsername"] = o.FallbackUsername.Get()
	}
	if o.FallbackPassword.IsSet() {
		toSerialize["fallbackPassword"] = o.FallbackPassword.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	return toSerialize, nil
}

type NullableUpdateCall struct {
	value *UpdateCall
	isSet bool
}

func (v NullableUpdateCall) Get() *UpdateCall {
	return v.value
}

func (v *NullableUpdateCall) Set(val *UpdateCall) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCall) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCall(val *UpdateCall) *NullableUpdateCall {
	return &NullableUpdateCall{value: val, isSet: true}
}

func (v NullableUpdateCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


