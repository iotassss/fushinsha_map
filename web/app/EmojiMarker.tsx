import { Marker, Popup } from "react-leaflet";
import { PersonSummary } from "./types/Persons";
import { GetPersonResponse, Person } from "./types/Person";

export function EmojiMarker({
  person,
  L,
  getPerson,
  setSelectedPerson,
  setIsPanelOpen,
}: {
  person: PersonSummary,
  L: typeof import('leaflet'),
  getPerson: (uuid: string) => Promise<GetPersonResponse>;
  setSelectedPerson: (person: Person | null) => void;
  setIsPanelOpen: (isOpen: boolean) => void;
}) {
  const handleButtonClick = async (personSummary: PersonSummary) => {
    const person = await getPerson(personSummary.uuid);
    console.log('Person:', person);
    setSelectedPerson(person.person);
    setIsPanelOpen(true);
  };

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
