<!DOCTYPE html>
<html>
  <head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <style type="text/css">
      html, body, #map-canvas { height: 100%; margin: 0; padding: 0;}
    </style>
    <script type="text/javascript"
      src="https://maps.googleapis.com/maps/api/js?key=AIzaSyDJPIZ6eYnHvslQk8FAVf3ea1eaAbVJqkw">
    </script>
   <script type="text/javascript" src="http://google-maps-utility-library-v3.googlecode.com/svn/tags/markerclusterer/1.0/src/markerclusterer.js"></script>
    <script type="text/javascript">
      var map, mc;
      
       function refreshMap() {
        if (markerClusterer) {
            markerClusterer.clearMarkers();
        }

        var markers = [];

        var markerImage = new google.maps.MarkerImage(imageUrl, new google.maps.Size(24, 32));

        for (var i = 0; i < 1000; ++i) {
            var latLng = new google.maps.LatLng(data.photos[i].latitude, data.photos[i].longitude)
            var marker = new google.maps.Marker({
                position: latLng,
                draggable: true,
                icon: markerImage
            });
            google.maps.event.addListener(marker, 'click', function() {
                var infowindow = new google.maps.InfoWindow({
                    content: 'Information!'
                });
                infowindow.open(map, this);
            });
            markers.push(marker);
        }

        var zoom = parseInt(document.getElementById('zoom').value, 10);
        var size = parseInt(document.getElementById('size').value, 10);
        var style = parseInt(document.getElementById('style').value, 10);
        zoom = zoom == -1 ? null : zoom;
        size = size == -1 ? null : size;
        style = style == -1 ? null : style;

        markerClusterer = new MarkerClusterer(map, markers, {
            maxZoom: zoom,
            gridSize: size,
            styles: styles[style]
        });
    }

      
       function initialize() {
        map = new google.maps.Map(document.getElementById('map'), {
            zoom: 2,
            center: new google.maps.LatLng(39.91, 116.38),
            mapTypeId: google.maps.MapTypeId.ROADMAP
        });

        var refresh = document.getElementById('refresh');
        google.maps.event.addDomListener(refresh, 'click', refreshMap);

        var clear = document.getElementById('clear');
        google.maps.event.addDomListener(clear, 'click', clearClusters);

        refreshMap();
      }
     


      google.maps.event.addDomListener(window, 'load', initialize);
    </script>
  </head>
  <body>

<div id="map-canvas"></div>
  </body>
</html>