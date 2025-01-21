# CreateCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The destination to call (must be an E.164 formatted number (e.g. &#x60;+15555551212&#x60;) or a SIP URI (e.g. &#x60;sip:user@server.example&#x60;)). | 
**From** | **string** | A Bandwidth phone number on your account the call should come from (must be in E.164 format, like &#x60;+15555551212&#x60;) even if &#x60;privacy&#x60; is set to true. | 
**Privacy** | Pointer to **NullableBool** | Hide the calling number. The &#x60;displayName&#x60; field can be used to customize the displayed name. | [optional] 
**DisplayName** | Pointer to **NullableString** | The caller display name to use when the call is created.  May not exceed 256 characters nor contain control characters such as new lines. If &#x60;privacy&#x60; is true, only the following values are valid: &#x60;Restricted&#x60;, &#x60;Anonymous&#x60;, &#x60;Private&#x60;, or &#x60;Unavailable&#x60;. | [optional] 
**Uui** | Pointer to **NullableString** | A comma-separated list of &#39;User-To-User&#39; headers to be sent in the INVITE when calling a SIP URI. Each value must end with an &#39;encoding&#39; parameter as described in &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc7433&#39;&gt;RFC 7433&lt;/a&gt;. Only &#39;jwt&#39;, &#39;base64&#39; and &#39;hex&#39; encodings are allowed. The entire value cannot exceed 350 characters, including parameters and separators. | [optional] 
**ApplicationId** | **string** | The id of the application associated with the &#x60;from&#x60; number. | 
**AnswerUrl** | **string** | The full URL to send the &lt;a href&#x3D;&#39;/docs/voice/webhooks/answer&#39;&gt;Answer&lt;/a&gt; event to when the called party answers. This endpoint should return the first &lt;a href&#x3D;&#39;/docs/voice/bxml&#39;&gt;BXML document&lt;/a&gt; to be executed in the call.  Must use &#x60;https&#x60; if specifying &#x60;username&#x60; and &#x60;password&#x60;. | 
**AnswerMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**Username** | Pointer to **NullableString** | Basic auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic auth password. | [optional] 
**AnswerFallbackUrl** | Pointer to **NullableString** | A fallback url which, if provided, will be used to retry the &#x60;answer&#x60; webhook delivery in case &#x60;answerUrl&#x60; fails to respond  Must use &#x60;https&#x60; if specifying &#x60;fallbackUsername&#x60; and &#x60;fallbackPassword&#x60;. | [optional] 
**AnswerFallbackMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**FallbackUsername** | Pointer to **NullableString** | Basic auth username. | [optional] 
**FallbackPassword** | Pointer to **NullableString** | Basic auth password. | [optional] 
**DisconnectUrl** | Pointer to **NullableString** | The URL to send the &lt;a href&#x3D;&#39;/docs/voice/webhooks/disconnect&#39;&gt;Disconnect&lt;/a&gt; event to when the call ends. This event does not expect a BXML response. | [optional] 
**DisconnectMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**CallTimeout** | Pointer to **NullableFloat64** | The timeout (in seconds) for the callee to answer the call after it starts ringing. If the call does not start ringing within 30s, the call will be cancelled regardless of this value.  Can be any numeric value (including decimals) between 1 and 300. | [optional] [default to 30]
**CallbackTimeout** | Pointer to **NullableFloat64** | This is the timeout (in seconds) to use when delivering webhooks for the call. Can be any numeric value (including decimals) between 1 and 25. | [optional] [default to 15]
**MachineDetection** | Pointer to [**MachineDetectionConfiguration**](MachineDetectionConfiguration.md) |  | [optional] 
**Priority** | Pointer to **NullableInt32** | The priority of this call over other calls from your account. For example, if during a call your application needs to place a new call and bridge it with the current call, you might want to create the call with priority 1 so that it will be the next call picked off your queue, ahead of other less time sensitive calls. A lower value means higher priority, so a priority 1 call takes precedence over a priority 2 call. | [optional] [default to 5]
**Tag** | Pointer to **NullableString** | A custom string that will be sent with all webhooks for this call unless overwritten by a future &lt;a href&#x3D;&#39;/docs/voice/bxml/tag&#39;&gt;&#x60;&lt;Tag&gt;&#x60;&lt;/a&gt; verb or &#x60;tag&#x60; attribute on another verb, or cleared.  May be cleared by setting &#x60;tag&#x3D;\&quot;\&quot;&#x60;  Max length 256 characters. | [optional] 

## Methods

### NewCreateCall

`func NewCreateCall(to string, from string, applicationId string, answerUrl string, ) *CreateCall`

NewCreateCall instantiates a new CreateCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCallWithDefaults

`func NewCreateCallWithDefaults() *CreateCall`

NewCreateCallWithDefaults instantiates a new CreateCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *CreateCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateCall) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *CreateCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateCall) SetFrom(v string)`

SetFrom sets From field to given value.


### GetPrivacy

`func (o *CreateCall) GetPrivacy() bool`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *CreateCall) GetPrivacyOk() (*bool, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *CreateCall) SetPrivacy(v bool)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *CreateCall) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### SetPrivacyNil

`func (o *CreateCall) SetPrivacyNil(b bool)`

 SetPrivacyNil sets the value for Privacy to be an explicit nil

### UnsetPrivacy
`func (o *CreateCall) UnsetPrivacy()`

UnsetPrivacy ensures that no value is present for Privacy, not even an explicit nil
### GetDisplayName

`func (o *CreateCall) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateCall) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateCall) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateCall) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateCall) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateCall) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUui

