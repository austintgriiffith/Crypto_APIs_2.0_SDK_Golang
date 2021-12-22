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

// GetWalletTransactionDetailsByTransactionIDRData struct for GetWalletTransactionDetailsByTransactionIDRData
type GetWalletTransactionDetailsByTransactionIDRData struct {
	Item GetWalletTransactionDetailsByTransactionIDRI `json:"item"`
}

// NewGetWalletTransactionDetailsByTransactionIDRData instantiates a new GetWalletTransactionDetailsByTransactionIDRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRData(item GetWalletTransactionDetailsByTransactionIDRI) *GetWalletTransactionDetailsByTransactionIDRData {
	this := GetWalletTransactionDetailsByTransactionIDRData{}
	this.Item = item
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRDataWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRDataWithDefaults() *GetWalletTransactionDetailsByTransactionIDRData {
	this := GetWalletTransactionDetailsByTransactionIDRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetWalletTransactionDetailsByTransactionIDRData) GetItem() GetWalletTransactionDetailsByTransactionIDRI {
	if o == nil {
		var ret GetWalletTransactionDetailsByTransactionIDRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRData) GetItemOk() (*GetWalletTransactionDetailsByTransactionIDRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRData) SetItem(v GetWalletTransactionDetailsByTransactionIDRI) {
	o.Item = v
}

func (o GetWalletTransactionDetailsByTransactionIDRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletTransactionDetailsByTransactionIDRData struct {
	value *GetWalletTransactionDetailsByTransactionIDRData
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRData) Get() *GetWalletTransactionDetailsByTransactionIDRData {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRData) Set(val *GetWalletTransactionDetailsByTransactionIDRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRData(val *GetWalletTransactionDetailsByTransactionIDRData) *NullableGetWalletTransactionDetailsByTransactionIDRData {
	return &NullableGetWalletTransactionDetailsByTransactionIDRData{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

