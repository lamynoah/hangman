
Hangman 

In this project, the player tries to guess a hidden word letter by letter. The game typically starts with a blank representation of the word, and the player has a limited number of attempts to guess the correct letters. For each incorrect guess, a part of a "hangman" figure is drawn. The goal is to guess the word before the entire hangman is drawn.

## Index

 - [Index](#index)
 - [Prerequisites](#prerequisites)
 - [Installation](#installation)
 - [Authors](#authors)

## Prerequisites

To install and run the project, you will need :

- [Go](https://go.dev/doc/install) installed on your local machine
- [Git](https://git-scm.com/downloads) installed on your local machine

## Installation

0. Ask for access from us to be added as a collaborator

1. Clone the Repository using git :
```bash
  git clone https://ytrack.learn.ynov.com/git/lnoah/hangman.git
```

2. Download the necessary packages with : 
```bash
  go mod tidy
```

3. For run the program in classic version with : 
```bash
  go run main.go file.txt
```
For run the program in Ascii version with :
```bash
go run main.go file.txt --letterfile AsciiFile.txt 
```
For Stop the program : 
in the Terminal : STOP 
For reload the program :
```bash
go run main.go --startwith save.txt
```
## Authors
- [@lnoah](https://ytrack.learn.ynov.com/git/lnoah)

