# \ProductListAPI

All URIs are relative to *https://product-list.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateProductListAssociations**](ProductListAPI.md#BulkUpdateProductListAssociations) | **Post** /productlist.ProductList/BulkUpdateProductListAssociations | Bulk Update Product-Collection Associations
[**CreateProductList**](ProductListAPI.md#CreateProductList) | **Post** /productlist.ProductList/CreateProductList | Create Collection
[**CreateProductListAssociation**](ProductListAPI.md#CreateProductListAssociation) | **Post** /productlist.ProductList/CreateProductListAssociation | Create Product-Collection Association
[**DeleteProductList**](ProductListAPI.md#DeleteProductList) | **Post** /productlist.ProductList/DeleteProductList | Delete Collection
[**DeleteProductListAssociation**](ProductListAPI.md#DeleteProductListAssociation) | **Post** /productlist.ProductList/DeleteProductListAssociation | Delete Product-Collection Association
[**GetProductListAssociationsByProductGrn**](ProductListAPI.md#GetProductListAssociationsByProductGrn) | **Post** /productlist.ProductList/GetProductListAssociationsByProductGrn | Retrieve Collections for Product by GRN
[**GetProductListByCode**](ProductListAPI.md#GetProductListByCode) | **Post** /productlist.ProductList/GetProductListByCode | Retrieve Collection by Code
[**GetProductListById**](ProductListAPI.md#GetProductListById) | **Post** /productlist.ProductList/GetProductListById | Retrieve Collection by ID
[**GetProductListByUrlKey**](ProductListAPI.md#GetProductListByUrlKey) | **Post** /productlist.ProductList/GetProductListByUrlKey | Retrieve Collection by URL Key
[**GetProductListsCount**](ProductListAPI.md#GetProductListsCount) | **Post** /productlist.ProductList/GetProductListsCount | Get Collection Product Count
[**ListProductListAssociations**](ProductListAPI.md#ListProductListAssociations) | **Post** /productlist.ProductList/ListProductListAssociations | List Product Associations in Collection
[**ListProductLists**](ProductListAPI.md#ListProductLists) | **Post** /productlist.ProductList/ListProductLists | List All Collections
[**SearchProductLists**](ProductListAPI.md#SearchProductLists) | **Post** /productlist.ProductList/SearchProductLists | Search Collections
[**SearchProductListsByIds**](ProductListAPI.md#SearchProductListsByIds) | **Post** /productlist.ProductList/SearchProductListsByIds | Search Collections by IDs
[**UpdateProductList**](ProductListAPI.md#UpdateProductList) | **Post** /productlist.ProductList/UpdateProductList | Update Collection



## BulkUpdateProductListAssociations

> RpcStatus BulkUpdateProductListAssociations(ctx).Body(body).Execute()

Bulk Update Product-Collection Associations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistBulkUpdateProductListAssociationsRequest() // ProductlistBulkUpdateProductListAssociationsRequest | Used to perform bulk updates of product associations within a collection, such as modifying their positions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.BulkUpdateProductListAssociations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.BulkUpdateProductListAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateProductListAssociations`: RpcStatus
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.BulkUpdateProductListAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateProductListAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistBulkUpdateProductListAssociationsRequest**](ProductlistBulkUpdateProductListAssociationsRequest.md) | Used to perform bulk updates of product associations within a collection, such as modifying their positions. | 

### Return type

[**RpcStatus**](RpcStatus.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductList

> ProductlistCreateProductListResponse CreateProductList(ctx).Body(body).Execute()

Create Collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistCreateProductListRequest() // ProductlistCreateProductListRequest | Used to create a new product collection with details like code, URL key, and attributes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.CreateProductList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.CreateProductList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductList`: ProductlistCreateProductListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.CreateProductList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistCreateProductListRequest**](ProductlistCreateProductListRequest.md) | Used to create a new product collection with details like code, URL key, and attributes. | 

### Return type

[**ProductlistCreateProductListResponse**](ProductlistCreateProductListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductListAssociation

> ProductlistCreateProductListAssociationResponse CreateProductListAssociation(ctx).Body(body).Execute()

Create Product-Collection Association



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistCreateProductListAssociationRequest() // ProductlistCreateProductListAssociationRequest | Used to associate a product with a collection, specifying details like position and GRN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.CreateProductListAssociation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.CreateProductListAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductListAssociation`: ProductlistCreateProductListAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.CreateProductListAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductListAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistCreateProductListAssociationRequest**](ProductlistCreateProductListAssociationRequest.md) | Used to associate a product with a collection, specifying details like position and GRN. | 

### Return type

[**ProductlistCreateProductListAssociationResponse**](ProductlistCreateProductListAssociationResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProductList

> ProductlistDeleteProductListResponse DeleteProductList(ctx).Body(body).Execute()

Delete Collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistDeleteProductListRequest() // ProductlistDeleteProductListRequest | Defines the structure for DeleteProductListRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.DeleteProductList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.DeleteProductList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProductList`: ProductlistDeleteProductListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.DeleteProductList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistDeleteProductListRequest**](ProductlistDeleteProductListRequest.md) | Defines the structure for DeleteProductListRequest. | 

### Return type

[**ProductlistDeleteProductListResponse**](ProductlistDeleteProductListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProductListAssociation

> ProductlistDeleteProductListAssociationResponse DeleteProductListAssociation(ctx).Body(body).Execute()

Delete Product-Collection Association



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistDeleteProductListAssociationRequest() // ProductlistDeleteProductListAssociationRequest | Defines the structure for DeleteProductListAssociationRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.DeleteProductListAssociation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.DeleteProductListAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProductListAssociation`: ProductlistDeleteProductListAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.DeleteProductListAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductListAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistDeleteProductListAssociationRequest**](ProductlistDeleteProductListAssociationRequest.md) | Defines the structure for DeleteProductListAssociationRequest. | 

### Return type

[**ProductlistDeleteProductListAssociationResponse**](ProductlistDeleteProductListAssociationResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListAssociationsByProductGrn

> ProductlistGetProductListAssociationsByProductGrnResponse GetProductListAssociationsByProductGrn(ctx).Body(body).Execute()

Retrieve Collections for Product by GRN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistGetProductListAssociationsByProductGrnRequest() // ProductlistGetProductListAssociationsByProductGrnRequest | Defines the structure for GetProductListAssociationsByProductGrnRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.GetProductListAssociationsByProductGrn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.GetProductListAssociationsByProductGrn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListAssociationsByProductGrn`: ProductlistGetProductListAssociationsByProductGrnResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.GetProductListAssociationsByProductGrn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListAssociationsByProductGrnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistGetProductListAssociationsByProductGrnRequest**](ProductlistGetProductListAssociationsByProductGrnRequest.md) | Defines the structure for GetProductListAssociationsByProductGrnRequest. | 

### Return type

[**ProductlistGetProductListAssociationsByProductGrnResponse**](ProductlistGetProductListAssociationsByProductGrnResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListByCode

> ProductlistGetProductListByCodeResponse GetProductListByCode(ctx).Body(body).Execute()

Retrieve Collection by Code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistGetProductListByCodeRequest() // ProductlistGetProductListByCodeRequest | Defines the structure for GetProductListByCodeRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.GetProductListByCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.GetProductListByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListByCode`: ProductlistGetProductListByCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.GetProductListByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistGetProductListByCodeRequest**](ProductlistGetProductListByCodeRequest.md) | Defines the structure for GetProductListByCodeRequest. | 

### Return type

[**ProductlistGetProductListByCodeResponse**](ProductlistGetProductListByCodeResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListById

> ProductlistGetProductListByIdResponse GetProductListById(ctx).Body(body).Execute()

Retrieve Collection by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistGetProductListByIdRequest() // ProductlistGetProductListByIdRequest | Defines the structure for GetProductListByIdRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.GetProductListById(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.GetProductListById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListById`: ProductlistGetProductListByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.GetProductListById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistGetProductListByIdRequest**](ProductlistGetProductListByIdRequest.md) | Defines the structure for GetProductListByIdRequest. | 

### Return type

[**ProductlistGetProductListByIdResponse**](ProductlistGetProductListByIdResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListByUrlKey

> ProductlistGetProductListByUrlKeyResponse GetProductListByUrlKey(ctx).Body(body).Execute()

Retrieve Collection by URL Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistGetProductListByUrlKeyRequest() // ProductlistGetProductListByUrlKeyRequest | Defines the structure for GetProductListByUrlKeyRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.GetProductListByUrlKey(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.GetProductListByUrlKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListByUrlKey`: ProductlistGetProductListByUrlKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.GetProductListByUrlKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListByUrlKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistGetProductListByUrlKeyRequest**](ProductlistGetProductListByUrlKeyRequest.md) | Defines the structure for GetProductListByUrlKeyRequest. | 

### Return type

[**ProductlistGetProductListByUrlKeyResponse**](ProductlistGetProductListByUrlKeyResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListsCount

> ProductlistGetProductListsCountResponse GetProductListsCount(ctx).Body(body).Execute()

Get Collection Product Count



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistGetProductListsCountRequest() // ProductlistGetProductListsCountRequest | Defines the structure for GetProductListsCountRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.GetProductListsCount(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.GetProductListsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListsCount`: ProductlistGetProductListsCountResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.GetProductListsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistGetProductListsCountRequest**](ProductlistGetProductListsCountRequest.md) | Defines the structure for GetProductListsCountRequest. | 

### Return type

[**ProductlistGetProductListsCountResponse**](ProductlistGetProductListsCountResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductListAssociations

> ProductlistListProductListAssociationsResponse ListProductListAssociations(ctx).Body(body).Execute()

List Product Associations in Collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistListProductListAssociationsRequest() // ProductlistListProductListAssociationsRequest | Defines the structure for ListProductListAssociationsRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.ListProductListAssociations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.ListProductListAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductListAssociations`: ProductlistListProductListAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.ListProductListAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductListAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistListProductListAssociationsRequest**](ProductlistListProductListAssociationsRequest.md) | Defines the structure for ListProductListAssociationsRequest. | 

### Return type

[**ProductlistListProductListAssociationsResponse**](ProductlistListProductListAssociationsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductLists

> ProductlistListProductListsResponse ListProductLists(ctx).Body(body).Execute()

List All Collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistListProductListsRequest() // ProductlistListProductListsRequest | Defines the structure for ListProductListsRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.ListProductLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.ListProductLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductLists`: ProductlistListProductListsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.ListProductLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistListProductListsRequest**](ProductlistListProductListsRequest.md) | Defines the structure for ListProductListsRequest. | 

### Return type

[**ProductlistListProductListsResponse**](ProductlistListProductListsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProductLists

> ProductlistSearchProductListsResponse SearchProductLists(ctx).Body(body).Execute()

Search Collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistSearchProductListsRequest() // ProductlistSearchProductListsRequest | Defines the structure for SearchProductListsRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.SearchProductLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.SearchProductLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProductLists`: ProductlistSearchProductListsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.SearchProductLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistSearchProductListsRequest**](ProductlistSearchProductListsRequest.md) | Defines the structure for SearchProductListsRequest. | 

### Return type

[**ProductlistSearchProductListsResponse**](ProductlistSearchProductListsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProductListsByIds

> ProductlistSearchProductListsByIdsResponse SearchProductListsByIds(ctx).Body(body).Execute()

Search Collections by IDs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistSearchProductListsByIdsRequest() // ProductlistSearchProductListsByIdsRequest | Defines the structure for SearchProductListsByIdsRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.SearchProductListsByIds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.SearchProductListsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProductListsByIds`: ProductlistSearchProductListsByIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.SearchProductListsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductListsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistSearchProductListsByIdsRequest**](ProductlistSearchProductListsByIdsRequest.md) | Defines the structure for SearchProductListsByIdsRequest. | 

### Return type

[**ProductlistSearchProductListsByIdsResponse**](ProductlistSearchProductListsByIdsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductList

> ProductlistUpdateProductListResponse UpdateProductList(ctx).Body(body).Execute()

Update Collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-list"
)

func main() {
	body := *openapiclient.NewProductlistUpdateProductListRequest() // ProductlistUpdateProductListRequest | Defines the structure for UpdateProductListRequest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductListAPI.UpdateProductList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductListAPI.UpdateProductList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductList`: ProductlistUpdateProductListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductListAPI.UpdateProductList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductlistUpdateProductListRequest**](ProductlistUpdateProductListRequest.md) | Defines the structure for UpdateProductListRequest. | 

### Return type

[**ProductlistUpdateProductListResponse**](ProductlistUpdateProductListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

