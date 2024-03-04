/*
product-list/bulk_update_product_list_associations.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-list

import (
	"encoding/json"
)

// checks if the ProductlistSetProductListAssociationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistSetProductListAssociationsRequest{}

// ProductlistSetProductListAssociationsRequest struct for ProductlistSetProductListAssociationsRequest
type ProductlistSetProductListAssociationsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ListId *string `json:"listId,omitempty"`
	ProductGrn *string `json:"productGrn,omitempty"`
	Associations []SetProductListAssociationsRequestAssociation `json:"associations,omitempty"`
}

// NewProductlistSetProductListAssociationsRequest instantiates a new ProductlistSetProductListAssociationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistSetProductListAssociationsRequest() *ProductlistSetProductListAssociationsRequest {
	this := ProductlistSetProductListAssociationsRequest{}
	return &this
}

// NewProductlistSetProductListAssociationsRequestWithDefaults instantiates a new ProductlistSetProductListAssociationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistSetProductListAssociationsRequestWithDefaults() *ProductlistSetProductListAssociationsRequest {
	this := ProductlistSetProductListAssociationsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistSetProductListAssociationsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSetProductListAssociationsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistSetProductListAssociationsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistSetProductListAssociationsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ProductlistSetProductListAssociationsRequest) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSetProductListAssociationsRequest) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ProductlistSetProductListAssociationsRequest) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ProductlistSetProductListAssociationsRequest) SetListId(v string) {
	o.ListId = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *ProductlistSetProductListAssociationsRequest) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSetProductListAssociationsRequest) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// HasProductGrn returns a boolean if a field has been set.
func (o *ProductlistSetProductListAssociationsRequest) HasProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *ProductlistSetProductListAssociationsRequest) SetProductGrn(v string) {
	o.ProductGrn = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *ProductlistSetProductListAssociationsRequest) GetAssociations() []SetProductListAssociationsRequestAssociation {
	if o == nil || IsNil(o.Associations) {
		var ret []SetProductListAssociationsRequestAssociation
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSetProductListAssociationsRequest) GetAssociationsOk() ([]SetProductListAssociationsRequestAssociation, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *ProductlistSetProductListAssociationsRequest) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []SetProductListAssociationsRequestAssociation and assigns it to the Associations field.
func (o *ProductlistSetProductListAssociationsRequest) SetAssociations(v []SetProductListAssociationsRequestAssociation) {
	o.Associations = v
}

func (o ProductlistSetProductListAssociationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistSetProductListAssociationsRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	return toSerialize, nil
}

type NullableProductlistSetProductListAssociationsRequest struct {
	value *ProductlistSetProductListAssociationsRequest
	isSet bool
}

func (v NullableProductlistSetProductListAssociationsRequest) Get() *ProductlistSetProductListAssociationsRequest {
	return v.value
}

func (v *NullableProductlistSetProductListAssociationsRequest) Set(val *ProductlistSetProductListAssociationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistSetProductListAssociationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistSetProductListAssociationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistSetProductListAssociationsRequest(val *ProductlistSetProductListAssociationsRequest) *NullableProductlistSetProductListAssociationsRequest {
	return &NullableProductlistSetProductListAssociationsRequest{value: val, isSet: true}
}

func (v NullableProductlistSetProductListAssociationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistSetProductListAssociationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


