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
	"time"
	"bytes"
	"fmt"
)

// CreateCallResponse struct for CreateCallResponse
type CreateCallResponse struct {
	// The id of the application associated with the `from` number.
	ApplicationId string `json:"applicationId"`
	// The bandwidth account ID associated with the call.
	AccountId string `json:"accountId"`
	// Programmable Voice API Call ID.
	CallId string `json:"callId"`
	// Recipient of the outgoing call.
	To string `json:"to"`
	// Phone number that created the outbound call.
	From string `json:"from"`
	// The time at which the call was accepted into the queue.
	EnqueuedTime utils.NullableTime `json:"enqueuedTime,omitempty"`
	// The URL to update this call's state.
	CallUrl string `json:"callUrl"`
	// The timeout (in seconds) for the callee to answer the call after it starts ringing.
	CallTimeout *float64 `json:"callTimeout,omitempty"`
	// This is the timeout (in seconds) to use when delivering webhooks for the call.
	CallbackTimeout *float64 `json:"callbackTimeout,omitempty"`
	// Custom tag value.
	Tag utils.NullableString `json:"tag,omitempty"`
	AnswerMethod NullableCallbackMethodEnum `json:"answerMethod"`
	// URL to deliver the `answer` event webhook.
	AnswerUrl string `json:"answerUrl"`
	AnswerFallbackMethod NullableCallbackMethodEnum `json:"answerFallbackMethod,omitempty"`
	// Fallback URL to deliver the `answer` event webhook.
	AnswerFallbackUrl utils.NullableString `json:"answerFallbackUrl,omitempty"`
	DisconnectMethod NullableCallbackMethodEnum `json:"disconnectMethod"`
	// URL to deliver the `disconnect` event webhook.
	DisconnectUrl utils.NullableString `json:"disconnectUrl,omitempty"`
	// Basic auth username.
	Username utils.NullableString `json:"username,omitempty"`
	// Basic auth password.
	Password utils.NullableString `json:"password,omitempty"`
	// Basic auth username.
	FallbackUsername utils.NullableString `json:"fallbackUsername,omitempty"`
	// Basic auth password.
	FallbackPassword utils.NullableString `json:"fallbackPassword,omitempty"`
	// The priority of this call over other calls from your account.
	Priority NullableInt32 `json:"priority,omitempty"`
}

type _CreateCallResponse CreateCallResponse

// NewCreateCallResponse instantiates a new CreateCallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCallResponse(applicationId string, accountId string, callId string, to string, from string, callUrl string, answerMethod NullableCallbackMethodEnum, answerUrl string, disconnectMethod NullableCallbackMethodEnum) *CreateCallResponse {
	this := CreateCallResponse{}
	this.ApplicationId = applicationId
	this.AccountId = accountId
	this.CallId = callId
	this.To = to
	this.From = from
	this.CallUrl = callUrl
	this.AnswerMethod = answerMethod
	this.AnswerUrl = answerUrl
	var answerFallbackMethod CallbackMethodEnum = POST
	this.AnswerFallbackMethod = *NewNullableCallbackMethodEnum(&answerFallbackMethod)
	this.DisconnectMethod = disconnectMethod
	return &this
}

// NewCreateCallResponseWithDefaults instantiates a new CreateCallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCallResponseWithDefaults() *CreateCallResponse {
	this := CreateCallResponse{}
	var answerMethod CallbackMethodEnum = POST
	this.AnswerMethod = *NewNullableCallbackMethodEnum(&answerMethod)
	var answerFallbackMethod CallbackMethodEnum = POST
	this.AnswerFallbackMethod = *NewNullableCallbackMethodEnum(&answerFallbackMethod)
	var disconnectMethod CallbackMethodEnum = POST
	this.DisconnectMethod = *NewNullableCallbackMethodEnum(&disconnectMethod)
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *CreateCallResponse) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CreateCallResponse) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetAccountId returns the AccountId field value
func (o *CreateCallResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CreateCallResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetCallId returns the CallId field value
func (o *CreateCallResponse) GetCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallId, true
}

