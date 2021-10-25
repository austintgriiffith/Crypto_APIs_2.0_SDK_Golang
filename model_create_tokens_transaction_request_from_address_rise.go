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

// CreateTokensTransactionRequestFromAddressRISE Ethereum Erc20 Token
type CreateTokensTransactionRequestFromAddressRISE struct {
	// Defines the contract address in the blockchain for an ERC20 token.
	ContractAddress string `json:"contractAddress"`
}

// NewCreateTokensTransactionRequestFromAddressRISE instantiates a new CreateTokensTransactionRequestFromAddressRISE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokensTransactionRequestFromAddressRISE(contractAddress string) *CreateTokensTransactionRequestFromAddressRISE {
	this := CreateTokensTransactionRequestFromAddressRISE{}
	this.ContractAddress = contractAddress
	return &this
}

// NewCreateTokensTransactionRequestFromAddressRISEWithDefaults instantiates a new CreateTokensTransactionRequestFromAddressRISE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokensTransactionRequestFromAddressRISEWithDefaults() *CreateTokensTransactionRequestFromAddressRISE {
	this := CreateTokensTransactionRequestFromAddressRISE{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *CreateTokensTransactionRequestFromAddressRISE) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *CreateTokensTransactionRequestFromAddressRISE) GetContractAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *CreateTokensTransactionRequestFromAddressRISE) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o CreateTokensTransactionRequestFromAddressRISE) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokensTransactionRequestFromAddressRISE struct {
	value *CreateTokensTransactionRequestFromAddressRISE
	isSet bool
}

func (v NullableCreateTokensTransactionRequestFromAddressRISE) Get() *CreateTokensTransactionRequestFromAddressRISE {
	return v.value
}

func (v *NullableCreateTokensTransactionRequestFromAddressRISE) Set(val *CreateTokensTransactionRequestFromAddressRISE) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokensTransactionRequestFromAddressRISE) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokensTransactionRequestFromAddressRISE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokensTransactionRequestFromAddressRISE(val *CreateTokensTransactionRequestFromAddressRISE) *NullableCreateTokensTransactionRequestFromAddressRISE {
	return &NullableCreateTokensTransactionRequestFromAddressRISE{value: val, isSet: true}
}

func (v NullableCreateTokensTransactionRequestFromAddressRISE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokensTransactionRequestFromAddressRISE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

