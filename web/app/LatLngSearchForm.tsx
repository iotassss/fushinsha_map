import { useState } from "react";

interface LatLngSearchFormProps {
  center: [number, number];
  setCenter: (center: [number, number]) => void;
}

export default function LatLngSearchForm({ center, setCenter }: LatLngSearchFormProps) {
  const [latlng, setLatlng] = useState<string>(center.join(', '));

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // カンマ区切り、スペース有無両対応
    const parts = latlng.split(',').map(s => s.trim());
    if (parts.length === 2) {
      const newLat = parseFloat(parts[0]);
      const newLng = parseFloat(parts[1]);
      if (!isNaN(newLat) && !isNaN(newLng)) {
        setCenter([newLat, newLng]);
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: '1rem' }}>
      <label>
        緯度,経度:
        <input
          type="text"
          value={latlng}
          onChange={e => setLatlng(e.target.value)}
          placeholder="35.9344771103889, 139.66427953115772"
          style={{ marginRight: '1rem', width: '300px' }}
        />
      </label>
      <button type="submit">座標に移動</button>
    </form>
  );
}
