# HangManGroupe

Le projet hangman consiste à réaliser un jeu du pendu dans le terminal.
Les consignes sont sur le repo suivant : https://github.com/Lyon-Ynov-Campus/YTrack/tree/master/subjects/hangman/hangman-classic

## Langage

 - Go

## Auteurs

- Andrea Taredelli
- Théo Colliot-Martinez
- Tao Chupin

## Installation 

Pour récuperer le dossier, ouvrir un emplacement où vous voulez le télécharger, ouvrir un
git bash puis taper la commande suivant : https://github.com/ctaoinfo/hangmanGroup.git

## Règles

Le jeu consiste à retrouer un mot séléctionné au hazard dans un fichier 'dictionnaire', pour jouer vous devez entrer une lettre (minuscules ou majuscules) 
et retouver le mot caché pour cela vous aurez 10 vies, après une tentative de lettre vous perdrez 1 vie jusqu'à arriver à 0 si cela arrive vous retrouverez 
votre personnage pendu.
Vous pourrez tout de même tenter de retrouver directement le mot en l'écrivant directement si celui-ci n'est pas bon pour perdrez 2 vies.
Les vies sont symbolisé par le pendu en construction.

## Instruction

Pour commencez il faut ouvir un terminal CMD ou bash puis rendez-vous dans le fichier src du dossier télécharger.

Puis lancer le programme avec la commande suivante : 'go run .\main.go'
si vous voulez utiliser le deuxième ou troisième dictionnaire ajouter le flag words derrière, exemple : 'go run ./main.go -words=words2' 
ou alors 'go run ./main.go -words=words3'.
Après cela vos pourrez commencez à jouer pour trouver le mot.
Bonne Chance !!
