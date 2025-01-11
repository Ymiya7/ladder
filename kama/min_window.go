package kama

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	start, end := 0, 0
	minLen := len(s) + 1
	idx := start
	for end < len(s) {
		sMap[s[end]]++
		for valid(sMap, tMap) {
			if end-start+1 < minLen {
				idx = start
				minLen = end - start + 1
			}
			sMap[s[start]]--
			if sMap[s[start]] == 0 {
				delete(sMap, s[start])
			}
			start++
		}
		end++
	}
	if minLen == len(s)+1 {
		return ""
	}
	return s[idx : idx+minLen]
}

func valid(sMap map[byte]int, tMap map[byte]int) bool {
	for bt, cnt := range tMap {
		if _, ok := sMap[bt]; !ok {
			return false
		}
		if cnt > sMap[bt] {
			return false
		}
	}
	return true
}
