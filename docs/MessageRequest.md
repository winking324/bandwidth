# MessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The ID of the Application your from number is associated with in the Bandwidth Phone Number Dashboard. | 
**To** | **[]string** | The phone number(s) the message should be sent to in E164 format. | 
**From** | **string** | Either an alphanumeric sender ID or the sender&#39;s Bandwidth phone number in E.164 format, which must be hosted within Bandwidth and linked to the account that is generating the message.  Alphanumeric Sender IDs can contain up to 11 characters, upper-case letters A-Z, lower-case letters a-z, numbers 0-9, space, hyphen -, plus +, underscore _ and ampersand &amp;. Alphanumeric Sender IDs must contain at least one letter. | 
**Text** | Pointer to **string** | The contents of the text message. Must be 2048 characters or less. | [optional] 
**Media** | Pointer to **[]string** | A list of URLs to include as media attachments as part of the message. Each URL can be at most 4096 characters. | [optional] 
**Tag** | Pointer to **string** | A custom string that will be included in callback events of the message. Max 1024 characters. | [optional] 
**Priority** | Pointer to [**PriorityEnum**](PriorityEnum.md) |  | [optional] 
**Expiration** | Pointer to **time.Time** | A string with the date/time value that the message will automatically expire by. This must be a valid RFC-3339 value, e.g., 2021-03-14T01:59:26Z or 2021-03-13T20:59:26-05:00. Must be a date-time in the future. Not supported on MMS. | [optional] 

## Methods

### NewMessageRequest

`func NewMessageRequest(applicationId string, to []string, from string, ) *MessageRequest`

NewMessageRequest instantiates a new MessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageRequestWithDefaults

`func NewMessageRequestWithDefaults() *MessageRequest`

NewMessageRequestWithDefaults instantiates a new MessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *MessageRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MessageRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MessageRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetTo

`func (o *MessageRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetFrom

`func (o *MessageRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetText

`func (o *MessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetMedia

`func (o *MessageRequest) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MessageRequest) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MessageRequest) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MessageRequest) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetTag

`func (o *MessageRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MessageRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MessageRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *MessageRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetPriority

`func (o *MessageRequest) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MessageRequest) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MessageRequest) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MessageRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetExpiration

`func (o *MessageRequest) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *MessageRequest) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *MessageRequest) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *MessageRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


