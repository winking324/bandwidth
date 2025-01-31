# BridgeCompleteCallback

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
**Cause** | Pointer to **string** | Reason the call failed - hangup, busy, timeout, cancel, rejected, callback-error, invalid-bxml, application-error, account-limit, node-capacity-exceeded, error, or unknown. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Text explaining the reason that caused the call to fail in case of errors. | [optional] 
**ErrorId** | Pointer to **NullableString** | Bandwidth&#39;s internal id that references the error event. | [optional] 

## Methods

### NewBridgeCompleteCallback

`func NewBridgeCompleteCallback() *BridgeCompleteCallback`

NewBridgeCompleteCallback instantiates a new BridgeCompleteCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeCompleteCallbackWithDefaults

`func NewBridgeCompleteCallbackWithDefaults() *BridgeCompleteCallback`

NewBridgeCompleteCallbackWithDefaults instantiates a new BridgeCompleteCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *BridgeCompleteCallback) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BridgeCompleteCallback) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BridgeCompleteCallback) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BridgeCompleteCallback) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventTime

`func (o *BridgeCompleteCallback) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *BridgeCompleteCallback) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *BridgeCompleteCallback) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *BridgeCompleteCallback) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetAccountId

`func (o *BridgeCompleteCallback) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BridgeCompleteCallback) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BridgeCompleteCallback) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BridgeCompleteCallback) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApplicationId

`func (o *BridgeCompleteCallback) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *BridgeCompleteCallback) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *BridgeCompleteCallback) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *BridgeCompleteCallback) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetFrom

`func (o *BridgeCompleteCallback) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BridgeCompleteCallback) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BridgeCompleteCallback) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BridgeCompleteCallback) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *BridgeCompleteCallback) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BridgeCompleteCallback) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BridgeCompleteCallback) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BridgeCompleteCallback) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDirection

`func (o *BridgeCompleteCallback) GetDirection() CallDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BridgeCompleteCallback) GetDirectionOk() (*CallDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BridgeCompleteCallback) SetDirection(v CallDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BridgeCompleteCallback) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetCallId

`func (o *BridgeCompleteCallback) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *BridgeCompleteCallback) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *BridgeCompleteCallback) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *BridgeCompleteCallback) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallUrl

`func (o *BridgeCompleteCallback) GetCallUrl() string`

GetCallUrl returns the CallUrl field if non-nil, zero value otherwise.

### GetCallUrlOk

`func (o *BridgeCompleteCallback) GetCallUrlOk() (*string, bool)`

GetCallUrlOk returns a tuple with the CallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallUrl

`func (o *BridgeCompleteCallback) SetCallUrl(v string)`

SetCallUrl sets CallUrl field to given value.

### HasCallUrl

`func (o *BridgeCompleteCallback) HasCallUrl() bool`

HasCallUrl returns a boolean if a field has been set.

### GetEnqueuedTime

`func (o *BridgeCompleteCallback) GetEnqueuedTime() time.Time`

GetEnqueuedTime returns the EnqueuedTime field if non-nil, zero value otherwise.

### GetEnqueuedTimeOk

`func (o *BridgeCompleteCallback) GetEnqueuedTimeOk() (*time.Time, bool)`

GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedTime

`func (o *BridgeCompleteCallback) SetEnqueuedTime(v time.Time)`

SetEnqueuedTime sets EnqueuedTime field to given value.

### HasEnqueuedTime

`func (o *BridgeCompleteCallback) HasEnqueuedTime() bool`

HasEnqueuedTime returns a boolean if a field has been set.

### SetEnqueuedTimeNil

`func (o *BridgeCompleteCallback) SetEnqueuedTimeNil(b bool)`

 SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil

### UnsetEnqueuedTime
`func (o *BridgeCompleteCallback) UnsetEnqueuedTime()`

UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
### GetStartTime

`func (o *BridgeCompleteCallback) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BridgeCompleteCallback) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BridgeCompleteCallback) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BridgeCompleteCallback) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAnswerTime

`func (o *BridgeCompleteCallback) GetAnswerTime() time.Time`

GetAnswerTime returns the AnswerTime field if non-nil, zero value otherwise.

### GetAnswerTimeOk

`func (o *BridgeCompleteCallback) GetAnswerTimeOk() (*time.Time, bool)`

GetAnswerTimeOk returns a tuple with the AnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerTime

`func (o *BridgeCompleteCallback) SetAnswerTime(v time.Time)`

SetAnswerTime sets AnswerTime field to given value.

### HasAnswerTime

`func (o *BridgeCompleteCallback) HasAnswerTime() bool`

HasAnswerTime returns a boolean if a field has been set.

### SetAnswerTimeNil

`func (o *BridgeCompleteCallback) SetAnswerTimeNil(b bool)`

 SetAnswerTimeNil sets the value for AnswerTime to be an explicit nil

### UnsetAnswerTime
`func (o *BridgeCompleteCallback) UnsetAnswerTime()`

UnsetAnswerTime ensures that no value is present for AnswerTime, not even an explicit nil
### GetTag

`func (o *BridgeCompleteCallback) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BridgeCompleteCallback) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BridgeCompleteCallback) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BridgeCompleteCallback) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *BridgeCompleteCallback) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *BridgeCompleteCallback) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetCause

`func (o *BridgeCompleteCallback) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BridgeCompleteCallback) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BridgeCompleteCallback) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *BridgeCompleteCallback) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BridgeCompleteCallback) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BridgeCompleteCallback) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BridgeCompleteCallback) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BridgeCompleteCallback) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BridgeCompleteCallback) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BridgeCompleteCallback) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorId

`func (o *BridgeCompleteCallback) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *BridgeCompleteCallback) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *BridgeCompleteCallback) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *BridgeCompleteCallback) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### SetErrorIdNil

`func (o *BridgeCompleteCallback) SetErrorIdNil(b bool)`

 SetErrorIdNil sets the value for ErrorId to be an explicit nil

### UnsetErrorId
`func (o *BridgeCompleteCallback) UnsetErrorId()`

UnsetErrorId ensures that no value is present for ErrorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


