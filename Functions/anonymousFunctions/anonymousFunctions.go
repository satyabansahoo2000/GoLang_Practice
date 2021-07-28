package anonymousFunctions

func Anonymous() func() int {
	// returns a func() which returns int value
	return func() int {
		return 45
	}
}