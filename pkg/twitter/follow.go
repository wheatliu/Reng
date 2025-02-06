package twitter

import (
	"encoding/json"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cast"
)

type FollowingReqVariables struct {
	UserID                 string `json:"userId"`
	Count                  int    `json:"count"`
	IncludePromotedContent bool   `json:"includePromotedContent"`
}

type FollowingReqFeatures struct {
	RwebTipjarConsumptionEnabled                                   bool `json:"rweb_tipjar_consumption_enabled"`
	ResponsiveWebGraphqlExcludeDirectiveEnabled                    bool `json:"responsive_web_graphql_exclude_directive_enabled"`
	VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
	CreatorSubscriptionsTweetPreviewApiEnabled                     bool `json:"creator_subscriptions_tweet_preview_api_enabled"`
	ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
	ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled      bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
	CommunitiesWebEnableTweetCommunityResultsFetch                 bool `json:"communities_web_enable_tweet_community_results_fetch"`
	C9STweetAnatomyModeratorBadgeEnabled                           bool `json:"c9s_tweet_anatomy_moderator_badge_enabled"`
	ArticlesPreviewEnabled                                         bool `json:"articles_preview_enabled"`
	ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
	GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
	ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
	LongformNotetweetsConsumptionEnabled                           bool `json:"longform_notetweets_consumption_enabled"`
	ResponsiveWebTwitterArticleTweetConsumptionEnabled             bool `json:"responsive_web_twitter_article_tweet_consumption_enabled"`
	TweetAwardsWebTippingEnabled                                   bool `json:"tweet_awards_web_tipping_enabled"`
	CreatorSubscriptionsQuoteTweetPreviewEnabled                   bool `json:"creator_subscriptions_quote_tweet_preview_enabled"`
	FreedomOfSpeechNotReachFetchEnabled                            bool `json:"freedom_of_speech_not_reach_fetch_enabled"`
	StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
	TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
	RwebVideoTimestampsEnabled                                     bool `json:"rweb_video_timestamps_enabled"`
	LongformNotetweetsRichTextReadEnabled                          bool `json:"longform_notetweets_rich_text_read_enabled"`
	LongformNotetweetsInlineMediaEnabled                           bool `json:"longform_notetweets_inline_media_enabled"`
	ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
}

func (f *FollowingReqFeatures) MarshalJsonWithDefaultValue() ([]byte, error) {
	f.RwebTipjarConsumptionEnabled = true
	f.ResponsiveWebGraphqlExcludeDirectiveEnabled = true
	f.VerifiedPhoneLabelEnabled = false
	f.CreatorSubscriptionsTweetPreviewApiEnabled = true
	f.ResponsiveWebGraphqlTimelineNavigationEnabled = true
	f.ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled = false
	f.CommunitiesWebEnableTweetCommunityResultsFetch = true
	f.C9STweetAnatomyModeratorBadgeEnabled = true
	f.ArticlesPreviewEnabled = true
	f.ResponsiveWebEditTweetApiEnabled = true
	f.GraphqlIsTranslatableRwebTweetIsTranslatableEnabled = true
	f.ViewCountsEverywhereApiEnabled = true
	f.LongformNotetweetsConsumptionEnabled = true
	f.ResponsiveWebTwitterArticleTweetConsumptionEnabled = true
	f.TweetAwardsWebTippingEnabled = false
	f.CreatorSubscriptionsQuoteTweetPreviewEnabled = false
	f.FreedomOfSpeechNotReachFetchEnabled = true
	f.StandardizedNudgesMisinfo = true
	f.TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled = true
	f.RwebVideoTimestampsEnabled = true
	f.LongformNotetweetsRichTextReadEnabled = true
	f.LongformNotetweetsInlineMediaEnabled = true
	f.ResponsiveWebEnhanceCardsEnabled = false
	return json.Marshal(f)
}

func (fv *FollowingReqVariables) MarshalJsonWithDefaultValue() ([]byte, error) {
	fv.Count = 20
	fv.IncludePromotedContent = true
	return json.Marshal(fv)
}

func (c *TwitterClient) GetUserFollowing(query *FollowingReqVariables) (resp *FollowingResp, err error) {
	resp = &FollowingResp{}
	features := &FollowingReqFeatures{}
	reqURL, err := c.generateReqURL(UserFollowingURL, query, features, nil)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient generateReqURL Error")
	}

	followListResp := &FollowListResp{}
	req, err := c.httpCli.R().SetResult(followListResp).Get(reqURL)
	if err != nil {
		fmt.Printf("Fetched User Following Failed, TwitterID: %s, Response Status: %d, resp: %s",
			query.UserID, req.StatusCode(), req.String())
		return nil, errors.Wrap(err, "TwitterClient Request GetUserTweets error")
	}
	headers := req.Header()
	fmt.Printf(
		"Fetched User Following Successful, TwitterID: %s, Response Status: %d, Header: limit: %s, reset: %s, time: %s, remaining: %s",
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
	resp.Data = followListResp
	return resp, nil
}
