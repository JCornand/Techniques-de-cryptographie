package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const colonnes = 8
const bloc = 5

func main() {
	log.Println("Application démarrée")

	texte := "Voici un exemple de texte pour tester le chiffrement scytale des spartiates !"
	fmt.Println("Texte original:", texte)

	// Nettoyage
	texteNettoye := nettoyerTexte(texte)
	fmt.Println("Texte nettoyé:", texteNettoye)

	// Chiffrement
	texteCrypte := crypter(texteNettoye, colonnes)
	fmt.Println("Texte crypté:", texteCrypte)

	// Déchiffrement
	texteDecrypte := decrypter(texteCrypte, colonnes)
	fmt.Println("Texte décrypté:", texteDecrypte)
}

// Nettoyer texte (minuscules, lettres et chiffres seulement)
func nettoyerTexte(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[a-z0-9]`)
	return strings.Join(re.FindAllString(s, -1), "")
}

// Ajouter tirets toutes les n lettres
func ajouterTirets(s string, n int) string {
	var builder strings.Builder
	for i, r := range s {
		if i > 0 && i%n == 0 {
			builder.WriteRune('-')
		}
		builder.WriteRune(r)
	}
	return builder.String()
}

// Chiffrement type Scytale
func crypter(s string, colonnes int) string {
	longueur := len(s)
	lignes := longueur / colonnes
	if longueur%colonnes != 0 {
		lignes++
	}

	// Création du tableau et remplissage
	tableau := make([][]rune, lignes)
	for i := range tableau {
		tableau[i] = make([]rune, colonnes)
		for j := 0; j < colonnes; j++ {
			index := i*colonnes + j
			if index < longueur {
				tableau[i][j] = rune(s[index])
			} else {
				tableau[i][j] = 0 // vide pour la lecture, pas de X visibles
			}
		}
	}

	// Lecture colonne par colonne
	var result []rune
	for j := 0; j < colonnes; j++ {
		for i := 0; i < lignes; i++ {
			if tableau[i][j] != 0 {
				result = append(result, tableau[i][j])
			}
		}
	}

	return ajouterTirets(string(result), bloc)
}

// Déchiffrement type Scytale
func decrypter(s string, colonnes int) string {
	// Supprimer les tirets
	s = strings.ReplaceAll(s, "-", "")
	longueur := len(s)
	lignes := longueur / colonnes
	if longueur%colonnes != 0 {
		lignes++
	}

	// Création du tableau vide
	tableau := make([][]rune, lignes)
	for i := range tableau {
		tableau[i] = make([]rune, colonnes)
	}

	// Remplissage colonne par colonne
	index := 0
	for j := 0; j < colonnes; j++ {
		for i := 0; i < lignes; i++ {
			if index < longueur {
				tableau[i][j] = rune(s[index])
				index++
			}
		}
	}

	// Lecture ligne par ligne
	var result []rune
	for i := 0; i < lignes; i++ {
		for j := 0; j < colonnes; j++ {
			if tableau[i][j] != 0 {
				result = append(result, tableau[i][j])
			}
		}
	}

	return string(result)
}


