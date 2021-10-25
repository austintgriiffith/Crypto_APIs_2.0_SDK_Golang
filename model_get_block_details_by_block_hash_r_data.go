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

// GetBlockDetailsByBlockHashRData struct for GetBlockDetailsByBlockHashRData
type GetBlockDetailsByBlockHashRData struct {
	Item GetBlockDetailsByBlockHashRI `json:"item"`
}

// NewGetBlockDetailsByBlockHashRData instantiates a new GetBlockDetailsByBlockHashRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashRData(item GetBlockDetailsByBlockHashRI) *GetBlockDetailsByBlockHashRData {
	this := GetBlockDetailsByBlockHashRData{}
	this.Item = item
	return &this
}

// NewGetBlockDetailsByBlockHashRDataWithDefaults instantiates a new GetBlockDetailsByBlockHashRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashRDataWithDefaults() *GetBlockDetailsByBlockHashRData {
	this := GetBlockDetailsByBlockHashRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetBlockDetailsByBlockHashRData) GetItem() GetBlockDetailsByBlockHashRI {
	if o == nil {
		var ret GetBlockDetailsByBlockHashRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRData) GetItemOk() (*GetBlockDetailsByBlockHashRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetBlockDetailsByBlockHashRData) SetItem(v GetBlockDetailsByBlockHashRI) {
	o.Item = v
}

func (o GetBlockDetailsByBlockHashRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashRData struct {
	value *GetBlockDetailsByBlockHashRData
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashRData) Get() *GetBlockDetailsByBlockHashRData {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashRData) Set(val *GetBlockDetailsByBlockHashRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashRData(val *GetBlockDetailsByBlockHashRData) *NullableGetBlockDetailsByBlockHashRData {
	return &NullableGetBlockDetailsByBlockHashRData{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


