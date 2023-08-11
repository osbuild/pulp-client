# DebReleaseArchitectureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Architecture** | **string** | Name of the architecture. | 
**Distribution** | **string** | Name of the distribution. | 
**Codename** | **string** |  | 
**Suite** | **string** |  | 

## Methods

### NewDebReleaseArchitectureResponse

`func NewDebReleaseArchitectureResponse(architecture string, distribution string, codename string, suite string, ) *DebReleaseArchitectureResponse`

NewDebReleaseArchitectureResponse instantiates a new DebReleaseArchitectureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebReleaseArchitectureResponseWithDefaults

`func NewDebReleaseArchitectureResponseWithDefaults() *DebReleaseArchitectureResponse`

NewDebReleaseArchitectureResponseWithDefaults instantiates a new DebReleaseArchitectureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *DebReleaseArchitectureResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *DebReleaseArchitectureResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *DebReleaseArchitectureResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *DebReleaseArchitectureResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *DebReleaseArchitectureResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *DebReleaseArchitectureResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *DebReleaseArchitectureResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *DebReleaseArchitectureResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetArchitecture

`func (o *DebReleaseArchitectureResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DebReleaseArchitectureResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DebReleaseArchitectureResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetDistribution

`func (o *DebReleaseArchitectureResponse) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DebReleaseArchitectureResponse) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DebReleaseArchitectureResponse) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.


### GetCodename

`func (o *DebReleaseArchitectureResponse) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *DebReleaseArchitectureResponse) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *DebReleaseArchitectureResponse) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetSuite

`func (o *DebReleaseArchitectureResponse) GetSuite() string`

GetSuite returns the Suite field if non-nil, zero value otherwise.

### GetSuiteOk

`func (o *DebReleaseArchitectureResponse) GetSuiteOk() (*string, bool)`

GetSuiteOk returns a tuple with the Suite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuite

`func (o *DebReleaseArchitectureResponse) SetSuite(v string)`

SetSuite sets Suite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


