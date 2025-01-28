# protoprivileges

`protoprivileges` provides a custom option to declare privileges required to process service methods.


## Usage with an example

Run `buf dep update` with the `buf.yaml` including a dependency `buf.build/googleapis/googleapis`.

```shell
```yaml
# buf.yaml
deps:
  - buf.build/googleapis/googleapis
```

Run `buf generate --include-imports` with protobuf schema file declaring required privileges.

```protobuf
service ExampleService {
  rpc GetExample(Req) returns (Res) {
    option (jumpaku.privileges.v1.privileges) = {
      required: [
        {resource: "example", actions: ["read", "write"]}
      ]
    };
  }
}
```

Extract the generated method extension in a source file.

```go
func main() {
	// The method descriptor can be accessed as follows:
	service := example.File_example_example_proto.Services().ByName("ExampleService")
	method := service.Methods().ByName("UpdateExample")

	// You can get the privileges extension as follows:
	privileges := proto.GetExtension(method.Options(), privilegesv1.E_Privileges).(*privilegesv1.Privileges)
	fmt.Println(privileges)
	// Output: required:{resource:"example" actions:"read" actions:"write"}
}
```

The full version of the above code fragments can be found in the https://github.com/Jumpaku/protoprivileges/tree/main/example .


## References

https://buf.build/jumpaku/privileges
https://github.com/Jumpaku/protoprivileges
