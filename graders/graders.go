package main

func grader(score int) string {

	if score >= 80 {
		return "A"
	}

	if score >= 70 {
		return "B"
	}

	if score >= 60 {
		return "C"
	}

	if score >= 50 {
		return "D"
	}

	if score >= 0 {
		return "F"
	}

	return "invalid score"
}
