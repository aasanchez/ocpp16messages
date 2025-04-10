// Package main demonstrates how to use validated CiStringXXType types
// from the ocpp16messages/types package.
package main

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

func main() {
	input := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla sed nisi et diam convallis placerat. Aenean maximus dui lacinia elit pharetra, nec aliquam felis faucibus. Cras risus risus, faucibus convallis tempor ut, vehicula ut velit. Etiam eleifend lorem auctor libero convallis, ut posuere nisl pellentesque. Duis vitae elit nibh. Vivamus vitae feugiat lorem, id accumsan magna. Sed porta tortor ut sapien lacinia pharetra. `

	cistr, err := types.CiString500(input)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("✅ Length:", len(cistr.String()))
	fmt.Println("📦 Snippet:", cistr.String()[:30], "...")
}
