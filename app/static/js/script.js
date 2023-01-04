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