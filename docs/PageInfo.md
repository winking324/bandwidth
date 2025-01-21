# PageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevPage** | Pointer to **string** | The link to the previous page for pagination. | [optional] 
**NextPage** | Pointer to **string** | The link to the next page for pagination. | [optional] 
**PrevPageToken** | Pointer to **string** | The isolated pagination token for the previous page. | [optional] 
**NextPageToken** | Pointer to **string** | The isolated pagination token for the next page. | [optional] 

## Methods

### NewPageInfo

`func NewPageInfo() *PageInfo`

NewPageInfo instantiates a new PageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInfoWithDefaults

`func NewPageInfoWithDefaults() *PageInfo`

NewPageInfoWithDefaults instantiates a new PageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevPage

`func (o *PageInfo) GetPrevPage() string`

GetPrevPage returns the PrevPage field if non-nil, zero value otherwise.

### GetPrevPageOk

`func (o *PageInfo) GetPrevPageOk() (*string, bool)`

GetPrevPageOk returns a tuple with the PrevPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPage

`func (o *PageInfo) SetPrevPage(v string)`

SetPrevPage sets PrevPage field to given value.

### HasPrevPage

`func (o *PageInfo) HasPrevPage() bool`

HasPrevPage returns a boolean if a field has been set.

### GetNextPage

`func (o *PageInfo) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PageInfo) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PageInfo) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PageInfo) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPrevPageToken

`func (o *PageInfo) GetPrevPageToken() string`

GetPrevPageToken returns the PrevPageToken field if non-nil, zero value otherwise.

### GetPrevPageTokenOk

`func (o *PageInfo) GetPrevPageTokenOk() (*string, bool)`

GetPrevPageTokenOk returns a tuple with the PrevPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageToken

`func (o *PageInfo) SetPrevPageToken(v string)`

SetPrevPageToken sets PrevPageToken field to given value.

### HasPrevPageToken

`func (o *PageInfo) HasPrevPageToken() bool`

HasPrevPageToken returns a boolean if a field has been set.

### GetNextPageToken

`func (o *PageInfo) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PageInfo) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PageInfo) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *PageInfo) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


