package constants

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)

	allErrors = append(allErrors, UserErrors...)
	allErrors = append(allErrors, GeneralErrors...)

	for _, e := range allErrors {
		if err.Error() == e.Error() {
			return true
		}
	}

	return false
}
