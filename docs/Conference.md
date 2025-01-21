# Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Bandwidth-generated conference ID. | [optional] 
**Name** | Pointer to **string** | The name of the conference, as specified by your application. | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time the conference was initiated, in ISO 8601 format. | [optional] 
**CompletedTime** | Pointer to **NullableTime** | The time the conference was terminated, in ISO 8601 format. | [optional] 
**ConferenceEventUrl** | Pointer to **NullableString** | The URL to send the conference-related events. | [optional] 
**ConferenceEventMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**Tag** | Pointer to **NullableString** | The custom string attached to the conference that will be sent with callbacks. | [optional] 
**ActiveMembers** | Pointer to [**[]ConferenceMember**](ConferenceMember.md) | A list of active members of the conference. Omitted if this is a response to the [Get Conferences endpoint](/apis/voice#tag/Conferences/operation/listConferences). | [optional] 

## Methods

### NewConference

`func NewConference() *Conference`

NewConference instantiates a new Conference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceWithDefaults

`func NewConferenceWithDefaults() *Conference`

NewConferenceWithDefaults instantiates a new Conference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Conference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Conference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Conference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Conference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Conference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Conference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Conference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Conference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Conference) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Conference) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Conference) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Conference) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCompletedTime

`func (o *Conference) GetCompletedTime() time.Time`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *Conference) GetCompletedTimeOk() (*time.Time, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *Conference) SetCompletedTime(v time.Time)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *Conference) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### SetCompletedTimeNil

`func (o *Conference) SetCompletedTimeNil(b bool)`

 SetCompletedTimeNil sets the value for CompletedTime to be an explicit nil

### UnsetCompletedTime
`func (o *Conference) UnsetCompletedTime()`

UnsetCompletedTime ensures that no value is present for CompletedTime, not even an explicit nil
### GetConferenceEventUrl

`func (o *Conference) GetConferenceEventUrl() string`

GetConferenceEventUrl returns the ConferenceEventUrl field if non-nil, zero value otherwise.

### GetConferenceEventUrlOk

`func (o *Conference) GetConferenceEventUrlOk() (*string, bool)`

GetConferenceEventUrlOk returns a tuple with the ConferenceEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceEventUrl

`func (o *Conference) SetConferenceEventUrl(v string)`

SetConferenceEventUrl sets ConferenceEventUrl field to given value.

### HasConferenceEventUrl

`func (o *Conference) HasConferenceEventUrl() bool`

HasConferenceEventUrl returns a boolean if a field has been set.

### SetConferenceEventUrlNil

`func (o *Conference) SetConferenceEventUrlNil(b bool)`

 SetConferenceEventUrlNil sets the value for ConferenceEventUrl to be an explicit nil

### UnsetConferenceEventUrl
`func (o *Conference) UnsetConferenceEventUrl()`

UnsetConferenceEventUrl ensures that no value is present for ConferenceEventUrl, not even an explicit nil
### GetConferenceEventMethod

`func (o *Conference) GetConferenceEventMethod() CallbackMethodEnum`

GetConferenceEventMethod returns the ConferenceEventMethod field if non-nil, zero value otherwise.

### GetConferenceEventMethodOk

`func (o *Conference) GetConferenceEventMethodOk() (*CallbackMethodEnum, bool)`

GetConferenceEventMethodOk returns a tuple with the ConferenceEventMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceEventMethod

`func (o *Conference) SetConferenceEventMethod(v CallbackMethodEnum)`

SetConferenceEventMethod sets ConferenceEventMethod field to given value.

### HasConferenceEventMethod

`func (o *Conference) HasConferenceEventMethod() bool`

HasConferenceEventMethod returns a boolean if a field has been set.

### SetConferenceEventMethodNil

`func (o *Conference) SetConferenceEventMethodNil(b bool)`

 SetConferenceEventMethodNil sets the value for ConferenceEventMethod to be an explicit nil

### UnsetConferenceEventMethod
`func (o *Conference) UnsetConferenceEventMethod()`

UnsetConferenceEventMethod ensures that no value is present for ConferenceEventMethod, not even an explicit nil
### GetTag

`func (o *Conference) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Conference) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Conference) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Conference) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *Conference) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *Conference) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetActiveMembers

`func (o *Conference) GetActiveMembers() []ConferenceMember`

GetActiveMembers returns the ActiveMembers field if non-nil, zero value otherwise.

### GetActiveMembersOk

`func (o *Conference) GetActiveMembersOk() (*[]ConferenceMember, bool)`

GetActiveMembersOk returns a tuple with the ActiveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMembers

`func (o *Conference) SetActiveMembers(v []ConferenceMember)`

SetActiveMembers sets ActiveMembers field to given value.

### HasActiveMembers

`func (o *Conference) HasActiveMembers() bool`

HasActiveMembers returns a boolean if a field has been set.

### SetActiveMembersNil

`func (o *Conference) SetActiveMembersNil(b bool)`

 SetActiveMembersNil sets the value for ActiveMembers to be an explicit nil

### UnsetActiveMembers
`func (o *Conference) UnsetActiveMembers()`

UnsetActiveMembers ensures that no value is present for ActiveMembers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


