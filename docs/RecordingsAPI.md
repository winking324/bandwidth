# \RecordingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecording**](RecordingsAPI.md#DeleteRecording) | **Delete** /accounts/{accountId}/calls/{callId}/recordings/{recordingId} | Delete Recording
[**DeleteRecordingMedia**](RecordingsAPI.md#DeleteRecordingMedia) | **Delete** /accounts/{accountId}/calls/{callId}/recordings/{recordingId}/media | Delete Recording Media
[**DeleteRecordingTranscription**](RecordingsAPI.md#DeleteRecordingTranscription) | **Delete** /accounts/{accountId}/calls/{callId}/recordings/{recordingId}/transcription | Delete Transcription
[**DownloadCallRecording**](RecordingsAPI.md#DownloadCallRecording) | **Get** /accounts/{accountId}/calls/{callId}/recordings/{recordingId}/media | Download Recording
[**GetCallRecording**](RecordingsAPI.md#GetCallRecording) | **Get** /accounts/{accountId}/calls/{callId}/recordings/{recordingId} | Get Call Recording
[**GetRecordingTranscription**](RecordingsAPI.md#GetRecordingTranscription) | **Get** /accounts/{accountId}/calls/{callId}/recordings/{recordingId}/transcription | Get Transcription
[**ListAccountCallRecordings**](RecordingsAPI.md#ListAccountCallRecordings) | **Get** /accounts/{accountId}/recordings | Get Call Recordings
[**ListCallRecordings**](RecordingsAPI.md#ListCallRecordings) | **Get** /accounts/{accountId}/calls/{callId}/recordings | List Call Recordings
[**TranscribeCallRecording**](RecordingsAPI.md#TranscribeCallRecording) | **Post** /accounts/{accountId}/calls/{callId}/recordings/{recordingId}/transcription | Create Transcription Request
[**UpdateCallRecordingState**](RecordingsAPI.md#UpdateCallRecordingState) | **Put** /accounts/{accountId}/calls/{callId}/recording | Update Recording



## DeleteRecording

> DeleteRecording(ctx, accountId, callId, recordingId).Execute()

Delete Recording



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordingsAPI.DeleteRecording(context.Background(), accountId, callId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.DeleteRecording``: %v\n", err)
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
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingRequest struct via the builder pattern


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


## DeleteRecordingMedia

> DeleteRecordingMedia(ctx, accountId, callId, recordingId).Execute()

Delete Recording Media



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordingsAPI.DeleteRecordingMedia(context.Background(), accountId, callId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.DeleteRecordingMedia``: %v\n", err)
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
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingMediaRequest struct via the builder pattern


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


## DeleteRecordingTranscription

> DeleteRecordingTranscription(ctx, accountId, callId, recordingId).Execute()

Delete Transcription



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordingsAPI.DeleteRecordingTranscription(context.Background(), accountId, callId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.DeleteRecordingTranscription``: %v\n", err)
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
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingTranscriptionRequest struct via the builder pattern


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


## DownloadCallRecording

> *os.File DownloadCallRecording(ctx, accountId, callId, recordingId).Execute()

Download Recording



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingsAPI.DownloadCallRecording(context.Background(), accountId, callId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.DownloadCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadCallRecording`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RecordingsAPI.DownloadCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/vnd.wave, audio/mpeg, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallRecording

> CallRecordingMetadata GetCallRecording(ctx, accountId, callId, recordingId).Execute()

Get Call Recording



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingsAPI.GetCallRecording(context.Background(), accountId, callId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.GetCallRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallRecording`: CallRecordingMetadata
	fmt.Fprintf(os.Stdout, "Response from `RecordingsAPI.GetCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CallRecordingMetadata**](CallRecordingMetadata.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingTranscription

> RecordingTranscriptions GetRecordingTranscription(ctx, accountId, callId, recordingId).Execute()

Get Transcription



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingsAPI.GetRecordingTranscription(context.Background(), accountId, callId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.GetRecordingTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordingTranscription`: RecordingTranscriptions
	fmt.Fprintf(os.Stdout, "Response from `RecordingsAPI.GetRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RecordingTranscriptions**](RecordingTranscriptions.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountCallRecordings

> []CallRecordingMetadata ListAccountCallRecordings(ctx, accountId).To(to).From(from).MinStartTime(minStartTime).MaxStartTime(maxStartTime).Execute()

Get Call Recordings



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
	minStartTime := "2022-06-21T19:13:21Z" // string | Filter results to recordings which have a `startTime` after or including `minStartTime` (in ISO8601 format). (optional)
	maxStartTime := "2022-06-21T19:13:21Z" // string | Filter results to recordings which have a `startTime` before `maxStartTime` (in ISO8601 format). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingsAPI.ListAccountCallRecordings(context.Background(), accountId).To(to).From(from).MinStartTime(minStartTime).MaxStartTime(maxStartTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.ListAccountCallRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountCallRecordings`: []CallRecordingMetadata
	fmt.Fprintf(os.Stdout, "Response from `RecordingsAPI.ListAccountCallRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountCallRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | Filter results by the &#x60;to&#x60; field. | 
 **from** | **string** | Filter results by the &#x60;from&#x60; field. | 
 **minStartTime** | **string** | Filter results to recordings which have a &#x60;startTime&#x60; after or including &#x60;minStartTime&#x60; (in ISO8601 format). | 
 **maxStartTime** | **string** | Filter results to recordings which have a &#x60;startTime&#x60; before &#x60;maxStartTime&#x60; (in ISO8601 format). | 

### Return type

[**[]CallRecordingMetadata**](CallRecordingMetadata.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallRecordings

> []CallRecordingMetadata ListCallRecordings(ctx, accountId, callId).Execute()

List Call Recordings



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
	resp, r, err := apiClient.RecordingsAPI.ListCallRecordings(context.Background(), accountId, callId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.ListCallRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCallRecordings`: []CallRecordingMetadata
	fmt.Fprintf(os.Stdout, "Response from `RecordingsAPI.ListCallRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**callId** | **string** | Programmable Voice API Call ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCallRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CallRecordingMetadata**](CallRecordingMetadata.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranscribeCallRecording

> TranscribeCallRecording(ctx, accountId, callId, recordingId).TranscribeRecording(transcribeRecording).Execute()

Create Transcription Request



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
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.
	transcribeRecording := *openapiclient.NewTranscribeRecording() // TranscribeRecording | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordingsAPI.TranscribeCallRecording(context.Background(), accountId, callId, recordingId).TranscribeRecording(transcribeRecording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.TranscribeCallRecording``: %v\n", err)
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
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranscribeCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transcribeRecording** | [**TranscribeRecording**](TranscribeRecording.md) |  | 

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


## UpdateCallRecordingState

> UpdateCallRecordingState(ctx, accountId, callId).UpdateCallRecording(updateCallRecording).Execute()

Update Recording



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
	updateCallRecording := *openapiclient.NewUpdateCallRecording(openapiclient.recordingStateEnum("paused")) // UpdateCallRecording | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecordingsAPI.UpdateCallRecordingState(context.Background(), accountId, callId).UpdateCallRecording(updateCallRecording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingsAPI.UpdateCallRecordingState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCallRecordingStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCallRecording** | [**UpdateCallRecording**](UpdateCallRecording.md) |  | 

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

