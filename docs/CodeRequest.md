# CodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The phone number to send the mfa code to. | 
**From** | **string** | The application phone number, the sender of the mfa code. | 
**ApplicationId** | **string** | The application unique ID, obtained from Bandwidth. | 
**Scope** | Pointer to **string** | An optional field to denote what scope or action the mfa code is addressing.  If not supplied, defaults to \&quot;2FA\&quot;. | [optional] 
**Message** | **string** | The message format of the mfa code.  There are three values that the system will replace \&quot;{CODE}\&quot;, \&quot;{NAME}\&quot;, \&quot;{SCOPE}\&quot;.  The \&quot;{SCOPE}\&quot; and \&quot;{NAME} value template are optional, while \&quot;{CODE}\&quot; must be supplied.  As the name would suggest, code will be replace with the actual mfa code.  Name is replaced with the application name, configured during provisioning of mfa.  The scope value is the same value sent during the call and partitioned by the server. | 
**Digits** | **int32** | The number of digits for your mfa code.  The valid number ranges from 2 to 8, inclusively. | 

## Methods

### NewCodeRequest

`func NewCodeRequest(to string, from string, applicationId string, message string, digits int32, ) *CodeRequest`

NewCodeRequest instantiates a new CodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeRequestWithDefaults

`func NewCodeRequestWithDefaults() *CodeRequest`

NewCodeRequestWithDefaults instantiates a new CodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *CodeRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CodeRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CodeRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *CodeRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CodeRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CodeRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetApplicationId

`func (o *CodeRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CodeRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CodeRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetScope

`func (o *CodeRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CodeRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CodeRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CodeRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetMessage

`func (o *CodeRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CodeRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CodeRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDigits

`func (o *CodeRequest) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *CodeRequest) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *CodeRequest) SetDigits(v int32)`

SetDigits sets Digits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


