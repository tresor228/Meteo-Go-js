document.addEventListener('DOMContentLoaded', function() {
    // Éléments DOM
    const cityInput = document.getElementById('city-input');
    const searchButton = document.getElementById('search-button');
    const cityName = document.getElementById('city-name');
    const weatherIcon = document.getElementById('weather-icon');
    const temperature = document.getElementById('temperature');
    const description = document.getElementById('description');
    const humidity = document.getElementById('humidity');
    const windSpeed = document.getElementById('wind-speed');
    const errorMessage = document.getElementById('error-message');
    const weatherContainer = document.getElementById('weather-container');

    // Fonction pour récupérer la météo
    async function getWeather(city) {
        try {
            weatherContainer.classList.add('loading');
            errorMessage.textContent = '';
            
            // Appel à notre API backend
            const response = await fetch(`/api/weather?city=${encodeURIComponent(city)}`);
            
            if (!response.ok) {
                throw new Error(`Erreur: ${response.status}`);
            }
            
            const data = await response.json();
            
            // Mise à jour de l'interface utilisateur
            cityName.textContent = data.city;
            temperature.textContent = Math.round(data.temperature);
            description.textContent = data.description;
            humidity.textContent = data.humidity;
            windSpeed.textContent = (data.windSpeed * 3.6).toFixed(1); // Conversion de m/s en km/h
            
            // Mise à jour de l'icône météo
            if (data.icon) {
                weatherIcon.src = `https://openweathermap.org/img/wn/${data.icon}@2x.png`;
                weatherIcon.alt = data.description;
            }
            
            weatherContainer.classList.remove('loading');
        } catch (error) {
            console.error('Erreur lors de la récupération des données météo:', error);
            errorMessage.textContent = 'Impossible de récupérer les données météo. Veuillez vérifier le nom de la ville et réessayer.';
            weatherContainer.classList.remove('loading');
        }
    }

    // Gestionnaire d'événement pour le bouton de recherche
    searchButton.addEventListener('click', function() {
        const city = cityInput.value.trim();
        if (city) {
            getWeather(city);
        } else {
            errorMessage.textContent = 'Veuillez entrer un nom de ville';
        }
    });

    // Gestionnaire d'événement pour la touche Entrée
    cityInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            const city = cityInput.value.trim();
            if (city) {
                getWeather(city);
            } else {
                errorMessage.textContent = 'Veuillez entrer un nom de ville';
            }
        }
    });

    // Charger la météo par défaut au chargement de la page
    getWeather('Lomé');
});