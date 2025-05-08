# ApiCreditCreditIdTransactionsGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to [**[]ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner**](ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner.md) |  | [optional] 
**Settled** | Pointer to [**[]ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner**](ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner.md) |  | [optional] 
**PendingTransactionBalance** | Pointer to [**ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner**](ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner.md) |  | [optional] 

## Methods

### NewApiCreditCreditIdTransactionsGet200ResponseData

`func NewApiCreditCreditIdTransactionsGet200ResponseData() *ApiCreditCreditIdTransactionsGet200ResponseData`

NewApiCreditCreditIdTransactionsGet200ResponseData instantiates a new ApiCreditCreditIdTransactionsGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCreditCreditIdTransactionsGet200ResponseDataWithDefaults

`func NewApiCreditCreditIdTransactionsGet200ResponseDataWithDefaults() *ApiCreditCreditIdTransactionsGet200ResponseData`

NewApiCreditCreditIdTransactionsGet200ResponseDataWithDefaults instantiates a new ApiCreditCreditIdTransactionsGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) GetPending() []ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) GetPendingOk() (*[]ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) SetPending(v []ApiCreditCreditIdTransactionsGet200ResponseDataPendingInner)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetSettled

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) GetSettled() []ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) GetSettledOk() (*[]ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) SetSettled(v []ApiCreditCreditIdTransactionsGet200ResponseDataSettledInner)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetPendingTransactionBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) GetPendingTransactionBalance() ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner`

GetPendingTransactionBalance returns the PendingTransactionBalance field if non-nil, zero value otherwise.

### GetPendingTransactionBalanceOk

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) GetPendingTransactionBalanceOk() (*ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner, bool)`

GetPendingTransactionBalanceOk returns a tuple with the PendingTransactionBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTransactionBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) SetPendingTransactionBalance(v ApiAccountsSummaryGet200ResponseDataProductsInnerPrimaryBalancesInner)`

SetPendingTransactionBalance sets PendingTransactionBalance field to given value.

### HasPendingTransactionBalance

`func (o *ApiCreditCreditIdTransactionsGet200ResponseData) HasPendingTransactionBalance() bool`

HasPendingTransactionBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


