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

// ListTransactionsByBlockHashRIBSZVin struct for ListTransactionsByBlockHashRIBSZVin
type ListTransactionsByBlockHashRIBSZVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase string `json:"coinbase"`
	ScriptSig ListTransactionsByBlockHashRIBSZScriptSig `json:"scriptSig"`
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

// NewListTransactionsByBlockHashRIBSZVin instantiates a new ListTransactionsByBlockHashRIBSZVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSZVin(addresses []string, coinbase string, scriptSig ListTransactionsByBlockHashRIBSZScriptSig, sequence int32, txid string, txinwitness []string, value string, vout int32) *ListTransactionsByBlockHashRIBSZVin {
	this := ListTransactionsByBlockHashRIBSZVin{}
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

// NewListTransactionsByBlockHashRIBSZVinWithDefaults instantiates a new ListTransactionsByBlockHashRIBSZVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSZVinWithDefaults() *ListTransactionsByBlockHashRIBSZVin {
	this := ListTransactionsByBlockHashRIBSZVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetCoinbase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetCoinbaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Coinbase, true
}

// SetCoinbase sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetCoinbase(v string) {
	o.Coinbase = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetScriptSig() ListTransactionsByBlockHashRIBSZScriptSig {
	if o == nil {
		var ret ListTransactionsByBlockHashRIBSZScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetScriptSigOk() (*ListTransactionsByBlockHashRIBSZScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetScriptSig(v ListTransactionsByBlockHashRIBSZScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetSequenceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetSequence(v int32) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetTxidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHashRIBSZVin) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSZVin) GetVoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHashRIBSZVin) SetVout(v int32) {
	o.Vout = v
}

func (o ListTransactionsByBlockHashRIBSZVin) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSZVin struct {
	value *ListTransactionsByBlockHashRIBSZVin
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSZVin) Get() *ListTransactionsByBlockHashRIBSZVin {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSZVin) Set(val *ListTransactionsByBlockHashRIBSZVin) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSZVin) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSZVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSZVin(val *ListTransactionsByBlockHashRIBSZVin) *NullableListTransactionsByBlockHashRIBSZVin {
	return &NullableListTransactionsByBlockHashRIBSZVin{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSZVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSZVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


