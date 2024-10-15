<h1>Hangman classic</h1>


This project's goal was to recreate the Hangman game in golang.

``Hangman`` is a very famous game that a player must guess the word by suggesting letters being limited by a low number of guesses. These numbers of attempts left are represented into a hangman hence the name of the game.

<h2>How to play</h2>


In order to start a game, you can either :

* Launching the game without any argument will launch the menu where you can start a new game as well as opening the menu in order to choose which word files you wanted to use for the game

- Launching the game by using arguments (this will skip the menu and launch the game) : 

  - Writing the word file name ```go run ./cmd words2.txt``` will use a word in the words2.txt file
  - Writing the argument ```--startWith``` followed by the name of the save file (in this case, data/save.txt) will resume the game that was previously save to this file


