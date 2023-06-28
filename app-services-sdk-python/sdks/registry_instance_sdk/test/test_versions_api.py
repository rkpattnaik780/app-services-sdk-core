"""
    Apicurio Registry API [v2]

    Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`.   # noqa: E501

    The version of the OpenAPI document: 2.4.x
    Contact: apicurio@lists.jboss.org
    Generated by: https://openapi-generator.tech
"""


import unittest

import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api.versions_api import VersionsApi  # noqa: E501


class TestVersionsApi(unittest.TestCase):
    """VersionsApi unit test stubs"""

    def setUp(self):
        self.api = VersionsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_artifact_version_comment(self):
        """Test case for add_artifact_version_comment

        Add new comment  # noqa: E501
        """
        pass

    def test_create_artifact_version(self):
        """Test case for create_artifact_version

        Create artifact version  # noqa: E501
        """
        pass

    def test_delete_artifact_version(self):
        """Test case for delete_artifact_version

        Delete artifact version  # noqa: E501
        """
        pass

    def test_delete_artifact_version_comment(self):
        """Test case for delete_artifact_version_comment

        Delete a single comment  # noqa: E501
        """
        pass

    def test_get_artifact_version(self):
        """Test case for get_artifact_version

        Get artifact version  # noqa: E501
        """
        pass

    def test_get_artifact_version_comments(self):
        """Test case for get_artifact_version_comments

        Get artifact version comments  # noqa: E501
        """
        pass

    def test_get_artifact_version_references(self):
        """Test case for get_artifact_version_references

        Get artifact version references  # noqa: E501
        """
        pass

    def test_list_artifact_versions(self):
        """Test case for list_artifact_versions

        List artifact versions  # noqa: E501
        """
        pass

    def test_update_artifact_version_comment(self):
        """Test case for update_artifact_version_comment

        Update a comment  # noqa: E501
        """
        pass

    def test_update_artifact_version_state(self):
        """Test case for update_artifact_version_state

        Update artifact version state  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
