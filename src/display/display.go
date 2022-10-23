package display

import "fmt"

func PrintHidden(word string, listLetter map[string]bool) {
	//Parcours le mot, vérifie dans la map si la lettre est trouvée. Si oui, print la lettre. Si non, print "_".
	for _, letter := range word {
		tmp := "_"
		if listLetter[string(letter)] {
			tmp = string(letter)
		}
		fmt.Print(tmp)
	}
	fmt.Print("\n")
}
