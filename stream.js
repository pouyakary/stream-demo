
const BYTE_ORDER_OF_MAGNITUDE = 1024;
const READ_SIZE = 10 * BYTE_ORDER_OF_MAGNITUDE;
const PATH_TO_FILE = __dirname + "/patrick.jpg"
const fs = require('fs');

async function * readChunksOfTheFileAsBASE64 () {
    const readable = fs.createReadStream(PATH_TO_FILE, { highWaterMark: READ_SIZE });
    for await (const chunk of readable) {
        yield chunk.toString("base64");
    }
}

async function savePartsInTheCasesFile () {
    let fileContent = ""
    for await (const part of readChunksOfTheFileAsBASE64()) {
        fileContent += part + "\n"
    }
    fs.writeFileSync( __dirname + "/parts.txt", fileContent.trim(), "utf8")
}

savePartsInTheCasesFile()
