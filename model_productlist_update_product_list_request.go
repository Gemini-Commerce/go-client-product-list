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

// checks if the ProductlistUpdateProductListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistUpdateProductListRequest{}

// ProductlistUpdateProductListRequest Defines the structure for UpdateProductListRequest.
type ProductlistUpdateProductListRequest struct {
	// Field tenant_id of type string.
	TenantId *string `json:"tenantId,omitempty"`
	// Field id of type string.
	Id     *string                   `json:"id,omitempty"`
	UrlKey *ProductlistLocalizedText `json:"urlKey,omitempty"`
	// Field attributes of type map<string, google.protobuf.Any>.
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	// Field attributes_mask of type google.protobuf.FieldMask.
	AttributesMask       *string `json:"attributesMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistUpdateProductListRequest ProductlistUpdateProductListRequest

// NewProductlistUpdateProductListRequest instantiates a new ProductlistUpdateProductListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistUpdateProductListRequest() *ProductlistUpdateProductListRequest {
	this := ProductlistUpdateProductListRequest{}
	return &this
}

// NewProductlistUpdateProductListRequestWithDefaults instantiates a new ProductlistUpdateProductListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistUpdateProductListRequestWithDefaults() *ProductlistUpdateProductListRequest {
	this := ProductlistUpdateProductListRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) IsSetTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistUpdateProductListRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistUpdateProductListRequest) SetId(v string) {
	o.Id = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetUrlKey() ProductlistLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductlistLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetUrlKeyOk() (*ProductlistLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) IsSetUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductlistLocalizedText and assigns it to the UrlKey field.
func (o *ProductlistUpdateProductListRequest) SetUrlKey(v ProductlistLocalizedText) {
	o.UrlKey = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) IsSetAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductlistUpdateProductListRequest) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetAttributesMask returns the AttributesMask field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetAttributesMask() string {
	if o == nil || IsNil(o.AttributesMask) {
		var ret string
		return ret
	}
	return *o.AttributesMask
}

// GetAttributesMaskOk returns a tuple with the AttributesMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetAttributesMaskOk() (*string, bool) {
	if o == nil || IsNil(o.AttributesMask) {
		return nil, false
	}
	return o.AttributesMask, true
}

// HasAttributesMask returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) IsSetAttributesMask() bool {
	if o != nil && !IsNil(o.AttributesMask) {
		return true
	}

	return false
}

// SetAttributesMask gets a reference to the given string and assigns it to the AttributesMask field.
func (o *ProductlistUpdateProductListRequest) SetAttributesMask(v string) {
	o.AttributesMask = &v
}

func (o ProductlistUpdateProductListRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistUpdateProductListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.AttributesMask) {
		toSerialize["attributesMask"] = o.AttributesMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistUpdateProductListRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistUpdateProductListRequest := _ProductlistUpdateProductListRequest{}

	err = json.Unmarshal(data, &varProductlistUpdateProductListRequest)

	if err != nil {
		return err
	}

	*o = ProductlistUpdateProductListRequest(varProductlistUpdateProductListRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "urlKey")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "attributesMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistUpdateProductListRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistUpdateProductListRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistUpdateProductListRequest struct {
	value *ProductlistUpdateProductListRequest
	isSet bool
}

func (v NullableProductlistUpdateProductListRequest) Get() *ProductlistUpdateProductListRequest {
	return v.value
}

func (v *NullableProductlistUpdateProductListRequest) Set(val *ProductlistUpdateProductListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistUpdateProductListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistUpdateProductListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistUpdateProductListRequest(val *ProductlistUpdateProductListRequest) *NullableProductlistUpdateProductListRequest {
	return &NullableProductlistUpdateProductListRequest{value: val, isSet: true}
}

func (v NullableProductlistUpdateProductListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistUpdateProductListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
