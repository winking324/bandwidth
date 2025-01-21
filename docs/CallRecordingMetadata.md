# CallRecordingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | The id of the application associated with the call. | [optional] 
**AccountId** | Pointer to **string** | The user account associated with the call. | [optional] 
**CallId** | Pointer to **string** | The call id associated with the event. | [optional] 
**ParentCallId** | Pointer to **string** | (optional) If the event is related to the B leg of a &lt;Transfer&gt;, the call id of the original call leg that executed the &lt;Transfer&gt;. Otherwise, this field will not be present. | [optional] 
**RecordingId** | Pointer to **string** | The unique ID of this recording | [optional] 
**To** | Pointer to **string** | The phone number that received the call, in E.164 format (e.g. +15555555555). | [optional] 
**From** | Pointer to **string** | The provided identifier of the caller. Must be a phone number in E.164 format (e.g. +15555555555). | [optional] 
**TransferCallerId** | Pointer to **string** | The phone number used as the from field of the B-leg call, in E.164 format (e.g. +15555555555). | [optional] 
**TransferTo** | Pointer to **string** | The phone number used as the to field of the B-leg call, in E.164 format (e.g. +15555555555). | [optional] 
**Duration** | Pointer to **string** | The duration of the recording in ISO-8601 format | [optional] 
**Direction** | Pointer to [**CallDirectionEnum**](CallDirectionEnum.md) |  | [optional] 
**Channels** | Pointer to **int32** | Always &#x60;1&#x60; for conference recordings; multi-channel recordings are not supported on conferences. | [optional] 
**StartTime** | Pointer to **time.Time** | Time the call was started, in ISO 8601 format. | [optional] 
**EndTime** | Pointer to **time.Time** | The time that the recording ended in ISO-8601 format | [optional] 
**FileFormat** | Pointer to [**FileFormatEnum**](FileFormatEnum.md) |  | [optional] 
**Status** | Pointer to **string** | The current status of the process. For recording, current possible values are &#39;processing&#39;, &#39;partial&#39;, &#39;complete&#39;, &#39;deleted&#39;, and &#39;error&#39;. For transcriptions, current possible values are &#39;none&#39;, &#39;processing&#39;, &#39;available&#39;, &#39;error&#39;, &#39;timeout&#39;, &#39;file-size-too-big&#39;, and &#39;file-size-too-small&#39;. Additional states may be added in the future, so your application must be tolerant of unknown values. | [optional] 
**MediaUrl** | Pointer to **NullableString** | The URL that can be used to download the recording. Only present if the recording is finished and may be downloaded. | [optional] 
**Transcription** | Pointer to [**NullableRecordingTranscriptionMetadata**](RecordingTranscriptionMetadata.md) |  | [optional] 
**RecordingName** | Pointer to **string** | A name to identify this recording. | [optional] 

## Methods

### NewCallRecordingMetadata

`func NewCallRecordingMetadata() *CallRecordingMetadata`

NewCallRecordingMetadata instantiates a new CallRecordingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordingMetadataWithDefaults

`func NewCallRecordingMetadataWithDefaults() *CallRecordingMetadata`

NewCallRecordingMetadataWithDefaults instantiates a new CallRecordingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CallRecordingMetadata) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CallRecordingMetadata) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CallRecordingMetadata) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CallRecordingMetadata) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAccountId

`func (o *CallRecordingMetadata) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CallRecordingMetadata) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CallRecordingMetadata) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CallRecordingMetadata) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCallId

`func (o *CallRecordingMetadata) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *CallRecordingMetadata) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *CallRecordingMetadata) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *CallRecordingMetadata) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetParentCallId

`func (o *CallRecordingMetadata) GetParentCallId() string`

GetParentCallId returns the ParentCallId field if non-nil, zero value otherwise.

### GetParentCallIdOk

`func (o *CallRecordingMetadata) GetParentCallIdOk() (*string, bool)`

GetParentCallIdOk returns a tuple with the ParentCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCallId

`func (o *CallRecordingMetadata) SetParentCallId(v string)`

SetParentCallId sets ParentCallId field to given value.

### HasParentCallId

`func (o *CallRecordingMetadata) HasParentCallId() bool`

HasParentCallId returns a boolean if a field has been set.

### GetRecordingId

