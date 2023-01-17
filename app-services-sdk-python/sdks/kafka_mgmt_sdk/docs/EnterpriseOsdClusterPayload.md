# EnterpriseOsdClusterPayload

Schema for the request body sent to /clusters POST

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_kafkas_via_private_network** | **bool** | Sets whether Kafkas created on this data plane cluster have to be accessed via private network | 
**cluster_id** | **str** | The data plane cluster ID. This is the ID of the cluster obtained from OpenShift Cluster Manager (OCM) API | 
**cluster_external_id** | **str** | external cluster ID. Can be obtained from the response JSON of OCM get /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt; | 
**cluster_ingress_dns_name** | **str** | dns name of the cluster. Can be obtained from the response JSON of the /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt;/ingresses (dns_name) | 
**kafka_machine_pool_node_count** | **int** | The node count given to the created kafka machine pool.  The machine pool must be created via /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt;/machine_pools prior to passing this value. The created machine pool must have a &#x60;bf2.org/kafkaInstanceProfileType&#x3D;standard&#x60; label and a &#x60;bf2.org/kafkaInstanceProfileType&#x3D;standard:NoExecute&#x60; taint. The name of the machine pool must be &#x60;kafka-standard&#x60;  The node count value has to be a multiple of 3 with a minimum of 3 nodes. | 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


