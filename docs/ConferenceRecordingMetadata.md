# ConferenceRecordingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The user account associated with the call. | [optional] 
**ConferenceId** | Pointer to **string** | The unique, Bandwidth-generated ID of the conference that was recorded | [optional] 
**Name** | Pointer to **string** | The user-specified name of the conference that was recorded | [optional] 
**RecordingId** | Pointer to **string** | The unique ID of this recording | [optional] 
**Duration** | Pointer to **string** | The duration of the recording in ISO-8601 format | [optional] 
**Channels** | Pointer to **int32** | Always &#x60;1&#x60; for conference recordings; multi-channel recordings are not supported on conferences. | [optional] 
**StartTime** | Pointer to **time.Time** | Time the call was started, in ISO 8601 format. | [optional] 
**EndTime** | Pointer to **time.Time** | The time that the recording ended in ISO-8601 format | [optional] 
**FileFormat** | Pointer to [**FileFormatEnum**](FileFormatEnum.md) |  | [optional] 
**Status** | Pointer to **string** | The current status of the process. For recording, current possible values are &#39;processing&#39;, &#39;partial&#39;, &#39;complete&#39;, &#39;deleted&#39;, and &#39;error&#39;. For transcriptions, current possible values are &#39;none&#39;, &#39;processing&#39;, &#39;available&#39;, &#39;error&#39;, &#39;timeout&#39;, &#39;file-size-too-big&#39;, and &#39;file-size-too-small&#39;. Additional states may be added in the future, so your application must be tolerant of unknown values. | [optional] 
**MediaUrl** | Pointer to **NullableString** | The URL that can be used to download the recording. Only present if the recording is finished and may be downloaded. | [optional] 
**RecordingName** | Pointer to **string** | A name to identify this recording. | [optional] 

## Methods

### NewConferenceRecordingMetadata

`func NewConferenceRecordingMetadata() *ConferenceRecordingMetadata`

NewConferenceRecordingMetadata instantiates a new ConferenceRecordingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceRecordingMetadataWithDefaults

`func NewConferenceRecordingMetadataWithDefaults() *ConferenceRecordingMetadata`

NewConferenceRecordingMetadataWithDefaults instantiates a new ConferenceRecordingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ConferenceRecordingMetadata) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ConferenceRecordingMetadata) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ConferenceRecordingMetadata) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ConferenceRecordingMetadata) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceRecordingMetadata) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceRecordingMetadata) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceRecordingMetadata) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceRecordingMetadata) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetName

`func (o *ConferenceRecordingMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConferenceRecordingMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConferenceRecordingMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConferenceRecordingMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordingId

`func (o *ConferenceRecordingMetadata) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *ConferenceRecordingMetadata) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *ConferenceRecordingMetadata) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *ConferenceRecordingMetadata) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.

### GetDuration

`func (o *ConferenceRecordingMetadata) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ConferenceRecordingMetadata) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ConferenceRecordingMetadata) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ConferenceRecordingMetadata) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetChannels

`func (o *ConferenceRecordingMetadata) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ConferenceRecordingMetadata) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ConferenceRecordingMetadata) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ConferenceRecordingMetadata) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetStartTime

`func (o *ConferenceRecordingMetadata) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ConferenceRecordingMetadata) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ConferenceRecordingMetadata) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ConferenceRecordingMetadata) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ConferenceRecordingMetadata) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ConferenceRecordingMetadata) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ConferenceRecordingMetadata) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ConferenceRecordingMetadata) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFileFormat

`func (o *ConferenceRecordingMetadata) GetFileFormat() FileFormatEnum`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *ConferenceRecordingMetadata) GetFileFormatOk() (*FileFormatEnum, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *ConferenceRecordingMetadata) SetFileFormat(v FileFormatEnum)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *ConferenceRecordingMetadata) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetStatus

`func (o *ConferenceRecordingMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConferenceRecordingMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConferenceRecordingMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConferenceRecordingMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMediaUrl

`func (o *ConferenceRecordingMetadata) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *ConferenceRecordingMetadata) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *ConferenceRecordingMetadata) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *ConferenceRecordingMetadata) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *ConferenceRecordingMetadata) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *ConferenceRecordingMetadata) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
### GetRecordingName

`func (o *ConferenceRecordingMetadata) GetRecordingName() string`

GetRecordingName returns the RecordingName field if non-nil, zero value otherwise.

### GetRecordingNameOk

`func (o *ConferenceRecordingMetadata) GetRecordingNameOk() (*string, bool)`

GetRecordingNameOk returns a tuple with the RecordingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingName

`func (o *ConferenceRecordingMetadata) SetRecordingName(v string)`

SetRecordingName sets RecordingName field to given value.

### HasRecordingName

`func (o *ConferenceRecordingMetadata) HasRecordingName() bool`

HasRecordingName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


