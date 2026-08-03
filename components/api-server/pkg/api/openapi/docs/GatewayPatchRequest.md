# GatewayPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FleetId** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ReleaseId** | Pointer to **string** |  | [optional] 
**DatabaseId** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ExternalDns** | Pointer to **string** |  | [optional] 
**TlsMode** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayPatchRequest

`func NewGatewayPatchRequest() *GatewayPatchRequest`

NewGatewayPatchRequest instantiates a new GatewayPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPatchRequestWithDefaults

`func NewGatewayPatchRequestWithDefaults() *GatewayPatchRequest`

NewGatewayPatchRequestWithDefaults instantiates a new GatewayPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewayPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFleetId

`func (o *GatewayPatchRequest) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *GatewayPatchRequest) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *GatewayPatchRequest) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.

### HasFleetId

`func (o *GatewayPatchRequest) HasFleetId() bool`

HasFleetId returns a boolean if a field has been set.

### GetClusterId

`func (o *GatewayPatchRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *GatewayPatchRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *GatewayPatchRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *GatewayPatchRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetReleaseId

`func (o *GatewayPatchRequest) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *GatewayPatchRequest) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *GatewayPatchRequest) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *GatewayPatchRequest) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetDatabaseId

`func (o *GatewayPatchRequest) GetDatabaseId() string`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *GatewayPatchRequest) GetDatabaseIdOk() (*string, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *GatewayPatchRequest) SetDatabaseId(v string)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *GatewayPatchRequest) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetNamespace

`func (o *GatewayPatchRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GatewayPatchRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GatewayPatchRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GatewayPatchRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetExternalDns

`func (o *GatewayPatchRequest) GetExternalDns() string`

GetExternalDns returns the ExternalDns field if non-nil, zero value otherwise.

### GetExternalDnsOk

`func (o *GatewayPatchRequest) GetExternalDnsOk() (*string, bool)`

GetExternalDnsOk returns a tuple with the ExternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDns

`func (o *GatewayPatchRequest) SetExternalDns(v string)`

SetExternalDns sets ExternalDns field to given value.

### HasExternalDns

`func (o *GatewayPatchRequest) HasExternalDns() bool`

HasExternalDns returns a boolean if a field has been set.

### GetTlsMode

`func (o *GatewayPatchRequest) GetTlsMode() string`

GetTlsMode returns the TlsMode field if non-nil, zero value otherwise.

### GetTlsModeOk

`func (o *GatewayPatchRequest) GetTlsModeOk() (*string, bool)`

GetTlsModeOk returns a tuple with the TlsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMode

`func (o *GatewayPatchRequest) SetTlsMode(v string)`

SetTlsMode sets TlsMode field to given value.

### HasTlsMode

`func (o *GatewayPatchRequest) HasTlsMode() bool`

HasTlsMode returns a boolean if a field has been set.

### GetServiceType

`func (o *GatewayPatchRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *GatewayPatchRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *GatewayPatchRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *GatewayPatchRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayPatchRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayPatchRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayPatchRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayPatchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPhase

`func (o *GatewayPatchRequest) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GatewayPatchRequest) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GatewayPatchRequest) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GatewayPatchRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


