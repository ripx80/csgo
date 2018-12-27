{{define "result"}}
<div class="row">
    <div class="col-md-5"></div>
    <div class="col-md-2 no-bg-img result">
		<div class="text-center"></div>
            {{ if .Result }}
                <div class="row">
                    <div class="col-md-offset-2 col-md-3">Cell:</div>
                    <div class="col-md-offset-1 col-md-3"><strong>{{.CellID}}</strong></div>
                </div>
                <div class="row">
                    <div class="col-md-offset-2 col-md-3">LAC:</div>
                    <div class="col-md-offset-1 col-md-3"><strong>{{.LacID}}</strong></div>
                </div>                    
                {{ range .Result }}
                <div class="row">
                    <div class="col-md-offset-2 col-md-3">MCC:</div>
                    <div class="col-md-offset-1 col-md-3"><strong>{{.Mcc}}</strong></div>
                </div>                         
                <div class="row">
                    <div class="col-md-offset-2 col-md-3">Longtitude:</div>
                    <div class="col-md-offset-1 col-md-3"><strong id="lon">{{.Lon}}</strong></div>
                </div>
                <div class="row">
                    <div class="col-md-offset-2 col-md-3">Latitude:</div>
                    <div class="col-md-offset-1 col-md-3"><strong id="lat">{{.Lat}}</strong></div>
                </div>
                <script src="/js/OpenLayers-2.13.1/OpenLayers.js"></script>                
                <script>
                    function render_map(){
                        map = new OpenLayers.Map("map");
                        map.addLayer(new OpenLayers.Layer.OSM());
                        var lonLat = new OpenLayers.LonLat($("#lon").html() ,$("#lat").html())            
                          .transform(
                            new OpenLayers.Projection("EPSG:4326"), // transform from WGS 1984
                            map.getProjectionObject() // to Spherical Mercator Projection
                          );          
                        var zoom=13;
                        var markers = new OpenLayers.Layer.Markers( "Markers" );
                        map.addLayer(markers);    
                        markers.addMarker(new OpenLayers.Marker(lonLat));    
                        map.setCenter (lonLat, zoom); 
                        $("html, body").scrollTop($("#map").offset().top);
                };
            </script>
            
                {{end}}
            {{ else }}
                <p>No Entry found</p>	
            {{ end }}
            </div>
        </div>
    </div>    
</div>
 <div class="row-fluid" style="margin-top:15px;">
        <div id="map" class="map"></div>
</div>

{{end}}
