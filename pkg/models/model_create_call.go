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

// CreateCall struct for CreateCall
type CreateCall struct {
	// The destination to call (must be an E.164 formatted number (e.g. `+15555551212`) or a SIP URI (e.g. `sip:user@server.example`)).
	To string `json:"to"`
	// A Bandwidth phone number on your account the call should come from (must be in E.164 format, like `+15555551212`) even if `privacy` is set to true.
	From string `json:"from"`
	// Hide the calling number. The `displayName` field can be used to customize the displayed name.
	Privacy NullableBool `json:"privacy,omitempty"`
	// The caller display name to use when the call is created.  May not exceed 256 characters nor contain control characters such as new lines. If `privacy` is true, only the following values are valid: `Restricted`, `Anonymous`, `Private`, or `Unavailable`.
	DisplayName utils.NullableString `json:"displayName,omitempty"`
	// A comma-separated list of 'User-To-User' headers to be sent in the INVITE when calling a SIP URI. Each value must end with an 'encoding' parameter as described in <a href='https://tools.ietf.org/html/rfc7433'>RFC 7433</a>. Only 'jwt', 'base64' and 'hex' encodings are allowed. The entire value cannot exceed 350 characters, including parameters and separators.
	Uui utils.NullableString `json:"uui,omitempty"`
	// The id of the application associated with the `from` number.
	ApplicationId string `json:"applicationId"`
	// The full URL to send the <a href='/docs/voice/webhooks/answer'>Answer</a> event to when the called party answers. This endpoint should return the first <a href='/docs/voice/bxml'>BXML document</a> to be executed in the call.  Must use `https` if specifying `username` and `password`.
	AnswerUrl string `json:"answerUrl"`
	AnswerMethod NullableCallbackMethodEnum `json:"answerMethod,omitempty"`
	// Basic auth username.
	Username utils.NullableString `json:"username,omitempty"`
	// Basic auth password.
	Password utils.NullableString `json:"password,omitempty"`
	// A fallback url which, if provided, will be used to retry the `answer` webhook delivery in case `answerUrl` fails to respond  Must use `https` if specifying `fallbackUsername` and `fallbackPassword`.
	AnswerFallbackUrl utils.NullableString `json:"answerFallbackUrl,omitempty"`
	AnswerFallbackMethod NullableCallbackMethodEnum `json:"answerFallbackMethod,omitempty"`
	// Basic auth username.
	FallbackUsername utils.NullableString `json:"fallbackUsername,omitempty"`
	// Basic auth password.
	FallbackPassword utils.NullableString `json:"fallbackPassword,omitempty"`
	// The URL to send the <a href='/docs/voice/webhooks/disconnect'>Disconnect</a> event to when the call ends. This event does not expect a BXML response.
	DisconnectUrl utils.NullableString `json:"disconnectUrl,omitempty"`
	DisconnectMethod NullableCallbackMethodEnum `json:"disconnectMethod,omitempty"`
	// The timeout (in seconds) for the callee to answer the call after it starts ringing. If the call does not start ringing within 30s, the call will be cancelled regardless of this value.  Can be any numeric value (including decimals) between 1 and 300.
	CallTimeout NullableFloat64 `json:"callTimeout,omitempty"`
	// This is the timeout (in seconds) to use when delivering webhooks for the call. Can be any numeric value (including decimals) between 1 and 25.
	CallbackTimeout NullableFloat64 `json:"callbackTimeout,omitempty"`
	MachineDetection *MachineDetectionConfiguration `json:"machineDetection,omitempty"`
	// The priority of this call over other calls from your account. For example, if during a call your application needs to place a new call and bridge it with the current call, you might want to create the call with priority 1 so that it will be the next call picked off your queue, ahead of other less time sensitive calls. A lower value means higher priority, so a priority 1 call takes precedence over a priority 2 call.
	Priority NullableInt32 `json:"priority,omitempty"`
	// A custom string that will be sent with all webhooks for this call unless overwritten by a future <a href='/docs/voice/bxml/tag'>`<Tag>`</a> verb or `tag` attribute on another verb, or cleared.  May be cleared by setting `tag=\"\"`  Max length 256 characters.
	Tag utils.NullableString `json:"tag,omitempty"`
}

