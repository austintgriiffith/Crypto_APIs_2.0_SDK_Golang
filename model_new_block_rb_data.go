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

// NewBlockRBData struct for NewBlockRBData
type NewBlockRBData struct {
	Item NewBlockRBDataItem `json:"item"`
}

// NewNewBlockRBData instantiates a new NewBlockRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewBlockRBData(item NewBlockRBDataItem) *NewBlockRBData {
	this := NewBlockRBData{}
	this.Item = item
	return &this
}

// NewNewBlockRBDataWithDefaults instantiates a new NewBlockRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewBlockRBDataWithDefaults() *NewBlockRBData {
	this := NewBlockRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewBlockRBData) GetItem() NewBlockRBDataItem {
	if o == nil {
		var ret NewBlockRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewBlockRBData) GetItemOk() (*NewBlockRBDataItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewBlockRBData) SetItem(v NewBlockRBDataItem) {
	o.Item = v
}

func (o NewBlockRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewBlockRBData struct {
	value *NewBlockRBData
	isSet bool
}

func (v NullableNewBlockRBData) Get() *NewBlockRBData {
	return v.value
}

func (v *NullableNewBlockRBData) Set(val *NewBlockRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockRBData(val *NewBlockRBData) *NullableNewBlockRBData {
	return &NullableNewBlockRBData{value: val, isSet: true}
}

func (v NullableNewBlockRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

