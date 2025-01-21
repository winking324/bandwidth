# UpdateConference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableConferenceStateEnum**](ConferenceStateEnum.md) |  | [optional] [default to ACTIVE]
**RedirectUrl** | Pointer to **NullableString** | The URL to send the [conferenceRedirect](/docs/voice/webhooks/conferenceRedirect) event which will provide new BXML. Not allowed if &#x60;state&#x60; is &#x60;completed&#x60;, but required if &#x60;state&#x60; is &#x60;active&#x60;. | [optional] 
**RedirectMethod** | Pointer to [**NullableRedirectMethodEnum**](RedirectMethodEnum.md) |  | [optional] [default to POST]
**Username** | Pointer to **NullableString** | Basic auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic auth password. | [optional] 
**RedirectFallbackUrl** | Pointer to **NullableString** | A fallback url which, if provided, will be used to retry the &#x60;conferenceRedirect&#x60; webhook delivery in case &#x60;redirectUrl&#x60; fails to respond.  Not allowed if &#x60;state&#x60; is &#x60;completed&#x60;. | [optional] 
**RedirectFallbackMethod** | Pointer to [**NullableRedirectMethodEnum**](RedirectMethodEnum.md) |  | [optional] [default to POST]
**FallbackUsername** | Pointer to **NullableString** | Basic auth username. | [optional] 
**FallbackPassword** | Pointer to **NullableString** | Basic auth password. | [optional] 

## Methods

### NewUpdateConference

`func NewUpdateConference() *UpdateConference`

NewUpdateConference instantiates a new UpdateConference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConferenceWithDefaults

`func NewUpdateConferenceWithDefaults() *UpdateConference`

NewUpdateConferenceWithDefaults instantiates a new UpdateConference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateConference) GetStatus() ConferenceStateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateConference) GetStatusOk() (*ConferenceStateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateConference) SetStatus(v ConferenceStateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateConference) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateConference) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateConference) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRedirectUrl

`func (o *UpdateConference) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateConference) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateConference) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateConference) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *UpdateConference) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *UpdateConference) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRedirectMethod

`func (o *UpdateConference) GetRedirectMethod() RedirectMethodEnum`

GetRedirectMethod returns the RedirectMethod field if non-nil, zero value otherwise.

### GetRedirectMethodOk

`func (o *UpdateConference) GetRedirectMethodOk() (*RedirectMethodEnum, bool)`

GetRedirectMethodOk returns a tuple with the RedirectMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectMethod

`func (o *UpdateConference) SetRedirectMethod(v RedirectMethodEnum)`

SetRedirectMethod sets RedirectMethod field to given value.

### HasRedirectMethod

`func (o *UpdateConference) HasRedirectMethod() bool`

HasRedirectMethod returns a boolean if a field has been set.

### SetRedirectMethodNil

`func (o *UpdateConference) SetRedirectMethodNil(b bool)`

 SetRedirectMethodNil sets the value for RedirectMethod to be an explicit nil

### UnsetRedirectMethod
`func (o *UpdateConference) UnsetRedirectMethod()`

UnsetRedirectMethod ensures that no value is present for RedirectMethod, not even an explicit nil
### GetUsername

`func (o *UpdateConference) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateConference) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateConference) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateConference) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateConference) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateConference) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateConference) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateConference) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateConference) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateConference) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateConference) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateConference) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRedirectFallbackUrl

`func (o *UpdateConference) GetRedirectFallbackUrl() string`

GetRedirectFallbackUrl returns the RedirectFallbackUrl field if non-nil, zero value otherwise.

### GetRedirectFallbackUrlOk

`func (o *UpdateConference) GetRedirectFallbackUrlOk() (*string, bool)`

GetRedirectFallbackUrlOk returns a tuple with the RedirectFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFallbackUrl

`func (o *UpdateConference) SetRedirectFallbackUrl(v string)`

SetRedirectFallbackUrl sets RedirectFallbackUrl field to given value.

