package main

import (
	"example/gen/pb/example"
	privilegesv1 "example/gen/pb/jumpaku/privileges/v1"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	// The method descriptor can be accessed as follows:
	service := example.File_example_example_proto.Services().ByName("ExampleService")
	method := service.Methods().ByName("UpdateExample")

	// You can get the privileges extension as follows:
	privileges := proto.GetExtension(method.Options(), privilegesv1.E_Privileges).(*privilegesv1.Privileges)
	fmt.Println(privileges)
	// Output: required:{resource:"example" actions:"read" actions:"write"}
}
