/*
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.14.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.openshift.cloud.api.kas.models;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import com.openshift.cloud.api.kas.models.SupportedKafkaInstanceType;
import com.openshift.cloud.api.kas.models.SupportedKafkaInstanceTypesListAllOf;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * SupportedKafkaInstanceTypesList
 */
@JsonPropertyOrder({
  SupportedKafkaInstanceTypesList.JSON_PROPERTY_INSTANCE_TYPES
})
@JsonTypeName("SupportedKafkaInstanceTypesList")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class SupportedKafkaInstanceTypesList {
  public static final String JSON_PROPERTY_INSTANCE_TYPES = "instance_types";
  private List<SupportedKafkaInstanceType> instanceTypes = null;

  public SupportedKafkaInstanceTypesList() { 
  }

  public SupportedKafkaInstanceTypesList instanceTypes(List<SupportedKafkaInstanceType> instanceTypes) {
    
    this.instanceTypes = instanceTypes;
    return this;
  }

  public SupportedKafkaInstanceTypesList addInstanceTypesItem(SupportedKafkaInstanceType instanceTypesItem) {
    if (this.instanceTypes == null) {
      this.instanceTypes = new ArrayList<>();
    }
    this.instanceTypes.add(instanceTypesItem);
    return this;
  }

   /**
   * Get instanceTypes
   * @return instanceTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_INSTANCE_TYPES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<SupportedKafkaInstanceType> getInstanceTypes() {
    return instanceTypes;
  }


  @JsonProperty(JSON_PROPERTY_INSTANCE_TYPES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setInstanceTypes(List<SupportedKafkaInstanceType> instanceTypes) {
    this.instanceTypes = instanceTypes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SupportedKafkaInstanceTypesList supportedKafkaInstanceTypesList = (SupportedKafkaInstanceTypesList) o;
    return Objects.equals(this.instanceTypes, supportedKafkaInstanceTypesList.instanceTypes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(instanceTypes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SupportedKafkaInstanceTypesList {\n");
    sb.append("    instanceTypes: ").append(toIndentedString(instanceTypes)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
