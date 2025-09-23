'use client';

import React, { useState } from 'react';
import axios from 'axios';
import ClientMap from "./ClientMap";
import LatLngSearchForm from './LatLngSearchForm';
import { Person, GetPersonResponse } from './types/Person';
import { GetPersonsResponse, PersonSummary } from './types/Persons';
import type { CreatePersonPayload } from './types/CreatePersonPayload';
import { GoogleMapSearch } from './GoogleMapSearch';
import styles from './page.module.css';

const DEFAULT_CENTER: [number, number] = [35.681236, 139.767125];
const apiUrl = process.env.NEXT_PUBLIC_RESOURCE_SERVER_BASE_URL;

const getPersons = async (uuid: string): Promise<GetPersonsResponse> => {
  try {
    const response = await axios.get<GetPersonsResponse>(`${apiUrl}/persons`);
    if (!response.data) throw new Error('No person data found');
    return response.data;
  } catch (error) {
    throw new Error('Failed to fetch persons');
  }
};

const getPerson = async (uuid: string): Promise<GetPersonResponse> => {
  try {
    const response = await axios.get<GetPersonResponse>(`${apiUrl}/persons/${uuid}`);
    if (!response.data) throw new Error('No person data found');
    return response.data;
  } catch (error) {
    throw new Error('Failed to fetch person');
  }
};

const createPerson = async (payload: CreatePersonPayload): Promise<void> => {
  try {
    const res = await axios.post(`${apiUrl}/persons`, payload);
    if (!res || res.status < 200 || res.status >= 300) {
      throw new Error('Failed to submit');
    }
  } catch (error) {
    throw new Error('Failed to submit');
  }
};

export default function Page() {
  const [center, setCenter] = useState<[number, number]>(DEFAULT_CENTER);

  return (
    <>
      <div className={styles.headerContainer}>
        <h1>
          🤪 不審者マップ
        </h1>
        <div className={styles.wordCoordinatesSearch}>
          <GoogleMapSearch />
          <p className={styles.arrow}>▶</p>
          <div className={styles.wordCoordinatesSearch}>
            <div className={styles.wordCoordinatesSearchText}>
              <p>Googleマップで</p>
              <p>右クリックして</p>
              <p>座標をコピペ</p>
            </div>
          </div>
          <p className={styles.arrow}>▶</p>
          <LatLngSearchForm center={center} setCenter={setCenter} />
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
