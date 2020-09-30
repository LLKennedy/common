package assertutil

import "github.com/stretchr/testify/assert"

// EqualError asserts that a specific error string will be returned, matching assert.ErrorAssertionFunc signature
func EqualError(err string) assert.ErrorAssertionFunc {
	return func(t assert.TestingT, e error, v ...interface{}) bool {
		return assert.EqualError(t, e, err, v...)
	}
}

// MatchError uses a custom matching function for the returned error, matching assert.ErrorAssertionFunc signature
func MatchError(match func(error) bool) assert.ErrorAssertionFunc {
	return func(t assert.TestingT, e error, v ...interface{}) bool {
		return assert.True(t, match(e), v...)
	}
}

// Error aliases assert.Error for ease of use
func Error(t assert.TestingT, e error, v ...interface{}) bool {
	return assert.Error(t, e, v...)
}

// NoError aliases assert.NoError for ease of use
func NoError(t assert.TestingT, e error, v ...interface{}) bool {
	return assert.NoError(t, e, v...)
}
