# # ProductlistCreateProductListRequest
Used to create a new product collection with details like code, URL key, and attributes.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** | The unique identifier of the tenant creating the collection.  | [optional]
**Code**| **string** | A unique code to identify the collection.  | [optional]
**UrlKey**| [**ProductlistLocalizedText**](ProductlistLocalizedText.md) |   | [optional]
**EntityType**| **string** | Specifies the type of entity associated with the collection.  | [optional]
**EntityCode**| **string** | The code of the associated entity within the collection.  | [optional]
**Attributes**| [**map[string]ProtobufAny**](ProtobufAny.md) | Additional custom attributes for the collection, stored as key-value pairs.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

