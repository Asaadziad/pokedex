package main

func Eval(s []string) error {
	err := Commands()[s[0]].callback()
	return err
}
