// Package twelve deals with the fact that people still like christmas music.
package twelve

import (
	"fmt"
	"strings"
)

type lyric struct {
	day  string
	text string
}

var lyrics = []lyric{
	lyric{"first", "a Partridge in a Pear Tree."},
	lyric{"second", "two Turtle Doves"},
	lyric{"third", "three French Hens"},
	lyric{"fourth", "four Calling Birds"},
	lyric{"fifth", "five Gold Rings"},
	lyric{"sixth", "six Geese-a-Laying"},
	lyric{"seventh", "seven Swans-a-Swimming"},
	lyric{"eighth", "eight Maids-a-Milking"},
	lyric{"ninth", "nine Ladies Dancing"},
	lyric{"tenth", "ten Lords-a-Leaping"},
	lyric{"eleventh", "eleven Pipers Piping"},
	lyric{"twelfth", "twelve Drummers Drumming"},
}

// Song sings a christmas song.
func Song() string {
	var result strings.Builder

    for i, item := range lyrics {
		result.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", item.day, item.text))

		for j := i - 1; j >= 0; j-- {
			prev := lyrics[j]
			result.WriteRune(',')

			if j == 0 {
				result.WriteString(" and")
			}
			result.WriteString(fmt.Sprintf(" %s", prev.text))
		}
		result.WriteRune('\n')
	}

	return strings.TrimRight(result.String(), "\n")
}

// Verse is stupid and don't like christmas songs, so he asks for whole song
// then when the song is done he just writes down the verse he liked.
func Verse(skip int) string {
	if skip > len(lyrics) {
		return ""
	}
	return strings.Split(Song(), "\n")[skip-1]
}
