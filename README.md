HangMan Groupe
Le projet hangman consiste à réaliser un jeu du pendu dans le terminal,
Pour lancer le jeu rendez vous dans le fichier du jeu: 'hangmanGroup/src',
puis lancer le programme avec la commande suivant: 'go run .\main.go -words=words'
si vous voulez utilisez le deuxième ou troisième dictionnaire ajouter un chiffre derrière, exemple 'go run ./main.go -words=words2' 
ou alors 'go run ./main.go -words=words3'.
Le jeu consiste à retrouer un mot séléctionné au hazard dans un fichier 'dictionnaire', pour jouer vous devez entrer une lettre (minuscules ou majuscules) 
et retouver le mot caché pour cela vous aurez 10 vies après une tentative de lettre vous perdrez 1 vie jusqu'à arriver à 0 si cela arrive vous retrouverez 
votre personnage pendu.
Vous pourrez tout de même tenter de retrouver directement le mot en l'écrivant directement si celui-ci n'est pas bon pour perdrez 2 vies.
