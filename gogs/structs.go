package gogs

import (
	"net/http"
	"time"
)

// Struct for handling the data received from the API endpoint
type ApiResult struct {
	Code        int
	Status      string
	Error       string
	ErrorString string
	Results     interface{}
}

// Struct for handling raw data received from the API endpoint
type RawApiResult struct {
	Code        int
	Status      string
	Error       string
	ErrorString string
	ResultData  []byte
}

// Struct for storing the credentials we receive from the REST API
type RestCredentials struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// Contains all the relevant credentials for both the REST and Socket API's
type Credentials struct {
	ClientID         string
	ClientSecret     string
	Username         string
	Password         string
	UserID           string
	ChatAuth         string
	UserJwt          string
	NotificationAuth string
	RestCredentials
}

// The primary server struct that contains all state related to the connection to OGS
type Server struct {
	IsAuthed    bool
	ApiVersion  string
	BaseUrl     string
	httpClient  *http.Client
	Credentials *Credentials
}

type PlayerId int

type Player struct {
	ID           PlayerId `json:"id"`
	Username     string   `json:"username"`
	Country      string   `json:"country"`
	Icon         string   `json:"icon"`
	Ratings      Ratings  `json:"ratings"`
	Ranking      float64  `json:"ranking"`
	Professional bool     `json:"professional"`
	UIClass      string   `json:"ui_class"`
}

type Ratings struct {
	Version int `json:"version"`
	Overall struct {
		Rating     float64 `json:"rating"`
		Deviation  float64 `json:"deviation"`
		Volatility float64 `json:"volatility"`
	} `json:"overall"`
}

// Game object from the /games endpoint
type Game struct {
	ID         int        `json:"id"`
	AllPlayers []PlayerId `json:"all_players"`
	Name       string     `json:"name"`
	Players    struct {
		Black Player `json:"black"`
		White Player `json:"white"`
	} `json:"players"`
	Related struct {
		Reviews string `json:"reviews"`
	} `json:"related"`
	Creator                PlayerId   `json:"creator"`
	Mode                   string     `json:"mode"`
	Source                 string     `json:"source"`
	Black                  PlayerId   `json:"black"`
	White                  PlayerId   `json:"white"`
	Width                  int        `json:"width"`
	Height                 int        `json:"height"`
	Rules                  string     `json:"rules"`
	Ranked                 bool       `json:"ranked"`
	Handicap               int        `json:"handicap"`
	HandicapRankDifference string     `json:"handicap_rank_difference"`
	Komi                   string     `json:"komi"`
	TimeControl            string     `json:"time_control"`
	BlackPlayerRank        int        `json:"black_player_rank"`
	BlackPlayerRating      string     `json:"black_player_rating"`
	WhitePlayerRank        int        `json:"white_player_rank"`
	WhitePlayerRating      string     `json:"white_player_rating"`
	TimePerMove            int        `json:"time_per_move"`
	TimeControlParameters  string     `json:"time_control_parameters"`
	DisableAnalysis        bool       `json:"disable_analysis"`
	Tournament             *int       `json:"tournament,omitempty"` //can be null or a Tournament ID
	TournamentRound        int        `json:"tournament_round"`
	Ladder                 *int       `json:"ladder,omitempty"` //can be null or the LadderId
	PauseOnWeekends        bool       `json:"pause_on_weekends"`
	Outcome                string     `json:"outcome"`
	BlackLost              bool       `json:"black_lost"`
	WhiteLost              bool       `json:"white_lost"`
	Annulled               bool       `json:"annulled"`
	AnnulmentReason        any        `json:"annulment_reason"` //TODO Find annulled game to get true type
	Started                time.Time  `json:"started"`
	Ended                  *time.Time `json:"ended,omitempty"` //can be null or a timestamp
	HistoricalRatings      struct {
		Black Player `json:"black"`
		White Player `json:"white"`
	} `json:"historical_ratings"`
	Gamedata            Gamedata `json:"gamedata"`
	Auth                string   `json:"auth"`
	Rengo               bool     `json:"rengo"`
	Flags               any      `json:"flags"`
	BotDetectionResults any      `json:"bot_detection_results"`
}

