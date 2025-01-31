# ConferenceCompletedCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The event type, value can be one of the following: answer, bridgeComplete, bridgeTargetComplete, conferenceCreated, conferenceRedirect, conferenceMemberJoin, conferenceMemberExit, conferenceCompleted, conferenceRecordingAvailable, disconnect, dtmf, gather, initiate, machineDetectionComplete, recordingComplete, recordingAvailable, redirect, transcriptionAvailable, transferAnswer, transferComplete, transferDisconnect. | [optional] 
**EventTime** | Pointer to **time.Time** | The approximate UTC date and time when the event was generated by the Bandwidth server, in ISO 8601 format. This may not be exactly the time of event execution. | [optional] 
**ConferenceId** | Pointer to **string** | The unique, Bandwidth-generated ID of the conference that was recorded | [optional] 
**Name** | Pointer to **string** | The user-specified name of the conference that was recorded | [optional] 
**Tag** | Pointer to **NullableString** | (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present. | [optional] 

## Methods

### NewConferenceCompletedCallback

`func NewConferenceCompletedCallback() *ConferenceCompletedCallback`

NewConferenceCompletedCallback instantiates a new ConferenceCompletedCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceCompletedCallbackWithDefaults

`func NewConferenceCompletedCallbackWithDefaults() *ConferenceCompletedCallback`

NewConferenceCompletedCallbackWithDefaults instantiates a new ConferenceCompletedCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ConferenceCompletedCallback) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ConferenceCompletedCallback) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ConferenceCompletedCallback) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ConferenceCompletedCallback) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventTime

`func (o *ConferenceCompletedCallback) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ConferenceCompletedCallback) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ConferenceCompletedCallback) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ConferenceCompletedCallback) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceCompletedCallback) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceCompletedCallback) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceCompletedCallback) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceCompletedCallback) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetName

`func (o *ConferenceCompletedCallback) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConferenceCompletedCallback) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConferenceCompletedCallback) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConferenceCompletedCallback) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTag

`func (o *ConferenceCompletedCallback) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ConferenceCompletedCallback) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ConferenceCompletedCallback) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ConferenceCompletedCallback) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ConferenceCompletedCallback) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ConferenceCompletedCallback) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


