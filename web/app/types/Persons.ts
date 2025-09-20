export interface PersonSummary {
  uuid: string;
  latitude: number;
  longitude: number;
  emoji: string;
  sign: string;
  sighting_count: number;
  sighting_time: string; // ISO8601文字列
}

export interface GetPersonsResponse {
  persons: PersonSummary[];
}
