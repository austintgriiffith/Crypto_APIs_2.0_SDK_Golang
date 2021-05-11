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

// GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin Bitcoin
type GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin struct {
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Represents a random value that can be adjusted to satisfy the proof of work
	Nonce string `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// A sub-unit of BTC equal to 0.000001 BTC, or 100 Satoshi, and is the same as microbitcoin (μBTC). Bits have two-decimal precision.
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	MerkleRoot string `json:"merkleRoot"`
	// Defines the numeric representation of the block size excluding the witness data.
	StrippedSize int32 `json:"strippedSize"`
	// Represents the version of the specific block on the blockchain.
	Version int32 `json:"version"`
	// Is the hexadecimal string representation of the block's version.
	VersionHex string `json:"versionHex"`
	// Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit.
	Weight int32 `json:"weight"`
}

// NewGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin instantiates a new GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin(difficulty string, nonce string, size int32, bits string, chainwork string, merkleRoot string, strippedSize int32, version int32, versionHex string, weight int32) *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin {
	this := GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin{}
	this.Difficulty = difficulty
	this.Nonce = nonce
	this.Size = size
	this.Bits = bits
	this.Chainwork = chainwork
	this.MerkleRoot = merkleRoot
	this.StrippedSize = strippedSize
	this.Version = version
	this.VersionHex = versionHex
	this.Weight = weight
	return &this
}

// NewGetLatestMinedBlockResponseItemBlockchainSpecificBitcoinWithDefaults instantiates a new GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLatestMinedBlockResponseItemBlockchainSpecificBitcoinWithDefaults() *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin {
	this := GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetDifficultyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetNonce returns the Nonce field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetNonceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetNonce(v string) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetSize(v int32) {
	o.Size = v
}

// GetBits returns the Bits field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetBitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetChainworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetChainwork(v string) {
	o.Chainwork = v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetMerkleRootOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetStrippedSize returns the StrippedSize field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetStrippedSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrippedSize
}

// GetStrippedSizeOk returns a tuple with the StrippedSize field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetStrippedSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StrippedSize, true
}

// SetStrippedSize sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetStrippedSize(v int32) {
	o.StrippedSize = v
}

// GetVersion returns the Version field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetVersion(v int32) {
	o.Version = v
}

// GetVersionHex returns the VersionHex field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetVersionHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionHex
}

// GetVersionHexOk returns a tuple with the VersionHex field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetVersionHexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VersionHex, true
}

// SetVersionHex sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetVersionHex(v string) {
	o.VersionHex = v
}

// GetWeight returns the Weight field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) GetWeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) SetWeight(v int32) {
	o.Weight = v
}

func (o GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) MarshalJSON() ([]byte, error) {
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
		toSerialize["bits"] = o.Bits
	}
	if true {
		toSerialize["chainwork"] = o.Chainwork
	}
	if true {
		toSerialize["merkleRoot"] = o.MerkleRoot
	}
	if true {
		toSerialize["strippedSize"] = o.StrippedSize
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["versionHex"] = o.VersionHex
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin struct {
	value *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin
	isSet bool
}

func (v NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) Get() *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin {
	return v.value
}

func (v *NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) Set(val *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin(val *GetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) *NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin {
	return &NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin{value: val, isSet: true}
}

func (v NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLatestMinedBlockResponseItemBlockchainSpecificBitcoin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

