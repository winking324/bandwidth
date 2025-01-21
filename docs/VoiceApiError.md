# VoiceApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoiceApiError

`func NewVoiceApiError() *VoiceApiError`

NewVoiceApiError instantiates a new VoiceApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceApiErrorWithDefaults

`func NewVoiceApiErrorWithDefaults() *VoiceApiError`

NewVoiceApiErrorWithDefaults instantiates a new VoiceApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VoiceApiError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoiceApiError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoiceApiError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VoiceApiError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *VoiceApiError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoiceApiError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoiceApiError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoiceApiError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *VoiceApiError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoiceApiError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoiceApiError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoiceApiError) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VoiceApiError) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VoiceApiError) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


