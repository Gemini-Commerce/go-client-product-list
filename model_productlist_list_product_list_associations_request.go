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

// checks if the ProductlistListProductListAssociationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistListProductListAssociationsRequest{}

// ProductlistListProductListAssociationsRequest Defines the structure for ListProductListAssociationsRequest.
type ProductlistListProductListAssociationsRequest struct {
	// Field tenant_id of type string.
	TenantId *string `json:"tenantId,omitempty"`
	// Field list_id of type string.
	ListId *string `json:"listId,omitempty"`
	// Field product_grn of type string.
	ProductGrn *string `json:"productGrn,omitempty"`
	// Field page_size of type uint32.
	PageSize *int64 `json:"pageSize,omitempty"`
	// Field page_token of type string.
	PageToken *string `json:"pageToken,omitempty"`
	// Field order_by of type repeated OrderBy.
	OrderBy              []ProductlistOrderBy `json:"orderBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistListProductListAssociationsRequest ProductlistListProductListAssociationsRequest

// NewProductlistListProductListAssociationsRequest instantiates a new ProductlistListProductListAssociationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistListProductListAssociationsRequest() *ProductlistListProductListAssociationsRequest {
	this := ProductlistListProductListAssociationsRequest{}
	return &this
}

// NewProductlistListProductListAssociationsRequestWithDefaults instantiates a new ProductlistListProductListAssociationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistListProductListAssociationsRequestWithDefaults() *ProductlistListProductListAssociationsRequest {
	this := ProductlistListProductListAssociationsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistListProductListAssociationsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsRequest) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsRequest) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsRequest) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ProductlistListProductListAssociationsRequest) SetListId(v string) {
	o.ListId = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsRequest) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsRequest) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// HasProductGrn returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsRequest) HasProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *ProductlistListProductListAssociationsRequest) SetProductGrn(v string) {
	o.ProductGrn = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ProductlistListProductListAssociationsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ProductlistListProductListAssociationsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsRequest) GetOrderBy() []ProductlistOrderBy {
	if o == nil || IsNil(o.OrderBy) {
		var ret []ProductlistOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsRequest) GetOrderByOk() ([]ProductlistOrderBy, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsRequest) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []ProductlistOrderBy and assigns it to the OrderBy field.
func (o *ProductlistListProductListAssociationsRequest) SetOrderBy(v []ProductlistOrderBy) {
	o.OrderBy = v
}

func (o ProductlistListProductListAssociationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistListProductListAssociationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ListId) {
		toSerialize["listId"] = o.ListId
	}
	if !IsNil(o.ProductGrn) {
		toSerialize["productGrn"] = o.ProductGrn
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistListProductListAssociationsRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistListProductListAssociationsRequest := _ProductlistListProductListAssociationsRequest{}

	err = json.Unmarshal(data, &varProductlistListProductListAssociationsRequest)

	if err != nil {
		return err
	}

	*o = ProductlistListProductListAssociationsRequest(varProductlistListProductListAssociationsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "listId")
		delete(additionalProperties, "productGrn")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		delete(additionalProperties, "orderBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistListProductListAssociationsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistListProductListAssociationsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistListProductListAssociationsRequest struct {
	value *ProductlistListProductListAssociationsRequest
	isSet bool
}

func (v NullableProductlistListProductListAssociationsRequest) Get() *ProductlistListProductListAssociationsRequest {
	return v.value
}

func (v *NullableProductlistListProductListAssociationsRequest) Set(val *ProductlistListProductListAssociationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistListProductListAssociationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistListProductListAssociationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistListProductListAssociationsRequest(val *ProductlistListProductListAssociationsRequest) *NullableProductlistListProductListAssociationsRequest {
	return &NullableProductlistListProductListAssociationsRequest{value: val, isSet: true}
}

func (v NullableProductlistListProductListAssociationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistListProductListAssociationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
