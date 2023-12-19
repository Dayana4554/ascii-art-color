package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// os.Args[1] == couleur
	// os.Args[2] == lettres a transformer
	// os.Args[3] == le mot en entier
	// flagPtr := flag.String("<color>", "<letters to be colored>")
	// je parcours l entree pour separe les lettres normales de celles en couleur
	if len(os.Args) == 4 {
		Ligne(os.Args[3])
	} else if len(os.Args) == 3 {
		Ligne(os.Args[2])
	} else {
		fmt.Println("Error")
	}
}

func WordbyWordScan(phrase string) {
	lines, err := ReadLines("standard.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	m := make(map[int][]string)
	values := make([]string, 0) // ici sinon ca fait des tableaux vides
	count := 0
	count2 := 0
	for i := 1; i <= len(lines); i++ {
		count++        // car chaque character fait 8 lignes de haut et une ligne de plus car un espace
		if count < 9 { // d'ou cette ligne
			values = append(values, lines[i]) // j'ajoute toutes les 10 lignes le character
		} else { // sinon j'ajoute le tableau de string (qui represente un character) dans un élément de la map
			key := 32 + count2
			m[key] = values
			count2++
			count = 0
		}
	}
	var s string
	var stock int
	var tab []string
	var tab2 [][]string
	// sert a prendre toutes les lettres de la string
	for _, r := range phrase { // je parcours mon entree
		new := int(r)               // je convertie pour parcourir la map
		for key, value := range m { // je parcours la map
			if key == new { // si je trouve la lettre
				for i := range value { // je prend le num de la derniere ligne
					stock = i
				}
				for i := stock - 7; i <= stock; i++ { // je parcours a l envers pour la lettre et j ajoute chaque ligne dans un tableau
					s += value[i]
					tab = append(tab, s)
					s = ""
				}
			}
		}
		tab2 = append(tab2, tab) // je fais un tableau qui a pour element chaque lettre
		tab = nil
	}
	errr := Flag(phrase, tab2)
	if errr != nil {
		log.Fatalf("Invalid arguments : %s", errr)
	}
	// je veux prendre les memes de chaque nombre
}

func PrintResult(color, ind, arg string, tab [][]string) {
	var tabind []int       // Tableau qui va stocker les indexs correspondant a ceux des lettres en ascii-art dans tab
	var tabos []string     // Tableau qui va stocker dans chaque case une lettre de l'os.Args qui contient le string à afficher
	toColor := false       // Bool qui permet de savoir si la lettre est a colorer ou pas
	if len(os.Args) == 3 { // Si pas de lettre à colorer spécifiée, tout colorier
		toColor = true
	}
	for _, r := range arg {
		tabos = append(tabos, string(r)) // Stocke les lettres dans tabos
	}
	for i := 0; i < len(tabos); i++ {
		for _, r := range ind {
			if string(r) == tabos[i] { // Si une lettre de l'os.Args qui contient les lettres a colorer correspond
				tabind = append(tabind, i) // à une lettre de l'os.Args qui contient la phrase à afficher
			} // alors stocker son index
		}
	}
	for i := 0; i < 8; i++ {
		for index := 0; index < len(tab); index++ {
			for j := 0; j < len(tabind); j++ { // Si l'index correspond à l'un des index stockés dans tabind
				if index == tabind[j] { // Alors passer toColor à true
					toColor = true
					break
				}
			}
			if toColor { // Pour colorer la lettre
				fmt.Print(color, tab[index][i])
				if len(os.Args) == 4 {
					toColor = false
				}
			} else { // Sinon, laisser en blanc
				fmt.Print("\033[37m", tab[index][i])
			}
		}
		fmt.Print("\n")
	}
}

func Flag(phrase string, tab [][]string) error {
	ErrInvalidArguments := errors.New("invalid color") // erreur si jamais la couleur rentrée est invalide
	var color string                                   // Variable qui va contenir le code couleur
	if strings.Contains(os.Args[1], "--color") {       // Si il y a le flag --color
		for i := 8; i < len(os.Args[1]); i++ { // Faire une boucle qui va stocker dans color la couleur demandée
			color += string(os.Args[1][i])
		}
		switch color { // Switch qui va regarder quelle couleur est demandée et va stocker le code couleur correspondant
		case "red":
			color = "\033[31m"
			break
		case "green":
			color = "\033[32m"
			break
		case "yellow":
			color = "\033[33m"
			break
		case "blue":
			color = "\033[34m"
			break
		case "purple":
			color = "\033[35m"
			break
		case "cyan":
			color = "\033[36m"
			break
		case "white":
			color = "\033[37m"
			break
		case "orange":
			color = "\033[38;5;202m"
		default:
			return ErrInvalidArguments
		}
	}
	if len(os.Args) == 4 { // Si les lettres a colorées sont spécifiées on colorie les bonnes lettres
		PrintResult(color, os.Args[2], phrase, tab)
	} else { // Sinon on colorie tout
		PrintResult(color, "", phrase, tab)
	}
	return nil
}

// je m occuper des \n
func Ligne(s string) {
	if strings.Contains(s, "\\n") || strings.Contains(s, "\n") { // la c quand y en a
		regex := regexp.MustCompile(`\\n|\n`)
		// test := regex.ReplaceAllString(s, "\n")
		// fmt.Println(test)
		test2 := regex.Split(s, -1)
		for _, r := range test2 { // la c est quand y en a plusieurs d affiler
			if r == "" {
				fmt.Printf("\n")
			} else { // la quand y en a pas
				WordbyWordScan(r)
			}
		}
	} else if s == "" { // la c quand la string en arg est vide
		os.Exit(0)
	} else { // la quand y  pas de \n
		if len(os.Args) == 4 {
			WordbyWordScan(os.Args[3])
		} else {
			WordbyWordScan(os.Args[2])
		}
	}
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
