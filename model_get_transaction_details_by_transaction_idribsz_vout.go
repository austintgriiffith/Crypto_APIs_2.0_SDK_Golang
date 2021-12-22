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

// GetTransactionDetailsByTransactionIDRIBSZVout struct for GetTransactionDetailsByTransactionIDRIBSZVout
type GetTransactionDetailsByTransactionIDRIBSZVout struct {
	// Defines whether the transaction output has been spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey GetTransactionDetailsByTransactionIDRIBSZScriptPubKey `json:"scriptPubKey"`
	// Represents the specific amount.
	Value string `json:"value"`
}

// NewGetTransactionDetailsByTransactionIDRIBSZVout instantiates a new GetTransactionDetailsByTransactionIDRIBSZVout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSZVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSZScriptPubKey, value string) *GetTransactionDetailsByTransactionIDRIBSZVout {
	this := GetTransactionDetailsByTransactionIDRIBSZVout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSZVoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSZVoutWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVout {
	this := GetTransactionDetailsByTransactionIDRIBSZVout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSZScriptPubKey {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSZScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSZScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSZScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSZVout) SetValue(v string) {
	o.Value = v
}

func (o GetTransactionDetailsByTransactionIDRIBSZVout) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDRIBSZVout struct {
	value *GetTransactionDetailsByTransactionIDRIBSZVout
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSZVout) Get() *GetTransactionDetailsByTransactionIDRIBSZVout {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSZVout) Set(val *GetTransactionDetailsByTransactionIDRIBSZVout) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSZVout) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSZVout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSZVout(val *GetTransactionDetailsByTransactionIDRIBSZVout) *NullableGetTransactionDetailsByTransactionIDRIBSZVout {
	return &NullableGetTransactionDetailsByTransactionIDRIBSZVout{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSZVout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSZVout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

