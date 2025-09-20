import { Person } from "./types/Person";

export function LeftSidePanel({
  isPanelOpen,
  setIsPanelOpen,
  selectedPerson
}: {
  isPanelOpen: boolean,
  setIsPanelOpen: React.Dispatch<React.SetStateAction<boolean>>,
  selectedPerson: Person | null
}) {
  // パネルのつまみクリックで開閉
  const handlePanelToggle = () => {
    setIsPanelOpen(open => !open);
  };

  return (
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
            <div style={{ marginBottom: 8 }}><span style={{ fontWeight: 'bold', marginRight: 8 }}>目撃時刻:</span>{selectedPerson.sighting_time}</div>
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
  );
}
