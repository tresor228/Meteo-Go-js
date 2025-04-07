# Application Météo

Une application météo simple avec un backend en Go et un frontend en HTML, CSS et JavaScript.

## Configuration

1. Assurez-vous que Go est installé sur votre système
2. Clonez ce dépôt
3. Créez un fichier .env à la racine du projet avec les variables suivantes:

   TON_API_KEY=e78a6edb3e54c39b2d6a58142763244d
   PORT=8080

## Installation des dépendances (si vous ne l'avez pas)

go mod download

## Lancement de l'application

go run main.go

L'application sera disponible à l'adresse http://localhost:8080.

## Fonctionnalités

- Recherche de la météo par nom de ville
- Affichage de la température actuelle
- Affichage de la description météo
- Affichage du taux d'humidité
- Affichage de la vitesse du vent