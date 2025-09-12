package error

import "errors"

func ErrMapping(err error) bool {
	if err == nil {
		return false // or handle nil appropriately
	}

	allErr := append(GeneralErrors[:], UserErrors[:]...)
	for _, targetErr := range allErr {
		if errors.Is(err, targetErr) {
			return true
		}
	}
	return false
}
