# InitiateCallback

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
**StartTime** | Pointer to **time.Time** | Time the call was started, in ISO 8601 format. | [optional] 
**Diversion** | Pointer to [**Diversion**](Diversion.md) |  | [optional] 
**StirShaken** | Pointer to [**StirShaken**](StirShaken.md) |  | [optional] 

## Methods

### NewInitiateCallback

`func NewInitiateCallback() *InitiateCallback`

NewInitiateCallback instantiates a new InitiateCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateCallbackWithDefaults

`func NewInitiateCallbackWithDefaults() *InitiateCallback`

NewInitiateCallbackWithDefaults instantiates a new InitiateCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *InitiateCallback) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *InitiateCallback) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *InitiateCallback) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *InitiateCallback) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventTime

`func (o *InitiateCallback) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *InitiateCallback) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *InitiateCallback) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *InitiateCallback) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetAccountId

`func (o *InitiateCallback) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InitiateCallback) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InitiateCallback) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InitiateCallback) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApplicationId

`func (o *InitiateCallback) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *InitiateCallback) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *InitiateCallback) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *InitiateCallback) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetFrom

`func (o *InitiateCallback) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InitiateCallback) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InitiateCallback) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *InitiateCallback) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *InitiateCallback) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InitiateCallback) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InitiateCallback) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *InitiateCallback) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDirection

`func (o *InitiateCallback) GetDirection() CallDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InitiateCallback) GetDirectionOk() (*CallDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InitiateCallback) SetDirection(v CallDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InitiateCallback) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetCallId

`func (o *InitiateCallback) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *InitiateCallback) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *InitiateCallback) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *InitiateCallback) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallUrl

`func (o *InitiateCallback) GetCallUrl() string`

GetCallUrl returns the CallUrl field if non-nil, zero value otherwise.

### GetCallUrlOk

`func (o *InitiateCallback) GetCallUrlOk() (*string, bool)`

GetCallUrlOk returns a tuple with the CallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallUrl

`func (o *InitiateCallback) SetCallUrl(v string)`

SetCallUrl sets CallUrl field to given value.

### HasCallUrl

`func (o *InitiateCallback) HasCallUrl() bool`

HasCallUrl returns a boolean if a field has been set.

### GetStartTime

`func (o *InitiateCallback) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InitiateCallback) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InitiateCallback) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InitiateCallback) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDiversion

`func (o *InitiateCallback) GetDiversion() Diversion`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *InitiateCallback) GetDiversionOk() (*Diversion, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *InitiateCallback) SetDiversion(v Diversion)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *InitiateCallback) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetStirShaken

`func (o *InitiateCallback) GetStirShaken() StirShaken`

GetStirShaken returns the StirShaken field if non-nil, zero value otherwise.

### GetStirShakenOk

`func (o *InitiateCallback) GetStirShakenOk() (*StirShaken, bool)`

GetStirShakenOk returns a tuple with the StirShaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStirShaken

`func (o *InitiateCallback) SetStirShaken(v StirShaken)`

SetStirShaken sets StirShaken field to given value.

### HasStirShaken

`func (o *InitiateCallback) HasStirShaken() bool`

HasStirShaken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


