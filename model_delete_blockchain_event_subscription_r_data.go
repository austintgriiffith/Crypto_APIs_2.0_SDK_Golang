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

// DeleteBlockchainEventSubscriptionRData struct for DeleteBlockchainEventSubscriptionRData
type DeleteBlockchainEventSubscriptionRData struct {
	Item DeleteBlockchainEventSubscriptionRI `json:"item"`
}

// NewDeleteBlockchainEventSubscriptionRData instantiates a new DeleteBlockchainEventSubscriptionRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBlockchainEventSubscriptionRData(item DeleteBlockchainEventSubscriptionRI) *DeleteBlockchainEventSubscriptionRData {
	this := DeleteBlockchainEventSubscriptionRData{}
	this.Item = item
	return &this
}

// NewDeleteBlockchainEventSubscriptionRDataWithDefaults instantiates a new DeleteBlockchainEventSubscriptionRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBlockchainEventSubscriptionRDataWithDefaults() *DeleteBlockchainEventSubscriptionRData {
	this := DeleteBlockchainEventSubscriptionRData{}
	return &this
}

// GetItem returns the Item field value
func (o *DeleteBlockchainEventSubscriptionRData) GetItem() DeleteBlockchainEventSubscriptionRI {
	if o == nil {
		var ret DeleteBlockchainEventSubscriptionRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *DeleteBlockchainEventSubscriptionRData) GetItemOk() (*DeleteBlockchainEventSubscriptionRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *DeleteBlockchainEventSubscriptionRData) SetItem(v DeleteBlockchainEventSubscriptionRI) {
	o.Item = v
}

func (o DeleteBlockchainEventSubscriptionRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteBlockchainEventSubscriptionRData struct {
	value *DeleteBlockchainEventSubscriptionRData
	isSet bool
}

func (v NullableDeleteBlockchainEventSubscriptionRData) Get() *DeleteBlockchainEventSubscriptionRData {
	return v.value
}

func (v *NullableDeleteBlockchainEventSubscriptionRData) Set(val *DeleteBlockchainEventSubscriptionRData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBlockchainEventSubscriptionRData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBlockchainEventSubscriptionRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBlockchainEventSubscriptionRData(val *DeleteBlockchainEventSubscriptionRData) *NullableDeleteBlockchainEventSubscriptionRData {
	return &NullableDeleteBlockchainEventSubscriptionRData{value: val, isSet: true}
}

func (v NullableDeleteBlockchainEventSubscriptionRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBlockchainEventSubscriptionRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


