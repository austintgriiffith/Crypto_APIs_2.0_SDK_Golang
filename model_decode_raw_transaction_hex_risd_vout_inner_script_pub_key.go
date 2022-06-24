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

// DecodeRawTransactionHexRISDVoutInnerScriptPubKey Represents the script public key.
type DecodeRawTransactionHexRISDVoutInnerScriptPubKey struct {
	// Represents the addresses which send the amount
	Address string `json:"address"`
	// Represents the assembly of the script public key of the address.
	Asm *string `json:"asm,omitempty"`
	// Represents the hex of the script public key of the address.
	Hex *string `json:"hex,omitempty"`
	// Represents the hex of the script public key of the address.
	Type *string `json:"type,omitempty"`
}

// NewDecodeRawTransactionHexRISDVoutInnerScriptPubKey instantiates a new DecodeRawTransactionHexRISDVoutInnerScriptPubKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISDVoutInnerScriptPubKey(address string) *DecodeRawTransactionHexRISDVoutInnerScriptPubKey {
	this := DecodeRawTransactionHexRISDVoutInnerScriptPubKey{}
	this.Address = address
	return &this
}

// NewDecodeRawTransactionHexRISDVoutInnerScriptPubKeyWithDefaults instantiates a new DecodeRawTransactionHexRISDVoutInnerScriptPubKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISDVoutInnerScriptPubKeyWithDefaults() *DecodeRawTransactionHexRISDVoutInnerScriptPubKey {
	this := DecodeRawTransactionHexRISDVoutInnerScriptPubKey{}
	return &this
}

// GetAddress returns the Address field value
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) SetAddress(v string) {
	o.Address = v
}

// GetAsm returns the Asm field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetAsm() string {
	if o == nil || o.Asm == nil {
		var ret string
		return ret
	}
	return *o.Asm
}

// GetAsmOk returns a tuple with the Asm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetAsmOk() (*string, bool) {
	if o == nil || o.Asm == nil {
		return nil, false
	}
	return o.Asm, true
}

// HasAsm returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) HasAsm() bool {
	if o != nil && o.Asm != nil {
		return true
	}

	return false
}

// SetAsm gets a reference to the given string and assigns it to the Asm field.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) SetAsm(v string) {
	o.Asm = &v
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetHex() string {
	if o == nil || o.Hex == nil {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetHexOk() (*string, bool) {
	if o == nil || o.Hex == nil {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) HasHex() bool {
	if o != nil && o.Hex != nil {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) SetHex(v string) {
	o.Hex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) SetType(v string) {
	o.Type = &v
}

func (o DecodeRawTransactionHexRISDVoutInnerScriptPubKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.Asm != nil {
		toSerialize["asm"] = o.Asm
	}
	if o.Hex != nil {
		toSerialize["hex"] = o.Hex
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey struct {
	value *DecodeRawTransactionHexRISDVoutInnerScriptPubKey
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey) Get() *DecodeRawTransactionHexRISDVoutInnerScriptPubKey {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey) Set(val *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey(val *DecodeRawTransactionHexRISDVoutInnerScriptPubKey) *NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey {
	return &NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISDVoutInnerScriptPubKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

