package errors

type HTTPErrors []*HTTPError

func ListErrors() HTTPErrors {
	errors := []*HTTPError{}

	for code, _ := range codeMessages {
		errors = append(errors, NewHTTPError(code, 000))
	}

	return errors
}
