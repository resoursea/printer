// A Printer package for Router
package printer

import (
	"fmt"
	"strings"

	"github.com/resoursea/api"
)

func Router(rt api.Router) {
	fmt.Println("\n--- PRINT ROUTER ---\n")
	router(rt, 0)
	fmt.Println("\n--- END PRINT ---\n")
}

func router(r api.Router, lvl int) {
	fmt.Printf("%sRoute: %s\n", strings.Repeat("	", lvl), r)

	for _, m := range r.Methods() {
		handler(m, lvl)
	}

	for _, c := range r.Children() {
		if r.IsSlice() {
			router(c, lvl)
		} else {
			router(c, lvl+1)
		}
	}
}

func handler(m api.Method, lvl int) {
	fmt.Printf("%s- Method: %s\n", strings.Repeat("	", lvl), m)
}
