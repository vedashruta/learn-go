package strings

import (
	"fmt"
	"strings"
)

func StringMethods() {
	print := fmt.Println
	var sample string = "This is a simple text"
	print("Contains this:", strings.Contains(sample, "this"))                          // false (case sesnsitive)
	print("Contains text:", strings.Contains(sample, "text"))                          // true
	print("Ocuurrence of i : ", strings.Count(sample, "i"))                            // 3
	print("Has Prefix is:", strings.HasPrefix(sample, "is"))                           // false
	print("Has Prefix This:", strings.HasPrefix(sample, "This"))                       // true
	print("Has Suffix ext:", strings.HasSuffix(sample, "ext"))                         // true
	print("Index of i:", strings.Index(sample, "i"))                                   // 2
	print("Joined String:", strings.Join([]string{"a", "b", "c", "d", "e", "f"}, "-")) // a-b-c-d-e-f
	print("Reapeat ab 4 times:", strings.Repeat("ab", 4))                              // abababab
	print("Replace 1st occurrence of is:", strings.Replace(sample, "is", "", 1))       // Th is a simple text
	print("Replace 0 occurrence of is:", strings.Replace(sample, "is", "", 0))         // This is a simple text
	print("Replace all occurrence of is:", strings.Replace(sample, "is", "", -1))      // Th  a simple text
	print(strings.Split(sample, " "))                                                  // [This is a simple text]
	print("Lowercase:", strings.ToLower(sample))                                       // this is a simple text
	print("Uppercase:", strings.ToUpper(sample))                                       // THIS IS A SIMPLE TEXT
	print("Trim - from main string: ", strings.Trim("-+"+sample+"-", "-"))             // +This is a simple text
}
