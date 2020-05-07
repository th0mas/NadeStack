package csgo

// Maps are supported active duty maps. Workshop maps can be used dynamically
var maps = []string{
	"de_mirage",
	"de_dust2",
	"de_cache",
	"de_inferno",
	"de_vertigo",
	"de_overpass",
}

func IsAvailiableMap(m string) bool {
	for _, v := range maps {
		if m == v {
			return true
		}
	}

	return false
}
