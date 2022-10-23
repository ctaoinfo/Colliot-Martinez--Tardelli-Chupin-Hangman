package game

import (
	"flag"
	"fmt"
	"src/display"
	"src/input"
	"time"
)

func Game() {
	hangState := 0
	pointHangState := &hangState

	flag := FlagInput()
	word, mapWord := input.WordChoice(input.ReturnListWord(flag))

	var letterUsed string
	listHangState := input.ReturnListHangmanStates()

	for hangState < 9 && !CheckIsFound(mapWord) {
		display.PrintGame(listHangState, hangState, letterUsed, word, mapWord)
		input.CheckInput(mapWord, pointHangState, word, &letterUsed)
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("=======================================")
	}

	if CheckIsFound(mapWord) {
		display.PrintWin(word, mapWord)
	}

	if hangState >= 9 {
		display.PrintLose(listHangState, hangState)
	}
}

func CheckIsFound(mapWordToFind map[string]bool) bool {
	//Vérifie si toutes les lettres du mot sont trouvées
	for key := range mapWordToFind {
		if !mapWordToFind[key] {
			return false
		}
	}
	return true
}

func FlagInput() string {
	//Récupère le flag -word
	flagPtr := flag.String("words", "words", "Mot à trouver")
	// si flag vide, alors erreur
	if *flagPtr == "" {
		fmt.Println("Aucun mot n'a été renseigné")
		fmt.Println("Veuillez renseigner un mot avec le flag -words")
		fmt.Println("Exemple : go run main.go -words=words ou =words2 ou =words3")

	}

	flag.Parse()
	return *flagPtr
}
