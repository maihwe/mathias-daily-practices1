package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ─────────────────────────────────────────────
// Q1 – Hex to Decimal
// "1E" -> 30,  "FF" -> 255
// ─────────────────────────────────────────────
func hexToDecimal(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

// ─────────────────────────────────────────────
// Q2 – Binary to Decimal
// "10" -> 2,  "1010" -> 10
// ─────────────────────────────────────────────
func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

// ─────────────────────────────────────────────
// Q3 – Capitalize Word (first letter up, rest down)
// "hELLO" -> "Hello"
// ─────────────────────────────────────────────
func cap(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// ─────────────────────────────────────────────
// Q4 – Uppercase last N words
// ["How","are","you"], n=2 -> "How ARE YOU"
// ─────────────────────────────────────────────
func upperN(words []string, n int) string {
	result := make([]string, len(words))
	copy(result, words)
	start := len(result) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(result); i++ {
		result[i] = strings.ToUpper(result[i])
	}
	return strings.Join(result, " ")
}

// ─────────────────────────────────────────────
// Q5 – Detect Punctuation
// "," -> true,  "x" -> false
// ─────────────────────────────────────────────
func punc(s string) bool {
	for _, c := range s {
		if !unicode.IsPunct(c) && c != '!' && c != '?' {
			return false
		}
	}
	return len(s) > 0
}

// ─────────────────────────────────────────────
// Q6 – Fix Article Spacing
// ["hello", ",", "world", "!"] -> "hello, world!"
// ─────────────────────────────────────────────
func fixArticles(words []string) string {
	return fixPunctuation(words)
}

// ─────────────────────────────────────────────
// Q7 – Determine article: "a" or "an"
// "apple" -> "an",  "book" -> "a",  "honest" -> "an"
// ─────────────────────────────────────────────
func article(word string) string {
	if word == "" {
		return "a"
	}
	first := unicode.ToLower(rune(word[0]))
	vowels := "aeiouAEIOU"
	// special words that start with a consonant but sound like a vowel
	special := []string{"honest", "hour", "honour", "heir"}
	lower := strings.ToLower(word)
	for _, w := range special {
		if lower == w || strings.HasPrefix(lower, w) {
			return "an"
		}
	}
	if strings.ContainsRune(vowels, first) {
		return "an"
	}
	return "a"
}

// ─────────────────────────────────────────────
// Q8 – Fix Articles in Sentence
// "A amazing rock." -> "An amazing rock."
// ─────────────────────────────────────────────
func fixSentenceArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		lower := strings.ToLower(words[i])
		if lower == "a" || lower == "an" {
			correct := article(words[i+1])
			// preserve original capitalisation
			if words[i][0] == 'A' {
				correct = strings.ToUpper(correct[:1]) + correct[1:]
			}
			words[i] = correct
		}
	}
	return strings.Join(words, " ")
}

// ─────────────────────────────────────────────
// Q9 – Fix Single Quotes
// "' awesome '" -> "'awesome'"
// ─────────────────────────────────────────────
func fixSingleQuotes(text string) string {
	runes := []rune(text)
	var out []rune
	i := 0
	for i < len(runes) {
		if runes[i] == '\'' {
			j := i + 1
			for j < len(runes) && runes[j] != '\'' {
				j++
			}
			if j < len(runes) {
				inner := strings.TrimSpace(string(runes[i+1 : j]))
				out = append(out, '\'')
				out = append(out, []rune(inner)...)
				out = append(out, '\'')
				i = j + 1
			} else {
				out = append(out, runes[i])
				i++
			}
		} else {
			out = append(out, runes[i])
			i++
		}
	}
	return string(out)
}

// ─────────────────────────────────────────────
// Q10 – Uppercase Conversion
// "hello" -> "HELLO"
// ─────────────────────────────────────────────
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// ─────────────────────────────────────────────
// Q11 – Lowercase Conversion
// "HELLO" -> "hello"
// ─────────────────────────────────────────────
func toLower(s string) string {
	return strings.ToLower(s)
}

// ─────────────────────────────────────────────
// Q12 – Detect Command Token
// "(up)" -> true,  "hello" -> false
// ─────────────────────────────────────────────
func isCommand(s string) bool {
	commands := []string{"(up)", "(low)", "(cap)", "(hex)", "(bin)"}
	lower := strings.ToLower(s)
	for _, cmd := range commands {
		if lower == cmd {
			return true
		}
	}
	// also match "(up, N)" style — starts with "(" and ends with ")"
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		return true
	}
	return false
}

// ─────────────────────────────────────────────
// Q13 – Remove Command from Slice
// ["go","(up)"] -> ["GO"]
// ─────────────────────────────────────────────
func removeCommand(words []string, index int) []string {
	return append(words[:index], words[index+1:]...)
}

