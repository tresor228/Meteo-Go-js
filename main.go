package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// WeatherData structure pour stocker les données météo
type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"windSpeed"`
	Icon        string  `json:"icon"`
}

// APIResponse structure pour la réponse de l'API OpenWeatherMap
type APIResponse struct {
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

func main() {
	// Charger les variables d'environnement depuis .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Erreur lors du chargement du fichier .env:", err)
	}

	// Configurer le router
	r := mux.NewRouter()

	// Servir les fichiers statiques
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Page d'accueil
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	// API endpoint pour récupérer les données météo
	r.HandleFunc("/api/weather", getWeather).Methods("GET")

	// Démarrer le serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Serveur démarré sur le port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// Fonction pour récupérer les données météo
func getWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Lomé" // Ville par défaut
	}

	apiKey := os.Getenv("MA_KEY")
	if apiKey == "" {
		http.Error(w, "Clé API non configurée", http.StatusInternalServerError)
		return
	}

	// Appel à l'API OpenWeatherMap
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&lang=fr&appid=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erreur lors de l'appel à l'API météo", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Erreur API: "+resp.Status, resp.StatusCode)
		return
	}

	// Décodage de la réponse
	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusInternalServerError)
		return
	}

	// Création de la structure de données à renvoyer
	weatherData := WeatherData{
		City:        apiResp.Name,
		Temperature: apiResp.Main.Temp,
		Description: "",
		Humidity:    apiResp.Main.Humidity,
		WindSpeed:   apiResp.Wind.Speed,
		Icon:        "",
	}

	// S'assurer qu'il y a au moins un élément dans le tableau Weather
	if len(apiResp.Weather) > 0 {
		weatherData.Description = apiResp.Weather[0].Description
		weatherData.Icon = apiResp.Weather[0].Icon
	}

	// Envoi de la réponse
	json.NewEncoder(w).Encode(weatherData)
}
