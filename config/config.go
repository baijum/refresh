/* Package config is taken from here:
https://github.com/pilu/config

The MIT License (MIT)

Copyright (c) 2013 Andrea Franz (http://gravityblast.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/

package config

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type Options map[string]string

type Sections map[string]Options

var commentSplitRegexp = regexp.MustCompile(`[#;]`)

var keyValueSplitRegexp = regexp.MustCompile(`(\s*(:|=)\s*)|\s+`)

func cleanLine(line string) string {
	chunks := commentSplitRegexp.Split(line, 2)
	return strings.TrimSpace(chunks[0])
}

func parse(reader *bufio.Reader, mainSectionName string) (Sections, error) {
	sections := make(Sections)
	section := mainSectionName
	options := make(Options)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return sections, err
		}

		line = cleanLine(line)

		if len(line) == 0 {
			continue
		}

		if line[0] == '[' && line[len(line)-1] == ']' {
			sections[section] = options
			section = line[1:(len(line) - 1)]
			options = sections[section] // check if section already exists
			if options == nil {
				options = make(Options)
			}
		} else {
			values := keyValueSplitRegexp.Split(line, 2)
			key := values[0]
			value := ""
			if len(values) == 2 {
				value = values[1]
			}

			options[key] = value
		}
	}

	sections[section] = options

	return sections, nil
}

func ParseFile(path string, mainSectionName string) (Sections, error) {
	file, err := os.Open(path)
	if err != nil {
		return make(Sections), err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	return parse(reader, mainSectionName)
}
