package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type stat struct {
	name   string
	wins   int
	loses  int
	draws  int
	points int
	played int
}

// Tally reads tournament scores from r,
// and writes the result as a tab formatted table.
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	teamToStat := map[string]stat{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		s := strings.Split(line, ";")
		if len(s) != 3 {
			return fmt.Errorf("Expecting line[ %q ] to have 3 items but got %d", line, len(s))
		}

		a, b, outcome := s[0], s[1], s[2]
		team1, team2 := teamToStat[a], teamToStat[b]

		team1.name = a
		team2.name = b

		switch outcome {
		case "win":
			team1.wins++
			team1.points += 3
			team2.loses++
		case "loss":
			team2.wins++
			team2.points += 3
			team1.loses++
		case "draw":
			team1.draws++
			team1.points++
			team2.draws++
			team2.points++
		default:
			return fmt.Errorf("Could not match %q to a valid outcome", outcome)
		}
		team1.played++
		team2.played++

		teamToStat[a], teamToStat[b] = team1, team2
	}

	fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")

	for _, result := range toOrderedSlice(teamToStat) {
		fmt.Fprintf(w, "%-31s", result.name)
		fmt.Fprintf(w, "|  %d |  %d |  %d |  %d |  %d\n", result.played, result.wins, result.draws, result.loses, result.points)
	}

	return nil
}

func toOrderedSlice(m map[string]stat) []stat {
	result := make([]stat, 0, len(m))

	for _, stat := range m {
		result = append(result, stat)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[j].points == result[i].points {
			return result[j].name > result[i].name
		}
		return result[i].points > result[j].points
	})

	return result
}