// ─────────────────────────────────────────────
// Q14 – Fix Punctuation Spacing
// ["hello", ",", "world"] -> "hello, world"
// ─────────────────────────────────────────────
func fixPunctuation(words []string) string {
	var sb strings.Builder
	for i, w := range words {
		if i == 0 {
			sb.WriteString(w)
			continue
		}
		if punc(w) {
			sb.WriteString(w)
		} else {
			sb.WriteString(" " + w)
		}
	}
	return sb.String()
}

// ─────────────────────────────────────────────
// Q15 – Detect Group Punctuation
// "..." -> true,  "!?" -> true,  "." -> false
// ─────────────────────────────────────────────
func isGroupPunctuation(s string) bool {
	if len(s) < 2 {
		return false
	}
	for _, c := range s {
		if !unicode.IsPunct(c) {
			return false
		}
	}
	return true
}

// ─────────────────────────────────────────────
// Q16 – Handle (up, n) — convert previous n words to uppercase
// ["this","is","cool","(up,","2)"] -> ["this","IS","COOL"]
// ─────────────────────────────────────────────
func upperNSlice(words []string, n int) []string {
	result := make([]string, len(words))
	copy(result, words)
	start := len(result) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(result); i++ {
		result[i] = strings.ToUpper(result[i])
	}
	return result
}

// ─────────────────────────────────────────────
// Q17 – Handle (low, n)
// ["I","LOVE","GO","(low,","2)"] -> ["I","love","go"]
// ─────────────────────────────────────────────
func lowerN(words []string, n int) []string {
	result := make([]string, len(words))
	copy(result, words)
	start := len(result) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(result); i++ {
		result[i] = strings.ToLower(result[i])
	}
	return result
}

// ─────────────────────────────────────────────
// Q18 – Handle (cap, n)
// ["welcome","to","the","brooklyn","bridge","(cap,","2)"]
// -> ["welcome","to","the","brooklyn","Bridge"]
// ─────────────────────────────────────────────
func capN(words []string, n int) []string {
	result := make([]string, len(words))
	copy(result, words)
	start := len(result) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(result); i++ {
		result[i] = cap(result[i])
	}
	return result
}

// ─────────────────────────────────────────────
// Q19 – Process (hex)
// ["1E","(hex)","files"] -> ["30","files"]
// ─────────────────────────────────────────────
func processHex(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == "(hex)" && i > 0 {
			val, err := hexToDecimal(result[len(result)-1])
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}
		} else {
			result = append(result, words[i])
		}
	}
	return result
}

// ─────────────────────────────────────────────
// Q20 – Process (bin)
// ["10","(bin)","years"] -> ["2","years"]
// ─────────────────────────────────────────────
func processBin(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == "(bin)" && i > 0 {
			val, err := binToDecimal(result[len(result)-1])
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}
		} else {
			result = append(result, words[i])
		}
	}
	return result
}

// ─────────────────────────────────────────────
// Q21 – Normalize Spaces
// "hello   world" -> "hello world"
// ─────────────────────────────────────────────
func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// ─────────────────────────────────────────────
// Q22 – Detect Vowel
// 'a' -> true,  'b' -> false
// ─────────────────────────────────────────────
func isVowel(c byte) bool {
	return strings.ContainsRune("aeiouAEIOU", rune(c))
}

// ─────────────────────────────────────────────
// Q23 – Process Full Sentence
// Apply all transformations to a sentence
// ─────────────────────────────────────────────
func processSentence(text string) string {
	words := strings.Fields(text)
	result := []string{}

	i := 0
	for i < len(words) {
		w := words[i]
		lower := strings.ToLower(w)

		switch {
		case lower == "(hex)" && len(result) > 0:
			val, err := hexToDecimal(result[len(result)-1])
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}

		case lower == "(bin)" && len(result) > 0:
			val, err := binToDecimal(result[len(result)-1])
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}

		case lower == "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}

		case lower == "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}

		case lower == "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = cap(result[len(result)-1])
			}

		// Handle (up, N) split across two tokens: "(up," and "2)"
		case strings.HasPrefix(lower, "(up,") && i+1 < len(words):
			nStr := strings.TrimRight(words[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err == nil {
				result = upperNSlice(result, n)
			}
			i++ // skip the "N)" token

		case strings.HasPrefix(lower, "(low,") && i+1 < len(words):
			nStr := strings.TrimRight(words[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err == nil {
				result = lowerN(result, n)
			}
			i++

		case strings.HasPrefix(lower, "(cap,") && i+1 < len(words):
			nStr := strings.TrimRight(words[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err == nil {
				result = capN(result, n)
			}
			i++

		default:
			result = append(result, w)
		}
		i++
	}

	// Fix articles (a/an) and punctuation spacing
	joined := strings.Join(result, " ")
	joined = fixSentenceArticles(joined)
	words2 := strings.Fields(joined)
	joined = fixPunctuation(words2)
	joined = fixSingleQuotes(joined)
	return normalizeSpaces(joined)
}

// ─────────────────────────────────────────────
// Q24 – Read File
// ─────────────────────────────────────────────
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ─────────────────────────────────────────────
// Q25 – Write File
// ─────────────────────────────────────────────
func writeFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// ═════════════════════════════════════════════
// IMAGE PROBLEMS
// ═════════════════════════════════════════════

// ─────────────────────────────────────────────
// Image 2 – ValidateBanner
// A well-formed banner map must have exactly 95 entries
// (printable ASCII runes 32–126), each with exactly 8 strings.
// ─────────────────────────────────────────────
func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95 {
		return fmt.Errorf("banner has %d entries, expected 95", len(banner))
	}
	// Check each rune 32–126 in consistent order
	for r := rune(32); r <= 126; r++ {
		lines, ok := banner[r]
		if !ok {
			return fmt.Errorf("missing character: %q (ASCII %d)", r, r)
		}
		if len(lines) != 8 {
			return fmt.Errorf("character %q has %d lines, expected 8", r, len(lines))
		}
	}
	return nil
}

// ─────────────────────────────────────────────
// Image 1 – MergeBanners
// Return a new map with all entries from base, overridden by priority.
// Neither input map is modified.
// ─────────────────────────────────────────────
func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)
	for k, v := range base {
		result[k] = v
	}
	for k, v := range priority {
		result[k] = v // priority wins on duplicate keys
	}
	return result
}

