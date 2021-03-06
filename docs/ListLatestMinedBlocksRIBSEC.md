# ListLatestMinedBlocksRIBSEC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Defines the total gas limit of all transactions in the block. | 
**GasUsed** | **string** | Represents the total amount of gas used by all transactions in this block. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in seconds. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the proof of work | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block. | 
**Uncles** | **[]string** |  | 

## Methods

### NewListLatestMinedBlocksRIBSEC

`func NewListLatestMinedBlocksRIBSEC(difficulty string, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, nonce string, sha3Uncles string, size int32, totalDifficulty string, uncles []string, ) *ListLatestMinedBlocksRIBSEC`

NewListLatestMinedBlocksRIBSEC instantiates a new ListLatestMinedBlocksRIBSEC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSECWithDefaults

`func NewListLatestMinedBlocksRIBSECWithDefaults() *ListLatestMinedBlocksRIBSEC`

NewListLatestMinedBlocksRIBSECWithDefaults instantiates a new ListLatestMinedBlocksRIBSEC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *ListLatestMinedBlocksRIBSEC) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *ListLatestMinedBlocksRIBSEC) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *ListLatestMinedBlocksRIBSEC) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetExtraData

`func (o *ListLatestMinedBlocksRIBSEC) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ListLatestMinedBlocksRIBSEC) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ListLatestMinedBlocksRIBSEC) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *ListLatestMinedBlocksRIBSEC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListLatestMinedBlocksRIBSEC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListLatestMinedBlocksRIBSEC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *ListLatestMinedBlocksRIBSEC) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListLatestMinedBlocksRIBSEC) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListLatestMinedBlocksRIBSEC) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *ListLatestMinedBlocksRIBSEC) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *ListLatestMinedBlocksRIBSEC) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *ListLatestMinedBlocksRIBSEC) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetNonce

`func (o *ListLatestMinedBlocksRIBSEC) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListLatestMinedBlocksRIBSEC) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListLatestMinedBlocksRIBSEC) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSha3Uncles

`func (o *ListLatestMinedBlocksRIBSEC) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *ListLatestMinedBlocksRIBSEC) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *ListLatestMinedBlocksRIBSEC) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetSize

`func (o *ListLatestMinedBlocksRIBSEC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListLatestMinedBlocksRIBSEC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListLatestMinedBlocksRIBSEC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotalDifficulty

`func (o *ListLatestMinedBlocksRIBSEC) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *ListLatestMinedBlocksRIBSEC) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *ListLatestMinedBlocksRIBSEC) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *ListLatestMinedBlocksRIBSEC) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *ListLatestMinedBlocksRIBSEC) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *ListLatestMinedBlocksRIBSEC) SetUncles(v []string)`

SetUncles sets Uncles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


