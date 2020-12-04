import { generateError } from "../utils"

/**
 * validateData will determine if the request.body
 * is valid, according to the rules and the request.body
 * that given
 *
 * @param {Object} ctx - the request context object
 * @param {Function} next - call next operation function
 */
export const validateData = async (ctx, next) => {
  try {
    console.log("Checks if request data is valid...")
    await next()
  } catch (e) {
    ctx.res = generateError(e, 400)
  }
}