// SetCallId sets field value
func (o *CreateCallResponse) SetCallId(v string) {
	o.CallId = v
}

// GetTo returns the To field value
func (o *CreateCallResponse) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *CreateCallResponse) SetTo(v string) {
	o.To = v
}

// GetFrom returns the From field value
func (o *CreateCallResponse) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CreateCallResponse) SetFrom(v string) {
	o.From = v
}

// GetEnqueuedTime returns the EnqueuedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetEnqueuedTime() time.Time {
	if o == nil || utils.IsNil(o.EnqueuedTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EnqueuedTime.Get()
}

// GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetEnqueuedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnqueuedTime.Get(), o.EnqueuedTime.IsSet()
}

// HasEnqueuedTime returns a boolean if a field has been set.
func (o *CreateCallResponse) HasEnqueuedTime() bool {
	if o != nil && o.EnqueuedTime.IsSet() {
		return true
	}

	return false
}

// SetEnqueuedTime gets a reference to the given utils.NullableTime and assigns it to the EnqueuedTime field.
func (o *CreateCallResponse) SetEnqueuedTime(v time.Time) {
	o.EnqueuedTime.Set(&v)
}
// SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil
func (o *CreateCallResponse) SetEnqueuedTimeNil() {
	o.EnqueuedTime.Set(nil)
}

// UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
func (o *CreateCallResponse) UnsetEnqueuedTime() {
	o.EnqueuedTime.Unset()
}

// GetCallUrl returns the CallUrl field value
func (o *CreateCallResponse) GetCallUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallUrl
}

// GetCallUrlOk returns a tuple with the CallUrl field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetCallUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallUrl, true
}

// SetCallUrl sets field value
func (o *CreateCallResponse) SetCallUrl(v string) {
	o.CallUrl = v
}

// GetCallTimeout returns the CallTimeout field value if set, zero value otherwise.
func (o *CreateCallResponse) GetCallTimeout() float64 {
	if o == nil || utils.IsNil(o.CallTimeout) {
		var ret float64
		return ret
	}
	return *o.CallTimeout
}

// GetCallTimeoutOk returns a tuple with the CallTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetCallTimeoutOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.CallTimeout) {
		return nil, false
	}
	return o.CallTimeout, true
}

// HasCallTimeout returns a boolean if a field has been set.
func (o *CreateCallResponse) HasCallTimeout() bool {
	if o != nil && !utils.IsNil(o.CallTimeout) {
		return true
	}

	return false
}

// SetCallTimeout gets a reference to the given float64 and assigns it to the CallTimeout field.
func (o *CreateCallResponse) SetCallTimeout(v float64) {
	o.CallTimeout = &v
}

// GetCallbackTimeout returns the CallbackTimeout field value if set, zero value otherwise.
func (o *CreateCallResponse) GetCallbackTimeout() float64 {
	if o == nil || utils.IsNil(o.CallbackTimeout) {
		var ret float64
		return ret
	}
	return *o.CallbackTimeout
}

// GetCallbackTimeoutOk returns a tuple with the CallbackTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetCallbackTimeoutOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.CallbackTimeout) {
		return nil, false
	}
	return o.CallbackTimeout, true
}

// HasCallbackTimeout returns a boolean if a field has been set.
func (o *CreateCallResponse) HasCallbackTimeout() bool {
	if o != nil && !utils.IsNil(o.CallbackTimeout) {
		return true
	}

	return false
}

// SetCallbackTimeout gets a reference to the given float64 and assigns it to the CallbackTimeout field.
func (o *CreateCallResponse) SetCallbackTimeout(v float64) {
	o.CallbackTimeout = &v
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetTag() string {
	if o == nil || utils.IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *CreateCallResponse) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given utils.NullableString and assigns it to the Tag field.
func (o *CreateCallResponse) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *CreateCallResponse) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *CreateCallResponse) UnsetTag() {
	o.Tag.Unset()
}

