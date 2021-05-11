/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// NewBlockResponseData struct for NewBlockResponseData
type NewBlockResponseData struct {
	Item NewBlockResponseItem `json:"item"`
}

// NewNewBlockResponseData instantiates a new NewBlockResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewBlockResponseData(item NewBlockResponseItem) *NewBlockResponseData {
	this := NewBlockResponseData{}
	this.Item = item
	return &this
}

// NewNewBlockResponseDataWithDefaults instantiates a new NewBlockResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewBlockResponseDataWithDefaults() *NewBlockResponseData {
	this := NewBlockResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewBlockResponseData) GetItem() NewBlockResponseItem {
	if o == nil {
		var ret NewBlockResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewBlockResponseData) GetItemOk() (*NewBlockResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewBlockResponseData) SetItem(v NewBlockResponseItem) {
	o.Item = v
}

func (o NewBlockResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewBlockResponseData struct {
	value *NewBlockResponseData
	isSet bool
}

func (v NullableNewBlockResponseData) Get() *NewBlockResponseData {
	return v.value
}

func (v *NullableNewBlockResponseData) Set(val *NewBlockResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockResponseData(val *NewBlockResponseData) *NullableNewBlockResponseData {
	return &NullableNewBlockResponseData{value: val, isSet: true}
}

func (v NullableNewBlockResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

