# RHOAS SDK for Java

Java packages and API clients for RHOAS services

## Prequisites

- [Maven 3.6.x](https://maven.apache.org) or above

## Usage

Add the required dependencies to your `pom.xml` file:

```xml
<repositories>
    <repository>
        <id>ossrh</id>
        <url>https://oss.sonatype.org/service/local/staging/deploy/maven2/</url>
    </repository>
</repositories>

<dependency>
  <groupId>com.github.redhat-developer.app-services-sdk-java</groupId>
  <artifactId>account-management-sdk</artifactId>
  <version>version</version>
</dependency>
<dependency>
  <groupId>com.github.redhat-developer.app-services-sdk-java</groupId>
  <artifactId>kafka-instance-sdk</artifactId>
  <version>version</version>
</dependency>
<dependency>
  <groupId>com.github.redhat-developer.app-services-sdk-java</groupId>
  <artifactId>kafka-management-sdk</artifactId>
  <version>version</version>
</dependency>
<dependency>
  <groupId>com.github.redhat-developer.app-services-sdk-java</groupId>
  <artifactId>registry-instance-sdk</artifactId>
  <version>version</version>
</dependency>
<dependency>
  <groupId>com.github.redhat-developer.app-services-sdk-java</groupId>
  <artifactId>registry-management-sdk</artifactId>
  <version>version</version>
</dependency>
<dependency>
  <groupId>com.github.redhat-developer.app-services-sdk-java</groupId>
  <artifactId>service-accounts-sdk</artifactId>
  <version>version</version>
</dependency>

```

To easily use RH SSO add also:
```xml
<dependency>
    <groupId>com.redhat.cloud</groupId>
    <artifactId>kiota-rh-auth</artifactId>
    <version>version</version>
</dependency>
```

After obtaining an offline token following the instructions [here](https://access.redhat.com/articles/3626371).
You can use the SDKs in your project like:

```java
OkHttpRequestAdapter adapter = new OkHttpRequestAdapter(new BaseBearerTokenAuthenticationProvider(new RHAccessTokenProvider(offline_token)));

// adapter.setBaseUrl("https://api.stage.openshift.com");
adapter.setBaseUrl("https://api.openshift.com");

var client = new ApiClient(adapter);

var user = client
        .api()
        .accounts_mgmt()
        .v1()
        .current_account()
        .get().get(3, TimeUnit.SECONDS);
```



## Docs

[Documentation](./docs)

## Contributing

Contributions are welcome. See [CONTRIBUTING](../CONTRIBUTING.md) for details.

[kafkagit]: https://github.com/redhat-developer/app-services-sdk-core/app-services-sdk-java/tree/main/packages/kafka-management-sdk
[sadoc]: https://github.com/redhat-developer/app-services-sdk-core/app-services-sdk-java/tree/main/packages/service-accounts-sdk
[kinstancegit]: https://github.com/redhat-developer/app-services-sdk-core/app-services-sdk-java/tree/main/packages/kafka-instance-sdk
[registrygit]: https://github.com/redhat-developer/app-services-sdk-core/app-services-sdk-java/tree/main/packages/registry-management-sdk
[connectorgit]: https://github.com/redhat-developer/app-services-sdk-core/app-services-sdk-java/tree/main/packages/connector-management-sdk
[smarteventsgit]: https://github.com/redhat-developer/app-services-sdk-core/app-services-sdk-java/tree/main/packages/smartevents-management-sdk