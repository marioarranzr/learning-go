## Maps

```go
lookup := make(map[string]int)
lookup["goku"] = 9001
power, exists := lookup["vegeta"]
fmt.Println(power, exists) // 0 false

fmt.Println(len(lookup)) // 1
delete(lookup, "goku")
fmt.Println(len(lookup)) // 0
```

>`go run main.go`
