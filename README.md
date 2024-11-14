# Go API client for productlist

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import productlist "github.com/Gemini-Commerce/go-client-product-list"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `productlist.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), productlist.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `productlist.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), productlist.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `productlist.ContextOperationServerIndices` and `productlist.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), productlist.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), productlist.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://product-list.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ProductListAPI* | [**CreateProductList**](docs/ProductListAPI.md#createproductlist) | **Post** /productlist.ProductList/CreateProductList | Create Collection
*ProductListAPI* | [**CreateProductListAssociation**](docs/ProductListAPI.md#createproductlistassociation) | **Post** /productlist.ProductList/CreateProductListAssociation | Create Collection/Product Association
*ProductListAPI* | [**DeleteProductList**](docs/ProductListAPI.md#deleteproductlist) | **Post** /productlist.ProductList/DeleteProductList | Delete Collection
*ProductListAPI* | [**DeleteProductListAssociation**](docs/ProductListAPI.md#deleteproductlistassociation) | **Post** /productlist.ProductList/DeleteProductListAssociation | Delete Collection/Product Association
*ProductListAPI* | [**GetProductListAssociationsByProductGrn**](docs/ProductListAPI.md#getproductlistassociationsbyproductgrn) | **Post** /productlist.ProductList/GetProductListAssociationsByProductGrn | Get Collection/Product Associations by Product GRN
*ProductListAPI* | [**GetProductListByCode**](docs/ProductListAPI.md#getproductlistbycode) | **Post** /productlist.ProductList/GetProductListByCode | Get Collection by Code
*ProductListAPI* | [**GetProductListById**](docs/ProductListAPI.md#getproductlistbyid) | **Post** /productlist.ProductList/GetProductListById | Get Collection by Id
*ProductListAPI* | [**GetProductListByUrlKey**](docs/ProductListAPI.md#getproductlistbyurlkey) | **Post** /productlist.ProductList/GetProductListByUrlKey | Get Collection by Url Key
*ProductListAPI* | [**GetProductListsCount**](docs/ProductListAPI.md#getproductlistscount) | **Post** /productlist.ProductList/GetProductListsCount | Get Collection Product Count
*ProductListAPI* | [**ListProductListAssociations**](docs/ProductListAPI.md#listproductlistassociations) | **Post** /productlist.ProductList/ListProductListAssociations | List Collection/Product Associations
*ProductListAPI* | [**ListProductLists**](docs/ProductListAPI.md#listproductlists) | **Post** /productlist.ProductList/ListProductLists | List Collections
*ProductListAPI* | [**ProductListBulkUpdateProductListAssociations**](docs/ProductListAPI.md#productlistbulkupdateproductlistassociations) | **Post** /productlist.ProductList/BulkUpdateProductListAssociations | 
*ProductListAPI* | [**SearchProductLists**](docs/ProductListAPI.md#searchproductlists) | **Post** /productlist.ProductList/SearchProductLists | Search Collections
*ProductListAPI* | [**SearchProductListsByIds**](docs/ProductListAPI.md#searchproductlistsbyids) | **Post** /productlist.ProductList/SearchProductListsByIds | Search Collections by Ids
*ProductListAPI* | [**SetProductListAssociations**](docs/ProductListAPI.md#setproductlistassociations) | **Post** /productlist.ProductList/SetProductListAssociations | Set Collection/Product Associations
*ProductListAPI* | [**UpdateProductList**](docs/ProductListAPI.md#updateproductlist) | **Post** /productlist.ProductList/UpdateProductList | Update Collection


## Documentation For Models

 - [OrderByDirection](docs/OrderByDirection.md)
 - [ProductlistBulkUpdateProductListAssociationsRequest](docs/ProductlistBulkUpdateProductListAssociationsRequest.md)
 - [ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation](docs/ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation.md)
 - [ProductlistCreateProductListAssociationRequest](docs/ProductlistCreateProductListAssociationRequest.md)
 - [ProductlistCreateProductListAssociationResponse](docs/ProductlistCreateProductListAssociationResponse.md)
 - [ProductlistCreateProductListRequest](docs/ProductlistCreateProductListRequest.md)
 - [ProductlistCreateProductListResponse](docs/ProductlistCreateProductListResponse.md)
 - [ProductlistDeleteProductListAssociationRequest](docs/ProductlistDeleteProductListAssociationRequest.md)
 - [ProductlistDeleteProductListAssociationResponse](docs/ProductlistDeleteProductListAssociationResponse.md)
 - [ProductlistDeleteProductListRequest](docs/ProductlistDeleteProductListRequest.md)
 - [ProductlistDeleteProductListResponse](docs/ProductlistDeleteProductListResponse.md)
 - [ProductlistGetProductListAssociationsByProductGrnRequest](docs/ProductlistGetProductListAssociationsByProductGrnRequest.md)
 - [ProductlistGetProductListAssociationsByProductGrnResponse](docs/ProductlistGetProductListAssociationsByProductGrnResponse.md)
 - [ProductlistGetProductListByCodeRequest](docs/ProductlistGetProductListByCodeRequest.md)
 - [ProductlistGetProductListByCodeResponse](docs/ProductlistGetProductListByCodeResponse.md)
 - [ProductlistGetProductListByIdRequest](docs/ProductlistGetProductListByIdRequest.md)
 - [ProductlistGetProductListByIdResponse](docs/ProductlistGetProductListByIdResponse.md)
 - [ProductlistGetProductListByUrlKeyRequest](docs/ProductlistGetProductListByUrlKeyRequest.md)
 - [ProductlistGetProductListByUrlKeyResponse](docs/ProductlistGetProductListByUrlKeyResponse.md)
 - [ProductlistGetProductListsCountRequest](docs/ProductlistGetProductListsCountRequest.md)
 - [ProductlistGetProductListsCountResponse](docs/ProductlistGetProductListsCountResponse.md)
 - [ProductlistListProductListAssociationsRequest](docs/ProductlistListProductListAssociationsRequest.md)
 - [ProductlistListProductListAssociationsResponse](docs/ProductlistListProductListAssociationsResponse.md)
 - [ProductlistListProductListsRequest](docs/ProductlistListProductListsRequest.md)
 - [ProductlistListProductListsResponse](docs/ProductlistListProductListsResponse.md)
 - [ProductlistLocalizedText](docs/ProductlistLocalizedText.md)
 - [ProductlistOrderBy](docs/ProductlistOrderBy.md)
 - [ProductlistProductListAssociation](docs/ProductlistProductListAssociation.md)
 - [ProductlistProductListAssociationError](docs/ProductlistProductListAssociationError.md)
 - [ProductlistProductListEntity](docs/ProductlistProductListEntity.md)
 - [ProductlistProductListError](docs/ProductlistProductListError.md)
 - [ProductlistSearchProductListsByIdsRequest](docs/ProductlistSearchProductListsByIdsRequest.md)
 - [ProductlistSearchProductListsByIdsResponse](docs/ProductlistSearchProductListsByIdsResponse.md)
 - [ProductlistSearchProductListsRequest](docs/ProductlistSearchProductListsRequest.md)
 - [ProductlistSearchProductListsResponse](docs/ProductlistSearchProductListsResponse.md)
 - [ProductlistSetProductListAssociationsRequest](docs/ProductlistSetProductListAssociationsRequest.md)
 - [ProductlistSetProductListAssociationsResponse](docs/ProductlistSetProductListAssociationsResponse.md)
 - [ProductlistUpdateProductListRequest](docs/ProductlistUpdateProductListRequest.md)
 - [ProductlistUpdateProductListResponse](docs/ProductlistUpdateProductListResponse.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [SearchProductListsRequestQuery](docs/SearchProductListsRequestQuery.md)
 - [SetProductListAssociationsRequestAssociation](docs/SetProductListAssociationsRequestAssociation.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		productlist.ContextAPIKeys,
		map[string]productlist.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### standardAuthorization


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), productlist.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, productlist.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

