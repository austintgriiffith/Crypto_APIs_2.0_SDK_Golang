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

// GenerateReceivingAddressRData struct for GenerateReceivingAddressRData
type GenerateReceivingAddressRData struct {
	Item GenerateReceivingAddressRI `json:"item"`
}

// NewGenerateReceivingAddressRData instantiates a new GenerateReceivingAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateReceivingAddressRData(item GenerateReceivingAddressRI) *GenerateReceivingAddressRData {
	this := GenerateReceivingAddressRData{}
	this.Item = item
	return &this
}

// NewGenerateReceivingAddressRDataWithDefaults instantiates a new GenerateReceivingAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateReceivingAddressRDataWithDefaults() *GenerateReceivingAddressRData {
	this := GenerateReceivingAddressRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GenerateReceivingAddressRData) GetItem() GenerateReceivingAddressRI {
	if o == nil {
		var ret GenerateReceivingAddressRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GenerateReceivingAddressRData) GetItemOk() (*GenerateReceivingAddressRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GenerateReceivingAddressRData) SetItem(v GenerateReceivingAddressRI) {
	o.Item = v
}

func (o GenerateReceivingAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateReceivingAddressRData struct {
	value *GenerateReceivingAddressRData
	isSet bool
}

func (v NullableGenerateReceivingAddressRData) Get() *GenerateReceivingAddressRData {
	return v.value
}

func (v *NullableGenerateReceivingAddressRData) Set(val *GenerateReceivingAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateReceivingAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateReceivingAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateReceivingAddressRData(val *GenerateReceivingAddressRData) *NullableGenerateReceivingAddressRData {
	return &NullableGenerateReceivingAddressRData{value: val, isSet: true}
}

func (v NullableGenerateReceivingAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateReceivingAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

