export interface CreatePersonPayload {
  latitude: number;
  longitude: number;
  emoji: string;
  sign: string;
  gender: string;
  clothing: string;
  accessories: string;
  vehicle: string;
  behavior: string;
  hairstyle: string;
  sightingTime: string; // ISO8601
  registerUUID?: string;
}
