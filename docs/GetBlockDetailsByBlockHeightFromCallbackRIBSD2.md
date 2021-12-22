# GetBlockDetailsByBlockHeightFromCallbackRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | **string** | Represents a specific sub-unit of Doge. Bits have two-decimal precision. | 
**Chainwork** | **string** | Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes. | 
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**MerkleRoot** | **string** | Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions&#39; hashes that are part of a blockchain block. | 
**Nonce** | **int32** | Represents a random value that can be adjusted to satisfy the proof of work | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**StrippedSize** | **int32** | Defines the numeric representation of the block size excluding the witness data. | 
**Version** | **int32** | Represents the version of the specific block on the blockchain. | 
**Weight** | **int32** | Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit. | 

## Methods

### NewGetBlockDetailsByBlockHeightFromCallbackRIBSD2

`func NewGetBlockDetailsByBlockHeightFromCallbackRIBSD2(bits string, chainwork string, difficulty string, merkleRoot string, nonce int32, size int32, strippedSize int32, version int32, weight int32, ) *GetBlockDetailsByBlockHeightFromCallbackRIBSD2`

NewGetBlockDetailsByBlockHeightFromCallbackRIBSD2 instantiates a new GetBlockDetailsByBlockHeightFromCallbackRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHeightFromCallbackRIBSD2WithDefaults

`func NewGetBlockDetailsByBlockHeightFromCallbackRIBSD2WithDefaults() *GetBlockDetailsByBlockHeightFromCallbackRIBSD2`

NewGetBlockDetailsByBlockHeightFromCallbackRIBSD2WithDefaults instantiates a new GetBlockDetailsByBlockHeightFromCallbackRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetBits() string`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetBitsOk() (*string, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetBits(v string)`

SetBits sets Bits field to given value.


### GetChainwork

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.


### GetDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetMerkleRoot

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStrippedSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetStrippedSize() int32`

GetStrippedSize returns the StrippedSize field if non-nil, zero value otherwise.

### GetStrippedSizeOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetStrippedSizeOk() (*int32, bool)`

GetStrippedSizeOk returns a tuple with the StrippedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrippedSize

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetStrippedSize(v int32)`

SetStrippedSize sets StrippedSize field to given value.


### GetVersion

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWeight

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetBlockDetailsByBlockHeightFromCallbackRIBSD2) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

