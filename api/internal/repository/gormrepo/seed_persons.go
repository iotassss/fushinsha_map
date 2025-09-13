package gormrepo

import "context"

func (r *PersonRepository) ResetTable(ctx context.Context) error {
	if err := r.db.WithContext(ctx).Exec("TRUNCATE TABLE persons").Error; err != nil {
		return err
	}
	return nil
}

func (r *PersonRepository) SeedDummyPersons(ctx context.Context) error {
	dummyPersons := []PersonModel{
		{
			UUID: "b3b1c7e2-8c2a-4e2a-9c1a-2b1c7e28c2a4", Emoji: "😀", Sign: "A", ResisterUUID: "a1a2b3c4-d5e6-7f89-0abc-def123456789", SightedCount: 5, SightingTime: "12:00",
			X: 139.6917, Y: 35.6895, Gender: "男性", Clothing: "スーツ", Accessories: "帽子", Vehicle: "自転車", Behavior: "徘徊", Hairstyle: "短髪",
		},
		{
			UUID: "c2d3e4f5-6789-4abc-8def-1234567890ab", Emoji: "😁", Sign: "あ", ResisterUUID: "a1a2b3c4-d5e6-7f89-0abc-def123456789", SightedCount: 3, SightingTime: "12:30",
			X: 135.5023, Y: 34.6937, Gender: "女性", Clothing: "私服", Accessories: "眼鏡", Vehicle: "徒歩", Behavior: "つきまとい", Hairstyle: "長髪",
		},
		{
			UUID: "d4e5f6a7-8901-4bcd-9efa-2345678901bc", Emoji: "😂", Sign: "東", ResisterUUID: "a1a2b3c4-d5e6-7f89-0abc-def123456789", SightedCount: 8, SightingTime: "13:00",
			X: 135.7681, Y: 35.0116, Gender: "不明", Clothing: "作業着", Accessories: "マスク", Vehicle: "バイク", Behavior: "大声", Hairstyle: "坊主",
		},
		{
			UUID: "e5f6a7b8-9012-4cde-af12-3456789012cd", Emoji: "😅", Sign: "Z", ResisterUUID: "a1a2b3c4-d5e6-7f89-0abc-def123456789", SightedCount: 2, SightingTime: "13:30",
			X: 136.9066, Y: 35.1815, Gender: "女性", Clothing: "制服", Accessories: "バッグ", Vehicle: "自動車", Behavior: "暴力", Hairstyle: "パーマ",
		},
	}
	for _, person := range dummyPersons {
		if err := r.db.WithContext(ctx).Create(&person).Error; err != nil {
			return err
		}
	}
	return nil
}
