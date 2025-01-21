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
)

// DtmfCallback The DTMF event is sent for every digit detected after a <StartGather> verb is executed. You may not respond to this event with BXML.
type DtmfCallback struct {
	// The event type, value can be one of the following: answer, bridgeComplete, bridgeTargetComplete, conferenceCreated, conferenceRedirect, conferenceMemberJoin, conferenceMemberExit, conferenceCompleted, conferenceRecordingAvailable, disconnect, dtmf, gather, initiate, machineDetectionComplete, recordingComplete, recordingAvailable, redirect, transcriptionAvailable, transferAnswer, transferComplete, transferDisconnect.
	EventType *string `json:"eventType,omitempty"`
	// The approximate UTC date and time when the event was generated by the Bandwidth server, in ISO 8601 format. This may not be exactly the time of event execution.
	EventTime *time.Time `json:"eventTime,omitempty"`
	// The user account associated with the call.
	AccountId *string `json:"accountId,omitempty"`
	// The id of the application associated with the call.
	ApplicationId *string `json:"applicationId,omitempty"`
	// The provided identifier of the caller. Must be a phone number in E.164 format (e.g. +15555555555).
	From *string `json:"from,omitempty"`
	// The phone number that received the call, in E.164 format (e.g. +15555555555).
	To *string `json:"to,omitempty"`
	// The call id associated with the event.
	CallId *string `json:"callId,omitempty"`
	Direction *CallDirectionEnum `json:"direction,omitempty"`
	// The digit collected in the call.
	Digit *string `json:"digit,omitempty"`
	// The URL of the call associated with the event.
	CallUrl *string `json:"callUrl,omitempty"`
	// (optional) If call queueing is enabled and this is an outbound call, time the call was queued, in ISO 8601 format.
	EnqueuedTime utils.NullableTime `json:"enqueuedTime,omitempty"`
	// Time the call was started, in ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Time the call was answered, in ISO 8601 format.
	AnswerTime utils.NullableTime `json:"answerTime,omitempty"`
	// (optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, this field will not be present.
	ParentCallId *string `json:"parentCallId,omitempty"`
	// The phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferCallerId *string `json:"transferCallerId,omitempty"`
	// The phone number used as the to field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferTo *string `json:"transferTo,omitempty"`
	// (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present.
	Tag utils.NullableString `json:"tag,omitempty"`
}

// NewDtmfCallback instantiates a new DtmfCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtmfCallback() *DtmfCallback {
	this := DtmfCallback{}
	return &this
}

// NewDtmfCallbackWithDefaults instantiates a new DtmfCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtmfCallbackWithDefaults() *DtmfCallback {
	this := DtmfCallback{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DtmfCallback) GetEventType() string {
	if o == nil || utils.IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetEventTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DtmfCallback) HasEventType() bool {
	if o != nil && !utils.IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *DtmfCallback) SetEventType(v string) {
	o.EventType = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *DtmfCallback) GetEventTime() time.Time {
	if o == nil || utils.IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *DtmfCallback) HasEventTime() bool {
	if o != nil && !utils.IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *DtmfCallback) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *DtmfCallback) GetAccountId() string {
	if o == nil || utils.IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetAccountIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *DtmfCallback) HasAccountId() bool {
	if o != nil && !utils.IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *DtmfCallback) SetAccountId(v string) {
	o.AccountId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *DtmfCallback) GetApplicationId() string {
	if o == nil || utils.IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetApplicationIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *DtmfCallback) HasApplicationId() bool {
	if o != nil && !utils.IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *DtmfCallback) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *DtmfCallback) GetFrom() string {
	if o == nil || utils.IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetFromOk() (*string, bool) {
	if o == nil || utils.IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *DtmfCallback) HasFrom() bool {
	if o != nil && !utils.IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *DtmfCallback) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *DtmfCallback) GetTo() string {
	if o == nil || utils.IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetToOk() (*string, bool) {
	if o == nil || utils.IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *DtmfCallback) HasTo() bool {
	if o != nil && !utils.IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *DtmfCallback) SetTo(v string) {
	o.To = &v
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *DtmfCallback) GetCallId() string {
	if o == nil || utils.IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetCallIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *DtmfCallback) HasCallId() bool {
	if o != nil && !utils.IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *DtmfCallback) SetCallId(v string) {
	o.CallId = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *DtmfCallback) GetDirection() CallDirectionEnum {
	if o == nil || utils.IsNil(o.Direction) {
		var ret CallDirectionEnum
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetDirectionOk() (*CallDirectionEnum, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *DtmfCallback) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given CallDirectionEnum and assigns it to the Direction field.
func (o *DtmfCallback) SetDirection(v CallDirectionEnum) {
	o.Direction = &v
}

// GetDigit returns the Digit field value if set, zero value otherwise.
func (o *DtmfCallback) GetDigit() string {
	if o == nil || utils.IsNil(o.Digit) {
		var ret string
		return ret
	}
	return *o.Digit
}

// GetDigitOk returns a tuple with the Digit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetDigitOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Digit) {
		return nil, false
	}
	return o.Digit, true
}

// HasDigit returns a boolean if a field has been set.
func (o *DtmfCallback) HasDigit() bool {
	if o != nil && !utils.IsNil(o.Digit) {
		return true
	}

	return false
}

// SetDigit gets a reference to the given string and assigns it to the Digit field.
func (o *DtmfCallback) SetDigit(v string) {
	o.Digit = &v
}

// GetCallUrl returns the CallUrl field value if set, zero value otherwise.
func (o *DtmfCallback) GetCallUrl() string {
	if o == nil || utils.IsNil(o.CallUrl) {
		var ret string
		return ret
	}
	return *o.CallUrl
}

// GetCallUrlOk returns a tuple with the CallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetCallUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CallUrl) {
		return nil, false
	}
	return o.CallUrl, true
}

