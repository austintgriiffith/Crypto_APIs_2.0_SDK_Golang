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

// GetTransactionDetailsByTransactionIDRIBSD2Vin struct for GetTransactionDetailsByTransactionIDRIBSD2Vin
type GetTransactionDetailsByTransactionIDRIBSD2Vin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSD2ScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int32 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout *int32 `json:"vout,omitempty"`
}

// NewGetTransactionDetailsByTransactionIDRIBSD2Vin instantiates a new GetTransactionDetailsByTransactionIDRIBSD2Vin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSD2Vin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSD2ScriptSig, sequence int32, txinwitness []string) *GetTransactionDetailsByTransactionIDRIBSD2Vin {
	this := GetTransactionDetailsByTransactionIDRIBSD2Vin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSD2VinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSD2Vin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSD2VinWithDefaults() *GetTransactionDetailsByTransactionIDRIBSD2Vin {
	this := GetTransactionDetailsByTransactionIDRIBSD2Vin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSD2ScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSD2ScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSD2ScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSD2ScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetSequenceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetSequence(v int32) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *GetTransactionDetailsByTransactionIDRIBSD2Vin) SetVout(v int32) {
	o.Vout = &v
}

func (o GetTransactionDetailsByTransactionIDRIBSD2Vin) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Vout != nil {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIBSD2Vin struct {
	value *GetTransactionDetailsByTransactionIDRIBSD2Vin
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSD2Vin) Get() *GetTransactionDetailsByTransactionIDRIBSD2Vin {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSD2Vin) Set(val *GetTransactionDetailsByTransactionIDRIBSD2Vin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSD2Vin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSD2Vin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSD2Vin(val *GetTransactionDetailsByTransactionIDRIBSD2Vin) *NullableGetTransactionDetailsByTransactionIDRIBSD2Vin {
	return &NullableGetTransactionDetailsByTransactionIDRIBSD2Vin{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSD2Vin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSD2Vin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


