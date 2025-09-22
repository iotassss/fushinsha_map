package seeder

import (
	"context"
	crand "crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/iotassss/fushinsha-map-api/internal/domain"
	"github.com/iotassss/fushinsha-map-api/internal/repository/gormrepo"
)

const (
	minLat = 35.2
	maxLat = 37.5
	minLon = 138.5
	maxLon = 141.0
)

type Seeder struct {
	personRepo *gormrepo.PersonRepository
}

func NewSeeder(personRepo *gormrepo.PersonRepository) *Seeder {
	return &Seeder{
		personRepo: personRepo,
	}
}

func (s *Seeder) SeedPersons(ctx context.Context, n int) error {
	// persons := make([]domain.Person, 0, n)
	for i := 0; i < n; i++ {
		person, err := randPerson()
		if err != nil {
			return fmt.Errorf("failed to generate random person: %v", err)
		}
		// persons = append(persons, person)
		err = s.personRepo.Create(context.Background(), person)
		if err != nil {
			var mysqlErr *mysql.MySQLError
			if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
				// Duplicate entry (unique constraint violation), skip
				continue
			}
			return fmt.Errorf("failed to create person: %v", err)
		}
	}
	return nil
}

func randPerson() (*domain.Person, error) {
	uuid := domain.GenerateUUID()
	emoji, err := domain.NewEmoji(randEmoji())
	if err != nil {
		return nil, err
	}
	rune, err := randRune()
	if err != nil {
		return nil, err
	}
	sign, err := domain.NewSign(string(rune))
	if err != nil {
		return nil, err
	}
	registerUUID := domain.GenerateUUID()
	randSightingCount, err := randInt(1000)
	if err != nil {
		return nil, err
	}
	sightingCount, err := domain.NewSightingCount(int(randSightingCount))
	if err != nil {
		return nil, err
	}
	randSightingTime, err := randTime()
	if err != nil {
		return nil, err
	}
	sightingTime, err := domain.NewSightingTime(randSightingTime)
	if err != nil {
		return nil, err
	}
	coordinates, err := domain.NewCoordinates(
		randInRange(minLat, maxLat),
		randInRange(minLon, maxLon),
	)
	if err != nil {
		return nil, err
	}
	gender, err := randGender()
	if err != nil {
		return nil, err
	}
	ageGroup, err := randAgeGroup()
	if err != nil {
		return nil, err
	}
	clothing, err := randClothing()
	if err != nil {
		return nil, err
	}
	accessories, err := randAccessories()
	if err != nil {
		return nil, err
	}
	randVehicle, err := randVehicle()
	if err != nil {
		return nil, err
	}
	behavior, err := randBehavior()
	if err != nil {
		return nil, err
	}
	hairstyle, err := randHairstyle()
	if err != nil {
		return nil, err
	}
	createdAt, err := randCreatedAt()
	if err != nil {
		return nil, err
	}

	person := domain.NewPerson(
		uuid,
		emoji,
		sign,
		registerUUID,
		sightingCount,
		sightingTime,
		coordinates,
		gender,
		ageGroup,
		clothing,
		accessories,
		randVehicle,
		behavior,
		hairstyle,
		createdAt,
	)
	return person, nil
}

func randEmoji() string {
	code := mrand.Intn(0x1F64A-0x1F600+1) + 0x1F600
	return string(rune(code))
}

const set = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろわゐゑをんァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロワヰヱヲン一右雨円王音下火花貝学気九休玉金空月犬見五口校左三山子四糸字耳七車手十出女小上森人水正生青夕石赤千川先早草足村大男竹中虫町天田土二日入年白八百文木本名目立力林六引羽雲園遠何科夏家歌画回会海絵外角楽活間丸岩顔汽記帰弓牛魚京強教近兄形計元言原戸古午後語工公広交光考行高黄合谷国黒今才細作算止市矢姉思紙寺自時室社弱首秋週春書少場色食心新親図数西声星晴切雪船線前組走多太体台地池知茶昼長鳥朝直通弟店点電刀冬当東答頭同道読内南肉馬売買麦半番父風分聞米歩母方北毎妹万明鳴毛門夜野友用曜来里理話悪安暗医委意育員院飲運泳駅央横屋温化荷界開階寒感漢館岸起期客究急級宮球去橋業曲局銀区苦具君係軽血決研県庫湖向幸港号根祭皿仕死使始指歯詩次事持式実写者主守取酒受州拾終習集住重宿所暑助昭消商章勝乗植申身神真深進世整昔全相送想息速族他打対待代第題炭短談着注柱丁帳調追定庭笛鉄転都度投豆島湯登等動童農波配倍箱畑発反坂板皮悲美鼻筆氷表秒病品負部服福物平返勉放味命面問役薬由油有遊予羊洋葉陽様落流旅両緑礼列練路和愛案以衣位囲胃印英栄塩億加果貨課芽改械害街各覚完官管関観願希季紀喜旗器機議求泣救給挙漁共協鏡競極訓軍郡径型景芸欠結建健験固功好候航康告差菜最材昨札刷殺察参産散残士氏史司試児治辞失借種周祝順初松笑唱焼象照賞臣信成省清静席積折節説浅戦選然争倉巣束側続卒孫帯隊達単置仲貯兆腸低底停的典伝徒努灯堂働特得毒熱念敗梅博飯飛費必票標不夫付府副粉兵別辺変便包法望牧末満未脈民無約勇要養浴利陸良料量輪類令冷例歴連老労録圧移因永営衛易益液演応往桜恩可仮価河過賀快解格確額刊幹慣眼基寄規技義逆久旧居許境均禁句群経潔件券険検限現減故個護効厚耕鉱構興講混査再災妻採際在財罪雑酸賛支志枝師資飼示似識質舎謝授修述術準序招承証条状常情織職制性政勢精製税責績接設舌絶銭祖素総造像増則測属率損退貸態団断築張提程適敵統銅導徳独任燃能破犯判版比肥非備俵評貧布婦富武復仏編弁保墓報豊防貿暴務夢迷綿輸余預容略留領異遺域宇映延沿我灰拡革閣割株干巻看簡危机揮貴疑吸供胸郷勤筋系敬警劇激穴絹権憲源厳己呼誤后孝皇紅降鋼刻穀骨困砂座済裁策冊蚕至私姿視詞誌磁射捨尺若樹収宗就衆従縦縮熟純処署諸除将傷障城蒸針仁垂推寸盛聖誠宣専泉洗染善奏窓創装層操蔵臓存尊宅担探誕段暖値宙忠著庁頂潮賃痛展討党糖届難乳認納脳派拝背肺俳班晩否批秘腹奮並陛閉片補暮宝訪亡忘棒枚幕密盟模訳郵優幼欲翌乱卵覧裏律臨朗論"

