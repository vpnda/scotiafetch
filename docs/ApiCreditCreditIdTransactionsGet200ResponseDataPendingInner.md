# ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionLines** | Pointer to **NullableString** |  | [optional] 
**SubDescription** | Pointer to **NullableString** |  | [optional] 
**CleanDescription** | Pointer to **NullableString** |  | [optional] 
**MnemonicCode** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**TsysCode** | Pointer to **NullableString** |  | [optional] 
**IsTsys** | Pointer to **bool** |  | [optional] 
**IsDisputable** | Pointer to **string** |  | [optional] 
**OriginalAmount** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**RunningBalance** | Pointer to **NullableString** |  | [optional] 
**AssociatedCardNumber** | Pointer to **string** |  | [optional] 
**PurchaseCountryCode** | Pointer to **NullableString** |  | [optional] 
**OutOfCountryIndicator** | Pointer to **NullableString** |  | [optional] 
**ReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RecurringPaymentIndicator** | Pointer to **string** |  | [optional] 
**StatementIndicator** | Pointer to **NullableString** |  | [optional] 
**FromAccount** | Pointer to **NullableString** |  | [optional] 
**ToAccount** | Pointer to **NullableString** |  | [optional] 
**PurchaseType** | Pointer to **NullableString** |  | [optional] 
**RewardsCategory** | Pointer to **NullableString** |  | [optional] 
**RewardCard** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**UserInputTag** | Pointer to **NullableString** |  | [optional] 
**Cheque** | Pointer to **NullableString** |  | [optional] 
**Merchant** | Pointer to [**ApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerMerchant**](ApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerMerchant.md) |  | [optional] 
**Enriched** | Pointer to **bool** |  | [optional] 
**AcquirerReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TransactionKey** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 
**TransactionCode** | Pointer to **NullableString** |  | [optional] 
**TransactionAmount** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**TransactionCategory** | Pointer to **NullableString** |  | [optional] 
**Subdescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiCreditCreditIdTransactionsGet200ResponseDataPendingInner

`func NewApiCreditCreditIdTransactionsGet200ResponseDataPendingInner() *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner`

NewApiCreditCreditIdTransactionsGet200ResponseDataPendingInner instantiates a new ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerWithDefaults

`func NewApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerWithDefaults() *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner`

NewApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerWithDefaults instantiates a new ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionLines

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetDescriptionLines() string`

GetDescriptionLines returns the DescriptionLines field if non-nil, zero value otherwise.

### GetDescriptionLinesOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetDescriptionLinesOk() (*string, bool)`

GetDescriptionLinesOk returns a tuple with the DescriptionLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLines

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetDescriptionLines(v string)`

SetDescriptionLines sets DescriptionLines field to given value.

### HasDescriptionLines

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasDescriptionLines() bool`

HasDescriptionLines returns a boolean if a field has been set.

### SetDescriptionLinesNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetDescriptionLinesNil(b bool)`

 SetDescriptionLinesNil sets the value for DescriptionLines to be an explicit nil

### UnsetDescriptionLines
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetDescriptionLines()`

UnsetDescriptionLines ensures that no value is present for DescriptionLines, not even an explicit nil
### GetSubDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetSubDescription() string`

GetSubDescription returns the SubDescription field if non-nil, zero value otherwise.

### GetSubDescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetSubDescriptionOk() (*string, bool)`

GetSubDescriptionOk returns a tuple with the SubDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetSubDescription(v string)`

SetSubDescription sets SubDescription field to given value.

### HasSubDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasSubDescription() bool`

HasSubDescription returns a boolean if a field has been set.

### SetSubDescriptionNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetSubDescriptionNil(b bool)`

 SetSubDescriptionNil sets the value for SubDescription to be an explicit nil

### UnsetSubDescription
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetSubDescription()`

UnsetSubDescription ensures that no value is present for SubDescription, not even an explicit nil
### GetCleanDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetCleanDescription() string`

GetCleanDescription returns the CleanDescription field if non-nil, zero value otherwise.

### GetCleanDescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetCleanDescriptionOk() (*string, bool)`

