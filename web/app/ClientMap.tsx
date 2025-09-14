'use client';

import dynamic from "next/dynamic";
import { MapProps } from "./Map";

const Map = dynamic(() => import("./Map"), { ssr: false });

export default function ClientMap(props: MapProps) {
  return <Map {...props} />;
}
