# DebReleaseArchitecture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Architecture** | **string** | Name of the architecture. | 
**Distribution** | **string** | Name of the distribution. | 
**Codename** | **string** |  | 
**Suite** | **string** |  | 

## Methods

### NewDebReleaseArchitecture

`func NewDebReleaseArchitecture(architecture string, distribution string, codename string, suite string, ) *DebReleaseArchitecture`

NewDebReleaseArchitecture instantiates a new DebReleaseArchitecture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseArchitectureWithDefaults

`func NewDebReleaseArchitectureWithDefaults() *DebReleaseArchitecture`

NewDebReleaseArchitectureWithDefaults instantiates a new DebReleaseArchitecture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebReleaseArchitecture) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebReleaseArchitecture) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebReleaseArchitecture) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebReleaseArchitecture) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebReleaseArchitecture) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebReleaseArchitecture) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebReleaseArchitecture) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetDistribution

`func (o *DebReleaseArchitecture) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseArchitecture) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseArchitecture) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetCodename

`func (o *DebReleaseArchitecture) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebReleaseArchitecture) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebReleaseArchitecture) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetSuite

`func (o *DebReleaseArchitecture) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebReleaseArchitecture) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebReleaseArchitecture) SetSuite(v string)`

SetSuite sets Suite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


