/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.16.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { SupportedKafkaSizeBytesValueItem } from './supported-kafka-size-bytes-value-item';

/**
 * 
 * @export
 * @interface KafkaRequestAllOf
 */
export interface KafkaRequestAllOf {
    /**
     * Values: [accepted, preparing, provisioning, ready, failed, deprovision, deleting, suspending, suspended, resuming] 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'status'?: string;
    /**
     * Name of Cloud used to deploy. For example AWS
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'cloud_provider'?: string;
    /**
     * 
     * @type {boolean}
     * @memberof KafkaRequestAllOf
     */
    'multi_az': boolean;
    /**
     * Values will be regions of specific cloud provider. For example: us-east-1 for AWS
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'region'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'owner'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'name'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'bootstrap_server_host'?: string;
    /**
     * The kafka admin server url to perform kafka admin operations e.g acl management etc. The value will be available when the Kafka has been fully provisioned i.e it reaches a \'ready\' state
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'admin_api_server_url'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'created_at'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'expires_at'?: string | null;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'updated_at'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'failed_reason'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'version'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'instance_type'?: string;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {string}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'instance_type_name'?: string;
    /**
     * 
     * @type {boolean}
     * @memberof KafkaRequestAllOf
     */
    'reauthentication_enabled': boolean;
    /**
     * 
     * @type {SupportedKafkaSizeBytesValueItem}
     * @memberof KafkaRequestAllOf
     */
    'max_data_retention_size'?: SupportedKafkaSizeBytesValueItem;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'browser_url'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'size_id'?: string;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {string}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'ingress_throughput_per_sec'?: string;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {string}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'egress_throughput_per_sec'?: string;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {number}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'total_max_connections'?: number;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {number}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'max_partitions'?: number;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {string}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'max_data_retention_period'?: string;
    /**
     * This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead.
     * @type {number}
     * @memberof KafkaRequestAllOf
     * @deprecated
     */
    'max_connection_attempts_per_sec'?: number;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'billing_cloud_account_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'marketplace'?: string;
    /**
     * 
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'billing_model'?: string;
    /**
     * Status of the Kafka request promotion. Possible values: [\'promoting\', \'failed\']. If unset it means no promotion is in progress.
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'promotion_status'?: string;
    /**
     * The ID of the data plane where Kafka is deployed on. This information is only returned for kafka whose billing model is enterprise
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'cluster_id'?: string | null;
    /**
     * Details of the Kafka request promotion. It can be set when a Kafka request promotion is in progress or has failed
     * @type {string}
     * @memberof KafkaRequestAllOf
     */
    'promotion_details'?: string;
}