// GetAnswerMethod returns the AnswerMethod field value
// If the value is explicit nil, the zero value for CallbackMethodEnum will be returned
func (o *CreateCallResponse) GetAnswerMethod() CallbackMethodEnum {
	if o == nil || o.AnswerMethod.Get() == nil {
		var ret CallbackMethodEnum
		return ret
	}

	return *o.AnswerMethod.Get()
}

// GetAnswerMethodOk returns a tuple with the AnswerMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetAnswerMethodOk() (*CallbackMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerMethod.Get(), o.AnswerMethod.IsSet()
}

// SetAnswerMethod sets field value
func (o *CreateCallResponse) SetAnswerMethod(v CallbackMethodEnum) {
	o.AnswerMethod.Set(&v)
}

// GetAnswerUrl returns the AnswerUrl field value
func (o *CreateCallResponse) GetAnswerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnswerUrl
}

// GetAnswerUrlOk returns a tuple with the AnswerUrl field value
// and a boolean to check if the value has been set.
func (o *CreateCallResponse) GetAnswerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerUrl, true
}

// SetAnswerUrl sets field value
func (o *CreateCallResponse) SetAnswerUrl(v string) {
	o.AnswerUrl = v
}

// GetAnswerFallbackMethod returns the AnswerFallbackMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetAnswerFallbackMethod() CallbackMethodEnum {
	if o == nil || utils.IsNil(o.AnswerFallbackMethod.Get()) {
		var ret CallbackMethodEnum
		return ret
	}
	return *o.AnswerFallbackMethod.Get()
}

// GetAnswerFallbackMethodOk returns a tuple with the AnswerFallbackMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetAnswerFallbackMethodOk() (*CallbackMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerFallbackMethod.Get(), o.AnswerFallbackMethod.IsSet()
}

// HasAnswerFallbackMethod returns a boolean if a field has been set.
func (o *CreateCallResponse) HasAnswerFallbackMethod() bool {
	if o != nil && o.AnswerFallbackMethod.IsSet() {
		return true
	}

	return false
}

// SetAnswerFallbackMethod gets a reference to the given NullableCallbackMethodEnum and assigns it to the AnswerFallbackMethod field.
func (o *CreateCallResponse) SetAnswerFallbackMethod(v CallbackMethodEnum) {
	o.AnswerFallbackMethod.Set(&v)
}
// SetAnswerFallbackMethodNil sets the value for AnswerFallbackMethod to be an explicit nil
func (o *CreateCallResponse) SetAnswerFallbackMethodNil() {
	o.AnswerFallbackMethod.Set(nil)
}

// UnsetAnswerFallbackMethod ensures that no value is present for AnswerFallbackMethod, not even an explicit nil
func (o *CreateCallResponse) UnsetAnswerFallbackMethod() {
	o.AnswerFallbackMethod.Unset()
}

// GetAnswerFallbackUrl returns the AnswerFallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetAnswerFallbackUrl() string {
	if o == nil || utils.IsNil(o.AnswerFallbackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AnswerFallbackUrl.Get()
}

// GetAnswerFallbackUrlOk returns a tuple with the AnswerFallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetAnswerFallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerFallbackUrl.Get(), o.AnswerFallbackUrl.IsSet()
}

