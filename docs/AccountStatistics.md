# AccountStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCallQueueSize** | Pointer to **int32** | The number of calls currently enqueued. | [optional] 
**MaxCallQueueSize** | Pointer to **int32** | The maximum size of the queue before outgoing calls start being rejected. | [optional] 

## Methods

### NewAccountStatistics

`func NewAccountStatistics() *AccountStatistics`

NewAccountStatistics instantiates a new AccountStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatisticsWithDefaults

`func NewAccountStatisticsWithDefaults() *AccountStatistics`

NewAccountStatisticsWithDefaults instantiates a new AccountStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCallQueueSize

`func (o *AccountStatistics) GetCurrentCallQueueSize() int32`

GetCurrentCallQueueSize returns the CurrentCallQueueSize field if non-nil, zero value otherwise.

### GetCurrentCallQueueSizeOk

`func (o *AccountStatistics) GetCurrentCallQueueSizeOk() (*int32, bool)`

GetCurrentCallQueueSizeOk returns a tuple with the CurrentCallQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCallQueueSize

`func (o *AccountStatistics) SetCurrentCallQueueSize(v int32)`

SetCurrentCallQueueSize sets CurrentCallQueueSize field to given value.

### HasCurrentCallQueueSize

`func (o *AccountStatistics) HasCurrentCallQueueSize() bool`

HasCurrentCallQueueSize returns a boolean if a field has been set.

### GetMaxCallQueueSize

`func (o *AccountStatistics) GetMaxCallQueueSize() int32`

GetMaxCallQueueSize returns the MaxCallQueueSize field if non-nil, zero value otherwise.

### GetMaxCallQueueSizeOk

`func (o *AccountStatistics) GetMaxCallQueueSizeOk() (*int32, bool)`

GetMaxCallQueueSizeOk returns a tuple with the MaxCallQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCallQueueSize

`func (o *AccountStatistics) SetMaxCallQueueSize(v int32)`

SetMaxCallQueueSize sets MaxCallQueueSize field to given value.

### HasMaxCallQueueSize

`func (o *AccountStatistics) HasMaxCallQueueSize() bool`

HasMaxCallQueueSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


