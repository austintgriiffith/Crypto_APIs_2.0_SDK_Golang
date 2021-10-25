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

// CreateCoinsTransactionRequestFromWalletRData struct for CreateCoinsTransactionRequestFromWalletRData
type CreateCoinsTransactionRequestFromWalletRData struct {
	Item CreateCoinsTransactionRequestFromWalletRI `json:"item"`
}

// NewCreateCoinsTransactionRequestFromWalletRData instantiates a new CreateCoinsTransactionRequestFromWalletRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionRequestFromWalletRData(item CreateCoinsTransactionRequestFromWalletRI) *CreateCoinsTransactionRequestFromWalletRData {
	this := CreateCoinsTransactionRequestFromWalletRData{}
	this.Item = item
	return &this
}

// NewCreateCoinsTransactionRequestFromWalletRDataWithDefaults instantiates a new CreateCoinsTransactionRequestFromWalletRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionRequestFromWalletRDataWithDefaults() *CreateCoinsTransactionRequestFromWalletRData {
	this := CreateCoinsTransactionRequestFromWalletRData{}
	return &this
}

// GetItem returns the Item field value
func (o *CreateCoinsTransactionRequestFromWalletRData) GetItem() CreateCoinsTransactionRequestFromWalletRI {
	if o == nil {
		var ret CreateCoinsTransactionRequestFromWalletRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromWalletRData) GetItemOk() (*CreateCoinsTransactionRequestFromWalletRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CreateCoinsTransactionRequestFromWalletRData) SetItem(v CreateCoinsTransactionRequestFromWalletRI) {
	o.Item = v
}

func (o CreateCoinsTransactionRequestFromWalletRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionRequestFromWalletRData struct {
	value *CreateCoinsTransactionRequestFromWalletRData
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromWalletRData) Get() *CreateCoinsTransactionRequestFromWalletRData {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromWalletRData) Set(val *CreateCoinsTransactionRequestFromWalletRData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromWalletRData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromWalletRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromWalletRData(val *CreateCoinsTransactionRequestFromWalletRData) *NullableCreateCoinsTransactionRequestFromWalletRData {
	return &NullableCreateCoinsTransactionRequestFromWalletRData{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromWalletRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromWalletRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