// HasAnswerFallbackUrl returns a boolean if a field has been set.
func (o *CreateCallResponse) HasAnswerFallbackUrl() bool {
	if o != nil && o.AnswerFallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetAnswerFallbackUrl gets a reference to the given utils.NullableString and assigns it to the AnswerFallbackUrl field.
func (o *CreateCallResponse) SetAnswerFallbackUrl(v string) {
	o.AnswerFallbackUrl.Set(&v)
}
// SetAnswerFallbackUrlNil sets the value for AnswerFallbackUrl to be an explicit nil
func (o *CreateCallResponse) SetAnswerFallbackUrlNil() {
	o.AnswerFallbackUrl.Set(nil)
}

// UnsetAnswerFallbackUrl ensures that no value is present for AnswerFallbackUrl, not even an explicit nil
func (o *CreateCallResponse) UnsetAnswerFallbackUrl() {
	o.AnswerFallbackUrl.Unset()
}

// GetDisconnectMethod returns the DisconnectMethod field value
// If the value is explicit nil, the zero value for CallbackMethodEnum will be returned
func (o *CreateCallResponse) GetDisconnectMethod() CallbackMethodEnum {
	if o == nil || o.DisconnectMethod.Get() == nil {
		var ret CallbackMethodEnum
		return ret
	}

	return *o.DisconnectMethod.Get()
}

// GetDisconnectMethodOk returns a tuple with the DisconnectMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetDisconnectMethodOk() (*CallbackMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisconnectMethod.Get(), o.DisconnectMethod.IsSet()
}

// SetDisconnectMethod sets field value
func (o *CreateCallResponse) SetDisconnectMethod(v CallbackMethodEnum) {
	o.DisconnectMethod.Set(&v)
}

// GetDisconnectUrl returns the DisconnectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetDisconnectUrl() string {
	if o == nil || utils.IsNil(o.DisconnectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DisconnectUrl.Get()
}

// GetDisconnectUrlOk returns a tuple with the DisconnectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetDisconnectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisconnectUrl.Get(), o.DisconnectUrl.IsSet()
}

// HasDisconnectUrl returns a boolean if a field has been set.
func (o *CreateCallResponse) HasDisconnectUrl() bool {
	if o != nil && o.DisconnectUrl.IsSet() {
		return true
	}

	return false
}

// SetDisconnectUrl gets a reference to the given utils.NullableString and assigns it to the DisconnectUrl field.
func (o *CreateCallResponse) SetDisconnectUrl(v string) {
	o.DisconnectUrl.Set(&v)
}
// SetDisconnectUrlNil sets the value for DisconnectUrl to be an explicit nil
func (o *CreateCallResponse) SetDisconnectUrlNil() {
	o.DisconnectUrl.Set(nil)
}

// UnsetDisconnectUrl ensures that no value is present for DisconnectUrl, not even an explicit nil
func (o *CreateCallResponse) UnsetDisconnectUrl() {
	o.DisconnectUrl.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetUsername() string {
	if o == nil || utils.IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateCallResponse) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given utils.NullableString and assigns it to the Username field.
func (o *CreateCallResponse) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *CreateCallResponse) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *CreateCallResponse) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetPassword() string {
	if o == nil || utils.IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateCallResponse) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given utils.NullableString and assigns it to the Password field.
func (o *CreateCallResponse) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CreateCallResponse) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CreateCallResponse) UnsetPassword() {
	o.Password.Unset()
}

// GetFallbackUsername returns the FallbackUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetFallbackUsername() string {
	if o == nil || utils.IsNil(o.FallbackUsername.Get()) {
		var ret string
		return ret
	}
	return *o.FallbackUsername.Get()
}

// GetFallbackUsernameOk returns a tuple with the FallbackUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetFallbackUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackUsername.Get(), o.FallbackUsername.IsSet()
}

// HasFallbackUsername returns a boolean if a field has been set.
func (o *CreateCallResponse) HasFallbackUsername() bool {
	if o != nil && o.FallbackUsername.IsSet() {
		return true
	}

	return false
}

// SetFallbackUsername gets a reference to the given utils.NullableString and assigns it to the FallbackUsername field.
func (o *CreateCallResponse) SetFallbackUsername(v string) {
	o.FallbackUsername.Set(&v)
}
// SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil
func (o *CreateCallResponse) SetFallbackUsernameNil() {
	o.FallbackUsername.Set(nil)
}

// UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
func (o *CreateCallResponse) UnsetFallbackUsername() {
	o.FallbackUsername.Unset()
}

// GetFallbackPassword returns the FallbackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetFallbackPassword() string {
	if o == nil || utils.IsNil(o.FallbackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.FallbackPassword.Get()
}

// GetFallbackPasswordOk returns a tuple with the FallbackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetFallbackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackPassword.Get(), o.FallbackPassword.IsSet()
}

// HasFallbackPassword returns a boolean if a field has been set.
func (o *CreateCallResponse) HasFallbackPassword() bool {
	if o != nil && o.FallbackPassword.IsSet() {
		return true
	}

	return false
}

// SetFallbackPassword gets a reference to the given utils.NullableString and assigns it to the FallbackPassword field.
func (o *CreateCallResponse) SetFallbackPassword(v string) {
	o.FallbackPassword.Set(&v)
}
// SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil
func (o *CreateCallResponse) SetFallbackPasswordNil() {
	o.FallbackPassword.Set(nil)
}

// UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil
func (o *CreateCallResponse) UnsetFallbackPassword() {
	o.FallbackPassword.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCallResponse) GetPriority() int32 {
	if o == nil || utils.IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCallResponse) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateCallResponse) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *CreateCallResponse) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *CreateCallResponse) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *CreateCallResponse) UnsetPriority() {
	o.Priority.Unset()
}

func (o CreateCallResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCallResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicationId"] = o.ApplicationId
	toSerialize["accountId"] = o.AccountId
	toSerialize["callId"] = o.CallId
	toSerialize["to"] = o.To
	toSerialize["from"] = o.From
	if o.EnqueuedTime.IsSet() {
		toSerialize["enqueuedTime"] = o.EnqueuedTime.Get()
	}
	toSerialize["callUrl"] = o.CallUrl
	if !utils.IsNil(o.CallTimeout) {
		toSerialize["callTimeout"] = o.CallTimeout
	}
	if !utils.IsNil(o.CallbackTimeout) {
		toSerialize["callbackTimeout"] = o.CallbackTimeout
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	toSerialize["answerMethod"] = o.AnswerMethod.Get()
	toSerialize["answerUrl"] = o.AnswerUrl
	if o.AnswerFallbackMethod.IsSet() {
		toSerialize["answerFallbackMethod"] = o.AnswerFallbackMethod.Get()
	}
	if o.AnswerFallbackUrl.IsSet() {
		toSerialize["answerFallbackUrl"] = o.AnswerFallbackUrl.Get()
	}
	toSerialize["disconnectMethod"] = o.DisconnectMethod.Get()
	if o.DisconnectUrl.IsSet() {
		toSerialize["disconnectUrl"] = o.DisconnectUrl.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.FallbackUsername.IsSet() {
		toSerialize["fallbackUsername"] = o.FallbackUsername.Get()
	}
	if o.FallbackPassword.IsSet() {
		toSerialize["fallbackPassword"] = o.FallbackPassword.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	return toSerialize, nil
}

func (o *CreateCallResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicationId",
		"accountId",
		"callId",
		"to",
		"from",
		"callUrl",
		"answerMethod",
		"answerUrl",
		"disconnectMethod",
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

	varCreateCallResponse := _CreateCallResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCallResponse)

	if err != nil {
		return err
	}

	*o = CreateCallResponse(varCreateCallResponse)

	return err
}

type NullableCreateCallResponse struct {
	value *CreateCallResponse
	isSet bool
}

func (v NullableCreateCallResponse) Get() *CreateCallResponse {
	return v.value
}

func (v *NullableCreateCallResponse) Set(val *CreateCallResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCallResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCallResponse(val *CreateCallResponse) *NullableCreateCallResponse {
	return &NullableCreateCallResponse{value: val, isSet: true}
}

func (v NullableCreateCallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


