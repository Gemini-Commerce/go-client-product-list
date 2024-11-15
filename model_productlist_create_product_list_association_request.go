/*
Collection Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
)

// checks if the ProductlistCreateProductListAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistCreateProductListAssociationRequest{}

// ProductlistCreateProductListAssociationRequest struct for ProductlistCreateProductListAssociationRequest
type ProductlistCreateProductListAssociationRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ListId *string `json:"listId,omitempty"`
	Position *int32 `json:"position,omitempty"`
	ProductGrn *string `json:"productGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistCreateProductListAssociationRequest ProductlistCreateProductListAssociationRequest

// NewProductlistCreateProductListAssociationRequest instantiates a new ProductlistCreateProductListAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistCreateProductListAssociationRequest() *ProductlistCreateProductListAssociationRequest {
	this := ProductlistCreateProductListAssociationRequest{}
	return &this
}

// NewProductlistCreateProductListAssociationRequestWithDefaults instantiates a new ProductlistCreateProductListAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistCreateProductListAssociationRequestWithDefaults() *ProductlistCreateProductListAssociationRequest {
	this := ProductlistCreateProductListAssociationRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistCreateProductListAssociationRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListAssociationRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistCreateProductListAssociationRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistCreateProductListAssociationRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ProductlistCreateProductListAssociationRequest) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListAssociationRequest) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ProductlistCreateProductListAssociationRequest) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ProductlistCreateProductListAssociationRequest) SetListId(v string) {
	o.ListId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductlistCreateProductListAssociationRequest) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListAssociationRequest) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductlistCreateProductListAssociationRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ProductlistCreateProductListAssociationRequest) SetPosition(v int32) {
	o.Position = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *ProductlistCreateProductListAssociationRequest) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListAssociationRequest) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// HasProductGrn returns a boolean if a field has been set.
func (o *ProductlistCreateProductListAssociationRequest) HasProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *ProductlistCreateProductListAssociationRequest) SetProductGrn(v string) {
	o.ProductGrn = &v
}

func (o ProductlistCreateProductListAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistCreateProductListAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ListId) {
		toSerialize["listId"] = o.ListId
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ProductGrn) {
		toSerialize["productGrn"] = o.ProductGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistCreateProductListAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistCreateProductListAssociationRequest := _ProductlistCreateProductListAssociationRequest{}

	err = json.Unmarshal(data, &varProductlistCreateProductListAssociationRequest)

	if err != nil {
		return err
	}

	*o = ProductlistCreateProductListAssociationRequest(varProductlistCreateProductListAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "listId")
		delete(additionalProperties, "position")
		delete(additionalProperties, "productGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistCreateProductListAssociationRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistCreateProductListAssociationRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistCreateProductListAssociationRequest struct {
	value *ProductlistCreateProductListAssociationRequest
	isSet bool
}

func (v NullableProductlistCreateProductListAssociationRequest) Get() *ProductlistCreateProductListAssociationRequest {
	return v.value
}

func (v *NullableProductlistCreateProductListAssociationRequest) Set(val *ProductlistCreateProductListAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistCreateProductListAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistCreateProductListAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistCreateProductListAssociationRequest(val *ProductlistCreateProductListAssociationRequest) *NullableProductlistCreateProductListAssociationRequest {
	return &NullableProductlistCreateProductListAssociationRequest{value: val, isSet: true}
}

func (v NullableProductlistCreateProductListAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistCreateProductListAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


