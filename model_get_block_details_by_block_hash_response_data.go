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

// GetBlockDetailsByBlockHashResponseData struct for GetBlockDetailsByBlockHashResponseData
type GetBlockDetailsByBlockHashResponseData struct {
	Item GetBlockDetailsByBlockHashResponseItem `json:"item"`
}

// NewGetBlockDetailsByBlockHashResponseData instantiates a new GetBlockDetailsByBlockHashResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashResponseData(item GetBlockDetailsByBlockHashResponseItem) *GetBlockDetailsByBlockHashResponseData {
	this := GetBlockDetailsByBlockHashResponseData{}
	this.Item = item
	return &this
}

// NewGetBlockDetailsByBlockHashResponseDataWithDefaults instantiates a new GetBlockDetailsByBlockHashResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashResponseDataWithDefaults() *GetBlockDetailsByBlockHashResponseData {
	this := GetBlockDetailsByBlockHashResponseData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetBlockDetailsByBlockHashResponseData) GetItem() GetBlockDetailsByBlockHashResponseItem {
	if o == nil {
		var ret GetBlockDetailsByBlockHashResponseItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashResponseData) GetItemOk() (*GetBlockDetailsByBlockHashResponseItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetBlockDetailsByBlockHashResponseData) SetItem(v GetBlockDetailsByBlockHashResponseItem) {
	o.Item = v
}

func (o GetBlockDetailsByBlockHashResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashResponseData struct {
	value *GetBlockDetailsByBlockHashResponseData
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashResponseData) Get() *GetBlockDetailsByBlockHashResponseData {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashResponseData) Set(val *GetBlockDetailsByBlockHashResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashResponseData(val *GetBlockDetailsByBlockHashResponseData) *NullableGetBlockDetailsByBlockHashResponseData {
	return &NullableGetBlockDetailsByBlockHashResponseData{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