`func (o *CreateCall) GetUui() string`

GetUui returns the Uui field if non-nil, zero value otherwise.

### GetUuiOk

`func (o *CreateCall) GetUuiOk() (*string, bool)`

GetUuiOk returns a tuple with the Uui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUui

`func (o *CreateCall) SetUui(v string)`

SetUui sets Uui field to given value.

### HasUui

`func (o *CreateCall) HasUui() bool`

HasUui returns a boolean if a field has been set.

### SetUuiNil

`func (o *CreateCall) SetUuiNil(b bool)`

 SetUuiNil sets the value for Uui to be an explicit nil

### UnsetUui
`func (o *CreateCall) UnsetUui()`

UnsetUui ensures that no value is present for Uui, not even an explicit nil
### GetApplicationId

`func (o *CreateCall) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CreateCall) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CreateCall) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetAnswerUrl

`func (o *CreateCall) GetAnswerUrl() string`

GetAnswerUrl returns the AnswerUrl field if non-nil, zero value otherwise.

### GetAnswerUrlOk

`func (o *CreateCall) GetAnswerUrlOk() (*string, bool)`

GetAnswerUrlOk returns a tuple with the AnswerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerUrl

`func (o *CreateCall) SetAnswerUrl(v string)`

SetAnswerUrl sets AnswerUrl field to given value.


### GetAnswerMethod

`func (o *CreateCall) GetAnswerMethod() CallbackMethodEnum`

GetAnswerMethod returns the AnswerMethod field if non-nil, zero value otherwise.

### GetAnswerMethodOk

`func (o *CreateCall) GetAnswerMethodOk() (*CallbackMethodEnum, bool)`

GetAnswerMethodOk returns a tuple with the AnswerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerMethod

`func (o *CreateCall) SetAnswerMethod(v CallbackMethodEnum)`

SetAnswerMethod sets AnswerMethod field to given value.

### HasAnswerMethod

`func (o *CreateCall) HasAnswerMethod() bool`

HasAnswerMethod returns a boolean if a field has been set.

### SetAnswerMethodNil

`func (o *CreateCall) SetAnswerMethodNil(b bool)`

 SetAnswerMethodNil sets the value for AnswerMethod to be an explicit nil

### UnsetAnswerMethod
`func (o *CreateCall) UnsetAnswerMethod()`

UnsetAnswerMethod ensures that no value is present for AnswerMethod, not even an explicit nil
### GetUsername

`func (o *CreateCall) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateCall) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateCall) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateCall) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateCall) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateCall) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateCall) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCall) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCall) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateCall) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateCall) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateCall) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetAnswerFallbackUrl

`func (o *CreateCall) GetAnswerFallbackUrl() string`

