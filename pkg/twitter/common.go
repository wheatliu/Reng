package twitter

import "time"

type TweetType string

const (
	// graphql url
	TweetDetailGraphqlURL = "https://x.com/i/api/graphql/U0HTv-bAWTBYylwEMT7x5A/TweetDetail"
	UserTweetsGraphqlURL  = "https://x.com/i/api/graphql/QWF3SzpHmykQHsQMixG0cg/UserTweets"
	UserFollowGraphqlURL  = "https://x.com/i/api/graphql/7oQrdmth4zE3EtD42ZxgOA/Following"

	// timeline instruction typename
	TimelineAddEntries = "TimelineAddEntries"
	// EntryIdPrefix
	NormalTweetPrefix       = "tweet-"                //tweet
	ConversationTweetPrefix = "profile-conversation-" //reply
	CursorPrefix            = "cursor-bottom"         //cursor
	// tweet created time format
	TimeFormat = time.RubyDate

	// twitter account status
	UserUnavailable = "UserUnavailable"
	Protected       = "Protected"

	// tweet type
	TweetTypeNormal  TweetType = "NORMAL"
	TweetTypeReply   TweetType = "REPLY"
	TweetTypeRetweet TweetType = "REPOST"
	TweetTypeQuote   TweetType = "QUOTE"
)

type Tweet struct {
	EntryID         string          `json:"entry_id"`
	RestID          string          `json:"rest_id"`
	FullText        string          `json:"full_text"`
	QuoteCount      int64           `json:"quote_count"`
	ReplyCount      int64           `json:"reply_count"`
	RetweetCount    int64           `json:"retweet_count"`
	FavoriteCount   int64           `json:"favorite_count"`
	ViewCount       int64           `json:"view_count"`
	CreatedAt       time.Time       `json:"created_at"`
	HashTags        []string        `json:"hashtags"`
	Retweeted       bool            `json:"retweeted"`
	Favorited       bool            `json:"favorited"`
	Followed        bool            `json:"followed"`
	Comment         []*TweetComment `json:"comment"`
	Type            TweetType       `json:"type"`
	ReferenceRestID string          `json:"reference_rest_id"`
}

type UserTweets struct {
	Tweets []*Tweet
	Cursor string
	Totals int
}

type TweetComment struct {
	UserID     string
	ScreenName string
	FullText   string
}
