# HangManGroupe

Le projet hangman consiste à réaliser un jeu du pendu dans le terminal.

## Langage

 - Go

## Auteurs

- Andrea Taredelli
- Théo
- Tao Chupin

## Installation 

Pour récuperer le dossier, ouvrir un emplacement où vous voulez le télécharger, ouvrir un
git bash puis taper la commande suivant : https://github.com/ctaoinfo/hangmanGroup.git

## Utilisation / Règles

Le jeu consiste à retrouer un mot séléctionné au hazard dans un fichier 'dictionnaire', pour jouer vous devez entrer une lettre (minuscules ou majuscules) 
et retouver le mot caché pour cela vous aurez 10 vies, après une tentative de lettre vous perdrez 1 vie jusqu'à arriver à 0 si cela arrive vous retrouverez 
votre personnage pendu.
Vous pourrez tout de même tenter de retrouver directement le mot en l'écrivant directement si celui-ci n'est pas bon pour perdrez 2 vies.
Les vies sont symbolisé par le pendu en construction.

## Instruction

Pour commencez il faut ouvir un terminal CMD ou bash puis rendez-vous dans le fichier src du dossier télécharger.

Puis lancer le programme avec la commande suivant: 'go run .\main.go -words=words'
si vous voulez utiliser le deuxième ou troisième dictionnaire ajouter un chiffre derrière, exemple 'go run ./main.go -words=words2' 
ou alors 'go run ./main.go -words=words3'.
Après cela vos pourrez commencez à jouer pour trouver le mot.