export interface CreatePersonPayload {
  latitude: number;
  longitude: number;
  emoji: string;
  sign: string;
  gender: string;
  ageGroup: string;
  clothing: string;
  accessories: string;
  vehicle: string;
  behavior: string;
  hairstyle: string;
  sightingTime: string; // ISO8601
  registerUUID?: string;
}