type _CreateCall CreateCall

// NewCreateCall instantiates a new CreateCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCall(to string, from string, applicationId string, answerUrl string) *CreateCall {
	this := CreateCall{}
	this.To = to
	this.From = from
	this.ApplicationId = applicationId
	this.AnswerUrl = answerUrl
	var answerMethod CallbackMethodEnum = POST
	this.AnswerMethod = *NewNullableCallbackMethodEnum(&answerMethod)
	var answerFallbackMethod CallbackMethodEnum = POST
	this.AnswerFallbackMethod = *NewNullableCallbackMethodEnum(&answerFallbackMethod)
	var disconnectMethod CallbackMethodEnum = POST
	this.DisconnectMethod = *NewNullableCallbackMethodEnum(&disconnectMethod)
	var callTimeout float64 = 30
	this.CallTimeout = *NewNullableFloat64(&callTimeout)
	var callbackTimeout float64 = 15
	this.CallbackTimeout = *NewNullableFloat64(&callbackTimeout)
	var priority int32 = 5
	this.Priority = *NewNullableInt32(&priority)
	return &this
}

// NewCreateCallWithDefaults instantiates a new CreateCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCallWithDefaults() *CreateCall {
	this := CreateCall{}
	var answerMethod CallbackMethodEnum = POST
	this.AnswerMethod = *NewNullableCallbackMethodEnum(&answerMethod)
	var answerFallbackMethod CallbackMethodEnum = POST
	this.AnswerFallbackMethod = *NewNullableCallbackMethodEnum(&answerFallbackMethod)
	var disconnectMethod CallbackMethodEnum = POST
	this.DisconnectMethod = *NewNullableCallbackMethodEnum(&disconnectMethod)
	var callTimeout float64 = 30
	this.CallTimeout = *NewNullableFloat64(&callTimeout)
	var callbackTimeout float64 = 15
	this.CallbackTimeout = *NewNullableFloat64(&callbackTimeout)
	var priority int32 = 5
	this.Priority = *NewNullableInt32(&priority)
	return &this
}

// GetTo returns the To field value
func (o *CreateCall) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *CreateCall) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *CreateCall) SetTo(v string) {
	o.To = v
}

// GetFrom returns the From field value
func (o *CreateCall) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateCall) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CreateCall) SetFrom(v string) {
	o.From = v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetPrivacy() bool {
	if o == nil || utils.IsNil(o.Privacy.Get()) {
		var ret bool
		return ret
	}
	return *o.Privacy.Get()
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetPrivacyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Privacy.Get(), o.Privacy.IsSet()
}

// HasPrivacy returns a boolean if a field has been set.
func (o *CreateCall) HasPrivacy() bool {
	if o != nil && o.Privacy.IsSet() {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given NullableBool and assigns it to the Privacy field.
func (o *CreateCall) SetPrivacy(v bool) {
	o.Privacy.Set(&v)
}
// SetPrivacyNil sets the value for Privacy to be an explicit nil
func (o *CreateCall) SetPrivacyNil() {
	o.Privacy.Set(nil)
}

// UnsetPrivacy ensures that no value is present for Privacy, not even an explicit nil
func (o *CreateCall) UnsetPrivacy() {
	o.Privacy.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetDisplayName() string {
	if o == nil || utils.IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateCall) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given utils.NullableString and assigns it to the DisplayName field.
func (o *CreateCall) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CreateCall) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CreateCall) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetUui returns the Uui field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetUui() string {
	if o == nil || utils.IsNil(o.Uui.Get()) {
		var ret string
		return ret
	}
	return *o.Uui.Get()
}

// GetUuiOk returns a tuple with the Uui field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetUuiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uui.Get(), o.Uui.IsSet()
}

// HasUui returns a boolean if a field has been set.
func (o *CreateCall) HasUui() bool {
	if o != nil && o.Uui.IsSet() {
		return true
	}

	return false
}

// SetUui gets a reference to the given utils.NullableString and assigns it to the Uui field.
func (o *CreateCall) SetUui(v string) {
	o.Uui.Set(&v)
}
// SetUuiNil sets the value for Uui to be an explicit nil
func (o *CreateCall) SetUuiNil() {
	o.Uui.Set(nil)
}

// UnsetUui ensures that no value is present for Uui, not even an explicit nil
func (o *CreateCall) UnsetUui() {
	o.Uui.Unset()
}

// GetApplicationId returns the ApplicationId field value
func (o *CreateCall) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CreateCall) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CreateCall) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetAnswerUrl returns the AnswerUrl field value
func (o *CreateCall) GetAnswerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnswerUrl
}

