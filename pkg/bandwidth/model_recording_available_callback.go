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

// checks if the RecordingAvailableCallback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingAvailableCallback{}

// RecordingAvailableCallback The Recording Available event is sent after a recording has been processed. It indicates that the recording is available for download.
type RecordingAvailableCallback struct {
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
	// (optional) If the event is related to the B leg of a <Transfer>, the call id of the original call leg that executed the <Transfer>. Otherwise, this field will not be present.
	ParentCallId *string `json:"parentCallId,omitempty"`
	// The unique ID of this recording
	RecordingId *string `json:"recordingId,omitempty"`
	// The URL that can be used to download the recording. Only present if the recording is finished and may be downloaded.
	MediaUrl NullableString `json:"mediaUrl,omitempty"`
	// (optional) If call queueing is enabled and this is an outbound call, time the call was queued, in ISO 8601 format.
	EnqueuedTime NullableTime `json:"enqueuedTime,omitempty"`
	// Time the call was started, in ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`
	// The time that the recording ended in ISO-8601 format
	EndTime *time.Time `json:"endTime,omitempty"`
	// The duration of the recording in ISO-8601 format
	Duration *string `json:"duration,omitempty"`
	FileFormat *FileFormatEnum `json:"fileFormat,omitempty"`
	// Always `1` for conference recordings; multi-channel recordings are not supported on conferences.
	Channels *int32 `json:"channels,omitempty"`
	// (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present.
	Tag NullableString `json:"tag,omitempty"`
	// The current status of the process. For recording, current possible values are 'processing', 'partial', 'complete', 'deleted', and 'error'. For transcriptions, current possible values are 'none', 'processing', 'available', 'error', 'timeout', 'file-size-too-big', and 'file-size-too-small'. Additional states may be added in the future, so your application must be tolerant of unknown values.
	Status *string `json:"status,omitempty"`
	// The phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferCallerId *string `json:"transferCallerId,omitempty"`
	// The phone number used as the to field of the B-leg call, in E.164 format (e.g. +15555555555).
	TransferTo *string `json:"transferTo,omitempty"`
}

// NewRecordingAvailableCallback instantiates a new RecordingAvailableCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingAvailableCallback() *RecordingAvailableCallback {
	this := RecordingAvailableCallback{}
	return &this
}

// NewRecordingAvailableCallbackWithDefaults instantiates a new RecordingAvailableCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingAvailableCallbackWithDefaults() *RecordingAvailableCallback {
	this := RecordingAvailableCallback{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *RecordingAvailableCallback) SetEventType(v string) {
	o.EventType = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetEventTime() time.Time {
	if o == nil || IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *RecordingAvailableCallback) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *RecordingAvailableCallback) SetAccountId(v string) {
	o.AccountId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *RecordingAvailableCallback) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *RecordingAvailableCallback) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *RecordingAvailableCallback) SetTo(v string) {
	o.To = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetDirection() CallDirectionEnum {
	if o == nil || IsNil(o.Direction) {
		var ret CallDirectionEnum
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetDirectionOk() (*CallDirectionEnum, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given CallDirectionEnum and assigns it to the Direction field.
func (o *RecordingAvailableCallback) SetDirection(v CallDirectionEnum) {
	o.Direction = &v
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *RecordingAvailableCallback) SetCallId(v string) {
	o.CallId = &v
}

// GetCallUrl returns the CallUrl field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetCallUrl() string {
	if o == nil || IsNil(o.CallUrl) {
		var ret string
		return ret
	}
	return *o.CallUrl
}

// GetCallUrlOk returns a tuple with the CallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetCallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallUrl) {
		return nil, false
	}
	return o.CallUrl, true
}

// HasCallUrl returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasCallUrl() bool {
	if o != nil && !IsNil(o.CallUrl) {
		return true
	}

	return false
}

// SetCallUrl gets a reference to the given string and assigns it to the CallUrl field.
func (o *RecordingAvailableCallback) SetCallUrl(v string) {
	o.CallUrl = &v
}

