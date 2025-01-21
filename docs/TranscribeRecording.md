# TranscribeRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | The URL to send the [TranscriptionAvailable](/docs/voice/webhooks/transcriptionAvailable) event to. You should not include sensitive or personally-identifiable information in the callbackUrl field! Always use the proper username and password fields for authorization. | [optional] 
**CallbackMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**Username** | Pointer to **NullableString** | Basic auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic auth password. | [optional] 
**Tag** | Pointer to **NullableString** | (optional) The tag specified on call creation. If no tag was specified or it was previously cleared, this field will not be present. | [optional] 
**CallbackTimeout** | Pointer to **NullableFloat64** | This is the timeout (in seconds) to use when delivering the webhook to &#x60;callbackUrl&#x60;. Can be any numeric value (including decimals) between 1 and 25. | [optional] [default to 15]
**DetectLanguage** | Pointer to **NullableBool** | A boolean value to indicate that the recording may not be in English, and the transcription service will need to detect the dominant language the recording is in and transcribe accordingly. Current supported languages are English, French, and Spanish. | [optional] [default to false]

## Methods

### NewTranscribeRecording

`func NewTranscribeRecording() *TranscribeRecording`

NewTranscribeRecording instantiates a new TranscribeRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscribeRecordingWithDefaults

`func NewTranscribeRecordingWithDefaults() *TranscribeRecording`

NewTranscribeRecordingWithDefaults instantiates a new TranscribeRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *TranscribeRecording) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *TranscribeRecording) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *TranscribeRecording) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *TranscribeRecording) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCallbackMethod

`func (o *TranscribeRecording) GetCallbackMethod() CallbackMethodEnum`

GetCallbackMethod returns the CallbackMethod field if non-nil, zero value otherwise.

### GetCallbackMethodOk

`func (o *TranscribeRecording) GetCallbackMethodOk() (*CallbackMethodEnum, bool)`

GetCallbackMethodOk returns a tuple with the CallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackMethod

`func (o *TranscribeRecording) SetCallbackMethod(v CallbackMethodEnum)`

SetCallbackMethod sets CallbackMethod field to given value.

### HasCallbackMethod

`func (o *TranscribeRecording) HasCallbackMethod() bool`

HasCallbackMethod returns a boolean if a field has been set.

### SetCallbackMethodNil

`func (o *TranscribeRecording) SetCallbackMethodNil(b bool)`

 SetCallbackMethodNil sets the value for CallbackMethod to be an explicit nil

### UnsetCallbackMethod
`func (o *TranscribeRecording) UnsetCallbackMethod()`

UnsetCallbackMethod ensures that no value is present for CallbackMethod, not even an explicit nil
### GetUsername

`func (o *TranscribeRecording) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TranscribeRecording) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TranscribeRecording) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TranscribeRecording) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TranscribeRecording) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TranscribeRecording) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *TranscribeRecording) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TranscribeRecording) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TranscribeRecording) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TranscribeRecording) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TranscribeRecording) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TranscribeRecording) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTag

`func (o *TranscribeRecording) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TranscribeRecording) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TranscribeRecording) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TranscribeRecording) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *TranscribeRecording) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *TranscribeRecording) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetCallbackTimeout

`func (o *TranscribeRecording) GetCallbackTimeout() float64`

GetCallbackTimeout returns the CallbackTimeout field if non-nil, zero value otherwise.

### GetCallbackTimeoutOk

`func (o *TranscribeRecording) GetCallbackTimeoutOk() (*float64, bool)`

GetCallbackTimeoutOk returns a tuple with the CallbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackTimeout

`func (o *TranscribeRecording) SetCallbackTimeout(v float64)`

SetCallbackTimeout sets CallbackTimeout field to given value.

### HasCallbackTimeout

`func (o *TranscribeRecording) HasCallbackTimeout() bool`

HasCallbackTimeout returns a boolean if a field has been set.

### SetCallbackTimeoutNil

`func (o *TranscribeRecording) SetCallbackTimeoutNil(b bool)`

 SetCallbackTimeoutNil sets the value for CallbackTimeout to be an explicit nil

### UnsetCallbackTimeout
`func (o *TranscribeRecording) UnsetCallbackTimeout()`

UnsetCallbackTimeout ensures that no value is present for CallbackTimeout, not even an explicit nil
### GetDetectLanguage

`func (o *TranscribeRecording) GetDetectLanguage() bool`

GetDetectLanguage returns the DetectLanguage field if non-nil, zero value otherwise.

### GetDetectLanguageOk

`func (o *TranscribeRecording) GetDetectLanguageOk() (*bool, bool)`

GetDetectLanguageOk returns a tuple with the DetectLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectLanguage

`func (o *TranscribeRecording) SetDetectLanguage(v bool)`

SetDetectLanguage sets DetectLanguage field to given value.

### HasDetectLanguage

`func (o *TranscribeRecording) HasDetectLanguage() bool`

HasDetectLanguage returns a boolean if a field has been set.

### SetDetectLanguageNil

`func (o *TranscribeRecording) SetDetectLanguageNil(b bool)`

 SetDetectLanguageNil sets the value for DetectLanguage to be an explicit nil

### UnsetDetectLanguage
`func (o *TranscribeRecording) UnsetDetectLanguage()`

UnsetDetectLanguage ensures that no value is present for DetectLanguage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


