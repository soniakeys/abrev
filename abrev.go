// Public domain.

// Abrev is an abbreviation expander suitable for command driven user
// interfaces.
//
// Abrev recognizes an abbreviation as a subsequence of letters in the full
// word.  It returns a list of all matching words unless the the input is found
// to be a full unabbreviated word, in which case it returns just that word.
// The idea is that the rule is simple and unambiguous and allows flexability
// with no predefined list of abbreviations.
package abrev

// History:  Decades ago a friend who worked at Prime computer was telling me
// about arguments over the command set of some command driven system.
// I came up with this suggestion off the top of my head.  My friend reported
// back to me later that my idea had gone into production on something or
// other!  It must have made sense to someone at the time at least.

import "strings"

type Set []string

func (s Set) Expand(abrev string) (matches []string) {
	for _, full := range s {
		switch {
		case abrev == full:
			return []string{full}
		case match(abrev, full):
			matches = append(matches, full)
		}
	}
	return
}

func match(a, f string) bool {
	for len(a) > 0 {
		x := strings.Index(f, a[:1])
		if x < 0 {
			return false
		}
		a = a[1:]
		f = f[x+1:]
	}
	return true
}
