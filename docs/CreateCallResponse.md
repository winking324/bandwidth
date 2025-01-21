# CreateCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The id of the application associated with the &#x60;from&#x60; number. | 
**AccountId** | **string** | The bandwidth account ID associated with the call. | 
**CallId** | **string** | Programmable Voice API Call ID. | 
**To** | **string** | Recipient of the outgoing call. | 
**From** | **string** | Phone number that created the outbound call. | 
**EnqueuedTime** | Pointer to **NullableTime** | The time at which the call was accepted into the queue. | [optional] 
**CallUrl** | **string** | The URL to update this call&#39;s state. | 
**CallTimeout** | Pointer to **float64** | The timeout (in seconds) for the callee to answer the call after it starts ringing. | [optional] 
**CallbackTimeout** | Pointer to **float64** | This is the timeout (in seconds) to use when delivering webhooks for the call. | [optional] 
**Tag** | Pointer to **NullableString** | Custom tag value. | [optional] 
**AnswerMethod** | [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [default to POST]
**AnswerUrl** | **string** | URL to deliver the &#x60;answer&#x60; event webhook. | 
**AnswerFallbackMethod** | Pointer to [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [optional] [default to POST]
**AnswerFallbackUrl** | Pointer to **NullableString** | Fallback URL to deliver the &#x60;answer&#x60; event webhook. | [optional] 
**DisconnectMethod** | [**NullableCallbackMethodEnum**](CallbackMethodEnum.md) |  | [default to POST]
**DisconnectUrl** | Pointer to **NullableString** | URL to deliver the &#x60;disconnect&#x60; event webhook. | [optional] 
**Username** | Pointer to **NullableString** | Basic auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic auth password. | [optional] 
**FallbackUsername** | Pointer to **NullableString** | Basic auth username. | [optional] 
**FallbackPassword** | Pointer to **NullableString** | Basic auth password. | [optional] 
**Priority** | Pointer to **NullableInt32** | The priority of this call over other calls from your account. | [optional] 

## Methods

### NewCreateCallResponse

`func NewCreateCallResponse(applicationId string, accountId string, callId string, to string, from string, callUrl string, answerMethod NullableCallbackMethodEnum, answerUrl string, disconnectMethod NullableCallbackMethodEnum, ) *CreateCallResponse`

NewCreateCallResponse instantiates a new CreateCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCallResponseWithDefaults

`func NewCreateCallResponseWithDefaults() *CreateCallResponse`

NewCreateCallResponseWithDefaults instantiates a new CreateCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CreateCallResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CreateCallResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CreateCallResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetAccountId

`func (o *CreateCallResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateCallResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateCallResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCallId

`func (o *CreateCallResponse) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *CreateCallResponse) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *CreateCallResponse) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetTo

`func (o *CreateCallResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateCallResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateCallResponse) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *CreateCallResponse) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateCallResponse) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateCallResponse) SetFrom(v string)`

SetFrom sets From field to given value.


### GetEnqueuedTime

`func (o *CreateCallResponse) GetEnqueuedTime() time.Time`

GetEnqueuedTime returns the EnqueuedTime field if non-nil, zero value otherwise.

### GetEnqueuedTimeOk

`func (o *CreateCallResponse) GetEnqueuedTimeOk() (*time.Time, bool)`

GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedTime

`func (o *CreateCallResponse) SetEnqueuedTime(v time.Time)`

SetEnqueuedTime sets EnqueuedTime field to given value.

### HasEnqueuedTime

`func (o *CreateCallResponse) HasEnqueuedTime() bool`

HasEnqueuedTime returns a boolean if a field has been set.

### SetEnqueuedTimeNil

`func (o *CreateCallResponse) SetEnqueuedTimeNil(b bool)`

 SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil

### UnsetEnqueuedTime
`func (o *CreateCallResponse) UnsetEnqueuedTime()`

UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
### GetCallUrl

`func (o *CreateCallResponse) GetCallUrl() string`

GetCallUrl returns the CallUrl field if non-nil, zero value otherwise.

### GetCallUrlOk

`func (o *CreateCallResponse) GetCallUrlOk() (*string, bool)`

GetCallUrlOk returns a tuple with the CallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallUrl

`func (o *CreateCallResponse) SetCallUrl(v string)`

SetCallUrl sets CallUrl field to given value.


### GetCallTimeout

`func (o *CreateCallResponse) GetCallTimeout() float64`

