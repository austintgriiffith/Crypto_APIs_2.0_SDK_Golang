# GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSBScriptSig**](GetTransactionDetailsByTransactionIDFromCallbackRIBSBScriptSig.md) |  | 
**Sequence** | **int32** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBVin

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDFromCallbackRIBSBScriptSig, sequence int32, txinwitness []string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBVin instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBVinWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBVinWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSBVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetScriptSig() GetTransactionDetailsByTransactionIDFromCallbackRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetScriptSig(v GetTransactionDetailsByTransactionIDFromCallbackRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSBVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

