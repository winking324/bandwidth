# \MFAAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateMessagingCode**](MFAAPI.md#GenerateMessagingCode) | **Post** /accounts/{accountId}/code/messaging | Messaging Authentication Code
[**GenerateVoiceCode**](MFAAPI.md#GenerateVoiceCode) | **Post** /accounts/{accountId}/code/voice | Voice Authentication Code
[**VerifyCode**](MFAAPI.md#VerifyCode) | **Post** /accounts/{accountId}/code/verify | Verify Authentication Code



## GenerateMessagingCode

> MessagingCodeResponse GenerateMessagingCode(ctx, accountId).CodeRequest(codeRequest).Execute()

Messaging Authentication Code



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
	codeRequest := *openapiclient.NewCodeRequest("+19195551234", "+19195554321", "66fd98ae-ac8d-a00f-7fcd-ba3280aeb9b1", "Your temporary {NAME} {SCOPE} code is {CODE}", int32(6)) // CodeRequest | MFA code request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.GenerateMessagingCode(context.Background(), accountId).CodeRequest(codeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.GenerateMessagingCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateMessagingCode`: MessagingCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.GenerateMessagingCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMessagingCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeRequest** | [**CodeRequest**](CodeRequest.md) | MFA code request body. | 

### Return type

[**MessagingCodeResponse**](MessagingCodeResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateVoiceCode

> VoiceCodeResponse GenerateVoiceCode(ctx, accountId).CodeRequest(codeRequest).Execute()

Voice Authentication Code



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
	codeRequest := *openapiclient.NewCodeRequest("+19195551234", "+19195554321", "66fd98ae-ac8d-a00f-7fcd-ba3280aeb9b1", "Your temporary {NAME} {SCOPE} code is {CODE}", int32(6)) // CodeRequest | MFA code request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.GenerateVoiceCode(context.Background(), accountId).CodeRequest(codeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.GenerateVoiceCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateVoiceCode`: VoiceCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.GenerateVoiceCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateVoiceCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeRequest** | [**CodeRequest**](CodeRequest.md) | MFA code request body. | 

### Return type

[**VoiceCodeResponse**](VoiceCodeResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCode

> VerifyCodeResponse VerifyCode(ctx, accountId).VerifyCodeRequest(verifyCodeRequest).Execute()

Verify Authentication Code



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
	verifyCodeRequest := *openapiclient.NewVerifyCodeRequest("+19195551234", float32(3), "123456") // VerifyCodeRequest | MFA code verify request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.VerifyCode(context.Background(), accountId).VerifyCodeRequest(verifyCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.VerifyCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyCode`: VerifyCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.VerifyCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyCodeRequest** | [**VerifyCodeRequest**](VerifyCodeRequest.md) | MFA code verify request body. | 

### Return type

[**VerifyCodeResponse**](VerifyCodeResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

