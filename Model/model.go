package Model

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	ID string `json:"id"`
	Account Account `json:"account"`
	Assets Assets `json:"assets"`
}

type Assets struct {
	Marios []Mario `json:"marios"`
}

type Account struct {
	UUID string `json:"uuid"`
	MemberShip bool `json:"memberShip"`
	Expired string `json:"expired"`
}

type Mario struct {
	ID string `json:"id"`                  // Mario ID
	Length uint64 `json:"length"`          // 长度
	Weight uint64 `json:"weight"`          // 体重
	Growing float64 `json:"growing"`       // 成长系数
	Nature string `json:"nature"`          // 性格
	UpdateTime string `json:"updateTime"`  // 更新时间
}

func NewUser(UUID string) *User {
	u := User{ID:UUID}
	u.Born()
	return &u
}

func (u *User) Born() {
	var marios []Mario
	for i := 0; i < 5; i++ {
		m := Mario{}
		m.init(u.ID, i)
		marios = append(marios, m)
	}
	u.Assets.Marios = marios
}

func (u *User) UpdateIfNeeded() {
	var marios []Mario
	for _, m := range u.Assets.Marios {
		m.UpdateIfNeeded()
		marios = append(marios, m)
	}
	u.Assets.Marios = marios
}

func (m Mario) CurrentTime() time.Time {
	return time.Now().UTC()
}

func (m Mario) DaysBetweenLastUpdateTime(c time.Time) uint64 {
	l, _ := time.Parse("2006-01-02 15:04:05", m.UpdateTime)
	return uint64(c.Sub(l) / time.Millisecond) //(24 * time.Hour)
}

func (m *Mario) init(UUID string, index int) {
	t := m.CurrentTime().Format("2006-01-02 15:04:05")
	// ID
	id := UUID
	id += " "
	id += t
	id += " "
	id += strconv.Itoa(index)
	m.ID = id
	// 成长系数
	m.Growing = 1 + math.Round(rand.Float64() * 100) / 100
	// 更新时间
	m.UpdateTime = t
	// 长度
	m.Length = 1
	// 体重
	m.Weight = 1
	// 性格
	m.Nature = "normal"
}

func (m *Mario) UpdateIfNeeded()  {
	c := m.CurrentTime()
	day := m.DaysBetweenLastUpdateTime(c)
	if day > 0 {
		m.Length += uint64(float64(day) * m.Growing)
		m.Weight += uint64(float64(day) * m.Growing)
		m.SetUpdateTime(c.Format("2006-01-02 15:04:05"))
	}
}

func (m *Mario) SetNature(nature string) {
	m.Nature = nature
}

func (m *Mario) SetUpdateTime(time string) {
	m.UpdateTime = time
}


