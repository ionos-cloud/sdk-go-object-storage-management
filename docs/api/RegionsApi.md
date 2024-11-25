# \RegionsApi

All URIs are relative to *https://s3.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegionsFindByRegion**](RegionsApi.md#RegionsFindByRegion) | **Get** /regions/{region} | Retrieve Region|
|[**RegionsGet**](RegionsApi.md#RegionsGet) | **Get** /regions | Retrieve all Regions|



## RegionsFindByRegion

```go
var result RegionRead = RegionsFindByRegion(ctx, region)
                      .Execute()
```

Retrieve Region



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-object-storage-management"
)

func main() {
    region := "eu-central-3" // string | The Region of the Region.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegionsApi.RegionsFindByRegion(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.RegionsFindByRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegionsFindByRegion`: RegionRead
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.RegionsFindByRegion`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**region** | **string** | The Region of the Region. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegionsFindByRegionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RegionRead**](../models/RegionRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegionsApiService.RegionsFindByRegion"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegionsApiService.RegionsFindByRegion": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegionsApiService.RegionsFindByRegion": {
    "port": "8443",
},
})
```


## RegionsGet

```go
var result RegionReadList = RegionsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all Regions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-object-storage-management"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegionsApi.RegionsGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionsApi.RegionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegionsGet`: RegionReadList
    fmt.Fprintf(os.Stdout, "Response from `RegionsApi.RegionsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRegionsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**RegionReadList**](../models/RegionReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegionsApiService.RegionsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegionsApiService.RegionsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegionsApiService.RegionsGet": {
    "port": "8443",
},
})
```

