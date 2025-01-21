# CreateLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The phone number lookup request ID from Bandwidth. | [optional] 
**Status** | Pointer to [**LookupStatusEnum**](LookupStatusEnum.md) |  | [optional] 

## Methods

### NewCreateLookupResponse

`func NewCreateLookupResponse() *CreateLookupResponse`

NewCreateLookupResponse instantiates a new CreateLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLookupResponseWithDefaults

`func NewCreateLookupResponseWithDefaults() *CreateLookupResponse`

NewCreateLookupResponseWithDefaults instantiates a new CreateLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateLookupResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateLookupResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateLookupResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateLookupResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateLookupResponse) GetStatus() LookupStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateLookupResponse) GetStatusOk() (*LookupStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateLookupResponse) SetStatus(v LookupStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateLookupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


