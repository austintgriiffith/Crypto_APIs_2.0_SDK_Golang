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

// GetWalletAssetDetailsRISentConfirmedAmount Specifies the confirmed amount that has been sent.
type GetWalletAssetDetailsRISentConfirmedAmount struct {
	// Specifies the confirmed amount that has been sent.
	Amount string `json:"amount"`
	// Specifies the unit of the confirmed amount that has been sent.
	Unit string `json:"unit"`
}

// NewGetWalletAssetDetailsRISentConfirmedAmount instantiates a new GetWalletAssetDetailsRISentConfirmedAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWalletAssetDetailsRISentConfirmedAmount(amount string, unit string) *GetWalletAssetDetailsRISentConfirmedAmount {
	this := GetWalletAssetDetailsRISentConfirmedAmount{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetWalletAssetDetailsRISentConfirmedAmountWithDefaults instantiates a new GetWalletAssetDetailsRISentConfirmedAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWalletAssetDetailsRISentConfirmedAmountWithDefaults() *GetWalletAssetDetailsRISentConfirmedAmount {
	this := GetWalletAssetDetailsRISentConfirmedAmount{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetWalletAssetDetailsRISentConfirmedAmount) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRISentConfirmedAmount) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetWalletAssetDetailsRISentConfirmedAmount) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetWalletAssetDetailsRISentConfirmedAmount) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetWalletAssetDetailsRISentConfirmedAmount) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetWalletAssetDetailsRISentConfirmedAmount) SetUnit(v string) {
	o.Unit = v
}

func (o GetWalletAssetDetailsRISentConfirmedAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetWalletAssetDetailsRISentConfirmedAmount struct {
	value *GetWalletAssetDetailsRISentConfirmedAmount
	isSet bool
}

func (v NullableGetWalletAssetDetailsRISentConfirmedAmount) Get() *GetWalletAssetDetailsRISentConfirmedAmount {
	return v.value
}

func (v *NullableGetWalletAssetDetailsRISentConfirmedAmount) Set(val *GetWalletAssetDetailsRISentConfirmedAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWalletAssetDetailsRISentConfirmedAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWalletAssetDetailsRISentConfirmedAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWalletAssetDetailsRISentConfirmedAmount(val *GetWalletAssetDetailsRISentConfirmedAmount) *NullableGetWalletAssetDetailsRISentConfirmedAmount {
	return &NullableGetWalletAssetDetailsRISentConfirmedAmount{value: val, isSet: true}
}

func (v NullableGetWalletAssetDetailsRISentConfirmedAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWalletAssetDetailsRISentConfirmedAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


