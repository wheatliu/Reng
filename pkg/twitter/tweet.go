package twitter

import (
	"encoding/json"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cast"
)

type CommonReqVariables struct {
	IncludePromotedContent                 bool `json:"includePromotedContent"`
	WithBirdwatchNotes                     bool `json:"withBirdwatchNotes"`
	WithQuickPromoteEligibilityTweetFields bool `json:"withQuickPromoteEligibilityTweetFields"`
	WithCommunity                          bool `json:"withCommunity"`
	WithVoice                              bool `json:"withVoice"`
	WithV2Timeline                         bool `json:"withV2Timeline"`
}

type TimeLineReqVariables struct {
	UserID string `json:"userId"`
	Count  int    `json:"count"`
	Cursor string `json:"cursor"`
	CommonReqVariables
}

type TweetDetailReqVariables struct {
	FocalTweetId      string `json:"focalTweetId"`
	WithRuxInjections bool   `json:"with_rux_injections"`
	CommonReqVariables
}

type FieldToggles struct {
	WithArticleRichContentState bool `json:"withArticleRichContentState"`
}

type CommonReqFeatures struct {
	ResponsiveWebGraphqlExcludeDirectiveEnabled                    bool `json:"responsive_web_graphql_exclude_directive_enabled"`
	VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
	CreatorSubscriptionsTweetPreviewApiEnabled                     bool `json:"creator_subscriptions_tweet_preview_api_enabled"`
	ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
	ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled      bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
	C9STweetAnatomyModeratorBadgeEnabled                           bool `json:"c9s_tweet_anatomy_moderator_badge_enabled"`
	TweetypieUnmentionOptimizationEnabled                          bool `json:"tweetypie_unmention_optimization_enabled"`
	ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
	GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
	ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
	LongformNotetweetsConsumptionEnabled                           bool `json:"longform_notetweets_consumption_enabled"`
	ResponsiveWebTwitterArticleTweetConsumptionEnabled             bool `json:"responsive_web_twitter_article_tweet_consumption_enabled"`
	TweetAwardsWebTippingEnabled                                   bool `json:"tweet_awards_web_tipping_enabled"`
	FreedomOfSpeechNotReachFetchEnabled                            bool `json:"freedom_of_speech_not_reach_fetch_enabled"`
	StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
	TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
	RwebVideoTimestampsEnabled                                     bool `json:"rweb_video_timestamps_enabled"`
	LongformNotetweetsRichTextReadEnabled                          bool `json:"longform_notetweets_rich_text_read_enabled"`
	LongformNotetweetsInlineMediaEnabled                           bool `json:"longform_notetweets_inline_media_enabled"`
	ResponsiveWebMediaDownloadVideoEnabled                         bool `json:"responsive_web_media_download_video_enabled"`
	ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
}

func (f *CommonReqFeatures) MarshalJsonWithDefaultValue() ([]byte, error) {
	f.ResponsiveWebGraphqlExcludeDirectiveEnabled = true
	f.CreatorSubscriptionsTweetPreviewApiEnabled = true
	f.ResponsiveWebGraphqlTimelineNavigationEnabled = true
	f.C9STweetAnatomyModeratorBadgeEnabled = true
	f.TweetypieUnmentionOptimizationEnabled = true
	f.ResponsiveWebEditTweetApiEnabled = true
	f.GraphqlIsTranslatableRwebTweetIsTranslatableEnabled = true
	f.ViewCountsEverywhereApiEnabled = true
	f.LongformNotetweetsConsumptionEnabled = true
	f.ResponsiveWebTwitterArticleTweetConsumptionEnabled = true
	f.FreedomOfSpeechNotReachFetchEnabled = true
	f.StandardizedNudgesMisinfo = true
	f.TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled = true
	f.RwebVideoTimestampsEnabled = true
	f.LongformNotetweetsRichTextReadEnabled = true
	f.LongformNotetweetsInlineMediaEnabled = true
	return json.Marshal(f)
}

func (ft *FieldToggles) MarshalJsonWithDefaultValue() ([]byte, error) {
	ft.WithArticleRichContentState = false
	return json.Marshal(ft)
}

