/* eslint-disable */
import jsonSchema from "../../../ops/data-files/schema/json-schema.json"

// Init json schema validetor
const Validator = require("jsonschema").Validator
const v = new Validator()

/**
 * validateData will determine if the request.body
 * is valid, according to the rules and the request.body
 * that given
 *
 * @param {string} identifier - validation identifier
 */
export const validateData = (identifier) => async (req, res, next) => {
  // validate request body
  const { errors } = v.validate(req.body, jsonSchema[identifier])

  // checks if there is any errors
  if (errors.length) {
    console.log("Opps! request.body is invalid.")

    // returns an validation error
    return res.status(400).json({
      error: {
        message: `unable to validate ${identifier}.`,
        ...formattedError(errors)
      },
    })
  }

  // All good!
  console.log("Validate Data is successfully!")
  return next()
}

/**
 * formattedError returns an formatted Error string
 *
 * @param {Array} errors
 * @returns {Object} - formatted errors object
 */
const formattedError = (errors) => {
  // define formatted erros object
  let formattedError = {}

  // build the formatted error object
  errors.forEach(err => {
    const [el, name] = err.property.split(".")
    formattedError[name] = err.message
  })

  // returns the error object
  return formattedError
}