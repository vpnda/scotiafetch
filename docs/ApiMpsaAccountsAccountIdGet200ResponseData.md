# ApiMpsaAccountsAccountIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DaysFromStartDate** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **map[string]interface{}** |  | [optional] 
**PrimaryBalance** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**AvailableBalance** | Pointer to **map[string]interface{}** |  | [optional] 
**Holds** | Pointer to **map[string]interface{}** |  | [optional] 
**Interest** | Pointer to [**ApiMpsaAccountsAccountIdGet200ResponseDataInterest**](ApiMpsaAccountsAccountIdGet200ResponseDataInterest.md) |  | [optional] 
**InterestRates** | Pointer to [**[]ApiMpsaAccountsAccountIdGet200ResponseDataInterestRatesInner**](ApiMpsaAccountsAccountIdGet200ResponseDataInterestRatesInner.md) |  | [optional] 
**BankTheRest** | Pointer to **map[string]interface{}** |  | [optional] 
**OverdraftLimit** | Pointer to **map[string]interface{}** |  | [optional] 
**OverdraftPlan** | Pointer to **map[string]interface{}** |  | [optional] 
**OpeningDate** | Pointer to **map[string]interface{}** |  | [optional] 
**AssociatedAccounts** | Pointer to **[]interface{}** |  | [optional] 
**BusinessFunctions** | Pointer to [**[]ApiMpsaAccountsAccountIdGet200ResponseDataBusinessFunctionsInner**](ApiMpsaAccountsAccountIdGet200ResponseDataBusinessFunctionsInner.md) |  | [optional] 
**ClosingDate** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiMpsaAccountsAccountIdGet200ResponseData

`func NewApiMpsaAccountsAccountIdGet200ResponseData() *ApiMpsaAccountsAccountIdGet200ResponseData`

NewApiMpsaAccountsAccountIdGet200ResponseData instantiates a new ApiMpsaAccountsAccountIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMpsaAccountsAccountIdGet200ResponseDataWithDefaults

`func NewApiMpsaAccountsAccountIdGet200ResponseDataWithDefaults() *ApiMpsaAccountsAccountIdGet200ResponseData`

NewApiMpsaAccountsAccountIdGet200ResponseDataWithDefaults instantiates a new ApiMpsaAccountsAccountIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDaysFromStartDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetDaysFromStartDate() string`

GetDaysFromStartDate returns the DaysFromStartDate field if non-nil, zero value otherwise.

### GetDaysFromStartDateOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetDaysFromStartDateOk() (*string, bool)`

GetDaysFromStartDateOk returns a tuple with the DaysFromStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysFromStartDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetDaysFromStartDate(v string)`

SetDaysFromStartDate sets DaysFromStartDate field to given value.

### HasDaysFromStartDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasDaysFromStartDate() bool`

HasDaysFromStartDate returns a boolean if a field has been set.

### GetAlias

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetAlias() map[string]interface{}`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetAliasOk() (*map[string]interface{}, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetAlias(v map[string]interface{})`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetPrimaryBalance

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetPrimaryBalance() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetPrimaryBalance returns the PrimaryBalance field if non-nil, zero value otherwise.

### GetPrimaryBalanceOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetPrimaryBalanceOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetPrimaryBalanceOk returns a tuple with the PrimaryBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBalance

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetPrimaryBalance(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetPrimaryBalance sets PrimaryBalance field to given value.

### HasPrimaryBalance

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasPrimaryBalance() bool`

HasPrimaryBalance returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetAvailableBalance() map[string]interface{}`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetAvailableBalanceOk() (*map[string]interface{}, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetAvailableBalance(v map[string]interface{})`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil
### GetHolds

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetHolds() map[string]interface{}`

GetHolds returns the Holds field if non-nil, zero value otherwise.

### GetHoldsOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetHoldsOk() (*map[string]interface{}, bool)`

GetHoldsOk returns a tuple with the Holds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolds

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetHolds(v map[string]interface{})`

SetHolds sets Holds field to given value.

### HasHolds

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasHolds() bool`

HasHolds returns a boolean if a field has been set.

### SetHoldsNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetHoldsNil(b bool)`

 SetHoldsNil sets the value for Holds to be an explicit nil

### UnsetHolds
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetHolds()`

UnsetHolds ensures that no value is present for Holds, not even an explicit nil
### GetInterest

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetInterest() ApiMpsaAccountsAccountIdGet200ResponseDataInterest`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetInterestOk() (*ApiMpsaAccountsAccountIdGet200ResponseDataInterest, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetInterest(v ApiMpsaAccountsAccountIdGet200ResponseDataInterest)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestRates

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetInterestRates() []ApiMpsaAccountsAccountIdGet200ResponseDataInterestRatesInner`

GetInterestRates returns the InterestRates field if non-nil, zero value otherwise.

### GetInterestRatesOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetInterestRatesOk() (*[]ApiMpsaAccountsAccountIdGet200ResponseDataInterestRatesInner, bool)`

GetInterestRatesOk returns a tuple with the InterestRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRates

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetInterestRates(v []ApiMpsaAccountsAccountIdGet200ResponseDataInterestRatesInner)`

SetInterestRates sets InterestRates field to given value.

### HasInterestRates

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasInterestRates() bool`

HasInterestRates returns a boolean if a field has been set.

### GetBankTheRest

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetBankTheRest() map[string]interface{}`

GetBankTheRest returns the BankTheRest field if non-nil, zero value otherwise.

### GetBankTheRestOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetBankTheRestOk() (*map[string]interface{}, bool)`

GetBankTheRestOk returns a tuple with the BankTheRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTheRest

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetBankTheRest(v map[string]interface{})`

SetBankTheRest sets BankTheRest field to given value.

### HasBankTheRest

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasBankTheRest() bool`

HasBankTheRest returns a boolean if a field has been set.

### SetBankTheRestNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetBankTheRestNil(b bool)`

 SetBankTheRestNil sets the value for BankTheRest to be an explicit nil

### UnsetBankTheRest
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetBankTheRest()`

UnsetBankTheRest ensures that no value is present for BankTheRest, not even an explicit nil
### GetOverdraftLimit

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetOverdraftLimit() map[string]interface{}`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetOverdraftLimitOk() (*map[string]interface{}, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetOverdraftLimit(v map[string]interface{})`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### SetOverdraftLimitNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetOverdraftLimitNil(b bool)`

 SetOverdraftLimitNil sets the value for OverdraftLimit to be an explicit nil

### UnsetOverdraftLimit
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetOverdraftLimit()`

UnsetOverdraftLimit ensures that no value is present for OverdraftLimit, not even an explicit nil
### GetOverdraftPlan

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetOverdraftPlan() map[string]interface{}`

GetOverdraftPlan returns the OverdraftPlan field if non-nil, zero value otherwise.

### GetOverdraftPlanOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetOverdraftPlanOk() (*map[string]interface{}, bool)`

GetOverdraftPlanOk returns a tuple with the OverdraftPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftPlan

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetOverdraftPlan(v map[string]interface{})`

SetOverdraftPlan sets OverdraftPlan field to given value.

### HasOverdraftPlan

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasOverdraftPlan() bool`

HasOverdraftPlan returns a boolean if a field has been set.

### SetOverdraftPlanNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetOverdraftPlanNil(b bool)`

 SetOverdraftPlanNil sets the value for OverdraftPlan to be an explicit nil

### UnsetOverdraftPlan
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetOverdraftPlan()`

UnsetOverdraftPlan ensures that no value is present for OverdraftPlan, not even an explicit nil
### GetOpeningDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetOpeningDate() map[string]interface{}`

GetOpeningDate returns the OpeningDate field if non-nil, zero value otherwise.

### GetOpeningDateOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetOpeningDateOk() (*map[string]interface{}, bool)`

GetOpeningDateOk returns a tuple with the OpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetOpeningDate(v map[string]interface{})`

SetOpeningDate sets OpeningDate field to given value.

### HasOpeningDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasOpeningDate() bool`

HasOpeningDate returns a boolean if a field has been set.

### SetOpeningDateNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetOpeningDateNil(b bool)`

 SetOpeningDateNil sets the value for OpeningDate to be an explicit nil

### UnsetOpeningDate
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetOpeningDate()`

UnsetOpeningDate ensures that no value is present for OpeningDate, not even an explicit nil
### GetAssociatedAccounts

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetAssociatedAccounts() []interface{}`

GetAssociatedAccounts returns the AssociatedAccounts field if non-nil, zero value otherwise.

### GetAssociatedAccountsOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetAssociatedAccountsOk() (*[]interface{}, bool)`

GetAssociatedAccountsOk returns a tuple with the AssociatedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedAccounts

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetAssociatedAccounts(v []interface{})`

SetAssociatedAccounts sets AssociatedAccounts field to given value.

### HasAssociatedAccounts

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasAssociatedAccounts() bool`

HasAssociatedAccounts returns a boolean if a field has been set.

### GetBusinessFunctions

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetBusinessFunctions() []ApiMpsaAccountsAccountIdGet200ResponseDataBusinessFunctionsInner`

GetBusinessFunctions returns the BusinessFunctions field if non-nil, zero value otherwise.

### GetBusinessFunctionsOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetBusinessFunctionsOk() (*[]ApiMpsaAccountsAccountIdGet200ResponseDataBusinessFunctionsInner, bool)`

GetBusinessFunctionsOk returns a tuple with the BusinessFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessFunctions

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetBusinessFunctions(v []ApiMpsaAccountsAccountIdGet200ResponseDataBusinessFunctionsInner)`

SetBusinessFunctions sets BusinessFunctions field to given value.

### HasBusinessFunctions

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasBusinessFunctions() bool`

HasBusinessFunctions returns a boolean if a field has been set.

### GetClosingDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetClosingDate() map[string]interface{}`

GetClosingDate returns the ClosingDate field if non-nil, zero value otherwise.

### GetClosingDateOk

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) GetClosingDateOk() (*map[string]interface{}, bool)`

GetClosingDateOk returns a tuple with the ClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetClosingDate(v map[string]interface{})`

SetClosingDate sets ClosingDate field to given value.

### HasClosingDate

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) HasClosingDate() bool`

HasClosingDate returns a boolean if a field has been set.

### SetClosingDateNil

`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) SetClosingDateNil(b bool)`

 SetClosingDateNil sets the value for ClosingDate to be an explicit nil

### UnsetClosingDate
`func (o *ApiMpsaAccountsAccountIdGet200ResponseData) UnsetClosingDate()`

UnsetClosingDate ensures that no value is present for ClosingDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


