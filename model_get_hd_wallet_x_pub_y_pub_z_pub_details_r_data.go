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

// GetHDWalletXPubYPubZPubDetailsRData struct for GetHDWalletXPubYPubZPubDetailsRData
type GetHDWalletXPubYPubZPubDetailsRData struct {
	Item GetHDWalletXPubYPubZPubDetailsRI `json:"item"`
}

// NewGetHDWalletXPubYPubZPubDetailsRData instantiates a new GetHDWalletXPubYPubZPubDetailsRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHDWalletXPubYPubZPubDetailsRData(item GetHDWalletXPubYPubZPubDetailsRI) *GetHDWalletXPubYPubZPubDetailsRData {
	this := GetHDWalletXPubYPubZPubDetailsRData{}
	this.Item = item
	return &this
}

// NewGetHDWalletXPubYPubZPubDetailsRDataWithDefaults instantiates a new GetHDWalletXPubYPubZPubDetailsRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHDWalletXPubYPubZPubDetailsRDataWithDefaults() *GetHDWalletXPubYPubZPubDetailsRData {
	this := GetHDWalletXPubYPubZPubDetailsRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetHDWalletXPubYPubZPubDetailsRData) GetItem() GetHDWalletXPubYPubZPubDetailsRI {
	if o == nil {
		var ret GetHDWalletXPubYPubZPubDetailsRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubDetailsRData) GetItemOk() (*GetHDWalletXPubYPubZPubDetailsRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetHDWalletXPubYPubZPubDetailsRData) SetItem(v GetHDWalletXPubYPubZPubDetailsRI) {
	o.Item = v
}

func (o GetHDWalletXPubYPubZPubDetailsRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetHDWalletXPubYPubZPubDetailsRData struct {
	value *GetHDWalletXPubYPubZPubDetailsRData
	isSet bool
}

func (v NullableGetHDWalletXPubYPubZPubDetailsRData) Get() *GetHDWalletXPubYPubZPubDetailsRData {
	return v.value
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsRData) Set(val *GetHDWalletXPubYPubZPubDetailsRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHDWalletXPubYPubZPubDetailsRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHDWalletXPubYPubZPubDetailsRData(val *GetHDWalletXPubYPubZPubDetailsRData) *NullableGetHDWalletXPubYPubZPubDetailsRData {
	return &NullableGetHDWalletXPubYPubZPubDetailsRData{value: val, isSet: true}
}

func (v NullableGetHDWalletXPubYPubZPubDetailsRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHDWalletXPubYPubZPubDetailsRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


