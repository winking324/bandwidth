# MfaUnauthorizedRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Unauthorized | [optional] 

## Methods

### NewMfaUnauthorizedRequestError

`func NewMfaUnauthorizedRequestError() *MfaUnauthorizedRequestError`

NewMfaUnauthorizedRequestError instantiates a new MfaUnauthorizedRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaUnauthorizedRequestErrorWithDefaults

`func NewMfaUnauthorizedRequestErrorWithDefaults() *MfaUnauthorizedRequestError`

NewMfaUnauthorizedRequestErrorWithDefaults instantiates a new MfaUnauthorizedRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MfaUnauthorizedRequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MfaUnauthorizedRequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MfaUnauthorizedRequestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MfaUnauthorizedRequestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


