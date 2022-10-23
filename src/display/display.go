package display

import (
	"fmt"
	"time"
)

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

func PrintWin(word string, mapWord map[string]bool) {
	PrintHidden(word, mapWord)
	fmt.Print("\n")
	fmt.Println("Bravo ! Tu as trouvé le mot")
	fmt.Println("Tes ancêtres sont fiers")
	time.Sleep(1500 * time.Millisecond)
}

func PrintLose(listHangState []string, hangState int) {
	fmt.Println(listHangState[hangState])
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Tu as maintenant la mort d'un stickman sur la conscience")
	fmt.Println("SHAME !")
	time.Sleep(1500 * time.Millisecond)
}

func PrintGame(listHangState []string, hangState int, letterUsed string) {
	fmt.Println(listHangState[hangState])
	PrintHidden(word, mapWord)
	fmt.Println("---------------------------------------")
	fmt.Println("Lettres utilisées :", letterUsed)
	fmt.Print("\n")
}