// ─────────────────────────────────────────────
// main – quick smoke tests
// ─────────────────────────────────────────────
func main() {
	// Q1
	v, _ := hexToDecimal("1E")
	fmt.Println("Q1 hexToDecimal(\"1E\"):", v) // 30

	// Q2
	v, _ = binToDecimal("1010")
	fmt.Println("Q2 binToDecimal(\"1010\"):", v) // 10

	// Q3
	fmt.Println("Q3 cap(\"hELLO\"):", cap("hELLO")) // Hello

	// Q4
	fmt.Println("Q4 upperN:", upperN([]string{"How", "are", "you"}, 2)) // How ARE YOU

	// Q5
	fmt.Println("Q5 punc(\",\"):", punc(",")) // true

	// Q6
	fmt.Println("Q6 fixArticles:", fixArticles([]string{"hello", ",", "world", "!"})) // hello, world!

	// Q7
	fmt.Println("Q7 article(\"apple\"):", article("apple"))   // an
	fmt.Println("Q7 article(\"book\"):", article("book"))     // a
	fmt.Println("Q7 article(\"honest\"):", article("honest")) // an

	// Q8
	fmt.Println("Q8:", fixSentenceArticles("There it was. A amazing rock. A honest man."))

	// Q9
	fmt.Println("Q9:", fixSingleQuotes("' awesome '")) // 'awesome'

	// Q10
	fmt.Println("Q10 toUpper:", toUpper("hello")) // HELLO

	// Q11
	fmt.Println("Q11 toLower:", toLower("HELLO")) // hello

	// Q12
	fmt.Println("Q12 isCommand(\"(up)\"):", isCommand("(up)"))   // true
	fmt.Println("Q12 isCommand(\"hello\"):", isCommand("hello")) // false

	// Q15
	fmt.Println("Q15 isGroupPunctuation(\"...\"):", isGroupPunctuation("...")) // true
	fmt.Println("Q15 isGroupPunctuation(\".\"):", isGroupPunctuation("."))     // false

	// Q16
	fmt.Println("Q16 upperNSlice:", upperNSlice([]string{"this", "is", "cool"}, 2)) // [this IS COOL]

	// Q17
	fmt.Println("Q17 lowerN:", lowerN([]string{"I", "LOVE", "GO"}, 2)) // [I love go]

	// Q18
	fmt.Println("Q18 capN:", capN([]string{"welcome", "to", "the", "brooklyn", "bridge"}, 2)) // [...Brooklyn Bridge]

	// Q19
	fmt.Println("Q19 processHex:", processHex([]string{"1E", "(hex)", "files"})) // [30 files]

	// Q20
	fmt.Println("Q20 processBin:", processBin([]string{"10", "(bin)", "years"})) // [2 years]

	// Q21
	fmt.Println("Q21 normalizeSpaces:", normalizeSpaces("hello   world")) // hello world

	// Q22
	fmt.Println("Q22 isVowel('a'):", isVowel('a')) // true
	fmt.Println("Q22 isVowel('b'):", isVowel('b')) // false

	// Q23
	fmt.Println("Q23 processSentence:", processSentence("1E (hex) files were added (up, 2)"))

	// ValidateBanner
	fmt.Println("\nValidateBanner(nil):", ValidateBanner(nil))

	// MergeBanners
	base := map[rune][]string{'A': {"line1"}}
	pri := map[rune][]string{'A': {"override"}, 'B': {"new"}}
	merged := MergeBanners(base, pri)
	fmt.Println("MergeBanners A:", merged['A']) // [override]
	fmt.Println("MergeBanners B:", merged['B']) // [new]
}
