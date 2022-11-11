const fs = require("fs");
const { cwd } = require("process");
const sdks = require(cwd() + "/.errors/index.js");

console.log("Generating errors sdks");

// will get extended t include more sdks, but process should be the same
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
    console.log(`Generated ${api} error definitions`);
}