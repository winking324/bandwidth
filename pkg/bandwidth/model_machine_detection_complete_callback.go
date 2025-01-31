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
	"time"
)

// checks if the MachineDetectionCompleteCallback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineDetectionCompleteCallback{}

// MachineDetectionCompleteCallback This event is sent to the url informed when requesting a machine detection operation. It contains the machine detection operation result, which can be: human, answering-machine, silence, timeout, error. This event is not sent when sync answering machine detection mode is chosen.
type MachineDetectionCompleteCallback struct {
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
	Direction *CallDirectionEnum `json:"direction,omitempty"`
	// The call id associated with the event.
	CallId *string `json:"callId,omitempty"`
	// The URL of the call associated with the event.
	CallUrl *string `json:"callUrl,omitempty"`
	// (optional) If call queueing is enabled and this is an outbound call, time the call was queued, in ISO 8601 format.
	EnqueuedTime NullableTime `json:"enqueuedTime,omitempty"`
	// Time the call was started, in ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Time the call was answered, in ISO 8601 format.
	AnswerTime NullableTime `json:"answerTime,omitempty"`
	// (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present.
	Tag NullableString `json:"tag,omitempty"`
	MachineDetectionResult NullableMachineDetectionResult `json:"machineDetectionResult,omitempty"`
}

// NewMachineDetectionCompleteCallback instantiates a new MachineDetectionCompleteCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineDetectionCompleteCallback() *MachineDetectionCompleteCallback {
	this := MachineDetectionCompleteCallback{}
	return &this
}

// NewMachineDetectionCompleteCallbackWithDefaults instantiates a new MachineDetectionCompleteCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineDetectionCompleteCallbackWithDefaults() *MachineDetectionCompleteCallback {
	this := MachineDetectionCompleteCallback{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MachineDetectionCompleteCallback) SetEventType(v string) {
	o.EventType = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetEventTime() time.Time {
	if o == nil || IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *MachineDetectionCompleteCallback) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *MachineDetectionCompleteCallback) SetAccountId(v string) {
	o.AccountId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *MachineDetectionCompleteCallback) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *MachineDetectionCompleteCallback) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *MachineDetectionCompleteCallback) SetTo(v string) {
	o.To = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetDirection() CallDirectionEnum {
	if o == nil || IsNil(o.Direction) {
		var ret CallDirectionEnum
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetDirectionOk() (*CallDirectionEnum, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given CallDirectionEnum and assigns it to the Direction field.
func (o *MachineDetectionCompleteCallback) SetDirection(v CallDirectionEnum) {
	o.Direction = &v
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *MachineDetectionCompleteCallback) SetCallId(v string) {
	o.CallId = &v
}

// GetCallUrl returns the CallUrl field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetCallUrl() string {
	if o == nil || IsNil(o.CallUrl) {
		var ret string
		return ret
	}
	return *o.CallUrl
}

// GetCallUrlOk returns a tuple with the CallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetCallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallUrl) {
		return nil, false
	}
	return o.CallUrl, true
}

// HasCallUrl returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasCallUrl() bool {
	if o != nil && !IsNil(o.CallUrl) {
		return true
	}

	return false
}

// SetCallUrl gets a reference to the given string and assigns it to the CallUrl field.
func (o *MachineDetectionCompleteCallback) SetCallUrl(v string) {
	o.CallUrl = &v
}

// GetEnqueuedTime returns the EnqueuedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineDetectionCompleteCallback) GetEnqueuedTime() time.Time {
	if o == nil || IsNil(o.EnqueuedTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EnqueuedTime.Get()
}

// GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineDetectionCompleteCallback) GetEnqueuedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnqueuedTime.Get(), o.EnqueuedTime.IsSet()
}

// HasEnqueuedTime returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasEnqueuedTime() bool {
	if o != nil && o.EnqueuedTime.IsSet() {
		return true
	}

	return false
}

// SetEnqueuedTime gets a reference to the given NullableTime and assigns it to the EnqueuedTime field.
func (o *MachineDetectionCompleteCallback) SetEnqueuedTime(v time.Time) {
	o.EnqueuedTime.Set(&v)
}
// SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil
func (o *MachineDetectionCompleteCallback) SetEnqueuedTimeNil() {
	o.EnqueuedTime.Set(nil)
}

// UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
func (o *MachineDetectionCompleteCallback) UnsetEnqueuedTime() {
	o.EnqueuedTime.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MachineDetectionCompleteCallback) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionCompleteCallback) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *MachineDetectionCompleteCallback) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetAnswerTime returns the AnswerTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineDetectionCompleteCallback) GetAnswerTime() time.Time {
	if o == nil || IsNil(o.AnswerTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AnswerTime.Get()
}

// GetAnswerTimeOk returns a tuple with the AnswerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineDetectionCompleteCallback) GetAnswerTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerTime.Get(), o.AnswerTime.IsSet()
}

// HasAnswerTime returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasAnswerTime() bool {
	if o != nil && o.AnswerTime.IsSet() {
		return true
	}

	return false
}

// SetAnswerTime gets a reference to the given NullableTime and assigns it to the AnswerTime field.
func (o *MachineDetectionCompleteCallback) SetAnswerTime(v time.Time) {
	o.AnswerTime.Set(&v)
}
// SetAnswerTimeNil sets the value for AnswerTime to be an explicit nil
func (o *MachineDetectionCompleteCallback) SetAnswerTimeNil() {
	o.AnswerTime.Set(nil)
}

// UnsetAnswerTime ensures that no value is present for AnswerTime, not even an explicit nil
func (o *MachineDetectionCompleteCallback) UnsetAnswerTime() {
	o.AnswerTime.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineDetectionCompleteCallback) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineDetectionCompleteCallback) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *MachineDetectionCompleteCallback) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *MachineDetectionCompleteCallback) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *MachineDetectionCompleteCallback) UnsetTag() {
	o.Tag.Unset()
}

// GetMachineDetectionResult returns the MachineDetectionResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineDetectionCompleteCallback) GetMachineDetectionResult() MachineDetectionResult {
	if o == nil || IsNil(o.MachineDetectionResult.Get()) {
		var ret MachineDetectionResult
		return ret
	}
	return *o.MachineDetectionResult.Get()
}

// GetMachineDetectionResultOk returns a tuple with the MachineDetectionResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineDetectionCompleteCallback) GetMachineDetectionResultOk() (*MachineDetectionResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineDetectionResult.Get(), o.MachineDetectionResult.IsSet()
}

// HasMachineDetectionResult returns a boolean if a field has been set.
func (o *MachineDetectionCompleteCallback) HasMachineDetectionResult() bool {
	if o != nil && o.MachineDetectionResult.IsSet() {
		return true
	}

	return false
}

// SetMachineDetectionResult gets a reference to the given NullableMachineDetectionResult and assigns it to the MachineDetectionResult field.
func (o *MachineDetectionCompleteCallback) SetMachineDetectionResult(v MachineDetectionResult) {
	o.MachineDetectionResult.Set(&v)
}
// SetMachineDetectionResultNil sets the value for MachineDetectionResult to be an explicit nil
func (o *MachineDetectionCompleteCallback) SetMachineDetectionResultNil() {
	o.MachineDetectionResult.Set(nil)
}

// UnsetMachineDetectionResult ensures that no value is present for MachineDetectionResult, not even an explicit nil
func (o *MachineDetectionCompleteCallback) UnsetMachineDetectionResult() {
	o.MachineDetectionResult.Unset()
}

func (o MachineDetectionCompleteCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineDetectionCompleteCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.CallId) {
		toSerialize["callId"] = o.CallId
	}
	if !IsNil(o.CallUrl) {
		toSerialize["callUrl"] = o.CallUrl
	}
	if o.EnqueuedTime.IsSet() {
		toSerialize["enqueuedTime"] = o.EnqueuedTime.Get()
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if o.AnswerTime.IsSet() {
		toSerialize["answerTime"] = o.AnswerTime.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if o.MachineDetectionResult.IsSet() {
		toSerialize["machineDetectionResult"] = o.MachineDetectionResult.Get()
	}
	return toSerialize, nil
}

type NullableMachineDetectionCompleteCallback struct {
	value *MachineDetectionCompleteCallback
	isSet bool
}

func (v NullableMachineDetectionCompleteCallback) Get() *MachineDetectionCompleteCallback {
	return v.value
}

func (v *NullableMachineDetectionCompleteCallback) Set(val *MachineDetectionCompleteCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineDetectionCompleteCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineDetectionCompleteCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineDetectionCompleteCallback(val *MachineDetectionCompleteCallback) *NullableMachineDetectionCompleteCallback {
	return &NullableMachineDetectionCompleteCallback{value: val, isSet: true}
}

func (v NullableMachineDetectionCompleteCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineDetectionCompleteCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


