package strings

func internalFunction() {
}

// Must begin with a capital letter in order to be exported
func CountOddEven(s string) (odds, evens int) {
    odds, evens = 0, 0
    for _, c := range s {
        if int(c)%2 == 0 {
            evens++
        } else {
            odds++
        }
    }
    return
}