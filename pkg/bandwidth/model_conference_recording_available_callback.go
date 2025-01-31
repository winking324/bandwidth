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

// checks if the ConferenceRecordingAvailableCallback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceRecordingAvailableCallback{}

// ConferenceRecordingAvailableCallback The Conference Recording Available event is sent after a conference recording has been processed. It indicates that the recording is available for download.
type ConferenceRecordingAvailableCallback struct {
	// The event type, value can be one of the following: answer, bridgeComplete, bridgeTargetComplete, conferenceCreated, conferenceRedirect, conferenceMemberJoin, conferenceMemberExit, conferenceCompleted, conferenceRecordingAvailable, disconnect, dtmf, gather, initiate, machineDetectionComplete, recordingComplete, recordingAvailable, redirect, transcriptionAvailable, transferAnswer, transferComplete, transferDisconnect.
	EventType *string `json:"eventType,omitempty"`
	// The approximate UTC date and time when the event was generated by the Bandwidth server, in ISO 8601 format. This may not be exactly the time of event execution.
	EventTime *time.Time `json:"eventTime,omitempty"`
	// The unique, Bandwidth-generated ID of the conference that was recorded
	ConferenceId *string `json:"conferenceId,omitempty"`
	// The user-specified name of the conference that was recorded
	Name *string `json:"name,omitempty"`
	// The user account associated with the call.
	AccountId *string `json:"accountId,omitempty"`
	// The unique ID of this recording
	RecordingId *string `json:"recordingId,omitempty"`
	// Always `1` for conference recordings; multi-channel recordings are not supported on conferences.
	Channels *int32 `json:"channels,omitempty"`
	// Time the call was started, in ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`
	// The time that the recording ended in ISO-8601 format
	EndTime *time.Time `json:"endTime,omitempty"`
	// The duration of the recording in ISO-8601 format
	Duration *string `json:"duration,omitempty"`
	FileFormat *FileFormatEnum `json:"fileFormat,omitempty"`
	// The URL that can be used to download the recording. Only present if the recording is finished and may be downloaded.
	MediaUrl NullableString `json:"mediaUrl,omitempty"`
	// (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present.
	Tag NullableString `json:"tag,omitempty"`
	// The current status of the process. For recording, current possible values are 'processing', 'partial', 'complete', 'deleted', and 'error'. For transcriptions, current possible values are 'none', 'processing', 'available', 'error', 'timeout', 'file-size-too-big', and 'file-size-too-small'. Additional states may be added in the future, so your application must be tolerant of unknown values.
	Status *string `json:"status,omitempty"`
}

// NewConferenceRecordingAvailableCallback instantiates a new ConferenceRecordingAvailableCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceRecordingAvailableCallback() *ConferenceRecordingAvailableCallback {
	this := ConferenceRecordingAvailableCallback{}
	return &this
}

// NewConferenceRecordingAvailableCallbackWithDefaults instantiates a new ConferenceRecordingAvailableCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceRecordingAvailableCallbackWithDefaults() *ConferenceRecordingAvailableCallback {
	this := ConferenceRecordingAvailableCallback{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *ConferenceRecordingAvailableCallback) SetEventType(v string) {
	o.EventType = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetEventTime() time.Time {
	if o == nil || IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *ConferenceRecordingAvailableCallback) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetConferenceId returns the ConferenceId field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetConferenceId() string {
	if o == nil || IsNil(o.ConferenceId) {
		var ret string
		return ret
	}
	return *o.ConferenceId
}

// GetConferenceIdOk returns a tuple with the ConferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetConferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConferenceId) {
		return nil, false
	}
	return o.ConferenceId, true
}

// HasConferenceId returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasConferenceId() bool {
	if o != nil && !IsNil(o.ConferenceId) {
		return true
	}

	return false
}

// SetConferenceId gets a reference to the given string and assigns it to the ConferenceId field.
func (o *ConferenceRecordingAvailableCallback) SetConferenceId(v string) {
	o.ConferenceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConferenceRecordingAvailableCallback) SetName(v string) {
	o.Name = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ConferenceRecordingAvailableCallback) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRecordingId returns the RecordingId field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetRecordingId() string {
	if o == nil || IsNil(o.RecordingId) {
		var ret string
		return ret
	}
	return *o.RecordingId
}

