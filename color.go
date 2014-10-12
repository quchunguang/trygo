package trygo

import (
	"strconv"
	"strings"
)

// 16 color for ANSI/VT100 Terminal Control Escape Sequences
// http://www.termsys.demon.co.uk/vtansi.htm
const (
	CRK = "\033[30m" // blacK - Regular
	CRR = "\033[31m" // Red
	CRG = "\033[32m" // Green
	CRY = "\033[33m" // Yellow
	CRB = "\033[34m" // Blue
	CRP = "\033[35m" // Purple
	CRC = "\033[36m" // Cyan
	CRW = "\033[37m" // White

	COK = "\033[1;30m" // blacK - bOld
	COR = "\033[1;31m" // Red
	COG = "\033[1;32m" // Green
	COY = "\033[1;33m" // Yellow
	COB = "\033[1;34m" // Blue
	COP = "\033[1;35m" // Purple
	COC = "\033[1;36m" // Cyan
	COW = "\033[1;37m" // White

	CUK = "\033[4;30m" // blacK - Underline
	CUR = "\033[4;31m" // Red
	CUG = "\033[4;32m" // Green
	CUY = "\033[4;33m" // Yellow
	CUB = "\033[4;34m" // Blue
	CUP = "\033[4;35m" // Purple
	CUC = "\033[4;36m" // Cyan
	CUW = "\033[4;37m" // White

	CBK = "\033[40m" // blacK - Background
	CBR = "\033[41m" // Red
	CBG = "\033[42m" // Green
	CBY = "\033[43m" // Yellow
	CBB = "\033[44m" // Blue
	CBP = "\033[45m" // Purple
	CBC = "\033[46m" // Cyan
	CBW = "\033[47m" // White

	CRD = "\033[0m" // Reset Default
)

// Return 256 (8-bit) Colors for ANSI escape code extended.
//     0 <= forcolor < 256, -1 for default
//     0 <= bkcolor < 256, -1 for default
// Reference:
// http://web.archive.org/web/20131009193526/http://bitmote.com/index.php?post/2012/11/19/Using-ANSI-Color-Codes-to-Colorize-Your-Bash-Prompt-on-Linux
// http://en.wikipedia.org/wiki/ANSI_escape_code
func ColorStr(s string, forcolor int, bkcolor int) (ret string) {
	var collect []string
	if forcolor >= 0 && forcolor < 256 {
		collect = append(collect, "38", "5", strconv.Itoa(forcolor))
	}
	if bkcolor >= 0 && bkcolor < 256 {
		collect = append(collect, "48", "5", strconv.Itoa(bkcolor))
	}
	if len(collect) > 0 {
		ret = "\033["
		ret += strings.Join(collect, ";")
		ret += "m" + s + "\033[m"
	} else {
		ret = s
	}
	return
}

func RedStr(s string) (ret string) {
	ret = CRR + s + CRD
	return
}

func GreenStr(s string) (ret string) {
	ret = CRG + s + CRD
	return
}

func BlueStr(s string) (ret string) {
	ret = CRB + s + CRD
	return
}

func YellowStr(s string) (ret string) {
	ret = CRY + s + CRD
	return
}

func BoldStr(s string) (ret string) {
	ret = COW + s + CRD
	return
}
