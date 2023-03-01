package com.openshift.cloud.api.registry.instance;

import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.Pair;

import javax.ws.rs.core.GenericType;

import com.openshift.cloud.api.registry.instance.models.CreateGroupMetaData;
import com.openshift.cloud.api.registry.instance.models.Error;
import com.openshift.cloud.api.registry.instance.models.GroupMetaData;
import com.openshift.cloud.api.registry.instance.models.GroupSearchResults;
import com.openshift.cloud.api.registry.instance.models.SortBy;
import com.openshift.cloud.api.registry.instance.models.SortOrder;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class GroupsApi {
  private ApiClient apiClient;

  public GroupsApi() {
    this(Configuration.getDefaultApiClient());
  }

  public GroupsApi(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Create a new group
   * Creates a new group.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) * The group already exist (HTTP error &#x60;409&#x60;) 
   * @param createGroupMetaData  (required)
   * @return a {@code GroupMetaData}
   * @throws ApiException if fails to make API call
   */
  public GroupMetaData createGroup(CreateGroupMetaData createGroupMetaData) throws ApiException {
    Object localVarPostBody = createGroupMetaData;
    
    // verify the required parameter 'createGroupMetaData' is set
    if (createGroupMetaData == null) {
      throw new ApiException(400, "Missing the required parameter 'createGroupMetaData' when calling createGroup");
    }
    
    // create path and map variables
    String localVarPath = "/groups".replaceAll("\\{format\\}","json");

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      "application/json"
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] {  };

    GenericType<GroupMetaData> localVarReturnType = new GenericType<GroupMetaData>() {};
    return apiClient.invokeAPI(localVarPath, "POST", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Delete a group by the specified ID.
   * Deletes a group by identifier.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) * The group does not exist (HTTP error &#x60;404&#x60;) 
   * @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. (required)
   * @throws ApiException if fails to make API call
   */
  public void deleteGroupById(String groupId) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'groupId' is set
    if (groupId == null) {
      throw new ApiException(400, "Missing the required parameter 'groupId' when calling deleteGroupById");
    }
    
    // create path and map variables
    String localVarPath = "/groups/{groupId}".replaceAll("\\{format\\}","json")
      .replaceAll("\\{" + "groupId" + "\\}", apiClient.escapeString(groupId.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] {  };


    apiClient.invokeAPI(localVarPath, "DELETE", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, null);
  }
  /**
   * Get a group by the specified ID.
   * Returns a group using the specified id.  This operation can fail for the following reasons:  * No group exists with the specified ID (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;)
   * @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. (required)
   * @return a {@code GroupMetaData}
   * @throws ApiException if fails to make API call
   */
  public GroupMetaData getGroupById(String groupId) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'groupId' is set
    if (groupId == null) {
      throw new ApiException(400, "Missing the required parameter 'groupId' when calling getGroupById");
    }
    
    // create path and map variables
    String localVarPath = "/groups/{groupId}".replaceAll("\\{format\\}","json")
      .replaceAll("\\{" + "groupId" + "\\}", apiClient.escapeString(groupId.toString()));

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();


    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] {  };

    GenericType<GroupMetaData> localVarReturnType = new GenericType<GroupMetaData>() {};
    return apiClient.invokeAPI(localVarPath, "GET", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * List groups
   * Returns a list of all groups.  This list is paged.
   * @param limit The number of groups to return.  Defaults to 20. (optional)
   * @param offset The number of groups to skip before starting the result set.  Defaults to 0. (optional)
   * @param order Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). (optional)
   * @param orderby The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  (optional)
   * @return a {@code GroupSearchResults}
   * @throws ApiException if fails to make API call
   */
  public GroupSearchResults listGroups(Integer limit, Integer offset, SortOrder order, SortBy orderby) throws ApiException {
    Object localVarPostBody = null;
    
    // create path and map variables
    String localVarPath = "/groups".replaceAll("\\{format\\}","json");

    // query params
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    localVarQueryParams.addAll(apiClient.parameterToPairs("", "limit", limit));
    localVarQueryParams.addAll(apiClient.parameterToPairs("", "offset", offset));
    localVarQueryParams.addAll(apiClient.parameterToPairs("", "order", order));
    localVarQueryParams.addAll(apiClient.parameterToPairs("", "orderby", orderby));

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] {  };

    GenericType<GroupSearchResults> localVarReturnType = new GenericType<GroupSearchResults>() {};
    return apiClient.invokeAPI(localVarPath, "GET", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
}
