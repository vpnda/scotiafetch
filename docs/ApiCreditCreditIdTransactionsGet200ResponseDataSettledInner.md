# ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PostedDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionLines** | Pointer to **NullableString** |  | [optional] 
**SubDescription** | Pointer to **string** |  | [optional] 
**CleanDescription** | Pointer to **string** |  | [optional] 
**MnemonicCode** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TsysCode** | Pointer to **string** |  | [optional] 
**IsTsys** | Pointer to **bool** |  | [optional] 
**IsDisputable** | Pointer to **string** |  | [optional] 
**OriginalAmount** | Pointer to **map[string]interface{}** |  | [optional] 
**RunningBalance** | Pointer to **NullableString** |  | [optional] 
**AssociatedCardNumber** | Pointer to **string** |  | [optional] 
**PurchaseCountryCode** | Pointer to **NullableString** |  | [optional] 
**OutOfCountryIndicator** | Pointer to **NullableString** |  | [optional] 
**ReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RecurringPaymentIndicator** | Pointer to **NullableString** |  | [optional] 
**StatementIndicator** | Pointer to **bool** |  | [optional] 
**FromAccount** | Pointer to **NullableString** |  | [optional] 
**ToAccount** | Pointer to **NullableString** |  | [optional] 
**PurchaseType** | Pointer to **NullableString** |  | [optional] 
**RewardsCategory** | Pointer to **string** |  | [optional] 
**RewardCard** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to [**ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory**](ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory.md) |  | [optional] 
**UserInputTag** | Pointer to **NullableString** |  | [optional] 
**Cheque** | Pointer to **NullableString** |  | [optional] 
**Merchant** | Pointer to [**ApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerMerchant**](ApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerMerchant.md) |  | [optional] 
**Enriched** | Pointer to **bool** |  | [optional] 
**AcquirerReferenceNumber** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TransactionKey** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 
**TransactionCode** | Pointer to **string** |  | [optional] 
**DescriptionLine1** | Pointer to **string** |  | [optional] 
**DescriptionLine2** | Pointer to **string** |  | [optional] 
**TransactionAmount** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**TransactionCategory** | Pointer to **string** |  | [optional] 
**TransactionCategoryCode** | Pointer to **string** |  | [optional] 
**Subdescription** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCreditCreditIdTransactionsGet200ResponseDataSettledInner

`func NewApiCreditCreditIdTransactionsGet200ResponseDataSettledInner() *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner`

NewApiCreditCreditIdTransactionsGet200ResponseDataSettledInner instantiates a new ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerWithDefaults

`func NewApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerWithDefaults() *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner`

NewApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerWithDefaults instantiates a new ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPostedDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionLines

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionLines() string`

GetDescriptionLines returns the DescriptionLines field if non-nil, zero value otherwise.

### GetDescriptionLinesOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionLinesOk() (*string, bool)`

GetDescriptionLinesOk returns a tuple with the DescriptionLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLines

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetDescriptionLines(v string)`

SetDescriptionLines sets DescriptionLines field to given value.

### HasDescriptionLines

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasDescriptionLines() bool`

HasDescriptionLines returns a boolean if a field has been set.

### SetDescriptionLinesNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetDescriptionLinesNil(b bool)`

 SetDescriptionLinesNil sets the value for DescriptionLines to be an explicit nil

### UnsetDescriptionLines
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetDescriptionLines()`

UnsetDescriptionLines ensures that no value is present for DescriptionLines, not even an explicit nil
### GetSubDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetSubDescription() string`

GetSubDescription returns the SubDescription field if non-nil, zero value otherwise.

### GetSubDescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetSubDescriptionOk() (*string, bool)`

GetSubDescriptionOk returns a tuple with the SubDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetSubDescription(v string)`

SetSubDescription sets SubDescription field to given value.

### HasSubDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasSubDescription() bool`

HasSubDescription returns a boolean if a field has been set.

### GetCleanDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetCleanDescription() string`

GetCleanDescription returns the CleanDescription field if non-nil, zero value otherwise.

### GetCleanDescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetCleanDescriptionOk() (*string, bool)`

