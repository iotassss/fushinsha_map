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

  // 地図クリック時に座標をアラートするコンポーネント
  // クリック位置のstateとポップアップ表示
  function MapClickHandler() {
    const [popupPos, setPopupPos] = useState<[number, number] | null>(null);
    const [popupMsg, setPopupMsg] = useState<string>('');
    useMapEvent('click', (event) => {
      const lat = Math.floor(event.latlng.lat * 10000) / 10000;
      const lng = Math.floor(event.latlng.lng * 10000) / 10000;
      setPopupPos([lat, lng]);
      setPopupMsg(`${lat}, ${lng}`);
    });
    return (
      <>
        {popupPos && (
          <Popup position={popupPos} eventHandlers={{ popupclose: () => setPopupPos(null) }}>
            <div>
              <a
                href={`https://www.google.com/maps?q=${popupPos[0]},${popupPos[1]}`}
                target="_blank"
                rel="noopener noreferrer"
                style={{ color: '#1976d2', textDecoration: 'underline' }}
              >
                {popupMsg}
              </a>
              <div
                style={{ marginTop: '8px', fontWeight: 'bold', cursor: 'pointer', color: '#d32f2f' }}
                onClick={() => {
                  setClickedPos(popupPos);
                  setShowModal(true);
                }}
              >
                👇️ここに不審者情報を投稿する
              </div>
            </div>
          </Popup>
        )}
      </>
    );
  }

  // person詳細ボタン（ダミー）
  const handleButtonClick = async (personSummary: PersonSummary) => {
    const person = await getPerson(personSummary.uuid);
    console.log('Person:', person);
    setSelectedPerson(person.person);
    setIsPanelOpen(true);
  };

  useEffect(() => {
    console.log('Persons data updated:', persons);
    persons.forEach(person => {
      console.log(`Person UUID: ${person.uuid}, Location: (${person.latitude}, ${person.longitude}), Emoji: ${person.emoji}, Sign: ${person.sign}, SightingCount: ${person.sighting_count}`);
    });
  }, [persons]);

  return (
    <div style={{ position: 'relative' }}>
      <LeftSidePanel isPanelOpen={isPanelOpen} setIsPanelOpen={setIsPanelOpen} selectedPerson={selectedPerson} />
      {/* 既存の地図部分 */}
      <div>
        <MapContainer
          center={center}
          zoom={13}
          scrollWheelZoom={true}
          touchZoom={true}
          wheelDebounceTime={10}
          zoomControl={false}
        >
          <MapClickHandler />
          <ZoomControl position="bottomright" />
          <ChangeMapCenter center={center} />
          <TileLayer
            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          />
          {/* Emojiを各personの位置に表示 */}
          {persons.map(person => (
            <EmojiMarker
              key={person.uuid}
              person={person}
              handleButtonClick={handleButtonClick}
              L={L}
            />
          ))}
          {/* 中心点のマーカーはそのまま残す場合 */}
          <Marker position={center}>
            <Popup>
              A pretty CSS3 popup. <br /> Easily customizable.
            </Popup>
          </Marker>
          <GetMapInstance setMapInstance={setMapInstance} />
        </MapContainer>
        {mapInstance && <SearchAreaButton setPersons={setPersons} map={mapInstance} />}
      </div>
      {/* 画面全体を覆う黒色透明オーバーレイ */}
      {/* <Backdrop /> */}
      モーダル表示
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
