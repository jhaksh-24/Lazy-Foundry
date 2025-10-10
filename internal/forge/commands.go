package forge

import (
	"fmt"
)

func Execute(command string) error {
	switch command {
	case "build":
		return Build()
	case "coverage":
		return Coverage()
	case "init":
		return Init()
	case "create":
		return Create()
	case "install":
		return Install()
	case "script":
		return Script()
	case "test":
		return Test()
	default:
		fmt.Println(command, "is yet to be implemented")
		return nil
	}
}
