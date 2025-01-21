# RecordingTranscriptionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique transcription ID | [optional] 
**Status** | Pointer to **string** | The current status of the process. For recording, current possible values are &#39;processing&#39;, &#39;partial&#39;, &#39;complete&#39;, &#39;deleted&#39;, and &#39;error&#39;. For transcriptions, current possible values are &#39;none&#39;, &#39;processing&#39;, &#39;available&#39;, &#39;error&#39;, &#39;timeout&#39;, &#39;file-size-too-big&#39;, and &#39;file-size-too-small&#39;. Additional states may be added in the future, so your application must be tolerant of unknown values. | [optional] 
**CompletedTime** | Pointer to **time.Time** | The time that the transcription was completed | [optional] 
**Url** | Pointer to **string** | The URL of the [transcription](#operation/getCallTranscription) | [optional] 

## Methods

### NewRecordingTranscriptionMetadata

`func NewRecordingTranscriptionMetadata() *RecordingTranscriptionMetadata`

NewRecordingTranscriptionMetadata instantiates a new RecordingTranscriptionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingTranscriptionMetadataWithDefaults

`func NewRecordingTranscriptionMetadataWithDefaults() *RecordingTranscriptionMetadata`

NewRecordingTranscriptionMetadataWithDefaults instantiates a new RecordingTranscriptionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecordingTranscriptionMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordingTranscriptionMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordingTranscriptionMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecordingTranscriptionMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *RecordingTranscriptionMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecordingTranscriptionMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecordingTranscriptionMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecordingTranscriptionMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompletedTime

`func (o *RecordingTranscriptionMetadata) GetCompletedTime() time.Time`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *RecordingTranscriptionMetadata) GetCompletedTimeOk() (*time.Time, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *RecordingTranscriptionMetadata) SetCompletedTime(v time.Time)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *RecordingTranscriptionMetadata) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetUrl

`func (o *RecordingTranscriptionMetadata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RecordingTranscriptionMetadata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RecordingTranscriptionMetadata) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RecordingTranscriptionMetadata) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