GetCleanDescriptionOk returns a tuple with the CleanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetCleanDescription(v string)`

SetCleanDescription sets CleanDescription field to given value.

### HasCleanDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasCleanDescription() bool`

HasCleanDescription returns a boolean if a field has been set.

### GetMnemonicCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetMnemonicCode() string`

GetMnemonicCode returns the MnemonicCode field if non-nil, zero value otherwise.

### GetMnemonicCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetMnemonicCodeOk() (*string, bool)`

GetMnemonicCodeOk returns a tuple with the MnemonicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetMnemonicCode(v string)`

SetMnemonicCode sets MnemonicCode field to given value.

### HasMnemonicCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasMnemonicCode() bool`

HasMnemonicCode returns a boolean if a field has been set.

### SetMnemonicCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetMnemonicCodeNil(b bool)`

 SetMnemonicCodeNil sets the value for MnemonicCode to be an explicit nil

### UnsetMnemonicCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetMnemonicCode()`

UnsetMnemonicCode ensures that no value is present for MnemonicCode, not even an explicit nil
### GetType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTsysCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTsysCode() string`

GetTsysCode returns the TsysCode field if non-nil, zero value otherwise.

### GetTsysCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTsysCodeOk() (*string, bool)`

GetTsysCodeOk returns a tuple with the TsysCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsysCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTsysCode(v string)`

SetTsysCode sets TsysCode field to given value.

### HasTsysCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTsysCode() bool`

HasTsysCode returns a boolean if a field has been set.

### GetIsTsys

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetIsTsys() bool`

GetIsTsys returns the IsTsys field if non-nil, zero value otherwise.

### GetIsTsysOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetIsTsysOk() (*bool, bool)`

GetIsTsysOk returns a tuple with the IsTsys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTsys

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetIsTsys(v bool)`

SetIsTsys sets IsTsys field to given value.

### HasIsTsys

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasIsTsys() bool`

HasIsTsys returns a boolean if a field has been set.

### GetIsDisputable

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetIsDisputable() string`

GetIsDisputable returns the IsDisputable field if non-nil, zero value otherwise.

### GetIsDisputableOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetIsDisputableOk() (*string, bool)`

GetIsDisputableOk returns a tuple with the IsDisputable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisputable

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetIsDisputable(v string)`

SetIsDisputable sets IsDisputable field to given value.

### HasIsDisputable

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasIsDisputable() bool`

HasIsDisputable returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetOriginalAmount() map[string]interface{}`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetOriginalAmountOk() (*map[string]interface{}, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetOriginalAmount(v map[string]interface{})`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### SetOriginalAmountNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetOriginalAmountNil(b bool)`

 SetOriginalAmountNil sets the value for OriginalAmount to be an explicit nil

### UnsetOriginalAmount
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetOriginalAmount()`

UnsetOriginalAmount ensures that no value is present for OriginalAmount, not even an explicit nil
### GetRunningBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRunningBalance() string`

GetRunningBalance returns the RunningBalance field if non-nil, zero value otherwise.

### GetRunningBalanceOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRunningBalanceOk() (*string, bool)`

GetRunningBalanceOk returns a tuple with the RunningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRunningBalance(v string)`

SetRunningBalance sets RunningBalance field to given value.

### HasRunningBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasRunningBalance() bool`

HasRunningBalance returns a boolean if a field has been set.

### SetRunningBalanceNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRunningBalanceNil(b bool)`

 SetRunningBalanceNil sets the value for RunningBalance to be an explicit nil

### UnsetRunningBalance
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetRunningBalance()`

UnsetRunningBalance ensures that no value is present for RunningBalance, not even an explicit nil
### GetAssociatedCardNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetAssociatedCardNumber() string`

GetAssociatedCardNumber returns the AssociatedCardNumber field if non-nil, zero value otherwise.

### GetAssociatedCardNumberOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetAssociatedCardNumberOk() (*string, bool)`

GetAssociatedCardNumberOk returns a tuple with the AssociatedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCardNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetAssociatedCardNumber(v string)`

SetAssociatedCardNumber sets AssociatedCardNumber field to given value.

### HasAssociatedCardNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasAssociatedCardNumber() bool`

