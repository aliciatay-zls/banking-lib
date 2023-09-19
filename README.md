## Installation
`go get github.com/udemy-go-1/banking-lib`

## Packages
### `logger`
* A global logger object is initialized the first time a method from this package is called. 
For example when running the [banking app](https://github.com/udemy-go-1/banking), this happens in `main.go` as `logger.Info()` is called there.
* To test log messages: call `ReplaceWithTestLogger()` which will initialize the logger object using the 
`zaptest/observer` package instead, and directly use the logs returned.
* To discard logs printed to console (such as during testing): call `MuteLogger()` at the start of the method or test. 
If needed call `UnmuteLogger()` to re-enable logging.
