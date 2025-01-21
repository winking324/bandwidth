# LookupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | Pointer to **int32** | Our vendor&#39;s response code. | [optional] 
**Message** | Pointer to **string** | Message associated with the response code. | [optional] 
**E164Format** | Pointer to **string** | The telephone number in E.164 format. | [optional] 
**Formatted** | Pointer to **string** | The formatted version of the telephone number. | [optional] 
**Country** | Pointer to **string** | The country of the telephone number. | [optional] 
**LineType** | Pointer to **string** | The line type of the telephone number. | [optional] 
**LineProvider** | Pointer to **string** | The messaging service provider of the telephone number. | [optional] 
**MobileCountryCode** | Pointer to **string** | The first half of the Home Network Identity (HNI). | [optional] 
**MobileNetworkCode** | Pointer to **string** | The second half of the HNI. | [optional] 

## Methods

### NewLookupResult

`func NewLookupResult() *LookupResult`

NewLookupResult instantiates a new LookupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupResultWithDefaults

`func NewLookupResultWithDefaults() *LookupResult`

NewLookupResultWithDefaults instantiates a new LookupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *LookupResult) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *LookupResult) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *LookupResult) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *LookupResult) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetMessage

`func (o *LookupResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LookupResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LookupResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LookupResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetE164Format

`func (o *LookupResult) GetE164Format() string`

GetE164Format returns the E164Format field if non-nil, zero value otherwise.

### GetE164FormatOk

`func (o *LookupResult) GetE164FormatOk() (*string, bool)`

GetE164FormatOk returns a tuple with the E164Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE164Format

`func (o *LookupResult) SetE164Format(v string)`

SetE164Format sets E164Format field to given value.

### HasE164Format

`func (o *LookupResult) HasE164Format() bool`

HasE164Format returns a boolean if a field has been set.

### GetFormatted

`func (o *LookupResult) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *LookupResult) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *LookupResult) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *LookupResult) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### GetCountry

`func (o *LookupResult) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LookupResult) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LookupResult) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LookupResult) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLineType

`func (o *LookupResult) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *LookupResult) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *LookupResult) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *LookupResult) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetLineProvider

`func (o *LookupResult) GetLineProvider() string`

GetLineProvider returns the LineProvider field if non-nil, zero value otherwise.

### GetLineProviderOk

`func (o *LookupResult) GetLineProviderOk() (*string, bool)`

GetLineProviderOk returns a tuple with the LineProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineProvider

`func (o *LookupResult) SetLineProvider(v string)`

SetLineProvider sets LineProvider field to given value.

### HasLineProvider

`func (o *LookupResult) HasLineProvider() bool`

HasLineProvider returns a boolean if a field has been set.

### GetMobileCountryCode

`func (o *LookupResult) GetMobileCountryCode() string`

GetMobileCountryCode returns the MobileCountryCode field if non-nil, zero value otherwise.

### GetMobileCountryCodeOk

`func (o *LookupResult) GetMobileCountryCodeOk() (*string, bool)`

GetMobileCountryCodeOk returns a tuple with the MobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileCountryCode

`func (o *LookupResult) SetMobileCountryCode(v string)`

SetMobileCountryCode sets MobileCountryCode field to given value.

### HasMobileCountryCode

`func (o *LookupResult) HasMobileCountryCode() bool`

HasMobileCountryCode returns a boolean if a field has been set.

### GetMobileNetworkCode

`func (o *LookupResult) GetMobileNetworkCode() string`

GetMobileNetworkCode returns the MobileNetworkCode field if non-nil, zero value otherwise.

### GetMobileNetworkCodeOk

`func (o *LookupResult) GetMobileNetworkCodeOk() (*string, bool)`

GetMobileNetworkCodeOk returns a tuple with the MobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNetworkCode

`func (o *LookupResult) SetMobileNetworkCode(v string)`

SetMobileNetworkCode sets MobileNetworkCode field to given value.

### HasMobileNetworkCode

`func (o *LookupResult) HasMobileNetworkCode() bool`

HasMobileNetworkCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


