'use client';

import 'leaflet/dist/leaflet.css';
import { MapContainer, Marker, Popup, TileLayer, useMap, ZoomControl, useMapEvent } from 'react-leaflet';
import L, { map } from 'leaflet';
import SearchAreaButton from './SearchAreaButton';
import { useEffect, useState } from 'react';
import type { GetPersonsResponse, PersonSummary } from './types/Persons';
import './initLeaflet';
import './Map.css';
import { GetPersonResponse, Person } from './types/Person';
import type { CreatePersonPayload } from "./types/CreatePersonPayload";
import { LeftSidePanel } from './LeftSidePanel';
import { CreatePersonModal } from './CreatePersonModel';
import { EmojiMarker } from './EmojiMarker';
import { CenterMarker } from './CenterMarker';
import { MapClickHandler } from './MapClickHandler';

// centerが変わったら地図を移動するコンポーネント
function ChangeMapCenter({ center }: { center: [number, number] }) {
  const map = useMap();
  useEffect(() => {
    map.setView(center);
  }, [center, map]);
  return null;
}

const GetMapInstance = (
  { setMapInstance }: { setMapInstance: (map: L.Map) => void }
) => {
  const map = useMap();
  useEffect(() => {
    console.log('Map instance:', map);
    setMapInstance(map);
  }, [map, setMapInstance]);
  return null;
};

export interface MapProps {
  center: [number, number];
  getPersons: (uuid: string) => Promise<GetPersonsResponse>;
  getPerson: (uuid: string) => Promise<GetPersonResponse>;
  createPerson: (payload: CreatePersonPayload) => Promise<void>;
}

export default function Map({ center, getPerson, createPerson }: MapProps) {
  // パネルの開閉状態
  const [isPanelOpen, setIsPanelOpen] = useState(false);
  // 地図・person関連のstate
  const [persons, setPersons] = useState<PersonSummary[]>([]);
  // person詳細表示用
  const [selectedPerson, setSelectedPerson] = useState<Person | null>(null);
  const [showModal, setShowModal] = useState(false);
  // クリック位置のstate（nullならcenterを使う）
  const [clickedPos, setClickedPos] = useState<[number, number] | null>(null);
  const [mapInstance, setMapInstance] = useState<L.Map | null>(null);

  useEffect(() => {
    console.log('Persons data updated:', persons);
    persons.forEach(person => {
      console.log(`Person UUID: ${person.uuid}, Location: (${person.latitude}, ${person.longitude}), Emoji: ${person.emoji}, Sign: ${person.sign}, SightingCount: ${person.sighting_count}`);
    });
  }, [persons]);

  return (
    <div style={{ position: 'relative' }}>
      <LeftSidePanel isPanelOpen={isPanelOpen} setIsPanelOpen={setIsPanelOpen} selectedPerson={selectedPerson} />
      <div>
        <MapContainer
          center={center}
          zoom={13}
          scrollWheelZoom={true}
          touchZoom={true}
          wheelDebounceTime={10}
          zoomControl={false}
        >
          <MapClickHandler
            setClickedPos={setClickedPos}
            setShowModal={setShowModal}
          />
          <ZoomControl position="bottomright" />
          <ChangeMapCenter center={center} />
          <TileLayer
            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          />
          {persons.map(person => (
            <EmojiMarker
              key={person.uuid}
              person={person}
              L={L}
              getPerson={getPerson}
              setSelectedPerson={setSelectedPerson}
              setIsPanelOpen={setIsPanelOpen}
            />
          ))}
          <CenterMarker center={center} />
          <GetMapInstance setMapInstance={setMapInstance} />
        </MapContainer>
        {mapInstance && <SearchAreaButton setPersons={setPersons} map={mapInstance} />}
      </div>
      {showModal && (
        <CreatePersonModal
          latitude={clickedPos ? clickedPos[0] : center[0]}
          longitude={clickedPos ? clickedPos[1] : center[1]}
          createPerson={createPerson}
          setShowModal={setShowModal}
        />
      )}
    </div>
  );
}