// GetAnswerUrlOk returns a tuple with the AnswerUrl field value
// and a boolean to check if the value has been set.
func (o *CreateCall) GetAnswerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerUrl, true
}

// SetAnswerUrl sets field value
func (o *CreateCall) SetAnswerUrl(v string) {
	o.AnswerUrl = v
}

// GetAnswerMethod returns the AnswerMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetAnswerMethod() CallbackMethodEnum {
	if o == nil || utils.IsNil(o.AnswerMethod.Get()) {
		var ret CallbackMethodEnum
		return ret
	}
	return *o.AnswerMethod.Get()
}

// GetAnswerMethodOk returns a tuple with the AnswerMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetAnswerMethodOk() (*CallbackMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerMethod.Get(), o.AnswerMethod.IsSet()
}

// HasAnswerMethod returns a boolean if a field has been set.
func (o *CreateCall) HasAnswerMethod() bool {
	if o != nil && o.AnswerMethod.IsSet() {
		return true
	}

	return false
}

// SetAnswerMethod gets a reference to the given NullableCallbackMethodEnum and assigns it to the AnswerMethod field.
func (o *CreateCall) SetAnswerMethod(v CallbackMethodEnum) {
	o.AnswerMethod.Set(&v)
}
// SetAnswerMethodNil sets the value for AnswerMethod to be an explicit nil
func (o *CreateCall) SetAnswerMethodNil() {
	o.AnswerMethod.Set(nil)
}

// UnsetAnswerMethod ensures that no value is present for AnswerMethod, not even an explicit nil
func (o *CreateCall) UnsetAnswerMethod() {
	o.AnswerMethod.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetUsername() string {
	if o == nil || utils.IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateCall) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given utils.NullableString and assigns it to the Username field.
func (o *CreateCall) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *CreateCall) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *CreateCall) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetPassword() string {
	if o == nil || utils.IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateCall) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given utils.NullableString and assigns it to the Password field.
func (o *CreateCall) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CreateCall) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CreateCall) UnsetPassword() {
	o.Password.Unset()
}

// GetAnswerFallbackUrl returns the AnswerFallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetAnswerFallbackUrl() string {
	if o == nil || utils.IsNil(o.AnswerFallbackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AnswerFallbackUrl.Get()
}

// GetAnswerFallbackUrlOk returns a tuple with the AnswerFallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetAnswerFallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerFallbackUrl.Get(), o.AnswerFallbackUrl.IsSet()
}

// HasAnswerFallbackUrl returns a boolean if a field has been set.
func (o *CreateCall) HasAnswerFallbackUrl() bool {
	if o != nil && o.AnswerFallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetAnswerFallbackUrl gets a reference to the given utils.NullableString and assigns it to the AnswerFallbackUrl field.
func (o *CreateCall) SetAnswerFallbackUrl(v string) {
	o.AnswerFallbackUrl.Set(&v)
}
// SetAnswerFallbackUrlNil sets the value for AnswerFallbackUrl to be an explicit nil
func (o *CreateCall) SetAnswerFallbackUrlNil() {
	o.AnswerFallbackUrl.Set(nil)
}

// UnsetAnswerFallbackUrl ensures that no value is present for AnswerFallbackUrl, not even an explicit nil
func (o *CreateCall) UnsetAnswerFallbackUrl() {
	o.AnswerFallbackUrl.Unset()
}

// GetAnswerFallbackMethod returns the AnswerFallbackMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetAnswerFallbackMethod() CallbackMethodEnum {
	if o == nil || utils.IsNil(o.AnswerFallbackMethod.Get()) {
		var ret CallbackMethodEnum
		return ret
	}
	return *o.AnswerFallbackMethod.Get()
}

