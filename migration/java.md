# Migrating a Java project

In this example project we use the kafka managment SDK to create kafkas. Each snippet is split into the old SDK and the new updated Kiota SDK.

## Migrating the pom.xml
### Old SDK
```xml
<dependencies>
    <dependency>
        <groupId>com.redhat.cloud</groupId>
        <artifactId>kafka-management-sdk</artifactId>
        <version>0.31.0</version>
        <scope>compile</scope>
    </dependency>
</dependencies>
```

### Kiota SDK
First add the sontatype remote repoistory so the project can access the required dependencies.
```xml
<repositories>
    <repository>
        <id>ossrh</id>
        <url>https://oss.sonatype.org/service/local/staging/deploy/maven2/</url>
    </repository>
</repositories>
```

Next add properties for the versions of the dependencies.
```xml
<properties>
    <sdk.version>1.0.1</sdk.version>
    <kiota.version>0.3.2</kiota.version>
    <auth.version>0.0.3</auth.version>
</properties>
```

Now add the required dependencies, this not only adds the kafka managment SDK but also the authentication dependency and the required Kiota dependencies for the SDK.
```xml
<dependencies>
        <dependency>
            <groupId>com.redhat.cloud</groupId>
            <artifactId>kafka-management-sdk</artifactId>
            <version>${sdk.version}</version>
        </dependency>
        <dependency>
            <groupId>com.microsoft.kiota</groupId>
            <artifactId>microsoft-kiota-abstractions</artifactId>
            <version>${kiota.version}</version>
        </dependency>
        <dependency>
            <groupId>com.microsoft.kiota</groupId>
            <artifactId>microsoft-kiota-serialization-json</artifactId>
            <version>${kiota.version}</version>
        </dependency>
        <dependency>
            <groupId>com.microsoft.kiota</groupId>
            <artifactId>microsoft-kiota-serialization-text</artifactId>
            <version>${kiota.version}</version>
        </dependency>
        <dependency>
            <groupId>com.microsoft.kiota</groupId>
            <artifactId>microsoft-kiota-serialization-form</artifactId>
            <version>${kiota.version}</version>
        </dependency>
        <dependency>
            <groupId>com.microsoft.kiota</groupId>
            <artifactId>microsoft-kiota-http-okHttp</artifactId>
            <version>${kiota.version}</version>
        </dependency>
        <dependency>
            <groupId>com.redhat.cloud</groupId>
            <artifactId>kiota-rh-auth</artifactId>
            <version>${auth.version}</version>
        </dependency>
    </dependencies>
```

## Migrating the code to create the API client
### Old SDK
```java
// setup the client to be used using the production API
ApiClient defaultClient = Configuration.getDefaultApiClient();
defaultClient.setBasePath("https://api.openshift.com");

// set the token to be used, this needs to be the access token
HttpBearerAuth Bearer = (HttpBearerAuth) defaultClient.getAuthentication("Bearer");
Bearer.setBearerToken(token);

DefaultApi apiInstance = new DefaultApi(defaultClient);
```

### Kiota SDK
```java
// setup the client to be used using the production API, and pass in the offline token
OkHttpRequestAdapter adapter = new OkHttpRequestAdapter(
        new BaseBearerTokenAuthenticationProvider(
                new RHAccessTokenProvider(token)
        )
);

adapter.setBaseUrl("https://api.openshift.com");
var client = new ApiClient(adapter);
```

## Migrating the code to create a Kafka
### Old SDK
```java
public static void createKafka(DefaultApi apiInstance) {
    KafkaRequestPayload payload = new KafkaRequestPayload().name("instance");

    try {
        KafkaRequest result = apiInstance.createKafka(true, payload);
    } catch (ApiException e) {
        e.printStackTrace();
    }
}
```

### Kiota SDK
```java
private static void createKafka(ApiClient client) throws Exception {
    KafkaRequestPayload payload = new KafkaRequestPayload();
    payload.setName("instance");
    var kafka = client
            .api()
            .kafkas_mgmt()
            .v1()
            .kafkas()
            .post(payload, (config) -> {
                config.queryParameters.async = true;
            })
            .get(3, TimeUnit.SECONDS);
}
```