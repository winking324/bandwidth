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
)




// MessagesList struct for MessagesList
type MessagesList struct {
	// The total number of messages matched by the search. When the request has limitTotalCount set to true this value is limited to 10,000.
	TotalCount *int32 `json:"totalCount,omitempty"`
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
	Messages []ListMessageItem `json:"messages,omitempty"`
}

// NewMessagesList instantiates a new MessagesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesList() *MessagesList {
	this := MessagesList{}
	return &this
}

// NewMessagesListWithDefaults instantiates a new MessagesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesListWithDefaults() *MessagesList {
	this := MessagesList{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *MessagesList) GetTotalCount() int32 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesList) GetTotalCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *MessagesList) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *MessagesList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *MessagesList) GetPageInfo() PageInfo {
	if o == nil || utils.IsNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesList) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || utils.IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *MessagesList) HasPageInfo() bool {
	if o != nil && !utils.IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *MessagesList) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *MessagesList) GetMessages() []ListMessageItem {
	if o == nil || utils.IsNil(o.Messages) {
		var ret []ListMessageItem
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesList) GetMessagesOk() ([]ListMessageItem, bool) {
	if o == nil || utils.IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *MessagesList) HasMessages() bool {
	if o != nil && !utils.IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ListMessageItem and assigns it to the Messages field.
func (o *MessagesList) SetMessages(v []ListMessageItem) {
	o.Messages = v
}

func (o MessagesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !utils.IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if !utils.IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableMessagesList struct {
	value *MessagesList
	isSet bool
}

func (v NullableMessagesList) Get() *MessagesList {
	return v.value
}

func (v *NullableMessagesList) Set(val *MessagesList) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesList) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesList(val *MessagesList) *NullableMessagesList {
	return &NullableMessagesList{value: val, isSet: true}
}

func (v NullableMessagesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


