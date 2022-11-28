import subprocess
import os

endpoints= [
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml",
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/connector_mgmt.yaml",
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/srs-fleet-manager/main/core/src/main/resources/srs-fleet-manager.json",
    "https://raw.githubusercontent.com/5733d9e2be6485d52ffa08870cabdee0/sandbox/main/openapi/smartevents_mgmt_v2.yaml",
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kafka-admin-api/main/kafka-admin/.openapi/kafka-admin-rest.yaml",
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kafka-admin-api/main/kafka-admin/.openapi/kafka-admin-rest.yaml",
    "https://sso.redhat.com/auth/realms/redhat-external/apis/openapi.yaml"
]

for url in endpoints:
    basename = os.path.basename(url)

    # these are both hacks
    if basename == "openapi.yaml":
        basename = "service-accounts.yaml"

    if basename == "smartevents_mgmt_v2.yaml":
        basename = "smartevents_mgmt.yaml"

            # get the directory the python script is in and then get relative path to output file
    output_path = f"{os.path.dirname(os.path.realpath(__file__))}/../.openapi/{basename}"

    subprocess.run([
        "curl",
        url,
        "--output",
        output_path
    ], capture_output=False)
