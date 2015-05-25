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
  var polylines = [];
  
  ws.onmessage = function(e) {
    clearMap();
    var data = JSON.parse(e.data);
    realtime.update(data);
    var polyline = L.polyline(data.properties.history, {color: 'red', weight: 2}).addTo(map);
    polylines.push(polyline);
    map.fitBounds(realtime.getBounds(), {maxZoom: 18});
  }; 

  function clearMap() {
    for(i = 0; i < polylines.length; i++) {
        map.removeLayer(polylines[i]);
    }
  }

});
