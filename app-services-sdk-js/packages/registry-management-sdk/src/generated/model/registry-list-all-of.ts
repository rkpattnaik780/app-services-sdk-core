/* tslint:disable */
/* eslint-disable */
/**
 * Service Registry Management API
 * Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: rhosak-eval-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Registry } from './registry';

/**
 * 
 * @export
 * @interface RegistryListAllOf
 */
export interface RegistryListAllOf {
    /**
     * 
     * @type {Array<Registry>}
     * @memberof RegistryListAllOf
     */
    'items': Array<Registry>;
}

