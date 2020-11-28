import express from "express"
import cors from "cors"
import bodyParser from "body-parser"

import fooRoutes from "./api/v1/foo"
import { checkIfAuthenticated } from "./middlewares/auth"

// Init express instance
const app = express()

// Define app's port
const port = process.env.PORT || 5000

//// Handle middlewares
// Cors settings
app.use(cors())

// Body parser to request body
app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())

// Handle authentication middleware
app.use(checkIfAuthenticated)

// Handle health check
app.get("/status", (req, res) => res.json({ status: "ok" }))

// Implement routes
app.use("/api/v1/foo", fooRoutes)

// listen to app instance
app.listen(port, () => console.log(`RESTful express server listen on port ${port}...`))