# MachineDetectionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**MachineDetectionModeEnum**](MachineDetectionModeEnum.md) |  | [optional] [default to ASYNC]
**DetectionTimeout** | Pointer to **NullableFloat64** | The timeout used for the whole operation, in seconds. If no result is determined in this period, a callback with a &#x60;timeout&#x60; result is sent. | [optional] [default to 15]
**SilenceTimeout** | Pointer to **NullableFloat64** | If no speech is detected in this period, a callback with a &#39;silence&#39; result is sent. | [optional] [default to 10]
**SpeechThreshold** | Pointer to **NullableFloat64** | When speech has ended and a result couldn&#39;t be determined based on the audio content itself, this value is used to determine if the speaker is a machine based on the speech duration. If the length of the speech detected is greater than or equal to this threshold, the result will be &#39;answering-machine&#39;. If the length of speech detected is below this threshold, the result will be &#39;human&#39;. | [optional] [default to 10]
**SpeechEndThreshold** | Pointer to **NullableFloat64** | Amount of silence (in seconds) before assuming the callee has finished speaking. | [optional] [default to 5]
**MachineSpeechEndThreshold** | Pointer to **NullableFloat64** | When an answering machine is detected, the amount of silence (in seconds) before assuming the message has finished playing.  If not provided it will default to the speechEndThreshold value. | [optional] 
**DelayResult** | Pointer to **NullableBool** | If set to &#39;true&#39; and if an answering machine is detected, the &#39;answering-machine&#39; callback will be delayed until the machine is done speaking, or an end of message tone is detected, or until the &#39;detectionTimeout&#39; is exceeded. If false, the &#39;answering-machine&#39; result is sent immediately. | [optional] [default to false]
**CallbackUrl** | Pointer to **NullableString** | The URL to send the &#39;machineDetectionComplete&#39; webhook when the detection is completed. Only for &#39;async&#39; mode. | [optional] 
**CallbackMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**Username** | Pointer to **NullableString** | Basic auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic auth password. | [optional] 
**FallbackUrl** | Pointer to **NullableString** | A fallback URL which, if provided, will be used to retry the machine detection complete webhook delivery in case &#x60;callbackUrl&#x60; fails to respond | [optional] 
**FallbackMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**FallbackUsername** | Pointer to **NullableString** | Basic auth username. | [optional] 
**FallbackPassword** | Pointer to **NullableString** | Basic auth password. | [optional] 

## Methods

### NewMachineDetectionConfiguration

`func NewMachineDetectionConfiguration() *MachineDetectionConfiguration`

NewMachineDetectionConfiguration instantiates a new MachineDetectionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineDetectionConfigurationWithDefaults

`func NewMachineDetectionConfigurationWithDefaults() *MachineDetectionConfiguration`

NewMachineDetectionConfigurationWithDefaults instantiates a new MachineDetectionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *MachineDetectionConfiguration) GetMode() MachineDetectionModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MachineDetectionConfiguration) GetModeOk() (*MachineDetectionModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MachineDetectionConfiguration) SetMode(v MachineDetectionModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MachineDetectionConfiguration) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDetectionTimeout

`func (o *MachineDetectionConfiguration) GetDetectionTimeout() float64`

GetDetectionTimeout returns the DetectionTimeout field if non-nil, zero value otherwise.

### GetDetectionTimeoutOk

`func (o *MachineDetectionConfiguration) GetDetectionTimeoutOk() (*float64, bool)`

GetDetectionTimeoutOk returns a tuple with the DetectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionTimeout

`func (o *MachineDetectionConfiguration) SetDetectionTimeout(v float64)`

SetDetectionTimeout sets DetectionTimeout field to given value.

### HasDetectionTimeout

`func (o *MachineDetectionConfiguration) HasDetectionTimeout() bool`

HasDetectionTimeout returns a boolean if a field has been set.

### SetDetectionTimeoutNil

`func (o *MachineDetectionConfiguration) SetDetectionTimeoutNil(b bool)`

 SetDetectionTimeoutNil sets the value for DetectionTimeout to be an explicit nil

### UnsetDetectionTimeout
`func (o *MachineDetectionConfiguration) UnsetDetectionTimeout()`

UnsetDetectionTimeout ensures that no value is present for DetectionTimeout, not even an explicit nil
### GetSilenceTimeout

`func (o *MachineDetectionConfiguration) GetSilenceTimeout() float64`

GetSilenceTimeout returns the SilenceTimeout field if non-nil, zero value otherwise.

### GetSilenceTimeoutOk

`func (o *MachineDetectionConfiguration) GetSilenceTimeoutOk() (*float64, bool)`

GetSilenceTimeoutOk returns a tuple with the SilenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilenceTimeout

`func (o *MachineDetectionConfiguration) SetSilenceTimeout(v float64)`

SetSilenceTimeout sets SilenceTimeout field to given value.

### HasSilenceTimeout

`func (o *MachineDetectionConfiguration) HasSilenceTimeout() bool`

