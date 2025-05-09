# ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionKey** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TransactionRecipient** | Pointer to **NullableString** |  | [optional] 
**TransactionTypeCode** | Pointer to **NullableString** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**TransactionAmount** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**RunningBalance** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 
**ChequeNumber** | Pointer to **NullableString** |  | [optional] 
**ChequeImage** | Pointer to **NullableString** |  | [optional] 
**ChequeImageStatus** | Pointer to **NullableString** |  | [optional] 
**ChequeImageKey** | Pointer to **NullableString** |  | [optional] 
**CardNumber** | Pointer to **string** |  | [optional] 
**MerchantName** | Pointer to **NullableString** |  | [optional] 
**TransactionCategory** | Pointer to **NullableString** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner

`func NewApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner() *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner`

NewApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner instantiates a new ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInnerWithDefaults

`func NewApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInnerWithDefaults() *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner`

NewApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInnerWithDefaults instantiates a new ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionKey

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionKey() string`

GetTransactionKey returns the TransactionKey field if non-nil, zero value otherwise.

### GetTransactionKeyOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionKeyOk() (*string, bool)`

GetTransactionKeyOk returns a tuple with the TransactionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionKey

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionKey(v string)`

SetTransactionKey sets TransactionKey field to given value.

### HasTransactionKey

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionKey() bool`

HasTransactionKey returns a boolean if a field has been set.

### GetTransactionId

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetDescription

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTransactionRecipient

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionRecipient() string`

GetTransactionRecipient returns the TransactionRecipient field if non-nil, zero value otherwise.

### GetTransactionRecipientOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionRecipientOk() (*string, bool)`

GetTransactionRecipientOk returns a tuple with the TransactionRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRecipient

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionRecipient(v string)`

SetTransactionRecipient sets TransactionRecipient field to given value.

### HasTransactionRecipient

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionRecipient() bool`

HasTransactionRecipient returns a boolean if a field has been set.

### SetTransactionRecipientNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionRecipientNil(b bool)`

 SetTransactionRecipientNil sets the value for TransactionRecipient to be an explicit nil

### UnsetTransactionRecipient
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetTransactionRecipient()`

UnsetTransactionRecipient ensures that no value is present for TransactionRecipient, not even an explicit nil
### GetTransactionTypeCode

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionTypeCode() string`

GetTransactionTypeCode returns the TransactionTypeCode field if non-nil, zero value otherwise.

### GetTransactionTypeCodeOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionTypeCodeOk() (*string, bool)`

GetTransactionTypeCodeOk returns a tuple with the TransactionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTypeCode

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionTypeCode(v string)`

SetTransactionTypeCode sets TransactionTypeCode field to given value.

### HasTransactionTypeCode

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionTypeCode() bool`

HasTransactionTypeCode returns a boolean if a field has been set.

### SetTransactionTypeCodeNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionTypeCodeNil(b bool)`

 SetTransactionTypeCodeNil sets the value for TransactionTypeCode to be an explicit nil

### UnsetTransactionTypeCode
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetTransactionTypeCode()`

UnsetTransactionTypeCode ensures that no value is present for TransactionTypeCode, not even an explicit nil
### GetTransactionType

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionAmount() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionAmountOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionAmount(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetRunningBalance

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetRunningBalance() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetRunningBalance returns the RunningBalance field if non-nil, zero value otherwise.

### GetRunningBalanceOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetRunningBalanceOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetRunningBalanceOk returns a tuple with the RunningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningBalance

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetRunningBalance(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetRunningBalance sets RunningBalance field to given value.

### HasRunningBalance

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasRunningBalance() bool`

HasRunningBalance returns a boolean if a field has been set.

### GetChequeNumber

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeNumber() string`

GetChequeNumber returns the ChequeNumber field if non-nil, zero value otherwise.

### GetChequeNumberOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeNumberOk() (*string, bool)`

GetChequeNumberOk returns a tuple with the ChequeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNumber

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeNumber(v string)`

SetChequeNumber sets ChequeNumber field to given value.

### HasChequeNumber

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasChequeNumber() bool`

HasChequeNumber returns a boolean if a field has been set.

### SetChequeNumberNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeNumberNil(b bool)`

 SetChequeNumberNil sets the value for ChequeNumber to be an explicit nil

### UnsetChequeNumber
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetChequeNumber()`

UnsetChequeNumber ensures that no value is present for ChequeNumber, not even an explicit nil
### GetChequeImage

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeImage() string`

GetChequeImage returns the ChequeImage field if non-nil, zero value otherwise.

### GetChequeImageOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeImageOk() (*string, bool)`

GetChequeImageOk returns a tuple with the ChequeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeImage

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeImage(v string)`

SetChequeImage sets ChequeImage field to given value.

### HasChequeImage

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasChequeImage() bool`

HasChequeImage returns a boolean if a field has been set.

### SetChequeImageNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeImageNil(b bool)`

 SetChequeImageNil sets the value for ChequeImage to be an explicit nil

### UnsetChequeImage
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetChequeImage()`

UnsetChequeImage ensures that no value is present for ChequeImage, not even an explicit nil
### GetChequeImageStatus

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeImageStatus() string`

GetChequeImageStatus returns the ChequeImageStatus field if non-nil, zero value otherwise.

### GetChequeImageStatusOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeImageStatusOk() (*string, bool)`

GetChequeImageStatusOk returns a tuple with the ChequeImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeImageStatus

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeImageStatus(v string)`

SetChequeImageStatus sets ChequeImageStatus field to given value.

### HasChequeImageStatus

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasChequeImageStatus() bool`

HasChequeImageStatus returns a boolean if a field has been set.

### SetChequeImageStatusNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeImageStatusNil(b bool)`

 SetChequeImageStatusNil sets the value for ChequeImageStatus to be an explicit nil

### UnsetChequeImageStatus
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetChequeImageStatus()`

UnsetChequeImageStatus ensures that no value is present for ChequeImageStatus, not even an explicit nil
### GetChequeImageKey

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeImageKey() string`

GetChequeImageKey returns the ChequeImageKey field if non-nil, zero value otherwise.

### GetChequeImageKeyOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetChequeImageKeyOk() (*string, bool)`

GetChequeImageKeyOk returns a tuple with the ChequeImageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeImageKey

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeImageKey(v string)`

SetChequeImageKey sets ChequeImageKey field to given value.

### HasChequeImageKey

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasChequeImageKey() bool`

HasChequeImageKey returns a boolean if a field has been set.

### SetChequeImageKeyNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetChequeImageKeyNil(b bool)`

 SetChequeImageKeyNil sets the value for ChequeImageKey to be an explicit nil

### UnsetChequeImageKey
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetChequeImageKey()`

UnsetChequeImageKey ensures that no value is present for ChequeImageKey, not even an explicit nil
### GetCardNumber

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetMerchantName

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetTransactionCategory

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionCategory() string`

GetTransactionCategory returns the TransactionCategory field if non-nil, zero value otherwise.

### GetTransactionCategoryOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionCategoryOk() (*string, bool)`

GetTransactionCategoryOk returns a tuple with the TransactionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCategory

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionCategory(v string)`

SetTransactionCategory sets TransactionCategory field to given value.

### HasTransactionCategory

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionCategory() bool`

HasTransactionCategory returns a boolean if a field has been set.

### SetTransactionCategoryNil

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionCategoryNil(b bool)`

 SetTransactionCategoryNil sets the value for TransactionCategory to be an explicit nil

### UnsetTransactionCategory
`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) UnsetTransactionCategory()`

UnsetTransactionCategory ensures that no value is present for TransactionCategory, not even an explicit nil
### GetTransactionDate

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


