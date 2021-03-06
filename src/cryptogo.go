package main

import (
	"fmt"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	fmt.Println(CG.SimpleSinglePrice("bitcoin", "eur"))
	fmt.Println(CG.SimpleSinglePrice("ethereum", "eur"))
}
