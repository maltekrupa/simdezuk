document.addEventListener("DOMContentLoaded", function(event) {

  var map = L.map('map').setView([50.119, 8.68211], 12);
  L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
    maxZoom: 18
  }).addTo(map);
  console.log("DOM fully loaded and parsed. Map attached");

  realtime = L.realtime({
    type: 'json'
  }, {
    start: false
  }).addTo(map);

  var ws = new WebSocket("ws://localhost:8080/ws");
  
  ws.onmessage = function(e) {
    console.log(e.data);
    var data = JSON.parse(e.data);
    realtime.update(data);
    map.fitBounds(realtime.getBounds(), {maxZoom: 17});
  }; 


//   var popup = L.popup();
//   
//   function DisplayNewPosition(lat,lng) {     
//       if (typeof marker != 'undefined') {
//           map.removeLayer(marker);  // delete previous marker
//           marker = L.marker([lat,lng]).addTo(map);  // add new marker
//       }
//       else {
//           marker = L.marker([lat,lng]).addTo(map);  // add new marker
//       }
//   }
//   

});
