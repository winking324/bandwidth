# MachineDetectionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Possible values are answering-machine, human, silence, timeout, or error. | [optional] 
**Duration** | Pointer to **string** | The amount of time it took to determine the result. | [optional] 

## Methods

### NewMachineDetectionResult

`func NewMachineDetectionResult() *MachineDetectionResult`

NewMachineDetectionResult instantiates a new MachineDetectionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineDetectionResultWithDefaults

`func NewMachineDetectionResultWithDefaults() *MachineDetectionResult`

NewMachineDetectionResultWithDefaults instantiates a new MachineDetectionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *MachineDetectionResult) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MachineDetectionResult) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MachineDetectionResult) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MachineDetectionResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDuration

`func (o *MachineDetectionResult) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MachineDetectionResult) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MachineDetectionResult) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MachineDetectionResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


