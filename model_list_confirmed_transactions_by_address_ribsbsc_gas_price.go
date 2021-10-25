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

// ListConfirmedTransactionsByAddressRIBSBSCGasPrice struct for ListConfirmedTransactionsByAddressRIBSBSCGasPrice
type ListConfirmedTransactionsByAddressRIBSBSCGasPrice struct {
	// Represents the price offered to the miner to purchase this amount of gas
	Amount string `json:"amount"`
	// Defines the unit of the gas price amount, e.g. BTC, ETH, XRP.
	Unit string `json:"unit"`
}

// NewListConfirmedTransactionsByAddressRIBSBSCGasPrice instantiates a new ListConfirmedTransactionsByAddressRIBSBSCGasPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSBSCGasPrice(amount string, unit string) *ListConfirmedTransactionsByAddressRIBSBSCGasPrice {
	this := ListConfirmedTransactionsByAddressRIBSBSCGasPrice{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSBSCGasPriceWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSBSCGasPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSBSCGasPriceWithDefaults() *ListConfirmedTransactionsByAddressRIBSBSCGasPrice {
	this := ListConfirmedTransactionsByAddressRIBSBSCGasPrice{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) SetUnit(v string) {
	o.Unit = v
}

func (o ListConfirmedTransactionsByAddressRIBSBSCGasPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice struct {
	value *ListConfirmedTransactionsByAddressRIBSBSCGasPrice
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice) Get() *ListConfirmedTransactionsByAddressRIBSBSCGasPrice {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice) Set(val *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSBSCGasPrice(val *ListConfirmedTransactionsByAddressRIBSBSCGasPrice) *NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice {
	return &NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSBSCGasPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