### HasRedirectFallbackUrl

`func (o *UpdateConference) HasRedirectFallbackUrl() bool`

HasRedirectFallbackUrl returns a boolean if a field has been set.

### SetRedirectFallbackUrlNil

`func (o *UpdateConference) SetRedirectFallbackUrlNil(b bool)`

 SetRedirectFallbackUrlNil sets the value for RedirectFallbackUrl to be an explicit nil

### UnsetRedirectFallbackUrl
`func (o *UpdateConference) UnsetRedirectFallbackUrl()`

UnsetRedirectFallbackUrl ensures that no value is present for RedirectFallbackUrl, not even an explicit nil
### GetRedirectFallbackMethod

`func (o *UpdateConference) GetRedirectFallbackMethod() RedirectMethodEnum`

GetRedirectFallbackMethod returns the RedirectFallbackMethod field if non-nil, zero value otherwise.

### GetRedirectFallbackMethodOk

`func (o *UpdateConference) GetRedirectFallbackMethodOk() (*RedirectMethodEnum, bool)`

GetRedirectFallbackMethodOk returns a tuple with the RedirectFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFallbackMethod

`func (o *UpdateConference) SetRedirectFallbackMethod(v RedirectMethodEnum)`

SetRedirectFallbackMethod sets RedirectFallbackMethod field to given value.

### HasRedirectFallbackMethod

`func (o *UpdateConference) HasRedirectFallbackMethod() bool`

HasRedirectFallbackMethod returns a boolean if a field has been set.

### SetRedirectFallbackMethodNil

`func (o *UpdateConference) SetRedirectFallbackMethodNil(b bool)`

 SetRedirectFallbackMethodNil sets the value for RedirectFallbackMethod to be an explicit nil

### UnsetRedirectFallbackMethod
`func (o *UpdateConference) UnsetRedirectFallbackMethod()`

UnsetRedirectFallbackMethod ensures that no value is present for RedirectFallbackMethod, not even an explicit nil
### GetFallbackUsername

`func (o *UpdateConference) GetFallbackUsername() string`

GetFallbackUsername returns the FallbackUsername field if non-nil, zero value otherwise.

### GetFallbackUsernameOk

`func (o *UpdateConference) GetFallbackUsernameOk() (*string, bool)`

GetFallbackUsernameOk returns a tuple with the FallbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUsername

`func (o *UpdateConference) SetFallbackUsername(v string)`

SetFallbackUsername sets FallbackUsername field to given value.

### HasFallbackUsername

`func (o *UpdateConference) HasFallbackUsername() bool`

HasFallbackUsername returns a boolean if a field has been set.

### SetFallbackUsernameNil

`func (o *UpdateConference) SetFallbackUsernameNil(b bool)`

 SetFallbackUsernameNil sets the value for FallbackUsername to be an explicit nil

### UnsetFallbackUsername
`func (o *UpdateConference) UnsetFallbackUsername()`

UnsetFallbackUsername ensures that no value is present for FallbackUsername, not even an explicit nil
### GetFallbackPassword

`func (o *UpdateConference) GetFallbackPassword() string`

GetFallbackPassword returns the FallbackPassword field if non-nil, zero value otherwise.

### GetFallbackPasswordOk

`func (o *UpdateConference) GetFallbackPasswordOk() (*string, bool)`

GetFallbackPasswordOk returns a tuple with the FallbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackPassword

`func (o *UpdateConference) SetFallbackPassword(v string)`

SetFallbackPassword sets FallbackPassword field to given value.

### HasFallbackPassword

`func (o *UpdateConference) HasFallbackPassword() bool`

HasFallbackPassword returns a boolean if a field has been set.

### SetFallbackPasswordNil

`func (o *UpdateConference) SetFallbackPasswordNil(b bool)`

 SetFallbackPasswordNil sets the value for FallbackPassword to be an explicit nil

### UnsetFallbackPassword
`func (o *UpdateConference) UnsetFallbackPassword()`

UnsetFallbackPassword ensures that no value is present for FallbackPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