GetAnswerFallbackUrl returns the AnswerFallbackUrl field if non-nil, zero value otherwise.

### GetAnswerFallbackUrlOk

`func (o *CreateCall) GetAnswerFallbackUrlOk() (*string, bool)`

GetAnswerFallbackUrlOk returns a tuple with the AnswerFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFallbackUrl

`func (o *CreateCall) SetAnswerFallbackUrl(v string)`

SetAnswerFallbackUrl sets AnswerFallbackUrl field to given value.

### HasAnswerFallbackUrl

`func (o *CreateCall) HasAnswerFallbackUrl() bool`

HasAnswerFallbackUrl returns a boolean if a field has been set.

### SetAnswerFallbackUrlNil

`func (o *CreateCall) SetAnswerFallbackUrlNil(b bool)`

 SetAnswerFallbackUrlNil sets the value for AnswerFallbackUrl to be an explicit nil

### UnsetAnswerFallbackUrl
`func (o *CreateCall) UnsetAnswerFallbackUrl()`

UnsetAnswerFallbackUrl ensures that no value is present for AnswerFallbackUrl, not even an explicit nil
### GetAnswerFallbackMethod

`func (o *CreateCall) GetAnswerFallbackMethod() CallbackMethodEnum`

GetAnswerFallbackMethod returns the AnswerFallbackMethod field if non-nil, zero value otherwise.

### GetAnswerFallbackMethodOk

`func (o *CreateCall) GetAnswerFallbackMethodOk() (*CallbackMethodEnum, bool)`

GetAnswerFallbackMethodOk returns a tuple with the AnswerFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFallbackMethod

`func (o *CreateCall) SetAnswerFallbackMethod(v CallbackMethodEnum)`

SetAnswerFallbackMethod sets AnswerFallbackMethod field to given value.

### HasAnswerFallbackMethod

`func (o *CreateCall) HasAnswerFallbackMethod() bool`

HasAnswerFallbackMethod returns a boolean if a field has been set.

### SetAnswerFallbackMethodNil

`func (o *CreateCall) SetAnswerFallbackMethodNil(b bool)`

 SetAnswerFallbackMethodNil sets the value for AnswerFallbackMethod to be an explicit nil

### UnsetAnswerFallbackMethod
`func (o *CreateCall) UnsetAnswerFallbackMethod()`

UnsetAnswerFallbackMethod ensures that no value is present for AnswerFallbackMethod, not even an explicit nil
### GetFallbackUsername

`func (o *CreateCall) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *CreateCall) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *CreateCall) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *CreateCall) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### SetFallbackUsernameNil

`func (o *CreateCall) SetFallbackUsernameNil(b bool)`

 SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil

### UnsetFallbackUsername
`func (o *CreateCall) UnsetFallbackUsername()`

UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
### GetFallbackPassword

`func (o *CreateCall) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *CreateCall) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *CreateCall) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *CreateCall) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### SetFallbackPasswordNil

`func (o *CreateCall) SetFallbackPasswordNil(b bool)`

 SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil

### UnsetFallbackPassword
`func (o *CreateCall) UnsetFallbackPassword()`

UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil
### GetDisconnectUrl

`func (o *CreateCall) GetDisconnectUrl() string`

GetDisconnectUrl returns the DisconnectUrl field if non-nil, zero value otherwise.

### GetDisconnectUrlOk

`func (o *CreateCall) GetDisconnectUrlOk() (*string, bool)`

GetDisconnectUrlOk returns a tuple with the DisconnectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectUrl

`func (o *CreateCall) SetDisconnectUrl(v string)`

SetDisconnectUrl sets DisconnectUrl field to given value.

### HasDisconnectUrl

`func (o *CreateCall) HasDisconnectUrl() bool`

HasDisconnectUrl returns a boolean if a field has been set.

### SetDisconnectUrlNil

`func (o *CreateCall) SetDisconnectUrlNil(b bool)`

 SetDisconnectUrlNil sets the value for DisconnectUrl to be an explicit nil

### UnsetDisconnectUrl
`func (o *CreateCall) UnsetDisconnectUrl()`

UnsetDisconnectUrl ensures that no value is present for DisconnectUrl, not even an explicit nil
### GetDisconnectMethod

