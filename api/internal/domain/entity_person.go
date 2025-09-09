package domain

type Person struct {
	uuid          UUID          // UUID
	emoji         Emoji         // 絵文字
	sign          Sign          // サイン
	registrarUUID UUID          // 登録者UUID
	sightingCount SightingCount // 目撃数
	sightingTime  SightingTime  // 目撃時刻
	coordinates   Coordinates   // 座標
	gender        Gender        // 性別
	clothing      Clothing      // 服装
	accessories   Accessories   // アクセサリー
	vehicle       Vehicle       // 乗り物
	behavior      Behavior      // 挙動
	hairstyle     Hairstyle     // 髪型
}

func NewPerson(
	uuid UUID,
	emoji Emoji,
	sign Sign,
	registrarUUID UUID,
	sightingCount SightingCount,
	sightingTime SightingTime,
	coordinates Coordinates,
	gender Gender,
	clothing Clothing,
	accessories Accessories,
	vehicle Vehicle,
	behavior Behavior,
	hairstyle Hairstyle,
) Person {
	return Person{
		uuid:          uuid,
		emoji:         emoji,
		sign:          sign,
		registrarUUID: registrarUUID,
		sightingCount: sightingCount,
		sightingTime:  sightingTime,
		coordinates:   coordinates,
		gender:        gender,
		clothing:      clothing,
		accessories:   accessories,
		vehicle:       vehicle,
		behavior:      behavior,
		hairstyle:     hairstyle,
	}
}

func (sp Person) UUID() UUID                   { return sp.uuid }
func (sp Person) Emoji() Emoji                 { return sp.emoji }
func (sp Person) Sign() Sign                   { return sp.sign }
func (sp Person) RegistrarUUID() UUID          { return sp.registrarUUID }
func (sp Person) SightingCount() SightingCount { return sp.sightingCount }
func (sp Person) SightingTime() SightingTime   { return sp.sightingTime }
func (sp Person) Coordinates() Coordinates     { return sp.coordinates }
func (sp Person) Gender() Gender               { return sp.gender }
func (sp Person) Clothing() Clothing           { return sp.clothing }
func (sp Person) Accessories() Accessories     { return sp.accessories }
func (sp Person) Vehicle() Vehicle             { return sp.vehicle }
func (sp Person) Behavior() Behavior           { return sp.behavior }
func (sp Person) Hairstyle() Hairstyle         { return sp.hairstyle }

func (sp *Person) SetSightingCount(count SightingCount) {
	sp.sightingCount = count
}

func (sp *Person) SetSightingTime(time SightingTime) {
	sp.sightingTime = time
}
