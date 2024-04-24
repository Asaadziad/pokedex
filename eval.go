package main

func Eval(cfg *Config, s []string) error {
	err := Commands()[s[0]].callback(cfg)
	return err
}