// GetParentCallId returns the ParentCallId field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetParentCallId() string {
	if o == nil || IsNil(o.ParentCallId) {
		var ret string
		return ret
	}
	return *o.ParentCallId
}

// GetParentCallIdOk returns a tuple with the ParentCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetParentCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCallId) {
		return nil, false
	}
	return o.ParentCallId, true
}

// HasParentCallId returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasParentCallId() bool {
	if o != nil && !IsNil(o.ParentCallId) {
		return true
	}

	return false
}

// SetParentCallId gets a reference to the given string and assigns it to the ParentCallId field.
func (o *RecordingAvailableCallback) SetParentCallId(v string) {
	o.ParentCallId = &v
}

// GetRecordingId returns the RecordingId field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetRecordingId() string {
	if o == nil || IsNil(o.RecordingId) {
		var ret string
		return ret
	}
	return *o.RecordingId
}

// GetRecordingIdOk returns a tuple with the RecordingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetRecordingIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordingId) {
		return nil, false
	}
	return o.RecordingId, true
}

// HasRecordingId returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasRecordingId() bool {
	if o != nil && !IsNil(o.RecordingId) {
		return true
	}

	return false
}

// SetRecordingId gets a reference to the given string and assigns it to the RecordingId field.
func (o *RecordingAvailableCallback) SetRecordingId(v string) {
	o.RecordingId = &v
}

// GetMediaUrl returns the MediaUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecordingAvailableCallback) GetMediaUrl() string {
	if o == nil || IsNil(o.MediaUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MediaUrl.Get()
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordingAvailableCallback) GetMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaUrl.Get(), o.MediaUrl.IsSet()
}

// HasMediaUrl returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasMediaUrl() bool {
	if o != nil && o.MediaUrl.IsSet() {
		return true
	}

	return false
}

// SetMediaUrl gets a reference to the given NullableString and assigns it to the MediaUrl field.
func (o *RecordingAvailableCallback) SetMediaUrl(v string) {
	o.MediaUrl.Set(&v)
}
// SetMediaUrlNil sets the value for MediaUrl to be an explicit nil
func (o *RecordingAvailableCallback) SetMediaUrlNil() {
	o.MediaUrl.Set(nil)
}

// UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
func (o *RecordingAvailableCallback) UnsetMediaUrl() {
	o.MediaUrl.Unset()
}

// GetEnqueuedTime returns the EnqueuedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecordingAvailableCallback) GetEnqueuedTime() time.Time {
	if o == nil || IsNil(o.EnqueuedTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EnqueuedTime.Get()
}

// GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordingAvailableCallback) GetEnqueuedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnqueuedTime.Get(), o.EnqueuedTime.IsSet()
}

// HasEnqueuedTime returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasEnqueuedTime() bool {
	if o != nil && o.EnqueuedTime.IsSet() {
		return true
	}

	return false
}

// SetEnqueuedTime gets a reference to the given NullableTime and assigns it to the EnqueuedTime field.
func (o *RecordingAvailableCallback) SetEnqueuedTime(v time.Time) {
	o.EnqueuedTime.Set(&v)
}
// SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil
func (o *RecordingAvailableCallback) SetEnqueuedTimeNil() {
	o.EnqueuedTime.Set(nil)
}

// UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
func (o *RecordingAvailableCallback) UnsetEnqueuedTime() {
	o.EnqueuedTime.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *RecordingAvailableCallback) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *RecordingAvailableCallback) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *RecordingAvailableCallback) SetDuration(v string) {
	o.Duration = &v
}

// GetFileFormat returns the FileFormat field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetFileFormat() FileFormatEnum {
	if o == nil || IsNil(o.FileFormat) {
		var ret FileFormatEnum
		return ret
	}
	return *o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetFileFormatOk() (*FileFormatEnum, bool) {
	if o == nil || IsNil(o.FileFormat) {
		return nil, false
	}
	return o.FileFormat, true
}

