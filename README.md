<h1>Hangman classic</h1>

***


This project's goal was to recreate the Hangman game in golang.

``Hangman`` is a very famous game that a player must guess the word by suggesting letters being limited by a low number of guesses. These numbers of attempts left are represented into a hangman hence the name of the game.

<h2>Requirements</h2>

***

In order to play this game without any problems, we highly recommend you to use Windows as your operating system

<h2>How to play</h2>

***

In order to start a game, you can either :

* Launching the game without any argument will launch the menu where you can start a new game as well as opening the settings in order to choose which word files you wanted to use for the game

````
  __  __       _
 |  \/  |     (_)
 | \  / | __ _ _ _ __    _ __ ___   ___ _ __  _   _
 | |\/| |/ _` | | '_ \  | '_ ` _ \ / _ \ '_ \| | | |
 | |  | | (_| | | | | | | | | | | |  __/ | | | |_| |
 |_|  |_|\__,_|_|_| |_| |_| |_| |_|\___|_| |_|\__,_|


s : launch a new game
o : opens the game's settings (change the words files, use an ASCII letter mode)
q : exit the game
Enter an input :
````
- Launching the game by using arguments (this will skip the menu and launch the game) :

  - Writing the words file name (ex: ```go run ./cmd words2.txt``` will use a word in the words2.txt file)
  - Writing the argument ```--startWith``` followed by the name of the save file will resume the game that was previous save written on the specified file (ex: ```--startWith data/save.txt```)


During a game, the program will ask the user to enter a valid character (capital or lower alphabetic character(s) or a "/") or an entire word if the player thinks he guessed the right word (however, if it happens to not be the correct word, 2 attempts will be subtracted from the player's remaining attempts). If the player writes a letter that was not already used and that happens to be in the word, the program will modify the list of underscores by replacing the one(s) where the letter is contained in the word. If he enters a wrong letter he looses an attempt and the hangman printed out in the terminal is more completed.

Example of a game :
````
Not present in the word, 8 attempts remaining
  +---+
      |
      |
      |
      |
      |
=========
e__de
Already guessed letters / words : [e d z m]
etude
Enter an input :
````
<h3>Special commands</h3>

* /r : restarts the game using the same words file
* /m : opens the main menu
* /s : saves the current game, writes the information in ```save/data.txt``` and leaves the program

***


**<a href="https://trello.com/invite/b/67052d4d9addbd1237aa2f53/ATTIf9dfc8cfabb4bd037d9451a424cac9346DEABC09/hangman-classic" target="_blank">Link to the trello</a><br>**
