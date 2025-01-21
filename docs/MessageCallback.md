# MessageCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** |  | 
**Type** | [**CallbackTypeEnum**](CallbackTypeEnum.md) |  | 
**To** | **string** |  | 
**Description** | **string** | A detailed description of the event described by the callback. | 
**Message** | [**MessageCallbackMessage**](MessageCallbackMessage.md) |  | 
**ErrorCode** | Pointer to **NullableInt32** | Optional error code, applicable only when type is &#x60;message-failed&#x60;. | [optional] 

## Methods

### NewMessageCallback

`func NewMessageCallback(time time.Time, type_ CallbackTypeEnum, to string, description string, message MessageCallbackMessage, ) *MessageCallback`

NewMessageCallback instantiates a new MessageCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCallbackWithDefaults

`func NewMessageCallbackWithDefaults() *MessageCallback`

NewMessageCallbackWithDefaults instantiates a new MessageCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *MessageCallback) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MessageCallback) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MessageCallback) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetType

`func (o *MessageCallback) GetType() CallbackTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageCallback) GetTypeOk() (*CallbackTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageCallback) SetType(v CallbackTypeEnum)`

SetType sets Type field to given value.


### GetTo

`func (o *MessageCallback) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageCallback) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageCallback) SetTo(v string)`

SetTo sets To field to given value.


### GetDescription

`func (o *MessageCallback) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageCallback) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageCallback) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMessage

`func (o *MessageCallback) GetMessage() MessageCallbackMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageCallback) GetMessageOk() (*MessageCallbackMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageCallback) SetMessage(v MessageCallbackMessage)`

SetMessage sets Message field to given value.


### GetErrorCode

`func (o *MessageCallback) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MessageCallback) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MessageCallback) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MessageCallback) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *MessageCallback) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *MessageCallback) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


