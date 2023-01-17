/* tslint:disable */
/* eslint-disable */
/**
 * Service Registry API
 * Service Registry Instance API  NOTE: This API cannot be called directly from the portal.
 *
 * The version of the OpenAPI document: 2.2.5.Final
 * Contact: apicurio@lists.jboss.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { RuleViolationCause } from './rule-violation-cause';

/**
 * 
 * @export
 * @interface RuleViolationErrorAllOf
 */
export interface RuleViolationErrorAllOf {
    /**
     * List of rule violation causes.
     * @type {Array<RuleViolationCause>}
     * @memberof RuleViolationErrorAllOf
     */
    'causes': Array<RuleViolationCause>;
}
