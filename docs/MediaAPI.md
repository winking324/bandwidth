# \MediaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMedia**](MediaAPI.md#DeleteMedia) | **Delete** /users/{accountId}/media/{mediaId} | Delete Media
[**GetMedia**](MediaAPI.md#GetMedia) | **Get** /users/{accountId}/media/{mediaId} | Get Media
[**ListMedia**](MediaAPI.md#ListMedia) | **Get** /users/{accountId}/media | List Media
[**UploadMedia**](MediaAPI.md#UploadMedia) | **Put** /users/{accountId}/media/{mediaId} | Upload Media



## DeleteMedia

> DeleteMedia(ctx, accountId, mediaId).Execute()

Delete Media



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
	mediaId := "14762070468292kw2fuqty55yp2b2/0/bw.png" // string | Media ID to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaAPI.DeleteMedia(context.Background(), accountId, mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.DeleteMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**mediaId** | **string** | Media ID to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaRequest struct via the builder pattern


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


## GetMedia

> *os.File GetMedia(ctx, accountId, mediaId).Execute()

Get Media



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
	mediaId := "14762070468292kw2fuqty55yp2b2/0/bw.png" // string | Media ID to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.GetMedia(context.Background(), accountId, mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.GetMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMedia`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.GetMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**mediaId** | **string** | Media ID to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMedia

> []Media ListMedia(ctx, accountId).ContinuationToken(continuationToken).Execute()

List Media



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
	continuationToken := "1XEi2tsFtLo1JbtLwETnM1ZJ+PqAa8w6ENvC5QKvwyrCDYII663Gy5M4s40owR1tjkuWUif6qbWvFtQJR5/ipqbUnfAqL254LKNlPy6tATCzioKSuHuOqgzloDkSwRtX0LtcL2otHS69hK343m+SjdL+vlj71tT39" // string | Continuation token used to retrieve subsequent media. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.ListMedia(context.Background(), accountId).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.ListMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMedia`: []Media
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.ListMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **string** | Continuation token used to retrieve subsequent media. | 

### Return type

[**[]Media**](Media.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMedia

> UploadMedia(ctx, accountId, mediaId).Body(body).ContentType(contentType).CacheControl(cacheControl).Execute()

Upload Media



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
	mediaId := "14762070468292kw2fuqty55yp2b2/0/bw.png" // string | Media ID to retrieve.
	body := os.NewFile(1234, "some_file") // *os.File | 
	contentType := "audio/wav" // string | The media type of the entity-body. (optional)
	cacheControl := "no-cache" // string | General-header field is used to specify directives that MUST be obeyed by all caching mechanisms along the request/response chain. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaAPI.UploadMedia(context.Background(), accountId, mediaId).Body(body).ContentType(contentType).CacheControl(cacheControl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.UploadMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 
**mediaId** | **string** | Media ID to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 
 **contentType** | **string** | The media type of the entity-body. | 
 **cacheControl** | **string** | General-header field is used to specify directives that MUST be obeyed by all caching mechanisms along the request/response chain. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, application/ogg, application/pdf, application/rtf, application/zip, application/x-tar, application/xml, application/gzip, application/x-bzip2, application/x-gzip, application/smil, application/javascript, audio/mp4, audio/mpeg, audio/ogg, audio/flac, audio/webm, audio/wav, audio/amr, audio/3gpp, image/bmp, image/gif, image/heic, image/heif, image/jpeg, image/pjpeg, image/png, image/svg+xml, image/tiff, image/webp, image/x-icon, text/css, text/csv, text/calendar, text/html, text/plain, text/javascript, text/vcard, text/vnd.wap.wml, text/xml, video/avi, video/mp4, video/mpeg, video/ogg, video/quicktime, video/webm, video/x-ms-wmv, video/x-flv
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

