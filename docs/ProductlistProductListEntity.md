# # ProductlistProductListEntity
Represents a product collection entity, including details like ID, code, attributes, and associations.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The unique identifier for the product collection.  | [optional]
**Grn**| **string** |   | [optional]
**Code**| **string** | A unique code representing the collection.  | [optional]
**UrlKey**| [**ProductlistLocalizedText**](ProductlistLocalizedText.md) |   | [optional]
**EntityType**| **string** | Specifies the type of entity associated with the collection.  | [optional]
**EntityCode**| **string** | The code of the associated entity in the collection.  | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) | The timestamp indicating when the collection was created.  | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) | The timestamp of the last update to the collection.  | [optional]
**Attributes**| [**map[string]ProtobufAny**](ProtobufAny.md) | A map of key-value pairs representing additional attributes of the collection.  | [optional]
**CountAssociations**| **int64** | The total count of products associated with the collection.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