HasAssociatedCardNumber returns a boolean if a field has been set.

### GetPurchaseCountryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetPurchaseCountryCode() string`

GetPurchaseCountryCode returns the PurchaseCountryCode field if non-nil, zero value otherwise.

### GetPurchaseCountryCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetPurchaseCountryCodeOk() (*string, bool)`

GetPurchaseCountryCodeOk returns a tuple with the PurchaseCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCountryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetPurchaseCountryCode(v string)`

SetPurchaseCountryCode sets PurchaseCountryCode field to given value.

### HasPurchaseCountryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasPurchaseCountryCode() bool`

HasPurchaseCountryCode returns a boolean if a field has been set.

### SetPurchaseCountryCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetPurchaseCountryCodeNil(b bool)`

 SetPurchaseCountryCodeNil sets the value for PurchaseCountryCode to be an explicit nil

### UnsetPurchaseCountryCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetPurchaseCountryCode()`

UnsetPurchaseCountryCode ensures that no value is present for PurchaseCountryCode, not even an explicit nil
### GetOutOfCountryIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetOutOfCountryIndicator() string`

GetOutOfCountryIndicator returns the OutOfCountryIndicator field if non-nil, zero value otherwise.

### GetOutOfCountryIndicatorOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetOutOfCountryIndicatorOk() (*string, bool)`

GetOutOfCountryIndicatorOk returns a tuple with the OutOfCountryIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfCountryIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetOutOfCountryIndicator(v string)`

SetOutOfCountryIndicator sets OutOfCountryIndicator field to given value.

### HasOutOfCountryIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasOutOfCountryIndicator() bool`

HasOutOfCountryIndicator returns a boolean if a field has been set.

### SetOutOfCountryIndicatorNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetOutOfCountryIndicatorNil(b bool)`

 SetOutOfCountryIndicatorNil sets the value for OutOfCountryIndicator to be an explicit nil

### UnsetOutOfCountryIndicator
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetOutOfCountryIndicator()`

UnsetOutOfCountryIndicator ensures that no value is present for OutOfCountryIndicator, not even an explicit nil
### GetReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetReasonCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetStatus

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRecurringPaymentIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRecurringPaymentIndicator() string`

GetRecurringPaymentIndicator returns the RecurringPaymentIndicator field if non-nil, zero value otherwise.

### GetRecurringPaymentIndicatorOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRecurringPaymentIndicatorOk() (*string, bool)`

GetRecurringPaymentIndicatorOk returns a tuple with the RecurringPaymentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPaymentIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRecurringPaymentIndicator(v string)`

SetRecurringPaymentIndicator sets RecurringPaymentIndicator field to given value.

### HasRecurringPaymentIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasRecurringPaymentIndicator() bool`

HasRecurringPaymentIndicator returns a boolean if a field has been set.

### SetRecurringPaymentIndicatorNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRecurringPaymentIndicatorNil(b bool)`

 SetRecurringPaymentIndicatorNil sets the value for RecurringPaymentIndicator to be an explicit nil

### UnsetRecurringPaymentIndicator
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetRecurringPaymentIndicator()`

UnsetRecurringPaymentIndicator ensures that no value is present for RecurringPaymentIndicator, not even an explicit nil
### GetStatementIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetStatementIndicator() bool`

GetStatementIndicator returns the StatementIndicator field if non-nil, zero value otherwise.

### GetStatementIndicatorOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetStatementIndicatorOk() (*bool, bool)`

GetStatementIndicatorOk returns a tuple with the StatementIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetStatementIndicator(v bool)`

SetStatementIndicator sets StatementIndicator field to given value.

### HasStatementIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasStatementIndicator() bool`

HasStatementIndicator returns a boolean if a field has been set.

### GetFromAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.

### HasFromAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasFromAccount() bool`

HasFromAccount returns a boolean if a field has been set.

### SetFromAccountNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetFromAccountNil(b bool)`

 SetFromAccountNil sets the value for FromAccount to be an explicit nil

### UnsetFromAccount
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetFromAccount()`

UnsetFromAccount ensures that no value is present for FromAccount, not even an explicit nil
### GetToAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.

### HasToAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasToAccount() bool`

