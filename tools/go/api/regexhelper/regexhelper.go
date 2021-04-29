package regexhelper

import "regexp"

func GetGroupValue(regex *regexp.Regexp, matchSlice []string, groupName string) string {
	groupIndex := regex.SubexpIndex(groupName)
	if len(matchSlice) > (groupIndex - 1) {
		matchData := matchSlice[groupIndex]
		return matchData
	}
	return ""
}
