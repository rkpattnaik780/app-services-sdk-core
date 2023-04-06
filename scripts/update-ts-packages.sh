(cd app-services-sdk-js && npm version $VERSION)

for entry in app-services-sdk-ts/packages/*
do
    (cd "$entry" && npm version $VERSION)
done