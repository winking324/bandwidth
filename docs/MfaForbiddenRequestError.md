# MfaForbiddenRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The message containing the reason behind the request being forbidden. | [optional] 

## Methods

### NewMfaForbiddenRequestError

`func NewMfaForbiddenRequestError() *MfaForbiddenRequestError`

NewMfaForbiddenRequestError instantiates a new MfaForbiddenRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaForbiddenRequestErrorWithDefaults

`func NewMfaForbiddenRequestErrorWithDefaults() *MfaForbiddenRequestError`

NewMfaForbiddenRequestErrorWithDefaults instantiates a new MfaForbiddenRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MfaForbiddenRequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MfaForbiddenRequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MfaForbiddenRequestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MfaForbiddenRequestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