HasSilenceTimeout returns a boolean if a field has been set.

### SetSilenceTimeoutNil

`func (o *MachineDetectionConfiguration) SetSilenceTimeoutNil(b bool)`

 SetSilenceTimeoutNil sets the value for SilenceTimeout to be an explicit nil

### UnsetSilenceTimeout
`func (o *MachineDetectionConfiguration) UnsetSilenceTimeout()`

UnsetSilenceTimeout ensures that no value is present for SilenceTimeout, not even an explicit nil
### GetSpeechThreshold

`func (o *MachineDetectionConfiguration) GetSpeechThreshold() float64`

GetSpeechThreshold returns the SpeechThreshold field if non-nil, zero value otherwise.

### GetSpeechThresholdOk

`func (o *MachineDetectionConfiguration) GetSpeechThresholdOk() (*float64, bool)`

GetSpeechThresholdOk returns a tuple with the SpeechThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechThreshold

`func (o *MachineDetectionConfiguration) SetSpeechThreshold(v float64)`

SetSpeechThreshold sets SpeechThreshold field to given value.

### HasSpeechThreshold

`func (o *MachineDetectionConfiguration) HasSpeechThreshold() bool`

HasSpeechThreshold returns a boolean if a field has been set.

### SetSpeechThresholdNil

`func (o *MachineDetectionConfiguration) SetSpeechThresholdNil(b bool)`

 SetSpeechThresholdNil sets the value for SpeechThreshold to be an explicit nil

### UnsetSpeechThreshold
`func (o *MachineDetectionConfiguration) UnsetSpeechThreshold()`

UnsetSpeechThreshold ensures that no value is present for SpeechThreshold, not even an explicit nil
### GetSpeechEndThreshold

`func (o *MachineDetectionConfiguration) GetSpeechEndThreshold() float64`

GetSpeechEndThreshold returns the SpeechEndThreshold field if non-nil, zero value otherwise.

### GetSpeechEndThresholdOk

`func (o *MachineDetectionConfiguration) GetSpeechEndThresholdOk() (*float64, bool)`

GetSpeechEndThresholdOk returns a tuple with the SpeechEndThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechEndThreshold

`func (o *MachineDetectionConfiguration) SetSpeechEndThreshold(v float64)`

SetSpeechEndThreshold sets SpeechEndThreshold field to given value.

### HasSpeechEndThreshold

`func (o *MachineDetectionConfiguration) HasSpeechEndThreshold() bool`

HasSpeechEndThreshold returns a boolean if a field has been set.

### SetSpeechEndThresholdNil

`func (o *MachineDetectionConfiguration) SetSpeechEndThresholdNil(b bool)`

 SetSpeechEndThresholdNil sets the value for SpeechEndThreshold to be an explicit nil

### UnsetSpeechEndThreshold
`func (o *MachineDetectionConfiguration) UnsetSpeechEndThreshold()`

UnsetSpeechEndThreshold ensures that no value is present for SpeechEndThreshold, not even an explicit nil
### GetMachineSpeechEndThreshold

`func (o *MachineDetectionConfiguration) GetMachineSpeechEndThreshold() float64`

GetMachineSpeechEndThreshold returns the MachineSpeechEndThreshold field if non-nil, zero value otherwise.

### GetMachineSpeechEndThresholdOk

`func (o *MachineDetectionConfiguration) GetMachineSpeechEndThresholdOk() (*float64, bool)`

GetMachineSpeechEndThresholdOk returns a tuple with the MachineSpeechEndThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSpeechEndThreshold

`func (o *MachineDetectionConfiguration) SetMachineSpeechEndThreshold(v float64)`

SetMachineSpeechEndThreshold sets MachineSpeechEndThreshold field to given value.

### HasMachineSpeechEndThreshold

`func (o *MachineDetectionConfiguration) HasMachineSpeechEndThreshold() bool`

HasMachineSpeechEndThreshold returns a boolean if a field has been set.

### SetMachineSpeechEndThresholdNil

`func (o *MachineDetectionConfiguration) SetMachineSpeechEndThresholdNil(b bool)`

 SetMachineSpeechEndThresholdNil sets the value for MachineSpeechEndThreshold to be an explicit nil

### UnsetMachineSpeechEndThreshold
`func (o *MachineDetectionConfiguration) UnsetMachineSpeechEndThreshold()`

UnsetMachineSpeechEndThreshold ensures that no value is present for MachineSpeechEndThreshold, not even an explicit nil
### GetDelayResult

`func (o *MachineDetectionConfiguration) GetDelayResult() bool`

GetDelayResult returns the DelayResult field if non-nil, zero value otherwise.

### GetDelayResultOk

`func (o *MachineDetectionConfiguration) GetDelayResultOk() (*bool, bool)`

GetDelayResultOk returns a tuple with the DelayResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayResult

`func (o *MachineDetectionConfiguration) SetDelayResult(v bool)`

SetDelayResult sets DelayResult field to given value.

### HasDelayResult