HasToAccount returns a boolean if a field has been set.

### SetToAccountNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetToAccountNil(b bool)`

 SetToAccountNil sets the value for ToAccount to be an explicit nil

### UnsetToAccount
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetToAccount()`

UnsetToAccount ensures that no value is present for ToAccount, not even an explicit nil
### GetPurchaseType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetPurchaseType() string`

GetPurchaseType returns the PurchaseType field if non-nil, zero value otherwise.

### GetPurchaseTypeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetPurchaseTypeOk() (*string, bool)`

GetPurchaseTypeOk returns a tuple with the PurchaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetPurchaseType(v string)`

SetPurchaseType sets PurchaseType field to given value.

### HasPurchaseType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasPurchaseType() bool`

HasPurchaseType returns a boolean if a field has been set.

### SetPurchaseTypeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetPurchaseTypeNil(b bool)`

 SetPurchaseTypeNil sets the value for PurchaseType to be an explicit nil

### UnsetPurchaseType
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetPurchaseType()`

UnsetPurchaseType ensures that no value is present for PurchaseType, not even an explicit nil
### GetRewardsCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRewardsCategory() string`

GetRewardsCategory returns the RewardsCategory field if non-nil, zero value otherwise.

### GetRewardsCategoryOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRewardsCategoryOk() (*string, bool)`

GetRewardsCategoryOk returns a tuple with the RewardsCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRewardsCategory(v string)`

SetRewardsCategory sets RewardsCategory field to given value.

### HasRewardsCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasRewardsCategory() bool`

HasRewardsCategory returns a boolean if a field has been set.

### GetRewardCard

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRewardCard() string`

GetRewardCard returns the RewardCard field if non-nil, zero value otherwise.

### GetRewardCardOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetRewardCardOk() (*string, bool)`

GetRewardCardOk returns a tuple with the RewardCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardCard

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRewardCard(v string)`

SetRewardCard sets RewardCard field to given value.

### HasRewardCard

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasRewardCard() bool`

HasRewardCard returns a boolean if a field has been set.

### SetRewardCardNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetRewardCardNil(b bool)`

 SetRewardCardNil sets the value for RewardCard to be an explicit nil

### UnsetRewardCard
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetRewardCard()`

UnsetRewardCard ensures that no value is present for RewardCard, not even an explicit nil
### GetCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetCategory() ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetCategoryOk() (*ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetCategory(v ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetUserInputTag

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetUserInputTag() string`

GetUserInputTag returns the UserInputTag field if non-nil, zero value otherwise.

### GetUserInputTagOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetUserInputTagOk() (*string, bool)`

GetUserInputTagOk returns a tuple with the UserInputTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputTag

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetUserInputTag(v string)`

SetUserInputTag sets UserInputTag field to given value.

### HasUserInputTag

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasUserInputTag() bool`

HasUserInputTag returns a boolean if a field has been set.

### SetUserInputTagNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetUserInputTagNil(b bool)`

 SetUserInputTagNil sets the value for UserInputTag to be an explicit nil

### UnsetUserInputTag
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetUserInputTag()`

UnsetUserInputTag ensures that no value is present for UserInputTag, not even an explicit nil
### GetCheque

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetCheque() string`

GetCheque returns the Cheque field if non-nil, zero value otherwise.

### GetChequeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetChequeOk() (*string, bool)`

GetChequeOk returns a tuple with the Cheque field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheque

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetCheque(v string)`

SetCheque sets Cheque field to given value.

### HasCheque

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasCheque() bool`

HasCheque returns a boolean if a field has been set.

### SetChequeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetChequeNil(b bool)`

 SetChequeNil sets the value for Cheque to be an explicit nil

### UnsetCheque
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) UnsetCheque()`

UnsetCheque ensures that no value is present for Cheque, not even an explicit nil
### GetMerchant

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetMerchant() ApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetMerchantOk() (*ApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetMerchant(v ApiCreditCreditIdTransactionsGet200ResponseDataSettledInnerMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetEnriched

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetEnriched() bool`

GetEnriched returns the Enriched field if non-nil, zero value otherwise.

### GetEnrichedOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetEnrichedOk() (*bool, bool)`

