import { Marker, Popup } from "react-leaflet";

export function CenterMarker({ center }: { center: [number, number] }) {
  return (
    <Marker position={center}>
      <Popup>
        A pretty CSS3 popup. <br /> Easily customizable.
      </Popup>
    </Marker>
  );
}
