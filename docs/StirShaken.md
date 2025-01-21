# StirShaken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verstat** | Pointer to **string** | (optional) The verification status indicating whether the verification was successful or not. Possible values are TN-Verification-Passed and TN-Verification-Failed. | [optional] 
**AttestationIndicator** | Pointer to **string** | (optional) The attestation level verified by Bandwidth. Possible values are A (full), B (partial) or C (gateway). | [optional] 
**OriginatingId** | Pointer to **string** | (optional) A unique origination identifier. | [optional] 

## Methods

### NewStirShaken

`func NewStirShaken() *StirShaken`

NewStirShaken instantiates a new StirShaken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStirShakenWithDefaults

`func NewStirShakenWithDefaults() *StirShaken`

NewStirShakenWithDefaults instantiates a new StirShaken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerstat

`func (o *StirShaken) GetVerstat() string`

GetVerstat returns the Verstat field if non-nil, zero value otherwise.

### GetVerstatOk

`func (o *StirShaken) GetVerstatOk() (*string, bool)`

GetVerstatOk returns a tuple with the Verstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerstat

`func (o *StirShaken) SetVerstat(v string)`

SetVerstat sets Verstat field to given value.

### HasVerstat

`func (o *StirShaken) HasVerstat() bool`

HasVerstat returns a boolean if a field has been set.

### GetAttestationIndicator

`func (o *StirShaken) GetAttestationIndicator() string`

GetAttestationIndicator returns the AttestationIndicator field if non-nil, zero value otherwise.

### GetAttestationIndicatorOk

`func (o *StirShaken) GetAttestationIndicatorOk() (*string, bool)`

GetAttestationIndicatorOk returns a tuple with the AttestationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationIndicator

`func (o *StirShaken) SetAttestationIndicator(v string)`

SetAttestationIndicator sets AttestationIndicator field to given value.

### HasAttestationIndicator

`func (o *StirShaken) HasAttestationIndicator() bool`

HasAttestationIndicator returns a boolean if a field has been set.

### GetOriginatingId

`func (o *StirShaken) GetOriginatingId() string`

GetOriginatingId returns the OriginatingId field if non-nil, zero value otherwise.

### GetOriginatingIdOk

`func (o *StirShaken) GetOriginatingIdOk() (*string, bool)`

GetOriginatingIdOk returns a tuple with the OriginatingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingId

`func (o *StirShaken) SetOriginatingId(v string)`

SetOriginatingId sets OriginatingId field to given value.

### HasOriginatingId

`func (o *StirShaken) HasOriginatingId() bool`

HasOriginatingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


