# ListMessageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | The message id | [optional] 
**AccountId** | Pointer to **string** | The account id associated with this message. | [optional] 
**SourceTn** | Pointer to **string** | The source phone number of the message. | [optional] 
**DestinationTn** | Pointer to **string** | The recipient phone number of the message. | [optional] 
**MessageStatus** | Pointer to [**MessageStatusEnum**](MessageStatusEnum.md) |  | [optional] 
**MessageDirection** | Pointer to [**ListMessageDirectionEnum**](ListMessageDirectionEnum.md) |  | [optional] 
**MessageType** | Pointer to [**MessageTypeEnum**](MessageTypeEnum.md) |  | [optional] 
**SegmentCount** | Pointer to **int32** | The number of segments the message was sent as. | [optional] 
**ErrorCode** | Pointer to **int32** | The numeric error code of the message. | [optional] 
**ReceiveTime** | Pointer to **time.Time** | The ISO 8601 datetime of the message. | [optional] 
**CarrierName** | Pointer to **NullableString** | The name of the carrier. Not currently supported for MMS coming soon. | [optional] 
**MessageSize** | Pointer to **NullableInt32** | The size of the message including message content and headers. | [optional] 
**MessageLength** | Pointer to **int32** | The length of the message content. | [optional] 
**AttachmentCount** | Pointer to **NullableInt32** | The number of attachments the message has. | [optional] 
**RecipientCount** | Pointer to **NullableInt32** | The number of recipients the message has. | [optional] 
**CampaignClass** | Pointer to **NullableString** | The campaign class of the message if it has one. | [optional] 
**CampaignId** | Pointer to **NullableString** | The campaign ID of the message if it has one. | [optional] 

## Methods

### NewListMessageItem

`func NewListMessageItem() *ListMessageItem`

NewListMessageItem instantiates a new ListMessageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessageItemWithDefaults

`func NewListMessageItemWithDefaults() *ListMessageItem`

NewListMessageItemWithDefaults instantiates a new ListMessageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *ListMessageItem) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ListMessageItem) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ListMessageItem) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ListMessageItem) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListMessageItem) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListMessageItem) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListMessageItem) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListMessageItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSourceTn

`func (o *ListMessageItem) GetSourceTn() string`

GetSourceTn returns the SourceTn field if non-nil, zero value otherwise.

### GetSourceTnOk

`func (o *ListMessageItem) GetSourceTnOk() (*string, bool)`

GetSourceTnOk returns a tuple with the SourceTn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTn

`func (o *ListMessageItem) SetSourceTn(v string)`

SetSourceTn sets SourceTn field to given value.

### HasSourceTn

`func (o *ListMessageItem) HasSourceTn() bool`

HasSourceTn returns a boolean if a field has been set.

### GetDestinationTn

`func (o *ListMessageItem) GetDestinationTn() string`

GetDestinationTn returns the DestinationTn field if non-nil, zero value otherwise.

### GetDestinationTnOk

`func (o *ListMessageItem) GetDestinationTnOk() (*string, bool)`

GetDestinationTnOk returns a tuple with the DestinationTn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTn

`func (o *ListMessageItem) SetDestinationTn(v string)`

SetDestinationTn sets DestinationTn field to given value.

### HasDestinationTn

`func (o *ListMessageItem) HasDestinationTn() bool`

HasDestinationTn returns a boolean if a field has been set.

### GetMessageStatus

`func (o *ListMessageItem) GetMessageStatus() MessageStatusEnum`

GetMessageStatus returns the MessageStatus field if non-nil, zero value otherwise.

### GetMessageStatusOk

`func (o *ListMessageItem) GetMessageStatusOk() (*MessageStatusEnum, bool)`

GetMessageStatusOk returns a tuple with the MessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatus

`func (o *ListMessageItem) SetMessageStatus(v MessageStatusEnum)`

SetMessageStatus sets MessageStatus field to given value.

### HasMessageStatus

`func (o *ListMessageItem) HasMessageStatus() bool`

HasMessageStatus returns a boolean if a field has been set.

### GetMessageDirection

`func (o *ListMessageItem) GetMessageDirection() ListMessageDirectionEnum`

GetMessageDirection returns the MessageDirection field if non-nil, zero value otherwise.

### GetMessageDirectionOk

`func (o *ListMessageItem) GetMessageDirectionOk() (*ListMessageDirectionEnum, bool)`

GetMessageDirectionOk returns a tuple with the MessageDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDirection

`func (o *ListMessageItem) SetMessageDirection(v ListMessageDirectionEnum)`

SetMessageDirection sets MessageDirection field to given value.

### HasMessageDirection

`func (o *ListMessageItem) HasMessageDirection() bool`

HasMessageDirection returns a boolean if a field has been set.

### GetMessageType

`func (o *ListMessageItem) GetMessageType() MessageTypeEnum`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ListMessageItem) GetMessageTypeOk() (*MessageTypeEnum, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ListMessageItem) SetMessageType(v MessageTypeEnum)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ListMessageItem) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSegmentCount

`func (o *ListMessageItem) GetSegmentCount() int32`

GetSegmentCount returns the SegmentCount field if non-nil, zero value otherwise.

### GetSegmentCountOk

`func (o *ListMessageItem) GetSegmentCountOk() (*int32, bool)`

GetSegmentCountOk returns a tuple with the SegmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentCount

`func (o *ListMessageItem) SetSegmentCount(v int32)`

SetSegmentCount sets SegmentCount field to given value.

### HasSegmentCount

`func (o *ListMessageItem) HasSegmentCount() bool`

HasSegmentCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *ListMessageItem) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ListMessageItem) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ListMessageItem) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ListMessageItem) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetReceiveTime

`func (o *ListMessageItem) GetReceiveTime() time.Time`

