package main

import (
	"fmt"
	"github.com/zach-klippenstein/goregen"
	"math/rand"
	"os"
	"regexp"
	"regexp/syntax"
)

func main() {
	// Example 1
	pattern := "[ab]{5}"
	str, _ := regen.Generate(pattern)

	if matched, _ := regexp.MatchString(pattern, str); matched {
		fmt.Println("Matches!")
	}

	// Example 2
	pattern = "[ab]{5}"
	generator, _ := regen.NewGenerator(pattern, &regen.GeneratorArgs{
		RngSource: rand.NewSource(0),
	})
	str = generator.Generate()
	if matched, _ := regexp.MatchString(pattern, str); matched {
		fmt.Println("Matches!")
	}

	// Example 3
	pattern = `Hello, (?P<firstname>[A-Z][a-z]{2,10}) (?P<lastname>[A-Z][a-z]{2,10})`
	generator, _ = regen.NewGenerator(pattern, &regen.GeneratorArgs{
		Flags: syntax.Perl,
		CaptureGroupHandler: func(index int, name string, group *syntax.Regexp, generator regen.Generator, args *regen.GeneratorArgs) string {
			if name == "firstname" {
				return fmt.Sprintf("FirstName (e.g. %s)", generator.Generate())
			}
			return fmt.Sprintf("LastName (e.g. %s)", generator.Generate())
		},
	})

	// Print to stderr since we're generating random output and can't assert equality.
	fmt.Fprintln(os.Stderr, generator.Generate())
}
