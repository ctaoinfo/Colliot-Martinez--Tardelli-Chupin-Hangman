package game

import (
	"fmt"
	"src/display"
	"src/input"
	"time"
)

func Game() {
	hangState := 0
	pointHangState := &hangState

	word, mapWord := input.WordChoice(input.ReturnListWord())
	var letterUsed string
	listHangState := input.ReturnListHangmanStates()

	for hangState < 9 && !CheckIsFound(mapWord) {
		fmt.Println(listHangState[hangState])
		display.PrintHidden(word, mapWord)
		fmt.Println("---------------------------------------")
		fmt.Println("Lettres utilisées :", letterUsed)
		fmt.Print("\n")
		input.CheckInput(mapWord, pointHangState, word, &letterUsed)
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("=======================================")
	}

	if CheckIsFound(mapWord) {
		display.PrintHidden(word, mapWord)
		fmt.Print("\n")
		fmt.Println("Bravo ! Tu as trouvé le mot")
		fmt.Println("Tes ancêtres sont fiers")
		time.Sleep(1500 * time.Millisecond)
	}

	if hangState >= 9 {
		fmt.Println(listHangState[hangState])
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Tu as maintenant la mort d'un stickman sur la conscience")
		fmt.Println("SHAME !")
		time.Sleep(1500 * time.Millisecond)
	}
}

func CheckIsFound(mapWordToFind map[string]bool) bool {
	//Vérifie si toutes les lettres du mot sont trouvées
	for key := range mapWordToFind {
		if mapWordToFind[key] == false {
			return false
		}
	}
	return true
}
