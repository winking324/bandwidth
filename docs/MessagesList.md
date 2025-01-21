# MessagesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The total number of messages matched by the search. When the request has limitTotalCount set to true this value is limited to 10,000. | [optional] 
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 
**Messages** | Pointer to [**[]ListMessageItem**](ListMessageItem.md) |  | [optional] 

## Methods

### NewMessagesList

`func NewMessagesList() *MessagesList`

NewMessagesList instantiates a new MessagesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesListWithDefaults

`func NewMessagesListWithDefaults() *MessagesList`

NewMessagesListWithDefaults instantiates a new MessagesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *MessagesList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MessagesList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MessagesList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *MessagesList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPageInfo

`func (o *MessagesList) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *MessagesList) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *MessagesList) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *MessagesList) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetMessages

`func (o *MessagesList) GetMessages() []ListMessageItem`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MessagesList) GetMessagesOk() (*[]ListMessageItem, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MessagesList) SetMessages(v []ListMessageItem)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MessagesList) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


