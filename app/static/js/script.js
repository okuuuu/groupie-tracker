var artistName;
var creationDate;
var firstAlbum;
var members;
var concertDates;
var concertLocations;
var imageSrc;
var coordinates;

// Modal Credits
var modal = document.getElementById("modalCredits");
var button = document.getElementById("creditsButton");
var span = document.getElementsByClassName("close")[0];

button.onclick = function () {
    modal.style.display = "block";
}

span.onclick = function () {
    modal.style.display = "none";
}

// Modal Card
var modalCard = document.getElementById("modalCard");
var spanCard = document.getElementsByClassName("close-card")[0];

spanCard.onclick = function () {
    modalCard.style.display = "none";
}

const form = document.querySelector('form');

form.addEventListener('submit', (event) => {
    event.preventDefault();

    // submit form
});

openModal = async function (id) {
    let response = await fetch(`http://localhost:8080/coordinates/${id}`);
    let data = await response.json();
    coordinates = data;
    sessionStorage.setItem('coordinates', JSON.stringify(coordinates));

    document.getElementById(`modal_${id}`).submit()
}

var locations = JSON.parse(sessionStorage.getItem('coordinates'));

var map = L.map('map').setView([
    51.505, -0.09
], 13).setZoom(1);

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 6,
    minZoom: 1,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);

locations.forEach(function (location) {
    var marker = L.marker([location.lat, location.lng]).addTo(map);
    marker.bindTooltip(location.location, {
        permanent: false,
        direction: `right`
    });

});
