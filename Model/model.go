package Model

import (
	"time"
)

type User struct {
	ID string `json:"id"`
	Account Account `json:"account"`
	Assets Assets `json:"assets"`
}

type Assets struct {
	Marios []Mario
}

type Account struct {
	UUID string
}

type Mario struct {
	ID string           // Mario ID
	Length uint64       // 长度
	Weight uint64       // 体重
	Growing float64     // 成长系数
	Nature string       // 性格
	UpdateTime string   // 更新时间
}

func (m Mario) CurrentTime() time.Time {
	return time.Now().UTC()
}

func (m Mario) DaysBetweenLastUpdateTime(c time.Time) uint64 {
	l, _ := time.Parse("2006-01-02 15:04:05", m.UpdateTime)
	return uint64(c.Sub(l) / (24 * time.Hour))
}

func (m *Mario) UpdateSizeIfNeeded()  {
	c := m.CurrentTime()
	day := m.DaysBetweenLastUpdateTime(c)
	if day > 0 {
		m.Length += day
		m.Weight += day
		m.SetUpdateTime(c.Format("2006-01-02 15:04:05"))
	}
}

func (m *Mario) SetNature(nature string) {
	m.Nature = nature
}

func (m *Mario) SetUpdateTime(time string) {
	m.UpdateTime = time
}


