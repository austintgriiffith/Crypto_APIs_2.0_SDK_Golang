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

// ListTransactionsByBlockHashRIBSD2Vout struct for ListTransactionsByBlockHashRIBSD2Vout
type ListTransactionsByBlockHashRIBSD2Vout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListTransactionsByBlockHashRIBSD2ScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListTransactionsByBlockHashRIBSD2Vout instantiates a new ListTransactionsByBlockHashRIBSD2Vout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSD2Vout(isSpent bool, scriptPubKey ListTransactionsByBlockHashRIBSD2ScriptPubKey, value string) *ListTransactionsByBlockHashRIBSD2Vout {
	this := ListTransactionsByBlockHashRIBSD2Vout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListTransactionsByBlockHashRIBSD2VoutWithDefaults instantiates a new ListTransactionsByBlockHashRIBSD2Vout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSD2VoutWithDefaults() *ListTransactionsByBlockHashRIBSD2Vout {
	this := ListTransactionsByBlockHashRIBSD2Vout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListTransactionsByBlockHashRIBSD2Vout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2Vout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListTransactionsByBlockHashRIBSD2Vout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListTransactionsByBlockHashRIBSD2Vout) GetScriptPubKey() ListTransactionsByBlockHashRIBSD2ScriptPubKey {
	if o == nil {
		var ret ListTransactionsByBlockHashRIBSD2ScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2Vout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashRIBSD2ScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListTransactionsByBlockHashRIBSD2Vout) SetScriptPubKey(v ListTransactionsByBlockHashRIBSD2ScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashRIBSD2Vout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD2Vout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashRIBSD2Vout) SetValue(v string) {
	o.Value = v
}

func (o ListTransactionsByBlockHashRIBSD2Vout) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSD2Vout struct {
	value *ListTransactionsByBlockHashRIBSD2Vout
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSD2Vout) Get() *ListTransactionsByBlockHashRIBSD2Vout {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSD2Vout) Set(val *ListTransactionsByBlockHashRIBSD2Vout) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSD2Vout) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSD2Vout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSD2Vout(val *ListTransactionsByBlockHashRIBSD2Vout) *NullableListTransactionsByBlockHashRIBSD2Vout {
	return &NullableListTransactionsByBlockHashRIBSD2Vout{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSD2Vout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSD2Vout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

