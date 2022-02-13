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