type Gamedata struct {
	AgaHandicapScoring     bool          `json:"aga_handicap_scoring"`
	AllowKo                bool          `json:"allow_ko"`
	AllowSelfCapture       bool          `json:"allow_self_capture"`
	AllowSuperko           bool          `json:"allow_superko"`
	AutomaticStoneRemoval  bool          `json:"automatic_stone_removal"`
	BlackPlayerID          PlayerId      `json:"black_player_id"`
	Clock                  GamedataClock `json:"clock"`
	DisableAnalysis        bool          `json:"disable_analysis"`
	EndTime                int           `json:"end_time"`
	FreeHandicapPlacement  bool          `json:"free_handicap_placement"`
	GameID                 int           `json:"game_id"`
	GameName               string        `json:"game_name"`
	GroupIds               []int         `json:"group_ids"`
	Handicap               int           `json:"handicap"`
	HandicapRankDifference int           `json:"handicap_rank_difference"`
	Height                 int           `json:"height"`
	InitialPlayer          string        `json:"initial_player"`
	InitialState           struct {
		Black string `json:"black"`
		White string `json:"white"`
	} `json:"initial_state"`
	Komi                          float64    `json:"komi"`
	Moves                         []GameMove `json:"moves"`
	OpponentPlaysFirstAfterResume bool       `json:"opponent_plays_first_after_resume"`
	OriginalDisableAnalysis       bool       `json:"original_disable_analysis"`
	PauseControl                  struct {
		StoneRemoval bool `json:"stone_removal"`
	} `json:"pause_control"`
	PauseOnWeekends bool                     `json:"pause_on_weekends"`
	PausedSince     int                      `json:"paused_since"`
	Phase           string                   `json:"phase"`
	PlayerPool      map[int]PlayerPoolPlayer `json:"player_pool"`
	Players         struct {
		Black GamedataPlayer `json:"black"`
		White GamedataPlayer `json:"white"`
	} `json:"players"`
	Private         bool     `json:"private"`
	Ranked          bool     `json:"ranked"`
	Removed         string   `json:"removed"`
	Rengo           bool     `json:"rengo"`
	RengoCasualMode bool     `json:"rengo_casual_mode"`
	RengoTeams      struct { //TODO: Find real type of rengo team slices
		Black []any `json:"black"`
		White []any `json:"white"`
	} `json:"rengo_teams"`
	Rules string `json:"rules"`
	Score struct {
		Black GamedataPlayerScore `json:"black"`
		White GamedataPlayerScore `json:"white"`
	} `json:"score"`
	ScoreHandicap        bool     `json:"score_handicap"`
	ScorePasses          bool     `json:"score_passes"`
	ScorePrisoners       bool     `json:"score_prisoners"`
	ScoreStones          bool     `json:"score_stones"`
	ScoreTerritory       bool     `json:"score_territory"`
	ScoreTerritoryInSeki bool     `json:"score_territory_in_seki"`
	StartTime            int      `json:"start_time"`
	StrictSekiMode       bool     `json:"strict_seki_mode"`
	SuperkoAlgorithm     string   `json:"superko_algorithm"`
	TimeControl          struct { //TODO: This will likely need to become an interface to deal with the different time controls
		System          string `json:"system"`
		TimeControl     string `json:"time_control"`
		Speed           string `json:"speed"`
		PauseOnWeekends bool   `json:"pause_on_weekends"`
	} `json:"time_control"`
	WhiteMustPassLast bool     `json:"white_must_pass_last"`
	WhitePlayerID     PlayerId `json:"white_player_id"`
	Width             int      `json:"width"`
	Winner            PlayerId `json:"winner"`
}

type PlayerPoolPlayer struct {
	Username               string  `json:"username"`
	Rank                   float64 `json:"rank"`
	Professional           bool    `json:"professional"`
	ID                     int     `json:"id"`
	AcceptedStones         string  `json:"accepted_stones"`
	AcceptedStrictSekiMode bool    `json:"acceppted_strict_seki_mode"`
}

type GamedataClock struct {
	GameID                 int          `json:"game_id"`
	CurrentPlayer          PlayerId     `json:"current_player"`
	BlackPlayerID          PlayerId     `json:"black_player_id"`
	WhitePlayerID          PlayerId     `json:"white_player_id"`
	Title                  string       `json:"title"`
	LastMove               int64        `json:"last_move"`
	Expiration             int64        `json:"expiration"`
	BlackTime              GamedataTime `json:"black_time"`
	WhiteTime              GamedataTime `json:"white_time"`
	PausedSince            int          `json:"paused_since"`
	PauseDelta             int          `json:"pause_delta"`
	ExpirationDelta        int          `json:"expiration_delta"`
	StoneRemovalMode       bool         `json:"stone_removal_mode"`
	StoneRemovalExpiration int          `json:"stone_removal_expiration"`
}

type GamedataTime struct {
	ThinkingTime float64 `json:"thinking_time"`
	SkipBonus    bool    `json:"skip_bonus"`
}

type GameMove []interface{} // Move is structured as [int, int, float64]

type GamedataPlayer struct {
	Username     string  `json:"username"`
	Rank         float64 `json:"rank"`
	Professional bool    `json:"professional"`
	ID           int     `json:"id"`
}

type GamedataPlayerScore struct {
	Total            float64 `json:"total"`
	Stones           int     `json:"stones"`
	Territory        int     `json:"territory"`
	Prisoners        int     `json:"prisoners"`
	ScoringPositions string  `json:"scoring_positions"`
	Handicap         int     `json:"handicap"`
	Komi             float64 `json:"komi"`
}
