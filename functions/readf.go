package ascii

import (
	"errors"  // Package utilisé pour gérer et créer des erreurs personnalisées
	"io"      // Package pour les opérations d'entrée/sortie comme la lecture de fichiers
	"os"      // Package pour travailler avec les fichiers et le système d'exploitation
)

// Fonction qui calcule le nombre de lignes dans un tableau de bytes (les données du fichier)
func CalculLine(dataIn []byte) int {
	lines := 0
	// Parcourt chaque caractère du tableau
	for _, char := range dataIn {
		// Incrémente le compteur chaque fois qu'un caractère '\n' (nouvelle ligne) est trouvé
		if char == '\n' {
			lines++
		}
	}
	// Retourne le nombre total de lignes
	return lines
}

// Fonction qui lit un fichier texte et le convertit en un tableau de 95 tableaux de bytes
func ReadF(name string) ([95][]byte, error) {
	// Ouvre le fichier à partir du dossier "banners" avec le nom passé en argument, et ajoute l'extension ".txt"
	file, err := os.Open("./banners/" + name + ".txt")
	// Si une erreur survient lors de l'ouverture du fichier, retourne un tableau vide et l'erreur
	if err != nil {
		return [95][]byte{}, err
	}
	// Ferme automatiquement le fichier après utilisation
	defer file.Close()

	// Lit le contenu entier du fichier
	dataIn, err := io.ReadAll(file)
	// Vérifie si le nombre de lignes dans le fichier est différent de 855 (le format attendu pour une bannière)
	if CalculLine(dataIn) != 855 {
		// Si ce n'est pas le cas, retourne une erreur "not a banner"
		return [95][]byte{}, errors.New("not a banner")
	}

	// Initialise une nouvelle slice vide pour stocker les données du fichier sans les caractères '\r' (retour à la ligne)
	data := []byte{}
	// Parcourt chaque caractère dans les données du fichier
	for _, char := range dataIn {
		// Ignore le caractère '\r' (retour chariot), car il n'est pas nécessaire pour notre traitement
		if char != '\r' {
			data = append(data, char)
		}
	}

	// Si une erreur survient lors de la lecture, retourne un tableau vide et l'erreur
	if err != nil {
		return [95][]byte{}, err
	}

	// Initialisation d'un tableau pour stocker chaque caractère ASCII (de 32 à 126)
	newTab := [95][]byte{}

	// Compteurs pour le processus de création du tableau
	count := 0    // Compteur de lignes pour chaque caractère
	curnnt := 0   // Indice pour chaque caractère dans le tableau (indice ASCII)

	// Parcourt les données du fichier à partir de l'indice 1
	for i := 1; i < len(data); i++ {
		// Ajoute chaque caractère au tableau `newTab` pour le caractère en cours de traitement
		newTab[curnnt] = append(newTab[curnnt], data[i])

		// Si le caractère est un saut de ligne (code ASCII 10), incrémente le compteur de lignes
		if data[i] == 10 {
			count++
		}

		// Si 8 lignes sont atteintes (car chaque caractère dans le fichier de bannière occupe 8 lignes), 
		// on passe au caractère suivant dans le tableau `newTab`
		if count == 8 {
			curnnt++  // Passe au caractère ASCII suivant
			count = 0  // Réinitialise le compteur de lignes
			i++  // Incrémente l'indice pour sauter la ligne vide entre chaque caractère
		}
	}

	// Retourne le tableau rempli avec les caractères ASCII et aucune erreur
	return newTab, nil
}