GetCleanDescriptionOk returns a tuple with the CleanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetCleanDescription(v string)`

SetCleanDescription sets CleanDescription field to given value.

### HasCleanDescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasCleanDescription() bool`

HasCleanDescription returns a boolean if a field has been set.

### SetCleanDescriptionNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetCleanDescriptionNil(b bool)`

 SetCleanDescriptionNil sets the value for CleanDescription to be an explicit nil

### UnsetCleanDescription
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetCleanDescription()`

UnsetCleanDescription ensures that no value is present for CleanDescription, not even an explicit nil
### GetMnemonicCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetMnemonicCode() string`

GetMnemonicCode returns the MnemonicCode field if non-nil, zero value otherwise.

### GetMnemonicCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetMnemonicCodeOk() (*string, bool)`

GetMnemonicCodeOk returns a tuple with the MnemonicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetMnemonicCode(v string)`

SetMnemonicCode sets MnemonicCode field to given value.

### HasMnemonicCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasMnemonicCode() bool`

HasMnemonicCode returns a boolean if a field has been set.

### SetMnemonicCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetMnemonicCodeNil(b bool)`

 SetMnemonicCodeNil sets the value for MnemonicCode to be an explicit nil

### UnsetMnemonicCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetMnemonicCode()`

UnsetMnemonicCode ensures that no value is present for MnemonicCode, not even an explicit nil
### GetType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTsysCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTsysCode() string`

GetTsysCode returns the TsysCode field if non-nil, zero value otherwise.

### GetTsysCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTsysCodeOk() (*string, bool)`

GetTsysCodeOk returns a tuple with the TsysCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsysCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTsysCode(v string)`

SetTsysCode sets TsysCode field to given value.

### HasTsysCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTsysCode() bool`

HasTsysCode returns a boolean if a field has been set.

### SetTsysCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTsysCodeNil(b bool)`

 SetTsysCodeNil sets the value for TsysCode to be an explicit nil

### UnsetTsysCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetTsysCode()`

UnsetTsysCode ensures that no value is present for TsysCode, not even an explicit nil
### GetIsTsys

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetIsTsys() bool`

GetIsTsys returns the IsTsys field if non-nil, zero value otherwise.

### GetIsTsysOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetIsTsysOk() (*bool, bool)`

GetIsTsysOk returns a tuple with the IsTsys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTsys

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetIsTsys(v bool)`

SetIsTsys sets IsTsys field to given value.

### HasIsTsys

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasIsTsys() bool`

HasIsTsys returns a boolean if a field has been set.

### GetIsDisputable

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetIsDisputable() string`

GetIsDisputable returns the IsDisputable field if non-nil, zero value otherwise.

### GetIsDisputableOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetIsDisputableOk() (*string, bool)`

GetIsDisputableOk returns a tuple with the IsDisputable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisputable

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetIsDisputable(v string)`

SetIsDisputable sets IsDisputable field to given value.

### HasIsDisputable

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasIsDisputable() bool`

