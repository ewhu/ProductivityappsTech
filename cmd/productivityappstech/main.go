// cmd/productivityappstech/main.go
package main

import (
"flag"
"log"
"os"

"productivityappstech/internal/productivityappstech"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := productivityappstech.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
