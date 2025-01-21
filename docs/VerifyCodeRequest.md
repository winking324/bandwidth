# VerifyCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The phone number to send the mfa code to. | 
**Scope** | Pointer to **string** | An optional field to denote what scope or action the mfa code is addressing.  If not supplied, defaults to \&quot;2FA\&quot;. | [optional] 
**ExpirationTimeInMinutes** | **float32** | The time period, in minutes, to validate the mfa code.  By setting this to 3 minutes, it will mean any code generated within the last 3 minutes are still valid.  The valid range for expiration time is between 0 and 15 minutes, exclusively and inclusively, respectively. | 
**Code** | **string** | The generated mfa code to check if valid. | 

## Methods

### NewVerifyCodeRequest

`func NewVerifyCodeRequest(to string, expirationTimeInMinutes float32, code string, ) *VerifyCodeRequest`

NewVerifyCodeRequest instantiates a new VerifyCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyCodeRequestWithDefaults

`func NewVerifyCodeRequestWithDefaults() *VerifyCodeRequest`

NewVerifyCodeRequestWithDefaults instantiates a new VerifyCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *VerifyCodeRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VerifyCodeRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VerifyCodeRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetScope

`func (o *VerifyCodeRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VerifyCodeRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VerifyCodeRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *VerifyCodeRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpirationTimeInMinutes

`func (o *VerifyCodeRequest) GetExpirationTimeInMinutes() float32`

GetExpirationTimeInMinutes returns the ExpirationTimeInMinutes field if non-nil, zero value otherwise.

### GetExpirationTimeInMinutesOk

`func (o *VerifyCodeRequest) GetExpirationTimeInMinutesOk() (*float32, bool)`

GetExpirationTimeInMinutesOk returns a tuple with the ExpirationTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimeInMinutes

`func (o *VerifyCodeRequest) SetExpirationTimeInMinutes(v float32)`

SetExpirationTimeInMinutes sets ExpirationTimeInMinutes field to given value.


### GetCode

`func (o *VerifyCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


