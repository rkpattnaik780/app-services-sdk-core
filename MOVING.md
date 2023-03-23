# Moving to Core SDK

## What's changed

### Go
While the code has moved the generation has stayed the same, which means updating to the Core SDK should be similar to updating to a new version in the old SDK

### Java
Java has been completly replaced with a new generator from Kiota you can learn more [here](https://microsoft.github.io/kiota/). This means any code in the previous Java SDK will no longer work and is no longer supported.

## How to update

### GO SDK

First you need to change all imports from the old SDK to now use the new Core SDK.
```bash
find . -type f -name "*.go" -not -path '*vendor*' -exec sed -i '' -e 's/github.com\/redhat-developer\/app-services-sdk-go/github.com\/redhat-developer\/app-services-sdk-core\/app-services-sdk-go/g' {} +
```

Then update the `go.mod` folder to use the new packages. 
```bash
go clean -modcache && go mod tidy
```

If you are using a vendor folder in your project then this is when you will need to re-vendor.
```bash
go mod vendor
```