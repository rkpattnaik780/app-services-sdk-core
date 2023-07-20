/* tslint:disable */
/* eslint-disable */
/**
 * Account Management Service API
 * Manage user subscriptions and clusters
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 
 * @export
 * @interface SubscriptionPatchRequest
 */
export interface SubscriptionPatchRequest {
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'billing_expiration_date'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'cloud_account_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'cloud_provider_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'cluster_billing_model'?: SubscriptionPatchRequestClusterBillingModelEnum;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'cluster_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'console_url'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'consumer_uuid'?: string;
    /**
     * 
     * @type {number}
     * @memberof SubscriptionPatchRequest
     */
    'cpu_total'?: number;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'creator_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'display_name'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'external_cluster_id'?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SubscriptionPatchRequest
     */
    'managed'?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'organization_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'plan_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'product_bundle'?: SubscriptionPatchRequestProductBundleEnum;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'provenance'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'region_id'?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SubscriptionPatchRequest
     */
    'released'?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'service_level'?: SubscriptionPatchRequestServiceLevelEnum;
    /**
     * 
     * @type {number}
     * @memberof SubscriptionPatchRequest
     */
    'socket_total'?: number;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'status'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'support_level'?: SubscriptionPatchRequestSupportLevelEnum;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'system_units'?: SubscriptionPatchRequestSystemUnitsEnum;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'trial_end_date'?: string;
    /**
     * 
     * @type {string}
     * @memberof SubscriptionPatchRequest
     */
    'usage'?: SubscriptionPatchRequestUsageEnum;
}

export const SubscriptionPatchRequestClusterBillingModelEnum = {
    Standard: 'standard',
    Marketplace: 'marketplace',
    MarketplaceAws: 'marketplace-aws',
    MarketplaceAzure: 'marketplace-azure',
    MarketplaceRhm: 'marketplace-rhm'
} as const;

export type SubscriptionPatchRequestClusterBillingModelEnum = typeof SubscriptionPatchRequestClusterBillingModelEnum[keyof typeof SubscriptionPatchRequestClusterBillingModelEnum];
export const SubscriptionPatchRequestProductBundleEnum = {
    Openshift: 'Openshift',
    JBossMiddleware: 'JBoss-Middleware',
    IbmCloudPak: 'IBM-CloudPak'
} as const;

export type SubscriptionPatchRequestProductBundleEnum = typeof SubscriptionPatchRequestProductBundleEnum[keyof typeof SubscriptionPatchRequestProductBundleEnum];
export const SubscriptionPatchRequestServiceLevelEnum = {
    L1L3: 'L1-L3',
    L3Only: 'L3-only'
} as const;

export type SubscriptionPatchRequestServiceLevelEnum = typeof SubscriptionPatchRequestServiceLevelEnum[keyof typeof SubscriptionPatchRequestServiceLevelEnum];
export const SubscriptionPatchRequestSupportLevelEnum = {
    Eval: 'Eval',
    Standard: 'Standard',
    Premium: 'Premium',
    SelfSupport: 'Self-Support',
    None: 'None'
} as const;

export type SubscriptionPatchRequestSupportLevelEnum = typeof SubscriptionPatchRequestSupportLevelEnum[keyof typeof SubscriptionPatchRequestSupportLevelEnum];
export const SubscriptionPatchRequestSystemUnitsEnum = {
    CoresVCpu: 'Cores/vCPU',
    Sockets: 'Sockets'
} as const;

export type SubscriptionPatchRequestSystemUnitsEnum = typeof SubscriptionPatchRequestSystemUnitsEnum[keyof typeof SubscriptionPatchRequestSystemUnitsEnum];
export const SubscriptionPatchRequestUsageEnum = {
    Production: 'Production',
    DevelopmentTest: 'Development/Test',
    DisasterRecovery: 'Disaster Recovery',
    Academic: 'Academic'
} as const;

export type SubscriptionPatchRequestUsageEnum = typeof SubscriptionPatchRequestUsageEnum[keyof typeof SubscriptionPatchRequestUsageEnum];

