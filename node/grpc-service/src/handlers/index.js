import { app } from "../init"
import { SayHello, SayHi } from "./matan-handlers"

export const initHandlers = () => {
  // init matan's service handlers
  app.use({ SayHello, SayHi })
}