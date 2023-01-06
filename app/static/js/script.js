var filterContainer = document.getElementById("filterContainer");
var isFilterContainerOpen = false;
var artistName;
var creationDate;
var firstAlbum;
var members;
var concertDates;
var concertLocations;
var imageSrc;

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


// Get the modal
var modal = document.getElementById("myModal");

// Get the button that opens the modal
var btn = document.getElementById("myBtn");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks on the button, open the modal
btn.onclick = function() {
  modal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
}

const modalCard = document.getElementById("modalCard")
const button = document.querySelector("button")
var span = document.getElementsByClassName("close-card")[0];

span.onclick = function() {
  modalCard.style.display = "none";
}

openModal = function(id) {
  artistName = document.getElementById(`artistName_${id}`).innerHTML;
  creationDate = document.getElementById(`creationDate_${id}`).innerHTML;
  firstAlbum = document.getElementById(`firstAlbum_${id}`).innerHTML;
  members = document.getElementById(`members_${id}`).innerHTML;
  dates = document.getElementById(`dates_${id}`).innerHTML;
  locations = document.getElementById(`locations_${id}`).innerHTML;
  imageSrc = document.getElementById(`discographyImage_${id}`).src;

  document.getElementById('artistNameModal').innerHTML = artistName;
  document.getElementById('creationDateModal').innerHTML = creationDate;
  document.getElementById('firstAlbumModal').innerHTML = firstAlbum;
  document.getElementById('membersModal').innerHTML = members;
  document.getElementById('datesModal').innerHTML = dates;
  document.getElementById('locationsModal').innerHTML = locations;
  document.getElementById('imageModal').src = imageSrc;

  
  modalCard.style.display = "block";
}