`func (o *CreateCall) GetDisconnectMethod() CallbackMethodEnum`

GetDisconnectMethod returns the DisconnectMethod field if non-nil, zero value otherwise.

### GetDisconnectMethodOk

`func (o *CreateCall) GetDisconnectMethodOk() (*CallbackMethodEnum, bool)`

GetDisconnectMethodOk returns a tuple with the DisconnectMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectMethod

`func (o *CreateCall) SetDisconnectMethod(v CallbackMethodEnum)`

SetDisconnectMethod sets DisconnectMethod field to given value.

### HasDisconnectMethod

`func (o *CreateCall) HasDisconnectMethod() bool`

HasDisconnectMethod returns a boolean if a field has been set.

### SetDisconnectMethodNil

`func (o *CreateCall) SetDisconnectMethodNil(b bool)`

 SetDisconnectMethodNil sets the value for DisconnectMethod to be an explicit nil

### UnsetDisconnectMethod
`func (o *CreateCall) UnsetDisconnectMethod()`

UnsetDisconnectMethod ensures that no value is present for DisconnectMethod, not even an explicit nil
### GetCallTimeout

`func (o *CreateCall) GetCallTimeout() float64`

GetCallTimeout returns the CallTimeout field if non-nil, zero value otherwise.

### GetCallTimeoutOk

`func (o *CreateCall) GetCallTimeoutOk() (*float64, bool)`

GetCallTimeoutOk returns a tuple with the CallTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTimeout

`func (o *CreateCall) SetCallTimeout(v float64)`

SetCallTimeout sets CallTimeout field to given value.

### HasCallTimeout

`func (o *CreateCall) HasCallTimeout() bool`

HasCallTimeout returns a boolean if a field has been set.

### SetCallTimeoutNil

`func (o *CreateCall) SetCallTimeoutNil(b bool)`

 SetCallTimeoutNil sets the value for CallTimeout to be an explicit nil

### UnsetCallTimeout
`func (o *CreateCall) UnsetCallTimeout()`

UnsetCallTimeout ensures that no value is present for CallTimeout, not even an explicit nil
### GetCallbackTimeout

`func (o *CreateCall) GetCallbackTimeout() float64`

GetCallbackTimeout returns the CallbackTimeout field if non-nil, zero value otherwise.

### GetCallbackTimeoutOk

`func (o *CreateCall) GetCallbackTimeoutOk() (*float64, bool)`

GetCallbackTimeoutOk returns a tuple with the CallbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackTimeout

`func (o *CreateCall) SetCallbackTimeout(v float64)`

SetCallbackTimeout sets CallbackTimeout field to given value.

### HasCallbackTimeout

`func (o *CreateCall) HasCallbackTimeout() bool`

HasCallbackTimeout returns a boolean if a field has been set.

### SetCallbackTimeoutNil

`func (o *CreateCall) SetCallbackTimeoutNil(b bool)`

 SetCallbackTimeoutNil sets the value for CallbackTimeout to be an explicit nil

### UnsetCallbackTimeout
`func (o *CreateCall) UnsetCallbackTimeout()`

UnsetCallbackTimeout ensures that no value is present for CallbackTimeout, not even an explicit nil
### GetMachineDetection

`func (o *CreateCall) GetMachineDetection() MachineDetectionConfiguration`

GetMachineDetection returns the MachineDetection field if non-nil, zero value otherwise.

### GetMachineDetectionOk

`func (o *CreateCall) GetMachineDetectionOk() (*MachineDetectionConfiguration, bool)`

GetMachineDetectionOk returns a tuple with the MachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetection

`func (o *CreateCall) SetMachineDetection(v MachineDetectionConfiguration)`

SetMachineDetection sets MachineDetection field to given value.

### HasMachineDetection

`func (o *CreateCall) HasMachineDetection() bool`

HasMachineDetection returns a boolean if a field has been set.

### GetPriority

`func (o *CreateCall) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateCall) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateCall) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateCall) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CreateCall) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CreateCall) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetTag

`func (o *CreateCall) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateCall) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateCall) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateCall) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *CreateCall) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *CreateCall) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


