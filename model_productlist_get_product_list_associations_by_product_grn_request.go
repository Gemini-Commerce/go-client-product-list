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
)

// checks if the ProductlistGetProductListAssociationsByProductGrnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistGetProductListAssociationsByProductGrnRequest{}

// ProductlistGetProductListAssociationsByProductGrnRequest Defines the structure for GetProductListAssociationsByProductGrnRequest.
type ProductlistGetProductListAssociationsByProductGrnRequest struct {
	// Field tenant_id of type string.
	TenantId *string `json:"tenantId,omitempty"`
	// Field product_grn of type string.
	ProductGrn           *string `json:"productGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistGetProductListAssociationsByProductGrnRequest ProductlistGetProductListAssociationsByProductGrnRequest

// NewProductlistGetProductListAssociationsByProductGrnRequest instantiates a new ProductlistGetProductListAssociationsByProductGrnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistGetProductListAssociationsByProductGrnRequest() *ProductlistGetProductListAssociationsByProductGrnRequest {
	this := ProductlistGetProductListAssociationsByProductGrnRequest{}
	return &this
}

// NewProductlistGetProductListAssociationsByProductGrnRequestWithDefaults instantiates a new ProductlistGetProductListAssociationsByProductGrnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistGetProductListAssociationsByProductGrnRequestWithDefaults() *ProductlistGetProductListAssociationsByProductGrnRequest {
	this := ProductlistGetProductListAssociationsByProductGrnRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// HasProductGrn returns a boolean if a field has been set.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) HasProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) SetProductGrn(v string) {
	o.ProductGrn = &v
}

func (o ProductlistGetProductListAssociationsByProductGrnRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistGetProductListAssociationsByProductGrnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ProductGrn) {
		toSerialize["productGrn"] = o.ProductGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistGetProductListAssociationsByProductGrnRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistGetProductListAssociationsByProductGrnRequest := _ProductlistGetProductListAssociationsByProductGrnRequest{}

	err = json.Unmarshal(data, &varProductlistGetProductListAssociationsByProductGrnRequest)

	if err != nil {
		return err
	}

	*o = ProductlistGetProductListAssociationsByProductGrnRequest(varProductlistGetProductListAssociationsByProductGrnRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "productGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistGetProductListAssociationsByProductGrnRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistGetProductListAssociationsByProductGrnRequest struct {
	value *ProductlistGetProductListAssociationsByProductGrnRequest
	isSet bool
}

func (v NullableProductlistGetProductListAssociationsByProductGrnRequest) Get() *ProductlistGetProductListAssociationsByProductGrnRequest {
	return v.value
}

func (v *NullableProductlistGetProductListAssociationsByProductGrnRequest) Set(val *ProductlistGetProductListAssociationsByProductGrnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistGetProductListAssociationsByProductGrnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistGetProductListAssociationsByProductGrnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistGetProductListAssociationsByProductGrnRequest(val *ProductlistGetProductListAssociationsByProductGrnRequest) *NullableProductlistGetProductListAssociationsByProductGrnRequest {
	return &NullableProductlistGetProductListAssociationsByProductGrnRequest{value: val, isSet: true}
}

func (v NullableProductlistGetProductListAssociationsByProductGrnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistGetProductListAssociationsByProductGrnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
