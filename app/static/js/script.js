var filterContainer = document.getElementById("filterContainer");
var isFilterContainerOpen = false;

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