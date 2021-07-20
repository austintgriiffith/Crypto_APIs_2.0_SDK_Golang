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

// GetTransactionDetailsByTransactionIDRIRecipients struct for GetTransactionDetailsByTransactionIDRIRecipients
type GetTransactionDetailsByTransactionIDRIRecipients struct {
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient.
	Address string `json:"address"`
	// Represents the amount received to this address.
	Amount string `json:"amount"`
}

// NewGetTransactionDetailsByTransactionIDRIRecipients instantiates a new GetTransactionDetailsByTransactionIDRIRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIRecipients(address string, amount string) *GetTransactionDetailsByTransactionIDRIRecipients {
	this := GetTransactionDetailsByTransactionIDRIRecipients{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIRecipientsWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIRecipientsWithDefaults() *GetTransactionDetailsByTransactionIDRIRecipients {
	this := GetTransactionDetailsByTransactionIDRIRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetTransactionDetailsByTransactionIDRIRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetTransactionDetailsByTransactionIDRIRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *GetTransactionDetailsByTransactionIDRIRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionDetailsByTransactionIDRIRecipients) SetAmount(v string) {
	o.Amount = v
}

func (o GetTransactionDetailsByTransactionIDRIRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIRecipients struct {
	value *GetTransactionDetailsByTransactionIDRIRecipients
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIRecipients) Get() *GetTransactionDetailsByTransactionIDRIRecipients {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIRecipients) Set(val *GetTransactionDetailsByTransactionIDRIRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIRecipients(val *GetTransactionDetailsByTransactionIDRIRecipients) *NullableGetTransactionDetailsByTransactionIDRIRecipients {
	return &NullableGetTransactionDetailsByTransactionIDRIRecipients{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


