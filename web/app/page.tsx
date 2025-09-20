'use client';

import React, { useState } from 'react';
import axios from 'axios';
import ClientMap from "./ClientMap";
import LatLngSearchForm from './LatLngSearchForm';
import { Person, GetPersonResponse } from './types/Person';
import { GetPersonsResponse, PersonSummary } from './types/Persons';
import type { CreatePersonPayload } from './types/CreatePersonPayload';
import { GoogleMapSearch } from './GoogleMapSearch';

const DEFAULT_CENTER: [number, number] = [35.681236, 139.767125];

const getPersons = async (uuid: string): Promise<GetPersonsResponse> => {
  try {
    const response = await axios.get<GetPersonsResponse>(`http://localhost:8080/api/persons/${uuid}`);
    if (!response.data) throw new Error('No person data found');
    return response.data;
  } catch (error) {
    throw new Error('Failed to fetch persons');
  }
};

const getPerson = async (uuid: string): Promise<GetPersonResponse> => {
  try {
    const response = await axios.get<GetPersonResponse>(`http://localhost:8080/api/persons/${uuid}`);
    if (!response.data) throw new Error('No person data found');
    return response.data;
  } catch (error) {
    throw new Error('Failed to fetch person');
  }
};

const createPerson = async (payload: CreatePersonPayload): Promise<void> => {
  try {
    const res = await axios.post('http://localhost:8080/api/persons', payload);
    if (!res || res.status < 200 || res.status >= 300) {
      throw new Error('Failed to submit');
    }
  } catch (error) {
    throw new Error('Failed to submit');
  }
};

export default function Home() {
  const [center, setCenter] = useState<[number, number]>(DEFAULT_CENTER);

  return (
    <>
      <div style={{ display: 'flex', alignItems: 'center', gap: '8px', margin: '0.5rem' }}>
        <h1 style={{ fontSize: '1.2rem', fontWeight: 'bold', margin: '1rem' }}>
          🤪 不審者マップ
        </h1>
        <div>
          <div style={{display: 'flex', alignItems: 'center', gap: '8px'}}>
            <GoogleMapSearch />
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', fontSize: '0.8rem' }}>
              <p>▶</p>
              <div>
                <p>Googleマップで</p>
                <p>右クリックして</p>
                <p>座標をコピペ</p>
              </div>
            </div>
            <p>▶</p>
            <LatLngSearchForm center={center} setCenter={setCenter} />
          </div>
        </div>
      </div>
      <ClientMap
        center={center}
        getPersons={getPersons}
        getPerson={getPerson}
        createPerson={createPerson}
      />
    </>
  );
}
