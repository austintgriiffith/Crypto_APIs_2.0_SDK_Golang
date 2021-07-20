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

// NewUnconfirmedTokensTransactionsRBData struct for NewUnconfirmedTokensTransactionsRBData
type NewUnconfirmedTokensTransactionsRBData struct {
	Item NewUnconfirmedTokensTransactionsRBDataItem `json:"item"`
}

// NewNewUnconfirmedTokensTransactionsRBData instantiates a new NewUnconfirmedTokensTransactionsRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUnconfirmedTokensTransactionsRBData(item NewUnconfirmedTokensTransactionsRBDataItem) *NewUnconfirmedTokensTransactionsRBData {
	this := NewUnconfirmedTokensTransactionsRBData{}
	this.Item = item
	return &this
}

// NewNewUnconfirmedTokensTransactionsRBDataWithDefaults instantiates a new NewUnconfirmedTokensTransactionsRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUnconfirmedTokensTransactionsRBDataWithDefaults() *NewUnconfirmedTokensTransactionsRBData {
	this := NewUnconfirmedTokensTransactionsRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewUnconfirmedTokensTransactionsRBData) GetItem() NewUnconfirmedTokensTransactionsRBDataItem {
	if o == nil {
		var ret NewUnconfirmedTokensTransactionsRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewUnconfirmedTokensTransactionsRBData) GetItemOk() (*NewUnconfirmedTokensTransactionsRBDataItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewUnconfirmedTokensTransactionsRBData) SetItem(v NewUnconfirmedTokensTransactionsRBDataItem) {
	o.Item = v
}

func (o NewUnconfirmedTokensTransactionsRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewUnconfirmedTokensTransactionsRBData struct {
	value *NewUnconfirmedTokensTransactionsRBData
	isSet bool
}

func (v NullableNewUnconfirmedTokensTransactionsRBData) Get() *NewUnconfirmedTokensTransactionsRBData {
	return v.value
}

func (v *NullableNewUnconfirmedTokensTransactionsRBData) Set(val *NewUnconfirmedTokensTransactionsRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUnconfirmedTokensTransactionsRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUnconfirmedTokensTransactionsRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUnconfirmedTokensTransactionsRBData(val *NewUnconfirmedTokensTransactionsRBData) *NullableNewUnconfirmedTokensTransactionsRBData {
	return &NullableNewUnconfirmedTokensTransactionsRBData{value: val, isSet: true}
}

func (v NullableNewUnconfirmedTokensTransactionsRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUnconfirmedTokensTransactionsRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


