package main

import "fmt"

type personalInformation struct {
	name        string
	fullName    string
	numberPhone int64
	address     string
}

type PersonalInformation interface {
}

type PersonalInformationBuilder interface {
	AddName(name string) PersonalInformationBuilder
	AddFullName(fullName string) PersonalInformationBuilder
	AddNumberPhone(numberPhone int64) PersonalInformationBuilder
	AddAddress(address string) PersonalInformationBuilder
	Build() PersonalInformation
}

func (info *personalInformation) AddName(name string) PersonalInformationBuilder {
	info.name = name
	return info
}

func (info *personalInformation) AddFullName(fullName string) PersonalInformationBuilder {
	info.fullName = fullName
	return info
}

func (info *personalInformation) AddNumberPhone(numberPhone int64) PersonalInformationBuilder {
	info.numberPhone = int64(numberPhone)
	return info
}

func (info *personalInformation) AddAddress(address string) PersonalInformationBuilder {
	info.address = address
	return info
}

func (info *personalInformation) Build() PersonalInformation {
	return info
}

func NewPersonalInformationBuilder() PersonalInformationBuilder {
	return &personalInformation{}
}

func main() {
	info := NewPersonalInformationBuilder().
		AddName("Nhat").
		AddFullName("Nguyen Trong Nhat").
		AddNumberPhone(12344).
		AddAddress("Thai Nguyen").
		Build()

	info1 := NewPersonalInformationBuilder().
		AddName("Nhan").
		AddFullName("Nguyen Trong Nhan").
		AddAddress("Ha Noi").
		Build()

	fmt.Println(info)
	fmt.Println(info1)
}
