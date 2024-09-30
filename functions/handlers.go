package ascii

import (
	"html/template" // Importation du package pour manipuler des templates HTML
	"net/http"      // Package pour créer un serveur HTTP et gérer les requêtes/réponses
	"strings"
)

// Définition de la structure Page, qui représente les données envoyées aux templates HTML
type Page struct {
	Title        string // Le titre du template HTML à charger
	Art          string // Contenu ASCII Art généré (stocké ici)
	Text         string // Texte envoyé par l'utilisateur
	Banner       string // Le type de bannière sélectionnée par l'utilisateur
	MessageError string // Message d'erreur à afficher en cas de problème
	Code         int
	HttpErr      string
}

// Variable globale pour la page principale, initialisée avec le titre "index.html"
var art = &Page{Title: "index.html"}

// Fonction pour afficher le template HTML. Elle prend en entrée un writer et une structure Page
func renderTemplate(w http.ResponseWriter, p *Page) {
	// ParseFiles charge le fichier template spécifié par Title (index.html par exemple)
	temp, err := template.ParseFiles("./templates/" + p.Title)
	if err != nil {
		// En cas d'erreur, renvoie une réponse HTTP avec le message d'erreur et le code 500 (erreur serveur)
		errorHandler(w, http.StatusInternalServerError, "Oops!!! Internal Server Error")
		return
	}
	// Execute applique les données de la structure Page au template chargé
	err = temp.Execute(w, p)
	art = &Page{Title: "index.html"}
	if err != nil {
		// Si une erreur se produit pendant l'exécution du template, une erreur 500 est renvoyée
		errorHandler(w, http.StatusInternalServerError, "Oops!!! Internal Server Error")
		return
	}
}

// Fonction qui gère la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Si c'est le cas, renvoie une erreur 405 (méthode non autorisée)
		errorHandler(w, http.StatusMethodNotAllowed, "Oops!! Method not allowed")
		return
	}
	// Vérifie si l'URL demandée est bien "/"
	if r.URL.Path != "/" {
		// Si ce n'est pas le cas, renvoie une erreur 404 (page non trouvée)
		if strings.ContainsRune(r.URL.Path[1:], '/') {
			http.Redirect(w, r, "/notFound", http.StatusFound)
			return
		}
		errorHandler(w, http.StatusNotFound, " Oops!!Page Not Found")
		return
	}
	// Affiche le template index.html en utilisant les données de la variable 'art'
	renderTemplate(w, art)
	// Remet à zéro le message d'erreur après affichage
	art.MessageError = ""
}

// Fonction qui gère la création et l'affichage de l'ASCII Art
func ArtHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifie si la méthode HTTP est "GET" (ce qui n'est pas permis ici)
	if r.Method != http.MethodPost {
		// Si c'est le cas, renvoie une erreur 405 (méthode non autorisée)
		errorHandler(w, http.StatusMethodNotAllowed, "Oops!! Method not allowed")
		return
	}

	// Récupère les valeurs "text" et "banner" envoyées par le formulaire
	text := r.FormValue("text")

	banner := r.FormValue("banner")

	// Si l'un des champs est vide, renvoie une erreur 400 (mauvaise requête)
	if len(text) > 3000 || text == "" || banner == "" {
		errorHandler(w, http.StatusBadRequest, "Oops!! Bad Request")
		return
	}

	// Appelle la fonction Art pour générer l'ASCII Art en fonction du texte et de la bannière
	data, err := Art(text, banner)
	if err != nil {
		// Gère les erreurs spécifiques renvoyées par la fonction Art
		if err.Error() == "not printable" {
			art.MessageError = "*your text is not printable"
		} else if err.Error() == "not a banner" {
			art.MessageError = "*This banner is not valide"
		} else if err.Error() == "this banner is not exists" {
			errorHandler(w, http.StatusBadRequest, "Oops!! Bad Request")
			return
		} else {
			// Si l'erreur n'est pas reconnue, renvoie une erreur 404 (ressource non trouvée)
			errorHandler(w, http.StatusNotFound, "Oops!!Page Not Found")
			return
		}
	}
	// Si tout est correct, stocke les données dans la variable 'art'
	art.Text = text
	art.Banner = banner
	art.Art = string(data)

	// Redirige l'utilisateur vers la page d'accueil pour afficher le résultat
	http.Redirect(w, r, "/", http.StatusFound)
}

// Fonction qui gère les erreurs HTTP et affiche les pages d'erreur correspondantes
func errorHandler(w http.ResponseWriter, status int, message string) {
	art.Code = status
	art.HttpErr = message
	// Définit le code de statut de la réponse HTTP
	w.WriteHeader(status)
	art.Title = "error.html"
	renderTemplate(w, art)
	art.Title = "index.html"
}
