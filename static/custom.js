document.addEventListener("DOMContentLoaded", function(event) {

  var map = L.map('map').setView([50.119, 8.68211], 12);
  L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
    maxZoom: 18
  }).addTo(map);
  console.log("DOM fully loaded and parsed. Map attached");
  var popup = L.popup();
  
  function DisplayNewPosition(lat,lng) {     
      if (typeof marker != 'undefined') {
          map.removeLayer(marker);  // delete previous marker
          marker = L.marker([lat,lng]).addTo(map);  // add new marker
      }
      else {
          marker = L.marker([lat,lng]).addTo(map);  // add new marker
      }
  }
  
  var ws = new WebSocket("ws://localhost:8080/ws");
  
  ws.onmessage = function(e) {
      if (typeof e.data === "string"){
          try {
              var object = JSON.parse(e.data);
              lat=object.Lat
              lon=object.Lon
              console.log("Lat: " + lat +", Lon: "+lon);
              DisplayNewPosition(lat,lon);
          }
          catch(err) {
              console.log("String message received (but not JSON): ", e.data)
          }
      } 
      else {
          console.log("Other message received (not string)", e.data);
      }
  }; 

});
