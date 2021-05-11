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

// AddressCoinsTransactionConfirmedEachConfirmationData Specifies all data, as attributes, included into the callback notification, which depends on the `event`.
type AddressCoinsTransactionConfirmedEachConfirmationData struct {
	// Represents the Crypto APIs 2.0 product which sends the callback.
	Product string `json:"product"`
	// Defines the specific event, for which a callback subscription is set.
	Event string `json:"event"`
	Item AddressCoinsTransactionConfirmedEachConfirmationDataItem `json:"item"`
}

// NewAddressCoinsTransactionConfirmedEachConfirmationData instantiates a new AddressCoinsTransactionConfirmedEachConfirmationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCoinsTransactionConfirmedEachConfirmationData(product string, event string, item AddressCoinsTransactionConfirmedEachConfirmationDataItem) *AddressCoinsTransactionConfirmedEachConfirmationData {
	this := AddressCoinsTransactionConfirmedEachConfirmationData{}
	this.Product = product
	this.Event = event
	this.Item = item
	return &this
}

// NewAddressCoinsTransactionConfirmedEachConfirmationDataWithDefaults instantiates a new AddressCoinsTransactionConfirmedEachConfirmationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCoinsTransactionConfirmedEachConfirmationDataWithDefaults() *AddressCoinsTransactionConfirmedEachConfirmationData {
	this := AddressCoinsTransactionConfirmedEachConfirmationData{}
	return &this
}

// GetProduct returns the Product field value
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) GetProductOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) SetProduct(v string) {
	o.Product = v
}

// GetEvent returns the Event field value
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) GetEventOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) SetEvent(v string) {
	o.Event = v
}

// GetItem returns the Item field value
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) GetItem() AddressCoinsTransactionConfirmedEachConfirmationDataItem {
	if o == nil {
		var ret AddressCoinsTransactionConfirmedEachConfirmationDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) GetItemOk() (*AddressCoinsTransactionConfirmedEachConfirmationDataItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *AddressCoinsTransactionConfirmedEachConfirmationData) SetItem(v AddressCoinsTransactionConfirmedEachConfirmationDataItem) {
	o.Item = v
}

func (o AddressCoinsTransactionConfirmedEachConfirmationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["product"] = o.Product
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableAddressCoinsTransactionConfirmedEachConfirmationData struct {
	value *AddressCoinsTransactionConfirmedEachConfirmationData
	isSet bool
}

func (v NullableAddressCoinsTransactionConfirmedEachConfirmationData) Get() *AddressCoinsTransactionConfirmedEachConfirmationData {
	return v.value
}

func (v *NullableAddressCoinsTransactionConfirmedEachConfirmationData) Set(val *AddressCoinsTransactionConfirmedEachConfirmationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCoinsTransactionConfirmedEachConfirmationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCoinsTransactionConfirmedEachConfirmationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCoinsTransactionConfirmedEachConfirmationData(val *AddressCoinsTransactionConfirmedEachConfirmationData) *NullableAddressCoinsTransactionConfirmedEachConfirmationData {
	return &NullableAddressCoinsTransactionConfirmedEachConfirmationData{value: val, isSet: true}
}

func (v NullableAddressCoinsTransactionConfirmedEachConfirmationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCoinsTransactionConfirmedEachConfirmationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

