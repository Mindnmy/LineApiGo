// Copyright (c) 2020 @Ch31212y
// Version 1.1 beta
// LastUpdate 2020/08/28

package lineapigo

// HOST URL
var lineHostURL string = "https://ga2.line.naver.jp"

// TALK PATH
var normal string = "/S4"

//POLLING PATH
var longPolling string = "/P4"

var authRegistration string = "/api/v4p/rs"
var newRegistration string = "/acct/lp/lgn/sq/v1"
var secondaryQrLogin string = "/acct/lgn/sq/v1"

var systemName string = "LineApiGo"

var systemVersion map[string]string = map[string]string{
	"LITE":   "10.0",
	"MAC":    "10.15.1",
	"CHROME": "1",
	"IOS":    "113.7",
	// "CHROME": "81.0",
}

var appVersion map[string]string = map[string]string{
	"LITE":   "2.14.0",
	"MAC":    "5.24.1",
	"CHROME": "2.3.9",
	"IOS":    "10.1.1",
}

// GetUserAgent This func will return UserAgent for line
// @param appType(string) string of your choiced line application
// @return string of User-Agent
func GetUserAgent(appType string) string {
	switch appType {
	case "LITE":
		return "LLA/" + systemVersion["LITE"] + " Galaxy Note 10+ " + systemVersion["LITE"]
	case "MAC":
		return "Line/" + systemVersion["MAC"]
	case "IOS":
		return "Line/" + appVersion["IOS"] + " iPhone8,1 " + systemVersion["IOS"]
	case "CHROME":
		return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"
	default:
		return GetUserAgent("LITE")
	}
}

// GetLineApplication This func will return string of X-Line-Application
// @param appType(string) string of your choiced line application type
// @return string of X-Line-Application
func GetLineApplication(appType string) string {
	switch appType {
	case "LITE":
		return "ANDROIDLITE\t" + appVersion["LITE"] + "\tAndroid OS\t" + systemVersion["LITE"]
	case "MAC":
		return "DESKTOPMAC\t" + appVersion["MAC"] + "\tOS X\t" + systemVersion["MAC"]
	case "IOS":
		return "IOS\t" + appVersion["IOS"] + "\tiPhone OS\t" + systemVersion["IOS"]
	case "CHROME":
		return "CHROMEOS\t" + appVersion["CHROME"] + "\tChrome_OS\t" + systemVersion["CHROME"]
	default:
		return GetLineApplication("LITE")
	}
}

// GetXLal This func will return string of x-lal
func GetXLal(appType string) string {
	switch appType {
	case "CHROME":
		return "ja"
	default:
		return "jp_ja"
	}
}
