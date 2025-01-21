# \CallsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCall**](CallsAPI.md#CreateCall) | **Post** /accounts/{accountId}/calls | Create Call
[**GetCallState**](CallsAPI.md#GetCallState) | **Get** /accounts/{accountId}/calls/{callId} | Get Call State Information
[**ListCalls**](CallsAPI.md#ListCalls) | **Get** /accounts/{accountId}/calls | Get Calls
[**UpdateCall**](CallsAPI.md#UpdateCall) | **Post** /accounts/{accountId}/calls/{callId} | Update Call
[**UpdateCallBxml**](CallsAPI.md#UpdateCallBxml) | **Put** /accounts/{accountId}/calls/{callId}/bxml | Update Call BXML



## CreateCall

> CreateCallResponse CreateCall(ctx, accountId).CreateCall(createCall).Execute()

Create Call



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
	createCall := *openapiclient.NewCreateCall("+19195551234", "+15555551212", "1234-qwer-5679-tyui", "https://www.myCallbackServer.example/webhooks/answer") // CreateCall | JSON object containing information to create an outbound call

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallsAPI.CreateCall(context.Background(), accountId).CreateCall(createCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.CreateCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCall`: CreateCallResponse
	fmt.Fprintf(os.Stdout, "Response from `CallsAPI.CreateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCall** | [**CreateCall**](CreateCall.md) | JSON object containing information to create an outbound call | 

### Return type

[**CreateCallResponse**](CreateCallResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallState

> CallState GetCallState(ctx, accountId, callId).Execute()

Get Call State Information



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
	callId := "c-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Call ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallsAPI.GetCallState(context.Background(), accountId, callId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.GetCallState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallState`: CallState
	fmt.Fprintf(os.Stdout, "Response from `CallsAPI.GetCallState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CallState**](CallState.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCalls

> []CallState ListCalls(ctx, accountId).To(to).From(from).MinStartTime(minStartTime).MaxStartTime(maxStartTime).DisconnectCause(disconnectCause).PageSize(pageSize).PageToken(pageToken).Execute()

Get Calls



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
	to := "%2b19195551234" // string | Filter results by the `to` field. (optional)
	from := "%2b19195554321" // string | Filter results by the `from` field. (optional)
	minStartTime := "2022-06-21T19:13:21Z" // string | Filter results to calls which have a `startTime` after or including `minStartTime` (in ISO8601 format). (optional)
	maxStartTime := "2022-06-21T19:13:21Z" // string | Filter results to calls which have a `startTime` before or including `maxStartTime` (in ISO8601 format). (optional)
	disconnectCause := "hangup" // string | Filter results to calls with specified call Disconnect Cause. (optional)
	pageSize := int32(500) // int32 | Specifies the max number of calls that will be returned. (optional) (default to 1000)
	pageToken := "pageToken_example" // string | Not intended for explicit use. To use pagination, follow the links in the `Link` header of the response, as indicated in the endpoint description. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallsAPI.ListCalls(context.Background(), accountId).To(to).From(from).MinStartTime(minStartTime).MaxStartTime(maxStartTime).DisconnectCause(disconnectCause).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.ListCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCalls`: []CallState
	fmt.Fprintf(os.Stdout, "Response from `CallsAPI.ListCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | Filter results by the &#x60;to&#x60; field. | 
 **from** | **string** | Filter results by the &#x60;from&#x60; field. | 
 **minStartTime** | **string** | Filter results to calls which have a &#x60;startTime&#x60; after or including &#x60;minStartTime&#x60; (in ISO8601 format). | 
 **maxStartTime** | **string** | Filter results to calls which have a &#x60;startTime&#x60; before or including &#x60;maxStartTime&#x60; (in ISO8601 format). | 
 **disconnectCause** | **string** | Filter results to calls with specified call Disconnect Cause. | 
 **pageSize** | **int32** | Specifies the max number of calls that will be returned. | [default to 1000]
 **pageToken** | **string** | Not intended for explicit use. To use pagination, follow the links in the &#x60;Link&#x60; header of the response, as indicated in the endpoint description. | 

### Return type

[**[]CallState**](CallState.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCall

> UpdateCall(ctx, accountId, callId).UpdateCall(updateCall).Execute()

Update Call



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
	callId := "c-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Call ID.
	updateCall := *openapiclient.NewUpdateCall() // UpdateCall | JSON object containing information to redirect an existing call to a new BXML document

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CallsAPI.UpdateCall(context.Background(), accountId, callId).UpdateCall(updateCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.UpdateCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCall** | [**UpdateCall**](UpdateCall.md) | JSON object containing information to redirect an existing call to a new BXML document | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallBxml

> UpdateCallBxml(ctx, accountId, callId).Body(body).Execute()

Update Call BXML



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
	callId := "c-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Call ID.
	body := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<Bxml>
  <SpeakSentence>This is a test sentence.</SpeakSentence>
</Bxml>" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CallsAPI.UpdateCallBxml(context.Background(), accountId, callId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallsAPI.UpdateCallBxml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallBxmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

