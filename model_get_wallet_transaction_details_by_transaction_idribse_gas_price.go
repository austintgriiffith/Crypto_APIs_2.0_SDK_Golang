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

// GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice struct for GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice
type GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice struct {
	// Represents the price offered to the miner to purchase this amount of gas.
	Amount string `json:"amount"`
	// Defines the unit of the gas price amount, e.g. BTC, ETH, XRP.
	Unit string `json:"unit"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice(amount string, unit string) *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice {
	this := GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSEGasPriceWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSEGasPriceWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice {
	this := GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) SetUnit(v string) {
	o.Unit = v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) Get() *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice(val *GetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) *NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSEGasPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