// GetAnswerFallbackMethodOk returns a tuple with the AnswerFallbackMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetAnswerFallbackMethodOk() (*CallbackMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerFallbackMethod.Get(), o.AnswerFallbackMethod.IsSet()
}

// HasAnswerFallbackMethod returns a boolean if a field has been set.
func (o *CreateCall) HasAnswerFallbackMethod() bool {
	if o != nil && o.AnswerFallbackMethod.IsSet() {
		return true
	}

	return false
}

// SetAnswerFallbackMethod gets a reference to the given NullableCallbackMethodEnum and assigns it to the AnswerFallbackMethod field.
func (o *CreateCall) SetAnswerFallbackMethod(v CallbackMethodEnum) {
	o.AnswerFallbackMethod.Set(&v)
}
// SetAnswerFallbackMethodNil sets the value for AnswerFallbackMethod to be an explicit nil
func (o *CreateCall) SetAnswerFallbackMethodNil() {
	o.AnswerFallbackMethod.Set(nil)
}

// UnsetAnswerFallbackMethod ensures that no value is present for AnswerFallbackMethod, not even an explicit nil
func (o *CreateCall) UnsetAnswerFallbackMethod() {
	o.AnswerFallbackMethod.Unset()
}

// GetFallbackUsername returns the FallbackUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetFallbackUsername() string {
	if o == nil || utils.IsNil(o.FallbackUsername.Get()) {
		var ret string
		return ret
	}
	return *o.FallbackUsername.Get()
}

// GetFallbackUsernameOk returns a tuple with the FallbackUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetFallbackUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackUsername.Get(), o.FallbackUsername.IsSet()
}

// HasFallbackUsername returns a boolean if a field has been set.
func (o *CreateCall) HasFallbackUsername() bool {
	if o != nil && o.FallbackUsername.IsSet() {
		return true
	}

	return false
}

// SetFallbackUsername gets a reference to the given utils.NullableString and assigns it to the FallbackUsername field.
func (o *CreateCall) SetFallbackUsername(v string) {
	o.FallbackUsername.Set(&v)
}
// SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil
func (o *CreateCall) SetFallbackUsernameNil() {
	o.FallbackUsername.Set(nil)
}

// UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
func (o *CreateCall) UnsetFallbackUsername() {
	o.FallbackUsername.Unset()
}

// GetFallbackPassword returns the FallbackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetFallbackPassword() string {
	if o == nil || utils.IsNil(o.FallbackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.FallbackPassword.Get()
}

// GetFallbackPasswordOk returns a tuple with the FallbackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetFallbackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FallbackPassword.Get(), o.FallbackPassword.IsSet()
}

// HasFallbackPassword returns a boolean if a field has been set.
func (o *CreateCall) HasFallbackPassword() bool {
	if o != nil && o.FallbackPassword.IsSet() {
		return true
	}

	return false
}

// SetFallbackPassword gets a reference to the given utils.NullableString and assigns it to the FallbackPassword field.
func (o *CreateCall) SetFallbackPassword(v string) {
	o.FallbackPassword.Set(&v)
}
// SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil
func (o *CreateCall) SetFallbackPasswordNil() {
	o.FallbackPassword.Set(nil)
}

// UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil
func (o *CreateCall) UnsetFallbackPassword() {
	o.FallbackPassword.Unset()
}

// GetDisconnectUrl returns the DisconnectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetDisconnectUrl() string {
	if o == nil || utils.IsNil(o.DisconnectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DisconnectUrl.Get()
}

// GetDisconnectUrlOk returns a tuple with the DisconnectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetDisconnectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisconnectUrl.Get(), o.DisconnectUrl.IsSet()
}

// HasDisconnectUrl returns a boolean if a field has been set.
func (o *CreateCall) HasDisconnectUrl() bool {
	if o != nil && o.DisconnectUrl.IsSet() {
		return true
	}

	return false
}

