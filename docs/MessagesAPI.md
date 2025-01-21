# \MessagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](MessagesAPI.md#CreateMessage) | **Post** /users/{accountId}/messages | Create Message
[**ListMessages**](MessagesAPI.md#ListMessages) | **Get** /users/{accountId}/messages | List Messages



## CreateMessage

> Message CreateMessage(ctx, accountId).MessageRequest(messageRequest).Execute()

Create Message



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
	messageRequest := *openapiclient.NewMessageRequest("93de2206-9669-4e07-948d-329f4b722ee2", []string{"To_example"}, "+15551113333") // MessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateMessage(context.Background(), accountId).MessageRequest(messageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessage`: Message
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageRequest** | [**MessageRequest**](MessageRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessages

> MessagesList ListMessages(ctx, accountId).MessageId(messageId).SourceTn(sourceTn).DestinationTn(destinationTn).MessageStatus(messageStatus).MessageDirection(messageDirection).CarrierName(carrierName).MessageType(messageType).ErrorCode(errorCode).FromDateTime(fromDateTime).ToDateTime(toDateTime).CampaignId(campaignId).Sort(sort).PageToken(pageToken).Limit(limit).LimitTotalCount(limitTotalCount).Execute()

List Messages



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
	messageId := "9e0df4ca-b18d-40d7-a59f-82fcdf5ae8e6" // string | The ID of the message to search for. Special characters need to be encoded using URL encoding. Message IDs could come in different formats, e.g., 9e0df4ca-b18d-40d7-a59f-82fcdf5ae8e6 and 1589228074636lm4k2je7j7jklbn2 are valid message ID formats. Note that you must include at least one query parameter. (optional)
	sourceTn := "%2B15554443333" // string | The phone number that sent the message. Accepted values are: a single full phone number a comma separated list of full phone numbers (maximum of 10) or a single partial phone number (minimum of 5 characters e.g. '%2B1919'). (optional)
	destinationTn := "%2B15554443333" // string | The phone number that received the message. Accepted values are: a single full phone number a comma separated list of full phone numbers (maximum of 10) or a single partial phone number (minimum of 5 characters e.g. '%2B1919'). (optional)
	messageStatus := openapiclient.messageStatusEnum("RECEIVED") // MessageStatusEnum | The status of the message. One of RECEIVED QUEUED SENDING SENT FAILED DELIVERED ACCEPTED UNDELIVERED. (optional)
	messageDirection := openapiclient.listMessageDirectionEnum("INBOUND") // ListMessageDirectionEnum | The direction of the message. One of INBOUND OUTBOUND. (optional)
	carrierName := "Verizon" // string | The name of the carrier used for this message. Possible values include but are not limited to Verizon and TMobile. Special characters need to be encoded using URL encoding (i.e. AT&T should be passed as AT%26T). (optional)
	messageType := openapiclient.messageTypeEnum("sms") // MessageTypeEnum | The type of message. Either sms or mms. (optional)
	errorCode := int32(9902) // int32 | The error code of the message. (optional)
	fromDateTime := "2022-09-14T18:20:16.000Z" // string | The start of the date range to search in ISO 8601 format. Uses the message receive time. The date range to search in is currently 14 days. (optional)
	toDateTime := "2022-09-14T18:20:16.000Z" // string | The end of the date range to search in ISO 8601 format. Uses the message receive time. The date range to search in is currently 14 days. (optional)
	campaignId := "CJEUMDK" // string | The campaign ID of the message. (optional)
	sort := "sourceTn:desc" // string | The field and direction to sort by combined with a colon. Direction is either asc or desc. (optional)
	pageToken := "gdEewhcJLQRB5" // string | A base64 encoded value used for pagination of results. (optional)
	limit := int32(50) // int32 | The maximum records requested in search result. Default 100. The sum of limit and after cannot be more than 10000. (optional)
	limitTotalCount := true // bool | When set to true, the response's totalCount field will have a maximum value of 10,000. When set to false, or excluded, this will give an accurate totalCount of all messages that match the provided filters. If you are experiencing latency, try using this parameter to limit your results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.ListMessages(context.Background(), accountId).MessageId(messageId).SourceTn(sourceTn).DestinationTn(destinationTn).MessageStatus(messageStatus).MessageDirection(messageDirection).CarrierName(carrierName).MessageType(messageType).ErrorCode(errorCode).FromDateTime(fromDateTime).ToDateTime(toDateTime).CampaignId(campaignId).Sort(sort).PageToken(pageToken).Limit(limit).LimitTotalCount(limitTotalCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ListMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMessages`: MessagesList
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Your Bandwidth Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageId** | **string** | The ID of the message to search for. Special characters need to be encoded using URL encoding. Message IDs could come in different formats, e.g., 9e0df4ca-b18d-40d7-a59f-82fcdf5ae8e6 and 1589228074636lm4k2je7j7jklbn2 are valid message ID formats. Note that you must include at least one query parameter. | 
 **sourceTn** | **string** | The phone number that sent the message. Accepted values are: a single full phone number a comma separated list of full phone numbers (maximum of 10) or a single partial phone number (minimum of 5 characters e.g. &#39;%2B1919&#39;). | 
 **destinationTn** | **string** | The phone number that received the message. Accepted values are: a single full phone number a comma separated list of full phone numbers (maximum of 10) or a single partial phone number (minimum of 5 characters e.g. &#39;%2B1919&#39;). | 
 **messageStatus** | [**MessageStatusEnum**](MessageStatusEnum.md) | The status of the message. One of RECEIVED QUEUED SENDING SENT FAILED DELIVERED ACCEPTED UNDELIVERED. | 
 **messageDirection** | [**ListMessageDirectionEnum**](ListMessageDirectionEnum.md) | The direction of the message. One of INBOUND OUTBOUND. | 
 **carrierName** | **string** | The name of the carrier used for this message. Possible values include but are not limited to Verizon and TMobile. Special characters need to be encoded using URL encoding (i.e. AT&amp;T should be passed as AT%26T). | 
 **messageType** | [**MessageTypeEnum**](MessageTypeEnum.md) | The type of message. Either sms or mms. | 
 **errorCode** | **int32** | The error code of the message. | 
 **fromDateTime** | **string** | The start of the date range to search in ISO 8601 format. Uses the message receive time. The date range to search in is currently 14 days. | 
 **toDateTime** | **string** | The end of the date range to search in ISO 8601 format. Uses the message receive time. The date range to search in is currently 14 days. | 
 **campaignId** | **string** | The campaign ID of the message. | 
 **sort** | **string** | The field and direction to sort by combined with a colon. Direction is either asc or desc. | 
 **pageToken** | **string** | A base64 encoded value used for pagination of results. | 
 **limit** | **int32** | The maximum records requested in search result. Default 100. The sum of limit and after cannot be more than 10000. | 
 **limitTotalCount** | **bool** | When set to true, the response&#39;s totalCount field will have a maximum value of 10,000. When set to false, or excluded, this will give an accurate totalCount of all messages that match the provided filters. If you are experiencing latency, try using this parameter to limit your results. | 

### Return type

[**MessagesList**](MessagesList.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

