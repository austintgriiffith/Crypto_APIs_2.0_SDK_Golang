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

// GetAddressDetailsRIConfirmedBalance struct for GetAddressDetailsRIConfirmedBalance
type GetAddressDetailsRIConfirmedBalance struct {
	// Defines the total balance of the address that is confirmed. It doesn't include unconfirmed transactions.
	Amount string `json:"amount"`
	// Defines the unit of the confirmed balance amount, e.g. BTC, ETH, XRP.
	Unit string `json:"unit"`
}

// NewGetAddressDetailsRIConfirmedBalance instantiates a new GetAddressDetailsRIConfirmedBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddressDetailsRIConfirmedBalance(amount string, unit string) *GetAddressDetailsRIConfirmedBalance {
	this := GetAddressDetailsRIConfirmedBalance{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetAddressDetailsRIConfirmedBalanceWithDefaults instantiates a new GetAddressDetailsRIConfirmedBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddressDetailsRIConfirmedBalanceWithDefaults() *GetAddressDetailsRIConfirmedBalance {
	this := GetAddressDetailsRIConfirmedBalance{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetAddressDetailsRIConfirmedBalance) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsRIConfirmedBalance) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetAddressDetailsRIConfirmedBalance) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetAddressDetailsRIConfirmedBalance) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetAddressDetailsRIConfirmedBalance) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetAddressDetailsRIConfirmedBalance) SetUnit(v string) {
	o.Unit = v
}

func (o GetAddressDetailsRIConfirmedBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetAddressDetailsRIConfirmedBalance struct {
	value *GetAddressDetailsRIConfirmedBalance
	isSet bool
}

func (v NullableGetAddressDetailsRIConfirmedBalance) Get() *GetAddressDetailsRIConfirmedBalance {
	return v.value
}

func (v *NullableGetAddressDetailsRIConfirmedBalance) Set(val *GetAddressDetailsRIConfirmedBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddressDetailsRIConfirmedBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddressDetailsRIConfirmedBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddressDetailsRIConfirmedBalance(val *GetAddressDetailsRIConfirmedBalance) *NullableGetAddressDetailsRIConfirmedBalance {
	return &NullableGetAddressDetailsRIConfirmedBalance{value: val, isSet: true}
}

func (v NullableGetAddressDetailsRIConfirmedBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddressDetailsRIConfirmedBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


