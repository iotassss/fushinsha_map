'use client';

import 'leaflet/dist/leaflet.css';
import { MapContainer, Marker, Popup, TileLayer, useMap, ZoomControl } from 'react-leaflet';
import L from 'leaflet';
import SearchAreaButton from './SearchAreaButton';
import { useEffect, useState } from 'react';

import type { GetPersonsResponse, PersonSummary } from './types/Persons';
import './initLeaflet';
import './Map.css';
import { GetPersonResponse, Person } from './types/Person';

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
  getPersons: (uuid: string) => Promise<GetPersonsResponse>;
  getPerson: (uuid: string) => Promise<GetPersonResponse>;
}


export default function Map({ center, getPerson }: MapProps) {
  // パネルの開閉状態
  const [isPanelOpen, setIsPanelOpen] = useState(false);

  // 地図・person関連のstate
  const [persons, setPersons] = useState<PersonSummary[]>([]);
  // person詳細表示用
  const [selectedPerson, setSelectedPerson] = useState<Person | null>(null);

  // パネルのつまみクリックで開閉
  const handlePanelToggle = () => {
    setIsPanelOpen(open => !open);
  };

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
      {/* サイドパネルとつまみ */}
      <div style={{
        position: 'fixed',
        top: 0,
        left: isPanelOpen ? 0 : -260,
        width: 260,
        height: '100vh',
        background: '#fff',
        boxShadow: isPanelOpen ? '2px 0 8px rgba(0,0,0,0.15)' : 'none',
        zIndex: 1000,
        transition: 'left 0.2s',
        display: 'flex',
        alignItems: 'flex-start',
      }}>
        {/* つまみボタン */}
        <button
          onClick={handlePanelToggle}
          style={{
            position: 'fixed',
            left: isPanelOpen ? 260 : 0,
            top: '50%',
            transform: 'translateY(-50%)',
            width: 32,
            height: 64,
            borderRadius: '0 8px 8px 0',
            border: '1px solid #ccc',
            background: '#fafafa',
            boxShadow: '1px 0 4px rgba(0,0,0,0.08)',
            cursor: 'pointer',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontSize: '1.5rem',
            zIndex: 1100,
            padding: 0,
            transition: 'left 0.2s, border-radius 0.2s',
          }}
          aria-label={isPanelOpen ? 'パネルを閉じる' : 'パネルを開く'}
        >
          {isPanelOpen ? '←' : '→'}
        </button>
        {/* パネル中身 */}
        <div style={{ padding: '32px 16px', width: '100%', position: 'relative' }}>
          {/* 閉じるボタン */}
          <button
            onClick={() => setIsPanelOpen(false)}
            style={{
              position: 'absolute',
              top: 8,
              right: 8,
              width: 32,
              height: 32,
              border: 'none',
              background: 'transparent',
              fontSize: '1.5rem',
              cursor: 'pointer',
              zIndex: 1200,
              lineHeight: 1,
              color: '#777',
              transition: 'color 0.15s',
            }}
            aria-label="閉じる"
            onMouseOver={e => (e.currentTarget.style.color = '#888')}
            onMouseOut={e => (e.currentTarget.style.color = '#bbb')}
          >✕</button>
          {selectedPerson ? (
            <>
              <div style={{ fontSize: '2.5rem', textAlign: 'center', marginBottom: 8 }}>{selectedPerson.emoji}</div>
              <div style={{ fontSize: '1.5rem', fontWeight: 'bold', textAlign: 'center', marginBottom: 16 }}>{selectedPerson.sign}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>UUID:</span>{selectedPerson.uuid}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>緯度:</span>{selectedPerson.latitude}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>経度:</span>{selectedPerson.longitude}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>目撃数:</span>{selectedPerson.sighting_count}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>目撃時刻:</span>{Array.isArray(selectedPerson.sighting_times) ? selectedPerson.sighting_times.join(', ') : ''}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>カテゴリ:</span>{Array.isArray(selectedPerson.categories) ? selectedPerson.categories.join(', ') : ''}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>性別:</span>{selectedPerson.gender}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>服装:</span>{selectedPerson.clothing}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>アクセサリー:</span>{selectedPerson.accessories}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>乗り物:</span>{selectedPerson.vehicle}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>行動:</span>{selectedPerson.behavior}</div>
              <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>髪型:</span>{selectedPerson.hairstyle}</div>
            </>
          ) : (
            <>
              <div style={{ fontWeight: 'bold', fontSize: '1.2rem' }}>サイドパネル</div>
              <div style={{ marginTop: 16, color: '#888' }}>ここに詳細情報などを表示できます</div>
            </>
          )}
        </div>
      </div>

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
                html: `<span style=\"font-size: 2rem;\">${person.emoji}</span>`
              })}
            >
              <Popup>
                <div>
                  <div>{person.emoji}サイン:  {person.sign}</div>
                  <div>👀目撃数: {person.sighting_count}</div>
                </div>
                <div style={{ marginTop: '8px' }}>
                  <button style={{ cursor: 'pointer' , fontWeight: 'bold' }} onClick={() => handleButtonClick(person)}>👉️詳細を見る</button>
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
    </div>
  );
}
