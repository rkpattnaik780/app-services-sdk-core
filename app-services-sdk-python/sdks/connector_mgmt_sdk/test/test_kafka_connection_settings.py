"""
    Connector Management API

    Connector Management API is a REST API to manage connectors.  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: rhosak-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_connector_mgmt_sdk
from rhoas_connector_mgmt_sdk.model.service_connection_settings import ServiceConnectionSettings
globals()['ServiceConnectionSettings'] = ServiceConnectionSettings
from rhoas_connector_mgmt_sdk.model.kafka_connection_settings import KafkaConnectionSettings


class TestKafkaConnectionSettings(unittest.TestCase):
    """KafkaConnectionSettings unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testKafkaConnectionSettings(self):
        """Test KafkaConnectionSettings"""
        # FIXME: construct object with mandatory attributes with example values
        # model = KafkaConnectionSettings()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()