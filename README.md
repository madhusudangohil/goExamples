# goExamples
#cheatsheet

Formats all the go files in the current folder.
gofmt -s -w .

execute the program

go run main.go

build -> this will produce executable, e.g. main.exe, you can execute it ./main.exe
go build main.go

This will install the main.exe in the bin folder, GOBIN path
go install main.go

Package management
https://pkg.go.dev/

Software dependency
https://research.swtch.com/deps

updating dependency 
go get rsc.io/sampler
go get rsc.io/sampler@v1.31.

shows all the dependencies
go list -m all 


You should make your function exportable with an uppercase like func

go modules
Module is go support for dependency management. 
A module by definition is a collection of related packages with go.mod at its root.  
The go.mod file defines the
Module import path.
The version of go with which the module is created
Dependency requirements of the module for a successful build. 
It defines both project’s dependencies requirement and also locks them to their correct version.
https://blog.golang.org/using-go-modules

go mod init "project_folder_name"

go.mod and go.sum should be checked in to version control.

loop and if

for i:= 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println("from loop")
		}
	}
  // function returns 
  func Println(a ...interface{}) (n int, err error)
  you can assing return to a variable
  n,_ = fmt.Println("test")
  fmt.Println(n)
  _ is a throw away variable, if you assign it to variable you have to use that in the code, otherwise complier will throw an error.


Go is a statically typed language
var y = 42
y = "hello" // this is not allowed

var z string = "hello"


https://blog.alexellis.io/golang-writing-unit-tests/
unit testing - create a test file with _test.go
/folder/calc.go
/folder/calc_test.go

to run the unit tests
go test -v 
go test -cover
to run the unit tests in all the subfolders, run the below command from the root
go test ./...

for html report of test coverage
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 

For mocking
https://duythhuynh.medium.com/gomock-unit-testing-made-easy-b59a0e947ba7

go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
mockgen -destination=mocks/mock_calculator.go -package=calculator -source="calculator/calc.go"


map[keytype]valuetype , e.g. map[string]int.
var m map[string]int , it doesn't point to an initialized map, writing to it will cause panic so don't do that
m = make(map[string]int) - make function allocates and initializes a hash map data structure and returns a map value.



