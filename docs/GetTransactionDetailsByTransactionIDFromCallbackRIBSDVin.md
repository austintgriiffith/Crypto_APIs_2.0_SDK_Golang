# GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDRIBSD2ScriptSig**](GetTransactionDetailsByTransactionIDRIBSD2ScriptSig.md) |  | 
**Sequence** | **int32** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | String representation of the txid | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSD2ScriptSig, sequence int32, txinwitness []string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVin instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVinWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVinWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSDVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSD2ScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSD2ScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSD2ScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSDVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

