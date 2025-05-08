# ApiAccountsSummaryGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]ApiAccountsSummaryGet200ResponseDataProductsInner**](ApiAccountsSummaryGet200ResponseDataProductsInner.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**CombinedBalances** | Pointer to [**ApiAccountsSummaryGet200ResponseDataCombinedBalances**](ApiAccountsSummaryGet200ResponseDataCombinedBalances.md) |  | [optional] 
**TotalBalances** | Pointer to [**[]ApiAccountsSummaryGet200ResponseDataTotalBalancesInner**](ApiAccountsSummaryGet200ResponseDataTotalBalancesInner.md) |  | [optional] 
**TotalHiddenAccounts** | Pointer to **float32** |  | [optional] 
**IsComboUser** | Pointer to **bool** |  | [optional] 
**OwnershipTypes** | Pointer to **[]string** |  | [optional] 
**IsItrade** | Pointer to **bool** |  | [optional] 
**IsWealth** | Pointer to **bool** |  | [optional] 
**IsRetail** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiAccountsSummaryGet200ResponseData

`func NewApiAccountsSummaryGet200ResponseData() *ApiAccountsSummaryGet200ResponseData`

NewApiAccountsSummaryGet200ResponseData instantiates a new ApiAccountsSummaryGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountsSummaryGet200ResponseDataWithDefaults

`func NewApiAccountsSummaryGet200ResponseDataWithDefaults() *ApiAccountsSummaryGet200ResponseData`

