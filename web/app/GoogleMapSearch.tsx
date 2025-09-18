import { useCallback, useMemo, useState, FormEvent } from 'react';

type GoogleMapSearchProps = {
  /** 近傍にバイアスしたい場合だけ渡す（Leaflet の Map インスタンス） */
  map?: L.Map;
  /** 初期値 */
  defaultQuery?: string;
  /** 新規タブで開く（既定:true） */
  newTab?: boolean;
};

export function GoogleMapSearch({ map, defaultQuery = '', newTab = true }: GoogleMapSearchProps) {
  const [q, setQ] = useState(defaultQuery);

  const centerAndZoom = useMemo(() => {
    if (!map) return null;
    const c = map.getCenter();
    const z = map.getZoom();
    // Google マップの z は 3–21 あたりを想定
    const gz = Math.min(21, Math.max(3, Math.round(z)));
    return { lat: c.lat, lng: c.lng, z: gz };
  }, [map]);

  const buildUrl = useCallback(
    (query: string) => {
      const trimmed = query.trim();
      if (!trimmed) return null;

      // 公式の v=1 URL: https://www.google.com/maps/search/?api=1&query=...
      // 位置バイアスを掛けたい時はパス形式で中心/ズームを付けるのが実用的。
      if (centerAndZoom) {
        return `https://www.google.com/maps/search/${encodeURIComponent(trimmed)}/@${centerAndZoom.lat},${centerAndZoom.lng},${centerAndZoom.z}z`;
      }
      return `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(trimmed)}`;
    },
    [centerAndZoom]
  );

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    const url = buildUrl(q);
    if (!url) return;
    if (newTab) {
      window.open(url, '_blank', 'noopener,noreferrer');
    } else {
      window.location.href = url;
    }
  };

  return (
    <div>
      <form
        onSubmit={onSubmit}
        style={{ display: 'flex', gap: '0.5rem', alignItems: 'center' }}
      >
        <input
          type="text"
          value={q}
          onChange={(e) => setQ(e.target.value)}
          placeholder="例）東京駅 "
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
          disabled={!q.trim()}
          style={{
            padding: '0.5rem 0.9rem',
            fontSize: '0.95rem',
            borderRadius: 6,
            background: '#1a73e8',
            color: '#fff',
            border: 'none',
            cursor: q.trim() ? 'pointer' : 'not-allowed',
          }}
          aria-label="Google マップで検索"
        >
          Googleマップ検索
        </button>
      </form>
    </div>
  );
}
