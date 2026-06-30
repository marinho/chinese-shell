package score

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	PointsCorrect      = 1
	PointsComboBonus   = 1 // extra per combo step
	PointsFirstOfDay   = 2 // first use of a command today
	PointsEnglish      = -1
	PointsUnknown      = -1
)

type Store struct {
	AllTime      int            `json:"all_time"`
	ByDay        map[string]int `json:"by_day"`
	// commands used today, to track first-of-day bonus
	UsedToday    map[string]bool `json:"used_today"`
	UsedTodayKey string          `json:"used_today_key"` // date the used_today map is for
	path         string
}

func Load() (*Store, error) {
	path := storePath()
	data, err := os.ReadFile(path)
	if err != nil && os.IsNotExist(err) {
		return &Store{
			ByDay:     make(map[string]int),
			UsedToday: make(map[string]bool),
			path:      path,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	var s Store
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	s.path = path
	if s.ByDay == nil {
		s.ByDay = make(map[string]int)
	}
	// reset used_today if it's for a different day
	today := todayKey()
	if s.UsedTodayKey != today {
		s.UsedToday = make(map[string]bool)
		s.UsedTodayKey = today
	}
	if s.UsedToday == nil {
		s.UsedToday = make(map[string]bool)
	}
	return &s, nil
}

func (s *Store) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0600)
}

// Add adds delta to all-time and today's score, clamping all-time and today to >= 0.
func (s *Store) Add(delta int) {
	today := todayKey()
	s.ByDay[today] += delta
	if s.ByDay[today] < 0 {
		s.ByDay[today] = 0
	}
	s.AllTime += delta
	if s.AllTime < 0 {
		s.AllTime = 0
	}
}

// IsFirstToday returns true if cmd has not been used today, and marks it as used.
func (s *Store) IsFirstToday(cmd string) bool {
	today := todayKey()
	if s.UsedTodayKey != today {
		s.UsedToday = make(map[string]bool)
		s.UsedTodayKey = today
	}
	if s.UsedToday[cmd] {
		return false
	}
	s.UsedToday[cmd] = true
	return true
}

// Today returns today's score.
func (s *Store) Today() int {
	return s.ByDay[todayKey()]
}

// Week returns the sum of the last 7 days.
func (s *Store) Week() int {
	total := 0
	now := time.Now()
	for i := 0; i < 7; i++ {
		day := now.AddDate(0, 0, -i).Format("2006-01-02")
		total += s.ByDay[day]
	}
	return total
}

func todayKey() string {
	return time.Now().Format("2006-01-02")
}

func storePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".zhell_score.json")
}
