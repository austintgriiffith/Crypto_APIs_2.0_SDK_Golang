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

// GetTransactionRequestDetailsRIRecipients struct for GetTransactionRequestDetailsRIRecipients
type GetTransactionRequestDetailsRIRecipients struct {
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient.
	Address string `json:"address"`
	// Represents the amount received to this address.
	Amount string `json:"amount"`
	// Defines the unit of the amount.
	Unit string `json:"unit"`
}

// NewGetTransactionRequestDetailsRIRecipients instantiates a new GetTransactionRequestDetailsRIRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionRequestDetailsRIRecipients(address string, amount string, unit string) *GetTransactionRequestDetailsRIRecipients {
	this := GetTransactionRequestDetailsRIRecipients{}
	this.Address = address
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetTransactionRequestDetailsRIRecipientsWithDefaults instantiates a new GetTransactionRequestDetailsRIRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionRequestDetailsRIRecipientsWithDefaults() *GetTransactionRequestDetailsRIRecipients {
	this := GetTransactionRequestDetailsRIRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetTransactionRequestDetailsRIRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetTransactionRequestDetailsRIRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetTransactionRequestDetailsRIRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *GetTransactionRequestDetailsRIRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionRequestDetailsRIRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionRequestDetailsRIRecipients) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetTransactionRequestDetailsRIRecipients) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionRequestDetailsRIRecipients) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetTransactionRequestDetailsRIRecipients) SetUnit(v string) {
	o.Unit = v
}

func (o GetTransactionRequestDetailsRIRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionRequestDetailsRIRecipients struct {
	value *GetTransactionRequestDetailsRIRecipients
	isSet bool
}

func (v NullableGetTransactionRequestDetailsRIRecipients) Get() *GetTransactionRequestDetailsRIRecipients {
	return v.value
}

func (v *NullableGetTransactionRequestDetailsRIRecipients) Set(val *GetTransactionRequestDetailsRIRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionRequestDetailsRIRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionRequestDetailsRIRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionRequestDetailsRIRecipients(val *GetTransactionRequestDetailsRIRecipients) *NullableGetTransactionRequestDetailsRIRecipients {
	return &NullableGetTransactionRequestDetailsRIRecipients{value: val, isSet: true}
}

func (v NullableGetTransactionRequestDetailsRIRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionRequestDetailsRIRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


