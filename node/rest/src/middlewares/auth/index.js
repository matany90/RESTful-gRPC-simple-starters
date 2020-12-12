/**
 * checkIfAuthenticated will process the request
 * authorization header and will determine if the user
 * is allowed to have access to the endpoint.
 *
 * @param {object} req - The request.
 * @param {object} res - The response.
 * @param {Function} next - Next function in chain.
 */
export const checkIfAuthenticated = (req, res, next) => {
  // whitelist paths are paths that
  // can skip auth. if the request source
  // is one of those paths, auth check will be skipped.

  const whitelist = [
    /* Some whitelist paths:
    e.g "/api/v1/foo-with-no-auth"
    */
  ]

  // check if path is in whitelist path
  if (whitelist.includes(req.path)) {
    next()
  }

  try {
    console.log("Checking if user is authenticate...")
    // Do an api call to validate user...
    console.log("User authenticate succssesfuly!")
    next()
  } catch (e) {
    return res.status(401).json({ error: "unable to authenticate user."})
  }
}