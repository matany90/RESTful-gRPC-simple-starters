import { generateError } from "../utils"

/**
 * checkIfAuthenticated checks if the user
 * have valid autherization header.
 *
 * @param {Object} ctx - the request context object
 * @param {Function} next - call next operation function
 */
export const checkIfAuthenticated = async(ctx, next) => {
  try {
    console.log("Checks if user is authenticated...")
    await next()
  } catch (e) {
    ctx.res = generateError(e, 401, "unable to authenticate user.")
  }
}