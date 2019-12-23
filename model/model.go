package model

type RspamdStatistics struct {
	ReadOnly bool  `json:"read_only"`
	Scanned  int32 `json:"scanned"`
	Learned  int32 `json:"learned"`
	Actions struct {
		Reject         int64 `json:"reject"`
		SoftReject     int64 `json:"soft reject"`
		RewriteSubject int64 `json:"rewrite subject"`
		AddHeader      int64 `json:"add header"`
		Greylist       int64 `json:"greylist"`
		NoAction       int64 `json:"no action"`
	} `json:"actions"`
	SpamCount             int64 `json:"spam_count"`
	HamCount              int64 `json:"ham_count"`
	Connections           int64 `json:"connections"`
	ControlConnections    int64 `json:"control_connections"`
	PoolsAllocated        int64 `json:"pools_allocated"`
	PoolsFreed            int64 `json:"pools_freed"`
	BytesAllocated        int64 `json:"bytes_allocated"`
	ChunksAllocated       int64 `json:"chunks_allocated"`
	SharedChunksAllocated int64 `json:"shared_chunks_allocated"`
	ChunksFreed           int64 `json:"chunks_freed"`
	ChunksOversized       int64 `json:"chunks_oversized"`
	Fragmented            int64 `json:"fragmented"`
	TotalLearns           int64 `json:"total_learns"`
	Statfiles []struct {
		Revision  int32  `json:"revision"`
		Used      int32  `json:"used"`
		Total     int32  `json:"total"`
		Size      int64  `json:"size"`
		Symbol    string `json:"symbol"`
		Type      string `json:"type"`
		Languages int32  `json:"languages"`
		Users     int32  `json:"users"`
	} `json:"statfiles"`
	FuzzyHashes struct {
		RspamdCom int `json:"rspamd.com"`
	} `json:"fuzzy_hashes"`
}

type InfluxDbSettings struct {
	dbUrl    string
	username string
	password string
	db       string
}

type RspamdSettings struct {
	rspamdUrl      string
	rspamdPassword string
}

type Settings struct {
	Interval         int
	RspamdSettings   RspamdSettings
	InfluxDbSettings InfluxDbSettings
}
