import { generateError } from "../utils"

/**
 * SayHello handler.
 *
 * @param {Object} ctx - the request context
 */
export const SayHello = (ctx) => {
  try {
    ctx.res = { message: `Say hello to ${ctx.req.name}`}
  } catch (e) {
    ctx.res = generateError(e, 500)
  }
}

/**
 * SayHi handler.
 *
 * @param {Object} ctx - the request context
 */
export const SayHi = (ctx) => {
  try {
    ctx.res = { message: `Say hi to ${ctx.req.name}`}
  } catch (e) {
    ctx.res = generateError(e, 500)
  }
}