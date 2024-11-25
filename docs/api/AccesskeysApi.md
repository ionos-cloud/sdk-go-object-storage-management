# \AccesskeysApi

All URIs are relative to *https://s3.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**AccesskeysDelete**](AccesskeysApi.md#AccesskeysDelete) | **Delete** /accesskeys/{accesskeyId} | Delete AccessKey|
|[**AccesskeysFindById**](AccesskeysApi.md#AccesskeysFindById) | **Get** /accesskeys/{accesskeyId} | Retrieve AccessKey|
|[**AccesskeysGet**](AccesskeysApi.md#AccesskeysGet) | **Get** /accesskeys | Retrieve all Accesskeys|
|[**AccesskeysPost**](AccesskeysApi.md#AccesskeysPost) | **Post** /accesskeys | Create AccessKey|
|[**AccesskeysPut**](AccesskeysApi.md#AccesskeysPut) | **Put** /accesskeys/{accesskeyId} | Ensure AccessKey|
|[**AccesskeysRenew**](AccesskeysApi.md#AccesskeysRenew) | **Put** /accesskeys/{accesskeyId}/renew | Ensure AccessKey|



## AccesskeysDelete

```go
var result  = AccesskeysDelete(ctx, accesskeyId)
                      .Execute()
```

Delete AccessKey



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
    accesskeyId := "fb68d39a-5706-51b4-b2b2-7b4bf9cbf0af" // string | The ID (UUID) of the AccessKey.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.AccesskeysApi.AccesskeysDelete(context.Background(), accesskeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccesskeysApi.AccesskeysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**accesskeyId** | **string** | The ID (UUID) of the AccessKey. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAccesskeysDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AccesskeysApiService.AccesskeysDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AccesskeysApiService.AccesskeysDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AccesskeysApiService.AccesskeysDelete": {
    "port": "8443",
},
})
```


## AccesskeysFindById

```go
var result AccessKeyRead = AccesskeysFindById(ctx, accesskeyId)
                      .Execute()
```

Retrieve AccessKey



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
    accesskeyId := "fb68d39a-5706-51b4-b2b2-7b4bf9cbf0af" // string | The ID (UUID) of the AccessKey.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AccesskeysApi.AccesskeysFindById(context.Background(), accesskeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccesskeysApi.AccesskeysFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AccesskeysFindById`: AccessKeyRead
    fmt.Fprintf(os.Stdout, "Response from `AccesskeysApi.AccesskeysFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**accesskeyId** | **string** | The ID (UUID) of the AccessKey. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAccesskeysFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**AccessKeyRead**](../models/AccessKeyRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AccesskeysApiService.AccesskeysFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AccesskeysApiService.AccesskeysFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AccesskeysApiService.AccesskeysFindById": {
    "port": "8443",
},
})
```


## AccesskeysGet

```go
var result AccessKeyReadList = AccesskeysGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterAccesskeyId(filterAccesskeyId)
                      .Execute()
```

Retrieve all Accesskeys



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
    filterAccesskeyId := "filterAccesskeyId_example" // string | The accesskey ID to filter by. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AccesskeysApi.AccesskeysGet(context.Background()).Offset(offset).Limit(limit).FilterAccesskeyId(filterAccesskeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccesskeysApi.AccesskeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AccesskeysGet`: AccessKeyReadList
    fmt.Fprintf(os.Stdout, "Response from `AccesskeysApi.AccesskeysGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiAccesskeysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterAccesskeyId** | **string** | The accesskey ID to filter by. | |

### Return type

[**AccessKeyReadList**](../models/AccessKeyReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AccesskeysApiService.AccesskeysGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AccesskeysApiService.AccesskeysGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AccesskeysApiService.AccesskeysGet": {
    "port": "8443",
},
})
```


## AccesskeysPost

```go
var result AccessKeyRead = AccesskeysPost(ctx)
                      .AccessKeyCreate(accessKeyCreate)
                      .Execute()