// HasCallUrl returns a boolean if a field has been set.
func (o *DtmfCallback) HasCallUrl() bool {
	if o != nil && !utils.IsNil(o.CallUrl) {
		return true
	}

	return false
}

// SetCallUrl gets a reference to the given string and assigns it to the CallUrl field.
func (o *DtmfCallback) SetCallUrl(v string) {
	o.CallUrl = &v
}

// GetEnqueuedTime returns the EnqueuedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DtmfCallback) GetEnqueuedTime() time.Time {
	if o == nil || utils.IsNil(o.EnqueuedTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EnqueuedTime.Get()
}

// GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DtmfCallback) GetEnqueuedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnqueuedTime.Get(), o.EnqueuedTime.IsSet()
}

// HasEnqueuedTime returns a boolean if a field has been set.
func (o *DtmfCallback) HasEnqueuedTime() bool {
	if o != nil && o.EnqueuedTime.IsSet() {
		return true
	}

	return false
}

// SetEnqueuedTime gets a reference to the given utils.NullableTime and assigns it to the EnqueuedTime field.
func (o *DtmfCallback) SetEnqueuedTime(v time.Time) {
	o.EnqueuedTime.Set(&v)
}
// SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil
func (o *DtmfCallback) SetEnqueuedTimeNil() {
	o.EnqueuedTime.Set(nil)
}

// UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
func (o *DtmfCallback) UnsetEnqueuedTime() {
	o.EnqueuedTime.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DtmfCallback) GetStartTime() time.Time {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DtmfCallback) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *DtmfCallback) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetAnswerTime returns the AnswerTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DtmfCallback) GetAnswerTime() time.Time {
	if o == nil || utils.IsNil(o.AnswerTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AnswerTime.Get()
}

// GetAnswerTimeOk returns a tuple with the AnswerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DtmfCallback) GetAnswerTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerTime.Get(), o.AnswerTime.IsSet()
}

// HasAnswerTime returns a boolean if a field has been set.
func (o *DtmfCallback) HasAnswerTime() bool {
	if o != nil && o.AnswerTime.IsSet() {
		return true
	}

	return false
}

// SetAnswerTime gets a reference to the given utils.NullableTime and assigns it to the AnswerTime field.
func (o *DtmfCallback) SetAnswerTime(v time.Time) {
	o.AnswerTime.Set(&v)
}
// SetAnswerTimeNil sets the value for AnswerTime to be an explicit nil
func (o *DtmfCallback) SetAnswerTimeNil() {
	o.AnswerTime.Set(nil)
}

// UnsetAnswerTime ensures that no value is present for AnswerTime, not even an explicit nil
func (o *DtmfCallback) UnsetAnswerTime() {
	o.AnswerTime.Unset()
}

