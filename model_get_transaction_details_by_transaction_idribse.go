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

// GetTransactionDetailsByTransactionIDRIBSE Ethereum
type GetTransactionDetailsByTransactionIDRIBSE struct {
	// Represents the specific transaction contract.
	Contract string `json:"contract"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	GasPrice GetTransactionDetailsByTransactionIDRIBSEGasPrice `json:"gasPrice"`
	// Represents the exact unit of gas that was used for the transaction.
	GasUsed string `json:"gasUsed"`
	// Represents additional information that is required for the transaction.
	InputData string `json:"inputData"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Represents the status of this transaction.
	TransactionStatus string `json:"transactionStatus"`
}

// NewGetTransactionDetailsByTransactionIDRIBSE instantiates a new GetTransactionDetailsByTransactionIDRIBSE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSE(contract string, gasLimit string, gasPrice GetTransactionDetailsByTransactionIDRIBSEGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string) *GetTransactionDetailsByTransactionIDRIBSE {
	this := GetTransactionDetailsByTransactionIDRIBSE{}
	this.Contract = contract
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.InputData = inputData
	this.Nonce = nonce
	this.TransactionStatus = transactionStatus
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSEWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSEWithDefaults() *GetTransactionDetailsByTransactionIDRIBSE {
	this := GetTransactionDetailsByTransactionIDRIBSE{}
	return &this
}

// GetContract returns the Contract field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetContractOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetContract(v string) {
	o.Contract = v
}

// GetGasLimit returns the GasLimit field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetGasLimitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetGasPrice() GetTransactionDetailsByTransactionIDRIBSEGasPrice {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSEGasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDRIBSEGasPrice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetGasPrice(v GetTransactionDetailsByTransactionIDRIBSEGasPrice) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetGasUsedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetInputData returns the InputData field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetInputData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetInputDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InputData, true
}

// SetInputData sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetInputData(v string) {
	o.InputData = v
}

// GetNonce returns the Nonce field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetNonceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetNonce(v int32) {
	o.Nonce = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSE) GetTransactionStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSE) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o GetTransactionDetailsByTransactionIDRIBSE) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIBSE struct {
	value *GetTransactionDetailsByTransactionIDRIBSE
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSE) Get() *GetTransactionDetailsByTransactionIDRIBSE {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSE) Set(val *GetTransactionDetailsByTransactionIDRIBSE) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSE) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSE(val *GetTransactionDetailsByTransactionIDRIBSE) *NullableGetTransactionDetailsByTransactionIDRIBSE {
	return &NullableGetTransactionDetailsByTransactionIDRIBSE{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

