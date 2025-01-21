/*
Bandwidth

Bandwidth's Communication APIs

API version: 1.0.0
Contact: letstalk@bandwidth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)




// ListMessageItem struct for ListMessageItem
type ListMessageItem struct {
	// The message id
	MessageId *string `json:"messageId,omitempty"`
	// The account id associated with this message.
	AccountId *string `json:"accountId,omitempty"`
	// The source phone number of the message.
	SourceTn *string `json:"sourceTn,omitempty"`
	// The recipient phone number of the message.
	DestinationTn *string `json:"destinationTn,omitempty"`
	MessageStatus *MessageStatusEnum `json:"messageStatus,omitempty"`
	MessageDirection *ListMessageDirectionEnum `json:"messageDirection,omitempty"`
	MessageType *MessageTypeEnum `json:"messageType,omitempty"`
	// The number of segments the message was sent as.
	SegmentCount *int32 `json:"segmentCount,omitempty"`
	// The numeric error code of the message.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// The ISO 8601 datetime of the message.
	ReceiveTime *time.Time `json:"receiveTime,omitempty"`
	// The name of the carrier. Not currently supported for MMS coming soon.
	CarrierName utils.NullableString `json:"carrierName,omitempty"`
	// The size of the message including message content and headers.
	MessageSize NullableInt32 `json:"messageSize,omitempty"`
	// The length of the message content.
	MessageLength *int32 `json:"messageLength,omitempty"`
	// The number of attachments the message has.
	AttachmentCount NullableInt32 `json:"attachmentCount,omitempty"`
	// The number of recipients the message has.
	RecipientCount NullableInt32 `json:"recipientCount,omitempty"`
	// The campaign class of the message if it has one.
	CampaignClass utils.NullableString `json:"campaignClass,omitempty"`
	// The campaign ID of the message if it has one.
	CampaignId utils.NullableString `json:"campaignId,omitempty"`
}

// NewListMessageItem instantiates a new ListMessageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMessageItem() *ListMessageItem {
	this := ListMessageItem{}
	return &this
}

// NewListMessageItemWithDefaults instantiates a new ListMessageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMessageItemWithDefaults() *ListMessageItem {
	this := ListMessageItem{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ListMessageItem) GetMessageId() string {
	if o == nil || utils.IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetMessageIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ListMessageItem) HasMessageId() bool {
	if o != nil && !utils.IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ListMessageItem) SetMessageId(v string) {
	o.MessageId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ListMessageItem) GetAccountId() string {
	if o == nil || utils.IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetAccountIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ListMessageItem) HasAccountId() bool {
	if o != nil && !utils.IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ListMessageItem) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSourceTn returns the SourceTn field value if set, zero value otherwise.
func (o *ListMessageItem) GetSourceTn() string {
	if o == nil || utils.IsNil(o.SourceTn) {
		var ret string
		return ret
	}
	return *o.SourceTn
}

// GetSourceTnOk returns a tuple with the SourceTn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetSourceTnOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceTn) {
		return nil, false
	}
	return o.SourceTn, true
}

// HasSourceTn returns a boolean if a field has been set.
func (o *ListMessageItem) HasSourceTn() bool {
	if o != nil && !utils.IsNil(o.SourceTn) {
		return true
	}

	return false
}

// SetSourceTn gets a reference to the given string and assigns it to the SourceTn field.
func (o *ListMessageItem) SetSourceTn(v string) {
	o.SourceTn = &v
}

// GetDestinationTn returns the DestinationTn field value if set, zero value otherwise.
func (o *ListMessageItem) GetDestinationTn() string {
	if o == nil || utils.IsNil(o.DestinationTn) {
		var ret string
		return ret
	}
	return *o.DestinationTn
}

// GetDestinationTnOk returns a tuple with the DestinationTn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetDestinationTnOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DestinationTn) {
		return nil, false
	}
	return o.DestinationTn, true
}

// HasDestinationTn returns a boolean if a field has been set.
func (o *ListMessageItem) HasDestinationTn() bool {
	if o != nil && !utils.IsNil(o.DestinationTn) {
		return true
	}

	return false
}

// SetDestinationTn gets a reference to the given string and assigns it to the DestinationTn field.
func (o *ListMessageItem) SetDestinationTn(v string) {
	o.DestinationTn = &v
}

// GetMessageStatus returns the MessageStatus field value if set, zero value otherwise.
func (o *ListMessageItem) GetMessageStatus() MessageStatusEnum {
	if o == nil || utils.IsNil(o.MessageStatus) {
		var ret MessageStatusEnum
		return ret
	}
	return *o.MessageStatus
}

// GetMessageStatusOk returns a tuple with the MessageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetMessageStatusOk() (*MessageStatusEnum, bool) {
	if o == nil || utils.IsNil(o.MessageStatus) {
		return nil, false
	}
	return o.MessageStatus, true
}

// HasMessageStatus returns a boolean if a field has been set.
func (o *ListMessageItem) HasMessageStatus() bool {
	if o != nil && !utils.IsNil(o.MessageStatus) {
		return true
	}

	return false
}

// SetMessageStatus gets a reference to the given MessageStatusEnum and assigns it to the MessageStatus field.
func (o *ListMessageItem) SetMessageStatus(v MessageStatusEnum) {
	o.MessageStatus = &v
}

// GetMessageDirection returns the MessageDirection field value if set, zero value otherwise.
func (o *ListMessageItem) GetMessageDirection() ListMessageDirectionEnum {
	if o == nil || utils.IsNil(o.MessageDirection) {
		var ret ListMessageDirectionEnum
		return ret
	}
	return *o.MessageDirection
}

// GetMessageDirectionOk returns a tuple with the MessageDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetMessageDirectionOk() (*ListMessageDirectionEnum, bool) {
	if o == nil || utils.IsNil(o.MessageDirection) {
		return nil, false
	}
	return o.MessageDirection, true
}

// HasMessageDirection returns a boolean if a field has been set.
func (o *ListMessageItem) HasMessageDirection() bool {
	if o != nil && !utils.IsNil(o.MessageDirection) {
		return true
	}

	return false
}

// SetMessageDirection gets a reference to the given ListMessageDirectionEnum and assigns it to the MessageDirection field.
func (o *ListMessageItem) SetMessageDirection(v ListMessageDirectionEnum) {
	o.MessageDirection = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ListMessageItem) GetMessageType() MessageTypeEnum {
	if o == nil || utils.IsNil(o.MessageType) {
		var ret MessageTypeEnum
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetMessageTypeOk() (*MessageTypeEnum, bool) {
	if o == nil || utils.IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ListMessageItem) HasMessageType() bool {
	if o != nil && !utils.IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given MessageTypeEnum and assigns it to the MessageType field.
func (o *ListMessageItem) SetMessageType(v MessageTypeEnum) {
	o.MessageType = &v
}

// GetSegmentCount returns the SegmentCount field value if set, zero value otherwise.
func (o *ListMessageItem) GetSegmentCount() int32 {
	if o == nil || utils.IsNil(o.SegmentCount) {
		var ret int32
		return ret
	}
	return *o.SegmentCount
}

// GetSegmentCountOk returns a tuple with the SegmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetSegmentCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SegmentCount) {
		return nil, false
	}
	return o.SegmentCount, true
}

// HasSegmentCount returns a boolean if a field has been set.
func (o *ListMessageItem) HasSegmentCount() bool {
	if o != nil && !utils.IsNil(o.SegmentCount) {
		return true
	}

	return false
}

// SetSegmentCount gets a reference to the given int32 and assigns it to the SegmentCount field.
func (o *ListMessageItem) SetSegmentCount(v int32) {
	o.SegmentCount = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ListMessageItem) GetErrorCode() int32 {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetErrorCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ListMessageItem) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *ListMessageItem) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetReceiveTime returns the ReceiveTime field value if set, zero value otherwise.
func (o *ListMessageItem) GetReceiveTime() time.Time {
	if o == nil || utils.IsNil(o.ReceiveTime) {
		var ret time.Time
		return ret
	}
	return *o.ReceiveTime
}

// GetReceiveTimeOk returns a tuple with the ReceiveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetReceiveTimeOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ReceiveTime) {
		return nil, false
	}
	return o.ReceiveTime, true
}

// HasReceiveTime returns a boolean if a field has been set.
func (o *ListMessageItem) HasReceiveTime() bool {
	if o != nil && !utils.IsNil(o.ReceiveTime) {
		return true
	}

	return false
}

// SetReceiveTime gets a reference to the given time.Time and assigns it to the ReceiveTime field.
func (o *ListMessageItem) SetReceiveTime(v time.Time) {
	o.ReceiveTime = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMessageItem) GetCarrierName() string {
	if o == nil || utils.IsNil(o.CarrierName.Get()) {
		var ret string
		return ret
	}
	return *o.CarrierName.Get()
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMessageItem) GetCarrierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CarrierName.Get(), o.CarrierName.IsSet()
}

// HasCarrierName returns a boolean if a field has been set.
func (o *ListMessageItem) HasCarrierName() bool {
	if o != nil && o.CarrierName.IsSet() {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given utils.NullableString and assigns it to the CarrierName field.
func (o *ListMessageItem) SetCarrierName(v string) {
	o.CarrierName.Set(&v)
}
// SetCarrierNameNil sets the value for CarrierName to be an explicit nil
func (o *ListMessageItem) SetCarrierNameNil() {
	o.CarrierName.Set(nil)
}

// UnsetCarrierName ensures that no value is present for CarrierName, not even an explicit nil
func (o *ListMessageItem) UnsetCarrierName() {
	o.CarrierName.Unset()
}

// GetMessageSize returns the MessageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMessageItem) GetMessageSize() int32 {
	if o == nil || utils.IsNil(o.MessageSize.Get()) {
		var ret int32
		return ret
	}
	return *o.MessageSize.Get()
}

// GetMessageSizeOk returns a tuple with the MessageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMessageItem) GetMessageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageSize.Get(), o.MessageSize.IsSet()
}

// HasMessageSize returns a boolean if a field has been set.
func (o *ListMessageItem) HasMessageSize() bool {
	if o != nil && o.MessageSize.IsSet() {
		return true
	}

	return false
}

// SetMessageSize gets a reference to the given NullableInt32 and assigns it to the MessageSize field.
func (o *ListMessageItem) SetMessageSize(v int32) {
	o.MessageSize.Set(&v)
}
// SetMessageSizeNil sets the value for MessageSize to be an explicit nil
func (o *ListMessageItem) SetMessageSizeNil() {
	o.MessageSize.Set(nil)
}

// UnsetMessageSize ensures that no value is present for MessageSize, not even an explicit nil
func (o *ListMessageItem) UnsetMessageSize() {
	o.MessageSize.Unset()
}

// GetMessageLength returns the MessageLength field value if set, zero value otherwise.
func (o *ListMessageItem) GetMessageLength() int32 {
	if o == nil || utils.IsNil(o.MessageLength) {
		var ret int32
		return ret
	}
	return *o.MessageLength
}

// GetMessageLengthOk returns a tuple with the MessageLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageItem) GetMessageLengthOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MessageLength) {
		return nil, false
	}
	return o.MessageLength, true
}

// HasMessageLength returns a boolean if a field has been set.
func (o *ListMessageItem) HasMessageLength() bool {
	if o != nil && !utils.IsNil(o.MessageLength) {
		return true
	}

	return false
}

// SetMessageLength gets a reference to the given int32 and assigns it to the MessageLength field.
func (o *ListMessageItem) SetMessageLength(v int32) {
	o.MessageLength = &v
}

// GetAttachmentCount returns the AttachmentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMessageItem) GetAttachmentCount() int32 {
	if o == nil || utils.IsNil(o.AttachmentCount.Get()) {
		var ret int32
		return ret
	}
	return *o.AttachmentCount.Get()
}

// GetAttachmentCountOk returns a tuple with the AttachmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMessageItem) GetAttachmentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentCount.Get(), o.AttachmentCount.IsSet()
}

// HasAttachmentCount returns a boolean if a field has been set.
func (o *ListMessageItem) HasAttachmentCount() bool {
	if o != nil && o.AttachmentCount.IsSet() {
		return true
	}

	return false
}

// SetAttachmentCount gets a reference to the given NullableInt32 and assigns it to the AttachmentCount field.
func (o *ListMessageItem) SetAttachmentCount(v int32) {
	o.AttachmentCount.Set(&v)
}
// SetAttachmentCountNil sets the value for AttachmentCount to be an explicit nil
func (o *ListMessageItem) SetAttachmentCountNil() {
	o.AttachmentCount.Set(nil)
}

// UnsetAttachmentCount ensures that no value is present for AttachmentCount, not even an explicit nil
func (o *ListMessageItem) UnsetAttachmentCount() {
	o.AttachmentCount.Unset()
}

// GetRecipientCount returns the RecipientCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMessageItem) GetRecipientCount() int32 {
	if o == nil || utils.IsNil(o.RecipientCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RecipientCount.Get()
}

// GetRecipientCountOk returns a tuple with the RecipientCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMessageItem) GetRecipientCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientCount.Get(), o.RecipientCount.IsSet()
}

// HasRecipientCount returns a boolean if a field has been set.
func (o *ListMessageItem) HasRecipientCount() bool {
	if o != nil && o.RecipientCount.IsSet() {
		return true
	}

	return false
}

// SetRecipientCount gets a reference to the given NullableInt32 and assigns it to the RecipientCount field.
func (o *ListMessageItem) SetRecipientCount(v int32) {
	o.RecipientCount.Set(&v)
}
// SetRecipientCountNil sets the value for RecipientCount to be an explicit nil
func (o *ListMessageItem) SetRecipientCountNil() {
	o.RecipientCount.Set(nil)
}

// UnsetRecipientCount ensures that no value is present for RecipientCount, not even an explicit nil
func (o *ListMessageItem) UnsetRecipientCount() {
	o.RecipientCount.Unset()
}

// GetCampaignClass returns the CampaignClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMessageItem) GetCampaignClass() string {
	if o == nil || utils.IsNil(o.CampaignClass.Get()) {
		var ret string
		return ret
	}
	return *o.CampaignClass.Get()
}

// GetCampaignClassOk returns a tuple with the CampaignClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMessageItem) GetCampaignClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignClass.Get(), o.CampaignClass.IsSet()
}

// HasCampaignClass returns a boolean if a field has been set.
func (o *ListMessageItem) HasCampaignClass() bool {
	if o != nil && o.CampaignClass.IsSet() {
		return true
	}

	return false
}

// SetCampaignClass gets a reference to the given utils.NullableString and assigns it to the CampaignClass field.
func (o *ListMessageItem) SetCampaignClass(v string) {
	o.CampaignClass.Set(&v)
}
// SetCampaignClassNil sets the value for CampaignClass to be an explicit nil
func (o *ListMessageItem) SetCampaignClassNil() {
	o.CampaignClass.Set(nil)
}

// UnsetCampaignClass ensures that no value is present for CampaignClass, not even an explicit nil
func (o *ListMessageItem) UnsetCampaignClass() {
	o.CampaignClass.Unset()
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMessageItem) GetCampaignId() string {
	if o == nil || utils.IsNil(o.CampaignId.Get()) {
		var ret string
		return ret
	}
	return *o.CampaignId.Get()
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMessageItem) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignId.Get(), o.CampaignId.IsSet()
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ListMessageItem) HasCampaignId() bool {
	if o != nil && o.CampaignId.IsSet() {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given utils.NullableString and assigns it to the CampaignId field.
func (o *ListMessageItem) SetCampaignId(v string) {
	o.CampaignId.Set(&v)
}
// SetCampaignIdNil sets the value for CampaignId to be an explicit nil
func (o *ListMessageItem) SetCampaignIdNil() {
	o.CampaignId.Set(nil)
}

// UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
func (o *ListMessageItem) UnsetCampaignId() {
	o.CampaignId.Unset()
}

func (o ListMessageItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMessageItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !utils.IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !utils.IsNil(o.SourceTn) {
		toSerialize["sourceTn"] = o.SourceTn
	}
	if !utils.IsNil(o.DestinationTn) {
		toSerialize["destinationTn"] = o.DestinationTn
	}
	if !utils.IsNil(o.MessageStatus) {
		toSerialize["messageStatus"] = o.MessageStatus
	}
	if !utils.IsNil(o.MessageDirection) {
		toSerialize["messageDirection"] = o.MessageDirection
	}
	if !utils.IsNil(o.MessageType) {
		toSerialize["messageType"] = o.MessageType
	}
	if !utils.IsNil(o.SegmentCount) {
		toSerialize["segmentCount"] = o.SegmentCount
	}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !utils.IsNil(o.ReceiveTime) {
		toSerialize["receiveTime"] = o.ReceiveTime
	}
	if o.CarrierName.IsSet() {
		toSerialize["carrierName"] = o.CarrierName.Get()
	}
	if o.MessageSize.IsSet() {
		toSerialize["messageSize"] = o.MessageSize.Get()
	}
	if !utils.IsNil(o.MessageLength) {
		toSerialize["messageLength"] = o.MessageLength
	}
	if o.AttachmentCount.IsSet() {
		toSerialize["attachmentCount"] = o.AttachmentCount.Get()
	}
	if o.RecipientCount.IsSet() {
		toSerialize["recipientCount"] = o.RecipientCount.Get()
	}
	if o.CampaignClass.IsSet() {
		toSerialize["campaignClass"] = o.CampaignClass.Get()
	}
	if o.CampaignId.IsSet() {
		toSerialize["campaignId"] = o.CampaignId.Get()
	}
	return toSerialize, nil
}

type NullableListMessageItem struct {
	value *ListMessageItem
	isSet bool
}

func (v NullableListMessageItem) Get() *ListMessageItem {
	return v.value
}

func (v *NullableListMessageItem) Set(val *ListMessageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListMessageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListMessageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMessageItem(val *ListMessageItem) *NullableListMessageItem {
	return &NullableListMessageItem{value: val, isSet: true}
}

func (v NullableListMessageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMessageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


