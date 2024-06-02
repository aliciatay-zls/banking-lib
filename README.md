# Banking Web App - library
The `logger` package and most of `errs` were built under the Udemy course 
["REST based microservices API development in Golang"](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/).

## Installation
`go get github.com/aliciatay-zls/banking-lib`

## Packages
### `logger`
* Wraps the `go.uber.org/zap` [logging package](https://github.com/uber-go/zap) with some added functionalities.
* To get a global logger object: call a method from this package
  * A logger is initialized as global object the first time any method from this package is called. 
  * Example: this happens in `main.go` of [banking app](https://github.com/aliciatay-zls/banking) as `logger.Info()` is called.
* To test log messages: call `ReplaceWithTestLogger()` which will initialize the logger object using the 
`zaptest/observer` package instead, and directly use the logs returned.
* To discard logs printed to console (such as during testing): call `MuteLogger()` at the start of the method or test. 
If needed call `UnmuteLogger()` to re-enable logging.

### `errs`
* Provides an `AppError` type and helper functions for common HTTP error codes.

### `formValidator`
* Wraps some functions from the `go-playground/validator` [struct validation package](https://github.com/go-playground/validator/v10).
* To get a global struct validator object: call `Create()`
* To get the name of a country from its country code: call `GetCountryFrom()`
* The alias `un` is reserved for a custom validator (`useCustomUsernameValidator`). 

### `clock`
* Provides actual time in UTC (wraps `time` package) for use in actual code and dummy time for use in test code.
* To prevent failing tests due to current time being different when mocking and when running actual code: 
pass around the `StaticClock` object in the code during testing in place of the `RealClock` object, so that the 
time is "frozen" within each test.
  * Example: `accountService.go` and `accountService_test.go` of [banking app](https://github.com/aliciatay-zls/banking).

## Development
Update all packages periodically to the latest version:
   ```
   go get -u all
   go mod tidy
   ```
