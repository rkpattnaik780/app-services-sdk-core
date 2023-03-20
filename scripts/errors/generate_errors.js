const fs = require("fs");
const { cwd } = require("process");
const sdks = require(cwd() + "/.errors/index.js");

console.log("===== GO SDK =====");

goSdk = sdks["go"]

for (api in goSdk) {

  apiJson = goSdk[api].definition;
  apiFileLocation = goSdk[api].file;

  if (!apiJson || !apiJson.items) {
    console.log("invalid error file detected", apiJson);
    exit(1);
  }

  stringBuffer=`
package error
// ${api} error codes 
  
const (
`
  
  apiJson.items.forEach(function(errorType) {
    stringBuffer += `  // ${errorType.reason}\n`
    stringBuffer += `  ERROR_${errorType.id} string = "${errorType.code}"`
    stringBuffer += `\n\n`
  })
  
  stringBuffer += `)`

  fs.writeFileSync(cwd() + "/" + apiFileLocation, stringBuffer, { encoding: "utf8" });
  console.log(`Generated ${api}`);
}

console.log("===== JS SDK =====");

jsSdk = sdks["js"]

for (api in jsSdk) {
  apiJson = jsSdk[api].definition;
  apiFileLocation = jsSdk[api].file;

  if (!apiJson || !apiJson.items) {
    console.log("invalid error file detected", apiJson);
    exit(1);
  }

  stringBuffer = `
/**
   ${api} error codes
    \`\`\`ts
    apiCall.then((data) => {
        console.log(data?.data.items)
    }).catch((err) => {
      if(APIErrorCodes.ERROR_5 == err.response?.data.code) {
        // Handle error
      }
    })
    \`\`\`
*/
export const APIErrorCodes = {
`;

  apiJson.items.forEach(function (errorType) {
    stringBuffer += `  /** ${errorType.reason}*/\n`;
    stringBuffer += `  ERROR_${errorType.id} : "${errorType.code}", \n\n`;
  });

  stringBuffer += `}`;

  fs.writeFileSync(
    cwd() + "/" + apiFileLocation, stringBuffer, { encoding: "utf8" });
  console.log(`Generated ${api}`);
}

console.log("===== PYTHON SDK =====");

pythonSdk = sdks["python"]

for (api in pythonSdk) {
  apiJson = pythonSdk[api].definition;
  apiFileLocation = pythonSdk[api].file;

  if (!apiJson || !apiJson.items) {
    console.log("invalid error file detected", apiJson);
    exit(1);
  }

  stringBuffer = 
`from enum import Enum 
#   ${api} error codes
class APIErrorCodes(Enum):
`;

  apiJson.items.forEach(function (errorType) {
    stringBuffer += `  # ${errorType.reason} \n`;
    stringBuffer += `  ERROR_${errorType.id} = "${errorType.code}" \n\n`;
  });


  fs.writeFileSync(
    cwd() + "/" + apiFileLocation, stringBuffer, { encoding: "utf8" });
  console.log(`Generated ${api}`);
}
