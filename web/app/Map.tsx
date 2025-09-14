'use client';

import 'leaflet/dist/leaflet.css';
import { MapContainer, Marker, Popup, TileLayer, useMap } from 'react-leaflet';
import { useEffect } from 'react';
import { useState } from 'react';
import './initLeaflet';
import './Map.css';

// centerが変わったら地図を移動するコンポーネント
function ChangeMapCenter({ center }: { center: [number, number] }) {
  const map = useMap();
  useEffect(() => {
    map.setView(center);
  }, [center, map]);
  return null;
}

export interface MapProps {
  center: [number, number];
}

export default function Map({ center }: MapProps) {
  return (
    <div>
      <MapContainer
        center={center}
        zoom={13}
        scrollWheelZoom={true}
        touchZoom={true}
        wheelDebounceTime={10}
      >
        <ChangeMapCenter center={center} />
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        <Marker position={center}>
          <Popup>
            A pretty CSS3 popup. <br /> Easily customizable.
          </Popup>
        </Marker>
      </MapContainer>
    </div>
  );
}
