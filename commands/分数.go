package commands

import (
	"fmt"
	"zhell/score"
)

func init() {
	Register(&scoreCmd{})
}

type scoreCmd struct{}

func (s *scoreCmd) Name() string        { return "分数" }
func (s *scoreCmd) Pinyin() string      { return "fēnshù" }
func (s *scoreCmd) Linux() string       { return "score" }
func (s *scoreCmd) Description() string { return "show your score and how the scoring system works" }
func (s *scoreCmd) Execute(_ []string) error {
	sc := Score()
	if sc == nil {
		fmt.Println("Score tracking is not available in this mode.")
		return nil
	}

	fmt.Println("=== 你的分数 / Your Score ===")
	fmt.Printf("  Today:    %d pts\n", sc.Today())
	fmt.Printf("  This week: %d pts\n", sc.Week())
	fmt.Printf("  All time:  %d pts\n", sc.AllTime)
	fmt.Println()
	fmt.Println("=== 计分规则 / How Scoring Works ===")
	fmt.Printf("  +%d  correct Chinese command\n", score.PointsCorrect)
	fmt.Printf("  +%d  first time using a command today (bonus)\n", score.PointsFirstOfDay)
	fmt.Printf("  +N  combo streak: each consecutive different command adds +N\n")
	fmt.Printf("  %d  typing an English equivalent (e.g. 'cat' instead of '猫')\n", score.PointsEnglish)
	fmt.Printf("  %d  unknown command\n", score.PointsUnknown)
	fmt.Println()
	fmt.Println("  Repeating the same command back-to-back scores no points and resets your combo streak.")
	fmt.Println("  Score never drops below 0.")
	return nil
}
