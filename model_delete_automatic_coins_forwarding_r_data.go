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

// DeleteAutomaticCoinsForwardingRData struct for DeleteAutomaticCoinsForwardingRData
type DeleteAutomaticCoinsForwardingRData struct {
	Item DeleteAutomaticCoinsForwardingRI `json:"item"`
}

// NewDeleteAutomaticCoinsForwardingRData instantiates a new DeleteAutomaticCoinsForwardingRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAutomaticCoinsForwardingRData(item DeleteAutomaticCoinsForwardingRI) *DeleteAutomaticCoinsForwardingRData {
	this := DeleteAutomaticCoinsForwardingRData{}
	this.Item = item
	return &this
}

// NewDeleteAutomaticCoinsForwardingRDataWithDefaults instantiates a new DeleteAutomaticCoinsForwardingRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAutomaticCoinsForwardingRDataWithDefaults() *DeleteAutomaticCoinsForwardingRData {
	this := DeleteAutomaticCoinsForwardingRData{}
	return &this
}

// GetItem returns the Item field value
func (o *DeleteAutomaticCoinsForwardingRData) GetItem() DeleteAutomaticCoinsForwardingRI {
	if o == nil {
		var ret DeleteAutomaticCoinsForwardingRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DeleteAutomaticCoinsForwardingRData) GetItemOk() (*DeleteAutomaticCoinsForwardingRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DeleteAutomaticCoinsForwardingRData) SetItem(v DeleteAutomaticCoinsForwardingRI) {
	o.Item = v
}

func (o DeleteAutomaticCoinsForwardingRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAutomaticCoinsForwardingRData struct {
	value *DeleteAutomaticCoinsForwardingRData
	isSet bool
}

func (v NullableDeleteAutomaticCoinsForwardingRData) Get() *DeleteAutomaticCoinsForwardingRData {
	return v.value
}

func (v *NullableDeleteAutomaticCoinsForwardingRData) Set(val *DeleteAutomaticCoinsForwardingRData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomaticCoinsForwardingRData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomaticCoinsForwardingRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomaticCoinsForwardingRData(val *DeleteAutomaticCoinsForwardingRData) *NullableDeleteAutomaticCoinsForwardingRData {
	return &NullableDeleteAutomaticCoinsForwardingRData{value: val, isSet: true}
}

func (v NullableDeleteAutomaticCoinsForwardingRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomaticCoinsForwardingRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


