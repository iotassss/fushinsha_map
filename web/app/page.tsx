'use client';

import React, { useState } from 'react';
import axios from 'axios';
import ClientMap from "./ClientMap";
import LatLngSearchForm from './LatLngSearchForm';
import { Person, GetPersonResponse } from './types/Person';
import { GetPersonsResponse, PersonSummary } from './types/Persons';

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

export default function Home() {
  const [center, setCenter] = useState<[number, number]>(DEFAULT_CENTER);

  return (
    <>
      <h1>不審者マップ</h1>
      <LatLngSearchForm center={center} setCenter={setCenter} />
      <ClientMap center={center} getPersons={getPersons} getPerson={getPerson} />
    </>
  );
}
