import Leaflet from 'leaflet';

const DefaultIcon = Leaflet.icon({
  iconUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png',
  iconRetinaUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon-2x.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
  iconSize: [25, 41],
  iconAnchor: [12, 41], // アイコンのオフセット
  popupAnchor: [0, -32], // ポップアップのオフセット
  shadowSize: [41, 41]
});
Leaflet.Marker.prototype.options.icon = DefaultIcon;