func (tlv *TimeLineReqVariables) MarshalJsonWithDefaultValue() ([]byte, error) {
	tlv.Count = 20
	tlv.IncludePromotedContent = true
	tlv.WithCommunity = true
	tlv.WithQuickPromoteEligibilityTweetFields = true
	tlv.WithBirdwatchNotes = true
	tlv.WithVoice = true
	tlv.WithV2Timeline = true
	return json.Marshal(tlv)
}

func (tdv *TweetDetailReqVariables) MarshalJsonWithDefaultValue() ([]byte, error) {
	tdv.WithRuxInjections = false
	tdv.IncludePromotedContent = true
	tdv.WithCommunity = true
	tdv.WithQuickPromoteEligibilityTweetFields = true
	tdv.WithBirdwatchNotes = true
	tdv.WithVoice = true
	tdv.WithV2Timeline = true
	return json.Marshal(tdv)
}

func (c *TwitterClient) GetUserTweets(query *TimeLineReqVariables) (resp *UserTweetsResponse, err error) {
	resp = &UserTweetsResponse{}
	features := &CommonReqFeatures{}
	reqURL, err := c.generateReqURL(UserTweetsURL, query, features, nil)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient generateReqURL Error")
	}

	userTweetsResp := &TimeLineResponse{}
	req, err := c.httpCli.R().SetResult(userTweetsResp).Get(reqURL)
	if err != nil {
		fmt.Printf("Fetched User Tweets Failed, TwitterID: %s, Response Status: %d, resp: %s",
			query.UserID, req.StatusCode(), req.String())
		return nil, errors.Wrap(err, "TwitterClient Request GetUserTweets error")
	}
	// fmt.Println("req status: %d, resp: %s", req.StatusCode(), req.String())

	headers := req.Header()

	fmt.Printf(
		"Fetched User Tweets Successful, TwitterID: %s, Response Status: %d, Header: limit: %s, reset: %s, time: %s, remaining: %s",
		query.UserID, req.StatusCode(),
		headers.Get("X-Rate-Limit-Limit"), headers.Get("X-Rate-Limit-Reset"),
		headers.Get("X-Response-Time"), headers.Get("X-Rate-Limit-Remaining"),
	)

	resp.RateLimit = &RateLimit{
		XRateLimitLimit:     cast.ToInt64(headers.Get("X-Rate-Limit-Limit")),
		XRateLimitReset:     cast.ToInt64(headers.Get("X-Rate-Limit-Reset")),
		XResponseTime:       cast.ToInt64(headers.Get("X-Response-Time")),
		XRateLimitRemaining: cast.ToInt64(headers.Get("X-Rate-Limit-Remaining")),
	}
	resp.StatusCode = req.StatusCode()
	resp.Data = userTweetsResp
	return resp, nil
}

func (c *TwitterClient) GetUserTweetDetail(tweetID string) (resp *TweetResponse, err error) {
	resp = &TweetResponse{}
	variables := TweetDetailReqVariables{
		FocalTweetId: tweetID,
	}
	features := &CommonReqFeatures{}
	fieldToggles := &FieldToggles{}
	reqURL, err := c.generateReqURL(TweetDetailURL, &variables, features, fieldToggles)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient generateReqURL error")
	}

	tweetDetailResp := &TweetDetailResponse{}
	req, err := c.httpCli.R().SetResult(&tweetDetailResp).Get(reqURL)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient Request GetUserTweets error")
	}
	headers := req.Header()
	fmt.Printf(
		"Header: limit: %s, reset: %s, time: %s, remaining: %s",
		headers.Get("X-Rate-Limit-Limit"), headers.Get("X-Rate-Limit-Reset"),
		headers.Get("X-Response-Time"), headers.Get("X-Rate-Limit-Remaining"),
	)

	resp.RateLimit = &RateLimit{
		XRateLimitLimit:     cast.ToInt64(headers.Get("x-rate-limit-limit")),
		XRateLimitReset:     cast.ToInt64(headers.Get("x-rate-limit-reset")),
		XResponseTime:       cast.ToInt64(headers.Get("x-response-time")),
		XRateLimitRemaining: cast.ToInt64(headers.Get("x-rate-limit-remaining")),
	}
	resp.StatusCode = req.StatusCode()
	resp.Data = tweetDetailResp
	return resp, nil
}