// SetDisconnectUrl gets a reference to the given utils.NullableString and assigns it to the DisconnectUrl field.
func (o *CreateCall) SetDisconnectUrl(v string) {
	o.DisconnectUrl.Set(&v)
}
// SetDisconnectUrlNil sets the value for DisconnectUrl to be an explicit nil
func (o *CreateCall) SetDisconnectUrlNil() {
	o.DisconnectUrl.Set(nil)
}

// UnsetDisconnectUrl ensures that no value is present for DisconnectUrl, not even an explicit nil
func (o *CreateCall) UnsetDisconnectUrl() {
	o.DisconnectUrl.Unset()
}

// GetDisconnectMethod returns the DisconnectMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetDisconnectMethod() CallbackMethodEnum {
	if o == nil || utils.IsNil(o.DisconnectMethod.Get()) {
		var ret CallbackMethodEnum
		return ret
	}
	return *o.DisconnectMethod.Get()
}

// GetDisconnectMethodOk returns a tuple with the DisconnectMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetDisconnectMethodOk() (*CallbackMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisconnectMethod.Get(), o.DisconnectMethod.IsSet()
}

// HasDisconnectMethod returns a boolean if a field has been set.
func (o *CreateCall) HasDisconnectMethod() bool {
	if o != nil && o.DisconnectMethod.IsSet() {
		return true
	}

	return false
}

// SetDisconnectMethod gets a reference to the given NullableCallbackMethodEnum and assigns it to the DisconnectMethod field.
func (o *CreateCall) SetDisconnectMethod(v CallbackMethodEnum) {
	o.DisconnectMethod.Set(&v)
}
// SetDisconnectMethodNil sets the value for DisconnectMethod to be an explicit nil
func (o *CreateCall) SetDisconnectMethodNil() {
	o.DisconnectMethod.Set(nil)
}

// UnsetDisconnectMethod ensures that no value is present for DisconnectMethod, not even an explicit nil
func (o *CreateCall) UnsetDisconnectMethod() {
	o.DisconnectMethod.Unset()
}

// GetCallTimeout returns the CallTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetCallTimeout() float64 {
	if o == nil || utils.IsNil(o.CallTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.CallTimeout.Get()
}

// GetCallTimeoutOk returns a tuple with the CallTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetCallTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallTimeout.Get(), o.CallTimeout.IsSet()
}

// HasCallTimeout returns a boolean if a field has been set.
func (o *CreateCall) HasCallTimeout() bool {
	if o != nil && o.CallTimeout.IsSet() {
		return true
	}

	return false
}

// SetCallTimeout gets a reference to the given NullableFloat64 and assigns it to the CallTimeout field.
func (o *CreateCall) SetCallTimeout(v float64) {
	o.CallTimeout.Set(&v)
}
// SetCallTimeoutNil sets the value for CallTimeout to be an explicit nil
func (o *CreateCall) SetCallTimeoutNil() {
	o.CallTimeout.Set(nil)
}

// UnsetCallTimeout ensures that no value is present for CallTimeout, not even an explicit nil
func (o *CreateCall) UnsetCallTimeout() {
	o.CallTimeout.Unset()
}

// GetCallbackTimeout returns the CallbackTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetCallbackTimeout() float64 {
	if o == nil || utils.IsNil(o.CallbackTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.CallbackTimeout.Get()
}

// GetCallbackTimeoutOk returns a tuple with the CallbackTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetCallbackTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallbackTimeout.Get(), o.CallbackTimeout.IsSet()
}

// HasCallbackTimeout returns a boolean if a field has been set.
func (o *CreateCall) HasCallbackTimeout() bool {
	if o != nil && o.CallbackTimeout.IsSet() {
		return true
	}

	return false
}

// SetCallbackTimeout gets a reference to the given NullableFloat64 and assigns it to the CallbackTimeout field.
func (o *CreateCall) SetCallbackTimeout(v float64) {
	o.CallbackTimeout.Set(&v)
}
// SetCallbackTimeoutNil sets the value for CallbackTimeout to be an explicit nil
func (o *CreateCall) SetCallbackTimeoutNil() {
	o.CallbackTimeout.Set(nil)
}

