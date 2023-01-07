var filterContainer = document.getElementById("filterContainer");
var isFilterContainerOpen = false;
var artistName;
var creationDate;
var firstAlbum;
var members;
var concertDates;
var concertLocations;
var imageSrc;


// Filters Drawer
showFilters = function () {
    if(isFilterContainerOpen) {
        filterContainer.style.display = "none"
    } else {
        filterContainer.style.display = "block"
    }
    isFilterContainerOpen = !isFilterContainerOpen;
}
// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
  if (event.target == filterContainer) {
    filterContainer.style.display = "none";
  }
}

//Modal Credits
var modal = document.getElementById("modalCredits");
var button= document.getElementById("creditsButton");
var span = document.getElementsByClassName("close")[0];

button.onclick = function() {
  modal.style.display = "block";
}

span.onclick = function() {
  modal.style.display = "none";
}

// Modal Card
var modalCard = document.getElementById("modalCard")
var spanCard = document.getElementsByClassName("close-card")[0];

spanCard.onclick = function() {
  modalCard.style.display = "none";
}

openModal = function(id) {
  //   setTimeout(function () {
  //     map.setZoom(1);
  //     window.dispatchEvent(new Event('resize'));
  // }, 1000);
  document.getElementById(`modal_${id}`).submit()
}

// Map
var locations = [
  {
    Name: 'Paris, France',
    Latitude: 48.856614,
    Longitude: 2.3522219,
    Type: 'city',
    Importance: 0.9
  },
  {
    Name: 'New York, USA',
    Latitude: 40.7127753,
    Longitude: -74.0059728,
    Type: 'city',
    Importance: 0.9
  },
  {
    Name: 'Tokyo, Japan',
    Latitude: 35.6894875,
    Longitude: 139.6917064,
    Type: 'city',
    Importance: 0.9
  }
];

var map = L.map('map').setView([51.505, -0.09], 13);
L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 6,
    minZoom: 1,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);

locations.forEach(function(location) {
  var marker = L.marker([location.Latitude, location.Longitude]).addTo(map);
  marker.bindTooltip(location.Name, {permanent:false, direction:`right`});
});
