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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSEC Ethereum Classic
type GetTransactionDetailsByTransactionIDFromCallbackRIBSEC struct {
	// Represents the specific transaction contract.
	Contract string `json:"contract"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	GasPrice GetTransactionDetailsByTransactionIDRIBSECGasPrice `json:"gasPrice"`
	// Represents the exact unit of gas that was used for the transaction.
	GasUsed string `json:"gasUsed"`
	// Represents additional information that is required for the transaction.
	InputData string `json:"inputData"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSEC instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSEC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSEC(contract string, gasLimit string, gasPrice GetTransactionDetailsByTransactionIDRIBSECGasPrice, gasUsed string, inputData string, nonce int32) *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSEC{}
	this.Contract = contract
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.InputData = inputData
	this.Nonce = nonce
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSECWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSEC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSECWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSEC{}
	return &this
}

// GetContract returns the Contract field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetContractOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) SetContract(v string) {
	o.Contract = v
}

// GetGasLimit returns the GasLimit field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetGasLimitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetGasPrice() GetTransactionDetailsByTransactionIDRIBSECGasPrice {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSECGasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDRIBSECGasPrice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) SetGasPrice(v GetTransactionDetailsByTransactionIDRIBSECGasPrice) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetGasUsedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetInputData returns the InputData field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetInputData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetInputDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InputData, true
}

// SetInputData sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) SetInputData(v string) {
	o.InputData = v
}

// GetNonce returns the Nonce field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) GetNonceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) SetNonce(v int32) {
	o.Nonce = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contract"] = o.Contract
	}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if true {
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["inputData"] = o.InputData
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSEC) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSEC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


