package vimparser

import (
	"bufio"
	"github.com/despreston/vimlytics/internal/plugged"
	"github.com/despreston/vimlytics/internal/vimoptions"
	"log"
	"regexp"
	"strings"
)

type Setting struct {
	Option      string `json:"option"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Duplicate   bool   `json:"duplicate"`
}

type Summary struct {
	Settings []Setting        `json:"settings"`
	Leader   string           `json:"leader"`
	Plugged  []plugged.Plugin `json:"plugged"`
	Vundle   []plugged.Plugin `json:"vundle"`
}

func parseSetting(line string) (Setting, int) {
	parsed := strings.Split(strings.TrimPrefix(line, "set "), "=")
	tabsSpacesReg := regexp.MustCompile(` |\t+`)

	// split on spaces again in case there are comments after the value
	name := tabsSpacesReg.Split(parsed[0], -1)[0]

	var i int
	var setting Setting

	// Lookup setting based on:
	// 1. long form
	// 2. short abbreviation form
	// 3. command prefixed with "no" e.g. "set noshowmode"
	if v, has := vimoptions.Longs[name]; has {
		i = v
	} else if v, has = vimoptions.Shorts[name]; has {
		i = v
	} else if strings.HasPrefix(name, "no") {
		if v, has = vimoptions.Longs[strings.TrimPrefix(name, "no")]; has {
			i = v
		}
	}

	if i == 0 {
		return setting, 0
	}

	desc := vimoptions.Descriptions[i-1]

	setting = Setting{
		Option:      name,
		Description: desc,
	}

	if len(parsed) > 1 {
		setting.Value = parsed[1]
	}

	return setting, i
}

func Parse(f *bufio.Reader) Summary {
	scanner := bufio.NewScanner(f)
	indices := map[int]bool{}

	var summary Summary
	var settings []Setting
	var leader string
	var pluggedList []string
	var vundleList []string

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")

		// Settings
		if strings.HasPrefix(line, "set ") {
			setting, i := parseSetting(line)

			if i > 0 {
				if indices[i] {
					setting.Duplicate = true
				}

				indices[i] = true
				settings = append(settings, setting)
			} else {
				log.Printf("No setting for line: %v\n", line)
			}

			continue
		}

		// Leader key
		if strings.HasPrefix(line, "let mapleader") {
			leader = strings.TrimPrefix(line, "let mapleader = ")
			continue
		}

		// vim-plugged
		if strings.HasPrefix(line, "Plug ") {
			name := strings.Trim(strings.TrimPrefix(line, "Plug "), "'")
			pluggedList = append(pluggedList, name)
			continue
		}

		// Vundle
		if strings.HasPrefix(line, "Plugin ") {
			name := strings.Trim(strings.TrimPrefix(line, "Plugin "), "'")
			vundleList = append(vundleList, name)
		}
	}

	summary.Settings = settings
	summary.Leader = leader
	summary.Plugged = plugged.FetchAll(pluggedList)
	summary.Vundle = plugged.FetchAll(vundleList)

	return summary
}
