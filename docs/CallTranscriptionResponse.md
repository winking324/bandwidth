# CallTranscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The user account associated with the call. | [optional] 
**CallId** | Pointer to **string** | The call id associated with the event. | [optional] 
**TranscriptionId** | Pointer to **string** | The programmable voice API transcription ID. | [optional] 
**Tracks** | Pointer to [**[]CallTranscription**](CallTranscription.md) |  | [optional] 

## Methods

### NewCallTranscriptionResponse

`func NewCallTranscriptionResponse() *CallTranscriptionResponse`

NewCallTranscriptionResponse instantiates a new CallTranscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTranscriptionResponseWithDefaults

`func NewCallTranscriptionResponseWithDefaults() *CallTranscriptionResponse`

NewCallTranscriptionResponseWithDefaults instantiates a new CallTranscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CallTranscriptionResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CallTranscriptionResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CallTranscriptionResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CallTranscriptionResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCallId

`func (o *CallTranscriptionResponse) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *CallTranscriptionResponse) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *CallTranscriptionResponse) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *CallTranscriptionResponse) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetTranscriptionId

`func (o *CallTranscriptionResponse) GetTranscriptionId() string`

GetTranscriptionId returns the TranscriptionId field if non-nil, zero value otherwise.

### GetTranscriptionIdOk

`func (o *CallTranscriptionResponse) GetTranscriptionIdOk() (*string, bool)`

GetTranscriptionIdOk returns a tuple with the TranscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionId

`func (o *CallTranscriptionResponse) SetTranscriptionId(v string)`

SetTranscriptionId sets TranscriptionId field to given value.

### HasTranscriptionId

`func (o *CallTranscriptionResponse) HasTranscriptionId() bool`

HasTranscriptionId returns a boolean if a field has been set.

### GetTracks

`func (o *CallTranscriptionResponse) GetTracks() []CallTranscription`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *CallTranscriptionResponse) GetTracksOk() (*[]CallTranscription, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *CallTranscriptionResponse) SetTracks(v []CallTranscription)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *CallTranscriptionResponse) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


