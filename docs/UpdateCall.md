# UpdateCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**NullableCallStateEnum**](CallStateEnum.md) |  | [optional] [default to ACTIVE]
**RedirectUrl** | Pointer to **NullableString** | The URL to send the [Redirect](/docs/voice/bxml/redirect) event to which will provide new BXML.  Required if &#x60;state&#x60; is &#x60;active&#x60;.  Not allowed if &#x60;state&#x60; is &#x60;completed&#x60;. | [optional] 
**RedirectMethod** | Pointer to [**NullableRedirectMethodEnum**](RedirectMethodEnum.md) |  | [optional] [default to POST]
**Username** | Pointer to **NullableString** | Basic auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic auth password. | [optional] 
**RedirectFallbackUrl** | Pointer to **NullableString** | A fallback url which, if provided, will be used to retry the redirect callback delivery in case &#x60;redirectUrl&#x60; fails to respond. | [optional] 
**RedirectFallbackMethod** | Pointer to [**NullableRedirectMethodEnum**](RedirectMethodEnum.md) |  | [optional] [default to POST]
**FallbackUsername** | Pointer to **NullableString** | Basic auth username. | [optional] 
**FallbackPassword** | Pointer to **NullableString** | Basic auth password. | [optional] 
**Tag** | Pointer to **NullableString** | A custom string that will be sent with this and all future callbacks unless overwritten by a future &#x60;tag&#x60; attribute or [&#x60;&lt;Tag&gt;&#x60;](/docs/voice/bxml/tag) verb, or cleared.  May be cleared by setting &#x60;tag&#x3D;\&quot;\&quot;&#x60;.  Max length 256 characters.  Not allowed if &#x60;state&#x60; is &#x60;completed&#x60;. | [optional] 

## Methods

### NewUpdateCall

`func NewUpdateCall() *UpdateCall`

NewUpdateCall instantiates a new UpdateCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallWithDefaults

`func NewUpdateCallWithDefaults() *UpdateCall`

NewUpdateCallWithDefaults instantiates a new UpdateCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *UpdateCall) GetState() CallStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCall) GetStateOk() (*CallStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateCall) SetState(v CallStateEnum)`

SetState sets State field to given value.

### HasState

`func (o *UpdateCall) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UpdateCall) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UpdateCall) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetRedirectUrl

`func (o *UpdateCall) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateCall) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateCall) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateCall) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *UpdateCall) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *UpdateCall) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRedirectMethod

`func (o *UpdateCall) GetRedirectMethod() RedirectMethodEnum`

GetRedirectMethod returns the RedirectMethod field if non-nil, zero value otherwise.

### GetRedirectMethodOk

`func (o *UpdateCall) GetRedirectMethodOk() (*RedirectMethodEnum, bool)`

GetRedirectMethodOk returns a tuple with the RedirectMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectMethod

`func (o *UpdateCall) SetRedirectMethod(v RedirectMethodEnum)`

SetRedirectMethod sets RedirectMethod field to given value.

### HasRedirectMethod

`func (o *UpdateCall) HasRedirectMethod() bool`

HasRedirectMethod returns a boolean if a field has been set.

### SetRedirectMethodNil

`func (o *UpdateCall) SetRedirectMethodNil(b bool)`

 SetRedirectMethodNil sets the value for RedirectMethod to be an explicit nil

### UnsetRedirectMethod
`func (o *UpdateCall) UnsetRedirectMethod()`

UnsetRedirectMethod ensures that no value is present for RedirectMethod, not even an explicit nil
### GetUsername

`func (o *UpdateCall) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCall) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCall) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateCall) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateCall) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateCall) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateCall) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCall) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCall) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateCall) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateCall) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateCall) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRedirectFallbackUrl

`func (o *UpdateCall) GetRedirectFallbackUrl() string`

GetRedirectFallbackUrl returns the RedirectFallbackUrl field if non-nil, zero value otherwise.

### GetRedirectFallbackUrlOk

`func (o *UpdateCall) GetRedirectFallbackUrlOk() (*string, bool)`

GetRedirectFallbackUrlOk returns a tuple with the RedirectFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFallbackUrl

`func (o *UpdateCall) SetRedirectFallbackUrl(v string)`

SetRedirectFallbackUrl sets RedirectFallbackUrl field to given value.

### HasRedirectFallbackUrl

`func (o *UpdateCall) HasRedirectFallbackUrl() bool`

HasRedirectFallbackUrl returns a boolean if a field has been set.

### SetRedirectFallbackUrlNil

`func (o *UpdateCall) SetRedirectFallbackUrlNil(b bool)`

 SetRedirectFallbackUrlNil sets the value for RedirectFallbackUrl to be an explicit nil

### UnsetRedirectFallbackUrl
`func (o *UpdateCall) UnsetRedirectFallbackUrl()`

UnsetRedirectFallbackUrl ensures that no value is present for RedirectFallbackUrl, not even an explicit nil
### GetRedirectFallbackMethod

`func (o *UpdateCall) GetRedirectFallbackMethod() RedirectMethodEnum`

GetRedirectFallbackMethod returns the RedirectFallbackMethod field if non-nil, zero value otherwise.

### GetRedirectFallbackMethodOk

`func (o *UpdateCall) GetRedirectFallbackMethodOk() (*RedirectMethodEnum, bool)`

GetRedirectFallbackMethodOk returns a tuple with the RedirectFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFallbackMethod

`func (o *UpdateCall) SetRedirectFallbackMethod(v RedirectMethodEnum)`

SetRedirectFallbackMethod sets RedirectFallbackMethod field to given value.

### HasRedirectFallbackMethod

`func (o *UpdateCall) HasRedirectFallbackMethod() bool`

HasRedirectFallbackMethod returns a boolean if a field has been set.

### SetRedirectFallbackMethodNil

`func (o *UpdateCall) SetRedirectFallbackMethodNil(b bool)`

 SetRedirectFallbackMethodNil sets the value for RedirectFallbackMethod to be an explicit nil

### UnsetRedirectFallbackMethod
`func (o *UpdateCall) UnsetRedirectFallbackMethod()`

UnsetRedirectFallbackMethod ensures that no value is present for RedirectFallbackMethod, not even an explicit nil
### GetFallbackUsername

`func (o *UpdateCall) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *UpdateCall) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *UpdateCall) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *UpdateCall) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### SetFallbackUsernameNil

`func (o *UpdateCall) SetFallbackUsernameNil(b bool)`

 SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil

### UnsetFallbackUsername
`func (o *UpdateCall) UnsetFallbackUsername()`

UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
### GetFallbackPassword

`func (o *UpdateCall) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *UpdateCall) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *UpdateCall) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *UpdateCall) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### SetFallbackPasswordNil

`func (o *UpdateCall) SetFallbackPasswordNil(b bool)`

 SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil

### UnsetFallbackPassword
`func (o *UpdateCall) UnsetFallbackPassword()`

UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil
### GetTag

`func (o *UpdateCall) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UpdateCall) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UpdateCall) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UpdateCall) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *UpdateCall) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *UpdateCall) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