HasIsDisputable returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetOriginalAmount() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetOriginalAmountOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetOriginalAmount(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetRunningBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRunningBalance() string`

GetRunningBalance returns the RunningBalance field if non-nil, zero value otherwise.

### GetRunningBalanceOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRunningBalanceOk() (*string, bool)`

GetRunningBalanceOk returns a tuple with the RunningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRunningBalance(v string)`

SetRunningBalance sets RunningBalance field to given value.

### HasRunningBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasRunningBalance() bool`

HasRunningBalance returns a boolean if a field has been set.

### SetRunningBalanceNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRunningBalanceNil(b bool)`

 SetRunningBalanceNil sets the value for RunningBalance to be an explicit nil

### UnsetRunningBalance
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetRunningBalance()`

UnsetRunningBalance ensures that no value is present for RunningBalance, not even an explicit nil
### GetAssociatedCardNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetAssociatedCardNumber() string`

GetAssociatedCardNumber returns the AssociatedCardNumber field if non-nil, zero value otherwise.

### GetAssociatedCardNumberOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetAssociatedCardNumberOk() (*string, bool)`

GetAssociatedCardNumberOk returns a tuple with the AssociatedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCardNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetAssociatedCardNumber(v string)`

SetAssociatedCardNumber sets AssociatedCardNumber field to given value.

### HasAssociatedCardNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasAssociatedCardNumber() bool`

HasAssociatedCardNumber returns a boolean if a field has been set.

### GetPurchaseCountryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetPurchaseCountryCode() string`

GetPurchaseCountryCode returns the PurchaseCountryCode field if non-nil, zero value otherwise.

### GetPurchaseCountryCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetPurchaseCountryCodeOk() (*string, bool)`

GetPurchaseCountryCodeOk returns a tuple with the PurchaseCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCountryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetPurchaseCountryCode(v string)`

SetPurchaseCountryCode sets PurchaseCountryCode field to given value.

### HasPurchaseCountryCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasPurchaseCountryCode() bool`

HasPurchaseCountryCode returns a boolean if a field has been set.

### SetPurchaseCountryCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetPurchaseCountryCodeNil(b bool)`

 SetPurchaseCountryCodeNil sets the value for PurchaseCountryCode to be an explicit nil

### UnsetPurchaseCountryCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetPurchaseCountryCode()`

UnsetPurchaseCountryCode ensures that no value is present for PurchaseCountryCode, not even an explicit nil
### GetOutOfCountryIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetOutOfCountryIndicator() string`

GetOutOfCountryIndicator returns the OutOfCountryIndicator field if non-nil, zero value otherwise.

### GetOutOfCountryIndicatorOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetOutOfCountryIndicatorOk() (*string, bool)`

GetOutOfCountryIndicatorOk returns a tuple with the OutOfCountryIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfCountryIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetOutOfCountryIndicator(v string)`

SetOutOfCountryIndicator sets OutOfCountryIndicator field to given value.

### HasOutOfCountryIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasOutOfCountryIndicator() bool`

HasOutOfCountryIndicator returns a boolean if a field has been set.

### SetOutOfCountryIndicatorNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetOutOfCountryIndicatorNil(b bool)`

 SetOutOfCountryIndicatorNil sets the value for OutOfCountryIndicator to be an explicit nil

### UnsetOutOfCountryIndicator
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetOutOfCountryIndicator()`

UnsetOutOfCountryIndicator ensures that no value is present for OutOfCountryIndicator, not even an explicit nil
### GetReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetReasonCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetStatus

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRecurringPaymentIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRecurringPaymentIndicator() string`

GetRecurringPaymentIndicator returns the RecurringPaymentIndicator field if non-nil, zero value otherwise.

### GetRecurringPaymentIndicatorOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRecurringPaymentIndicatorOk() (*string, bool)`

GetRecurringPaymentIndicatorOk returns a tuple with the RecurringPaymentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPaymentIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRecurringPaymentIndicator(v string)`

SetRecurringPaymentIndicator sets RecurringPaymentIndicator field to given value.

### HasRecurringPaymentIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasRecurringPaymentIndicator() bool`

HasRecurringPaymentIndicator returns a boolean if a field has been set.

### GetStatementIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetStatementIndicator() string`

GetStatementIndicator returns the StatementIndicator field if non-nil, zero value otherwise.

### GetStatementIndicatorOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetStatementIndicatorOk() (*string, bool)`

GetStatementIndicatorOk returns a tuple with the StatementIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetStatementIndicator(v string)`

SetStatementIndicator sets StatementIndicator field to given value.

### HasStatementIndicator

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasStatementIndicator() bool`

HasStatementIndicator returns a boolean if a field has been set.

### SetStatementIndicatorNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetStatementIndicatorNil(b bool)`

 SetStatementIndicatorNil sets the value for StatementIndicator to be an explicit nil

### UnsetStatementIndicator
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetStatementIndicator()`

UnsetStatementIndicator ensures that no value is present for StatementIndicator, not even an explicit nil
### GetFromAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.

### HasFromAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasFromAccount() bool`

HasFromAccount returns a boolean if a field has been set.

### SetFromAccountNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetFromAccountNil(b bool)`

 SetFromAccountNil sets the value for FromAccount to be an explicit nil

### UnsetFromAccount
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetFromAccount()`

UnsetFromAccount ensures that no value is present for FromAccount, not even an explicit nil
### GetToAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetToAccount() string`

GetToAccount returns the ToAccount field if non-nil, zero value otherwise.

### GetToAccountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetToAccountOk() (*string, bool)`

GetToAccountOk returns a tuple with the ToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetToAccount(v string)`

SetToAccount sets ToAccount field to given value.

### HasToAccount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasToAccount() bool`

HasToAccount returns a boolean if a field has been set.

### SetToAccountNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetToAccountNil(b bool)`

 SetToAccountNil sets the value for ToAccount to be an explicit nil

### UnsetToAccount
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetToAccount()`

UnsetToAccount ensures that no value is present for ToAccount, not even an explicit nil
### GetPurchaseType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetPurchaseType() string`

GetPurchaseType returns the PurchaseType field if non-nil, zero value otherwise.

### GetPurchaseTypeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetPurchaseTypeOk() (*string, bool)`

GetPurchaseTypeOk returns a tuple with the PurchaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetPurchaseType(v string)`

SetPurchaseType sets PurchaseType field to given value.

### HasPurchaseType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasPurchaseType() bool`

HasPurchaseType returns a boolean if a field has been set.

### SetPurchaseTypeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetPurchaseTypeNil(b bool)`

 SetPurchaseTypeNil sets the value for PurchaseType to be an explicit nil

### UnsetPurchaseType
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetPurchaseType()`

UnsetPurchaseType ensures that no value is present for PurchaseType, not even an explicit nil
### GetRewardsCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRewardsCategory() string`

GetRewardsCategory returns the RewardsCategory field if non-nil, zero value otherwise.

### GetRewardsCategoryOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRewardsCategoryOk() (*string, bool)`

GetRewardsCategoryOk returns a tuple with the RewardsCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRewardsCategory(v string)`

SetRewardsCategory sets RewardsCategory field to given value.

### HasRewardsCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasRewardsCategory() bool`

HasRewardsCategory returns a boolean if a field has been set.

### SetRewardsCategoryNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRewardsCategoryNil(b bool)`

 SetRewardsCategoryNil sets the value for RewardsCategory to be an explicit nil

### UnsetRewardsCategory
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetRewardsCategory()`

UnsetRewardsCategory ensures that no value is present for RewardsCategory, not even an explicit nil
### GetRewardCard

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRewardCard() string`

GetRewardCard returns the RewardCard field if non-nil, zero value otherwise.

### GetRewardCardOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetRewardCardOk() (*string, bool)`

GetRewardCardOk returns a tuple with the RewardCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardCard

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRewardCard(v string)`

SetRewardCard sets RewardCard field to given value.

### HasRewardCard

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasRewardCard() bool`

HasRewardCard returns a boolean if a field has been set.

### SetRewardCardNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetRewardCardNil(b bool)`

 SetRewardCardNil sets the value for RewardCard to be an explicit nil

### UnsetRewardCard
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetRewardCard()`

UnsetRewardCard ensures that no value is present for RewardCard, not even an explicit nil
### GetCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetUserInputTag

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetUserInputTag() string`

GetUserInputTag returns the UserInputTag field if non-nil, zero value otherwise.

### GetUserInputTagOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetUserInputTagOk() (*string, bool)`

GetUserInputTagOk returns a tuple with the UserInputTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputTag

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetUserInputTag(v string)`

SetUserInputTag sets UserInputTag field to given value.

### HasUserInputTag

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasUserInputTag() bool`

HasUserInputTag returns a boolean if a field has been set.

### SetUserInputTagNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetUserInputTagNil(b bool)`

 SetUserInputTagNil sets the value for UserInputTag to be an explicit nil

### UnsetUserInputTag
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetUserInputTag()`

UnsetUserInputTag ensures that no value is present for UserInputTag, not even an explicit nil
### GetCheque

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetCheque() string`

GetCheque returns the Cheque field if non-nil, zero value otherwise.

### GetChequeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetChequeOk() (*string, bool)`

GetChequeOk returns a tuple with the Cheque field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheque

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetCheque(v string)`

SetCheque sets Cheque field to given value.

### HasCheque

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasCheque() bool`

HasCheque returns a boolean if a field has been set.

### SetChequeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetChequeNil(b bool)`

 SetChequeNil sets the value for Cheque to be an explicit nil

### UnsetCheque
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetCheque()`

UnsetCheque ensures that no value is present for Cheque, not even an explicit nil
### GetMerchant

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetMerchant() ApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetMerchantOk() (*ApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetMerchant(v ApiCreditCreditIdTransactionsGet200ResponseDataPendingInnerMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetEnriched

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetEnriched() bool`

GetEnriched returns the Enriched field if non-nil, zero value otherwise.

### GetEnrichedOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetEnrichedOk() (*bool, bool)`

GetEnrichedOk returns a tuple with the Enriched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnriched

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetEnriched(v bool)`

SetEnriched sets Enriched field to given value.

### HasEnriched

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasEnriched() bool`

HasEnriched returns a boolean if a field has been set.

### GetAcquirerReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetAcquirerReferenceNumber() string`

GetAcquirerReferenceNumber returns the AcquirerReferenceNumber field if non-nil, zero value otherwise.

### GetAcquirerReferenceNumberOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetAcquirerReferenceNumberOk() (*string, bool)`

GetAcquirerReferenceNumberOk returns a tuple with the AcquirerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetAcquirerReferenceNumber(v string)`

SetAcquirerReferenceNumber sets AcquirerReferenceNumber field to given value.

### HasAcquirerReferenceNumber

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasAcquirerReferenceNumber() bool`

HasAcquirerReferenceNumber returns a boolean if a field has been set.

### SetAcquirerReferenceNumberNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetAcquirerReferenceNumberNil(b bool)`

 SetAcquirerReferenceNumberNil sets the value for AcquirerReferenceNumber to be an explicit nil

### UnsetAcquirerReferenceNumber
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetAcquirerReferenceNumber()`

UnsetAcquirerReferenceNumber ensures that no value is present for AcquirerReferenceNumber, not even an explicit nil
### GetTransactionId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionKey() string`

GetTransactionKey returns the TransactionKey field if non-nil, zero value otherwise.

### GetTransactionKeyOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionKeyOk() (*string, bool)`

GetTransactionKeyOk returns a tuple with the TransactionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionKey(v string)`

SetTransactionKey sets TransactionKey field to given value.

### HasTransactionKey

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionKey() bool`

HasTransactionKey returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionCode() string`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionCodeOk() (*string, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionCode(v string)`

SetTransactionCode sets TransactionCode field to given value.

### HasTransactionCode

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionCode() bool`

HasTransactionCode returns a boolean if a field has been set.

### SetTransactionCodeNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionCodeNil(b bool)`

 SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil

### UnsetTransactionCode
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetTransactionCode()`

UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
### GetTransactionAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionAmount() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionAmountOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionAmount(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionCategory() string`

GetTransactionCategory returns the TransactionCategory field if non-nil, zero value otherwise.

### GetTransactionCategoryOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetTransactionCategoryOk() (*string, bool)`

GetTransactionCategoryOk returns a tuple with the TransactionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionCategory(v string)`

SetTransactionCategory sets TransactionCategory field to given value.

### HasTransactionCategory

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasTransactionCategory() bool`

HasTransactionCategory returns a boolean if a field has been set.

### SetTransactionCategoryNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetTransactionCategoryNil(b bool)`

 SetTransactionCategoryNil sets the value for TransactionCategory to be an explicit nil

### UnsetTransactionCategory
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetTransactionCategory()`

UnsetTransactionCategory ensures that no value is present for TransactionCategory, not even an explicit nil
### GetSubdescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetSubdescription() string`

GetSubdescription returns the Subdescription field if non-nil, zero value otherwise.

### GetSubdescriptionOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) GetSubdescriptionOk() (*string, bool)`

GetSubdescriptionOk returns a tuple with the Subdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetSubdescription(v string)`

SetSubdescription sets Subdescription field to given value.

### HasSubdescription

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) HasSubdescription() bool`

HasSubdescription returns a boolean if a field has been set.

### SetSubdescriptionNil

`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) SetSubdescriptionNil(b bool)`

 SetSubdescriptionNil sets the value for Subdescription to be an explicit nil

### UnsetSubdescription
`func (o *ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner) UnsetSubdescription()`

UnsetSubdescription ensures that no value is present for Subdescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


