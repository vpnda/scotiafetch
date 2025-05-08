# ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PostedDate** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SubDescription** | Pointer to **map[string]interface{}** |  | [optional] 
**CleanDescription** | Pointer to **string** |  | [optional] 
**TypeCode** | Pointer to **string** |  | [optional] 
**MnemonicCode** | Pointer to **string** |  | [optional] 
**TsysCode** | Pointer to **map[string]interface{}** |  | [optional] 
**IsTsys** | Pointer to **bool** |  | [optional] 
**IsDisputable** | Pointer to **map[string]interface{}** |  | [optional] 
**OriginalAmount** | Pointer to [**ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerOriginalAmount**](ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerOriginalAmount.md) |  | [optional] 
**RunningBalance** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**AssociatedCardNumber** | Pointer to **map[string]interface{}** |  | [optional] 
**PurchaseCountryCode** | Pointer to **map[string]interface{}** |  | [optional] 
**OutOfCountryIndicator** | Pointer to **bool** |  | [optional] 
**ReasonCode** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 
**RecurringPaymentIndicator** | Pointer to **map[string]interface{}** |  | [optional] 
**DirectionIndicator** | Pointer to **string** |  | [optional] 
**StatementIndicator** | Pointer to **map[string]interface{}** |  | [optional] 
**FromAccount** | Pointer to **map[string]interface{}** |  | [optional] 
**ToAccount** | Pointer to **map[string]interface{}** |  | [optional] 
**PurchaseType** | Pointer to **map[string]interface{}** |  | [optional] 
**RewardsCategory** | Pointer to **map[string]interface{}** |  | [optional] 
**RewardCard** | Pointer to **map[string]interface{}** |  | [optional] 
**Category** | Pointer to [**ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory**](ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory.md) |  | [optional] 
**UserInputTag** | Pointer to **map[string]interface{}** |  | [optional] 
**Cheque** | Pointer to **map[string]interface{}** |  | [optional] 
**Merchant** | Pointer to [**ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant**](ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant.md) |  | [optional] 
**Enriched** | Pointer to **bool** |  | [optional] 
**AcquirerReferenceNumber** | Pointer to **map[string]interface{}** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Subdescription** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**TransactionKey** | Pointer to **string** |  | [optional] 
**TransactionAmount** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**TransactionCategory** | Pointer to **string** |  | [optional] 
**TransactionCategoryCode** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**TransactionTypeCode** | Pointer to **string** |  | [optional] 
**DescriptionLine1** | Pointer to **string** |  | [optional] 
**DescriptionLine2** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner

`func NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner() *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner`

NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner instantiates a new ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerWithDefaults

`func NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerWithDefaults() *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner`

NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerWithDefaults instantiates a new ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPostedDate

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetPostedDate() map[string]interface{}`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetPostedDateOk() (*map[string]interface{}, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetPostedDate(v map[string]interface{})`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### SetPostedDateNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetPostedDateNil(b bool)`

 SetPostedDateNil sets the value for PostedDate to be an explicit nil

### UnsetPostedDate
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetPostedDate()`

UnsetPostedDate ensures that no value is present for PostedDate, not even an explicit nil
### GetDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetSubDescription() map[string]interface{}`

GetSubDescription returns the SubDescription field if non-nil, zero value otherwise.

### GetSubDescriptionOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetSubDescriptionOk() (*map[string]interface{}, bool)`

GetSubDescriptionOk returns a tuple with the SubDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetSubDescription(v map[string]interface{})`

SetSubDescription sets SubDescription field to given value.

### HasSubDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasSubDescription() bool`

HasSubDescription returns a boolean if a field has been set.

### SetSubDescriptionNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetSubDescriptionNil(b bool)`

 SetSubDescriptionNil sets the value for SubDescription to be an explicit nil

### UnsetSubDescription
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetSubDescription()`

UnsetSubDescription ensures that no value is present for SubDescription, not even an explicit nil
### GetCleanDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetCleanDescription() string`

GetCleanDescription returns the CleanDescription field if non-nil, zero value otherwise.

### GetCleanDescriptionOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetCleanDescriptionOk() (*string, bool)`

GetCleanDescriptionOk returns a tuple with the CleanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetCleanDescription(v string)`

SetCleanDescription sets CleanDescription field to given value.

### HasCleanDescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasCleanDescription() bool`

HasCleanDescription returns a boolean if a field has been set.

### GetTypeCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetMnemonicCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetMnemonicCode() string`

GetMnemonicCode returns the MnemonicCode field if non-nil, zero value otherwise.

### GetMnemonicCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetMnemonicCodeOk() (*string, bool)`

GetMnemonicCodeOk returns a tuple with the MnemonicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetMnemonicCode(v string)`

SetMnemonicCode sets MnemonicCode field to given value.

### HasMnemonicCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasMnemonicCode() bool`

HasMnemonicCode returns a boolean if a field has been set.

### GetTsysCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTsysCode() map[string]interface{}`

GetTsysCode returns the TsysCode field if non-nil, zero value otherwise.

### GetTsysCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTsysCodeOk() (*map[string]interface{}, bool)`

GetTsysCodeOk returns a tuple with the TsysCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsysCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTsysCode(v map[string]interface{})`

SetTsysCode sets TsysCode field to given value.

### HasTsysCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTsysCode() bool`

HasTsysCode returns a boolean if a field has been set.

### SetTsysCodeNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTsysCodeNil(b bool)`

 SetTsysCodeNil sets the value for TsysCode to be an explicit nil

### UnsetTsysCode
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetTsysCode()`

UnsetTsysCode ensures that no value is present for TsysCode, not even an explicit nil
### GetIsTsys

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetIsTsys() bool`

GetIsTsys returns the IsTsys field if non-nil, zero value otherwise.

### GetIsTsysOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetIsTsysOk() (*bool, bool)`

GetIsTsysOk returns a tuple with the IsTsys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTsys

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetIsTsys(v bool)`

SetIsTsys sets IsTsys field to given value.

### HasIsTsys

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasIsTsys() bool`

HasIsTsys returns a boolean if a field has been set.

### GetIsDisputable

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetIsDisputable() map[string]interface{}`

GetIsDisputable returns the IsDisputable field if non-nil, zero value otherwise.

### GetIsDisputableOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetIsDisputableOk() (*map[string]interface{}, bool)`

GetIsDisputableOk returns a tuple with the IsDisputable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisputable

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetIsDisputable(v map[string]interface{})`

SetIsDisputable sets IsDisputable field to given value.

### HasIsDisputable

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasIsDisputable() bool`

HasIsDisputable returns a boolean if a field has been set.

### SetIsDisputableNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetIsDisputableNil(b bool)`

 SetIsDisputableNil sets the value for IsDisputable to be an explicit nil

