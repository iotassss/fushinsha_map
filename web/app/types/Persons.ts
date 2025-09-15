export interface PersonSummary {
  uuid: string;
  latitude: number;
  longitude: number;
  emoji: string;
  sign: string;
  sighting_count: number;
}

export interface GetPersonsResponse {
  persons: PersonSummary[];
}
