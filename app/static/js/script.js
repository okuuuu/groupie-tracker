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
var modalCard = document.getElementById("modalCard");
var spanCard = document.getElementsByClassName("close-card")[0];

spanCard.onclick = function() {
  modalCard.style.display = "none";
}

openModal = function(id) {
  document.getElementById(`modal_${id}`).submit()
}

// Map
var locations = document.getElementById("locations");
var map = L.map('map').setView([51.505, -0.09], 13).setZoom(1);

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 6,
    minZoom: 1,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);

locations.forEach(function(location) {
  var marker = L.marker([location.Latitude, location.Longitude]).addTo(map);
  marker.bindTooltip(location.Name, {permanent:false, direction:`right`});
});