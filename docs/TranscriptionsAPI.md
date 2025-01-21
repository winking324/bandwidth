# \TranscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRealTimeTranscription**](TranscriptionsAPI.md#DeleteRealTimeTranscription) | **Delete** /accounts/{accountId}/calls/{callId}/transcriptions/{transcriptionId} | Delete a specific transcription
[**GetRealTimeTranscription**](TranscriptionsAPI.md#GetRealTimeTranscription) | **Get** /accounts/{accountId}/calls/{callId}/transcriptions/{transcriptionId} | Retrieve a specific transcription
[**ListRealTimeTranscriptions**](TranscriptionsAPI.md#ListRealTimeTranscriptions) | **Get** /accounts/{accountId}/calls/{callId}/transcriptions | Enumerate transcriptions made with StartTranscription



## DeleteRealTimeTranscription

> DeleteRealTimeTranscription(ctx, accountId, callId, transcriptionId).Execute()

Delete a specific transcription



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
	transcriptionId := "t-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Transcription ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TranscriptionsAPI.DeleteRealTimeTranscription(context.Background(), accountId, callId, transcriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranscriptionsAPI.DeleteRealTimeTranscription``: %v\n", err)
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
**transcriptionId** | **string** | Programmable Voice API Transcription ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRealTimeTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealTimeTranscription

> CallTranscriptionResponse GetRealTimeTranscription(ctx, accountId, callId, transcriptionId).Execute()

Retrieve a specific transcription



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
	transcriptionId := "t-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Transcription ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranscriptionsAPI.GetRealTimeTranscription(context.Background(), accountId, callId, transcriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranscriptionsAPI.GetRealTimeTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealTimeTranscription`: CallTranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `TranscriptionsAPI.GetRealTimeTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 
**transcriptionId** | **string** | Programmable Voice API Transcription ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealTimeTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CallTranscriptionResponse**](CallTranscriptionResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealTimeTranscriptions

> []CallTranscriptionMetadata ListRealTimeTranscriptions(ctx, accountId, callId).Execute()

Enumerate transcriptions made with StartTranscription



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
	resp, r, err := apiClient.TranscriptionsAPI.ListRealTimeTranscriptions(context.Background(), accountId, callId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranscriptionsAPI.ListRealTimeTranscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRealTimeTranscriptions`: []CallTranscriptionMetadata
	fmt.Fprintf(os.Stdout, "Response from `TranscriptionsAPI.ListRealTimeTranscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRealTimeTranscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CallTranscriptionMetadata**](CallTranscriptionMetadata.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

