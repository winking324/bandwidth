# BridgeTargetCompleteCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The event type, value can be one of the following: answer, bridgeComplete, bridgeTargetComplete, conferenceCreated, conferenceRedirect, conferenceMemberJoin, conferenceMemberExit, conferenceCompleted, conferenceRecordingAvailable, disconnect, dtmf, gather, initiate, machineDetectionComplete, recordingComplete, recordingAvailable, redirect, transcriptionAvailable, transferAnswer, transferComplete, transferDisconnect. | [optional] 
**EventTime** | Pointer to **time.Time** | The approximate UTC date and time when the event was generated by the Bandwidth server, in ISO 8601 format. This may not be exactly the time of event execution. | [optional] 
**AccountId** | Pointer to **string** | The user account associated with the call. | [optional] 
**ApplicationId** | Pointer to **string** | The id of the application associated with the call. | [optional] 
**From** | Pointer to **string** | The provided identifier of the caller. Must be a phone number in E.164 format (e.g. +15555555555). | [optional] 
**To** | Pointer to **string** | The phone number that received the call, in E.164 format (e.g. +15555555555). | [optional] 
**Direction** | Pointer to [**CallDirectionEnum**](CallDirectionEnum.md) |  | [optional] 
**CallId** | Pointer to **string** | The call id associated with the event. | [optional] 
**CallUrl** | Pointer to **string** | The URL of the call associated with the event. | [optional] 
**EnqueuedTime** | Pointer to **NullableTime** | (optional) If call queueing is enabled and this is an outbound call, time the call was queued, in ISO 8601 format. | [optional] 
**StartTime** | Pointer to **time.Time** | Time the call was started, in ISO 8601 format. | [optional] 
**AnswerTime** | Pointer to **NullableTime** | Time the call was answered, in ISO 8601 format. | [optional] 
**Tag** | Pointer to **NullableString** | (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present. | [optional] 

## Methods

### NewBridgeTargetCompleteCallback

`func NewBridgeTargetCompleteCallback() *BridgeTargetCompleteCallback`

NewBridgeTargetCompleteCallback instantiates a new BridgeTargetCompleteCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeTargetCompleteCallbackWithDefaults

`func NewBridgeTargetCompleteCallbackWithDefaults() *BridgeTargetCompleteCallback`

NewBridgeTargetCompleteCallbackWithDefaults instantiates a new BridgeTargetCompleteCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BridgeTargetCompleteCallback) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BridgeTargetCompleteCallback) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BridgeTargetCompleteCallback) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BridgeTargetCompleteCallback) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventTime

`func (o *BridgeTargetCompleteCallback) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *BridgeTargetCompleteCallback) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *BridgeTargetCompleteCallback) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *BridgeTargetCompleteCallback) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetAccountId

`func (o *BridgeTargetCompleteCallback) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BridgeTargetCompleteCallback) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BridgeTargetCompleteCallback) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BridgeTargetCompleteCallback) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApplicationId

`func (o *BridgeTargetCompleteCallback) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *BridgeTargetCompleteCallback) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *BridgeTargetCompleteCallback) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *BridgeTargetCompleteCallback) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetFrom

`func (o *BridgeTargetCompleteCallback) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BridgeTargetCompleteCallback) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BridgeTargetCompleteCallback) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BridgeTargetCompleteCallback) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *BridgeTargetCompleteCallback) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BridgeTargetCompleteCallback) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BridgeTargetCompleteCallback) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BridgeTargetCompleteCallback) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDirection

`func (o *BridgeTargetCompleteCallback) GetDirection() CallDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BridgeTargetCompleteCallback) GetDirectionOk() (*CallDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BridgeTargetCompleteCallback) SetDirection(v CallDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BridgeTargetCompleteCallback) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetCallId

`func (o *BridgeTargetCompleteCallback) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *BridgeTargetCompleteCallback) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *BridgeTargetCompleteCallback) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *BridgeTargetCompleteCallback) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallUrl

`func (o *BridgeTargetCompleteCallback) GetCallUrl() string`

GetCallUrl returns the CallUrl field if non-nil, zero value otherwise.

### GetCallUrlOk

`func (o *BridgeTargetCompleteCallback) GetCallUrlOk() (*string, bool)`

GetCallUrlOk returns a tuple with the CallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallUrl

`func (o *BridgeTargetCompleteCallback) SetCallUrl(v string)`

SetCallUrl sets CallUrl field to given value.

### HasCallUrl

`func (o *BridgeTargetCompleteCallback) HasCallUrl() bool`

HasCallUrl returns a boolean if a field has been set.

### GetEnqueuedTime

`func (o *BridgeTargetCompleteCallback) GetEnqueuedTime() time.Time`

GetEnqueuedTime returns the EnqueuedTime field if non-nil, zero value otherwise.

### GetEnqueuedTimeOk

`func (o *BridgeTargetCompleteCallback) GetEnqueuedTimeOk() (*time.Time, bool)`

GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedTime

`func (o *BridgeTargetCompleteCallback) SetEnqueuedTime(v time.Time)`

SetEnqueuedTime sets EnqueuedTime field to given value.

### HasEnqueuedTime

`func (o *BridgeTargetCompleteCallback) HasEnqueuedTime() bool`

HasEnqueuedTime returns a boolean if a field has been set.

### SetEnqueuedTimeNil

`func (o *BridgeTargetCompleteCallback) SetEnqueuedTimeNil(b bool)`

 SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil

### UnsetEnqueuedTime
`func (o *BridgeTargetCompleteCallback) UnsetEnqueuedTime()`

UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
### GetStartTime

`func (o *BridgeTargetCompleteCallback) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BridgeTargetCompleteCallback) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BridgeTargetCompleteCallback) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BridgeTargetCompleteCallback) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAnswerTime

`func (o *BridgeTargetCompleteCallback) GetAnswerTime() time.Time`

GetAnswerTime returns the AnswerTime field if non-nil, zero value otherwise.

### GetAnswerTimeOk

`func (o *BridgeTargetCompleteCallback) GetAnswerTimeOk() (*time.Time, bool)`

GetAnswerTimeOk returns a tuple with the AnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerTime

`func (o *BridgeTargetCompleteCallback) SetAnswerTime(v time.Time)`

SetAnswerTime sets AnswerTime field to given value.

### HasAnswerTime

`func (o *BridgeTargetCompleteCallback) HasAnswerTime() bool`

HasAnswerTime returns a boolean if a field has been set.

### SetAnswerTimeNil

`func (o *BridgeTargetCompleteCallback) SetAnswerTimeNil(b bool)`

 SetAnswerTimeNil sets the value for AnswerTime to be an explicit nil

### UnsetAnswerTime
`func (o *BridgeTargetCompleteCallback) UnsetAnswerTime()`

UnsetAnswerTime ensures that no value is present for AnswerTime, not even an explicit nil
### GetTag

`func (o *BridgeTargetCompleteCallback) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BridgeTargetCompleteCallback) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BridgeTargetCompleteCallback) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BridgeTargetCompleteCallback) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *BridgeTargetCompleteCallback) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *BridgeTargetCompleteCallback) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


