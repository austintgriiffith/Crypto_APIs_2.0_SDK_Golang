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

// CreateAutomaticCoinsForwardingRBData struct for CreateAutomaticCoinsForwardingRBData
type CreateAutomaticCoinsForwardingRBData struct {
	Item CreateAutomaticCoinsForwardingRBDataItem `json:"item"`
}

// NewCreateAutomaticCoinsForwardingRBData instantiates a new CreateAutomaticCoinsForwardingRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticCoinsForwardingRBData(item CreateAutomaticCoinsForwardingRBDataItem) *CreateAutomaticCoinsForwardingRBData {
	this := CreateAutomaticCoinsForwardingRBData{}
	this.Item = item
	return &this
}

// NewCreateAutomaticCoinsForwardingRBDataWithDefaults instantiates a new CreateAutomaticCoinsForwardingRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticCoinsForwardingRBDataWithDefaults() *CreateAutomaticCoinsForwardingRBData {
	this := CreateAutomaticCoinsForwardingRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *CreateAutomaticCoinsForwardingRBData) GetItem() CreateAutomaticCoinsForwardingRBDataItem {
	if o == nil {
		var ret CreateAutomaticCoinsForwardingRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRBData) GetItemOk() (*CreateAutomaticCoinsForwardingRBDataItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CreateAutomaticCoinsForwardingRBData) SetItem(v CreateAutomaticCoinsForwardingRBDataItem) {
	o.Item = v
}

func (o CreateAutomaticCoinsForwardingRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticCoinsForwardingRBData struct {
	value *CreateAutomaticCoinsForwardingRBData
	isSet bool
}

func (v NullableCreateAutomaticCoinsForwardingRBData) Get() *CreateAutomaticCoinsForwardingRBData {
	return v.value
}

func (v *NullableCreateAutomaticCoinsForwardingRBData) Set(val *CreateAutomaticCoinsForwardingRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticCoinsForwardingRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticCoinsForwardingRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticCoinsForwardingRBData(val *CreateAutomaticCoinsForwardingRBData) *NullableCreateAutomaticCoinsForwardingRBData {
	return &NullableCreateAutomaticCoinsForwardingRBData{value: val, isSet: true}
}

func (v NullableCreateAutomaticCoinsForwardingRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticCoinsForwardingRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

