# ListAllUnconfirmedTransactionsRIBSD2VinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig**](ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSD2VinInner

`func NewListAllUnconfirmedTransactionsRIBSD2VinInner(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32, ) *ListAllUnconfirmedTransactionsRIBSD2VinInner`

NewListAllUnconfirmedTransactionsRIBSD2VinInner instantiates a new ListAllUnconfirmedTransactionsRIBSD2VinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSD2VinInnerWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSD2VinInnerWithDefaults() *ListAllUnconfirmedTransactionsRIBSD2VinInner`

NewListAllUnconfirmedTransactionsRIBSD2VinInnerWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSD2VinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetScriptSig() ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSD2VinInner) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


