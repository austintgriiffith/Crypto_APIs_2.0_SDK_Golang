# ListXRPRippleTransactionsByBlockHashRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** | Represents any additional data that may be needed. | [optional] 
**Index** | **int32** | Represents the index position of the transaction in the specific block. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]ListXRPRippleTransactionsByBlockHashRIRecipients**](ListXRPRippleTransactionsByBlockHashRIRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]ListXRPRippleTransactionsByBlockHashRISenders**](ListXRPRippleTransactionsByBlockHashRISenders.md) | Represents an object of addresses that provide the funds. | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | **string** | Defines the status of the transaction. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**Type** | **string** | Defines the type of the transaction. | 
**Fee** | [**ListXRPRippleTransactionsByBlockHashRIFee**](ListXRPRippleTransactionsByBlockHashRIFee.md) |  | 
**Offer** | [**ListXRPRippleTransactionsByBlockHashRIOffer**](ListXRPRippleTransactionsByBlockHashRIOffer.md) |  | 
**Receive** | [**ListXRPRippleTransactionsByBlockHashRIReceive**](ListXRPRippleTransactionsByBlockHashRIReceive.md) |  | 
**Value** | [**ListXRPRippleTransactionsByBlockHashRIValue**](ListXRPRippleTransactionsByBlockHashRIValue.md) |  | 

## Methods

### NewListXRPRippleTransactionsByBlockHashRI

`func NewListXRPRippleTransactionsByBlockHashRI(index int32, minedInBlockHeight int32, recipients []ListXRPRippleTransactionsByBlockHashRIRecipients, senders []ListXRPRippleTransactionsByBlockHashRISenders, sequence int32, status string, timestamp int32, transactionHash string, type_ string, fee ListXRPRippleTransactionsByBlockHashRIFee, offer ListXRPRippleTransactionsByBlockHashRIOffer, receive ListXRPRippleTransactionsByBlockHashRIReceive, value ListXRPRippleTransactionsByBlockHashRIValue, ) *ListXRPRippleTransactionsByBlockHashRI`

NewListXRPRippleTransactionsByBlockHashRI instantiates a new ListXRPRippleTransactionsByBlockHashRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListXRPRippleTransactionsByBlockHashRIWithDefaults

`func NewListXRPRippleTransactionsByBlockHashRIWithDefaults() *ListXRPRippleTransactionsByBlockHashRI`

NewListXRPRippleTransactionsByBlockHashRIWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHashRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHashRI) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetIndex

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetRecipients() []ListXRPRippleTransactionsByBlockHashRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetRecipientsOk() (*[]ListXRPRippleTransactionsByBlockHashRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetRecipients(v []ListXRPRippleTransactionsByBlockHashRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetSenders() []ListXRPRippleTransactionsByBlockHashRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetSendersOk() (*[]ListXRPRippleTransactionsByBlockHashRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetSenders(v []ListXRPRippleTransactionsByBlockHashRISenders)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetFee() ListXRPRippleTransactionsByBlockHashRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetFeeOk() (*ListXRPRippleTransactionsByBlockHashRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetFee(v ListXRPRippleTransactionsByBlockHashRIFee)`

SetFee sets Fee field to given value.


### GetOffer

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetOffer() ListXRPRippleTransactionsByBlockHashRIOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetOfferOk() (*ListXRPRippleTransactionsByBlockHashRIOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetOffer(v ListXRPRippleTransactionsByBlockHashRIOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetReceive() ListXRPRippleTransactionsByBlockHashRIReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetReceiveOk() (*ListXRPRippleTransactionsByBlockHashRIReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetReceive(v ListXRPRippleTransactionsByBlockHashRIReceive)`

SetReceive sets Receive field to given value.


### GetValue

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetValue() ListXRPRippleTransactionsByBlockHashRIValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListXRPRippleTransactionsByBlockHashRI) GetValueOk() (*ListXRPRippleTransactionsByBlockHashRIValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListXRPRippleTransactionsByBlockHashRI) SetValue(v ListXRPRippleTransactionsByBlockHashRIValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

