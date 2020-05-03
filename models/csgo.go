package models

// CsgoMatch is a struct reflecting the config format for get_5 plugin, and should be encoded via JSON
// Possibly could persist some of these fields to database in future, so should mke sure ID fields etc are unique.
type CsgoMatch struct {
	MatchID        string   `json:"matchid"`
	NumMaps        int      `json:"num_maps"`
	MapList        []string `json:"map_list"`
	SkipVeto       bool     `json:"skip_veto"`
	SideType       string   `json:"side_type"`
	PlayersPerTeam int      `json:"players_per_team"`
	CVars          []string `json:"cvars"`
	TeamOne        CsgoTeam `json:"team1"`
	TeamTwo        CsgoTeam `json:"team2"`
}

// CsgoTeam is a team in a CsgoMatch. Needs a name and a list of players.
// Other fields that can be added in future/if needed are: tag, flag, logo
type CsgoTeam struct {
	Name    string   `json:"name"`
	Players []string `json:"players"` // Players is a list of STEAMIDs on that team - prefer steamID 64

}