var runeSet = []rune(set)

func randRune() (rune, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(int64(len(runeSet))))
	if err != nil {
		return 0, err
	}
	return runeSet[n.Int64()], nil
}

func randInt(max int64) (int64, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(max+1))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// randTime returns a random time string between 00:00 and 23:59.
func randTime() (time.Time, error) {
	// ランダムな分数 (0〜1439)
	n, err := crand.Int(crand.Reader, big.NewInt(24*60))
	if err != nil {
		return time.Time{}, err
	}
	minutes := n.Int64()
	hour := minutes / 60
	min := minutes % 60
	return time.Parse("15:04", fmt.Sprintf("%02d:%02d", hour, min))
}

func randInRange(min, max float64) float64 {
	v := min + mrand.Float64()*(max-min)
	return math.Round(v*10000) / 10000
}

func randGender() (domain.Gender, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(3))
	if err != nil {
		return domain.Gender(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewGender("男性")
	case 1:
		return domain.NewGender("女性")
	default:
		return domain.NewGender("不明")
	}
}

func randAgeGroup() (domain.AgeGroup, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(6))
	if err != nil {
		return domain.AgeGroup(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewAgeGroup("未成年")
	case 1:
		return domain.NewAgeGroup("20代")
	case 2:
		return domain.NewAgeGroup("30代")
	case 3:
		return domain.NewAgeGroup("40代")
	case 4:
		return domain.NewAgeGroup("50代")
	default:
		return domain.NewAgeGroup("60代以上")
	}
}

func randClothing() (domain.Clothing, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(6))
	if err != nil {
		return domain.Clothing(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewClothing("")
	case 1:
		return domain.NewClothing("スーツ")
	case 2:
		return domain.NewClothing("制服")
	case 3:
		return domain.NewClothing("私服")
	case 4:
		return domain.NewClothing("作業着")
	default:
		return domain.NewClothing("その他")
	}
}

func randAccessories() (domain.Accessories, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(6))
	if err != nil {
		return domain.Accessories(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewAccessories("")
	case 1:
		return domain.NewAccessories("帽子")
	case 2:
		return domain.NewAccessories("眼鏡")
	case 3:
		return domain.NewAccessories("マスク")
	case 4:
		return domain.NewAccessories("バッグ")
	default:
		return domain.NewAccessories("なし")
	}
}

func randVehicle() (domain.Vehicle, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(6))
	if err != nil {
		return domain.Vehicle(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewVehicle("")
	case 1:
		return domain.NewVehicle("自転車")
	case 2:
		return domain.NewVehicle("バイク")
	case 3:
		return domain.NewVehicle("自動車")
	case 4:
		return domain.NewVehicle("徒歩")
	default:
		return domain.NewVehicle("その他")
	}
}

func randBehavior() (domain.Behavior, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(6))
	if err != nil {
		return domain.Behavior(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewBehavior("")
	case 1:
		return domain.NewBehavior("徘徊")
	case 2:
		return domain.NewBehavior("大声")
	case 3:
		return domain.NewBehavior("暴力")
	case 4:
		return domain.NewBehavior("つきまとい")
	default:
		return domain.NewBehavior("その他")
	}
}

func randHairstyle() (domain.Hairstyle, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(6))
	if err != nil {
		return domain.Hairstyle(""), err
	}
	switch n.Int64() {
	case 0:
		return domain.NewHairstyle("")
	case 1:
		return domain.NewHairstyle("短髪")
	case 2:
		return domain.NewHairstyle("長髪")
	case 3:
		return domain.NewHairstyle("坊主")
	case 4:
		return domain.NewHairstyle("パーマ")
	default:
		return domain.NewHairstyle("その他")
	}
}

func randDateTime(start, end time.Time) (time.Time, error) {
	if end.Before(start) {
		return time.Time{}, fmt.Errorf("end must be after start")
	}
	// 差分秒数を計算
	diff := end.Unix() - start.Unix()
	// 乱数を生成（0〜diff）
	n, err := crand.Int(crand.Reader, big.NewInt(diff+1))
	if err != nil {
		return time.Time{}, err
	}
	// startに加算して返す
	return start.Add(time.Duration(n.Int64()) * time.Second), nil
}

func randCreatedAt() (domain.CreatedAt, error) {
	// 適当な範囲で乱数を生成（例: 2020年1月1日から現在まで）
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	t, err := randDateTime(start, end)
	if err != nil {
		return domain.CreatedAt{}, err
	}
	return domain.NewCreatedAt(t)
}
