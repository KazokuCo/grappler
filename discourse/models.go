package discourse

// I don't even know why I'm grabbing half of these fields.
// Not even bothering parsing any dates, screw that.

type Post struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Username          string `json:"username"`
	AvatarTemplate    string `json:"avatar_template"`
	Cooked            string `json:"cooked"`
	PostNumber        int    `json:"post_number"`
	PostType          int    `json:"post_type"`
	ReplyCount        int    `json:"reply_count"`
	ReplyToPostNumber *int   `json:"reply_to_post_number"`
	QuoteCount        int    `json:"quote_count"`
	IncomingLinkCount int    `json:"incoming_link_count"`
	Reads             int    `json:"reads"`
	Score             int    `json:"score"`
	TopicID           int64  `json:"topic_id"`
	TopicSlug         string `json:"topic_slug"`
	DisplayUsername   string `json:"display_username"`
	Version           int    `json:"version"`
	UserTitle         string `json:"user_title"`
	Moderator         bool   `json:"moderator"`
	Admin             bool   `json:"admin"`
	Staff             bool   `json:"staff"`
	UserID            int64  `json:"id"`
	Hidden            bool   `json:"hidden"`
	TrustLevel        int    `json:"trust_level"`
	Wiki              bool   `json:"wiki"`
}

type Topic struct {
	PostStream struct {
		Posts []Post `json:"posts"`
	} `json:"post_stream"`

	ID               int64  `json:"id"`
	Title            string `json:"title"`
	FancyTitle       string `json:"fancy_title"`
	PostsCount       int    `json:"posts_count"`
	Views            int    `json:"views"`
	ReplyCount       int    `json:"reply_count"`
	ParticipantCount int    `json:"participant_count"`
	LikeCount        int    `json:"like_count"`
	Visible          bool   `json:"visible"`
	Closed           bool   `json:"closed"`
	Archived         bool   `json:"archived"`
	HasSummary       bool   `json:"summary"`
	Archetype        string `json:"archetype"`
	Slug             string `json:"slug"`
	CategoryID       int    `json:"category_id"`
	WordCount        int    `json:"word_count"`
	UserID           int    `json:"user_id"`
	PinnedGlobally   bool   `json:"pinned_globally"`
	Pinned           bool   `json:"pinned"`
}

type User struct {
	ID                int64  `json:"id"`
	Username          string `json:"username"`
	AvatarTemplate    string `json:"avatar_template"`
	Name              string `json:"name"`
	Bio               string `json:"bio_raw"`
	ProfileBackground string `json:"profile_background"`
	CardBackground    string `json:"card_background"`
	TrustLevel        int    `json:"trust_level"`
	Moderator         bool   `json:"moderator"`
	Admin             bool   `json:"admin"`
	Title             string `json:"title"`
}

type Envelope struct {
	Topic Topic `json:"topic"`
	User  User  `json:"user"`
	Post  Post  `json:"post"`
}
