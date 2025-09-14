'use client';
import dynamic from "next/dynamic";

const ClientMap = dynamic(() => import("./Map"), { ssr: false });

export default ClientMap;
