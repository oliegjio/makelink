package main

import "os"
import "time"
import "fmt"
import "math/rand"

func main() {
    source := rand.NewSource(time.Now().UnixNano())
    random := rand.New(source)

    var symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    var symbolsNumber = 62
    var codeLength = 10
    var code = ""
    for i := 0; i < codeLength; i++ {
        var randomIndex = random.Intn(symbolsNumber)
        code += string(symbols[randomIndex])
    }

    var link = ""
    var linkFromStart = "[<"
    var linkFromEnd = ">]"
    var linkToStart = "[>"
    var linkToEnd = "<]"

    if len(os.Args) != 2 {
        fmt.Println("ERROR (no arguments provided): You should provide either `from` or `to` argument to this command!")
        return
    }
    var argument = os.Args[1]

    if argument == "from" {
        link = linkFromStart + code + linkFromEnd
    } else if argument == "to" {
        link = linkToStart + code + linkToEnd
    } else {
        fmt.Println("ERROR (wrong argument): You should provide either `from` or `to` argument to this command!")
        return
    }

    fmt.Println(link)
}
