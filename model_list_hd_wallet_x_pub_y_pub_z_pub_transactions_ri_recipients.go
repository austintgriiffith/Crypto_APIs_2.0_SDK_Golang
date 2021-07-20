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

// ListHDWalletXPubYPubZPubTransactionsRIRecipients struct for ListHDWalletXPubYPubZPubTransactionsRIRecipients
type ListHDWalletXPubYPubZPubTransactionsRIRecipients struct {
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient.
	Address string `json:"address"`
	// Represents the amount received to this address.
	Amount string `json:"amount"`
	// Defines whether an address is a child address derived from the HD wallet (xPub, yPub, zPub) as boolean.
	IsMember bool `json:"isMember"`
}

// NewListHDWalletXPubYPubZPubTransactionsRIRecipients instantiates a new ListHDWalletXPubYPubZPubTransactionsRIRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHDWalletXPubYPubZPubTransactionsRIRecipients(address string, amount string, isMember bool) *ListHDWalletXPubYPubZPubTransactionsRIRecipients {
	this := ListHDWalletXPubYPubZPubTransactionsRIRecipients{}
	this.Address = address
	this.Amount = amount
	this.IsMember = isMember
	return &this
}

// NewListHDWalletXPubYPubZPubTransactionsRIRecipientsWithDefaults instantiates a new ListHDWalletXPubYPubZPubTransactionsRIRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHDWalletXPubYPubZPubTransactionsRIRecipientsWithDefaults() *ListHDWalletXPubYPubZPubTransactionsRIRecipients {
	this := ListHDWalletXPubYPubZPubTransactionsRIRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) SetAmount(v string) {
	o.Amount = v
}

// GetIsMember returns the IsMember field value
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetIsMember() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMember
}

// GetIsMemberOk returns a tuple with the IsMember field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetIsMemberOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsMember, true
}

// SetIsMember sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) SetIsMember(v bool) {
	o.IsMember = v
}

func (o ListHDWalletXPubYPubZPubTransactionsRIRecipients) MarshalJSON() ([]byte, error) {
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

type NullableListHDWalletXPubYPubZPubTransactionsRIRecipients struct {
	value *ListHDWalletXPubYPubZPubTransactionsRIRecipients
	isSet bool
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRIRecipients) Get() *ListHDWalletXPubYPubZPubTransactionsRIRecipients {
	return v.value
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRIRecipients) Set(val *ListHDWalletXPubYPubZPubTransactionsRIRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRIRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRIRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHDWalletXPubYPubZPubTransactionsRIRecipients(val *ListHDWalletXPubYPubZPubTransactionsRIRecipients) *NullableListHDWalletXPubYPubZPubTransactionsRIRecipients {
	return &NullableListHDWalletXPubYPubZPubTransactionsRIRecipients{value: val, isSet: true}
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRIRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRIRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

