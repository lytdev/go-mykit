package gid

import (
	"strings"
)

func Uuid1() string {
	uuid, _ := NewV1()
	return uuid.String()
}
func Uuid2() string {
	uuid, _ := NewV2(DomainPerson)
	return uuid.String()
}
func Uuid3() string {
	return NewV3(NamespaceDNS, "www.example.com").String()
}

func Uuid4() string {
	uuid, _ := NewV4()
	return uuid.String()
}

func Uuid5() string {
	uuid := NewV5(NamespaceDNS, "www.example.com")
	return uuid.String()
}

func FastUuid() string {
	return strings.Replace(Uuid4(), "-", "", -1)
}
