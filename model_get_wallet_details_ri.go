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

// GetWalletDetailsRI struct for GetWalletDetailsRI
type GetWalletDetailsRI struct {
	// Specifies the count of deposit addresses in the Wallet.
	DepositAddressesCount int32 `json:"depositAddressesCount"`
	// Defines the name of the Wallet given to it by the user.
	Name string `json:"name"`
}

// NewGetWalletDetailsRI instantiates a new GetWalletDetailsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletDetailsRI(depositAddressesCount int32, name string) *GetWalletDetailsRI {
	this := GetWalletDetailsRI{}
	this.DepositAddressesCount = depositAddressesCount
	this.Name = name
	return &this
}

// NewGetWalletDetailsRIWithDefaults instantiates a new GetWalletDetailsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletDetailsRIWithDefaults() *GetWalletDetailsRI {
	this := GetWalletDetailsRI{}
	return &this
}

// GetDepositAddressesCount returns the DepositAddressesCount field value
func (o *GetWalletDetailsRI) GetDepositAddressesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DepositAddressesCount
}

// GetDepositAddressesCountOk returns a tuple with the DepositAddressesCount field value
// and a boolean to check if the value has been set.
func (o *GetWalletDetailsRI) GetDepositAddressesCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositAddressesCount, true
}

// SetDepositAddressesCount sets field value
func (o *GetWalletDetailsRI) SetDepositAddressesCount(v int32) {
	o.DepositAddressesCount = v
}

// GetName returns the Name field value
func (o *GetWalletDetailsRI) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetWalletDetailsRI) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetWalletDetailsRI) SetName(v string) {
	o.Name = v
}

func (o GetWalletDetailsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["depositAddressesCount"] = o.DepositAddressesCount
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletDetailsRI struct {
	value *GetWalletDetailsRI
	isSet bool
}

func (v NullableGetWalletDetailsRI) Get() *GetWalletDetailsRI {
	return v.value
}

func (v *NullableGetWalletDetailsRI) Set(val *GetWalletDetailsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletDetailsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletDetailsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletDetailsRI(val *GetWalletDetailsRI) *NullableGetWalletDetailsRI {
	return &NullableGetWalletDetailsRI{value: val, isSet: true}
}

func (v NullableGetWalletDetailsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletDetailsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


