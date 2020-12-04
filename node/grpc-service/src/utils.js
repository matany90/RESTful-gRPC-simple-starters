/**
 * generateError returns an error
 * object for handler function.
 *
 * @param {Error} error - the error thrown by the hanlder
 * @param {Number} statusCode - error status code
 * @param {string} message - the error message
 */
export const generateError = (error, statusCode, message) => {
  // define error object
  const e = new Error(message || error.message || "unable to complete task.")
  e.statusCode = statusCode || 500

  // returns error
  return e
}