// GetRecordingIdOk returns a tuple with the RecordingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetRecordingIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordingId) {
		return nil, false
	}
	return o.RecordingId, true
}

// HasRecordingId returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasRecordingId() bool {
	if o != nil && !IsNil(o.RecordingId) {
		return true
	}

	return false
}

// SetRecordingId gets a reference to the given string and assigns it to the RecordingId field.
func (o *ConferenceRecordingAvailableCallback) SetRecordingId(v string) {
	o.RecordingId = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetChannels() int32 {
	if o == nil || IsNil(o.Channels) {
		var ret int32
		return ret
	}
	return *o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetChannelsOk() (*int32, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given int32 and assigns it to the Channels field.
func (o *ConferenceRecordingAvailableCallback) SetChannels(v int32) {
	o.Channels = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ConferenceRecordingAvailableCallback) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ConferenceRecordingAvailableCallback) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *ConferenceRecordingAvailableCallback) SetDuration(v string) {
	o.Duration = &v
}

// GetFileFormat returns the FileFormat field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetFileFormat() FileFormatEnum {
	if o == nil || IsNil(o.FileFormat) {
		var ret FileFormatEnum
		return ret
	}
	return *o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetFileFormatOk() (*FileFormatEnum, bool) {
	if o == nil || IsNil(o.FileFormat) {
		return nil, false
	}
	return o.FileFormat, true
}

// HasFileFormat returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasFileFormat() bool {
	if o != nil && !IsNil(o.FileFormat) {
		return true
	}

	return false
}

// SetFileFormat gets a reference to the given FileFormatEnum and assigns it to the FileFormat field.
func (o *ConferenceRecordingAvailableCallback) SetFileFormat(v FileFormatEnum) {
	o.FileFormat = &v
}

// GetMediaUrl returns the MediaUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceRecordingAvailableCallback) GetMediaUrl() string {
	if o == nil || IsNil(o.MediaUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MediaUrl.Get()
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceRecordingAvailableCallback) GetMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaUrl.Get(), o.MediaUrl.IsSet()
}

// HasMediaUrl returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasMediaUrl() bool {
	if o != nil && o.MediaUrl.IsSet() {
		return true
	}

	return false
}

// SetMediaUrl gets a reference to the given NullableString and assigns it to the MediaUrl field.
func (o *ConferenceRecordingAvailableCallback) SetMediaUrl(v string) {
	o.MediaUrl.Set(&v)
}
// SetMediaUrlNil sets the value for MediaUrl to be an explicit nil
func (o *ConferenceRecordingAvailableCallback) SetMediaUrlNil() {
	o.MediaUrl.Set(nil)
}

// UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
func (o *ConferenceRecordingAvailableCallback) UnsetMediaUrl() {
	o.MediaUrl.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceRecordingAvailableCallback) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceRecordingAvailableCallback) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *ConferenceRecordingAvailableCallback) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *ConferenceRecordingAvailableCallback) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *ConferenceRecordingAvailableCallback) UnsetTag() {
	o.Tag.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConferenceRecordingAvailableCallback) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingAvailableCallback) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConferenceRecordingAvailableCallback) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConferenceRecordingAvailableCallback) SetStatus(v string) {
	o.Status = &v
}

func (o ConferenceRecordingAvailableCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceRecordingAvailableCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.ConferenceId) {
		toSerialize["conferenceId"] = o.ConferenceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.RecordingId) {
		toSerialize["recordingId"] = o.RecordingId
	}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
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
	if o.MediaUrl.IsSet() {
		toSerialize["mediaUrl"] = o.MediaUrl.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableConferenceRecordingAvailableCallback struct {
	value *ConferenceRecordingAvailableCallback
	isSet bool
}

func (v NullableConferenceRecordingAvailableCallback) Get() *ConferenceRecordingAvailableCallback {
	return v.value
}

func (v *NullableConferenceRecordingAvailableCallback) Set(val *ConferenceRecordingAvailableCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceRecordingAvailableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceRecordingAvailableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceRecordingAvailableCallback(val *ConferenceRecordingAvailableCallback) *NullableConferenceRecordingAvailableCallback {
	return &NullableConferenceRecordingAvailableCallback{value: val, isSet: true}
}

func (v NullableConferenceRecordingAvailableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceRecordingAvailableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


