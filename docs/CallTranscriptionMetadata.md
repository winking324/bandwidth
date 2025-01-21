# CallTranscriptionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranscriptionId** | Pointer to **string** | The programmable voice API transcription ID. | [optional] 
**TranscriptionName** | Pointer to **string** | The programmable voice API transcription name. This name could be provided by the user when creating the transcription. | [optional] 
**TranscriptionUrl** | Pointer to **string** | A URL that may be used to retrieve the transcription itself. This points to the [Get Call Transcription](/apis/voice/#operation/getCallTranscription) endpoint. | [optional] 

## Methods

### NewCallTranscriptionMetadata

`func NewCallTranscriptionMetadata() *CallTranscriptionMetadata`

NewCallTranscriptionMetadata instantiates a new CallTranscriptionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTranscriptionMetadataWithDefaults

`func NewCallTranscriptionMetadataWithDefaults() *CallTranscriptionMetadata`

NewCallTranscriptionMetadataWithDefaults instantiates a new CallTranscriptionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranscriptionId

`func (o *CallTranscriptionMetadata) GetTranscriptionId() string`

GetTranscriptionId returns the TranscriptionId field if non-nil, zero value otherwise.

### GetTranscriptionIdOk

`func (o *CallTranscriptionMetadata) GetTranscriptionIdOk() (*string, bool)`

GetTranscriptionIdOk returns a tuple with the TranscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionId

`func (o *CallTranscriptionMetadata) SetTranscriptionId(v string)`

SetTranscriptionId sets TranscriptionId field to given value.

### HasTranscriptionId

`func (o *CallTranscriptionMetadata) HasTranscriptionId() bool`

HasTranscriptionId returns a boolean if a field has been set.

### GetTranscriptionName

`func (o *CallTranscriptionMetadata) GetTranscriptionName() string`

GetTranscriptionName returns the TranscriptionName field if non-nil, zero value otherwise.

### GetTranscriptionNameOk

`func (o *CallTranscriptionMetadata) GetTranscriptionNameOk() (*string, bool)`

GetTranscriptionNameOk returns a tuple with the TranscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionName

`func (o *CallTranscriptionMetadata) SetTranscriptionName(v string)`

SetTranscriptionName sets TranscriptionName field to given value.

### HasTranscriptionName

`func (o *CallTranscriptionMetadata) HasTranscriptionName() bool`

HasTranscriptionName returns a boolean if a field has been set.

### GetTranscriptionUrl

`func (o *CallTranscriptionMetadata) GetTranscriptionUrl() string`

GetTranscriptionUrl returns the TranscriptionUrl field if non-nil, zero value otherwise.

### GetTranscriptionUrlOk

`func (o *CallTranscriptionMetadata) GetTranscriptionUrlOk() (*string, bool)`

GetTranscriptionUrlOk returns a tuple with the TranscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionUrl

`func (o *CallTranscriptionMetadata) SetTranscriptionUrl(v string)`

SetTranscriptionUrl sets TranscriptionUrl field to given value.

### HasTranscriptionUrl

`func (o *CallTranscriptionMetadata) HasTranscriptionUrl() bool`

HasTranscriptionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


