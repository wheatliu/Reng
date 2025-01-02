package twitter

import "encoding/json"

type twitterAPIParams interface {
	MarshalJsonWithDefaultValue() ([]byte, error)
}

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

type FollowingReqVariables struct {
	UserID                 string `json:"userId"`
	Count                  int    `json:"count"`
	IncludePromotedContent bool   `json:"includePromotedContent"`
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

func (fv *FollowingReqVariables) MarshalJsonWithDefaultValue() ([]byte, error) {
	fv.Count = 20
	fv.IncludePromotedContent = true
	return json.Marshal(fv)
}

func (ft *FieldToggles) MarshalJsonWithDefaultValue() ([]byte, error) {
	ft.WithArticleRichContentState = false
	return json.Marshal(ft)
}
