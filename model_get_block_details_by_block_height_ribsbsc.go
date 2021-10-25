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

// GetBlockDetailsByBlockHeightRIBSBSC Binance Smart Chain
type GetBlockDetailsByBlockHeightRIBSBSC struct {
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Represents a random value that can be adjusted to satisfy the Proof of Work.
	Nonce string `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// Represents any data that can be included by the miner in the block.
	ExtraData string `json:"extraData"`
	// Defines the total gas limit of all transactions in the block.
	GasLimit string `json:"gasLimit"`
	// Represents the total amount of gas used by all transactions in this block.
	GasUsed string `json:"gasUsed"`
	// Specifies the amount of time required for the block to be mined in seconds.
	MinedInSeconds int32 `json:"minedInSeconds"`
	// Defines the combined hash of all uncles for a given parent.
	Sha3Uncles string `json:"sha3Uncles"`
	// Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block.
	TotalDifficulty string `json:"totalDifficulty"`
}

// NewGetBlockDetailsByBlockHeightRIBSBSC instantiates a new GetBlockDetailsByBlockHeightRIBSBSC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHeightRIBSBSC(difficulty string, nonce string, size int32, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, sha3Uncles string, totalDifficulty string) *GetBlockDetailsByBlockHeightRIBSBSC {
	this := GetBlockDetailsByBlockHeightRIBSBSC{}
	this.Difficulty = difficulty
	this.Nonce = nonce
	this.Size = size
	this.ExtraData = extraData
	this.GasLimit = gasLimit
	this.GasUsed = gasUsed
	this.MinedInSeconds = minedInSeconds
	this.Sha3Uncles = sha3Uncles
	this.TotalDifficulty = totalDifficulty
	return &this
}

// NewGetBlockDetailsByBlockHeightRIBSBSCWithDefaults instantiates a new GetBlockDetailsByBlockHeightRIBSBSC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHeightRIBSBSCWithDefaults() *GetBlockDetailsByBlockHeightRIBSBSC {
	this := GetBlockDetailsByBlockHeightRIBSBSC{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetDifficultyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetNonce returns the Nonce field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetNonceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetNonce(v string) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetSize(v int32) {
	o.Size = v
}

// GetExtraData returns the ExtraData field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetExtraData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetExtraDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExtraData, true
}

// SetExtraData sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetExtraData(v string) {
	o.ExtraData = v
}

// GetGasLimit returns the GasLimit field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetGasLimitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetGasUsedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetMinedInSeconds returns the MinedInSeconds field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetMinedInSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInSeconds
}

// GetMinedInSecondsOk returns a tuple with the MinedInSeconds field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetMinedInSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinedInSeconds, true
}

// SetMinedInSeconds sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetMinedInSeconds(v int32) {
	o.MinedInSeconds = v
}

// GetSha3Uncles returns the Sha3Uncles field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetSha3Uncles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha3Uncles
}

// GetSha3UnclesOk returns a tuple with the Sha3Uncles field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetSha3UnclesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sha3Uncles, true
}

// SetSha3Uncles sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetSha3Uncles(v string) {
	o.Sha3Uncles = v
}

// GetTotalDifficulty returns the TotalDifficulty field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetTotalDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalDifficulty
}

// GetTotalDifficultyOk returns a tuple with the TotalDifficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSBSC) GetTotalDifficultyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalDifficulty, true
}

// SetTotalDifficulty sets field value
func (o *GetBlockDetailsByBlockHeightRIBSBSC) SetTotalDifficulty(v string) {
	o.TotalDifficulty = v
}

func (o GetBlockDetailsByBlockHeightRIBSBSC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["difficulty"] = o.Difficulty
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["extraData"] = o.ExtraData
	}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if true {
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["minedInSeconds"] = o.MinedInSeconds
	}
	if true {
		toSerialize["sha3Uncles"] = o.Sha3Uncles
	}
	if true {
		toSerialize["totalDifficulty"] = o.TotalDifficulty
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHeightRIBSBSC struct {
	value *GetBlockDetailsByBlockHeightRIBSBSC
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHeightRIBSBSC) Get() *GetBlockDetailsByBlockHeightRIBSBSC {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHeightRIBSBSC) Set(val *GetBlockDetailsByBlockHeightRIBSBSC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHeightRIBSBSC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHeightRIBSBSC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHeightRIBSBSC(val *GetBlockDetailsByBlockHeightRIBSBSC) *NullableGetBlockDetailsByBlockHeightRIBSBSC {
	return &NullableGetBlockDetailsByBlockHeightRIBSBSC{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHeightRIBSBSC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHeightRIBSBSC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


