package numstar

import (
	"strings"
)

// NumStar is numstar
type NumStar struct {
	space   string
	star    string
	gapSize int
}

// New returns NumStar pointer
func New() *NumStar {
	return &NumStar{
		space:   " ",
		star:    "*",
		gapSize: 2,
	}
}

// Space sets space char
func (ns *NumStar) Space(s string) {
	ns.space = s
}

// Star sets star char
func (ns *NumStar) Star(s string) {
	ns.star = s
}

// Gap sets gap size
func (ns *NumStar) Gap(g int) {
	ns.gapSize = g
}

// Big converts a string of digits into big one with default space size
func (ns *NumStar) Big(n string) string {
	return ns.BigS(n, ns.gapSize)
}

// BigS converts a string of digits into big one with spaces of s
func (ns *NumStar) BigS(n string, s int) string {
	d := strings.Split(n, "")
	buff := ""
	for _, v := range d {
		buff = ns.Concat(buff, ns.BigNum(v), s)
	}

	return buff
}

// Concat returns 2 big strings with spacing of size s
func (ns *NumStar) Concat(m, n string, s int) string {
	mLines := strings.Split(m, "\n")
	nLines := strings.Split(n, "\n")
	mLineCount := len(mLines)
	nLineCount := len(nLines)
	mWidth := len(mLines[0])
	nWidth := len(nLines[0])

	if mWidth == 0 {
		return n
	}

	if nWidth == 0 {
		return m
	}

	spaces := strings.Repeat(ns.space, s)
	if mLineCount < nLineCount {
		for i := mLineCount; i < nLineCount; i++ {
			mLines = append(mLines, strings.Repeat(ns.space, mWidth))
		}
	} else {
		if nLineCount < mLineCount {
			for i := nLineCount; i < mLineCount; i++ {
				nLines = append(nLines, strings.Repeat(ns.space, nWidth))
			}
		}
	}

	buff := []string{}

	for i, n := 0, len(mLines); i < n; i++ {
		buff = append(buff, mLines[i]+spaces+nLines[i])
	}

	return strings.Join(buff, "\n")
}

// BigNum returns a number with big size
func (ns *NumStar) BigNum(n string) string {
	num := map[string]string{
		"1": `*
*
*
*
*`,
		"2": ` *** 
*   *
  ** 
 *   
*****`,
		"3": ` *** 
*   *
   * 
*   *
 *** `,
		"4": `   * 
  ** 
 * * 
*****
   * `,
		"5": `*****
*    
**** 
    *
**** `,
		"6": ` *** 
*    
**** 
*   *
 *** `,
		"7": `*****
    *
   * 
  *  
 *   `,
		"8": ` *** 
*   *
 *** 
*   *
 *** `,
		"9": ` *** 
*   *
 ****
    *
 *** `,
		"0": ` *** 
*   *
*   *
*   *
 *** `,
	}

	if v, ok := num[n]; ok {
		return strings.Replace(strings.Replace(v, " ", ns.space, -1), "*", ns.star, -1)
	}
	return ""
}
