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

// GetTransactionDetailsByTransactionIDRIBSDVout struct for GetTransactionDetailsByTransactionIDRIBSDVout
type GetTransactionDetailsByTransactionIDRIBSDVout struct {
	// Defines whether the output is spent or not.
	IsSpent bool `json:"isSpent"`
	ScriptPubKey GetTransactionDetailsByTransactionIDRIBSDScriptPubKey `json:"scriptPubKey"`
	// Represents the sent/received amount.
	Value string `json:"value"`
}

// NewGetTransactionDetailsByTransactionIDRIBSDVout instantiates a new GetTransactionDetailsByTransactionIDRIBSDVout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSDVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSDScriptPubKey, value string) *GetTransactionDetailsByTransactionIDRIBSDVout {
	this := GetTransactionDetailsByTransactionIDRIBSDVout{}
	this.IsSpent = isSpent
	this.ScriptPubKey = scriptPubKey
	this.Value = value
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSDVoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSDVout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSDVoutWithDefaults() *GetTransactionDetailsByTransactionIDRIBSDVout {
	this := GetTransactionDetailsByTransactionIDRIBSDVout{}
	return &this
}

// GetIsSpent returns the IsSpent field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) GetIsSpent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSpent
}

// GetIsSpentOk returns a tuple with the IsSpent field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) GetIsSpentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSpent, true
}

// SetIsSpent sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) SetIsSpent(v bool) {
	o.IsSpent = v
}

// GetScriptPubKey returns the ScriptPubKey field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSDScriptPubKey {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSDScriptPubKey
		return ret
	}

	return o.ScriptPubKey
}

// GetScriptPubKeyOk returns a tuple with the ScriptPubKey field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSDScriptPubKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScriptPubKey, true
}

// SetScriptPubKey sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSDScriptPubKey) {
	o.ScriptPubKey = v
}

// GetValue returns the Value field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSDVout) SetValue(v string) {
	o.Value = v
}

func (o GetTransactionDetailsByTransactionIDRIBSDVout) MarshalJSON() ([]byte, error) {
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

type NullableGetTransactionDetailsByTransactionIDRIBSDVout struct {
	value *GetTransactionDetailsByTransactionIDRIBSDVout
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSDVout) Get() *GetTransactionDetailsByTransactionIDRIBSDVout {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSDVout) Set(val *GetTransactionDetailsByTransactionIDRIBSDVout) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSDVout) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSDVout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSDVout(val *GetTransactionDetailsByTransactionIDRIBSDVout) *NullableGetTransactionDetailsByTransactionIDRIBSDVout {
	return &NullableGetTransactionDetailsByTransactionIDRIBSDVout{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSDVout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSDVout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


