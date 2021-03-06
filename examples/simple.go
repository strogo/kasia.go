package main

import (
    "kasia"
    "os"
    "fmt"
)

type Ctx struct {
    h, w string
}

func main() {
    ctx := &Ctx{"Hello", "world"}

    tpl := kasia.New()
    err := tpl.Parse("$h $w!\n")
    if err != nil {
        fmt.Println(err)
        return
    }

    err = tpl.Run(os.Stdout, ctx)
    if err != nil {
        fmt.Println(err)
    }
}
