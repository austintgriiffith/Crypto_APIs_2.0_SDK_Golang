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

// ListConfirmedTransactionsByAddressRIBSZVin struct for ListConfirmedTransactionsByAddressRIBSZVin
type ListConfirmedTransactionsByAddressRIBSZVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase string `json:"coinbase"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSZScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int32 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Defines the specific amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewListConfirmedTransactionsByAddressRIBSZVin instantiates a new ListConfirmedTransactionsByAddressRIBSZVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSZVin(addresses []string, coinbase string, scriptSig GetTransactionDetailsByTransactionIDRIBSZScriptSig, sequence int32, txid string, txinwitness []string, value string, vout int32) *ListConfirmedTransactionsByAddressRIBSZVin {
	this := ListConfirmedTransactionsByAddressRIBSZVin{}
	this.Addresses = addresses
	this.Coinbase = coinbase
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSZVinWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSZVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSZVinWithDefaults() *ListConfirmedTransactionsByAddressRIBSZVin {
	this := ListConfirmedTransactionsByAddressRIBSZVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetCoinbase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetCoinbaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Coinbase, true
}

// SetCoinbase sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetCoinbase(v string) {
	o.Coinbase = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSZScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSZScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSZScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSZScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetSequenceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetSequence(v int32) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetTxidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSZVin) GetVoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListConfirmedTransactionsByAddressRIBSZVin) SetVout(v int32) {
	o.Vout = v
}

func (o ListConfirmedTransactionsByAddressRIBSZVin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if true {
		toSerialize["coinbase"] = o.Coinbase
	}
	if true {
		toSerialize["scriptSig"] = o.ScriptSig
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if true {
		toSerialize["txid"] = o.Txid
	}
	if true {
		toSerialize["txinwitness"] = o.Txinwitness
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSZVin struct {
	value *ListConfirmedTransactionsByAddressRIBSZVin
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSZVin) Get() *ListConfirmedTransactionsByAddressRIBSZVin {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSZVin) Set(val *ListConfirmedTransactionsByAddressRIBSZVin) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSZVin) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSZVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSZVin(val *ListConfirmedTransactionsByAddressRIBSZVin) *NullableListConfirmedTransactionsByAddressRIBSZVin {
	return &NullableListConfirmedTransactionsByAddressRIBSZVin{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSZVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSZVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


