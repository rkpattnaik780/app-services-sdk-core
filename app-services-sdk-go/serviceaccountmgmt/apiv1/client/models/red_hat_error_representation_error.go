package models
import (
    "errors"
)
// 
type RedHatErrorRepresentation_error int

const (
    SERVICE_ACCOUNT_LIMIT_EXCEEDED_REDHATERRORREPRESENTATION_ERROR RedHatErrorRepresentation_error = iota
    SERVICE_ACCOUNT_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
    SERVICE_ACCOUNT_USER_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
    SERVICE_ACCOUNT_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
    ACS_TENANT_LIMIT_EXCEEDED_REDHATERRORREPRESENTATION_ERROR
    ACS_TENANT_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
    ACS_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
    ACS_INVALID_REDIRECT_URI_REDHATERRORREPRESENTATION_ERROR
    ACS_INVALID_CLIENT_REDHATERRORREPRESENTATION_ERROR
    ACS_DISABLED_REDHATERRORREPRESENTATION_ERROR
    SMOKETEST_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
    SMOKETEST_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
    GENERAL_FAILURE_REDHATERRORREPRESENTATION_ERROR
    ORGANIZATION_API_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
)

func (i RedHatErrorRepresentation_error) String() string {
    return []string{"service_account_limit_exceeded", "service_account_not_found", "service_account_user_not_found", "service_account_access_invalid", "acs_tenant_limit_exceeded", "acs_tenant_not_found", "acs_access_invalid", "acs_invalid_redirect_uri", "acs_invalid_client", "acs_disabled", "smoketest_access_invalid", "smoketest_not_found", "general_failure", "organization_api_access_invalid"}[i]
}
func ParseRedHatErrorRepresentation_error(v string) (any, error) {
    result := SERVICE_ACCOUNT_LIMIT_EXCEEDED_REDHATERRORREPRESENTATION_ERROR
    switch v {
        case "service_account_limit_exceeded":
            result = SERVICE_ACCOUNT_LIMIT_EXCEEDED_REDHATERRORREPRESENTATION_ERROR
        case "service_account_not_found":
            result = SERVICE_ACCOUNT_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
        case "service_account_user_not_found":
            result = SERVICE_ACCOUNT_USER_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
        case "service_account_access_invalid":
            result = SERVICE_ACCOUNT_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
        case "acs_tenant_limit_exceeded":
            result = ACS_TENANT_LIMIT_EXCEEDED_REDHATERRORREPRESENTATION_ERROR
        case "acs_tenant_not_found":
            result = ACS_TENANT_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
        case "acs_access_invalid":
            result = ACS_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
        case "acs_invalid_redirect_uri":
            result = ACS_INVALID_REDIRECT_URI_REDHATERRORREPRESENTATION_ERROR
        case "acs_invalid_client":
            result = ACS_INVALID_CLIENT_REDHATERRORREPRESENTATION_ERROR
        case "acs_disabled":
            result = ACS_DISABLED_REDHATERRORREPRESENTATION_ERROR
        case "smoketest_access_invalid":
            result = SMOKETEST_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
        case "smoketest_not_found":
            result = SMOKETEST_NOT_FOUND_REDHATERRORREPRESENTATION_ERROR
        case "general_failure":
            result = GENERAL_FAILURE_REDHATERRORREPRESENTATION_ERROR
        case "organization_api_access_invalid":
            result = ORGANIZATION_API_ACCESS_INVALID_REDHATERRORREPRESENTATION_ERROR
        default:
            return 0, errors.New("Unknown RedHatErrorRepresentation_error value: " + v)
    }
    return &result, nil
}
func SerializeRedHatErrorRepresentation_error(values []RedHatErrorRepresentation_error) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