NewApiAccountsSummaryGet200ResponseDataWithDefaults instantiates a new ApiAccountsSummaryGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *ApiAccountsSummaryGet200ResponseData) GetProducts() []ApiAccountsSummaryGet200ResponseDataProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetProductsOk() (*[]ApiAccountsSummaryGet200ResponseDataProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ApiAccountsSummaryGet200ResponseData) SetProducts(v []ApiAccountsSummaryGet200ResponseDataProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ApiAccountsSummaryGet200ResponseData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetUserId

`func (o *ApiAccountsSummaryGet200ResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiAccountsSummaryGet200ResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiAccountsSummaryGet200ResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCombinedBalances

`func (o *ApiAccountsSummaryGet200ResponseData) GetCombinedBalances() ApiAccountsSummaryGet200ResponseDataCombinedBalances`

GetCombinedBalances returns the CombinedBalances field if non-nil, zero value otherwise.

### GetCombinedBalancesOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetCombinedBalancesOk() (*ApiAccountsSummaryGet200ResponseDataCombinedBalances, bool)`

GetCombinedBalancesOk returns a tuple with the CombinedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedBalances

`func (o *ApiAccountsSummaryGet200ResponseData) SetCombinedBalances(v ApiAccountsSummaryGet200ResponseDataCombinedBalances)`

SetCombinedBalances sets CombinedBalances field to given value.

### HasCombinedBalances

`func (o *ApiAccountsSummaryGet200ResponseData) HasCombinedBalances() bool`

HasCombinedBalances returns a boolean if a field has been set.

### GetTotalBalances

`func (o *ApiAccountsSummaryGet200ResponseData) GetTotalBalances() []ApiAccountsSummaryGet200ResponseDataTotalBalancesInner`

GetTotalBalances returns the TotalBalances field if non-nil, zero value otherwise.

### GetTotalBalancesOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetTotalBalancesOk() (*[]ApiAccountsSummaryGet200ResponseDataTotalBalancesInner, bool)`

GetTotalBalancesOk returns a tuple with the TotalBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalances

`func (o *ApiAccountsSummaryGet200ResponseData) SetTotalBalances(v []ApiAccountsSummaryGet200ResponseDataTotalBalancesInner)`

SetTotalBalances sets TotalBalances field to given value.

### HasTotalBalances

`func (o *ApiAccountsSummaryGet200ResponseData) HasTotalBalances() bool`

HasTotalBalances returns a boolean if a field has been set.

### GetTotalHiddenAccounts

`func (o *ApiAccountsSummaryGet200ResponseData) GetTotalHiddenAccounts() float32`

GetTotalHiddenAccounts returns the TotalHiddenAccounts field if non-nil, zero value otherwise.

### GetTotalHiddenAccountsOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetTotalHiddenAccountsOk() (*float32, bool)`

GetTotalHiddenAccountsOk returns a tuple with the TotalHiddenAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHiddenAccounts

`func (o *ApiAccountsSummaryGet200ResponseData) SetTotalHiddenAccounts(v float32)`

SetTotalHiddenAccounts sets TotalHiddenAccounts field to given value.

### HasTotalHiddenAccounts

`func (o *ApiAccountsSummaryGet200ResponseData) HasTotalHiddenAccounts() bool`

HasTotalHiddenAccounts returns a boolean if a field has been set.

### GetIsComboUser

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsComboUser() bool`

GetIsComboUser returns the IsComboUser field if non-nil, zero value otherwise.

### GetIsComboUserOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsComboUserOk() (*bool, bool)`

GetIsComboUserOk returns a tuple with the IsComboUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComboUser

`func (o *ApiAccountsSummaryGet200ResponseData) SetIsComboUser(v bool)`

SetIsComboUser sets IsComboUser field to given value.

### HasIsComboUser

`func (o *ApiAccountsSummaryGet200ResponseData) HasIsComboUser() bool`

HasIsComboUser returns a boolean if a field has been set.

### GetOwnershipTypes

`func (o *ApiAccountsSummaryGet200ResponseData) GetOwnershipTypes() []string`

GetOwnershipTypes returns the OwnershipTypes field if non-nil, zero value otherwise.

### GetOwnershipTypesOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetOwnershipTypesOk() (*[]string, bool)`

GetOwnershipTypesOk returns a tuple with the OwnershipTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipTypes

`func (o *ApiAccountsSummaryGet200ResponseData) SetOwnershipTypes(v []string)`

SetOwnershipTypes sets OwnershipTypes field to given value.

### HasOwnershipTypes

`func (o *ApiAccountsSummaryGet200ResponseData) HasOwnershipTypes() bool`

HasOwnershipTypes returns a boolean if a field has been set.

### GetIsItrade

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsItrade() bool`

GetIsItrade returns the IsItrade field if non-nil, zero value otherwise.

### GetIsItradeOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsItradeOk() (*bool, bool)`

GetIsItradeOk returns a tuple with the IsItrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsItrade

`func (o *ApiAccountsSummaryGet200ResponseData) SetIsItrade(v bool)`

SetIsItrade sets IsItrade field to given value.

### HasIsItrade

`func (o *ApiAccountsSummaryGet200ResponseData) HasIsItrade() bool`

HasIsItrade returns a boolean if a field has been set.

### GetIsWealth

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsWealth() bool`

GetIsWealth returns the IsWealth field if non-nil, zero value otherwise.

### GetIsWealthOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsWealthOk() (*bool, bool)`

GetIsWealthOk returns a tuple with the IsWealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWealth

`func (o *ApiAccountsSummaryGet200ResponseData) SetIsWealth(v bool)`

SetIsWealth sets IsWealth field to given value.

### HasIsWealth

`func (o *ApiAccountsSummaryGet200ResponseData) HasIsWealth() bool`

HasIsWealth returns a boolean if a field has been set.

### GetIsRetail

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsRetail() bool`

GetIsRetail returns the IsRetail field if non-nil, zero value otherwise.

### GetIsRetailOk

`func (o *ApiAccountsSummaryGet200ResponseData) GetIsRetailOk() (*bool, bool)`

GetIsRetailOk returns a tuple with the IsRetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetail

`func (o *ApiAccountsSummaryGet200ResponseData) SetIsRetail(v bool)`

SetIsRetail sets IsRetail field to given value.

### HasIsRetail

`func (o *ApiAccountsSummaryGet200ResponseData) HasIsRetail() bool`

HasIsRetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


