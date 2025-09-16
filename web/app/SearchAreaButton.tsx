
import { useMap } from 'react-leaflet';
import axios from 'axios';
import type { GetPersonsResponse, PersonSummary } from './types/Persons';

interface SearchAreaButtonProps {
  setPersons: (persons: PersonSummary[]) => void;
  map: L.Map;
}

export default function SearchAreaButton({ setPersons, map }: SearchAreaButtonProps) {
  const handleClick = async () => {
    const bounds = map.getBounds();
    const params = {
      lx: bounds.getSouthWest().lng,
      rx: bounds.getNorthEast().lng,
      ty: bounds.getNorthEast().lat,
      by: bounds.getSouthWest().lat,
    };
    try {
      const res = await axios.get<GetPersonsResponse>('http://localhost:8080/api/persons', { params });
      setPersons(res.data.persons || []);
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <button
      onClick={handleClick}
      style={{
        position: 'absolute',
        zIndex: 1000,
        top: '5%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        padding: '0.5rem 1.0rem',
        fontSize: '1rem',
        borderRadius: '8px',
        background: '#fff',
        border: '1px solid #ccc',
        boxShadow: '0 2px 8px rgba(0,0,0,0.15)',
        cursor: 'pointer',
      }}
    >
      現在の範囲で検索
    </button>
  );
}
