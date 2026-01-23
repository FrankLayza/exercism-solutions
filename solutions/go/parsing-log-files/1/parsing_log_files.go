package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	startsTRC := strings.HasPrefix(text, "[TRC]")
	startsDBG := strings.HasPrefix(text, "[DBG]")
	startsINF := strings.HasPrefix(text, "[INF]")
	startsWRN := strings.HasPrefix(text, "[WRN]")
	startsERR := strings.HasPrefix(text, "[ERR]")
	startsFTL := strings.HasPrefix(text, "[FTL]")
	if(startsTRC || startsDBG || startsERR || startsFTL || startsINF || startsWRN){
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	splittedText := re.Split(text, -1)
	return splittedText
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, value := range lines{
		if re.MatchString(value){
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := []string{}
	re := regexp.MustCompile(`User\s+(\S+)`)
	for _, value := range lines{
		foundIt := re.FindStringSubmatch(value)
		if foundIt == nil{
			result = append(result, value)
		}else{
			username := foundIt[1]
			newValue := fmt.Sprintf("[USR] %v %v", username,value)
			result = append(result, newValue)
		}
	}
	return result
}
