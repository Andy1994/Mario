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
	UpdateTime string `json:"updateTime"`
}

func NewUser(UUID string) *User {
	u := User{ID:UUID}
	u.Born()
	u.GetMushroom()
	return &u
}

func CurrentTime() time.Time {
	return time.Now().UTC()
}

func DaysBetweenLastUpdateTime(c time.Time, updateTime string) uint64 {
	u, _ := time.Parse(TimeFormat, updateTime)
	return uint64(c.Sub(u) / time.Hour)
}

func TimeStringAddDays(updateTime string, days uint64) string {
	u, _ := time.Parse(TimeFormat, updateTime)
	return u.Add(time.Duration(days) * time.Hour).Format(TimeFormat)
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
	c := CurrentTime()
	var marios []Mario
	for _, m := range u.Assets.Marios {
		days := DaysBetweenLastUpdateTime(c, m.UpdateTime)
		growing := uint64(float64(days) * m.Growing * 100)
		if days > 0 && u.Assets.Mushroom.Value != 0 {
			if u.Assets.Mushroom.Value >= growing {
				u.Assets.Mushroom.Value -= growing
				m.Length += growing
				m.Weight += growing
				m.UpdateTime = TimeStringAddDays(m.UpdateTime, days)
			} else {
				m.Length += u.Assets.Mushroom.Value
				m.Weight += u.Assets.Mushroom.Value
				m.UpdateTime = TimeStringAddDays(m.UpdateTime, days)
				u.Assets.Mushroom.Value = 0
			}
		}
		marios = append(marios, m)
	}
	u.Assets.Marios = marios
}

func (u *User) GetMushroom() {
	if len(u.Assets.Mushroom.ID) == 0 {
		u.Assets.Mushroom.Init()
	} else if DaysBetweenLastUpdateTime(CurrentTime(), u.Assets.Mushroom.UpdateTime) > 0 {
		u.Assets.Mushroom.Add()
	}
}

func (m *Mushroom) Init() {
	m.ID = "2333"
	m.Type = "normal"
	m.Value = 10000
	m.UpdateTime = CurrentTime().Format(TimeFormat)
}

func (m *Mushroom) Add() {
	m.Value += 10000
	m.UpdateTime = CurrentTime().Format(TimeFormat)
}

func (m *Mario) init(UUID string, index int) {
	t := CurrentTime().Format(TimeFormat)
	// ID
	id := UUID
	id += " "
	id += t
	id += " "
	id += strconv.Itoa(index)
	m.ID = id
	// 成长系数
	rand.Seed(time.Now().UnixNano())
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


