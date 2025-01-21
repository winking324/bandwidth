# MessageCallbackMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Owner** | **string** |  | 
**ApplicationId** | **string** |  | 
**Time** | **time.Time** |  | 
**SegmentCount** | **int32** |  | 
**Direction** | [**MessageDirectionEnum**](MessageDirectionEnum.md) |  | 
**To** | **[]string** |  | 
**From** | **string** |  | 
**Text** | **string** |  | 
**Tag** | Pointer to **string** |  | [optional] 
**Media** | Pointer to **[]string** | Optional media, applicable only for mms | [optional] 
**Priority** | Pointer to [**PriorityEnum**](PriorityEnum.md) |  | [optional] 

## Methods

### NewMessageCallbackMessage

`func NewMessageCallbackMessage(id string, owner string, applicationId string, time time.Time, segmentCount int32, direction MessageDirectionEnum, to []string, from string, text string, ) *MessageCallbackMessage`

NewMessageCallbackMessage instantiates a new MessageCallbackMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCallbackMessageWithDefaults

`func NewMessageCallbackMessageWithDefaults() *MessageCallbackMessage`

NewMessageCallbackMessageWithDefaults instantiates a new MessageCallbackMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageCallbackMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageCallbackMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageCallbackMessage) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *MessageCallbackMessage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MessageCallbackMessage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MessageCallbackMessage) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetApplicationId

`func (o *MessageCallbackMessage) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MessageCallbackMessage) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MessageCallbackMessage) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetTime

`func (o *MessageCallbackMessage) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MessageCallbackMessage) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MessageCallbackMessage) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetSegmentCount

`func (o *MessageCallbackMessage) GetSegmentCount() int32`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *MessageCallbackMessage) GetSegmentCountOk() (*int32, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *MessageCallbackMessage) SetSegmentCount(v int32)`

SetSegmentCount sets SegmentCount field to given value.


### GetDirection

`func (o *MessageCallbackMessage) GetDirection() MessageDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MessageCallbackMessage) GetDirectionOk() (*MessageDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MessageCallbackMessage) SetDirection(v MessageDirectionEnum)`

SetDirection sets Direction field to given value.


### GetTo

`func (o *MessageCallbackMessage) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageCallbackMessage) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageCallbackMessage) SetTo(v []string)`

SetTo sets To field to given value.


### GetFrom

`func (o *MessageCallbackMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageCallbackMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageCallbackMessage) SetFrom(v string)`

SetFrom sets From field to given value.


### GetText

`func (o *MessageCallbackMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageCallbackMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageCallbackMessage) SetText(v string)`

SetText sets Text field to given value.


### GetTag

`func (o *MessageCallbackMessage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MessageCallbackMessage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MessageCallbackMessage) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *MessageCallbackMessage) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetMedia

`func (o *MessageCallbackMessage) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MessageCallbackMessage) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MessageCallbackMessage) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MessageCallbackMessage) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *MessageCallbackMessage) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *MessageCallbackMessage) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetPriority

`func (o *MessageCallbackMessage) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MessageCallbackMessage) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MessageCallbackMessage) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MessageCallbackMessage) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


