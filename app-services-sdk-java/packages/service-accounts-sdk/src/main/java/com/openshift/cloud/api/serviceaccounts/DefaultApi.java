package com.openshift.cloud.api.serviceaccounts;

import com.openshift.cloud.api.serviceaccounts.invoker.ApiException;
import com.openshift.cloud.api.serviceaccounts.invoker.ApiClient;
import com.openshift.cloud.api.serviceaccounts.invoker.Configuration;
import com.openshift.cloud.api.serviceaccounts.invoker.Pair;

import javax.ws.rs.core.GenericType;

import com.openshift.cloud.api.serviceaccounts.models.AuthenticationPolicy;
import com.openshift.cloud.api.serviceaccounts.models.Error;
import com.openshift.cloud.api.serviceaccounts.models.SSOHealthResult;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class DefaultApi {
  private ApiClient apiClient;

  public DefaultApi() {
    this(Configuration.getDefaultApiClient());
  }

  public DefaultApi(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Get current authentication policy information
   * 
   * @param id  (required)
   * @return a {@code AuthenticationPolicy}
   * @throws ApiException if fails to make API call
   */
  public AuthenticationPolicy getAuthenticationPolicy(String id) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'id' is set
    if (id == null) {
      throw new ApiException(400, "Missing the required parameter 'id' when calling getAuthenticationPolicy");
    }
    
    // create path and map variables
    String localVarPath = "/apis/organizations/v1/{id}/authentication-policy".replaceAll("\\{format\\}","json")
      .replaceAll("\\{" + "id" + "\\}", apiClient.escapeString(id.toString()));

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

    String[] localVarAuthNames = new String[] { "authFlow", "bearerAuth", "serviceAccounts" };

    GenericType<AuthenticationPolicy> localVarReturnType = new GenericType<AuthenticationPolicy>() {};
    return apiClient.invokeAPI(localVarPath, "GET", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * 
   * 
   * @param checkName  (required)
   * @return a {@code SSOHealthResult}
   * @throws ApiException if fails to make API call
   */
  public SSOHealthResult getSmoketestByName(String checkName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'checkName' is set
    if (checkName == null) {
      throw new ApiException(400, "Missing the required parameter 'checkName' when calling getSmoketestByName");
    }
    
    // create path and map variables
    String localVarPath = "/apis/smoketest/v1/smoketests/{checkName}".replaceAll("\\{format\\}","json")
      .replaceAll("\\{" + "checkName" + "\\}", apiClient.escapeString(checkName.toString()));

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

    String[] localVarAuthNames = new String[] { "authFlow", "bearerAuth", "serviceAccounts" };

    GenericType<SSOHealthResult> localVarReturnType = new GenericType<SSOHealthResult>() {};
    return apiClient.invokeAPI(localVarPath, "GET", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Update current authentication policy information
   * 
   * @param id  (required)
   * @param authenticationPolicy  (optional)
   * @return a {@code AuthenticationPolicy}
   * @throws ApiException if fails to make API call
   */
  public AuthenticationPolicy setAuthenticationPolicy(String id, AuthenticationPolicy authenticationPolicy) throws ApiException {
    Object localVarPostBody = authenticationPolicy;
    
    // verify the required parameter 'id' is set
    if (id == null) {
      throw new ApiException(400, "Missing the required parameter 'id' when calling setAuthenticationPolicy");
    }
    
    // create path and map variables
    String localVarPath = "/apis/organizations/v1/{id}/authentication-policy".replaceAll("\\{format\\}","json")
      .replaceAll("\\{" + "id" + "\\}", apiClient.escapeString(id.toString()));

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

    String[] localVarAuthNames = new String[] { "authFlow", "bearerAuth", "serviceAccounts" };

    GenericType<AuthenticationPolicy> localVarReturnType = new GenericType<AuthenticationPolicy>() {};
    return apiClient.invokeAPI(localVarPath, "POST", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
}
