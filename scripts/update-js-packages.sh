(cd app-services-sdk-js && npm version $VERSION)

for entry in app-services-sdk-js/packages/*
do
    (cd "$entry" && npm version $VERSION)
done