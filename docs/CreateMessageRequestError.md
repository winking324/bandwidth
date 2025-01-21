# CreateMessageRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Description** | **string** |  | 
**FieldErrors** | Pointer to [**[]FieldError**](FieldError.md) |  | [optional] 

## Methods

### NewCreateMessageRequestError

`func NewCreateMessageRequestError(type_ string, description string, ) *CreateMessageRequestError`

NewCreateMessageRequestError instantiates a new CreateMessageRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageRequestErrorWithDefaults

`func NewCreateMessageRequestErrorWithDefaults() *CreateMessageRequestError`

NewCreateMessageRequestErrorWithDefaults instantiates a new CreateMessageRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateMessageRequestError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMessageRequestError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMessageRequestError) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *CreateMessageRequestError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMessageRequestError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMessageRequestError) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFieldErrors

`func (o *CreateMessageRequestError) GetFieldErrors() []FieldError`

GetFieldErrors returns the FieldErrors field if non-nil, zero value otherwise.

### GetFieldErrorsOk

`func (o *CreateMessageRequestError) GetFieldErrorsOk() (*[]FieldError, bool)`

GetFieldErrorsOk returns a tuple with the FieldErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldErrors

`func (o *CreateMessageRequestError) SetFieldErrors(v []FieldError)`

SetFieldErrors sets FieldErrors field to given value.

### HasFieldErrors

`func (o *CreateMessageRequestError) HasFieldErrors() bool`

HasFieldErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


