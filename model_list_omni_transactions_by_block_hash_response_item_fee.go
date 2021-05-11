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

// ListOmniTransactionsByBlockHashResponseItemFee struct for ListOmniTransactionsByBlockHashResponseItemFee
type ListOmniTransactionsByBlockHashResponseItemFee struct {
	// Defines the amount of the fee.
	Amount string `json:"amount"`
	// Defines the unit of the fee.
	Unit string `json:"unit"`
}

// NewListOmniTransactionsByBlockHashResponseItemFee instantiates a new ListOmniTransactionsByBlockHashResponseItemFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOmniTransactionsByBlockHashResponseItemFee(amount string, unit string) *ListOmniTransactionsByBlockHashResponseItemFee {
	this := ListOmniTransactionsByBlockHashResponseItemFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListOmniTransactionsByBlockHashResponseItemFeeWithDefaults instantiates a new ListOmniTransactionsByBlockHashResponseItemFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOmniTransactionsByBlockHashResponseItemFeeWithDefaults() *ListOmniTransactionsByBlockHashResponseItemFee {
	this := ListOmniTransactionsByBlockHashResponseItemFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListOmniTransactionsByBlockHashResponseItemFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHashResponseItemFee) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListOmniTransactionsByBlockHashResponseItemFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListOmniTransactionsByBlockHashResponseItemFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByBlockHashResponseItemFee) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListOmniTransactionsByBlockHashResponseItemFee) SetUnit(v string) {
	o.Unit = v
}

func (o ListOmniTransactionsByBlockHashResponseItemFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListOmniTransactionsByBlockHashResponseItemFee struct {
	value *ListOmniTransactionsByBlockHashResponseItemFee
	isSet bool
}

func (v NullableListOmniTransactionsByBlockHashResponseItemFee) Get() *ListOmniTransactionsByBlockHashResponseItemFee {
	return v.value
}

func (v *NullableListOmniTransactionsByBlockHashResponseItemFee) Set(val *ListOmniTransactionsByBlockHashResponseItemFee) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByBlockHashResponseItemFee) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByBlockHashResponseItemFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByBlockHashResponseItemFee(val *ListOmniTransactionsByBlockHashResponseItemFee) *NullableListOmniTransactionsByBlockHashResponseItemFee {
	return &NullableListOmniTransactionsByBlockHashResponseItemFee{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByBlockHashResponseItemFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByBlockHashResponseItemFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

