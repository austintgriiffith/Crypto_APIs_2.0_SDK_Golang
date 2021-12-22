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

// GetBlockDetailsByBlockHashRIBSBC Bitcoin Cash
type GetBlockDetailsByBlockHashRIBSBC struct {
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Represents a random value that can be adjusted to satisfy the Proof of Work.
	Nonce string `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// A sub-unit of BTC equal to 0.000001 BTC, or 100 Satoshi, and is the same as microbitcoin (μBTC). Bits have two-decimal precision.
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	MerkleRoot string `json:"merkleRoot"`
	// Represents the version of the specific block on the blockchain.
	Version int32 `json:"version"`
	// Is the hexadecimal string representation of the block's version.
	VersionHex string `json:"versionHex"`
}

// NewGetBlockDetailsByBlockHashRIBSBC instantiates a new GetBlockDetailsByBlockHashRIBSBC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashRIBSBC(difficulty string, nonce string, size int32, bits string, chainwork string, merkleRoot string, version int32, versionHex string) *GetBlockDetailsByBlockHashRIBSBC {
	this := GetBlockDetailsByBlockHashRIBSBC{}
	this.Difficulty = difficulty
	this.Nonce = nonce
	this.Size = size
	this.Bits = bits
	this.Chainwork = chainwork
	this.MerkleRoot = merkleRoot
	this.Version = version
	this.VersionHex = versionHex
	return &this
}

// NewGetBlockDetailsByBlockHashRIBSBCWithDefaults instantiates a new GetBlockDetailsByBlockHashRIBSBC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashRIBSBCWithDefaults() *GetBlockDetailsByBlockHashRIBSBC {
	this := GetBlockDetailsByBlockHashRIBSBC{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetDifficultyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetNonce returns the Nonce field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetNonceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetNonce(v string) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetSize(v int32) {
	o.Size = v
}

// GetBits returns the Bits field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetBitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetChainworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetChainwork(v string) {
	o.Chainwork = v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetMerkleRootOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetVersion returns the Version field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetVersion(v int32) {
	o.Version = v
}

// GetVersionHex returns the VersionHex field value
func (o *GetBlockDetailsByBlockHashRIBSBC) GetVersionHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionHex
}

// GetVersionHexOk returns a tuple with the VersionHex field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSBC) GetVersionHexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VersionHex, true
}

// SetVersionHex sets field value
func (o *GetBlockDetailsByBlockHashRIBSBC) SetVersionHex(v string) {
	o.VersionHex = v
}

func (o GetBlockDetailsByBlockHashRIBSBC) MarshalJSON() ([]byte, error) {
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
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["versionHex"] = o.VersionHex
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashRIBSBC struct {
	value *GetBlockDetailsByBlockHashRIBSBC
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashRIBSBC) Get() *GetBlockDetailsByBlockHashRIBSBC {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashRIBSBC) Set(val *GetBlockDetailsByBlockHashRIBSBC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashRIBSBC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashRIBSBC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashRIBSBC(val *GetBlockDetailsByBlockHashRIBSBC) *NullableGetBlockDetailsByBlockHashRIBSBC {
	return &NullableGetBlockDetailsByBlockHashRIBSBC{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashRIBSBC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashRIBSBC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