```

Create AccessKey



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
    accessKeyCreate := *openapiclient.NewAccessKeyCreate(*openapiclient.NewAccessKey("Access key details", "00000Du4KiV8A1tGUbjI2Ky1YmXqXnCpaQpgWw6_1V5fhMPRnAAAAAEIkfcQAAAAAAiR9xAioHCB_FWuNJ4un864Pouw", "tFVkUARsoeCdntQs2jVSyGG6TMPfPZ+ghnsWj/gG")) // AccessKeyCreate | AccessKey to create.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AccesskeysApi.AccesskeysPost(context.Background()).AccessKeyCreate(accessKeyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccesskeysApi.AccesskeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AccesskeysPost`: AccessKeyRead
    fmt.Fprintf(os.Stdout, "Response from `AccesskeysApi.AccesskeysPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiAccesskeysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **accessKeyCreate** | [**AccessKeyCreate**](../models/AccessKeyCreate.md) | AccessKey to create. | |

### Return type

[**AccessKeyRead**](../models/AccessKeyRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AccesskeysApiService.AccesskeysPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AccesskeysApiService.AccesskeysPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AccesskeysApiService.AccesskeysPost": {
    "port": "8443",
},
})
```


## AccesskeysPut

```go
var result AccessKeyRead = AccesskeysPut(ctx, accesskeyId)
                      .AccessKeyEnsure(accessKeyEnsure)
                      .Execute()
```

Ensure AccessKey



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
    accesskeyId := "fb68d39a-5706-51b4-b2b2-7b4bf9cbf0af" // string | The ID (UUID) of the AccessKey.
    accessKeyEnsure := *openapiclient.NewAccessKeyEnsure("fb68d39a-5706-51b4-b2b2-7b4bf9cbf0af", *openapiclient.NewAccessKey("Access key details", "00000Du4KiV8A1tGUbjI2Ky1YmXqXnCpaQpgWw6_1V5fhMPRnAAAAAEIkfcQAAAAAAiR9xAioHCB_FWuNJ4un864Pouw", "tFVkUARsoeCdntQs2jVSyGG6TMPfPZ+ghnsWj/gG")) // AccessKeyEnsure | update AccessKey

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AccesskeysApi.AccesskeysPut(context.Background(), accesskeyId).AccessKeyEnsure(accessKeyEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccesskeysApi.AccesskeysPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AccesskeysPut`: AccessKeyRead
    fmt.Fprintf(os.Stdout, "Response from `AccesskeysApi.AccesskeysPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**accesskeyId** | **string** | The ID (UUID) of the AccessKey. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAccesskeysPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **accessKeyEnsure** | [**AccessKeyEnsure**](../models/AccessKeyEnsure.md) | update AccessKey | |

### Return type

[**AccessKeyRead**](../models/AccessKeyRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AccesskeysApiService.AccesskeysPut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AccesskeysApiService.AccesskeysPut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AccesskeysApiService.AccesskeysPut": {
    "port": "8443",
},
})
```


## AccesskeysRenew

```go
var result AccessKeyRead = AccesskeysRenew(ctx, accesskeyId)
                      .Execute()
```

Ensure AccessKey



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
    accesskeyId := "fb68d39a-5706-51b4-b2b2-7b4bf9cbf0af" // string | The ID (UUID) of the AccessKey that should be ensured.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AccesskeysApi.AccesskeysRenew(context.Background(), accesskeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccesskeysApi.AccesskeysRenew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AccesskeysRenew`: AccessKeyRead
    fmt.Fprintf(os.Stdout, "Response from `AccesskeysApi.AccesskeysRenew`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**accesskeyId** | **string** | The ID (UUID) of the AccessKey that should be ensured. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAccesskeysRenewRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**AccessKeyRead**](../models/AccessKeyRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AccesskeysApiService.AccesskeysRenew"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AccesskeysApiService.AccesskeysRenew": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AccesskeysApiService.AccesskeysRenew": {
    "port": "8443",
},
})
```

