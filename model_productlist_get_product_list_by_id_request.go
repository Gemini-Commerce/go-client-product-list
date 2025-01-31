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

// checks if the ProductlistGetProductListByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistGetProductListByIdRequest{}

// ProductlistGetProductListByIdRequest Defines the structure for GetProductListByIdRequest.
type ProductlistGetProductListByIdRequest struct {
	// Field tenant_id of type string.
	TenantId *string `json:"tenantId,omitempty"`
	// Field id of type string.
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistGetProductListByIdRequest ProductlistGetProductListByIdRequest

// NewProductlistGetProductListByIdRequest instantiates a new ProductlistGetProductListByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistGetProductListByIdRequest() *ProductlistGetProductListByIdRequest {
	this := ProductlistGetProductListByIdRequest{}
	return &this
}

// NewProductlistGetProductListByIdRequestWithDefaults instantiates a new ProductlistGetProductListByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistGetProductListByIdRequestWithDefaults() *ProductlistGetProductListByIdRequest {
	this := ProductlistGetProductListByIdRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistGetProductListByIdRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListByIdRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistGetProductListByIdRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistGetProductListByIdRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistGetProductListByIdRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListByIdRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistGetProductListByIdRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistGetProductListByIdRequest) SetId(v string) {
	o.Id = &v
}

func (o ProductlistGetProductListByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistGetProductListByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistGetProductListByIdRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistGetProductListByIdRequest := _ProductlistGetProductListByIdRequest{}

	err = json.Unmarshal(data, &varProductlistGetProductListByIdRequest)

	if err != nil {
		return err
	}

	*o = ProductlistGetProductListByIdRequest(varProductlistGetProductListByIdRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistGetProductListByIdRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistGetProductListByIdRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistGetProductListByIdRequest struct {
	value *ProductlistGetProductListByIdRequest
	isSet bool
}

func (v NullableProductlistGetProductListByIdRequest) Get() *ProductlistGetProductListByIdRequest {
	return v.value
}

func (v *NullableProductlistGetProductListByIdRequest) Set(val *ProductlistGetProductListByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistGetProductListByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistGetProductListByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistGetProductListByIdRequest(val *ProductlistGetProductListByIdRequest) *NullableProductlistGetProductListByIdRequest {
	return &NullableProductlistGetProductListByIdRequest{value: val, isSet: true}
}

func (v NullableProductlistGetProductListByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistGetProductListByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
