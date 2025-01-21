# ConferenceMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | Pointer to **string** | The call id associated with the event. | [optional] 
**ConferenceId** | Pointer to **string** | The unique, Bandwidth-generated ID of the conference that was recorded | [optional] 
**MemberUrl** | Pointer to **string** | A URL that may be used to retrieve information about or update the state of this conference member. This is the URL of this member&#39;s [Get Conference Member](/apis/voice/#operation/getConferenceMember) endpoint and [Modify Conference Member](/apis/voice/#operation/updateConferenceMember) endpoint. | [optional] 
**Mute** | Pointer to **bool** | Whether or not this member is currently muted. Members who are muted are still able to hear other participants.  If used in a PUT request, updates this member&#39;s mute status. Has no effect if omitted. | [optional] 
**Hold** | Pointer to **bool** | Whether or not this member is currently on hold. Members who are on hold are not able to hear or speak in the conference.  If used in a PUT request, updates this member&#39;s hold status. Has no effect if omitted. | [optional] 
**CallIdsToCoach** | Pointer to **[]string** | If this member had a value set for &#x60;callIdsToCoach&#x60; in its [Conference](/docs/voice/bxml/conference) verb or this list was added with a previous PUT request to modify the member, this is that list of calls.  If present in a PUT request, modifies the calls that this member is coaching. Has no effect if omitted. See the documentation for the [Conference](/docs/voice/bxml/conference) verb for more details about coaching. Note that this will not add the matching calls to the conference; each call must individually execute a Conference verb to join. | [optional] 

## Methods

### NewConferenceMember

`func NewConferenceMember() *ConferenceMember`

NewConferenceMember instantiates a new ConferenceMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceMemberWithDefaults

`func NewConferenceMemberWithDefaults() *ConferenceMember`

NewConferenceMemberWithDefaults instantiates a new ConferenceMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *ConferenceMember) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *ConferenceMember) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *ConferenceMember) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *ConferenceMember) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceMember) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceMember) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceMember) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceMember) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetMemberUrl

`func (o *ConferenceMember) GetMemberUrl() string`

GetMemberUrl returns the MemberUrl field if non-nil, zero value otherwise.

### GetMemberUrlOk

`func (o *ConferenceMember) GetMemberUrlOk() (*string, bool)`

GetMemberUrlOk returns a tuple with the MemberUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUrl

`func (o *ConferenceMember) SetMemberUrl(v string)`

SetMemberUrl sets MemberUrl field to given value.

### HasMemberUrl

`func (o *ConferenceMember) HasMemberUrl() bool`

HasMemberUrl returns a boolean if a field has been set.

### GetMute

`func (o *ConferenceMember) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *ConferenceMember) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *ConferenceMember) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *ConferenceMember) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetHold

`func (o *ConferenceMember) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *ConferenceMember) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *ConferenceMember) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *ConferenceMember) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetCallIdsToCoach

`func (o *ConferenceMember) GetCallIdsToCoach() []string`

GetCallIdsToCoach returns the CallIdsToCoach field if non-nil, zero value otherwise.

### GetCallIdsToCoachOk

`func (o *ConferenceMember) GetCallIdsToCoachOk() (*[]string, bool)`

GetCallIdsToCoachOk returns a tuple with the CallIdsToCoach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIdsToCoach

`func (o *ConferenceMember) SetCallIdsToCoach(v []string)`

SetCallIdsToCoach sets CallIdsToCoach field to given value.

### HasCallIdsToCoach

`func (o *ConferenceMember) HasCallIdsToCoach() bool`

HasCallIdsToCoach returns a boolean if a field has been set.

### SetCallIdsToCoachNil

`func (o *ConferenceMember) SetCallIdsToCoachNil(b bool)`

 SetCallIdsToCoachNil sets the value for CallIdsToCoach to be an explicit nil

### UnsetCallIdsToCoach
`func (o *ConferenceMember) UnsetCallIdsToCoach()`

UnsetCallIdsToCoach ensures that no value is present for CallIdsToCoach, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


