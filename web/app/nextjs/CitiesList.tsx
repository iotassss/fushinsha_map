"use client";
import { use, useEffect, useState } from "react";

export default function CitiesList() {
  const [cities, setCities] = useState<{ id: string; name: string }[]>([]);
  useEffect(() => {
    fetch("http://localhost:8080/api/cities")
      .then((res) => res.json())
      .then((data) => {
        if (data.cities) setCities(data.cities);
      });
  }, []);

  useEffect(() => {
    console.log("Fetched cities:", cities);
  }, [cities]);

  return (
    <div className="w-full max-w-md">
      <h2 className="font-bold mb-2">Cities from API:</h2>
      <ul className="list-disc pl-5">
        {cities.map((city, index) => (
          <li key={index}>{city.name}</li>
        ))}
      </ul>
    </div>
  );
}