// GetParentCallId returns the ParentCallId field value if set, zero value otherwise.
func (o *DtmfCallback) GetParentCallId() string {
	if o == nil || utils.IsNil(o.ParentCallId) {
		var ret string
		return ret
	}
	return *o.ParentCallId
}

// GetParentCallIdOk returns a tuple with the ParentCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetParentCallIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ParentCallId) {
		return nil, false
	}
	return o.ParentCallId, true
}

// HasParentCallId returns a boolean if a field has been set.
func (o *DtmfCallback) HasParentCallId() bool {
	if o != nil && !utils.IsNil(o.ParentCallId) {
		return true
	}

	return false
}

// SetParentCallId gets a reference to the given string and assigns it to the ParentCallId field.
func (o *DtmfCallback) SetParentCallId(v string) {
	o.ParentCallId = &v
}

// GetTransferCallerId returns the TransferCallerId field value if set, zero value otherwise.
func (o *DtmfCallback) GetTransferCallerId() string {
	if o == nil || utils.IsNil(o.TransferCallerId) {
		var ret string
		return ret
	}
	return *o.TransferCallerId
}

// GetTransferCallerIdOk returns a tuple with the TransferCallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetTransferCallerIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TransferCallerId) {
		return nil, false
	}
	return o.TransferCallerId, true
}

// HasTransferCallerId returns a boolean if a field has been set.
func (o *DtmfCallback) HasTransferCallerId() bool {
	if o != nil && !utils.IsNil(o.TransferCallerId) {
		return true
	}

	return false
}

// SetTransferCallerId gets a reference to the given string and assigns it to the TransferCallerId field.
func (o *DtmfCallback) SetTransferCallerId(v string) {
	o.TransferCallerId = &v
}

// GetTransferTo returns the TransferTo field value if set, zero value otherwise.
func (o *DtmfCallback) GetTransferTo() string {
	if o == nil || utils.IsNil(o.TransferTo) {
		var ret string
		return ret
	}
	return *o.TransferTo
}

// GetTransferToOk returns a tuple with the TransferTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtmfCallback) GetTransferToOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TransferTo) {
		return nil, false
	}
	return o.TransferTo, true
}

// HasTransferTo returns a boolean if a field has been set.
func (o *DtmfCallback) HasTransferTo() bool {
	if o != nil && !utils.IsNil(o.TransferTo) {
		return true
	}

	return false
}

// SetTransferTo gets a reference to the given string and assigns it to the TransferTo field.
func (o *DtmfCallback) SetTransferTo(v string) {
	o.TransferTo = &v
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DtmfCallback) GetTag() string {
	if o == nil || utils.IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DtmfCallback) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *DtmfCallback) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given utils.NullableString and assigns it to the Tag field.
func (o *DtmfCallback) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *DtmfCallback) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *DtmfCallback) UnsetTag() {
	o.Tag.Unset()
}

func (o DtmfCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtmfCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !utils.IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !utils.IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !utils.IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !utils.IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !utils.IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !utils.IsNil(o.CallId) {
		toSerialize["callId"] = o.CallId
	}
	if !utils.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !utils.IsNil(o.Digit) {
		toSerialize["digit"] = o.Digit
	}
	if !utils.IsNil(o.CallUrl) {
		toSerialize["callUrl"] = o.CallUrl
	}
	if o.EnqueuedTime.IsSet() {
		toSerialize["enqueuedTime"] = o.EnqueuedTime.Get()
	}
	if !utils.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if o.AnswerTime.IsSet() {
		toSerialize["answerTime"] = o.AnswerTime.Get()
	}
	if !utils.IsNil(o.ParentCallId) {
		toSerialize["parentCallId"] = o.ParentCallId
	}
	if !utils.IsNil(o.TransferCallerId) {
		toSerialize["transferCallerId"] = o.TransferCallerId
	}
	if !utils.IsNil(o.TransferTo) {
		toSerialize["transferTo"] = o.TransferTo
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	return toSerialize, nil
}

type NullableDtmfCallback struct {
	value *DtmfCallback
	isSet bool
}

func (v NullableDtmfCallback) Get() *DtmfCallback {
	return v.value
}

func (v *NullableDtmfCallback) Set(val *DtmfCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableDtmfCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableDtmfCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtmfCallback(val *DtmfCallback) *NullableDtmfCallback {
	return &NullableDtmfCallback{value: val, isSet: true}
}

func (v NullableDtmfCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtmfCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


