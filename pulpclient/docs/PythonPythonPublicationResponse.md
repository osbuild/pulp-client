# PythonPythonPublicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**Distributions** | Pointer to **[]string** | This publication is currently being hosted as configured by these distributions. | [optional] [readonly] 

## Methods

### NewPythonPythonPublicationResponse

`func NewPythonPythonPublicationResponse() *PythonPythonPublicationResponse`

NewPythonPythonPublicationResponse instantiates a new PythonPythonPublicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPythonPublicationResponseWithDefaults

`func NewPythonPythonPublicationResponseWithDefaults() *PythonPythonPublicationResponse`

NewPythonPythonPublicationResponseWithDefaults instantiates a new PythonPythonPublicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *PythonPythonPublicationResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *PythonPythonPublicationResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *PythonPythonPublicationResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *PythonPythonPublicationResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *PythonPythonPublicationResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *PythonPythonPublicationResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *PythonPythonPublicationResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *PythonPythonPublicationResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *PythonPythonPublicationResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *PythonPythonPublicationResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *PythonPythonPublicationResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *PythonPythonPublicationResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *PythonPythonPublicationResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PythonPythonPublicationResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PythonPythonPublicationResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PythonPythonPublicationResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetDistributions

`func (o *PythonPythonPublicationResponse) GetDistributions() []string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *PythonPythonPublicationResponse) GetDistributionsOk() (*[]string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *PythonPythonPublicationResponse) SetDistributions(v []string)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *PythonPythonPublicationResponse) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


