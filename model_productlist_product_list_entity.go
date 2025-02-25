/*
Collection Service

The Collection Service provides a comprehensive API for managing product collections and categories. It supports the creation, retrieval, updating, and deletion of product collections, as well as managing associations between collections and products. Key functionalities include: - **Collection Management**: Create, update, and delete collections of products, enabling businesses to organize products into meaningful groups for easier browsing and marketing. - **Retrieval Operations**: Retrieve collections using various identifiers such as ID, URL key, or code, and list all existing collections with optional search and filter capabilities. - **Association Management**: Establish and manage associations between collections and products, allowing products to belong to one or multiple collections. - **Bulk Operations**: Perform bulk updates on product collection associations, optimizing the management of large-scale collections. - **Count and Search**: Get counts of products in a collection and search collections using filters or product identifiers for efficient management. - **Non-Hierarchical Structure**: While the service supports flat collections and categories, it does not yet include hierarchical relationships between categories. Designed for flexibility and scalability, the Collection Service empowers businesses to organize and present their product offerings effectively while maintaining a streamlined management process.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
	"time"
)

// checks if the ProductlistProductListEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistProductListEntity{}

// ProductlistProductListEntity Represents a product collection entity, including details like ID, code, attributes, and associations.
type ProductlistProductListEntity struct {
	// The unique identifier for the product collection.
	Id  *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	// A unique code representing the collection.
	Code   *string                   `json:"code,omitempty"`
	UrlKey *ProductlistLocalizedText `json:"urlKey,omitempty"`
	// Specifies the type of entity associated with the collection.
	EntityType *string `json:"entityType,omitempty"`
	// The code of the associated entity in the collection.
	EntityCode *string `json:"entityCode,omitempty"`
	// The timestamp indicating when the collection was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The timestamp of the last update to the collection.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A map of key-value pairs representing additional attributes of the collection.
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	// The total count of products associated with the collection.
	CountAssociations    *int64 `json:"countAssociations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistProductListEntity ProductlistProductListEntity

// NewProductlistProductListEntity instantiates a new ProductlistProductListEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistProductListEntity() *ProductlistProductListEntity {
	this := ProductlistProductListEntity{}
	return &this
}

// NewProductlistProductListEntityWithDefaults instantiates a new ProductlistProductListEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistProductListEntityWithDefaults() *ProductlistProductListEntity {
	this := ProductlistProductListEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistProductListEntity) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ProductlistProductListEntity) SetGrn(v string) {
	o.Grn = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductlistProductListEntity) SetCode(v string) {
	o.Code = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetUrlKey() ProductlistLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductlistLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetUrlKeyOk() (*ProductlistLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductlistLocalizedText and assigns it to the UrlKey field.
func (o *ProductlistProductListEntity) SetUrlKey(v ProductlistLocalizedText) {
	o.UrlKey = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *ProductlistProductListEntity) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetEntityCode() string {
	if o == nil || IsNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetEntityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityCode) {
		return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetEntityCode() bool {
	if o != nil && !IsNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *ProductlistProductListEntity) SetEntityCode(v string) {
	o.EntityCode = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProductlistProductListEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProductlistProductListEntity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductlistProductListEntity) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetCountAssociations returns the CountAssociations field value if set, zero value otherwise.
func (o *ProductlistProductListEntity) GetCountAssociations() int64 {
	if o == nil || IsNil(o.CountAssociations) {
		var ret int64
		return ret
	}
	return *o.CountAssociations
}

// GetCountAssociationsOk returns a tuple with the CountAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListEntity) GetCountAssociationsOk() (*int64, bool) {
	if o == nil || IsNil(o.CountAssociations) {
		return nil, false
	}
	return o.CountAssociations, true
}

// HasCountAssociations returns a boolean if a field has been set.
func (o *ProductlistProductListEntity) IsSetCountAssociations() bool {
	if o != nil && !IsNil(o.CountAssociations) {
		return true
	}

	return false
}

// SetCountAssociations gets a reference to the given int64 and assigns it to the CountAssociations field.
func (o *ProductlistProductListEntity) SetCountAssociations(v int64) {
	o.CountAssociations = &v
}

func (o ProductlistProductListEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistProductListEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.CountAssociations) {
		toSerialize["countAssociations"] = o.CountAssociations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistProductListEntity) UnmarshalJSON(data []byte) (err error) {
	varProductlistProductListEntity := _ProductlistProductListEntity{}

	err = json.Unmarshal(data, &varProductlistProductListEntity)

	if err != nil {
		return err
	}

	*o = ProductlistProductListEntity(varProductlistProductListEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "code")
		delete(additionalProperties, "urlKey")
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "entityCode")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "countAssociations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistProductListEntity) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistProductListEntity) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistProductListEntity struct {
	value *ProductlistProductListEntity
	isSet bool
}

func (v NullableProductlistProductListEntity) Get() *ProductlistProductListEntity {
	return v.value
}

func (v *NullableProductlistProductListEntity) Set(val *ProductlistProductListEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistProductListEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistProductListEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistProductListEntity(val *ProductlistProductListEntity) *NullableProductlistProductListEntity {
	return &NullableProductlistProductListEntity{value: val, isSet: true}
}

func (v NullableProductlistProductListEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistProductListEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
