# \DefaultAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHypershellV1FleetsGet**](DefaultAPI.md#ApiHypershellV1FleetsGet) | **Get** /api/hypershell/v1/fleets | Returns a list of fleets
[**ApiHypershellV1FleetsIdGet**](DefaultAPI.md#ApiHypershellV1FleetsIdGet) | **Get** /api/hypershell/v1/fleets/{id} | Get an fleet by id
[**ApiHypershellV1FleetsIdPatch**](DefaultAPI.md#ApiHypershellV1FleetsIdPatch) | **Patch** /api/hypershell/v1/fleets/{id} | Update an fleet
[**ApiHypershellV1FleetsPost**](DefaultAPI.md#ApiHypershellV1FleetsPost) | **Post** /api/hypershell/v1/fleets | Create a new fleet
[**ApiHypershellV1GatewayNetworksGet**](DefaultAPI.md#ApiHypershellV1GatewayNetworksGet) | **Get** /api/hypershell/v1/gateway_networks | Returns a list of gatewayNetworks
[**ApiHypershellV1GatewayNetworksIdGet**](DefaultAPI.md#ApiHypershellV1GatewayNetworksIdGet) | **Get** /api/hypershell/v1/gateway_networks/{id} | Get an gatewayNetwork by id
[**ApiHypershellV1GatewayNetworksIdPatch**](DefaultAPI.md#ApiHypershellV1GatewayNetworksIdPatch) | **Patch** /api/hypershell/v1/gateway_networks/{id} | Update an gatewayNetwork
[**ApiHypershellV1GatewayNetworksPost**](DefaultAPI.md#ApiHypershellV1GatewayNetworksPost) | **Post** /api/hypershell/v1/gateway_networks | Create a new gatewayNetwork
[**ApiHypershellV1GatewayReleasesGet**](DefaultAPI.md#ApiHypershellV1GatewayReleasesGet) | **Get** /api/hypershell/v1/gateway_releases | Returns a list of gatewayReleases
[**ApiHypershellV1GatewayReleasesIdGet**](DefaultAPI.md#ApiHypershellV1GatewayReleasesIdGet) | **Get** /api/hypershell/v1/gateway_releases/{id} | Get an gatewayRelease by id
[**ApiHypershellV1GatewayReleasesIdPatch**](DefaultAPI.md#ApiHypershellV1GatewayReleasesIdPatch) | **Patch** /api/hypershell/v1/gateway_releases/{id} | Update an gatewayRelease
[**ApiHypershellV1GatewayReleasesPost**](DefaultAPI.md#ApiHypershellV1GatewayReleasesPost) | **Post** /api/hypershell/v1/gateway_releases | Create a new gatewayRelease
[**ApiHypershellV1GatewaysGet**](DefaultAPI.md#ApiHypershellV1GatewaysGet) | **Get** /api/hypershell/v1/gateways | Returns a list of gateways
[**ApiHypershellV1GatewaysIdGet**](DefaultAPI.md#ApiHypershellV1GatewaysIdGet) | **Get** /api/hypershell/v1/gateways/{id} | Get an gateway by id
[**ApiHypershellV1GatewaysIdPatch**](DefaultAPI.md#ApiHypershellV1GatewaysIdPatch) | **Patch** /api/hypershell/v1/gateways/{id} | Update an gateway
[**ApiHypershellV1GatewaysPost**](DefaultAPI.md#ApiHypershellV1GatewaysPost) | **Post** /api/hypershell/v1/gateways | Create a new gateway
[**ApiHypershellV1ManagedClustersGet**](DefaultAPI.md#ApiHypershellV1ManagedClustersGet) | **Get** /api/hypershell/v1/managed_clusters | Returns a list of managedClusters
[**ApiHypershellV1ManagedClustersIdGet**](DefaultAPI.md#ApiHypershellV1ManagedClustersIdGet) | **Get** /api/hypershell/v1/managed_clusters/{id} | Get an managedCluster by id
[**ApiHypershellV1ManagedClustersIdPatch**](DefaultAPI.md#ApiHypershellV1ManagedClustersIdPatch) | **Patch** /api/hypershell/v1/managed_clusters/{id} | Update an managedCluster
[**ApiHypershellV1ManagedClustersPost**](DefaultAPI.md#ApiHypershellV1ManagedClustersPost) | **Post** /api/hypershell/v1/managed_clusters | Create a new managedCluster
[**ApiHypershellV1ManagedDatabasesGet**](DefaultAPI.md#ApiHypershellV1ManagedDatabasesGet) | **Get** /api/hypershell/v1/managed_databases | Returns a list of managedDatabases
[**ApiHypershellV1ManagedDatabasesIdGet**](DefaultAPI.md#ApiHypershellV1ManagedDatabasesIdGet) | **Get** /api/hypershell/v1/managed_databases/{id} | Get an managedDatabase by id
[**ApiHypershellV1ManagedDatabasesIdPatch**](DefaultAPI.md#ApiHypershellV1ManagedDatabasesIdPatch) | **Patch** /api/hypershell/v1/managed_databases/{id} | Update an managedDatabase
[**ApiHypershellV1ManagedDatabasesPost**](DefaultAPI.md#ApiHypershellV1ManagedDatabasesPost) | **Post** /api/hypershell/v1/managed_databases | Create a new managedDatabase
[**GetMetadata**](DefaultAPI.md#GetMetadata) | **Get** /api/hypershell/v1/metadata | Service metadata



## ApiHypershellV1FleetsGet

> FleetList ApiHypershellV1FleetsGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()

Returns a list of fleets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
	size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
	search := "search_example" // string | Specifies the search criteria (optional)
	orderBy := "orderBy_example" // string | Specifies the order by criteria (optional)
	fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1FleetsGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1FleetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1FleetsGet`: FleetList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1FleetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1FleetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria | 
 **orderBy** | **string** | Specifies the order by criteria | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned | 

### Return type

[**FleetList**](FleetList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1FleetsIdGet

> Fleet ApiHypershellV1FleetsIdGet(ctx, id).Execute()

Get an fleet by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1FleetsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1FleetsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1FleetsIdGet`: Fleet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1FleetsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1FleetsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Fleet**](Fleet.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1FleetsIdPatch

> Fleet ApiHypershellV1FleetsIdPatch(ctx, id).FleetPatchRequest(fleetPatchRequest).Execute()

Update an fleet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record
	fleetPatchRequest := *openapiclient.NewFleetPatchRequest() // FleetPatchRequest | Updated fleet data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1FleetsIdPatch(context.Background(), id).FleetPatchRequest(fleetPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1FleetsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1FleetsIdPatch`: Fleet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1FleetsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1FleetsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetPatchRequest** | [**FleetPatchRequest**](FleetPatchRequest.md) | Updated fleet data | 

### Return type

[**Fleet**](Fleet.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1FleetsPost

> Fleet ApiHypershellV1FleetsPost(ctx).Fleet(fleet).Execute()

Create a new fleet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fleet := *openapiclient.NewFleet("Name_example") // Fleet | Fleet data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1FleetsPost(context.Background()).Fleet(fleet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1FleetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1FleetsPost`: Fleet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1FleetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1FleetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleet** | [**Fleet**](Fleet.md) | Fleet data | 

### Return type

[**Fleet**](Fleet.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayNetworksGet

> GatewayNetworkList ApiHypershellV1GatewayNetworksGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()

Returns a list of gatewayNetworks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
	size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
	search := "search_example" // string | Specifies the search criteria (optional)
	orderBy := "orderBy_example" // string | Specifies the order by criteria (optional)
	fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayNetworksGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayNetworksGet`: GatewayNetworkList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria | 
 **orderBy** | **string** | Specifies the order by criteria | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned | 

### Return type

[**GatewayNetworkList**](GatewayNetworkList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayNetworksIdGet

> GatewayNetwork ApiHypershellV1GatewayNetworksIdGet(ctx, id).Execute()

Get an gatewayNetwork by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayNetworksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayNetworksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayNetworksIdGet`: GatewayNetwork
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayNetworksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayNetworksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayNetwork**](GatewayNetwork.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayNetworksIdPatch

> GatewayNetwork ApiHypershellV1GatewayNetworksIdPatch(ctx, id).GatewayNetworkPatchRequest(gatewayNetworkPatchRequest).Execute()

Update an gatewayNetwork

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record
	gatewayNetworkPatchRequest := *openapiclient.NewGatewayNetworkPatchRequest() // GatewayNetworkPatchRequest | Updated gatewayNetwork data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayNetworksIdPatch(context.Background(), id).GatewayNetworkPatchRequest(gatewayNetworkPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayNetworksIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayNetworksIdPatch`: GatewayNetwork
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayNetworksIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayNetworksIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayNetworkPatchRequest** | [**GatewayNetworkPatchRequest**](GatewayNetworkPatchRequest.md) | Updated gatewayNetwork data | 

### Return type

[**GatewayNetwork**](GatewayNetwork.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayNetworksPost

> GatewayNetwork ApiHypershellV1GatewayNetworksPost(ctx).GatewayNetwork(gatewayNetwork).Execute()

Create a new gatewayNetwork

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gatewayNetwork := *openapiclient.NewGatewayNetwork("Name_example", "FleetId_example") // GatewayNetwork | GatewayNetwork data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayNetworksPost(context.Background()).GatewayNetwork(gatewayNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayNetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayNetworksPost`: GatewayNetwork
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayNetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayNetwork** | [**GatewayNetwork**](GatewayNetwork.md) | GatewayNetwork data | 

### Return type

[**GatewayNetwork**](GatewayNetwork.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayReleasesGet

> GatewayReleaseList ApiHypershellV1GatewayReleasesGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()

Returns a list of gatewayReleases

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
	size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
	search := "search_example" // string | Specifies the search criteria (optional)
	orderBy := "orderBy_example" // string | Specifies the order by criteria (optional)
	fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayReleasesGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayReleasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayReleasesGet`: GatewayReleaseList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayReleasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayReleasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria | 
 **orderBy** | **string** | Specifies the order by criteria | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned | 

### Return type

[**GatewayReleaseList**](GatewayReleaseList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayReleasesIdGet

> GatewayRelease ApiHypershellV1GatewayReleasesIdGet(ctx, id).Execute()

Get an gatewayRelease by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayReleasesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayReleasesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayReleasesIdGet`: GatewayRelease
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayReleasesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayReleasesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayRelease**](GatewayRelease.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayReleasesIdPatch

> GatewayRelease ApiHypershellV1GatewayReleasesIdPatch(ctx, id).GatewayReleasePatchRequest(gatewayReleasePatchRequest).Execute()

Update an gatewayRelease

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record
	gatewayReleasePatchRequest := *openapiclient.NewGatewayReleasePatchRequest() // GatewayReleasePatchRequest | Updated gatewayRelease data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayReleasesIdPatch(context.Background(), id).GatewayReleasePatchRequest(gatewayReleasePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayReleasesIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayReleasesIdPatch`: GatewayRelease
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayReleasesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayReleasesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayReleasePatchRequest** | [**GatewayReleasePatchRequest**](GatewayReleasePatchRequest.md) | Updated gatewayRelease data | 

### Return type

[**GatewayRelease**](GatewayRelease.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewayReleasesPost

> GatewayRelease ApiHypershellV1GatewayReleasesPost(ctx).GatewayRelease(gatewayRelease).Execute()

Create a new gatewayRelease

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gatewayRelease := *openapiclient.NewGatewayRelease("Name_example", "FleetId_example", "Image_example") // GatewayRelease | GatewayRelease data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewayReleasesPost(context.Background()).GatewayRelease(gatewayRelease).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewayReleasesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewayReleasesPost`: GatewayRelease
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewayReleasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewayReleasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayRelease** | [**GatewayRelease**](GatewayRelease.md) | GatewayRelease data | 

### Return type

[**GatewayRelease**](GatewayRelease.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewaysGet

> GatewayList ApiHypershellV1GatewaysGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()

Returns a list of gateways

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
	size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
	search := "search_example" // string | Specifies the search criteria (optional)
	orderBy := "orderBy_example" // string | Specifies the order by criteria (optional)
	fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewaysGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewaysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewaysGet`: GatewayList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewaysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewaysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria | 
 **orderBy** | **string** | Specifies the order by criteria | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned | 

### Return type

[**GatewayList**](GatewayList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewaysIdGet

> Gateway ApiHypershellV1GatewaysIdGet(ctx, id).Execute()

Get an gateway by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewaysIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewaysIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewaysIdGet`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewaysIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewaysIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewaysIdPatch

> Gateway ApiHypershellV1GatewaysIdPatch(ctx, id).GatewayPatchRequest(gatewayPatchRequest).Execute()

Update an gateway

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record
	gatewayPatchRequest := *openapiclient.NewGatewayPatchRequest() // GatewayPatchRequest | Updated gateway data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewaysIdPatch(context.Background(), id).GatewayPatchRequest(gatewayPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewaysIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewaysIdPatch`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewaysIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewaysIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayPatchRequest** | [**GatewayPatchRequest**](GatewayPatchRequest.md) | Updated gateway data | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1GatewaysPost

> Gateway ApiHypershellV1GatewaysPost(ctx).Gateway(gateway).Execute()

Create a new gateway

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gateway := *openapiclient.NewGateway("Name_example", "FleetId_example", "ClusterId_example", "ReleaseId_example", "DatabaseId_example", "Namespace_example") // Gateway | Gateway data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1GatewaysPost(context.Background()).Gateway(gateway).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1GatewaysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1GatewaysPost`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1GatewaysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1GatewaysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gateway** | [**Gateway**](Gateway.md) | Gateway data | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedClustersGet

> ManagedClusterList ApiHypershellV1ManagedClustersGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()

Returns a list of managedClusters

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
	size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
	search := "search_example" // string | Specifies the search criteria (optional)
	orderBy := "orderBy_example" // string | Specifies the order by criteria (optional)
	fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedClustersGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedClustersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedClustersGet`: ManagedClusterList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedClustersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedClustersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria | 
 **orderBy** | **string** | Specifies the order by criteria | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned | 

### Return type

[**ManagedClusterList**](ManagedClusterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedClustersIdGet

> ManagedCluster ApiHypershellV1ManagedClustersIdGet(ctx, id).Execute()

Get an managedCluster by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedClustersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedClustersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedClustersIdGet`: ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedClustersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedClustersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedClustersIdPatch

> ManagedCluster ApiHypershellV1ManagedClustersIdPatch(ctx, id).ManagedClusterPatchRequest(managedClusterPatchRequest).Execute()

Update an managedCluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record
	managedClusterPatchRequest := *openapiclient.NewManagedClusterPatchRequest() // ManagedClusterPatchRequest | Updated managedCluster data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedClustersIdPatch(context.Background(), id).ManagedClusterPatchRequest(managedClusterPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedClustersIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedClustersIdPatch`: ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedClustersIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedClustersIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedClusterPatchRequest** | [**ManagedClusterPatchRequest**](ManagedClusterPatchRequest.md) | Updated managedCluster data | 

### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedClustersPost

> ManagedCluster ApiHypershellV1ManagedClustersPost(ctx).ManagedCluster(managedCluster).Execute()

Create a new managedCluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	managedCluster := *openapiclient.NewManagedCluster("Name_example", "FleetId_example", "Provider_example", "KubeconfigSecret_example") // ManagedCluster | ManagedCluster data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedClustersPost(context.Background()).ManagedCluster(managedCluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedClustersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedClustersPost`: ManagedCluster
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedClustersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedClustersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedCluster** | [**ManagedCluster**](ManagedCluster.md) | ManagedCluster data | 

### Return type

[**ManagedCluster**](ManagedCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedDatabasesGet

> ManagedDatabaseList ApiHypershellV1ManagedDatabasesGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()

Returns a list of managedDatabases

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
	size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
	search := "search_example" // string | Specifies the search criteria (optional)
	orderBy := "orderBy_example" // string | Specifies the order by criteria (optional)
	fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedDatabasesGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedDatabasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedDatabasesGet`: ManagedDatabaseList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedDatabasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedDatabasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria | 
 **orderBy** | **string** | Specifies the order by criteria | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned | 

### Return type

[**ManagedDatabaseList**](ManagedDatabaseList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedDatabasesIdGet

> ManagedDatabase ApiHypershellV1ManagedDatabasesIdGet(ctx, id).Execute()

Get an managedDatabase by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedDatabasesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedDatabasesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedDatabasesIdGet`: ManagedDatabase
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedDatabasesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedDatabasesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedDatabase**](ManagedDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedDatabasesIdPatch

> ManagedDatabase ApiHypershellV1ManagedDatabasesIdPatch(ctx, id).ManagedDatabasePatchRequest(managedDatabasePatchRequest).Execute()

Update an managedDatabase

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The id of record
	managedDatabasePatchRequest := *openapiclient.NewManagedDatabasePatchRequest() // ManagedDatabasePatchRequest | Updated managedDatabase data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedDatabasesIdPatch(context.Background(), id).ManagedDatabasePatchRequest(managedDatabasePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedDatabasesIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedDatabasesIdPatch`: ManagedDatabase
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedDatabasesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedDatabasesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedDatabasePatchRequest** | [**ManagedDatabasePatchRequest**](ManagedDatabasePatchRequest.md) | Updated managedDatabase data | 

### Return type

[**ManagedDatabase**](ManagedDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHypershellV1ManagedDatabasesPost

> ManagedDatabase ApiHypershellV1ManagedDatabasesPost(ctx).ManagedDatabase(managedDatabase).Execute()

Create a new managedDatabase

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	managedDatabase := *openapiclient.NewManagedDatabase("Name_example", "FleetId_example", "Provider_example") // ManagedDatabase | ManagedDatabase data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHypershellV1ManagedDatabasesPost(context.Background()).ManagedDatabase(managedDatabase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHypershellV1ManagedDatabasesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHypershellV1ManagedDatabasesPost`: ManagedDatabase
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHypershellV1ManagedDatabasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHypershellV1ManagedDatabasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedDatabase** | [**ManagedDatabase**](ManagedDatabase.md) | ManagedDatabase data | 

### Return type

[**ManagedDatabase**](ManagedDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadata

> ObjectReference GetMetadata(ctx).Execute()

Service metadata

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadata`: ObjectReference
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


### Return type

[**ObjectReference**](ObjectReference.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

