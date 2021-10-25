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

// ListHDWalletXPubYPubZPubTransactionsRISenders struct for ListHDWalletXPubYPubZPubTransactionsRISenders
type ListHDWalletXPubYPubZPubTransactionsRISenders struct {
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Address string `json:"address"`
	// Represents the amount sent by this address.
	Amount string `json:"amount"`
	// Defines whether an address is a child address derived from the HD wallet (xPub, yPub, zPub) as boolean.
	IsMember bool `json:"isMember"`
}

// NewListHDWalletXPubYPubZPubTransactionsRISenders instantiates a new ListHDWalletXPubYPubZPubTransactionsRISenders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHDWalletXPubYPubZPubTransactionsRISenders(address string, amount string, isMember bool) *ListHDWalletXPubYPubZPubTransactionsRISenders {
	this := ListHDWalletXPubYPubZPubTransactionsRISenders{}
	this.Address = address
	this.Amount = amount
	this.IsMember = isMember
	return &this
}

// NewListHDWalletXPubYPubZPubTransactionsRISendersWithDefaults instantiates a new ListHDWalletXPubYPubZPubTransactionsRISenders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHDWalletXPubYPubZPubTransactionsRISendersWithDefaults() *ListHDWalletXPubYPubZPubTransactionsRISenders {
	this := ListHDWalletXPubYPubZPubTransactionsRISenders{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) SetAmount(v string) {
	o.Amount = v
}

// GetIsMember returns the IsMember field value
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) GetIsMember() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMember
}

// GetIsMemberOk returns a tuple with the IsMember field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) GetIsMemberOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsMember, true
}

// SetIsMember sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRISenders) SetIsMember(v bool) {
	o.IsMember = v
}

func (o ListHDWalletXPubYPubZPubTransactionsRISenders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["isMember"] = o.IsMember
	}
	return json.Marshal(toSerialize)
}

type NullableListHDWalletXPubYPubZPubTransactionsRISenders struct {
	value *ListHDWalletXPubYPubZPubTransactionsRISenders
	isSet bool
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRISenders) Get() *ListHDWalletXPubYPubZPubTransactionsRISenders {
	return v.value
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRISenders) Set(val *ListHDWalletXPubYPubZPubTransactionsRISenders) {
	v.value = val
	v.isSet = true
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRISenders) IsSet() bool {
	return v.isSet
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRISenders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHDWalletXPubYPubZPubTransactionsRISenders(val *ListHDWalletXPubYPubZPubTransactionsRISenders) *NullableListHDWalletXPubYPubZPubTransactionsRISenders {
	return &NullableListHDWalletXPubYPubZPubTransactionsRISenders{value: val, isSet: true}
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRISenders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRISenders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


