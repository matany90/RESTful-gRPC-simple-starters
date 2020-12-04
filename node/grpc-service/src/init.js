import path from "path"
import Mali from "mali"

// Init app using proto file
const PROTO_PATH = path.resolve(__dirname, "../proto/matan.proto")
const app = new Mali(PROTO_PATH, "MatanService")

// define port and host
const HOST = process.env.HOST || "0.0.0.0"
const PORT = process.env.PORT || 50001

// exports app defenitions
export {
  app,
  HOST,
  PORT
}