`func (o *MachineDetectionConfiguration) HasDelayResult() bool`

HasDelayResult returns a boolean if a field has been set.

### SetDelayResultNil

`func (o *MachineDetectionConfiguration) SetDelayResultNil(b bool)`

 SetDelayResultNil sets the value for DelayResult to be an explicit nil

### UnsetDelayResult
`func (o *MachineDetectionConfiguration) UnsetDelayResult()`

UnsetDelayResult ensures that no value is present for DelayResult, not even an explicit nil
### GetCallbackUrl

`func (o *MachineDetectionConfiguration) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *MachineDetectionConfiguration) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *MachineDetectionConfiguration) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *MachineDetectionConfiguration) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *MachineDetectionConfiguration) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *MachineDetectionConfiguration) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetCallbackMethod

`func (o *MachineDetectionConfiguration) GetCallbackMethod() CallbackMethodEnum`

GetCallbackMethod returns the CallbackMethod field if non-nil, zero value otherwise.

### GetCallbackMethodOk

`func (o *MachineDetectionConfiguration) GetCallbackMethodOk() (*CallbackMethodEnum, bool)`

GetCallbackMethodOk returns a tuple with the CallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackMethod

`func (o *MachineDetectionConfiguration) SetCallbackMethod(v CallbackMethodEnum)`

SetCallbackMethod sets CallbackMethod field to given value.

### HasCallbackMethod

`func (o *MachineDetectionConfiguration) HasCallbackMethod() bool`

HasCallbackMethod returns a boolean if a field has been set.

### SetCallbackMethodNil

`func (o *MachineDetectionConfiguration) SetCallbackMethodNil(b bool)`

 SetCallbackMethodNil sets the value for CallbackMethod to be an explicit nil

### UnsetCallbackMethod
`func (o *MachineDetectionConfiguration) UnsetCallbackMethod()`

UnsetCallbackMethod ensures that no value is present for CallbackMethod, not even an explicit nil
### GetUsername

`func (o *MachineDetectionConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MachineDetectionConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MachineDetectionConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MachineDetectionConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *MachineDetectionConfiguration) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *MachineDetectionConfiguration) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *MachineDetectionConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MachineDetectionConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MachineDetectionConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MachineDetectionConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MachineDetectionConfiguration) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MachineDetectionConfiguration) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetFallbackUrl

`func (o *MachineDetectionConfiguration) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *MachineDetectionConfiguration) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *MachineDetectionConfiguration) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *MachineDetectionConfiguration) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### SetFallbackUrlNil

`func (o *MachineDetectionConfiguration) SetFallbackUrlNil(b bool)`

 SetFallbackUrlNil sets the value for FallbackUrl to be an explicit nil

### UnsetFallbackUrl
`func (o *MachineDetectionConfiguration) UnsetFallbackUrl()`

UnsetFallbackUrl ensures that no value is present for FallbackUrl, not even an explicit nil
### GetFallbackMethod

`func (o *MachineDetectionConfiguration) GetFallbackMethod() CallbackMethodEnum`

GetFallbackMethod returns the FallbackMethod field if non-nil, zero value otherwise.

### GetFallbackMethodOk

`func (o *MachineDetectionConfiguration) GetFallbackMethodOk() (*CallbackMethodEnum, bool)`

GetFallbackMethodOk returns a tuple with the FallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMethod

`func (o *MachineDetectionConfiguration) SetFallbackMethod(v CallbackMethodEnum)`

SetFallbackMethod sets FallbackMethod field to given value.

### HasFallbackMethod

`func (o *MachineDetectionConfiguration) HasFallbackMethod() bool`

HasFallbackMethod returns a boolean if a field has been set.

### SetFallbackMethodNil

`func (o *MachineDetectionConfiguration) SetFallbackMethodNil(b bool)`

 SetFallbackMethodNil sets the value for FallbackMethod to be an explicit nil

### UnsetFallbackMethod
`func (o *MachineDetectionConfiguration) UnsetFallbackMethod()`

UnsetFallbackMethod ensures that no value is present for FallbackMethod, not even an explicit nil
### GetFallbackUsername

`func (o *MachineDetectionConfiguration) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *MachineDetectionConfiguration) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *MachineDetectionConfiguration) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *MachineDetectionConfiguration) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### SetFallbackUsernameNil

`func (o *MachineDetectionConfiguration) SetFallbackUsernameNil(b bool)`

 SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil

### UnsetFallbackUsername
`func (o *MachineDetectionConfiguration) UnsetFallbackUsername()`

UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
### GetFallbackPassword

`func (o *MachineDetectionConfiguration) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *MachineDetectionConfiguration) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *MachineDetectionConfiguration) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *MachineDetectionConfiguration) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### SetFallbackPasswordNil

`func (o *MachineDetectionConfiguration) SetFallbackPasswordNil(b bool)`

 SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil

### UnsetFallbackPassword
`func (o *MachineDetectionConfiguration) UnsetFallbackPassword()`

UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


