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

// ListTransactionsByBlockHeightRIBSD2VinInner struct for ListTransactionsByBlockHeightRIBSD2VinInner
type ListTransactionsByBlockHeightRIBSD2VinInner struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence string `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value string `json:"value"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewListTransactionsByBlockHeightRIBSD2VinInner instantiates a new ListTransactionsByBlockHeightRIBSD2VinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHeightRIBSD2VinInner(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig, sequence string, txinwitness []string, value string, vout int32) *ListTransactionsByBlockHeightRIBSD2VinInner {
	this := ListTransactionsByBlockHeightRIBSD2VinInner{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	this.Value = value
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHeightRIBSD2VinInnerWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSD2VinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHeightRIBSD2VinInnerWithDefaults() *ListTransactionsByBlockHeightRIBSD2VinInner {
	this := ListTransactionsByBlockHeightRIBSD2VinInner{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSDVinInnerScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetSequenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetSequence(v string) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetTxinwitnessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetValue(v string) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHeightRIBSD2VinInner) SetVout(v int32) {
	o.Vout = v
}

func (o ListTransactionsByBlockHeightRIBSD2VinInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Coinbase != nil {
		toSerialize["coinbase"] = o.Coinbase
	}
	if true {
		toSerialize["scriptSig"] = o.ScriptSig
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Txid != nil {
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

type NullableListTransactionsByBlockHeightRIBSD2VinInner struct {
	value *ListTransactionsByBlockHeightRIBSD2VinInner
	isSet bool
}

func (v NullableListTransactionsByBlockHeightRIBSD2VinInner) Get() *ListTransactionsByBlockHeightRIBSD2VinInner {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightRIBSD2VinInner) Set(val *ListTransactionsByBlockHeightRIBSD2VinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightRIBSD2VinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightRIBSD2VinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightRIBSD2VinInner(val *ListTransactionsByBlockHeightRIBSD2VinInner) *NullableListTransactionsByBlockHeightRIBSD2VinInner {
	return &NullableListTransactionsByBlockHeightRIBSD2VinInner{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightRIBSD2VinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightRIBSD2VinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

