# Learning Go (golang)

[Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/go-the-complete-developers-guide/)

## 01-helloworld
>`go run main.go`

## 02-cards
```go
var card = "Ace of Spades"
card = "Ace of Spades"

card2 := "Ace of Spades"
card2 = "Ace of Spades2"

card3 := newCard()

// slice
cards := []string{newCard(), "Ace of Diamonds"}
cards = append(cards, "Six of Spades")

// for loop
for i, card := range cards {
    fmt.Print(i, " "+card+", ")
}
```

>`go run main.go`
>
>`go run main.go state.go`
