# ListAssetsDetailsRISC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1HourPriceChangeInPercentage** | **string** | Represents the percentage of the asset&#39;s current price against the its price from 1 hour ago. | 
**Var1WeekPriceChangeInPercentage** | **string** | Represents the percentage of the asset&#39;s current price against the its price from 1 week ago. | 
**Var24HoursPriceChangeInPercentage** | **string** | Represents the percentage of the asset&#39;s current price against the its price from 24 hours ago. | 
**Var24HoursTradingVolume** | **string** | Represents the trading volume of the asset for the time frame of 24 hours. | 
**AssetType** | **string** | Represent a subtype of the crypto assets. Could be COIN or TOKEN. | 
**CirculatingSupply** | **string** | Represents the amount of the asset that is circulating on the market and in public hands. | 
**MarketCapInUSD** | **string** | Defines the total market value of the asset&#39;s circulating supply in USD. | 
**MaxSupply** | **string** | Represents the maximum amount of all coins of a specific asset that will ever exist in its lifetime. | 

## Methods

### NewListAssetsDetailsRISC

`func NewListAssetsDetailsRISC(var1HourPriceChangeInPercentage string, var1WeekPriceChangeInPercentage string, var24HoursPriceChangeInPercentage string, var24HoursTradingVolume string, assetType string, circulatingSupply string, marketCapInUSD string, maxSupply string, ) *ListAssetsDetailsRISC`

NewListAssetsDetailsRISC instantiates a new ListAssetsDetailsRISC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetsDetailsRISCWithDefaults

`func NewListAssetsDetailsRISCWithDefaults() *ListAssetsDetailsRISC`

NewListAssetsDetailsRISCWithDefaults instantiates a new ListAssetsDetailsRISC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1HourPriceChangeInPercentage

`func (o *ListAssetsDetailsRISC) GetVar1HourPriceChangeInPercentage() string`

GetVar1HourPriceChangeInPercentage returns the Var1HourPriceChangeInPercentage field if non-nil, zero value otherwise.

### GetVar1HourPriceChangeInPercentageOk

`func (o *ListAssetsDetailsRISC) GetVar1HourPriceChangeInPercentageOk() (*string, bool)`

GetVar1HourPriceChangeInPercentageOk returns a tuple with the Var1HourPriceChangeInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1HourPriceChangeInPercentage

`func (o *ListAssetsDetailsRISC) SetVar1HourPriceChangeInPercentage(v string)`

SetVar1HourPriceChangeInPercentage sets Var1HourPriceChangeInPercentage field to given value.


### GetVar1WeekPriceChangeInPercentage

`func (o *ListAssetsDetailsRISC) GetVar1WeekPriceChangeInPercentage() string`

GetVar1WeekPriceChangeInPercentage returns the Var1WeekPriceChangeInPercentage field if non-nil, zero value otherwise.

### GetVar1WeekPriceChangeInPercentageOk

`func (o *ListAssetsDetailsRISC) GetVar1WeekPriceChangeInPercentageOk() (*string, bool)`

GetVar1WeekPriceChangeInPercentageOk returns a tuple with the Var1WeekPriceChangeInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1WeekPriceChangeInPercentage

`func (o *ListAssetsDetailsRISC) SetVar1WeekPriceChangeInPercentage(v string)`

SetVar1WeekPriceChangeInPercentage sets Var1WeekPriceChangeInPercentage field to given value.


### GetVar24HoursPriceChangeInPercentage

`func (o *ListAssetsDetailsRISC) GetVar24HoursPriceChangeInPercentage() string`

GetVar24HoursPriceChangeInPercentage returns the Var24HoursPriceChangeInPercentage field if non-nil, zero value otherwise.

### GetVar24HoursPriceChangeInPercentageOk

`func (o *ListAssetsDetailsRISC) GetVar24HoursPriceChangeInPercentageOk() (*string, bool)`

GetVar24HoursPriceChangeInPercentageOk returns a tuple with the Var24HoursPriceChangeInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar24HoursPriceChangeInPercentage

`func (o *ListAssetsDetailsRISC) SetVar24HoursPriceChangeInPercentage(v string)`

SetVar24HoursPriceChangeInPercentage sets Var24HoursPriceChangeInPercentage field to given value.


### GetVar24HoursTradingVolume

`func (o *ListAssetsDetailsRISC) GetVar24HoursTradingVolume() string`

GetVar24HoursTradingVolume returns the Var24HoursTradingVolume field if non-nil, zero value otherwise.

### GetVar24HoursTradingVolumeOk

`func (o *ListAssetsDetailsRISC) GetVar24HoursTradingVolumeOk() (*string, bool)`

GetVar24HoursTradingVolumeOk returns a tuple with the Var24HoursTradingVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar24HoursTradingVolume

`func (o *ListAssetsDetailsRISC) SetVar24HoursTradingVolume(v string)`

SetVar24HoursTradingVolume sets Var24HoursTradingVolume field to given value.


### GetAssetType

`func (o *ListAssetsDetailsRISC) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *ListAssetsDetailsRISC) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *ListAssetsDetailsRISC) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCirculatingSupply

`func (o *ListAssetsDetailsRISC) GetCirculatingSupply() string`

GetCirculatingSupply returns the CirculatingSupply field if non-nil, zero value otherwise.

### GetCirculatingSupplyOk

`func (o *ListAssetsDetailsRISC) GetCirculatingSupplyOk() (*string, bool)`

GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCirculatingSupply

`func (o *ListAssetsDetailsRISC) SetCirculatingSupply(v string)`

SetCirculatingSupply sets CirculatingSupply field to given value.


### GetMarketCapInUSD

`func (o *ListAssetsDetailsRISC) GetMarketCapInUSD() string`

GetMarketCapInUSD returns the MarketCapInUSD field if non-nil, zero value otherwise.

### GetMarketCapInUSDOk

`func (o *ListAssetsDetailsRISC) GetMarketCapInUSDOk() (*string, bool)`

GetMarketCapInUSDOk returns a tuple with the MarketCapInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapInUSD

`func (o *ListAssetsDetailsRISC) SetMarketCapInUSD(v string)`

SetMarketCapInUSD sets MarketCapInUSD field to given value.


### GetMaxSupply

`func (o *ListAssetsDetailsRISC) GetMaxSupply() string`

GetMaxSupply returns the MaxSupply field if non-nil, zero value otherwise.

### GetMaxSupplyOk

`func (o *ListAssetsDetailsRISC) GetMaxSupplyOk() (*string, bool)`

GetMaxSupplyOk returns a tuple with the MaxSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupply

`func (o *ListAssetsDetailsRISC) SetMaxSupply(v string)`

SetMaxSupply sets MaxSupply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