// UnsetCallbackTimeout ensures that no value is present for CallbackTimeout, not even an explicit nil
func (o *CreateCall) UnsetCallbackTimeout() {
	o.CallbackTimeout.Unset()
}

// GetMachineDetection returns the MachineDetection field value if set, zero value otherwise.
func (o *CreateCall) GetMachineDetection() MachineDetectionConfiguration {
	if o == nil || utils.IsNil(o.MachineDetection) {
		var ret MachineDetectionConfiguration
		return ret
	}
	return *o.MachineDetection
}

// GetMachineDetectionOk returns a tuple with the MachineDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCall) GetMachineDetectionOk() (*MachineDetectionConfiguration, bool) {
	if o == nil || utils.IsNil(o.MachineDetection) {
		return nil, false
	}
	return o.MachineDetection, true
}

// HasMachineDetection returns a boolean if a field has been set.
func (o *CreateCall) HasMachineDetection() bool {
	if o != nil && !utils.IsNil(o.MachineDetection) {
		return true
	}

	return false
}

// SetMachineDetection gets a reference to the given MachineDetectionConfiguration and assigns it to the MachineDetection field.
func (o *CreateCall) SetMachineDetection(v MachineDetectionConfiguration) {
	o.MachineDetection = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetPriority() int32 {
	if o == nil || utils.IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateCall) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *CreateCall) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *CreateCall) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *CreateCall) UnsetPriority() {
	o.Priority.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCall) GetTag() string {
	if o == nil || utils.IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCall) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *CreateCall) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given utils.NullableString and assigns it to the Tag field.
func (o *CreateCall) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *CreateCall) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *CreateCall) UnsetTag() {
	o.Tag.Unset()
}

func (o CreateCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["to"] = o.To
	toSerialize["from"] = o.From
	if o.Privacy.IsSet() {
		toSerialize["privacy"] = o.Privacy.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Uui.IsSet() {
		toSerialize["uui"] = o.Uui.Get()
	}
	toSerialize["applicationId"] = o.ApplicationId
	toSerialize["answerUrl"] = o.AnswerUrl
	if o.AnswerMethod.IsSet() {
		toSerialize["answerMethod"] = o.AnswerMethod.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.AnswerFallbackUrl.IsSet() {
		toSerialize["answerFallbackUrl"] = o.AnswerFallbackUrl.Get()
	}
	if o.AnswerFallbackMethod.IsSet() {
		toSerialize["answerFallbackMethod"] = o.AnswerFallbackMethod.Get()
	}
	if o.FallbackUsername.IsSet() {
		toSerialize["fallbackUsername"] = o.FallbackUsername.Get()
	}
	if o.FallbackPassword.IsSet() {
		toSerialize["fallbackPassword"] = o.FallbackPassword.Get()
	}
	if o.DisconnectUrl.IsSet() {
		toSerialize["disconnectUrl"] = o.DisconnectUrl.Get()
	}
	if o.DisconnectMethod.IsSet() {
		toSerialize["disconnectMethod"] = o.DisconnectMethod.Get()
	}
	if o.CallTimeout.IsSet() {
		toSerialize["callTimeout"] = o.CallTimeout.Get()
	}
	if o.CallbackTimeout.IsSet() {
		toSerialize["callbackTimeout"] = o.CallbackTimeout.Get()
	}
	if !utils.IsNil(o.MachineDetection) {
		toSerialize["machineDetection"] = o.MachineDetection
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	return toSerialize, nil
}

func (o *CreateCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"to",
		"from",
		"applicationId",
		"answerUrl",
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

	varCreateCall := _CreateCall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCall)

	if err != nil {
		return err
	}

	*o = CreateCall(varCreateCall)

	return err
}

type NullableCreateCall struct {
	value *CreateCall
	isSet bool
}

func (v NullableCreateCall) Get() *CreateCall {
	return v.value
}

func (v *NullableCreateCall) Set(val *CreateCall) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCall) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCall(val *CreateCall) *NullableCreateCall {
	return &NullableCreateCall{value: val, isSet: true}
}

func (v NullableCreateCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


