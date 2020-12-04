import { app } from "../init"
import { checkIfAuthenticated } from "./auth"
import { validateData } from "./data-validation"

/**
 * Expose app middlewares.
 * includes checkIfAuthenticated middleware
 * and data valdiation.
 */
export const initMiddlewares = () => {
  // checkIfAutheticated middleware init
  app.use(checkIfAuthenticated)

  // validateData middleware init
  app.use(validateData)
}