# errscope
`errscope` finds code which can narrow the scope of error type.  
```
func f() error {
  return errors.New("this is example")
}

func main() {
  err := f() // it can narrow the scope of error
  if err != nil {
    // error handling
  }
}
```
If you write code like below, `errscope` does not point out.
```
func main() {
  if err := f(); err != nil { // it is ok
    // error handling
  }
}
```

# usage
```
$ go install gihtub.com/kimuson13/errscope/cmd/errscope@latest
$ errscope ./...
```
# implement soon...
I want to implement the function that can ignore some result.
For example
```
func main() {
  err := f() // ignore 
  if err != nil {
    // error handling
  }
}
```
This example code can narrow the scope of error type.  
But corresponding line has comment such as "// ignore".  
I expected that this line don't report by `errscope`.
