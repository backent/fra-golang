package helpers

func PanifIfError(err interface{}) {
	if err != nil {
		panic(err)
	}
}
