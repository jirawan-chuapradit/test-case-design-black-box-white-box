package main

func nameChecker(name string) string {
	valid := "valid"
	invalid := "invalid"
	if len(name) < 1 {
		return invalid
	}
	if len(name) > 10 {
		return invalid
	}
	return valid
}
