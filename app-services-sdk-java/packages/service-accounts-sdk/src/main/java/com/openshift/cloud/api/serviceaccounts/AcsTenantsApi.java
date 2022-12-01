package com.openshift.cloud.api.serviceaccounts;

import com.openshift.cloud.api.serviceaccounts.invoker.ApiException;
import com.openshift.cloud.api.serviceaccounts.invoker.ApiClient;
import com.openshift.cloud.api.serviceaccounts.invoker.Configuration;
import com.openshift.cloud.api.serviceaccounts.invoker.Pair;

import javax.ws.rs.core.GenericType;

import com.openshift.cloud.api.serviceaccounts.models.AcsClientRequestData;
import com.openshift.cloud.api.serviceaccounts.models.AcsClientResponseData;
import com.openshift.cloud.api.serviceaccounts.models.Error;
import com.openshift.cloud.api.serviceaccounts.models.RedHatErrorRepresentation;
import com.openshift.cloud.api.serviceaccounts.models.ValidationExceptionData;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class AcsTenantsApi {
  private ApiClient apiClient;

  public AcsTenantsApi() {
    this(Configuration.getDefaultApiClient());
  }

  public AcsTenantsApi(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Create ACS managed central client
   * Create an ACS managed central client. Created ACS managed central clients are associated with the supplied organization id.
   * @param acsClientRequestData The name, redirect URIs and the organization id of the ACS managed central client (required)
   * @return a {@code AcsClientResponseData}
   * @throws ApiException if fails to make API call
   */
  public AcsClientResponseData createAcsClient(AcsClientRequestData acsClientRequestData) throws ApiException {
    Object localVarPostBody = acsClientRequestData;
    
    // verify the required parameter 'acsClientRequestData' is set
    if (acsClientRequestData == null) {
      throw new ApiException(400, "Missing the required parameter 'acsClientRequestData' when calling createAcsClient");
    }
    
    // create path and map variables
    String localVarPath = "/apis/beta/acs/v1".replaceAll("\\{format\\}","json");

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

    String[] localVarAuthNames = new String[] { "serviceAccounts" };

    GenericType<AcsClientResponseData> localVarReturnType = new GenericType<AcsClientResponseData>() {};
    return apiClient.invokeAPI(localVarPath, "POST", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, localVarReturnType);
      }
  /**
   * Delete ACS managed central client
   * Delete ACS managed central client by clientId. Throws not found exception if the client is not found
   * @param clientId  (required)
   * @throws ApiException if fails to make API call
   */
  public void deleteAcsClient(String clientId) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'clientId' is set
    if (clientId == null) {
      throw new ApiException(400, "Missing the required parameter 'clientId' when calling deleteAcsClient");
    }
    
    // create path and map variables
    String localVarPath = "/apis/beta/acs/v1/{clientId}".replaceAll("\\{format\\}","json")
      .replaceAll("\\{" + "clientId" + "\\}", apiClient.escapeString(clientId.toString()));

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

    String[] localVarAuthNames = new String[] { "serviceAccounts" };


    apiClient.invokeAPI(localVarPath, "DELETE", localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAccept, localVarContentType, localVarAuthNames, null);
  }
}
