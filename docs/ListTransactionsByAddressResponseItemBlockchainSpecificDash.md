# ListTransactionsByAddressResponseItemBlockchainSpecificDash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVin**](ListTransactionsByAddressResponseItemBlockchainSpecificDashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVout**](ListTransactionsByAddressResponseItemBlockchainSpecificDashVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecificDash

`func NewListTransactionsByAddressResponseItemBlockchainSpecificDash(locktime int32, size int32, version int32, vin []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin, vout []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout, ) *ListTransactionsByAddressResponseItemBlockchainSpecificDash`

NewListTransactionsByAddressResponseItemBlockchainSpecificDash instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificDashWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificDashWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificDash`

NewListTransactionsByAddressResponseItemBlockchainSpecificDashWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVin() []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVinOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetVin(v []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVout() []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) GetVoutOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDash) SetVout(v []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

