package Model

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type User struct {
	ID string `json:"id"`
	Account Account `json:"account"`
	Assets Assets `json:"assets"`
}

type Assets struct {
	Marios []Mario `json:"marios"`
	Mushroom Mushroom `json:"mushroom"`
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
	Hunger uint64 `json:"hunger"`          // 饥饿值
	Growing float64 `json:"growing"`       // 成长系数
	Nature string `json:"nature"`          // 性格
	Level uint64 `json:"level"`            // 等级
	UpdateTime string `json:"updateTime"`  // 更新时间
}

type Mushroom struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Value uint64 `json:"value"`
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
	l, _ := time.Parse(TimeFormat, m.UpdateTime)
	return uint64(c.Sub(l) / time.Minute) //(24 * time.Hour)
}

func (m *Mario) init(UUID string, index int) {
	t := m.CurrentTime().Format(TimeFormat)
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
	m.Length = 100
	// 体重
	m.Weight = 100
	// 饥饿值
	m.Hunger = 0
	// 等级
	m.Level = 0
	// 性格
	m.Nature = "normal"
}

func (m *Mario) UpdateIfNeeded()  {
	c := m.CurrentTime()
	day := m.DaysBetweenLastUpdateTime(c)
	if day > 0 {
		m.Length += uint64(float64(day) * m.Growing * 100)
		m.Weight += uint64(float64(day) * m.Growing * 100)
		m.UpdateTime = c.Format(TimeFormat)
	}
}


