package elastic

type IndexErrorResponse struct {
	Error struct {
		Index     string `json:"index"`
		IndexUUID string `json:"index_uuid"`
		Reason    string `json:"reason"`
		RootCause []struct {
			Index  string `json:"index"`
			Reason string `json:"reason"`
			Type   string `json:"type"`
		} `json:"root_cause"`
		Type string `json:"type"`
	} `json:"error"`
	Status int `json:"status"`
}

type IndexResponse struct {
	Acknowledged       bool   `json:"acknowledged"`
	Index              string `json:"index"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	ID     string                 `json:"_id"`
	Index  string                 `json:"_index"`
	Score  float64                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}

type Shards struct {
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
	Total      int `json:"total"`
}

type Response struct {
	Shards   Shards `json:"_shards"`
	HitsData Hits   `json:"hits"`
	TimedOut bool   `json:"timed_out"`
	Took     int    `json:"took"`
}

type Total struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}