GetReceiveTime returns the ReceiveTime field if non-nil, zero value otherwise.

### GetReceiveTimeOk

`func (o *ListMessageItem) GetReceiveTimeOk() (*time.Time, bool)`

GetReceiveTimeOk returns a tuple with the ReceiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTime

`func (o *ListMessageItem) SetReceiveTime(v time.Time)`

SetReceiveTime sets ReceiveTime field to given value.

### HasReceiveTime

`func (o *ListMessageItem) HasReceiveTime() bool`

HasReceiveTime returns a boolean if a field has been set.

### GetCarrierName

`func (o *ListMessageItem) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *ListMessageItem) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *ListMessageItem) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *ListMessageItem) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### SetCarrierNameNil

`func (o *ListMessageItem) SetCarrierNameNil(b bool)`

 SetCarrierNameNil sets the value for CarrierName to be an explicit nil

### UnsetCarrierName
`func (o *ListMessageItem) UnsetCarrierName()`

UnsetCarrierName ensures that no value is present for CarrierName, not even an explicit nil
### GetMessageSize

`func (o *ListMessageItem) GetMessageSize() int32`

GetMessageSize returns the MessageSize field if non-nil, zero value otherwise.

### GetMessageSizeOk

`func (o *ListMessageItem) GetMessageSizeOk() (*int32, bool)`

GetMessageSizeOk returns a tuple with the MessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSize

`func (o *ListMessageItem) SetMessageSize(v int32)`

SetMessageSize sets MessageSize field to given value.

### HasMessageSize

`func (o *ListMessageItem) HasMessageSize() bool`

HasMessageSize returns a boolean if a field has been set.

### SetMessageSizeNil

`func (o *ListMessageItem) SetMessageSizeNil(b bool)`

 SetMessageSizeNil sets the value for MessageSize to be an explicit nil

### UnsetMessageSize
`func (o *ListMessageItem) UnsetMessageSize()`

UnsetMessageSize ensures that no value is present for MessageSize, not even an explicit nil
### GetMessageLength

`func (o *ListMessageItem) GetMessageLength() int32`

GetMessageLength returns the MessageLength field if non-nil, zero value otherwise.

### GetMessageLengthOk

`func (o *ListMessageItem) GetMessageLengthOk() (*int32, bool)`

GetMessageLengthOk returns a tuple with the MessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageLength

`func (o *ListMessageItem) SetMessageLength(v int32)`

SetMessageLength sets MessageLength field to given value.

### HasMessageLength

`func (o *ListMessageItem) HasMessageLength() bool`

HasMessageLength returns a boolean if a field has been set.

### GetAttachmentCount

`func (o *ListMessageItem) GetAttachmentCount() int32`

GetAttachmentCount returns the AttachmentCount field if non-nil, zero value otherwise.

### GetAttachmentCountOk

`func (o *ListMessageItem) GetAttachmentCountOk() (*int32, bool)`

GetAttachmentCountOk returns a tuple with the AttachmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentCount

`func (o *ListMessageItem) SetAttachmentCount(v int32)`

SetAttachmentCount sets AttachmentCount field to given value.

### HasAttachmentCount

`func (o *ListMessageItem) HasAttachmentCount() bool`

HasAttachmentCount returns a boolean if a field has been set.

### SetAttachmentCountNil

`func (o *ListMessageItem) SetAttachmentCountNil(b bool)`

 SetAttachmentCountNil sets the value for AttachmentCount to be an explicit nil

### UnsetAttachmentCount
`func (o *ListMessageItem) UnsetAttachmentCount()`

UnsetAttachmentCount ensures that no value is present for AttachmentCount, not even an explicit nil
### GetRecipientCount

`func (o *ListMessageItem) GetRecipientCount() int32`

GetRecipientCount returns the RecipientCount field if non-nil, zero value otherwise.

### GetRecipientCountOk

`func (o *ListMessageItem) GetRecipientCountOk() (*int32, bool)`

GetRecipientCountOk returns a tuple with the RecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCount

`func (o *ListMessageItem) SetRecipientCount(v int32)`

SetRecipientCount sets RecipientCount field to given value.

### HasRecipientCount

`func (o *ListMessageItem) HasRecipientCount() bool`

HasRecipientCount returns a boolean if a field has been set.

### SetRecipientCountNil

`func (o *ListMessageItem) SetRecipientCountNil(b bool)`

 SetRecipientCountNil sets the value for RecipientCount to be an explicit nil

### UnsetRecipientCount
`func (o *ListMessageItem) UnsetRecipientCount()`

UnsetRecipientCount ensures that no value is present for RecipientCount, not even an explicit nil
### GetCampaignClass

`func (o *ListMessageItem) GetCampaignClass() string`

GetCampaignClass returns the CampaignClass field if non-nil, zero value otherwise.

### GetCampaignClassOk

`func (o *ListMessageItem) GetCampaignClassOk() (*string, bool)`

GetCampaignClassOk returns a tuple with the CampaignClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignClass

`func (o *ListMessageItem) SetCampaignClass(v string)`

SetCampaignClass sets CampaignClass field to given value.

### HasCampaignClass

`func (o *ListMessageItem) HasCampaignClass() bool`

HasCampaignClass returns a boolean if a field has been set.

### SetCampaignClassNil

`func (o *ListMessageItem) SetCampaignClassNil(b bool)`

 SetCampaignClassNil sets the value for CampaignClass to be an explicit nil

### UnsetCampaignClass
`func (o *ListMessageItem) UnsetCampaignClass()`

UnsetCampaignClass ensures that no value is present for CampaignClass, not even an explicit nil
### GetCampaignId

`func (o *ListMessageItem) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ListMessageItem) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ListMessageItem) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ListMessageItem) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *ListMessageItem) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *ListMessageItem) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


