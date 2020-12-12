package middleware

import (
	_"errors"

	"github.com/gin-gonic/gin"

	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/logging"
	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/utils"
)

// VerifyAccess verify user's access
// to the server, and throw error if user is
// not allowed
func VerifyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		* Here you'll pull auth header, and send it
		* to checkIfAutherize. if auth is invalid, checkIfAutherize
		* will return an error.
		*/
		// get authorization header
		authorizationHeader := c.GetHeader("authorization")

		if err := checkIfAutherize(authorizationHeader); err == nil {
			logging.Log.Debug("user authorize succesfuly.")
			c.Next()
			return
		}

		logging.Log.Error("user in unauthorized.")
		utils.HandleError(c, 401, "unauthorized")
	}
}

// checkIfAutherize checks if user is autherize
// by aut header.
// NOTE: for this simpleStarter purposes, the function will
// allways pass authentication.
func checkIfAutherize(authHeader string) error {
	logging.Log.Info("got an auth header: ", authHeader)
	/*
	e.g for return statement that
	throw 401:
	return errors.New("error. user is unauthorized.")
	*/
	return nil
}