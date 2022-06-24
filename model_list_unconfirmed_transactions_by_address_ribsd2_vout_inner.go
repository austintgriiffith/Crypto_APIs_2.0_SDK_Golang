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

// ListUnconfirmedTransactionsByAddressRIBSD2VoutInner struct for ListUnconfirmedTransactionsByAddressRIBSD2VoutInner
type ListUnconfirmedTransactionsByAddressRIBSD2VoutInner struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey `json:"scriptPubKey"`
	// String representation of the amount
	Value string `json:"value"`
}

// NewListUnconfirmedTransactionsByAddressRIBSD2VoutInner instantiates a new ListUnconfirmedTransactionsByAddressRIBSD2VoutInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSD2VoutInner(isSpent bool, scriptPubKey ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey, value string) *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner {
	this := ListUnconfirmedTransactionsByAddressRIBSD2VoutInner{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSD2VoutInnerWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSD2VoutInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSD2VoutInnerWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner {
	this := ListUnconfirmedTransactionsByAddressRIBSD2VoutInner{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) GetIsSpentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) GetScriptPubKey() ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) GetScriptPubKeyOk() (*ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) SetScriptPubKey(v ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) SetValue(v string) {
	o.Value = v
}

func (o ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) MarshalJSON() ([]byte, error) {
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

type NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner struct {
	value *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner) Get() *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner) Set(val *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner(val *ListUnconfirmedTransactionsByAddressRIBSD2VoutInner) *NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner {
	return &NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD2VoutInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

