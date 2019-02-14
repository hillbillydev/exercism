package scale

import (
	"strings"
	"unicode"
)

type Note int
type Pitch int

const (
	Flat Pitch = iota
	Sharp
	None
)

var flatNotes = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharpNotes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

func Scale(tonic, interval string) []string {
	var notes []string

	switch determinePitch(tonic) {
	case Sharp:
		notes = sharpNotes
	case None:
		notes = sharpNotes
	case Flat:
		notes = flatNotes
	}

	return getNotesWithInterval(notes, tonic, interval)
}

func normalizeTonic(tonic string) string {
	builder := strings.Builder{}
	builder.WriteRune(unicode.ToUpper(rune(tonic[0])))
	builder.WriteString(tonic[1:])
	return builder.String()
}

func getNotesWithInterval(notes []string, tonic string, interval string) []string {
	tonic = normalizeTonic(tonic)
	startIndex := getStartPostion(notes, tonic)

	if len(interval) == 0 {
		return getNotesWithoutInterval(notes, startIndex)
	}

	return getNotesFromInterval(notes, interval, startIndex)
}
func getStartPostion(notes []string, tonic string) int {
	for i, note := range notes {
		if note == tonic {
			tonic = note
			return i
		}
	}
	return 0
}
func getNotesFromInterval(notes []string, interval string, startIndex int) []string {
	var result []string
	startNote := notes[startIndex]
	result = append(result, notes[startIndex])

	for _, token := range interval {
		switch token {
		case 'A':
			startIndex += 3
		case 'M':
			startIndex += 2
		case 'm':
			startIndex++
		}
		note := notes[startIndex%len(notes)]
		if note == startNote {
			break
		}

		result = append(result, note)
	}
	return result
}

func getNotesWithoutInterval(notes []string, startIndex int) []string {
	var result []string
	startNote := notes[startIndex]
	result = append(result, notes[startIndex])

	for i := startIndex + 1; ; i++ {
		currentNote := notes[i%len(notes)]
		if currentNote == startNote {
			break
		}
		result = append(result, currentNote)
	}
	return result
}

func determinePitch(tonic string) Pitch {
	switch tonic {
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		return Sharp
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return Flat
	default:
		return None
	}
}
