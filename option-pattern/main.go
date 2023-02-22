package main

import (
	"fmt"
)

type Info struct {
	name        string
	fullName    string
	phoneNumber int
}

func New(options ...func(*Info)) *Info {
	svr := &Info{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func WithName(name string) func(*Info) {
	return func(s *Info) {
		s.name = name
	}
}

func WithFullname(fullName string) func(*Info) {
	return func(s *Info) {
		s.fullName = fullName
	}
}

func WithPhoneNumber(phoneNumber int) func(*Info) {
	return func(s *Info) {
		s.phoneNumber = phoneNumber
	}
}

func main() {
	info := New(
		WithName("Nhat"),
		WithFullname("Nguyen Trong Nhat"),
		WithPhoneNumber(01234567),
	)
	fmt.Println(info)
}
