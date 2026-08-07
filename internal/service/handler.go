package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type User struct {
	Id                  string               `json:"id"`
	Username            string               `json:"username"`
	Name                string               `json:"name"`
	Honor               int                  `json:"honor"`
	Clan                string               `json:"clan"`
	LeaderboardPosition int                  `json:"leaderboardPosition"`
	Skills              []string             `json:"skills"`
	Ranks               RanksStruct          `json:"ranks"`
	CodeChallenges      CodeChallengesStruct `json:"codeChallenges"`
}

type RanksStruct struct {
	OverallStruct  RankInfo            `json:"overall"`
	LanguageStruct map[string]RankInfo `json:"languages"`
}

type RankInfo struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Score int    `json:"score"`
}

type CodeChallengesStruct struct {
	TotalAuthored  int `json:"totalAuthored"`
	TotalCompleted int `json:"totalCompleted"`
}

type LanguageEntry struct {
	Name string
	Info RankInfo
}

// ---

func Svg(w http.ResponseWriter, r *http.Request) {
	size, err_s := strconv.Atoi(r.FormValue("size"))
	if err_s != nil || size < 1 {
		http.Error(w, "400 size error", http.StatusBadRequest)
		return
	}

	resp, err_r := http.Get("https://www.codewars.com/api/v1/users/" + url.QueryEscape(r.FormValue("username")))

	if err_r != nil || resp.StatusCode != 200 {
		http.Error(w, "404 Codewars user no found", http.StatusNotFound)
		return
	}

	defer resp.Body.Close()

	var user User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, "500 Error processing response from codewars", http.StatusInternalServerError)
		return
	}

	const baseWidth = 540.0
	const baseHeight = 310.0

	width := int(baseWidth * size)
	height := int(baseHeight * size)

	s := func(val float64) int {
		return int(val * float64(size))
	}

	var sortedLangs []LanguageEntry
	for name, info := range user.Ranks.LanguageStruct {
		sortedLangs = append(sortedLangs, LanguageEntry{Name: name, Info: info})
	}

	sort.Slice(sortedLangs, func(i, j int) bool {
		return sortedLangs[i].Info.Score > sortedLangs[j].Info.Score
	})

	if len(sortedLangs) > 3 {
		sortedLangs = sortedLangs[:3]
	}

	canvas := svg.New(w)
	canvas.Start(width, height)

	overallColor := mapCodewarsColor(user.Ranks.OverallStruct.Color)

	style := fmt.Sprintf(`
		@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600;700&amp;display=swap');
		
		text {
			font-family: 'JetBrains Mono', monospace;
			dominant-baseline: middle;
		}
		
		.bg { fill: #0F172A; stroke: #1E293B; stroke-width: %dpx; }
		.rank-badge { font-size: %dpx; font-weight: 700; fill: %s; }
		.username { font-size: %dpx; font-weight: 700; fill: #F8FAFC; }
		.clan-text { font-size: %dpx; fill: #64748B; }
		
		.stat-label { font-size: %dpx; fill: #64748B; letter-spacing: 0.5px; }
		.stat-value { font-size: %dpx; font-weight: 600; fill: #F1F5F9; }
		
		.section-header { font-size: %dpx; font-weight: 600; fill: #38BDF8; letter-spacing: 1px; }
		.lang-name { font-size: %dpx; font-weight: 600; fill: #E2E8F0; }
		.lang-info { font-size: %dpx; fill: #94A3B8; }
		
		.divider { stroke: #334155; stroke-width: %dpx; }
	`,
		s(2),
		s(16), overallColor,
		s(22),
		s(12),
		s(11),
		s(15),
		s(12),
		s(13),
		s(12),
		s(1),
	)

	canvas.Style("text/css", style)

	canvas.Roundrect(0, 0, width, height, s(16), s(16), "class=\"bg\"")

	paddingX := 28.0

	rankName := strings.ToUpper(user.Ranks.OverallStruct.Name)
	if rankName == "" {
		rankName = "N/A"
	}

	canvas.Text(s(paddingX), s(42), rankName, "class=\"rank-badge\"")

	canvas.Text(s(paddingX+85), s(42), user.Username, "class=\"username\"")

	clanStr := "No Clan"
	if strings.TrimSpace(user.Clan) != "" {
		clanStr = fmt.Sprintf("Clan: %s", user.Clan)
	}
	canvas.Text(s(paddingX), s(68), clanStr, "class=\"clan-text\"")

	canvas.Line(s(paddingX), s(90), s(baseWidth-paddingX), s(90), "class=\"divider\"")

	col1 := paddingX
	col2 := paddingX + 160.0
	col3 := paddingX + 320.0
	lblY := 115.0
	valY := 135.0

	canvas.Text(s(col1), s(lblY), "LEADERBOARD", "class=\"stat-label\"")
	leaderboardStr := "N/A"
	if user.LeaderboardPosition > 0 {
		leaderboardStr = fmt.Sprintf("#%d", user.LeaderboardPosition)
	}
	canvas.Text(s(col1), s(valY), leaderboardStr, "class=\"stat-value\"")

	canvas.Text(s(col2), s(lblY), "HONOR", "class=\"stat-label\"")
	canvas.Text(s(col2), s(valY), fmt.Sprintf("%d", user.Honor), "class=\"stat-value\"")

	canvas.Text(s(col3), s(lblY), "COMPLETED", "class=\"stat-label\"")
	canvas.Text(s(col3), s(valY), fmt.Sprintf("%d", user.CodeChallenges.TotalCompleted), "class=\"stat-value\"")

	canvas.Line(s(paddingX), s(165), s(baseWidth-paddingX), s(165), "class=\"divider\"")

	canvas.Text(s(paddingX), s(190), "TOP LANGUAGES", "class=\"section-header\"")

	langStartY := 220.0
	langStepY := 26.0

	for i, lang := range sortedLangs {
		currentY := langStartY + float64(i)*langStepY

		canvas.Text(s(paddingX), s(currentY), lang.Name, "class=\"lang-name\"")

		detail := fmt.Sprintf("%s / %d pts", lang.Info.Name, lang.Info.Score)
		canvas.Text(s(paddingX+180.0), s(currentY), detail, "class=\"lang-info\"")
	}

	canvas.End()
}

func mapCodewarsColor(color string) string {
	switch strings.ToLower(color) {
	case "white":
		return "#E6E6E6"
	case "yellow":
		return "#ECB613"
	case "blue":
		return "#3C7EBB"
	case "purple":
		return "#866CC7"
	case "black":
		return "#515151"
	case "red":
		return "#BB4343"
	default:
		return "#38BDF8"
	}
}
