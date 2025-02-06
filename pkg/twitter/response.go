package twitter

type HashTag struct {
	Text string `json:"text"`
}

type Entities struct {
	HashTags []HashTag `json:"hashtags"`
}

type QuotedStatusResultData struct {
	RestID string `json:"rest_id"`
}
type QuotedStatusResult struct {
	Result *QuotedStatusResultData `json:"result"`
}

type RetweetedStatusResultData struct {
	RestID string `json:"rest_id"`
}
type RetweetedStatusResult struct {
	Result *RetweetedStatusResultData `json:"result"`
}

type Views struct {
	Count string `json:"count"`
}

type Legacy struct {
	Entities              Entities               `json:"entities"`
	CreatedAt             string                 `json:"created_at"`
	FullText              string                 `json:"full_text"`
	IsQuoteStatus         bool                   `json:"is_quote_status"`
	ReplyCount            int64                  `json:"reply_count"`
	RetweetCount          int64                  `json:"retweet_count"`
	FavoriteCount         int64                  `json:"favorite_count"`
	QuoteCount            int64                  `json:"quote_count"`
	Retweeted             bool                   `json:"retweeted"`
	Favorited             bool                   `json:"favorited"`
	RetweetedStatusResult *RetweetedStatusResult `json:"retweeted_status_result"`
}

type TweetResult struct {
	Legacy             *Legacy             `json:"legacy"`
	QuotedStatusResult *QuotedStatusResult `json:"quoted_status_result"`
	Views              *Views              `json:"views"`
	RestID             string              `json:"rest_id"`
	TypeName           string              `json:"__typename"`
}

type TweetResults struct {
	Result *TweetResult `json:"result"`
}

type ItemContent struct {
	ItemType         string        `json:"itemType"`
	TweetDisplayType string        `json:"tweetDisplayType"`
	TweetResults     *TweetResults `json:"tweet_results"`
}

type EntryContent struct {
	EntryType   string              `json:"entryType"`
	DisplayType string              `json:"displayType"`
	EntryID     string              `json:"entryId"`
	ItemContent *ItemContent        `json:"itemContent"`
	Items       []*CommentItemEntry `json:"items"`
	Value       string              `json:"value"`
}

type Entry struct {
	EntryID string        `json:"entryId"`
	Content *EntryContent `json:"content"`
}

type CommentLegacy struct {
	FullText  string `json:"full_text"`
	UserIDStr string `json:"user_id_str"`
}

type CommentResult struct {
	Legacy *CommentLegacy `json:"legacy"`
}

type CommentTweetResults struct {
	Result CommentResult `json:"result"`
}

type CommentItemContent struct {
	TweetResults CommentTweetResults `json:"tweet_results"`
}

type CommentItem struct {
	Content CommentItemContent `json:"itemContent"`
}

type CommentItemEntry struct {
	EntryID string      `json:"entryId"`
	Item    CommentItem `json:"item"`
}

type Instruction struct {
	Type    string   `json:"type"`
	Entries []*Entry `json:"entries"`
}

type Timeline struct {
	Instructions []*Instruction `json:"instructions"`
}

type TimelineV2 struct {
	Timeline *Timeline `json:"timeline"`
}

type Result struct {
	TimelineV2 *TimelineV2 `json:"timeline_v2"`
	TypeName   string      `json:"__typename"`
}

type User struct {
	Result *Result `json:"result"`
}

type TimeLineData struct {
	User *User `json:"user"`
}

type TimeLineResponse struct {
	Data *TimeLineData `json:"data"`
}

type ThreadedConversationWithInjectionsV2 struct {
	Instructions []*Instruction `json:"instructions"`
}

type TweetDetailData struct {
	Data *ThreadedConversationWithInjectionsV2 `json:"threaded_conversation_with_injections_v2"`
}

type TweetDetailResponse struct {
	Data *TweetDetailData `json:"data"`
}

type RateLimit struct {
	XRateLimitLimit     int64 `json:"x-rate-limit-limit"`
	XRateLimitRemaining int64 `json:"x-rate-limit-remaining"`
	XRateLimitReset     int64 `json:"x-rate-limit-reset"`
	XResponseTime       int64 `json:"x-response-time"`
}

type UserTweetsResponse struct {
	StatusCode int               `json:"status_code"`
	RateLimit  *RateLimit        `json:"rate_limit"`
	Data       *TimeLineResponse `json:"data"`
}

type TweetResponse struct {
	StatusCode int                  `json:"status_code"`
	RateLimit  *RateLimit           `json:"rate_limit"`
	Data       *TweetDetailResponse `json:"data"`
}

type FollowListInstructions struct {
	Type      string             `json:"type"`
	Direction string             `json:"direction,omitempty"`
	Entries   []*FollowListEntry `json:"entries,omitempty"`
}

type FollowListEntry struct {
	Content *FollowListContent `json:"content"`
}

type FollowListContent struct {
	ItemContent *FollowListItemContent `json:"itemContent"`
}

type FollowListItemContent struct {
	ItemType    string                 `json:"itemType"`
	Typename    string                 `json:"__typename"`
	UserResults *FollowListUserResults `json:"user_results"`
}

type FollowListUserResults struct {
	Result *FollowListResult `json:"result"`
}

type FollowListResult struct {
	RestID string `json:"rest_id"`
	Legacy struct {
		Following  bool   `json:"following"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	} `json:"legacy"`
}

type FollowingSubTimeline struct {
	Instructions []*FollowListInstructions `json:"instructions"`
}
type FollowListTimelineTimeline struct {
	Timeline *FollowingSubTimeline `json:"timeline"`
}

type FollowListTimeline struct {
	Timeline *FollowListTimelineTimeline `json:"timeline"`
}

type FollowListUser struct {
	Result *FollowListTimeline `json:"result"`
}

type FollowListUserData struct {
	User *FollowListUser `json:"user"`
}

type FollowListResp struct {
	StatusCode int                 `json:"status_code"`
	RateLimit  *RateLimit          `json:"rate_limit"`
	Data       *FollowListUserData `json:"data"`
}

type FollowingResp struct {
	StatusCode int             `json:"status_code"`
	RateLimit  *RateLimit      `json:"rate_limit"`
	Data       *FollowListResp `json:"data"`
}
type UserByScreenNameData struct {
	Data UserByScreenNameDataUser `json:"data"`
}

type UserByScreenNameDataUser struct {
	User UserByScreenNameResult `json:"user"`
}

type UserByScreenNameResult struct {
	Result UserResult `json:"result"`
}

type UserResult struct {
	TypeName string     `json:"__typename"`
	RestID   string     `json:"rest_id"`
	Legacy   UserLegacy `json:"legacy"`
}

type UserLegacy struct {
	Following bool `json:"following"`
}