### UnsetIsDisputable
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetIsDisputable()`

UnsetIsDisputable ensures that no value is present for IsDisputable, not even an explicit nil
### GetOriginalAmount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetOriginalAmount() ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerOriginalAmount`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetOriginalAmountOk() (*ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerOriginalAmount, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetOriginalAmount(v ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerOriginalAmount)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetRunningBalance

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRunningBalance() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetRunningBalance returns the RunningBalance field if non-nil, zero value otherwise.

### GetRunningBalanceOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRunningBalanceOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetRunningBalanceOk returns a tuple with the RunningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningBalance

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRunningBalance(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetRunningBalance sets RunningBalance field to given value.

### HasRunningBalance

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasRunningBalance() bool`

HasRunningBalance returns a boolean if a field has been set.

### GetAssociatedCardNumber

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetAssociatedCardNumber() map[string]interface{}`

GetAssociatedCardNumber returns the AssociatedCardNumber field if non-nil, zero value otherwise.

### GetAssociatedCardNumberOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetAssociatedCardNumberOk() (*map[string]interface{}, bool)`

GetAssociatedCardNumberOk returns a tuple with the AssociatedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCardNumber

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetAssociatedCardNumber(v map[string]interface{})`

SetAssociatedCardNumber sets AssociatedCardNumber field to given value.

### HasAssociatedCardNumber

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasAssociatedCardNumber() bool`

HasAssociatedCardNumber returns a boolean if a field has been set.

### SetAssociatedCardNumberNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetAssociatedCardNumberNil(b bool)`

 SetAssociatedCardNumberNil sets the value for AssociatedCardNumber to be an explicit nil

### UnsetAssociatedCardNumber
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetAssociatedCardNumber()`

UnsetAssociatedCardNumber ensures that no value is present for AssociatedCardNumber, not even an explicit nil
### GetPurchaseCountryCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetPurchaseCountryCode() map[string]interface{}`

GetPurchaseCountryCode returns the PurchaseCountryCode field if non-nil, zero value otherwise.

### GetPurchaseCountryCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetPurchaseCountryCodeOk() (*map[string]interface{}, bool)`

GetPurchaseCountryCodeOk returns a tuple with the PurchaseCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCountryCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetPurchaseCountryCode(v map[string]interface{})`

SetPurchaseCountryCode sets PurchaseCountryCode field to given value.

### HasPurchaseCountryCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasPurchaseCountryCode() bool`

HasPurchaseCountryCode returns a boolean if a field has been set.

### SetPurchaseCountryCodeNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetPurchaseCountryCodeNil(b bool)`

 SetPurchaseCountryCodeNil sets the value for PurchaseCountryCode to be an explicit nil

### UnsetPurchaseCountryCode
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetPurchaseCountryCode()`

UnsetPurchaseCountryCode ensures that no value is present for PurchaseCountryCode, not even an explicit nil
### GetOutOfCountryIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetOutOfCountryIndicator() bool`

GetOutOfCountryIndicator returns the OutOfCountryIndicator field if non-nil, zero value otherwise.

### GetOutOfCountryIndicatorOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetOutOfCountryIndicatorOk() (*bool, bool)`

GetOutOfCountryIndicatorOk returns a tuple with the OutOfCountryIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfCountryIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetOutOfCountryIndicator(v bool)`

SetOutOfCountryIndicator sets OutOfCountryIndicator field to given value.

### HasOutOfCountryIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasOutOfCountryIndicator() bool`

HasOutOfCountryIndicator returns a boolean if a field has been set.

### GetReasonCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetReasonCode() map[string]interface{}`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetReasonCodeOk() (*map[string]interface{}, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetReasonCode(v map[string]interface{})`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetStatus

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRecurringPaymentIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRecurringPaymentIndicator() map[string]interface{}`

GetRecurringPaymentIndicator returns the RecurringPaymentIndicator field if non-nil, zero value otherwise.

### GetRecurringPaymentIndicatorOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRecurringPaymentIndicatorOk() (*map[string]interface{}, bool)`

GetRecurringPaymentIndicatorOk returns a tuple with the RecurringPaymentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPaymentIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRecurringPaymentIndicator(v map[string]interface{})`

SetRecurringPaymentIndicator sets RecurringPaymentIndicator field to given value.

### HasRecurringPaymentIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasRecurringPaymentIndicator() bool`

HasRecurringPaymentIndicator returns a boolean if a field has been set.

### SetRecurringPaymentIndicatorNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRecurringPaymentIndicatorNil(b bool)`

 SetRecurringPaymentIndicatorNil sets the value for RecurringPaymentIndicator to be an explicit nil

### UnsetRecurringPaymentIndicator
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetRecurringPaymentIndicator()`

UnsetRecurringPaymentIndicator ensures that no value is present for RecurringPaymentIndicator, not even an explicit nil
### GetDirectionIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDirectionIndicator() string`

GetDirectionIndicator returns the DirectionIndicator field if non-nil, zero value otherwise.

### GetDirectionIndicatorOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDirectionIndicatorOk() (*string, bool)`

GetDirectionIndicatorOk returns a tuple with the DirectionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetDirectionIndicator(v string)`

SetDirectionIndicator sets DirectionIndicator field to given value.

### HasDirectionIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasDirectionIndicator() bool`

HasDirectionIndicator returns a boolean if a field has been set.

### GetStatementIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetStatementIndicator() map[string]interface{}`

GetStatementIndicator returns the StatementIndicator field if non-nil, zero value otherwise.

### GetStatementIndicatorOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetStatementIndicatorOk() (*map[string]interface{}, bool)`

GetStatementIndicatorOk returns a tuple with the StatementIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetStatementIndicator(v map[string]interface{})`

SetStatementIndicator sets StatementIndicator field to given value.

### HasStatementIndicator

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasStatementIndicator() bool`

HasStatementIndicator returns a boolean if a field has been set.

### SetStatementIndicatorNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetStatementIndicatorNil(b bool)`

 SetStatementIndicatorNil sets the value for StatementIndicator to be an explicit nil

### UnsetStatementIndicator
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetStatementIndicator()`

UnsetStatementIndicator ensures that no value is present for StatementIndicator, not even an explicit nil
### GetFromAccount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetFromAccount() map[string]interface{}`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetFromAccountOk() (*map[string]interface{}, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetFromAccount(v map[string]interface{})`

SetFromAccount sets FromAccount field to given value.

### HasFromAccount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasFromAccount() bool`

HasFromAccount returns a boolean if a field has been set.

### SetFromAccountNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetFromAccountNil(b bool)`

 SetFromAccountNil sets the value for FromAccount to be an explicit nil

### UnsetFromAccount
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetFromAccount()`

UnsetFromAccount ensures that no value is present for FromAccount, not even an explicit nil
### GetToAccount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetToAccount() map[string]interface{}`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetToAccountOk() (*map[string]interface{}, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetToAccount(v map[string]interface{})`

SetToAccount sets ToAccount field to given value.

### HasToAccount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasToAccount() bool`

HasToAccount returns a boolean if a field has been set.

### SetToAccountNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetToAccountNil(b bool)`

 SetToAccountNil sets the value for ToAccount to be an explicit nil

### UnsetToAccount
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetToAccount()`

UnsetToAccount ensures that no value is present for ToAccount, not even an explicit nil
### GetPurchaseType

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetPurchaseType() map[string]interface{}`

GetPurchaseType returns the PurchaseType field if non-nil, zero value otherwise.

### GetPurchaseTypeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetPurchaseTypeOk() (*map[string]interface{}, bool)`

GetPurchaseTypeOk returns a tuple with the PurchaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseType

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetPurchaseType(v map[string]interface{})`

SetPurchaseType sets PurchaseType field to given value.

### HasPurchaseType

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasPurchaseType() bool`

HasPurchaseType returns a boolean if a field has been set.

### SetPurchaseTypeNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetPurchaseTypeNil(b bool)`

 SetPurchaseTypeNil sets the value for PurchaseType to be an explicit nil

### UnsetPurchaseType
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetPurchaseType()`

UnsetPurchaseType ensures that no value is present for PurchaseType, not even an explicit nil
### GetRewardsCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRewardsCategory() map[string]interface{}`

GetRewardsCategory returns the RewardsCategory field if non-nil, zero value otherwise.

### GetRewardsCategoryOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRewardsCategoryOk() (*map[string]interface{}, bool)`

GetRewardsCategoryOk returns a tuple with the RewardsCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRewardsCategory(v map[string]interface{})`

SetRewardsCategory sets RewardsCategory field to given value.

### HasRewardsCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasRewardsCategory() bool`

HasRewardsCategory returns a boolean if a field has been set.

### SetRewardsCategoryNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRewardsCategoryNil(b bool)`

 SetRewardsCategoryNil sets the value for RewardsCategory to be an explicit nil

### UnsetRewardsCategory
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetRewardsCategory()`

UnsetRewardsCategory ensures that no value is present for RewardsCategory, not even an explicit nil
### GetRewardCard

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRewardCard() map[string]interface{}`

GetRewardCard returns the RewardCard field if non-nil, zero value otherwise.

### GetRewardCardOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetRewardCardOk() (*map[string]interface{}, bool)`

GetRewardCardOk returns a tuple with the RewardCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardCard

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRewardCard(v map[string]interface{})`

SetRewardCard sets RewardCard field to given value.

### HasRewardCard

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasRewardCard() bool`

HasRewardCard returns a boolean if a field has been set.

### SetRewardCardNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetRewardCardNil(b bool)`

 SetRewardCardNil sets the value for RewardCard to be an explicit nil

### UnsetRewardCard
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetRewardCard()`

UnsetRewardCard ensures that no value is present for RewardCard, not even an explicit nil
### GetCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetCategory() ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetCategoryOk() (*ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetCategory(v ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetUserInputTag

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetUserInputTag() map[string]interface{}`

GetUserInputTag returns the UserInputTag field if non-nil, zero value otherwise.

### GetUserInputTagOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetUserInputTagOk() (*map[string]interface{}, bool)`

GetUserInputTagOk returns a tuple with the UserInputTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputTag

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetUserInputTag(v map[string]interface{})`

SetUserInputTag sets UserInputTag field to given value.

### HasUserInputTag

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasUserInputTag() bool`

HasUserInputTag returns a boolean if a field has been set.

### SetUserInputTagNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetUserInputTagNil(b bool)`

 SetUserInputTagNil sets the value for UserInputTag to be an explicit nil

### UnsetUserInputTag
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetUserInputTag()`

UnsetUserInputTag ensures that no value is present for UserInputTag, not even an explicit nil
### GetCheque

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetCheque() map[string]interface{}`

GetCheque returns the Cheque field if non-nil, zero value otherwise.

### GetChequeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetChequeOk() (*map[string]interface{}, bool)`

GetChequeOk returns a tuple with the Cheque field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheque

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetCheque(v map[string]interface{})`

SetCheque sets Cheque field to given value.

### HasCheque

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasCheque() bool`

HasCheque returns a boolean if a field has been set.

### SetChequeNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetChequeNil(b bool)`

 SetChequeNil sets the value for Cheque to be an explicit nil

### UnsetCheque
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetCheque()`

UnsetCheque ensures that no value is present for Cheque, not even an explicit nil
### GetMerchant

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetMerchant() ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetMerchantOk() (*ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetMerchant(v ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetEnriched

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetEnriched() bool`

GetEnriched returns the Enriched field if non-nil, zero value otherwise.

### GetEnrichedOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetEnrichedOk() (*bool, bool)`

GetEnrichedOk returns a tuple with the Enriched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnriched

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetEnriched(v bool)`

SetEnriched sets Enriched field to given value.

### HasEnriched

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasEnriched() bool`

HasEnriched returns a boolean if a field has been set.

### GetAcquirerReferenceNumber

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetAcquirerReferenceNumber() map[string]interface{}`

GetAcquirerReferenceNumber returns the AcquirerReferenceNumber field if non-nil, zero value otherwise.

### GetAcquirerReferenceNumberOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetAcquirerReferenceNumberOk() (*map[string]interface{}, bool)`

GetAcquirerReferenceNumberOk returns a tuple with the AcquirerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceNumber

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetAcquirerReferenceNumber(v map[string]interface{})`

SetAcquirerReferenceNumber sets AcquirerReferenceNumber field to given value.

### HasAcquirerReferenceNumber

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasAcquirerReferenceNumber() bool`

HasAcquirerReferenceNumber returns a boolean if a field has been set.

### SetAcquirerReferenceNumberNil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetAcquirerReferenceNumberNil(b bool)`

 SetAcquirerReferenceNumberNil sets the value for AcquirerReferenceNumber to be an explicit nil

### UnsetAcquirerReferenceNumber
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetAcquirerReferenceNumber()`

UnsetAcquirerReferenceNumber ensures that no value is present for AcquirerReferenceNumber, not even an explicit nil
### GetTransactionId

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetSubdescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetSubdescription() string`

GetSubdescription returns the Subdescription field if non-nil, zero value otherwise.

### GetSubdescriptionOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetSubdescriptionOk() (*string, bool)`

GetSubdescriptionOk returns a tuple with the Subdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetSubdescription(v string)`

SetSubdescription sets Subdescription field to given value.

### HasSubdescription

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasSubdescription() bool`

HasSubdescription returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetBalance

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetBalance() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetBalanceOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetBalance(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetTransactionKey

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionKey() string`

GetTransactionKey returns the TransactionKey field if non-nil, zero value otherwise.

### GetTransactionKeyOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionKeyOk() (*string, bool)`

GetTransactionKeyOk returns a tuple with the TransactionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionKey

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionKey(v string)`

SetTransactionKey sets TransactionKey field to given value.

### HasTransactionKey

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionKey() bool`

HasTransactionKey returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionAmount() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionAmountOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionAmount(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionCategory() string`

GetTransactionCategory returns the TransactionCategory field if non-nil, zero value otherwise.

### GetTransactionCategoryOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionCategoryOk() (*string, bool)`

GetTransactionCategoryOk returns a tuple with the TransactionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionCategory(v string)`

SetTransactionCategory sets TransactionCategory field to given value.

### HasTransactionCategory

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionCategory() bool`

HasTransactionCategory returns a boolean if a field has been set.

### GetTransactionCategoryCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionCategoryCode() string`

GetTransactionCategoryCode returns the TransactionCategoryCode field if non-nil, zero value otherwise.

### GetTransactionCategoryCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionCategoryCodeOk() (*string, bool)`

GetTransactionCategoryCodeOk returns a tuple with the TransactionCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionCategoryCode(v string)`

SetTransactionCategoryCode sets TransactionCategoryCode field to given value.

### HasTransactionCategoryCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionCategoryCode() bool`

HasTransactionCategoryCode returns a boolean if a field has been set.

### GetTransactionType

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionTypeCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionTypeCode() string`

GetTransactionTypeCode returns the TransactionTypeCode field if non-nil, zero value otherwise.

### GetTransactionTypeCodeOk

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetTransactionTypeCodeOk() (*string, bool)`

GetTransactionTypeCodeOk returns a tuple with the TransactionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTypeCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetTransactionTypeCode(v string)`

SetTransactionTypeCode sets TransactionTypeCode field to given value.

### HasTransactionTypeCode

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasTransactionTypeCode() bool`

HasTransactionTypeCode returns a boolean if a field has been set.

### GetDescriptionLine1

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDescriptionLine1() string`

GetDescriptionLine1 returns the DescriptionLine1 field if non-nil, zero value otherwise.

### GetDescriptionLine1Ok

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDescriptionLine1Ok() (*string, bool)`

GetDescriptionLine1Ok returns a tuple with the DescriptionLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLine1

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetDescriptionLine1(v string)`

SetDescriptionLine1 sets DescriptionLine1 field to given value.

### HasDescriptionLine1

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasDescriptionLine1() bool`

HasDescriptionLine1 returns a boolean if a field has been set.

### GetDescriptionLine2

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDescriptionLine2() map[string]interface{}`

GetDescriptionLine2 returns the DescriptionLine2 field if non-nil, zero value otherwise.

### GetDescriptionLine2Ok

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) GetDescriptionLine2Ok() (*map[string]interface{}, bool)`

GetDescriptionLine2Ok returns a tuple with the DescriptionLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLine2

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetDescriptionLine2(v map[string]interface{})`

SetDescriptionLine2 sets DescriptionLine2 field to given value.

### HasDescriptionLine2

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) HasDescriptionLine2() bool`

HasDescriptionLine2 returns a boolean if a field has been set.

### SetDescriptionLine2Nil

`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) SetDescriptionLine2Nil(b bool)`

 SetDescriptionLine2Nil sets the value for DescriptionLine2 to be an explicit nil

### UnsetDescriptionLine2
`func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInner) UnsetDescriptionLine2()`

UnsetDescriptionLine2 ensures that no value is present for DescriptionLine2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


