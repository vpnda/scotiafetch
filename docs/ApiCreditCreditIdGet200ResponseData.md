# ApiCreditCreditIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ProductCategory** | Pointer to **string** |  | [optional] 
**DisplayId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PrimaryBalance** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**CreditLimit** | Pointer to [**[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**AvailableCredit** | Pointer to [**[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**BalanceOwed** | Pointer to [**[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**LastPaymentDate** | Pointer to **string** |  | [optional] 
**NextDueDate** | Pointer to **string** |  | [optional] 
**LastStatementBalance** | Pointer to [**[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**MinimumMonthlyPayment** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**LastPayment** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**TotalDue** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**InstallmentPlanEligibility** | Pointer to [**ApiCreditCreditIdGet200ResponseDataInstallmentPlanEligibility**](ApiCreditCreditIdGet200ResponseDataInstallmentPlanEligibility.md) |  | [optional] 
**SelectpayEligible** | Pointer to **bool** |  | [optional] 
**BusinessFunctions** | Pointer to [**[]ApiCreditCreditIdGet200ResponseDataBusinessFunctionsInner**](ApiCreditCreditIdGet200ResponseDataBusinessFunctionsInner.md) |  | [optional] 
**OpeningDate** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**PbpInfo** | Pointer to [**ApiCreditCreditIdGet200ResponseDataPbpInfo**](ApiCreditCreditIdGet200ResponseDataPbpInfo.md) |  | [optional] 
**InterestRate** | Pointer to **float32** |  | [optional] 
**CashAdvanceRate** | Pointer to **float32** |  | [optional] 
**CashAdvRateAdjFactor** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCreditCreditIdGet200ResponseData

`func NewApiCreditCreditIdGet200ResponseData() *ApiCreditCreditIdGet200ResponseData`

NewApiCreditCreditIdGet200ResponseData instantiates a new ApiCreditCreditIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCreditCreditIdGet200ResponseDataWithDefaults

`func NewApiCreditCreditIdGet200ResponseDataWithDefaults() *ApiCreditCreditIdGet200ResponseData`

NewApiCreditCreditIdGet200ResponseDataWithDefaults instantiates a new ApiCreditCreditIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiCreditCreditIdGet200ResponseData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiCreditCreditIdGet200ResponseData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiCreditCreditIdGet200ResponseData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *ApiCreditCreditIdGet200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCreditCreditIdGet200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiCreditCreditIdGet200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApiCreditCreditIdGet200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiCreditCreditIdGet200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiCreditCreditIdGet200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProductCategory

`func (o *ApiCreditCreditIdGet200ResponseData) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ApiCreditCreditIdGet200ResponseData) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *ApiCreditCreditIdGet200ResponseData) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### GetDisplayId

`func (o *ApiCreditCreditIdGet200ResponseData) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *ApiCreditCreditIdGet200ResponseData) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *ApiCreditCreditIdGet200ResponseData) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetDescription

`func (o *ApiCreditCreditIdGet200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCreditCreditIdGet200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCreditCreditIdGet200ResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimaryBalance

`func (o *ApiCreditCreditIdGet200ResponseData) GetPrimaryBalance() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetPrimaryBalance returns the PrimaryBalance field if non-nil, zero value otherwise.

### GetPrimaryBalanceOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetPrimaryBalanceOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetPrimaryBalanceOk returns a tuple with the PrimaryBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBalance

`func (o *ApiCreditCreditIdGet200ResponseData) SetPrimaryBalance(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetPrimaryBalance sets PrimaryBalance field to given value.

### HasPrimaryBalance

`func (o *ApiCreditCreditIdGet200ResponseData) HasPrimaryBalance() bool`

HasPrimaryBalance returns a boolean if a field has been set.

### GetCreditLimit

`func (o *ApiCreditCreditIdGet200ResponseData) GetCreditLimit() []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetCreditLimitOk() (*[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *ApiCreditCreditIdGet200ResponseData) SetCreditLimit(v []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *ApiCreditCreditIdGet200ResponseData) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetAvailableCredit

`func (o *ApiCreditCreditIdGet200ResponseData) GetAvailableCredit() []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetAvailableCreditOk() (*[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *ApiCreditCreditIdGet200ResponseData) SetAvailableCredit(v []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetAvailableCredit sets AvailableCredit field to given value.

### HasAvailableCredit

`func (o *ApiCreditCreditIdGet200ResponseData) HasAvailableCredit() bool`

HasAvailableCredit returns a boolean if a field has been set.

### GetBalanceOwed

`func (o *ApiCreditCreditIdGet200ResponseData) GetBalanceOwed() []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetBalanceOwed returns the BalanceOwed field if non-nil, zero value otherwise.

### GetBalanceOwedOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetBalanceOwedOk() (*[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetBalanceOwedOk returns a tuple with the BalanceOwed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceOwed

`func (o *ApiCreditCreditIdGet200ResponseData) SetBalanceOwed(v []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetBalanceOwed sets BalanceOwed field to given value.

### HasBalanceOwed

`func (o *ApiCreditCreditIdGet200ResponseData) HasBalanceOwed() bool`

HasBalanceOwed returns a boolean if a field has been set.

### GetLastPaymentDate

`func (o *ApiCreditCreditIdGet200ResponseData) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *ApiCreditCreditIdGet200ResponseData) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.

### HasLastPaymentDate

`func (o *ApiCreditCreditIdGet200ResponseData) HasLastPaymentDate() bool`

HasLastPaymentDate returns a boolean if a field has been set.

### GetNextDueDate

`func (o *ApiCreditCreditIdGet200ResponseData) GetNextDueDate() string`

GetNextDueDate returns the NextDueDate field if non-nil, zero value otherwise.

### GetNextDueDateOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetNextDueDateOk() (*string, bool)`

GetNextDueDateOk returns a tuple with the NextDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDueDate

`func (o *ApiCreditCreditIdGet200ResponseData) SetNextDueDate(v string)`

SetNextDueDate sets NextDueDate field to given value.

### HasNextDueDate

`func (o *ApiCreditCreditIdGet200ResponseData) HasNextDueDate() bool`

HasNextDueDate returns a boolean if a field has been set.

### GetLastStatementBalance

`func (o *ApiCreditCreditIdGet200ResponseData) GetLastStatementBalance() []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetLastStatementBalance returns the LastStatementBalance field if non-nil, zero value otherwise.

### GetLastStatementBalanceOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetLastStatementBalanceOk() (*[]ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetLastStatementBalanceOk returns a tuple with the LastStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementBalance

`func (o *ApiCreditCreditIdGet200ResponseData) SetLastStatementBalance(v []ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetLastStatementBalance sets LastStatementBalance field to given value.

### HasLastStatementBalance

`func (o *ApiCreditCreditIdGet200ResponseData) HasLastStatementBalance() bool`

HasLastStatementBalance returns a boolean if a field has been set.

### GetMinimumMonthlyPayment

`func (o *ApiCreditCreditIdGet200ResponseData) GetMinimumMonthlyPayment() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetMinimumMonthlyPayment returns the MinimumMonthlyPayment field if non-nil, zero value otherwise.

### GetMinimumMonthlyPaymentOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetMinimumMonthlyPaymentOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetMinimumMonthlyPaymentOk returns a tuple with the MinimumMonthlyPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumMonthlyPayment

`func (o *ApiCreditCreditIdGet200ResponseData) SetMinimumMonthlyPayment(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetMinimumMonthlyPayment sets MinimumMonthlyPayment field to given value.

### HasMinimumMonthlyPayment

`func (o *ApiCreditCreditIdGet200ResponseData) HasMinimumMonthlyPayment() bool`

HasMinimumMonthlyPayment returns a boolean if a field has been set.

### GetLastPayment

`func (o *ApiCreditCreditIdGet200ResponseData) GetLastPayment() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetLastPayment returns the LastPayment field if non-nil, zero value otherwise.

### GetLastPaymentOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetLastPaymentOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetLastPaymentOk returns a tuple with the LastPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPayment

`func (o *ApiCreditCreditIdGet200ResponseData) SetLastPayment(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetLastPayment sets LastPayment field to given value.

### HasLastPayment

`func (o *ApiCreditCreditIdGet200ResponseData) HasLastPayment() bool`

HasLastPayment returns a boolean if a field has been set.

### GetTotalDue

`func (o *ApiCreditCreditIdGet200ResponseData) GetTotalDue() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetTotalDue returns the TotalDue field if non-nil, zero value otherwise.

### GetTotalDueOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetTotalDueOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetTotalDueOk returns a tuple with the TotalDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDue

`func (o *ApiCreditCreditIdGet200ResponseData) SetTotalDue(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetTotalDue sets TotalDue field to given value.

### HasTotalDue

`func (o *ApiCreditCreditIdGet200ResponseData) HasTotalDue() bool`

HasTotalDue returns a boolean if a field has been set.

### GetInstallmentPlanEligibility

`func (o *ApiCreditCreditIdGet200ResponseData) GetInstallmentPlanEligibility() ApiCreditCreditIdGet200ResponseDataInstallmentPlanEligibility`

GetInstallmentPlanEligibility returns the InstallmentPlanEligibility field if non-nil, zero value otherwise.

### GetInstallmentPlanEligibilityOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetInstallmentPlanEligibilityOk() (*ApiCreditCreditIdGet200ResponseDataInstallmentPlanEligibility, bool)`

GetInstallmentPlanEligibilityOk returns a tuple with the InstallmentPlanEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPlanEligibility

`func (o *ApiCreditCreditIdGet200ResponseData) SetInstallmentPlanEligibility(v ApiCreditCreditIdGet200ResponseDataInstallmentPlanEligibility)`

SetInstallmentPlanEligibility sets InstallmentPlanEligibility field to given value.

### HasInstallmentPlanEligibility

`func (o *ApiCreditCreditIdGet200ResponseData) HasInstallmentPlanEligibility() bool`

HasInstallmentPlanEligibility returns a boolean if a field has been set.

### GetSelectpayEligible

`func (o *ApiCreditCreditIdGet200ResponseData) GetSelectpayEligible() bool`

GetSelectpayEligible returns the SelectpayEligible field if non-nil, zero value otherwise.

### GetSelectpayEligibleOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetSelectpayEligibleOk() (*bool, bool)`

GetSelectpayEligibleOk returns a tuple with the SelectpayEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectpayEligible

`func (o *ApiCreditCreditIdGet200ResponseData) SetSelectpayEligible(v bool)`

SetSelectpayEligible sets SelectpayEligible field to given value.

### HasSelectpayEligible

`func (o *ApiCreditCreditIdGet200ResponseData) HasSelectpayEligible() bool`

HasSelectpayEligible returns a boolean if a field has been set.

### GetBusinessFunctions

`func (o *ApiCreditCreditIdGet200ResponseData) GetBusinessFunctions() []ApiCreditCreditIdGet200ResponseDataBusinessFunctionsInner`

GetBusinessFunctions returns the BusinessFunctions field if non-nil, zero value otherwise.

### GetBusinessFunctionsOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetBusinessFunctionsOk() (*[]ApiCreditCreditIdGet200ResponseDataBusinessFunctionsInner, bool)`

GetBusinessFunctionsOk returns a tuple with the BusinessFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessFunctions

`func (o *ApiCreditCreditIdGet200ResponseData) SetBusinessFunctions(v []ApiCreditCreditIdGet200ResponseDataBusinessFunctionsInner)`

SetBusinessFunctions sets BusinessFunctions field to given value.

### HasBusinessFunctions

`func (o *ApiCreditCreditIdGet200ResponseData) HasBusinessFunctions() bool`

HasBusinessFunctions returns a boolean if a field has been set.

### GetOpeningDate

`func (o *ApiCreditCreditIdGet200ResponseData) GetOpeningDate() string`

GetOpeningDate returns the OpeningDate field if non-nil, zero value otherwise.

### GetOpeningDateOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetOpeningDateOk() (*string, bool)`

GetOpeningDateOk returns a tuple with the OpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDate

`func (o *ApiCreditCreditIdGet200ResponseData) SetOpeningDate(v string)`

SetOpeningDate sets OpeningDate field to given value.

### HasOpeningDate

`func (o *ApiCreditCreditIdGet200ResponseData) HasOpeningDate() bool`

HasOpeningDate returns a boolean if a field has been set.

### GetProvider

`func (o *ApiCreditCreditIdGet200ResponseData) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiCreditCreditIdGet200ResponseData) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApiCreditCreditIdGet200ResponseData) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPbpInfo

`func (o *ApiCreditCreditIdGet200ResponseData) GetPbpInfo() ApiCreditCreditIdGet200ResponseDataPbpInfo`

GetPbpInfo returns the PbpInfo field if non-nil, zero value otherwise.

### GetPbpInfoOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetPbpInfoOk() (*ApiCreditCreditIdGet200ResponseDataPbpInfo, bool)`

GetPbpInfoOk returns a tuple with the PbpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbpInfo

`func (o *ApiCreditCreditIdGet200ResponseData) SetPbpInfo(v ApiCreditCreditIdGet200ResponseDataPbpInfo)`

SetPbpInfo sets PbpInfo field to given value.

### HasPbpInfo

`func (o *ApiCreditCreditIdGet200ResponseData) HasPbpInfo() bool`

HasPbpInfo returns a boolean if a field has been set.

### GetInterestRate

`func (o *ApiCreditCreditIdGet200ResponseData) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *ApiCreditCreditIdGet200ResponseData) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *ApiCreditCreditIdGet200ResponseData) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetCashAdvanceRate

`func (o *ApiCreditCreditIdGet200ResponseData) GetCashAdvanceRate() float32`

GetCashAdvanceRate returns the CashAdvanceRate field if non-nil, zero value otherwise.

### GetCashAdvanceRateOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetCashAdvanceRateOk() (*float32, bool)`

GetCashAdvanceRateOk returns a tuple with the CashAdvanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAdvanceRate

`func (o *ApiCreditCreditIdGet200ResponseData) SetCashAdvanceRate(v float32)`

SetCashAdvanceRate sets CashAdvanceRate field to given value.

### HasCashAdvanceRate

`func (o *ApiCreditCreditIdGet200ResponseData) HasCashAdvanceRate() bool`

HasCashAdvanceRate returns a boolean if a field has been set.

### GetCashAdvRateAdjFactor

`func (o *ApiCreditCreditIdGet200ResponseData) GetCashAdvRateAdjFactor() string`

GetCashAdvRateAdjFactor returns the CashAdvRateAdjFactor field if non-nil, zero value otherwise.

### GetCashAdvRateAdjFactorOk

`func (o *ApiCreditCreditIdGet200ResponseData) GetCashAdvRateAdjFactorOk() (*string, bool)`

GetCashAdvRateAdjFactorOk returns a tuple with the CashAdvRateAdjFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAdvRateAdjFactor

`func (o *ApiCreditCreditIdGet200ResponseData) SetCashAdvRateAdjFactor(v string)`

SetCashAdvRateAdjFactor sets CashAdvRateAdjFactor field to given value.

### HasCashAdvRateAdjFactor

`func (o *ApiCreditCreditIdGet200ResponseData) HasCashAdvRateAdjFactor() bool`

HasCashAdvRateAdjFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


