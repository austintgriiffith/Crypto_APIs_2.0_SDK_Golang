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

// ListAllUnconfirmedTransactionsRIBSLVin struct for ListAllUnconfirmedTransactionsRIBSLVin
type ListAllUnconfirmedTransactionsRIBSLVin struct {
	Addresses []string `json:"addresses"`
	ScriptSig ListConfirmedTransactionsByAddressRIBSLScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// Defines the vout of the transaction output, i.e. which output to spend.
	Vout int32 `json:"vout"`
}

// NewListAllUnconfirmedTransactionsRIBSLVin instantiates a new ListAllUnconfirmedTransactionsRIBSLVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllUnconfirmedTransactionsRIBSLVin(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSLScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32) *ListAllUnconfirmedTransactionsRIBSLVin {
	this := ListAllUnconfirmedTransactionsRIBSLVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewListAllUnconfirmedTransactionsRIBSLVinWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSLVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllUnconfirmedTransactionsRIBSLVinWithDefaults() *ListAllUnconfirmedTransactionsRIBSLVin {
	this := ListAllUnconfirmedTransactionsRIBSLVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSLScriptSig {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBSLScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSLScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSLScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetSequenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetVoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetVout(v int32) {
	o.Vout = v
}

func (o ListAllUnconfirmedTransactionsRIBSLVin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
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

type NullableListAllUnconfirmedTransactionsRIBSLVin struct {
	value *ListAllUnconfirmedTransactionsRIBSLVin
	isSet bool
}

func (v NullableListAllUnconfirmedTransactionsRIBSLVin) Get() *ListAllUnconfirmedTransactionsRIBSLVin {
	return v.value
}

func (v *NullableListAllUnconfirmedTransactionsRIBSLVin) Set(val *ListAllUnconfirmedTransactionsRIBSLVin) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllUnconfirmedTransactionsRIBSLVin) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllUnconfirmedTransactionsRIBSLVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllUnconfirmedTransactionsRIBSLVin(val *ListAllUnconfirmedTransactionsRIBSLVin) *NullableListAllUnconfirmedTransactionsRIBSLVin {
	return &NullableListAllUnconfirmedTransactionsRIBSLVin{value: val, isSet: true}
}

func (v NullableListAllUnconfirmedTransactionsRIBSLVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllUnconfirmedTransactionsRIBSLVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