// HasFileFormat returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasFileFormat() bool {
	if o != nil && !IsNil(o.FileFormat) {
		return true
	}

	return false
}

// SetFileFormat gets a reference to the given FileFormatEnum and assigns it to the FileFormat field.
func (o *RecordingAvailableCallback) SetFileFormat(v FileFormatEnum) {
	o.FileFormat = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetChannels() int32 {
	if o == nil || IsNil(o.Channels) {
		var ret int32
		return ret
	}
	return *o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetChannelsOk() (*int32, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given int32 and assigns it to the Channels field.
func (o *RecordingAvailableCallback) SetChannels(v int32) {
	o.Channels = &v
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecordingAvailableCallback) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordingAvailableCallback) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *RecordingAvailableCallback) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *RecordingAvailableCallback) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *RecordingAvailableCallback) UnsetTag() {
	o.Tag.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RecordingAvailableCallback) SetStatus(v string) {
	o.Status = &v
}

// GetTransferCallerId returns the TransferCallerId field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetTransferCallerId() string {
	if o == nil || IsNil(o.TransferCallerId) {
		var ret string
		return ret
	}
	return *o.TransferCallerId
}

// GetTransferCallerIdOk returns a tuple with the TransferCallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetTransferCallerIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransferCallerId) {
		return nil, false
	}
	return o.TransferCallerId, true
}

// HasTransferCallerId returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasTransferCallerId() bool {
	if o != nil && !IsNil(o.TransferCallerId) {
		return true
	}

	return false
}

// SetTransferCallerId gets a reference to the given string and assigns it to the TransferCallerId field.
func (o *RecordingAvailableCallback) SetTransferCallerId(v string) {
	o.TransferCallerId = &v
}

// GetTransferTo returns the TransferTo field value if set, zero value otherwise.
func (o *RecordingAvailableCallback) GetTransferTo() string {
	if o == nil || IsNil(o.TransferTo) {
		var ret string
		return ret
	}
	return *o.TransferTo
}

// GetTransferToOk returns a tuple with the TransferTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingAvailableCallback) GetTransferToOk() (*string, bool) {
	if o == nil || IsNil(o.TransferTo) {
		return nil, false
	}
	return o.TransferTo, true
}

// HasTransferTo returns a boolean if a field has been set.
func (o *RecordingAvailableCallback) HasTransferTo() bool {
	if o != nil && !IsNil(o.TransferTo) {
		return true
	}

	return false
}

// SetTransferTo gets a reference to the given string and assigns it to the TransferTo field.
func (o *RecordingAvailableCallback) SetTransferTo(v string) {
	o.TransferTo = &v
}

func (o RecordingAvailableCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingAvailableCallback) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ParentCallId) {
		toSerialize["parentCallId"] = o.ParentCallId
	}
	if !IsNil(o.RecordingId) {
		toSerialize["recordingId"] = o.RecordingId
	}
	if o.MediaUrl.IsSet() {
		toSerialize["mediaUrl"] = o.MediaUrl.Get()
	}
	if o.EnqueuedTime.IsSet() {
		toSerialize["enqueuedTime"] = o.EnqueuedTime.Get()
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.FileFormat) {
		toSerialize["fileFormat"] = o.FileFormat
	}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransferCallerId) {
		toSerialize["transferCallerId"] = o.TransferCallerId
	}
	if !IsNil(o.TransferTo) {
		toSerialize["transferTo"] = o.TransferTo
	}
	return toSerialize, nil
}

type NullableRecordingAvailableCallback struct {
	value *RecordingAvailableCallback
	isSet bool
}

func (v NullableRecordingAvailableCallback) Get() *RecordingAvailableCallback {
	return v.value
}

func (v *NullableRecordingAvailableCallback) Set(val *RecordingAvailableCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingAvailableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingAvailableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingAvailableCallback(val *RecordingAvailableCallback) *NullableRecordingAvailableCallback {
	return &NullableRecordingAvailableCallback{value: val, isSet: true}
}

func (v NullableRecordingAvailableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingAvailableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


