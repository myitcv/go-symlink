package mylib2

import "bitbucket.org/pkg/inflect"

func DoSomethingDifferent() string {
	return inflect.Camelize("hello_world")
}
