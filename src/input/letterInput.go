package input

import (
	"fmt"
	"strings"
)

func CheckInput(mapWordToFind map[string]bool, hangState *int, word string, letterUsed *string) {
	//Récupère l'input, si mauvais input, erreur.
	//Parcours la map des lettres découvertes, si input == key de la map, alors la valeur de clé = true
	//Si après parcours de la map pas de lettre sont trouvé, alors message d'échec
	var input string

	fmt.Scanln(&input)
	input = ToUpper(input)

	if !isAlpha(input) {
		fmt.Println("Caractère non valide")
		return
	}

	if len(input) > 1 {
		if input == word {
			for k := range mapWordToFind {
				mapWordToFind[k] = true
			}
		} else {
			*hangState += 2
		}
	} else {
		if strings.Contains(*letterUsed, input) {
			fmt.Println("Cette lettre a déjà été utilisée")
		} else {
			for k := range mapWordToFind {
				if k == input {
					mapWordToFind[k] = true
					*letterUsed += input
					return
				}
			}
			*letterUsed += input
			fmt.Println("Mauvaise lettre")
			*hangState += 1
		}
	}
}
