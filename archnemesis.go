package archnemesis

import (
	"fmt"
	"github.com/fatih/color"
)

func Missing(target string, inventory map[string]int){
	missing(target, inventory, 0, true)
}

func missing(target string, inventory map[string]int, level int, isMissing bool) {

	if level == 0 {
		color.Magenta(target)
	}

	for _, i := range RecipeList[target] {
		for j := 0; j < level+1; j++ {
			fmt.Printf("    ")
		}
		if inventory[i] > 0  {
			color.Green(fmt.Sprintf("%s: %d\n", i, inventory[i]))
		} else {
			if _, ok := RecipeList[i]; !ok {
				color.Red(fmt.Sprintf("%s: %d\n", i, inventory[i]))
			}else{
				color.White(fmt.Sprintf("%s: %d\n", i, inventory[i]))
				missing(i, inventory, level+1, false)
			}
		}
	}
}

