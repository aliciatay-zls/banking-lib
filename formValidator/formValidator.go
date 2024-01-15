package formValidator

import (
	"bufio"
	"bytes"
	"embed"
	_ "embed"
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const UsernameAlias = "un"

var appValidator *validator.Validate
var codeCountryMap map[string]string

//go:embed codes.txt countries.txt
var data embed.FS

// Create initializes a global validator that uses a custom username validator, as well as a global map of
// countries' codes to their names.
func Create() {
	appValidator = validator.New(validator.WithRequiredStructEnabled())
	useCustomUsernameValidator()
	createCodeCountryMap()
}

// useCustomUsernameValidator registers a custom validator for a username field that needs to be filled,
// starts with alphabet, only alphanumeric and underscore allowed, min 6 chars and max 20 chars long.
func useCustomUsernameValidator() {
	unRegex := regexp.MustCompile("^[A-Za-z]\\w{5,19}$")

	if err := appValidator.RegisterValidation("validusername", func(fl validator.FieldLevel) bool {
		return unRegex.MatchString(fl.Field().String())
	}); err != nil {
		panic(err)
	}

	appValidator.RegisterAlias(UsernameAlias, "required,min=6,max=20,validusername")
}

// createCodeCountryMap uses the country code and name pairs provided by the ISO 3166-1 alpha-2 standard.
// Retrieved 8/1/2024: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
func createCodeCountryMap() {
	codeCountryMap = map[string]string{}

	f1, err := data.ReadFile("codes.txt")
	if err != nil {
		panic(err)
	}

	f2, err := data.ReadFile("countries.txt")
	if err != nil {
		panic(err)
	}

	s1 := bufio.NewScanner(bytes.NewReader(f1))
	s2 := bufio.NewScanner(bytes.NewReader(f2))
	for s1.Scan() {
		s2.Scan()
		codeCountryMap[s1.Text()] = s2.Text()
	}
}

func GetCountryFrom(code string) string {
	return codeCountryMap[code]
}

// Wrapped functions from go-playground/validator:

func Struct(s interface{}) validator.ValidationErrors {
	errs := appValidator.Struct(s)
	if errs != nil {
		var errsArr validator.ValidationErrors
		errors.As(errs, &errsArr)
		return errsArr
	}
	return nil
}
