import { useState } from "react";
import { Popup, useMapEvent } from "react-leaflet";

// 地図クリック時に座標をアラートするコンポーネント
// クリック位置のstateとポップアップ表示
export function MapClickHandler({
  setClickedPos,
  setShowModal
}: {
  setClickedPos: (pos: [number, number]) => void,
  setShowModal: (show: boolean) => void
}) {
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
