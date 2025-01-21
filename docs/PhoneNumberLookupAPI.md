# \PhoneNumberLookupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLookup**](PhoneNumberLookupAPI.md#CreateLookup) | **Post** /accounts/{accountId}/tnlookup | Create Lookup
[**GetLookupStatus**](PhoneNumberLookupAPI.md#GetLookupStatus) | **Get** /accounts/{accountId}/tnlookup/{requestId} | Get Lookup Request Status



## CreateLookup

> CreateLookupResponse CreateLookup(ctx, accountId).LookupRequest(lookupRequest).Execute()

Create Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/winking324/bandwidth"
)

func main() {
	accountId := "9900000" // string | Your Bandwidth Account ID.
	lookupRequest := *openapiclient.NewLookupRequest([]string{"Tns_example"}) // LookupRequest | Phone number lookup request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberLookupAPI.CreateLookup(context.Background(), accountId).LookupRequest(lookupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberLookupAPI.CreateLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLookup`: CreateLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberLookupAPI.CreateLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lookupRequest** | [**LookupRequest**](LookupRequest.md) | Phone number lookup request. | 

### Return type

[**CreateLookupResponse**](CreateLookupResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLookupStatus

> LookupStatus GetLookupStatus(ctx, accountId, requestId).Execute()

Get Lookup Request Status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/winking324/bandwidth"
)

func main() {
	accountId := "9900000" // string | Your Bandwidth Account ID.
	requestId := "004223a0-8b17-41b1-bf81-20732adf5590" // string | The phone number lookup request ID from Bandwidth.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberLookupAPI.GetLookupStatus(context.Background(), accountId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberLookupAPI.GetLookupStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLookupStatus`: LookupStatus
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberLookupAPI.GetLookupStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**requestId** | **string** | The phone number lookup request ID from Bandwidth. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLookupStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LookupStatus**](LookupStatus.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

