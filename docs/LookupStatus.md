# LookupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The requestId. | [optional] 
**Status** | Pointer to [**LookupStatusEnum**](LookupStatusEnum.md) |  | [optional] 
**Result** | Pointer to [**[]LookupResult**](LookupResult.md) | The carrier information results for the specified telephone number. | [optional] 
**FailedTelephoneNumbers** | Pointer to **[]string** | The telephone numbers whose lookup failed. | [optional] 

## Methods

### NewLookupStatus

`func NewLookupStatus() *LookupStatus`

NewLookupStatus instantiates a new LookupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupStatusWithDefaults

`func NewLookupStatusWithDefaults() *LookupStatus`

NewLookupStatusWithDefaults instantiates a new LookupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *LookupStatus) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LookupStatus) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LookupStatus) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *LookupStatus) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *LookupStatus) GetStatus() LookupStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LookupStatus) GetStatusOk() (*LookupStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LookupStatus) SetStatus(v LookupStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LookupStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *LookupStatus) GetResult() []LookupResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LookupStatus) GetResultOk() (*[]LookupResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LookupStatus) SetResult(v []LookupResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *LookupStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFailedTelephoneNumbers

`func (o *LookupStatus) GetFailedTelephoneNumbers() []string`

GetFailedTelephoneNumbers returns the FailedTelephoneNumbers field if non-nil, zero value otherwise.

### GetFailedTelephoneNumbersOk

`func (o *LookupStatus) GetFailedTelephoneNumbersOk() (*[]string, bool)`

GetFailedTelephoneNumbersOk returns a tuple with the FailedTelephoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTelephoneNumbers

`func (o *LookupStatus) SetFailedTelephoneNumbers(v []string)`

SetFailedTelephoneNumbers sets FailedTelephoneNumbers field to given value.

### HasFailedTelephoneNumbers

`func (o *LookupStatus) HasFailedTelephoneNumbers() bool`

HasFailedTelephoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


