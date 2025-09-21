import { useState } from "react";
import { CreatePersonPayload } from "./types/CreatePersonPayload";
import { Backdrop } from "./Backdrop";

export function CreatePersonModal({
  latitude,
  longitude,
  createPerson,
  setShowModal,
}: {
  latitude: number,
  longitude: number,
  createPerson: (payload: CreatePersonPayload) => Promise<void>,
  setShowModal: (show: boolean) => void,
}) {
  // 入力state
  const [sign, setSign] = useState('');
  const [gender, setGender] = useState('');
  const [clothing, setClothing] = useState('');
  const [accessories, setAccessories] = useState('');
  const [vehicle, setVehicle] = useState('');
  const [behavior, setBehavior] = useState('');
  const [hairstyle, setHairstyle] = useState('');
  const [sightingTime, setSightingTime] = useState('');
  const [selectedEmoji, setSelectedEmoji] = useState<string>('');

  const handleSubmit = async () => {
    // TODO: バリデーションは後で実装

    // sightingTimeをISO8601形式に変換（今日の日付を付与）
    let isoSightingTime = '';
    if (sightingTime) {
      // sightingTimeはJST（日本標準時）で入力されている前提
      const now = new Date();
      const [hh, mm] = sightingTime.split(':');
      // JSTの年月日を取得
      const year = now.getFullYear();
      const month = now.getMonth() + 1;
      const day = now.getDate();
      isoSightingTime = `${year.toString().padStart(4, '0')}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}T${hh.padStart(2, '0')}:${mm.padStart(2, '0')}:00+09:00`;
    }
    console.log(isoSightingTime)
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
      <Backdrop zIndex={1999} />
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
