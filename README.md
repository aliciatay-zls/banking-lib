## Installation
`go get github.com/udemy-go-1/banking-lib`

## Packages
### `logger`
* A global logger object is initialized the first time a method from this package is called. 
For example when running the [banking app](https://github.com/udemy-go-1/banking), this happens in `main.go`.
* Call `TurnOffLogger()` to disable logging to console (such as during testing) and `TurnOnLogger()` to re-enable logging.