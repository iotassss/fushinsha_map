import { Marker, Popup } from "react-leaflet";
import { PersonSummary } from "./types/Persons";

export function EmojiMarker({
  person,
  handleButtonClick,
  L,
}: {
  person: PersonSummary,
  handleButtonClick: (person: PersonSummary) => void,
  L: typeof import('leaflet'),
}) {
  return (
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
          <div>🕒️目撃時刻: {person.sighting_time}</div>
        </div>
        <div style={{ marginTop: '8px' }}>
          <button style={{ cursor: 'pointer' , fontWeight: 'bold' }} onClick={() => handleButtonClick(person)}>👉️詳細を見る</button>
        </div>
      </Popup>
    </Marker>
  );
}
