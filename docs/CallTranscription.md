# CallTranscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectedLanguage** | Pointer to [**CallTranscriptionDetectedLanguageEnum**](CallTranscriptionDetectedLanguageEnum.md) |  | [optional] 
**Track** | Pointer to [**CallTranscriptionTrackEnum**](CallTranscriptionTrackEnum.md) |  | [optional] 
**Transcript** | Pointer to **string** | The transcription itself. | [optional] 
**Confidence** | Pointer to **float64** | How confident the transcription engine was in transcribing the associated audio (from &#x60;0&#x60; to &#x60;1&#x60;). | [optional] 

## Methods

### NewCallTranscription

`func NewCallTranscription() *CallTranscription`

NewCallTranscription instantiates a new CallTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTranscriptionWithDefaults

`func NewCallTranscriptionWithDefaults() *CallTranscription`

NewCallTranscriptionWithDefaults instantiates a new CallTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectedLanguage

`func (o *CallTranscription) GetDetectedLanguage() CallTranscriptionDetectedLanguageEnum`

GetDetectedLanguage returns the DetectedLanguage field if non-nil, zero value otherwise.

### GetDetectedLanguageOk

`func (o *CallTranscription) GetDetectedLanguageOk() (*CallTranscriptionDetectedLanguageEnum, bool)`

GetDetectedLanguageOk returns a tuple with the DetectedLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedLanguage

`func (o *CallTranscription) SetDetectedLanguage(v CallTranscriptionDetectedLanguageEnum)`

SetDetectedLanguage sets DetectedLanguage field to given value.

### HasDetectedLanguage

`func (o *CallTranscription) HasDetectedLanguage() bool`

HasDetectedLanguage returns a boolean if a field has been set.

### GetTrack

`func (o *CallTranscription) GetTrack() CallTranscriptionTrackEnum`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *CallTranscription) GetTrackOk() (*CallTranscriptionTrackEnum, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *CallTranscription) SetTrack(v CallTranscriptionTrackEnum)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *CallTranscription) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetTranscript

`func (o *CallTranscription) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *CallTranscription) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *CallTranscription) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *CallTranscription) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetConfidence

`func (o *CallTranscription) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CallTranscription) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CallTranscription) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CallTranscription) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


