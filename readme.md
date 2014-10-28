Abrev is an abbreviation expander suitable for command driven user
interfaces.

Abrev recognizes an abbreviation as a subsequence of letters in the full
word.  It returns a list of all matching words unless the the input is found
to be a full unabbreviated word, in which case it returns just that word.
The idea is that the rule is simple and unambiguous and allows flexability
with no predefined list of abbreviations.

Example: Command set {open export expand pan please}

* op → open, it's unique
* o → {open export} both words contain an o
* xr → export. a subsequence is not a substring
* xn → expand
* pan → pan, a full match excludes partial matches
* plz → &lt;nothing> no magic
