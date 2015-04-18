package unflatten

import "github.com/imdario/mergo"

// TokenizerFunc This function is used to tokenize the keys in the flattened data structure.
//
// The following example uses strings.Split to tokenize based on .
//  func(k string) []string { return strings.Split(k, ".") }
type TokenizerFunc func(string) []string

// Unflatten This function will unflatten a map with keys which are comprised of multiple tokens which
// are segmented by the tokenizer function.
func Unflatten(m map[string]interface{}, tf TokenizerFunc) map[string]interface{} {
	tree := make(map[string]interface{})

	c := make(chan map[string]interface{})

	go mapify(m, c, tf)

	for n := range c {
		mergo.Merge(&tree, n)
	}

	return tree
}

func mapify(m map[string]interface{}, c chan map[string]interface{}, tf TokenizerFunc) {

	for k, v := range m {

		tokens := tf(k)

		var (
			z map[string]interface{}
			t string
		)

		// we are going to use pop to go backwards through the tokens
		for {
			// pop
			t, tokens = tokens[len(tokens)-1], tokens[:len(tokens)-1]

			// start by appending the actual value.
			if z == nil {
				z = map[string]interface{}{
					t: v,
				}
				continue
			}

			z = map[string]interface{}{
				t: z,
			}

			// all done?
			if len(tokens) == 0 {
				c <- z
				break
			}
		}
	}

	close(c)
}
