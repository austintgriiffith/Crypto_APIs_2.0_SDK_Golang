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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin struct for GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin
type GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetTransactionDetailsByTransactionIDRIBSD2ScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int32 `json:"sequence"`
	// String representation of the txid
	Txid *string `json:"txid,omitempty"`
	Txinwitness []string `json:"txinwitness"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout *int32 `json:"vout,omitempty"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSD2ScriptSig, sequence int32, txinwitness []string) *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txinwitness = txinwitness
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVinWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSD2ScriptSig {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSD2ScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSD2ScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSD2ScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetSequenceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetSequence(v int32) {
	o.Sequence = v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetTxid(v string) {
	o.Txid = &v
}

// GetTxinwitness returns the Txinwitness field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxinwitness() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txinwitness, true
}

// SetTxinwitness sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetTxinwitness(v []string) {
	o.Txinwitness = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetVout(v int32) {
	o.Vout = &v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