GetEnrichedOk returns a tuple with the Enriched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnriched

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetEnriched(v bool)`

SetEnriched sets Enriched field to given value.

### HasEnriched

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasEnriched() bool`

HasEnriched returns a boolean if a field has been set.

### GetAcquirerReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetAcquirerReferenceNumber() string`

GetAcquirerReferenceNumber returns the AcquirerReferenceNumber field if non-nil, zero value otherwise.

### GetAcquirerReferenceNumberOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetAcquirerReferenceNumberOk() (*string, bool)`

GetAcquirerReferenceNumberOk returns a tuple with the AcquirerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetAcquirerReferenceNumber(v string)`

SetAcquirerReferenceNumber sets AcquirerReferenceNumber field to given value.

### HasAcquirerReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasAcquirerReferenceNumber() bool`

HasAcquirerReferenceNumber returns a boolean if a field has been set.

### GetTransactionId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionKey() string`

GetTransactionKey returns the TransactionKey field if non-nil, zero value otherwise.

### GetTransactionKeyOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionKeyOk() (*string, bool)`

GetTransactionKeyOk returns a tuple with the TransactionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionKey(v string)`

SetTransactionKey sets TransactionKey field to given value.

### HasTransactionKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionKey() bool`

HasTransactionKey returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionCode() string`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionCodeOk() (*string, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionCode(v string)`

SetTransactionCode sets TransactionCode field to given value.

### HasTransactionCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionCode() bool`

HasTransactionCode returns a boolean if a field has been set.

### GetDescriptionLine1

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionLine1() string`

GetDescriptionLine1 returns the DescriptionLine1 field if non-nil, zero value otherwise.

### GetDescriptionLine1Ok

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionLine1Ok() (*string, bool)`

GetDescriptionLine1Ok returns a tuple with the DescriptionLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLine1

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetDescriptionLine1(v string)`

SetDescriptionLine1 sets DescriptionLine1 field to given value.

### HasDescriptionLine1

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasDescriptionLine1() bool`

HasDescriptionLine1 returns a boolean if a field has been set.

### GetDescriptionLine2

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionLine2() string`

GetDescriptionLine2 returns the DescriptionLine2 field if non-nil, zero value otherwise.

### GetDescriptionLine2Ok

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetDescriptionLine2Ok() (*string, bool)`

GetDescriptionLine2Ok returns a tuple with the DescriptionLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLine2

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetDescriptionLine2(v string)`

SetDescriptionLine2 sets DescriptionLine2 field to given value.

### HasDescriptionLine2

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasDescriptionLine2() bool`

HasDescriptionLine2 returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionAmount() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionAmountOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionAmount(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionCategory() string`

GetTransactionCategory returns the TransactionCategory field if non-nil, zero value otherwise.

### GetTransactionCategoryOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionCategoryOk() (*string, bool)`

GetTransactionCategoryOk returns a tuple with the TransactionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionCategory(v string)`

SetTransactionCategory sets TransactionCategory field to given value.

### HasTransactionCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionCategory() bool`

HasTransactionCategory returns a boolean if a field has been set.

### GetTransactionCategoryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionCategoryCode() string`

GetTransactionCategoryCode returns the TransactionCategoryCode field if non-nil, zero value otherwise.

### GetTransactionCategoryCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetTransactionCategoryCodeOk() (*string, bool)`

GetTransactionCategoryCodeOk returns a tuple with the TransactionCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategoryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetTransactionCategoryCode(v string)`

SetTransactionCategoryCode sets TransactionCategoryCode field to given value.

### HasTransactionCategoryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasTransactionCategoryCode() bool`

HasTransactionCategoryCode returns a boolean if a field has been set.

### GetSubdescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetSubdescription() string`

GetSubdescription returns the Subdescription field if non-nil, zero value otherwise.

### GetSubdescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) GetSubdescriptionOk() (*string, bool)`

GetSubdescriptionOk returns a tuple with the Subdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) SetSubdescription(v string)`

SetSubdescription sets Subdescription field to given value.

### HasSubdescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner) HasSubdescription() bool`

HasSubdescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


