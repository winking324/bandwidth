# Transcription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | The transcribed text | [optional] 
**Confidence** | Pointer to **float64** | The confidence on the recognized content, ranging from &#x60;0.0&#x60; to &#x60;1.0&#x60; with &#x60;1.0&#x60; being the highest confidence. | [optional] 

## Methods

### NewTranscription

`func NewTranscription() *Transcription`

NewTranscription instantiates a new Transcription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionWithDefaults

`func NewTranscriptionWithDefaults() *Transcription`

NewTranscriptionWithDefaults instantiates a new Transcription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *Transcription) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Transcription) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Transcription) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Transcription) HasText() bool`

HasText returns a boolean if a field has been set.

### GetConfidence

`func (o *Transcription) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Transcription) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Transcription) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *Transcription) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


