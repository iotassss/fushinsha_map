export interface Person {
	uuid: string; // UUID文字列
	latitude: number; // 座標値
	longitude: number;
	emoji: string;
	sign: string;
	sighting_count: number;
	sighting_time: string; // ISO8601文字列
	categories: string[];
	gender: string;
	clothing: string;
	accessories: string;
	vehicle: string;
	behavior: string;
	hairstyle: string;
	created_at: string; // ISO8601文字列
}

export interface GetPersonResponse {
  person: Person;
}
