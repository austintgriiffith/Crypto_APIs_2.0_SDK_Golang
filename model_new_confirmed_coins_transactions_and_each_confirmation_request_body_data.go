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

// NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData struct for NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData
type NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData struct {
	Item NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem `json:"item"`
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData(item NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData{}
	this.Item = item
	return &this
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataWithDefaults instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataWithDefaults() *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData{}
	return &this
}

// GetItem returns the Item field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) GetItem() NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem {
	if o == nil {
		var ret NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) GetItemOk() (*NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) SetItem(v NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) {
	o.Item = v
}

func (o NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData struct {
	value *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) Get() *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) Set(val *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData(val *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData {
	return &NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