`func (o *CallRecordingMetadata) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *CallRecordingMetadata) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *CallRecordingMetadata) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *CallRecordingMetadata) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.

### GetTo

`func (o *CallRecordingMetadata) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallRecordingMetadata) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallRecordingMetadata) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallRecordingMetadata) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *CallRecordingMetadata) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallRecordingMetadata) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallRecordingMetadata) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallRecordingMetadata) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTransferCallerId

`func (o *CallRecordingMetadata) GetTransferCallerId() string`

GetTransferCallerId returns the TransferCallerId field if non-nil, zero value otherwise.

### GetTransferCallerIdOk

`func (o *CallRecordingMetadata) GetTransferCallerIdOk() (*string, bool)`

GetTransferCallerIdOk returns a tuple with the TransferCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferCallerId

`func (o *CallRecordingMetadata) SetTransferCallerId(v string)`

SetTransferCallerId sets TransferCallerId field to given value.

### HasTransferCallerId

`func (o *CallRecordingMetadata) HasTransferCallerId() bool`

HasTransferCallerId returns a boolean if a field has been set.

### GetTransferTo

`func (o *CallRecordingMetadata) GetTransferTo() string`

GetTransferTo returns the TransferTo field if non-nil, zero value otherwise.

### GetTransferToOk

`func (o *CallRecordingMetadata) GetTransferToOk() (*string, bool)`

GetTransferToOk returns a tuple with the TransferTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTo

`func (o *CallRecordingMetadata) SetTransferTo(v string)`

SetTransferTo sets TransferTo field to given value.

### HasTransferTo

`func (o *CallRecordingMetadata) HasTransferTo() bool`

HasTransferTo returns a boolean if a field has been set.

### GetDuration

`func (o *CallRecordingMetadata) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CallRecordingMetadata) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CallRecordingMetadata) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CallRecordingMetadata) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDirection

`func (o *CallRecordingMetadata) GetDirection() CallDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CallRecordingMetadata) GetDirectionOk() (*CallDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CallRecordingMetadata) SetDirection(v CallDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CallRecordingMetadata) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetChannels

`func (o *CallRecordingMetadata) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CallRecordingMetadata) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CallRecordingMetadata) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CallRecordingMetadata) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetStartTime

`func (o *CallRecordingMetadata) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallRecordingMetadata) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallRecordingMetadata) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallRecordingMetadata) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *CallRecordingMetadata) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CallRecordingMetadata) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CallRecordingMetadata) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CallRecordingMetadata) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFileFormat

`func (o *CallRecordingMetadata) GetFileFormat() FileFormatEnum`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *CallRecordingMetadata) GetFileFormatOk() (*FileFormatEnum, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *CallRecordingMetadata) SetFileFormat(v FileFormatEnum)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *CallRecordingMetadata) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetStatus

`func (o *CallRecordingMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallRecordingMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallRecordingMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallRecordingMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMediaUrl

`func (o *CallRecordingMetadata) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *CallRecordingMetadata) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *CallRecordingMetadata) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *CallRecordingMetadata) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *CallRecordingMetadata) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *CallRecordingMetadata) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
### GetTranscription

`func (o *CallRecordingMetadata) GetTranscription() RecordingTranscriptionMetadata`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *CallRecordingMetadata) GetTranscriptionOk() (*RecordingTranscriptionMetadata, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *CallRecordingMetadata) SetTranscription(v RecordingTranscriptionMetadata)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *CallRecordingMetadata) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### SetTranscriptionNil

`func (o *CallRecordingMetadata) SetTranscriptionNil(b bool)`

 SetTranscriptionNil sets the value for Transcription to be an explicit nil

### UnsetTranscription
`func (o *CallRecordingMetadata) UnsetTranscription()`

UnsetTranscription ensures that no value is present for Transcription, not even an explicit nil
### GetRecordingName

`func (o *CallRecordingMetadata) GetRecordingName() string`

GetRecordingName returns the RecordingName field if non-nil, zero value otherwise.

### GetRecordingNameOk

`func (o *CallRecordingMetadata) GetRecordingNameOk() (*string, bool)`

GetRecordingNameOk returns a tuple with the RecordingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingName

`func (o *CallRecordingMetadata) SetRecordingName(v string)`

SetRecordingName sets RecordingName field to given value.

### HasRecordingName

`func (o *CallRecordingMetadata) HasRecordingName() bool`

HasRecordingName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


