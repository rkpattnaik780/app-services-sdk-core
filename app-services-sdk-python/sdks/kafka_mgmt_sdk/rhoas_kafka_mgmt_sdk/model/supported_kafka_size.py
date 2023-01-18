"""
    Kafka Management API

    Kafka Management API is a REST API to manage Kafka instances  # noqa: E501

    The version of the OpenAPI document: 1.14.0
    Contact: rhosak-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from rhoas_kafka_mgmt_sdk.model_utils import (  # noqa: F401
    ApiTypeError,
    ModelComposed,
    ModelNormal,
    ModelSimple,
    cached_property,
    change_keys_js_to_python,
    convert_js_args_to_python_args,
    date,
    datetime,
    file_type,
    none_type,
    validate_get_composed_info,
    OpenApiModel
)
from rhoas_kafka_mgmt_sdk.exceptions import ApiAttributeError


def lazy_import():
    from rhoas_kafka_mgmt_sdk.model.supported_kafka_size_bytes_value_item import SupportedKafkaSizeBytesValueItem
    globals()['SupportedKafkaSizeBytesValueItem'] = SupportedKafkaSizeBytesValueItem


class SupportedKafkaSize(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
    }

    validations = {
    }

    @cached_property
    def additional_properties_type():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded
        """
        lazy_import()
        return (bool, date, datetime, dict, float, int, list, str, none_type,)  # noqa: E501

    _nullable = False

    @cached_property
    def openapi_types():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded

        Returns
            openapi_types (dict): The key is attribute name
                and the value is attribute type.
        """
        lazy_import()
        return {
            'id': (str,),  # noqa: E501
            'display_name': (str,),  # noqa: E501
            'ingress_throughput_per_sec': (SupportedKafkaSizeBytesValueItem,),  # noqa: E501
            'egress_throughput_per_sec': (SupportedKafkaSizeBytesValueItem,),  # noqa: E501
            'total_max_connections': (int,),  # noqa: E501
            'max_data_retention_size': (SupportedKafkaSizeBytesValueItem,),  # noqa: E501
            'max_partitions': (int,),  # noqa: E501
            'max_data_retention_period': (str,),  # noqa: E501
            'max_connection_attempts_per_sec': (int,),  # noqa: E501
            'max_message_size': (SupportedKafkaSizeBytesValueItem,),  # noqa: E501
            'min_in_sync_replicas': (int,),  # noqa: E501
            'replication_factor': (int,),  # noqa: E501
            'supported_az_modes': ([str],),  # noqa: E501
            'lifespan_seconds': (int, none_type,),  # noqa: E501
            'quota_consumed': (int,),  # noqa: E501
            'quota_type': (str,),  # noqa: E501
            'capacity_consumed': (int,),  # noqa: E501
            'maturity_status': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'id': 'id',  # noqa: E501
        'display_name': 'display_name',  # noqa: E501
        'ingress_throughput_per_sec': 'ingress_throughput_per_sec',  # noqa: E501
        'egress_throughput_per_sec': 'egress_throughput_per_sec',  # noqa: E501
        'total_max_connections': 'total_max_connections',  # noqa: E501
        'max_data_retention_size': 'max_data_retention_size',  # noqa: E501
        'max_partitions': 'max_partitions',  # noqa: E501
        'max_data_retention_period': 'max_data_retention_period',  # noqa: E501
        'max_connection_attempts_per_sec': 'max_connection_attempts_per_sec',  # noqa: E501
        'max_message_size': 'max_message_size',  # noqa: E501
        'min_in_sync_replicas': 'min_in_sync_replicas',  # noqa: E501
        'replication_factor': 'replication_factor',  # noqa: E501
        'supported_az_modes': 'supported_az_modes',  # noqa: E501
        'lifespan_seconds': 'lifespan_seconds',  # noqa: E501
        'quota_consumed': 'quota_consumed',  # noqa: E501
        'quota_type': 'quota_type',  # noqa: E501
        'capacity_consumed': 'capacity_consumed',  # noqa: E501
        'maturity_status': 'maturity_status',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """SupportedKafkaSize - a model defined in OpenAPI

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            id (str): Unique identifier of this Kafka instance size.. [optional]  # noqa: E501
            display_name (str): Display name of this Kafka instance size.. [optional]  # noqa: E501
            ingress_throughput_per_sec (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            egress_throughput_per_sec (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            total_max_connections (int): Maximum amount of total connections available to this Kafka instance size.. [optional]  # noqa: E501
            max_data_retention_size (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            max_partitions (int): Maximum amount of total partitions available to this Kafka instance size.. [optional]  # noqa: E501
            max_data_retention_period (str): Maximum data retention period available to this Kafka instance size.. [optional]  # noqa: E501
            max_connection_attempts_per_sec (int): Maximium connection attempts per second available to this Kafka instance size.. [optional]  # noqa: E501
            max_message_size (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            min_in_sync_replicas (int): Minimum number of in-sync replicas.. [optional]  # noqa: E501
            replication_factor (int): Replication factor available to this Kafka instance size.. [optional]  # noqa: E501
            supported_az_modes ([str]): List of Availability Zone modes that this Kafka instance size supports. The possible values are \"single\", \"multi\".. [optional]  # noqa: E501
            lifespan_seconds (int, none_type): The limit lifespan of the kafka instance in seconds. If not specified then the instance never expires.. [optional]  # noqa: E501
            quota_consumed (int): Quota consumed by this Kafka instance size.. [optional]  # noqa: E501
            quota_type (str): Quota type used by this Kafka instance size. This is now deprecated, please refer to supported_billing_models at instance-type level instead.. [optional]  # noqa: E501
            capacity_consumed (int): Data plane cluster capacity consumed by this Kafka instance size.. [optional]  # noqa: E501
            maturity_status (str): Maturity level of the size. Can be \"stable\" or \"preview\".. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        self = super(OpenApiModel, cls).__new__(cls)

        if args:
            raise ApiTypeError(
                "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                    args,
                    self.__class__.__name__,
                ),
                path_to_item=_path_to_item,
                valid_classes=(self.__class__,),
            )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
        return self

    required_properties = set([
        '_data_store',
        '_check_type',
        '_spec_property_naming',
        '_path_to_item',
        '_configuration',
        '_visited_composed_classes',
    ])

    @convert_js_args_to_python_args
    def __init__(self, *args, **kwargs):  # noqa: E501
        """SupportedKafkaSize - a model defined in OpenAPI

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            id (str): Unique identifier of this Kafka instance size.. [optional]  # noqa: E501
            display_name (str): Display name of this Kafka instance size.. [optional]  # noqa: E501
            ingress_throughput_per_sec (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            egress_throughput_per_sec (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            total_max_connections (int): Maximum amount of total connections available to this Kafka instance size.. [optional]  # noqa: E501
            max_data_retention_size (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            max_partitions (int): Maximum amount of total partitions available to this Kafka instance size.. [optional]  # noqa: E501
            max_data_retention_period (str): Maximum data retention period available to this Kafka instance size.. [optional]  # noqa: E501
            max_connection_attempts_per_sec (int): Maximium connection attempts per second available to this Kafka instance size.. [optional]  # noqa: E501
            max_message_size (SupportedKafkaSizeBytesValueItem): [optional]  # noqa: E501
            min_in_sync_replicas (int): Minimum number of in-sync replicas.. [optional]  # noqa: E501
            replication_factor (int): Replication factor available to this Kafka instance size.. [optional]  # noqa: E501
            supported_az_modes ([str]): List of Availability Zone modes that this Kafka instance size supports. The possible values are \"single\", \"multi\".. [optional]  # noqa: E501
            lifespan_seconds (int, none_type): The limit lifespan of the kafka instance in seconds. If not specified then the instance never expires.. [optional]  # noqa: E501
            quota_consumed (int): Quota consumed by this Kafka instance size.. [optional]  # noqa: E501
            quota_type (str): Quota type used by this Kafka instance size. This is now deprecated, please refer to supported_billing_models at instance-type level instead.. [optional]  # noqa: E501
            capacity_consumed (int): Data plane cluster capacity consumed by this Kafka instance size.. [optional]  # noqa: E501
            maturity_status (str): Maturity level of the size. Can be \"stable\" or \"preview\".. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        if args:
            raise ApiTypeError(
                "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                    args,
                    self.__class__.__name__,
                ),
                path_to_item=_path_to_item,
                valid_classes=(self.__class__,),
            )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")
