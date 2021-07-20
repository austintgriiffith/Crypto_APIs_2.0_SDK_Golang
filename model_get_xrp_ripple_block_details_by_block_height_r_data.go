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

// GetXRPRippleBlockDetailsByBlockHeightRData struct for GetXRPRippleBlockDetailsByBlockHeightRData
type GetXRPRippleBlockDetailsByBlockHeightRData struct {
	Item GetXRPRippleBlockDetailsByBlockHeightRI `json:"item"`
}

// NewGetXRPRippleBlockDetailsByBlockHeightRData instantiates a new GetXRPRippleBlockDetailsByBlockHeightRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleBlockDetailsByBlockHeightRData(item GetXRPRippleBlockDetailsByBlockHeightRI) *GetXRPRippleBlockDetailsByBlockHeightRData {
	this := GetXRPRippleBlockDetailsByBlockHeightRData{}
	this.Item = item
	return &this
}

// NewGetXRPRippleBlockDetailsByBlockHeightRDataWithDefaults instantiates a new GetXRPRippleBlockDetailsByBlockHeightRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleBlockDetailsByBlockHeightRDataWithDefaults() *GetXRPRippleBlockDetailsByBlockHeightRData {
	this := GetXRPRippleBlockDetailsByBlockHeightRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRData) GetItem() GetXRPRippleBlockDetailsByBlockHeightRI {
	if o == nil {
		var ret GetXRPRippleBlockDetailsByBlockHeightRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHeightRData) GetItemOk() (*GetXRPRippleBlockDetailsByBlockHeightRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRData) SetItem(v GetXRPRippleBlockDetailsByBlockHeightRI) {
	o.Item = v
}

func (o GetXRPRippleBlockDetailsByBlockHeightRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleBlockDetailsByBlockHeightRData struct {
	value *GetXRPRippleBlockDetailsByBlockHeightRData
	isSet bool
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRData) Get() *GetXRPRippleBlockDetailsByBlockHeightRData {
	return v.value
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRData) Set(val *GetXRPRippleBlockDetailsByBlockHeightRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleBlockDetailsByBlockHeightRData(val *GetXRPRippleBlockDetailsByBlockHeightRData) *NullableGetXRPRippleBlockDetailsByBlockHeightRData {
	return &NullableGetXRPRippleBlockDetailsByBlockHeightRData{value: val, isSet: true}
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


