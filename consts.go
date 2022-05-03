package main

type Species struct {
	name        string
	probability float64
}

var (
	languages = []string{
		"Low Imperial",
		"High Imperial",
		"Khadish",
		"Selan",
		"Varlish",
		"Veton",
		"Fae",
		"Grimmark",
		"Draconic",
	}

	speciess = []Species{
		{"Human", 0.8},
		{"Dwarf", 0.2},
	}

	affiliations = []string{
		"Seekers of the Unknown Fire",
		"Church of the New Sun",
		"Witch Cult",
		"Unseeing Monks",
		"Crimson Court",
		"Foremost Armigers",
		"Great House vet-Doranar",
		"Great House vet-Elvan",
		"Great House vet-Haan",
		"Low House vet-Aldar",
		"Low House vet-Akash",
		"Explorers' Guild",
		"Black Dogs",
		"Scrivenwood Varlings",
		"Guild of Torturers",
		"White Cloaks",
		"Wardens of the Crystal Tower",
		"Silver Mountain Tribe",
		"Smoldering Heart Tribe",
		"Crystal Lake Tribe",
		"Star Watchers",
	}

	attunementLocations = []string{
		"Caverns of Zakhar",
		"Silver Mountain",
		"Crystal Lake",
		"Scrivenwood",
		"Isolate Tower",
	}
)
