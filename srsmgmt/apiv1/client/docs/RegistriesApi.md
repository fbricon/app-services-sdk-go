# \RegistriesApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistriesApi.md#CreateRegistry) | **Post** /api/serviceregistry_mgmt/v1/registries | Create a new Registry instance
[**DeleteRegistry**](RegistriesApi.md#DeleteRegistry) | **Delete** /api/serviceregistry_mgmt/v1/{registryId} | Delete a Registry
[**GetRegistries**](RegistriesApi.md#GetRegistries) | **Get** /api/serviceregistry_mgmt/v1/registries | Get the list of all registries.
[**GetRegistry**](RegistriesApi.md#GetRegistry) | **Get** /api/serviceregistry_mgmt/v1/{registryId} | Get a Registry



## CreateRegistry

> Registry CreateRegistry(ctx).RegistryCreate(registryCreate).Execute()

Create a new Registry instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryCreate := *openapiclient.NewRegistryCreate() // RegistryCreate | A new `Registry` to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.CreateRegistry(context.Background()).RegistryCreate(registryCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.CreateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistry`: Registry
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.CreateRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryCreate** | [**RegistryCreate**](RegistryCreate.md) | A new &#x60;Registry&#x60; to be created. | 

### Return type

[**Registry**](Registry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> DeleteRegistry(ctx, registryId).Execute()

Delete a Registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryId := int32(56) // int32 | A unique identifier for a `Registry`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.DeleteRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **int32** | A unique identifier for a &#x60;Registry&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistries

> []Registry GetRegistries(ctx).Execute()

Get the list of all registries.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.GetRegistries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.GetRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistries`: []Registry
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.GetRegistries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistriesRequest struct via the builder pattern


### Return type

[**[]Registry**](Registry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> Registry GetRegistry(ctx, registryId).Execute()

Get a Registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryId := int32(56) // int32 | A unique identifier for a `Registry`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.GetRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.GetRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistry`: Registry
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **int32** | A unique identifier for a &#x60;Registry&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Registry**](Registry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
