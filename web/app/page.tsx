'use client';

import React, { useState } from 'react';
import ClientMap from "./ClientMap";
import LatLngSearchForm from './LatLngSearchForm';

const DEFAULT_CENTER: [number, number] = [35.681236, 139.767125];

export default function Home() {
  const [center, setCenter] = useState<[number, number]>(DEFAULT_CENTER);

  return (
    <>
      <h1>不審者マップ</h1>
      <LatLngSearchForm center={center} setCenter={setCenter} />
      <ClientMap center={center} />
    </>
  );
}
