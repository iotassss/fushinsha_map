import { useState } from "react";

interface LatLngSearchFormProps {
  center: [number, number];
  setCenter: (center: [number, number]) => void;
}

export default function LatLngSearchForm({ center, setCenter }: LatLngSearchFormProps) {
  const [latlng, setLatlng] = useState<string>('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // カッコを除去し、カンマ区切り、スペース有無両対応
    const cleaned = latlng.replace(/[()]/g, '');
    const parts = cleaned.split(',').map(s => s.trim());
    if (parts.length === 2) {
      const newLat = parseFloat(parts[0]);
      const newLng = parseFloat(parts[1]);
      if (!isNaN(newLat) && !isNaN(newLng)) {
        setCenter([newLat, newLng]);
      }
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '0.5rem', alignItems: 'center' }}>
        <input
          type="text"
          value={latlng}
          onChange={e => setLatlng(e.target.value)}
          placeholder="35.9344771103889, 139.66427953115772"
          style={{
            width: 260,
            padding: '0.5rem 0.75rem',
            fontSize: '1rem',
            border: '1px solid #ddd',
            borderRadius: 6,
            outline: 'none',
          }}
        />
        <button
          type="submit"
          disabled={!latlng.trim()}
          style={{
            padding: '0.5rem 0.9rem',
            fontSize: '0.95rem',
            borderRadius: 6,
            background: '#1a73e8',
            color: '#fff',
            border: 'none',
            cursor: latlng.trim() ? 'pointer' : 'not-allowed',
          }}
          aria-label="座標に移動"
        >
          座標に移動
        </button>
      </form>
    </div>
  );
}