GetCallTimeout returns the CallTimeout field if non-nil, zero value otherwise.

### GetCallTimeoutOk

`func (o *CreateCallResponse) GetCallTimeoutOk() (*float64, bool)`

GetCallTimeoutOk returns a tuple with the CallTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTimeout

`func (o *CreateCallResponse) SetCallTimeout(v float64)`

SetCallTimeout sets CallTimeout field to given value.

### HasCallTimeout

`func (o *CreateCallResponse) HasCallTimeout() bool`

HasCallTimeout returns a boolean if a field has been set.

### GetCallbackTimeout

`func (o *CreateCallResponse) GetCallbackTimeout() float64`

GetCallbackTimeout returns the CallbackTimeout field if non-nil, zero value otherwise.

### GetCallbackTimeoutOk

`func (o *CreateCallResponse) GetCallbackTimeoutOk() (*float64, bool)`

GetCallbackTimeoutOk returns a tuple with the CallbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackTimeout

`func (o *CreateCallResponse) SetCallbackTimeout(v float64)`

SetCallbackTimeout sets CallbackTimeout field to given value.

### HasCallbackTimeout

`func (o *CreateCallResponse) HasCallbackTimeout() bool`

HasCallbackTimeout returns a boolean if a field has been set.

### GetTag

`func (o *CreateCallResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateCallResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateCallResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateCallResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *CreateCallResponse) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *CreateCallResponse) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetAnswerMethod

`func (o *CreateCallResponse) GetAnswerMethod() CallbackMethodEnum`

GetAnswerMethod returns the AnswerMethod field if non-nil, zero value otherwise.

### GetAnswerMethodOk

`func (o *CreateCallResponse) GetAnswerMethodOk() (*CallbackMethodEnum, bool)`

GetAnswerMethodOk returns a tuple with the AnswerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerMethod

`func (o *CreateCallResponse) SetAnswerMethod(v CallbackMethodEnum)`

SetAnswerMethod sets AnswerMethod field to given value.


### SetAnswerMethodNil

`func (o *CreateCallResponse) SetAnswerMethodNil(b bool)`

 SetAnswerMethodNil sets the value for AnswerMethod to be an explicit nil

### UnsetAnswerMethod
`func (o *CreateCallResponse) UnsetAnswerMethod()`

UnsetAnswerMethod ensures that no value is present for AnswerMethod, not even an explicit nil
### GetAnswerUrl

`func (o *CreateCallResponse) GetAnswerUrl() string`

GetAnswerUrl returns the AnswerUrl field if non-nil, zero value otherwise.

### GetAnswerUrlOk

`func (o *CreateCallResponse) GetAnswerUrlOk() (*string, bool)`

GetAnswerUrlOk returns a tuple with the AnswerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerUrl

`func (o *CreateCallResponse) SetAnswerUrl(v string)`

SetAnswerUrl sets AnswerUrl field to given value.


### GetAnswerFallbackMethod

`func (o *CreateCallResponse) GetAnswerFallbackMethod() CallbackMethodEnum`

GetAnswerFallbackMethod returns the AnswerFallbackMethod field if non-nil, zero value otherwise.

### GetAnswerFallbackMethodOk

`func (o *CreateCallResponse) GetAnswerFallbackMethodOk() (*CallbackMethodEnum, bool)`

GetAnswerFallbackMethodOk returns a tuple with the AnswerFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFallbackMethod

`func (o *CreateCallResponse) SetAnswerFallbackMethod(v CallbackMethodEnum)`

SetAnswerFallbackMethod sets AnswerFallbackMethod field to given value.

### HasAnswerFallbackMethod

`func (o *CreateCallResponse) HasAnswerFallbackMethod() bool`

HasAnswerFallbackMethod returns a boolean if a field has been set.

### SetAnswerFallbackMethodNil

`func (o *CreateCallResponse) SetAnswerFallbackMethodNil(b bool)`

 SetAnswerFallbackMethodNil sets the value for AnswerFallbackMethod to be an explicit nil

### UnsetAnswerFallbackMethod
`func (o *CreateCallResponse) UnsetAnswerFallbackMethod()`

UnsetAnswerFallbackMethod ensures that no value is present for AnswerFallbackMethod, not even an explicit nil
### GetAnswerFallbackUrl

`func (o *CreateCallResponse) GetAnswerFallbackUrl() string`

GetAnswerFallbackUrl returns the AnswerFallbackUrl field if non-nil, zero value otherwise.

### GetAnswerFallbackUrlOk

`func (o *CreateCallResponse) GetAnswerFallbackUrlOk() (*string, bool)`

GetAnswerFallbackUrlOk returns a tuple with the AnswerFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFallbackUrl

`func (o *CreateCallResponse) SetAnswerFallbackUrl(v string)`

SetAnswerFallbackUrl sets AnswerFallbackUrl field to given value.

### HasAnswerFallbackUrl

`func (o *CreateCallResponse) HasAnswerFallbackUrl() bool`

HasAnswerFallbackUrl returns a boolean if a field has been set.

### SetAnswerFallbackUrlNil

`func (o *CreateCallResponse) SetAnswerFallbackUrlNil(b bool)`

 SetAnswerFallbackUrlNil sets the value for AnswerFallbackUrl to be an explicit nil

### UnsetAnswerFallbackUrl
`func (o *CreateCallResponse) UnsetAnswerFallbackUrl()`

UnsetAnswerFallbackUrl ensures that no value is present for AnswerFallbackUrl, not even an explicit nil
### GetDisconnectMethod

`func (o *CreateCallResponse) GetDisconnectMethod() CallbackMethodEnum`

GetDisconnectMethod returns the DisconnectMethod field if non-nil, zero value otherwise.

### GetDisconnectMethodOk

`func (o *CreateCallResponse) GetDisconnectMethodOk() (*CallbackMethodEnum, bool)`

GetDisconnectMethodOk returns a tuple with the DisconnectMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectMethod

`func (o *CreateCallResponse) SetDisconnectMethod(v CallbackMethodEnum)`

SetDisconnectMethod sets DisconnectMethod field to given value.


### SetDisconnectMethodNil

`func (o *CreateCallResponse) SetDisconnectMethodNil(b bool)`

 SetDisconnectMethodNil sets the value for DisconnectMethod to be an explicit nil

### UnsetDisconnectMethod
`func (o *CreateCallResponse) UnsetDisconnectMethod()`

UnsetDisconnectMethod ensures that no value is present for DisconnectMethod, not even an explicit nil
### GetDisconnectUrl

`func (o *CreateCallResponse) GetDisconnectUrl() string`

GetDisconnectUrl returns the DisconnectUrl field if non-nil, zero value otherwise.

### GetDisconnectUrlOk

`func (o *CreateCallResponse) GetDisconnectUrlOk() (*string, bool)`

GetDisconnectUrlOk returns a tuple with the DisconnectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectUrl

`func (o *CreateCallResponse) SetDisconnectUrl(v string)`

SetDisconnectUrl sets DisconnectUrl field to given value.

### HasDisconnectUrl

`func (o *CreateCallResponse) HasDisconnectUrl() bool`

HasDisconnectUrl returns a boolean if a field has been set.

### SetDisconnectUrlNil

`func (o *CreateCallResponse) SetDisconnectUrlNil(b bool)`

 SetDisconnectUrlNil sets the value for DisconnectUrl to be an explicit nil

### UnsetDisconnectUrl
`func (o *CreateCallResponse) UnsetDisconnectUrl()`

UnsetDisconnectUrl ensures that no value is present for DisconnectUrl, not even an explicit nil
### GetUsername

`func (o *CreateCallResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateCallResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateCallResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateCallResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateCallResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateCallResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateCallResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCallResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCallResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateCallResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateCallResponse) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateCallResponse) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetFallbackUsername

`func (o *CreateCallResponse) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *CreateCallResponse) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *CreateCallResponse) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *CreateCallResponse) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### SetFallbackUsernameNil

`func (o *CreateCallResponse) SetFallbackUsernameNil(b bool)`

 SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil

### UnsetFallbackUsername
`func (o *CreateCallResponse) UnsetFallbackUsername()`

UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
### GetFallbackPassword

`func (o *CreateCallResponse) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *CreateCallResponse) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *CreateCallResponse) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *CreateCallResponse) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### SetFallbackPasswordNil

`func (o *CreateCallResponse) SetFallbackPasswordNil(b bool)`

 SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil

### UnsetFallbackPassword
`func (o *CreateCallResponse) UnsetFallbackPassword()`

UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil
### GetPriority

`func (o *CreateCallResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateCallResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateCallResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateCallResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CreateCallResponse) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CreateCallResponse) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


