'use client';

import 'leaflet/dist/leaflet.css';
import { MapContainer, Marker, Popup, TileLayer, useMap, ZoomControl, useMapEvent } from 'react-leaflet';
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

// 画面全体を覆う黒色透明オーバーレイ
const Overlay = ({ zIndex }: { zIndex: number }) => (
  <div style={{
    position: 'fixed',
    top: 0,
    left: 0,
    width: '100vw',
    height: '100vh',
    backgroundColor: 'rgba(0,0,0,0.5)',
    zIndex: zIndex,
    pointerEvents: 'auto',
  }} />
);

import type { CreatePersonPayload } from "./types/CreatePersonPayload";
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
  const [selectedEmoji, setSelectedEmoji] = useState<string>('');
  // クリック位置のstate（nullならcenterを使う）
  const [clickedPos, setClickedPos] = useState<[number, number] | null>(null);

  // 地図クリック時に座標をアラートするコンポーネント
  // クリック位置のstateとポップアップ表示
  function MapClickHandler() {
    const [popupPos, setPopupPos] = useState<[number, number] | null>(null);
    const [popupMsg, setPopupMsg] = useState<string>('');
    useMapEvent('click', (event) => {
      setPopupPos([event.latlng.lat, event.latlng.lng]);
      setPopupMsg(`${event.latlng.lat}, ${event.latlng.lng}`);
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

  function CreatePersonModal({ latitude, longitude }: { latitude: number, longitude: number }) {
    // 入力state
    const [sign, setSign] = useState('');
    const [gender, setGender] = useState('');
    const [clothing, setClothing] = useState('');
    const [accessories, setAccessories] = useState('');
    const [vehicle, setVehicle] = useState('');
    const [behavior, setBehavior] = useState('');
    const [hairstyle, setHairstyle] = useState('');
    const [sightingTime, setSightingTime] = useState('');
    const [features, setFeatures] = useState('');
    const [categories, setCategories] = useState('');

    const handleSubmit = async () => {
      // TODO: バリデーションは後で実装
      // sightingTimeをISO8601形式に変換（今日の日付を付与）
      let isoSightingTime = '';
      if (sightingTime) {
        const today = new Date();
        const [hh, mm] = sightingTime.split(':');
        today.setHours(Number(hh), Number(mm), 0, 0);
        isoSightingTime = today.toISOString();
      }
      const payload: CreatePersonPayload = {
        latitude,
        longitude,
        emoji: selectedEmoji,
        sign,
        gender,
        clothing,
        accessories,
        vehicle,
        behavior,
        hairstyle,
        sightingTime: isoSightingTime,
        registerUUID: '',
      };
      try {
        await createPerson(payload);
        setShowModal(false);
      } catch (err) {
        alert('送信に失敗しました');
        console.error(err);
      }
    };

    return (
      <>
        <Overlay zIndex={1999}/>
        <div style={{
          position: 'fixed',
          top: 0,
          left: 0,
          width: '100vw',
          height: '100vh',
          zIndex: 2000,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}>
          <div style={{
            background: '#fff',
            borderRadius: 8,
            padding: 32,
            minWidth: 320,
            maxWidth: '90vw',
            maxHeight: '80vh',
            overflowY: 'auto',
            boxShadow: '0 4px 24px rgba(0,0,0,0.18)',
            position: 'relative',
          }}>
            <button
              onClick={() => setShowModal(false)}
              style={{
                position: 'absolute',
                top: 12,
                right: 12,
                background: 'transparent',
                border: 'none',
                fontSize: 24,
                cursor: 'pointer',
                color: '#888',
              }}
              aria-label="閉じる"
            >✕</button>
            <h2 style={{ marginBottom: 16 }}>不審者情報を投稿</h2>
            <form>
              {/* 緯度経度は地図クリック位置から取得する想定。ここではcenterを表示 */}
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>緯度・経度</label>
                <div style={{ color: '#555', fontSize: 14 }}>{latitude}, {longitude}</div>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>顔絵文字</label>
                <div style={{
                  display: 'flex', flexWrap: 'wrap', gap: 4, maxHeight: 120, overflowY: 'auto', border: '1px solid #eee', borderRadius: 4, padding: 4, marginBottom: 8
                }}>
                  {Array.from({ length: 0x1F64A - 0x1F600 + 1 }, (_, i) => 0x1F600 + i).map(code => {
                    const emoji = String.fromCodePoint(code);
                    return (
                      <button
                        type="button"
                        key={code}
                        onClick={() => setSelectedEmoji(emoji)}
                        style={{
                          fontSize: 24,
                          padding: 2,
                          border: selectedEmoji === emoji ? '2px solid #1976d2' : '1px solid #ccc',
                          borderRadius: 4,
                          background: selectedEmoji === emoji ? '#e3f2fd' : '#fff',
                          cursor: 'pointer',
                        }}
                        aria-label={emoji}
                      >{emoji}</button>
                    );
                  })}
                </div>
                {selectedEmoji && <div style={{ marginTop: 4 }}>選択中: <span style={{ fontSize: 20 }}>{selectedEmoji}</span></div>}
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>サイン（1文字）</label>
                <input type="text" maxLength={1} value={sign} onChange={e => setSign(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }} />
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>性別</label>
                <select value={gender} onChange={e => setGender(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }}>
                  <option value="">未選択</option>
                  <option value="男性">男性</option>
                  <option value="女性">女性</option>
                  <option value="不明">不明</option>
                </select>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>服装</label>
                <select value={clothing} onChange={e => setClothing(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }}>
                  <option value="">未選択</option>
                  <option value="スーツ">スーツ</option>
                  <option value="制服">制服</option>
                  <option value="私服">私服</option>
                  <option value="作業着">作業着</option>
                  <option value="その他">その他</option>
                </select>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>アクセサリー</label>
                <select value={accessories} onChange={e => setAccessories(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }}>
                  <option value="">未選択</option>
                  <option value="帽子">帽子</option>
                  <option value="眼鏡">眼鏡</option>
                  <option value="マスク">マスク</option>
                  <option value="バッグ">バッグ</option>
                  <option value="なし">なし</option>
                </select>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>乗り物</label>
                <select value={vehicle} onChange={e => setVehicle(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }}>
                  <option value="">未選択</option>
                  <option value="自転車">自転車</option>
                  <option value="バイク">バイク</option>
                  <option value="自動車">自動車</option>
                  <option value="徒歩">徒歩</option>
                  <option value="その他">その他</option>
                </select>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>挙動</label>
                <select value={behavior} onChange={e => setBehavior(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }}>
                  <option value="">未選択</option>
                  <option value="徘徊">徘徊</option>
                  <option value="大声">大声</option>
                  <option value="暴力">暴力</option>
                  <option value="つきまとい">つきまとい</option>
                  <option value="その他">その他</option>
                </select>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>髪型</label>
                <select value={hairstyle} onChange={e => setHairstyle(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }}>
                  <option value="">未選択</option>
                  <option value="短髪">短髪</option>
                  <option value="長髪">長髪</option>
                  <option value="坊主">坊主</option>
                  <option value="パーマ">パーマ</option>
                  <option value="その他">その他</option>
                </select>
              </div>
              <div style={{ marginBottom: 12 }}>
                <label style={{ display: 'block', marginBottom: 4 }}>目撃時刻</label>
                <input type="time" value={sightingTime} onChange={e => setSightingTime(e.target.value)} style={{ width: '100%', padding: 8, borderRadius: 4, border: '1px solid #ccc' }} />
              </div>
              <button type="button" style={{
                background: '#1976d2',
                color: '#fff',
                border: 'none',
                borderRadius: 4,
                padding: '8px 24px',
                fontWeight: 'bold',
                fontSize: 16,
                cursor: 'pointer',
              }}
              onClick={handleSubmit}
              >送信</button>
            </form>
          </div>
        </div>
      </>
    );
  }

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
          <MapClickHandler />
          <ZoomControl position="bottomright" />
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
          <SearchAreaButton setPersons={setPersons} />
        </MapContainer>
      </div>
      {/* 画面全体を覆う黒色透明オーバーレイ */}
      {/* <Overlay /> */}
      モーダル表示
      {showModal && (
        <CreatePersonModal
          latitude={clickedPos ? clickedPos[0] : center[0]}
          longitude={clickedPos ? clickedPos[1] : center[1]}
        />
      )}
    </div>
  );
}
