'use client';

import 'leaflet/dist/leaflet.css';
import { MapContainer, Marker, Popup, TileLayer, useMap } from 'react-leaflet';
import L from 'leaflet';
import SearchAreaButton from './SearchAreaButton';
import { useEffect } from 'react';
import { useState } from 'react';

import type { PersonSummary } from './types/PersonSummary';
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
  const [persons, setPersons] = useState<PersonSummary[]>([]);

  useEffect(() => {
    console.log('Persons data updated:', persons);
    persons.forEach(person => {
      console.log(`Person UUID: ${person.uuid}, Location: (${person.latitude}, ${person.longitude}), Emoji: ${person.emoji}, Sign: ${person.sign}, SightingCount: ${person.sightingCount}`);
    });
  }, [persons]);

  return (
    <div>
      <MapContainer
        center={center}
        zoom={13}
        scrollWheelZoom={true}
        touchZoom={true}
        wheelDebounceTime={10}
        zoomControl={false}
      >
        <ZoomControl position="bottomright" />
<SearchAreaButton setPersons={setPersons} />
        <ChangeMapCenter center={center} />
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
{/* Emojiを各personの位置に表示 */}
        {persons.map(person => (
          <Marker
            key={person.uuid}
            position={[person.latitude, person.longitude]}
            icon={L.divIcon({
              className: 'emoji-marker',
              html: `<span style="font-size: 2rem;">${person.emoji}</span>`
            })}
          >
            <Popup>
              <div>
                <div>Sign: {person.sign}</div>
                <div>SightingCount: {person.sightingCount}</div>
              </div>
            </Popup>
          </Marker>
        ))}
        {/* 中心点のマーカーはそのまま残す場合 */}
        <Marker position={center}>
          <Popup>
            A pretty CSS3 popup. <br /> Easily customizable.
          </Popup>
        </Marker>
      </MapContainer>
    </div>
  );
}
