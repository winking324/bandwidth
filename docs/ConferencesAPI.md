# \ConferencesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadConferenceRecording**](ConferencesAPI.md#DownloadConferenceRecording) | **Get** /accounts/{accountId}/conferences/{conferenceId}/recordings/{recordingId}/media | Download Conference Recording
[**GetConference**](ConferencesAPI.md#GetConference) | **Get** /accounts/{accountId}/conferences/{conferenceId} | Get Conference Information
[**GetConferenceMember**](ConferencesAPI.md#GetConferenceMember) | **Get** /accounts/{accountId}/conferences/{conferenceId}/members/{memberId} | Get Conference Member
[**GetConferenceRecording**](ConferencesAPI.md#GetConferenceRecording) | **Get** /accounts/{accountId}/conferences/{conferenceId}/recordings/{recordingId} | Get Conference Recording Information
[**ListConferenceRecordings**](ConferencesAPI.md#ListConferenceRecordings) | **Get** /accounts/{accountId}/conferences/{conferenceId}/recordings | Get Conference Recordings
[**ListConferences**](ConferencesAPI.md#ListConferences) | **Get** /accounts/{accountId}/conferences | Get Conferences
[**UpdateConference**](ConferencesAPI.md#UpdateConference) | **Post** /accounts/{accountId}/conferences/{conferenceId} | Update Conference
[**UpdateConferenceBxml**](ConferencesAPI.md#UpdateConferenceBxml) | **Put** /accounts/{accountId}/conferences/{conferenceId}/bxml | Update Conference BXML
[**UpdateConferenceMember**](ConferencesAPI.md#UpdateConferenceMember) | **Put** /accounts/{accountId}/conferences/{conferenceId}/members/{memberId} | Update Conference Member



## DownloadConferenceRecording

> *os.File DownloadConferenceRecording(ctx, accountId, conferenceId, recordingId).Execute()

Download Conference Recording



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferencesAPI.DownloadConferenceRecording(context.Background(), accountId, conferenceId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.DownloadConferenceRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadConferenceRecording`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.DownloadConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadConferenceRecordingRequest struct via the builder pattern


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


## GetConference

> Conference GetConference(ctx, accountId, conferenceId).Execute()

Get Conference Information



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferencesAPI.GetConference(context.Background(), accountId, conferenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.GetConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConference`: Conference
	fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.GetConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Conference**](Conference.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConferenceMember

> ConferenceMember GetConferenceMember(ctx, accountId, conferenceId, memberId).Execute()

Get Conference Member



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.
	memberId := "c-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Conference Member ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferencesAPI.GetConferenceMember(context.Background(), accountId, conferenceId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.GetConferenceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConferenceMember`: ConferenceMember
	fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.GetConferenceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 
**memberId** | **string** | Programmable Voice API Conference Member ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConferenceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConferenceMember**](ConferenceMember.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConferenceRecording

> ConferenceRecordingMetadata GetConferenceRecording(ctx, accountId, conferenceId, recordingId).Execute()

Get Conference Recording Information



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.
	recordingId := "r-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Recording ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferencesAPI.GetConferenceRecording(context.Background(), accountId, conferenceId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.GetConferenceRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConferenceRecording`: ConferenceRecordingMetadata
	fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.GetConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 
**recordingId** | **string** | Programmable Voice API Recording ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConferenceRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConferenceRecordingMetadata**](ConferenceRecordingMetadata.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferenceRecordings

> []ConferenceRecordingMetadata ListConferenceRecordings(ctx, accountId, conferenceId).Execute()

Get Conference Recordings



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferencesAPI.ListConferenceRecordings(context.Background(), accountId, conferenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListConferenceRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConferenceRecordings`: []ConferenceRecordingMetadata
	fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListConferenceRecordings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConferenceRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ConferenceRecordingMetadata**](ConferenceRecordingMetadata.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferences

> []Conference ListConferences(ctx, accountId).Name(name).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).PageSize(pageSize).PageToken(pageToken).Execute()

Get Conferences



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
	name := "my-custom-name" // string | Filter results by the `name` field. (optional)
	minCreatedTime := "2022-06-21T19:13:21Z" // string | Filter results to conferences which have a `createdTime` after or at `minCreatedTime` (in ISO8601 format). (optional)
	maxCreatedTime := "2022-06-21T19:13:21Z" // string | Filter results to conferences which have a `createdTime` before or at `maxCreatedTime` (in ISO8601 format). (optional)
	pageSize := int32(500) // int32 | Specifies the max number of conferences that will be returned. (optional) (default to 1000)
	pageToken := "pageToken_example" // string | Not intended for explicit use. To use pagination, follow the links in the `Link` header of the response, as indicated in the endpoint description. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConferencesAPI.ListConferences(context.Background(), accountId).Name(name).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListConferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConferences`: []Conference
	fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListConferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter results by the &#x60;name&#x60; field. | 
 **minCreatedTime** | **string** | Filter results to conferences which have a &#x60;createdTime&#x60; after or at &#x60;minCreatedTime&#x60; (in ISO8601 format). | 
 **maxCreatedTime** | **string** | Filter results to conferences which have a &#x60;createdTime&#x60; before or at &#x60;maxCreatedTime&#x60; (in ISO8601 format). | 
 **pageSize** | **int32** | Specifies the max number of conferences that will be returned. | [default to 1000]
 **pageToken** | **string** | Not intended for explicit use. To use pagination, follow the links in the &#x60;Link&#x60; header of the response, as indicated in the endpoint description. | 

### Return type

[**[]Conference**](Conference.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConference

> UpdateConference(ctx, accountId, conferenceId).UpdateConference(updateConference).Execute()

Update Conference



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.
	updateConference := *openapiclient.NewUpdateConference() // UpdateConference | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConferencesAPI.UpdateConference(context.Background(), accountId, conferenceId).UpdateConference(updateConference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UpdateConference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConference** | [**UpdateConference**](UpdateConference.md) |  | 

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


## UpdateConferenceBxml

> UpdateConferenceBxml(ctx, accountId, conferenceId).Body(body).Execute()

Update Conference BXML



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.
	body := "<?xml version="1.0" encoding="UTF-8"?>
<Bxml>
    <StopRecording/>
</Bxml>" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConferencesAPI.UpdateConferenceBxml(context.Background(), accountId, conferenceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UpdateConferenceBxml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceBxmlRequest struct via the builder pattern


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


## UpdateConferenceMember

> UpdateConferenceMember(ctx, accountId, conferenceId, memberId).UpdateConferenceMember(updateConferenceMember).Execute()

Update Conference Member



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
	conferenceId := "conf-fe23a767-a75a5b77-20c5-4cca-b581-cbbf0776eca9" // string | Programmable Voice API Conference ID.
	memberId := "c-15ac29a2-1331029c-2cb0-4a07-b215-b22865662d85" // string | Programmable Voice API Conference Member ID.
	updateConferenceMember := *openapiclient.NewUpdateConferenceMember() // UpdateConferenceMember | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConferencesAPI.UpdateConferenceMember(context.Background(), accountId, conferenceId, memberId).UpdateConferenceMember(updateConferenceMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UpdateConferenceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**conferenceId** | **string** | Programmable Voice API Conference ID. | 
**memberId** | **string** | Programmable Voice API Conference Member ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConferenceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConferenceMember** | [**UpdateConferenceMember**](UpdateConferenceMember.md) |  | 

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

