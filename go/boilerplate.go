package main

import (
	"github.com/zellyn/kooky"
)

func main() {
	// uses registered finders to find cookie store files in default locations
	// applies the passed filters "Valid", "DomainHasSuffix()" and "Name()" in order to the cookies
	kooky.RegisterFinder("firefox")
	kooky.FindAllCookies(kooky.Valid, kooky.DomainHasSuffix("adventofcode"), kooky.Name("PREF"))
}
