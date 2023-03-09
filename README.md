# RHOAS Core SDKs

| Supported Langauge  | SDK |
| ----------- | ----------- |
| JavaScript/TypeScript  | [app-services-sdk-js](/app-services-sdk-js/) |
| Python   | [app-services-sdk-python](/app-services-sdk-python/)    |
| Java   | [app-services-sdk-java](/app-services-sdk-java/)        |
|  Go   | [app-services-sdk-go](/app-services-sdk-go/)        |

## Purpose of this repository

Repository contains source code for openapi based generator along with automation scripts that generate all
underlying RHOAS SDKS. 

# RHOAS SDK for Typescript and JavaScript

Typescript and JavaScript packages and API clients for RHOAS services

## Prequisites

- [NodeJS 14.x lts](https://nodejs.org/en/about/releases/) or above

## Introduction

RHOAS TypeScript SDK's are delivered as set of the individual npm packages.
SDK's are divided into two types:

- **Management SDKs** - support creating services
- **Instance SDKs** -  support interaction with services (limited to capabilities not supported by upstream tools)

Typically developers need to integrate with Management SDK's to create or manage service instances 
and later use upstream libraries or management SDK's to interact with services.

You can view all SDK packages directly in [npm](https://www.npmjs.com/search?q=keywords:rhoas)

## Management SDK's

| Name             | Package name                        | Status | Version            | NPM                 | Docs                |
| ---------------- | :---------------------------------- | :----- | ------------------ | ------------------- | ------------------- |
| kafka            | **@rhoas/kafka-management-sdk**     | beta   | ![kafkashield]     | [npm][kafkanpm]     | [doc][kafkagit]     |
| service-accounts            | **@rhoas/service-accounts-sdk**     | alpha   | ![sashield]     | [npm][sanpm]     | [doc][sagit]     
| connector        | **@rhoas/connector-management-sdk** | alpha  | ![connectorshield] | [npm][connectornpm] | [doc][connectorgit] |
| service registry | **@rhoas/registry-management-sdk**  | alpha  | ![registryshield]  | [npm][registrynpm]  | [doc][registrygit]  |
| smart events | **@rhoas/smart-events-management-sdk**  | alpha  | ![smarteventsshield]| [npm][smarteventsnpm]  | [doc][smarteventsgit] |


## Instances SDK's

| Name             | Package name                     | Status | Version             | NPM                 | Docs                |
| ---------------- | :------------------------------- | :----- | ------------------- | ------------------- | ------------------- |
| kafka            | **@rhoas/kafka-instance-sdk**    | beta   | ![kinstanceshield]  | [npm][kinstancenpm] | [doc][kinstancegit] |
| service registry | **@rhoas/registry-instance-sdk** | beta   | ![srinstanceshield] | [npm][srinstancenpm]| [doc][srinstancegit]|

## Docs

[Documentation](/app-services-sdk-js/docs) 

## Examples

[Examples](/app-services-sdk-js/examples) 

# RHOAS SDK for Python

Python packages and API clients for Red Had OpenShift Application Services (RHOAS) 

## Prequisites

- [Python 3.9](https://docs.python.org/3/) or above
- [pip](https://pypi.org/project/pip/) for installing packages

## Installation

Currently all RHOAS SDKs are bundled together. To install the RHOAS SDK with the pip package installer:

```shell
$ python -m pip install rhoas-sdks
```

## RHOAS App Services SDK for Python

> NOTE: Some of these APIs are under development and may sometimes cause backwards-incompatible changes.

All packages are now available and can be accessed by just importing them as shown below:


| API                       | Status | Package                                                                                                                                                         |
| :------------------------ | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [KafkaManagement](/app-services-sdk-python/sdks/kafka_mgmt_sdk/README.md)           | beta   | `import python rhoas_kafka_mgmt_sdk`          |
| [ServiceRegistryManagement](app-services-sdk-python/sdks/registry_mgmt_sdk/README.md)  | alpha   | `import rhoas_service_registry_mgmt_sdk`         |
| [ConnectorManagement](app-services-sdk-python/sdks/connector_mgmt_sdk/README.md)       | alpha  | `import rhoas_connector_mgmt_sdk`  |
| [ServiceAccounts](app-services-sdk-python/sdks/service_accounts_mgmt_sdk/README.md) | alpha | `import rhoas_service_accounts_mgmt_sdk` |

 
 ## Instances SDKs

| API              | Status | Package                                                                                                                                                                               |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [KafkaInstance](app-services-sdk-python/sdks/kafka_instance_sdk/README.md)    | beta   | `import rhoas_kafka_instance_sdk`|
| [RegistryInstance](app-services-sdk-python/sdks/registry_instance_sdk/README.md) | beta   | `import rhoas_registry_instance_sdk` |


## Documentation
[Documentation](/app-services-sdk-python/docs)
## Examples
[Examples](/app-services-sdk-python/examples)

# RHOAS SDK for Java

Java packages and API clients for RHOAS services

## Prequisites

- [NodeJS 14.x lts](https://nodejs.org/en/about/releases/) or above
- [Maven 3.6.x](https://maven.apache.org) or above
- [YQ 4.6.x](https://github.com/mikefarah/yq#install) or above

## Introduction

RHOAS Java SDK's are delivered as set of the individual maven packages.
SDK's are divided into two types:

- **Management SDKs** - support creating services
- **Instance SDKs** -  support interaction with services (limited to capabilities not supported by upstream tools)

Typically developers need to integrate with Management SDK's to create or manage service instances
and later use upstream libraries or management SDK's to interact with services.

You can view all SDK packages directly in [GitHub](https://github.com/orgs/redhat-developer/packages)

## Management SDK's

| Name             | Artifact Id                        | Status | Version           | Package            | Docs               |
| ---------------- | :--------------------------------- | :----- | ----------------- | ------------------ | ------------------ |
| kafka            | **kafka-management-sdk**           | beta   | 1.1.2             | [info](https://search.maven.org/search?q=a:kafka-management-sdk)                | [doc][kafkagit]    |
| service accounts            | **service-accounts-sdk**           | alpha   | 1.1.2             | [info](https://search.maven.org/search?q=a:service-accounts-sdk)                | [doc][sadoc]    |
| service registry | **registry-management-sdk**        | alpha  | 0.0.6             | [info](https://search.maven.org/search?q=a:registry-management-sdk)                | [doc][smarteventsgit] |
|  connector-management | **connector-management-sdk**        | alpha  | 0.0.6             | [info](https://search.maven.org/search?q=a:connector-management-sdk)                | [doc][connectorgit] |
|  smart-events-management | **smart-events-management-sdk**        | alpha  | 0.0.6             | [info](https://search.maven.org/search?q=a:smartevents-management-sdk)                | [doc][smarteventsgit] |

## Instances SDK's

| Name             | Package name                  | Status | Version            |  Package            | Docs                |
| ---------------- | :---------------------------- | :----- | ------------------ | ------------------- | ------------------- |
| kafka            | **kafka-instance-sdk**        | beta   | 1.1.2              | [info](https://search.maven.org/search?q=a:kafka-instance-sdk)                 | [doc][kinstancegit] |
| service registry | TBA                           | TBA    | TBA                | TBA                 | TBA                 |

## Docs

[Documentation](/app-services-sdk-java/docs)

# RHOAS SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/redhat-developer/app-services-sdk-go.svg)](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go)

Go packages and API clients for Red Hat OpenShift Application Services (RHOAS) 

## Prequisites

- [Go 1.15](https://golang.org/doc/go1.15) or above

## Installation

Install the RHOAS SDK with `go get`:

```shell
$ go get github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go
```

Import:

```go
import "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go"
```

## Management SDKs

> NOTE: Some of these APIs are under development and may sometimes cause backwards-incompatible changes.


| API                       | Status | Package                                                                                                                                                         |
| :------------------------ | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| KafkaManagement           | beta   | [`github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1)         |
| Service Account Management  | alpha   | [`github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1)         |
| ServiceRegistryManagement | alpha  | [`github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1)   |
| ConnectorManagement       | alpha  | [`github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1) |

 
 ## Instances SDKs

| API              | Status | Package                                                                                                                                                                               |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| KafkaInstance    | beta   | [`github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkainstance/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkainstance/apiv1internal)       |
| RegistryInstance | beta   | [`github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal) |


## Documentation

[Documentation](/app-services-sdk-go/docs)

## Examples

[Examples](/app-services-sdk-go/examples)

[//]: # (Java SDK variables)
[kafkagit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-java/packages/kafka-management-sdk
[sadoc]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-java/packages/service-accounts-sdk
[kinstancegit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-java/packages/kafka-instance-sdk
[registrygit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-java/packages/registry-management-sdk
[connectorgit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-java/packages/connector-management-sdk
[smarteventsgit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-java/packages/smartevents-management-sdk


[//]: # (JS SDK variables)
[kafkagit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/kafka-management-sdk 
[kafkanpm]: https://www.npmjs.com/package/@rhoas/kafka-management-sdk
[kafkashield]: https://img.shields.io/npm/v/@rhoas/kafka-management-sdk
[sagit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/service-accounts-sdk 
[sanpm]: https://www.npmjs.com/package/@rhoas/service-accounts-sdk
[sashield]: https://img.shields.io/npm/v/@rhoas/service-accounts-sdk
[kinstancegit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/kafka-instance-sdk 
[kinstancenpm]: https://www.npmjs.com/package/@rhoas/kafka-instance-sdk
[kinstanceshield]: https://img.shields.io/npm/v/@rhoas/kafka-instance-sdk
[srinstancegit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/registry-instance-sdk 
[srinstancenpm]: https://www.npmjs.com/package/@rhoas/registry-instance-sdk
[srinstanceshield]: https://img.shields.io/npm/v/@rhoas/registry-instance-sdk
[registrygit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/registry-management-sdk 
[registrynpm]: https://www.npmjs.com/package/@rhoas/registry-management-sdk
[registryshield]: https://img.shields.io/npm/v/@rhoas/registry-management-sdk
[connectorgit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/connector-management-sdk 
[connectornpm]: https://www.npmjs.com/package/@rhoas/connector-management-sdk
[connectorshield]: https://img.shields.io/npm/v/@rhoas/connector-management-sdk
[smarteventsgit]: https://github.com/redhat-developer/app-services-sdk-core/tree/main/app-services-sdk-js/packages/smart-events-management-sdk 
[smarteventsnpm]: https://www.npmjs.com/package/@rhoas/smart-events-management-sdk
[smarteventsshield]: https://img.shields.io/npm/v/@rhoas/smart-events-management-sdk

