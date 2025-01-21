# Media

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**ContentLength** | Pointer to **int32** |  | [optional] 
**MediaName** | Pointer to **string** |  | [optional] 

## Methods

### NewMedia

`func NewMedia() *Media`

NewMedia instantiates a new Media object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaWithDefaults

`func NewMediaWithDefaults() *Media`

NewMediaWithDefaults instantiates a new Media object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Media) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Media) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Media) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Media) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentLength

`func (o *Media) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *Media) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *Media) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *Media) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### GetMediaName

`func (o *Media) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *Media) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *Media) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *Media) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


