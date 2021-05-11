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

// ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout struct for ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout
type ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificDashScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout(isSpent bool, scriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificDashScriptPubKey, value string) *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout {
	this := ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashVoutWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashVoutWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout {
	this := ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) GetScriptPubKey() ListTransactionsByBlockHashResponseItemBlockchainSpecificDashScriptPubKey {
	if o == nil {
		var ret ListTransactionsByBlockHashResponseItemBlockchainSpecificDashScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecificDashScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) SetScriptPubKey(v ListTransactionsByBlockHashResponseItemBlockchainSpecificDashScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) SetValue(v string) {
	o.Value = v
}

func (o ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout struct {
	value *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout
	isSet bool
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) Get() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout {
	return v.value
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) Set(val *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout(val *ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout {
	return &NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

