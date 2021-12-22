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

// GetWalletTransactionDetailsByTransactionIDRIBSBVin struct for GetWalletTransactionDetailsByTransactionIDRIBSBVin
type GetWalletTransactionDetailsByTransactionIDRIBSBVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int32 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness *[]string `json:"txinwitness,omitempty"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout int32 `json:"vout"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSBVin instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSBVin(addresses []string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig, sequence int32, txid string, vout int32) *GetWalletTransactionDetailsByTransactionIDRIBSBVin {
	this := GetWalletTransactionDetailsByTransactionIDRIBSBVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	this.Vout = vout
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSBVinWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSBVinWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSBVin {
	this := GetWalletTransactionDetailsByTransactionIDRIBSBVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig {
	if o == nil {
		var ret GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetSequenceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetSequence(v int32) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxinwitness() []string {
	if o == nil || o.Txinwitness == nil {
		var ret []string
		return ret
	}
	return *o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil || o.Txinwitness == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// HasTxinwitness returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) HasTxinwitness() bool {
	if o != nil && o.Txinwitness != nil {
		return true
	}

	return false
}

// SetTxinwitness gets a reference to the given []string and assigns it to the Txinwitness field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetTxinwitness(v []string) {
	o.Txinwitness = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetVoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetVout(v int32) {
	o.Vout = v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSBVin) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["txid"] = o.Txid
	}
	if o.Txinwitness != nil {
		toSerialize["txinwitness"] = o.Txinwitness
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSBVin
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin) Get() *GetWalletTransactionDetailsByTransactionIDRIBSBVin {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSBVin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSBVin(val *GetWalletTransactionDetailsByTransactionIDRIBSBVin) *NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSBVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


