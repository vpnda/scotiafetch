# \DefaultAPI

All URIs are relative to *https://secure.scotiabank.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsSummaryGet**](DefaultAPI.md#ApiAccountsSummaryGet) | **Get** /api/accounts/summary | GET summary
[**ApiCampaignsGet**](DefaultAPI.md#ApiCampaignsGet) | **Get** /api/campaigns | GET campaigns
[**ApiCardsCardIdGet**](DefaultAPI.md#ApiCardsCardIdGet) | **Get** /api/cards/{cardId} | GET cards by cardId
[**ApiCreditCreditIdGet**](DefaultAPI.md#ApiCreditCreditIdGet) | **Get** /api/credit/{creditId} | GET credit by creditId
[**ApiCreditCreditIdTransactionsGet**](DefaultAPI.md#ApiCreditCreditIdTransactionsGet) | **Get** /api/credit/{creditId}/transactions | GET transactions by creditId
[**ApiMpsaAccountsAccountIdGet**](DefaultAPI.md#ApiMpsaAccountsAccountIdGet) | **Get** /api/mpsa-accounts/{accountId} | GET mpsa-accounts by accountId
[**ApiMpsaAccountsAccountIdTransactionsGet**](DefaultAPI.md#ApiMpsaAccountsAccountIdTransactionsGet) | **Get** /api/mpsa-accounts/{accountId}/transactions | GET transactions by accountId
[**ApiRewardsRewardIdGet**](DefaultAPI.md#ApiRewardsRewardIdGet) | **Get** /api/rewards/{rewardId} | GET rewards by rewardId
[**ApiTransactionsDepositAccountsDepositAccountIdGet**](DefaultAPI.md#ApiTransactionsDepositAccountsDepositAccountIdGet) | **Get** /api/transactions/deposit-accounts/{depositAccountId} | GET deposit-accounts by depositAccountId



## ApiAccountsSummaryGet

> ApiAccountsSummaryGet200Response ApiAccountsSummaryGet(ctx).Execute()

GET summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiAccountsSummaryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiAccountsSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAccountsSummaryGet`: ApiAccountsSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiAccountsSummaryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsSummaryGetRequest struct via the builder pattern


### Return type

[**ApiAccountsSummaryGet200Response**](ApiAccountsSummaryGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCampaignsGet

> ApiCampaignsGet200Response ApiCampaignsGet(ctx).Page(page).Execute()

GET campaigns

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := "page_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiCampaignsGet(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiCampaignsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCampaignsGet`: ApiCampaignsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiCampaignsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCampaignsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** |  | 

### Return type

[**ApiCampaignsGet200Response**](ApiCampaignsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCardsCardIdGet

> ApiCardsCardIdGet200Response ApiCardsCardIdGet(ctx, cardId).Execute()

GET cards by cardId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cardId := "cardId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiCardsCardIdGet(context.Background(), cardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiCardsCardIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCardsCardIdGet`: ApiCardsCardIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiCardsCardIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCardsCardIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCardsCardIdGet200Response**](ApiCardsCardIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCreditCreditIdGet

> ApiCreditCreditIdGet200Response ApiCreditCreditIdGet(ctx, creditId).Execute()

GET credit by creditId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	creditId := "creditId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiCreditCreditIdGet(context.Background(), creditId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiCreditCreditIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCreditCreditIdGet`: ApiCreditCreditIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiCreditCreditIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCreditCreditIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCreditCreditIdGet200Response**](ApiCreditCreditIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCreditCreditIdTransactionsGet

> ApiCreditCreditIdTransactionsGet200Response ApiCreditCreditIdTransactionsGet(ctx, creditId).Execute()

GET transactions by creditId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	creditId := "creditId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiCreditCreditIdTransactionsGet(context.Background(), creditId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiCreditCreditIdTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCreditCreditIdTransactionsGet`: ApiCreditCreditIdTransactionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiCreditCreditIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCreditCreditIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCreditCreditIdTransactionsGet200Response**](ApiCreditCreditIdTransactionsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMpsaAccountsAccountIdGet

> ApiMpsaAccountsAccountIdGet200Response ApiMpsaAccountsAccountIdGet(ctx, accountId).Execute()

GET mpsa-accounts by accountId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiMpsaAccountsAccountIdGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiMpsaAccountsAccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiMpsaAccountsAccountIdGet`: ApiMpsaAccountsAccountIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiMpsaAccountsAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMpsaAccountsAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiMpsaAccountsAccountIdGet200Response**](ApiMpsaAccountsAccountIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMpsaAccountsAccountIdTransactionsGet

> ApiMpsaAccountsAccountIdTransactionsGet200Response ApiMpsaAccountsAccountIdTransactionsGet(ctx, accountId).FromDate(fromDate).ToDate(toDate).Execute()

GET transactions by accountId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountId := "accountId_example" // string | 
	fromDate := "fromDate_example" // string |  (optional)
	toDate := "toDate_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiMpsaAccountsAccountIdTransactionsGet(context.Background(), accountId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiMpsaAccountsAccountIdTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiMpsaAccountsAccountIdTransactionsGet`: ApiMpsaAccountsAccountIdTransactionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiMpsaAccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMpsaAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**ApiMpsaAccountsAccountIdTransactionsGet200Response**](ApiMpsaAccountsAccountIdTransactionsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRewardsRewardIdGet

> ApiRewardsRewardIdGet200Response ApiRewardsRewardIdGet(ctx, rewardId).Execute()

GET rewards by rewardId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	rewardId := "rewardId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiRewardsRewardIdGet(context.Background(), rewardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiRewardsRewardIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRewardsRewardIdGet`: ApiRewardsRewardIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiRewardsRewardIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rewardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRewardsRewardIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiRewardsRewardIdGet200Response**](ApiRewardsRewardIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTransactionsDepositAccountsDepositAccountIdGet

> ApiTransactionsDepositAccountsDepositAccountIdGet200Response ApiTransactionsDepositAccountsDepositAccountIdGet(ctx, depositAccountId).Execute()

GET deposit-accounts by depositAccountId

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	depositAccountId := "depositAccountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiTransactionsDepositAccountsDepositAccountIdGet(context.Background(), depositAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiTransactionsDepositAccountsDepositAccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTransactionsDepositAccountsDepositAccountIdGet`: ApiTransactionsDepositAccountsDepositAccountIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiTransactionsDepositAccountsDepositAccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**depositAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTransactionsDepositAccountsDepositAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiTransactionsDepositAccountsDepositAccountIdGet200Response**](ApiTransactionsDepositAccountsDepositAccountIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

