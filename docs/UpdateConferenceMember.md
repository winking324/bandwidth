# UpdateConferenceMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mute** | Pointer to **bool** | Whether or not this member is currently muted. Members who are muted are still able to hear other participants.  Updates this member&#39;s mute status. Has no effect if omitted. | [optional] 
**Hold** | Pointer to **bool** | Whether or not this member is currently on hold. Members who are on hold are not able to hear or speak in the conference.  Updates this member&#39;s hold status. Has no effect if omitted. | [optional] 
**CallIdsToCoach** | Pointer to **[]string** | If this member had a value set for &#x60;callIdsToCoach&#x60; in its [Conference](/docs/voice/bxml/conference) verb or this list was added with a previous PUT request to modify the member, this is that list of calls.  Modifies the calls that this member is coaching. Has no effect if omitted. See the documentation for the [Conference](/docs/voice/bxml/conference) verb for more details about coaching.  Note that this will not add the matching calls to the conference; each call must individually execute a Conference verb to join. | [optional] 

## Methods

### NewUpdateConferenceMember

`func NewUpdateConferenceMember() *UpdateConferenceMember`

NewUpdateConferenceMember instantiates a new UpdateConferenceMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConferenceMemberWithDefaults

`func NewUpdateConferenceMemberWithDefaults() *UpdateConferenceMember`

NewUpdateConferenceMemberWithDefaults instantiates a new UpdateConferenceMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMute

`func (o *UpdateConferenceMember) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *UpdateConferenceMember) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *UpdateConferenceMember) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *UpdateConferenceMember) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetHold

`func (o *UpdateConferenceMember) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *UpdateConferenceMember) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *UpdateConferenceMember) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *UpdateConferenceMember) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetCallIdsToCoach

`func (o *UpdateConferenceMember) GetCallIdsToCoach() []string`

GetCallIdsToCoach returns the CallIdsToCoach field if non-nil, zero value otherwise.

### GetCallIdsToCoachOk

`func (o *UpdateConferenceMember) GetCallIdsToCoachOk() (*[]string, bool)`

GetCallIdsToCoachOk returns a tuple with the CallIdsToCoach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIdsToCoach

`func (o *UpdateConferenceMember) SetCallIdsToCoach(v []string)`

SetCallIdsToCoach sets CallIdsToCoach field to given value.

### HasCallIdsToCoach

`func (o *UpdateConferenceMember) HasCallIdsToCoach() bool`

HasCallIdsToCoach returns a boolean if a field has been set.

### SetCallIdsToCoachNil

`func (o *UpdateConferenceMember) SetCallIdsToCoachNil(b bool)`

 SetCallIdsToCoachNil sets the value for CallIdsToCoach to be an explicit nil

### UnsetCallIdsToCoach
`func (o *UpdateConferenceMember) UnsetCallIdsToCoach()`

UnsetCallIdsToCoach ensures that no value is present for CallIdsToCoach, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


