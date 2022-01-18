package services

import "strings"

// GetOsTypeFromUserAgent returns approximate os type determined based of UserAgent
func GetOsTypeFromUserAgent(userAgent string) string {
	tmpString := strings.Split(userAgent, "(")
	systemInfo := ""
	if len(tmpString) == 3 {
		tmpString2 := strings.Split(tmpString[1], ")")
		if len(tmpString2) == 2 {
			systemInfo = tmpString2[0]
		}
	}
	return systemInfo
}
