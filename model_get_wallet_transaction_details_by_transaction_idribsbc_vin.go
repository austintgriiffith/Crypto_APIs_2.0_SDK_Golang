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

// GetWalletTransactionDetailsByTransactionIDRIBSBCVin struct for GetWalletTransactionDetailsByTransactionIDRIBSBCVin
type GetWalletTransactionDetailsByTransactionIDRIBSBCVin struct {
	Addresses []string `json:"addresses"`
	// Represents the coinbase hex.
	Coinbase *string `json:"coinbase,omitempty"`
	ScriptSig GetWalletTransactionDetailsByTransactionIDRIBSBCScriptSig `json:"scriptSig"`
	// Represents the script sequence number.
	Sequence int32 `json:"sequence"`
	// Represents the reference transaction identifier.
	Txid string `json:"txid"`
	Txinwitness *[]string `json:"txinwitness,omitempty"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
	// It refers to the index of the output address of this transaction. The index starts from 0.
	Vout *int32 `json:"vout,omitempty"`
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSBCVin instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBCVin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletTransactionDetailsByTransactionIDRIBSBCVin(addresses []string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSBCScriptSig, sequence int32, txid string) *GetWalletTransactionDetailsByTransactionIDRIBSBCVin {
	this := GetWalletTransactionDetailsByTransactionIDRIBSBCVin{}
	this.Addresses = addresses
	this.ScriptSig = scriptSig
	this.Sequence = sequence
	this.Txid = txid
	return &this
}

// NewGetWalletTransactionDetailsByTransactionIDRIBSBCVinWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBCVin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletTransactionDetailsByTransactionIDRIBSBCVinWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSBCVin {
	this := GetWalletTransactionDetailsByTransactionIDRIBSBCVin{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetAddressesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetAddresses(v []string) {
	o.Addresses = v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetCoinbase(v string) {
	o.Coinbase = &v
}

// GetScriptSig returns the ScriptSig field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSBCScriptSig {
	if o == nil {
		var ret GetWalletTransactionDetailsByTransactionIDRIBSBCScriptSig
		return ret
	}

	return o.ScriptSig
}

// GetScriptSigOk returns a tuple with the ScriptSig field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSBCScriptSig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptSig, true
}

// SetScriptSig sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSBCScriptSig) {
	o.ScriptSig = v
}

// GetSequence returns the Sequence field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetSequenceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetSequence(v int32) {
	o.Sequence = v
}

// GetTxid returns the Txid field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetTxidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetTxid(v string) {
	o.Txid = v
}

// GetTxinwitness returns the Txinwitness field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetTxinwitness() []string {
	if o == nil || o.Txinwitness == nil {
		var ret []string
		return ret
	}
	return *o.Txinwitness
}

// GetTxinwitnessOk returns a tuple with the Txinwitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetTxinwitnessOk() (*[]string, bool) {
	if o == nil || o.Txinwitness == nil {
		return nil, false
	}
	return o.Txinwitness, true
}

// HasTxinwitness returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) HasTxinwitness() bool {
	if o != nil && o.Txinwitness != nil {
		return true
	}

	return false
}

// SetTxinwitness gets a reference to the given []string and assigns it to the Txinwitness field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetTxinwitness(v []string) {
	o.Txinwitness = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetValue(v string) {
	o.Value = &v
}

// GetVout returns the Vout field value if set, zero value otherwise.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetVout() int32 {
	if o == nil || o.Vout == nil {
		var ret int32
		return ret
	}
	return *o.Vout
}

// GetVoutOk returns a tuple with the Vout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) GetVoutOk() (*int32, bool) {
	if o == nil || o.Vout == nil {
		return nil, false
	}
	return o.Vout, true
}

// HasVout returns a boolean if a field has been set.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) HasVout() bool {
	if o != nil && o.Vout != nil {
		return true
	}

	return false
}

// SetVout gets a reference to the given int32 and assigns it to the Vout field.
func (o *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) SetVout(v int32) {
	o.Vout = &v
}

func (o GetWalletTransactionDetailsByTransactionIDRIBSBCVin) MarshalJSON() ([]byte, error) {
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
	if o.Vout != nil {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin struct {
	value *GetWalletTransactionDetailsByTransactionIDRIBSBCVin
	isSet bool
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin) Get() *GetWalletTransactionDetailsByTransactionIDRIBSBCVin {
	return v.value
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin) Set(val *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin(val *GetWalletTransactionDetailsByTransactionIDRIBSBCVin) *NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin {
	return &NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin{value: val, isSet: true}
}

func (v NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletTransactionDetailsByTransactionIDRIBSBCVin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


