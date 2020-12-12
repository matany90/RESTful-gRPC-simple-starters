import { app, HOST, PORT } from "./init"
import { initMiddlewares } from "./middlewares"
import { initHandlers } from "./handlers"

// Init app middlewares
initMiddlewares()

// init app handlers
initHandlers()

// start app
console.log(`gRPC Mali server listen on port ${PORT}...`)
app.start(`${HOST}:${PORT}`)