package view

func viewFail(value string) {
	fmt.Print("Content-Type: text/plain\n\n")
	fmt.Print("Error:" + value)
}
