# MfaRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | A message describing the error with your request. | [optional] 
**RequestId** | Pointer to **string** | The associated requestId from AWS. | [optional] 

## Methods

### NewMfaRequestError

`func NewMfaRequestError() *MfaRequestError`

NewMfaRequestError instantiates a new MfaRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaRequestErrorWithDefaults

`func NewMfaRequestErrorWithDefaults() *MfaRequestError`

NewMfaRequestErrorWithDefaults instantiates a new MfaRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MfaRequestError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MfaRequestError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MfaRequestError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MfaRequestError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *MfaRequestError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MfaRequestError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MfaRequestError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *MfaRequestError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


