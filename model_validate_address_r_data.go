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

// ValidateAddressRData struct for ValidateAddressRData
type ValidateAddressRData struct {
	Item ValidateAddressRI `json:"item"`
}

// NewValidateAddressRData instantiates a new ValidateAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateAddressRData(item ValidateAddressRI) *ValidateAddressRData {
	this := ValidateAddressRData{}
	this.Item = item
	return &this
}

// NewValidateAddressRDataWithDefaults instantiates a new ValidateAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateAddressRDataWithDefaults() *ValidateAddressRData {
	this := ValidateAddressRData{}
	return &this
}

// GetItem returns the Item field value
func (o *ValidateAddressRData) GetItem() ValidateAddressRI {
	if o == nil {
		var ret ValidateAddressRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ValidateAddressRData) GetItemOk() (*ValidateAddressRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *ValidateAddressRData) SetItem(v ValidateAddressRI) {
	o.Item = v
}

func (o ValidateAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableValidateAddressRData struct {
	value *ValidateAddressRData
	isSet bool
}

func (v NullableValidateAddressRData) Get() *ValidateAddressRData {
	return v.value
}

func (v *NullableValidateAddressRData) Set(val *ValidateAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateAddressRData(val *ValidateAddressRData) *NullableValidateAddressRData {
	return &NullableValidateAddressRData{value: val, isSet: true}
}

func (v NullableValidateAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

