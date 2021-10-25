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

// ListConfirmedTransactionsByAddressRIBSD2Vout struct for ListConfirmedTransactionsByAddressRIBSD2Vout
type ListConfirmedTransactionsByAddressRIBSD2Vout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListConfirmedTransactionsByAddressRIBSD2Vout instantiates a new ListConfirmedTransactionsByAddressRIBSD2Vout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSD2Vout(isSpent bool, scriptPubKey ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey, value string) *ListConfirmedTransactionsByAddressRIBSD2Vout {
	this := ListConfirmedTransactionsByAddressRIBSD2Vout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSD2VoutWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSD2Vout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSD2VoutWithDefaults() *ListConfirmedTransactionsByAddressRIBSD2Vout {
	this := ListConfirmedTransactionsByAddressRIBSD2Vout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetScriptPubKey() ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetScriptPubKeyOk() (*ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) SetScriptPubKey(v ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) SetValue(v string) {
	o.Value = v
}

func (o ListConfirmedTransactionsByAddressRIBSD2Vout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isSpent"] = o.IsSpent
	}
	if true {
		toSerialize["scriptPubKey"] = o.ScriptPubKey
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSD2Vout struct {
	value *ListConfirmedTransactionsByAddressRIBSD2Vout
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2Vout) Get() *ListConfirmedTransactionsByAddressRIBSD2Vout {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2Vout) Set(val *ListConfirmedTransactionsByAddressRIBSD2Vout) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2Vout) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2Vout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSD2Vout(val *ListConfirmedTransactionsByAddressRIBSD2Vout) *NullableListConfirmedTransactionsByAddressRIBSD2Vout {
	return &NullableListConfirmedTransactionsByAddressRIBSD2Vout{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSD2Vout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSD2Vout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

