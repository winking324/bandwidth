# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the message. | [optional] 
**Owner** | Pointer to **string** | The Bandwidth phone number associated with the message. | [optional] 
**ApplicationId** | Pointer to **string** | The application ID associated with the message. | [optional] 
**Time** | Pointer to **time.Time** | The datetime stamp of the message in ISO 8601 | [optional] 
**SegmentCount** | Pointer to **int32** | The number of segments the original message from the user is broken into before sending over to carrier networks. | [optional] 
**Direction** | Pointer to [**MessageDirectionEnum**](MessageDirectionEnum.md) |  | [optional] 
**To** | Pointer to **[]string** | The phone number recipients of the message. | [optional] 
**From** | Pointer to **string** | The phone number the message was sent from. | [optional] 
**Media** | Pointer to **[]string** | The list of media URLs sent in the message. Including a &#x60;filename&#x60; field in the &#x60;Content-Disposition&#x60; header of the media linked with a URL will set the displayed file name. This is a best practice to ensure that your media has a readable file name. | [optional] 
**Text** | Pointer to **string** | The contents of the message. | [optional] 
**Tag** | Pointer to **string** | The custom string set by the user. | [optional] 
**Priority** | Pointer to [**PriorityEnum**](PriorityEnum.md) |  | [optional] 
**Expiration** | Pointer to **time.Time** | The expiration date-time set by the user. | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Message) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Message) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Message) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Message) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *Message) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Message) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Message) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Message) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetApplicationId

`func (o *Message) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Message) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Message) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Message) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetTime

`func (o *Message) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Message) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Message) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Message) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetSegmentCount

`func (o *Message) GetSegmentCount() int32`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *Message) GetSegmentCountOk() (*int32, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *Message) SetSegmentCount(v int32)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *Message) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetDirection

`func (o *Message) GetDirection() MessageDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Message) GetDirectionOk() (*MessageDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Message) SetDirection(v MessageDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Message) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTo

`func (o *Message) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Message) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Message) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *Message) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *Message) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Message) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Message) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Message) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMedia

`func (o *Message) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Message) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Message) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Message) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetText

`func (o *Message) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Message) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Message) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Message) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTag

`func (o *Message) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Message) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Message) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Message) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetPriority

`func (o *Message) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Message) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Message) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Message) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetExpiration

`func (o *Message) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Message) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Message) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Message) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


