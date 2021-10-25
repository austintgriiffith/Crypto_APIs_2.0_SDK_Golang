/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GenerateDepositAddressRData struct for GenerateDepositAddressRData
type GenerateDepositAddressRData struct {
	Item GenerateDepositAddressRI `json:"item"`
}

// NewGenerateDepositAddressRData instantiates a new GenerateDepositAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDepositAddressRData(item GenerateDepositAddressRI) *GenerateDepositAddressRData {
	this := GenerateDepositAddressRData{}
	this.Item = item
	return &this
}

// NewGenerateDepositAddressRDataWithDefaults instantiates a new GenerateDepositAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDepositAddressRDataWithDefaults() *GenerateDepositAddressRData {
	this := GenerateDepositAddressRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GenerateDepositAddressRData) GetItem() GenerateDepositAddressRI {
	if o == nil {
		var ret GenerateDepositAddressRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GenerateDepositAddressRData) GetItemOk() (*GenerateDepositAddressRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GenerateDepositAddressRData) SetItem(v GenerateDepositAddressRI) {
	o.Item = v
}

func (o GenerateDepositAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateDepositAddressRData struct {
	value *GenerateDepositAddressRData
	isSet bool
}

func (v NullableGenerateDepositAddressRData) Get() *GenerateDepositAddressRData {
	return v.value
}

func (v *NullableGenerateDepositAddressRData) Set(val *GenerateDepositAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDepositAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDepositAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDepositAddressRData(val *GenerateDepositAddressRData) *NullableGenerateDepositAddressRData {
	return &NullableGenerateDepositAddressRData{value: val, isSet: true}
}

func (v NullableGenerateDepositAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDepositAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


