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

// ListTransactionsByBlockHashRIBSLVout struct for ListTransactionsByBlockHashRIBSLVout
type ListTransactionsByBlockHashRIBSLVout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListTransactionsByBlockHashRIBSLScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListTransactionsByBlockHashRIBSLVout instantiates a new ListTransactionsByBlockHashRIBSLVout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSLVout(isSpent bool, scriptPubKey ListTransactionsByBlockHashRIBSLScriptPubKey, value string) *ListTransactionsByBlockHashRIBSLVout {
	this := ListTransactionsByBlockHashRIBSLVout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListTransactionsByBlockHashRIBSLVoutWithDefaults instantiates a new ListTransactionsByBlockHashRIBSLVout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSLVoutWithDefaults() *ListTransactionsByBlockHashRIBSLVout {
	this := ListTransactionsByBlockHashRIBSLVout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListTransactionsByBlockHashRIBSLVout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSLVout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListTransactionsByBlockHashRIBSLVout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListTransactionsByBlockHashRIBSLVout) GetScriptPubKey() ListTransactionsByBlockHashRIBSLScriptPubKey {
	if o == nil {
		var ret ListTransactionsByBlockHashRIBSLScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSLVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashRIBSLScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListTransactionsByBlockHashRIBSLVout) SetScriptPubKey(v ListTransactionsByBlockHashRIBSLScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashRIBSLVout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSLVout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashRIBSLVout) SetValue(v string) {
	o.Value = v
}

func (o ListTransactionsByBlockHashRIBSLVout) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSLVout struct {
	value *ListTransactionsByBlockHashRIBSLVout
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSLVout) Get() *ListTransactionsByBlockHashRIBSLVout {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSLVout) Set(val *ListTransactionsByBlockHashRIBSLVout) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSLVout) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSLVout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSLVout(val *ListTransactionsByBlockHashRIBSLVout) *NullableListTransactionsByBlockHashRIBSLVout {
	return &NullableListTransactionsByBlockHashRIBSLVout{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSLVout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSLVout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


