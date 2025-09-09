package domain

type SuspiciousPerson struct {
	uuid          UUID          // UUID
	emoji         Emoji         // 絵文字
	sign          Sign          // サイン
	registrarUUID UUID          // 登録者UUID
	sightingCount SightingCount // 目撃数
	sightingTime  SightingTime  // 目撃時刻
	coordinates   Coordinates   // 座標
	gender        *Gender       // 性別
	clothing      *Clothing     // 服装
	accessories   *Accessories  // アクセサリー
	vehicle       *Vehicle      // 乗り物
	behavior      *Behavior     // 挙動
	hairstyle     *Hairstyle    // 髪型
}

func NewSuspiciousPerson(
	uuid UUID,
	emoji Emoji,
	sign Sign,
	registrarUUID UUID,
	sightingCount SightingCount,
	sightingTime SightingTime,
	coordinates Coordinates,
	gender *Gender,
	clothing *Clothing,
	accessories *Accessories,
	vehicle *Vehicle,
	behavior *Behavior,
	hairstyle *Hairstyle,
) SuspiciousPerson {
	return SuspiciousPerson{
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

func (sp SuspiciousPerson) UUID() UUID                   { return sp.uuid }
func (sp SuspiciousPerson) Emoji() Emoji                 { return sp.emoji }
func (sp SuspiciousPerson) Sign() Sign                   { return sp.sign }
func (sp SuspiciousPerson) RegistrarUUID() UUID          { return sp.registrarUUID }
func (sp SuspiciousPerson) SightingCount() SightingCount { return sp.sightingCount }
func (sp SuspiciousPerson) SightingTime() SightingTime   { return sp.sightingTime }
func (sp SuspiciousPerson) Coordinates() Coordinates     { return sp.coordinates }
func (sp SuspiciousPerson) Gender() *Gender              { return sp.gender }
func (sp SuspiciousPerson) Clothing() *Clothing          { return sp.clothing }
func (sp SuspiciousPerson) Accessories() *Accessories    { return sp.accessories }
func (sp SuspiciousPerson) Vehicle() *Vehicle            { return sp.vehicle }
func (sp SuspiciousPerson) Behavior() *Behavior          { return sp.behavior }
func (sp SuspiciousPerson) Hairstyle() *Hairstyle        { return sp.hairstyle }

func (sp *SuspiciousPerson) SetSightingCount(count SightingCount) {
	sp.sightingCount = count
}

func (sp *SuspiciousPerson) SetSightingTime(time SightingTime) {
	sp.sightingTime = time
}
