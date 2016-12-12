package main

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
