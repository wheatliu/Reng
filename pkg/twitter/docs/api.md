### Twitter Graphql API

#### UserProfile

##### Request

```javascript
fetch(
  "https://x.com/i/api/graphql/QGIw94L0abhuohrr76cSbw/UserByScreenName?variables=%7B%22screen_name%22%3A%22visegrad24%22%7D&features=%7B%22hidden_profile_subscriptions_enabled%22%3Atrue%2C%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22subscriptions_verification_info_is_identity_verified_enabled%22%3Atrue%2C%22subscriptions_verification_info_verified_since_enabled%22%3Atrue%2C%22highlights_tweets_tab_ui_enabled%22%3Atrue%2C%22responsive_web_twitter_article_notes_tab_enabled%22%3Atrue%2C%22subscriptions_feature_can_gift_premium%22%3Atrue%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%7D&fieldToggles=%7B%22withAuxiliaryUserLabels%22%3Afalse%7D",
  {
    headers: {
      accept: "*/*",
      "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
      authorization:
        "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
      "content-type": "application/json",
      priority: "u=1, i",
      "sec-ch-ua":
        '"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"',
      "sec-ch-ua-mobile": "?0",
      "sec-ch-ua-platform": '"macOS"',
      "sec-fetch-dest": "empty",
      "sec-fetch-mode": "cors",
      "sec-fetch-site": "same-origin",
      "x-client-transaction-id":
        "hyTVV0x0lEOTpFzmluIF94SALBgnVDdFslWVtZcJy8dxpxmDh07/NWuMbfCmcBMhepl+ooQDCWnDYveYeXzyxndsEn8chA",
      "x-csrf-token":
        "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
      "x-twitter-active-user": "yes",
      "x-twitter-auth-type": "OAuth2Session",
      "x-twitter-client-language": "en",
      Referer: "https://x.com/visegrad24",
      "Referrer-Policy": "strict-origin-when-cross-origin",
    },
    body: null,
    method: "GET",
  }
);
```

###### Payload

```javascript

variables: {"screen_name":"visegrad24"}
features: {"hidden_profile_subscriptions_enabled":true,"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"subscriptions_verification_info_is_identity_verified_enabled":true,"subscriptions_verification_info_verified_since_enabled":true,"highlights_tweets_tab_ui_enabled":true,"responsive_web_twitter_article_notes_tab_enabled":true,"subscriptions_feature_can_gift_premium":true,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"responsive_web_graphql_timeline_navigation_enabled":true}
fieldToggles: {"withAuxiliaryUserLabels":false}
```

###### Response

```javascript
{
    "data": {
        "user": {
            "result": {
                "__typename": "User",
                "id": "VXNlcjoxMjIyNzczMzAyNDQxMTQ4NDE2",
                "rest_id": "1222773302441148416",
                "affiliates_highlighted_label": {},
                "has_graduated_access": true,
                "is_blue_verified": true,
                "profile_image_shape": "Circle",
                "legacy": {
                    "following": true,
                    "can_dm": true,
                    "can_media_tag": false,
                    "created_at": "Thu Jan 30 06:48:04 +0000 2020",
                    "default_profile": true,
                    "default_profile_image": false,
                    "description": "Aggregating and curating news, politics and current affairs. #BTC",
                    "entities": {
                        "description": {
                            "urls": []
                        },
                        "url": {
                            "urls": [
                                {
                                    "display_url": "visegrad24.com",
                                    "expanded_url": "http://www.visegrad24.com",
                                    "url": "https://t.co/qpd6ZYX46V",
                                    "indices": [
                                        0,
                                        23
                                    ]
                                }
                            ]
                        }
                    },
                    "fast_followers_count": 0,
                    "favourites_count": 66015,
                    "followers_count": 1253647,
                    "friends_count": 1882,
                    "has_custom_timelines": true,
                    "is_translator": false,
                    "listed_count": 9347,
                    "location": "Visegrad",
                    "media_count": 37641,
                    "name": "Visegrád 24",
                    "normal_followers_count": 1253647,
                    "pinned_tweet_ids_str": [
                        "1874129693806649859"
                    ],
                    "possibly_sensitive": false,
                    "profile_banner_url": "https://pbs.twimg.com/profile_banners/1222773302441148416/1715766741",
                    "profile_image_url_https": "https://pbs.twimg.com/profile_images/1798263200648994816/qFYydb3B_normal.jpg",
                    "profile_interstitial_type": "",
                    "screen_name": "visegrad24",
                    "statuses_count": 60812,
                    "translator_type": "none",
                    "url": "https://t.co/qpd6ZYX46V",
                    "verified": false,
                    "want_retweets": true,
                    "withheld_in_countries": []
                },
                "professional": {
                    "rest_id": "1679839390804131841",
                    "professional_type": "Creator",
                    "category": [
                        {
                            "id": 580,
                            "name": "Media & News Company",
                            "icon_name": "IconBriefcaseStroke"
                        }
                    ]
                },
                "tipjar_settings": {
                    "is_enabled": true,
                    "bitcoin_handle": "35DXndp5iFAecrhWxJh54xmMKjknoHuJsD",
                    "ethereum_handle": "0xBcA55830bd9EDAe7Ca5ACCf8369BDD8f5364f2dc"
                },
                "legacy_extended_profile": {},
                "is_profile_translatable": false,
                "has_hidden_subscriptions_on_profile": false,
                "verification_info": {
                    "is_identity_verified": false,
                    "reason": {
                        "description": {
                            "text": "This account is verified. Learn more",
                            "entities": [
                                {
                                    "from_index": 26,
                                    "to_index": 36,
                                    "ref": {
                                        "url": "https://help.twitter.com/managing-your-account/about-twitter-verified-accounts",
                                        "url_type": "ExternalUrl"
                                    }
                                }
                            ]
                        },
                        "verified_since_msec": "1678958739464"
                    }
                },
                "highlights_info": {
                    "can_highlight_tweets": true,
                    "highlighted_tweets": "82"
                },
                "user_seed_tweet_count": 0,
                "premium_gifting_eligible": false,
                "business_account": {},
                "creator_subscriptions_count": 0
            }
        }
    }
}
```


##### Follower

###### Request

```javascript
fetch("https://x.com/i/api/graphql/r4fuEJKOqqzaYcvJU5ZWVA/Followers?variables=%7B%22userId%22%3A%221222773302441148416%22%2C%22count%22%3A20%2C%22includePromotedContent%22%3Afalse%7D&features=%7B%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22premium_content_api_read_enabled%22%3Afalse%2C%22communities_web_enable_tweet_community_results_fetch%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_button_fetch_trends_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_post_followups_enabled%22%3Afalse%2C%22responsive_web_grok_share_attachment_enabled%22%3Atrue%2C%22articles_preview_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22creator_subscriptions_quote_tweet_preview_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "E7BBw9jgANcHMMhyAnaRYxAUuIyzwKPRJsEBIQOdX1PlM40XE9prof8Y+WQy5Ie17tbpNhAbd8JwIFiGHGER3OCBWP8iEA",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/visegrad24/followers",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

###### Payload

```javascript
variables: {"userId":"1222773302441148416","count":20,"includePromotedContent":false}
features: {"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"premium_content_api_read_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"responsive_web_grok_analyze_button_fetch_trends_enabled":true,"responsive_web_grok_analyze_post_followups_enabled":false,"responsive_web_grok_share_attachment_enabled":true,"articles_preview_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"rweb_video_timestamps_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"responsive_web_enhance_cards_enabled":false}
```

###### Response
``` javascript
{
    "data": {
        "user": {
            "result": {
                "__typename": "User",
                "timeline": {
                    "timeline": {
                        "instructions": [
                            {
                                "type": "TimelineClearCache"
                            },
                            {
                                "type": "TimelineTerminateTimeline",
                                "direction": "Top"
                            },
                            {
                                "type": "TimelineTerminateTimeline",
                                "direction": "Bottom"
                            },
                            {
                                "type": "TimelineAddEntries",
                                "entries": [
                                    {
                                        "entryId": "user-1537130875296563200",
                                        "sortIndex": "1874476833932574720",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTM3MTMwODc1Mjk2NTYzMjAw",
                                                        "rest_id": "1537130875296563200",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jun 15 17:54:58 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "American in Ukraine",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 50782,
                                                            "followers_count": 130935,
                                                            "friends_count": 1849,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 2146,
                                                            "location": "Kyiv",
                                                            "media_count": 13032,
                                                            "name": "Jay in Kyiv",
                                                            "normal_followers_count": 130935,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1803814183118557184/V4jlh9YN_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "JayinKyiv",
                                                            "statuses_count": 39964,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {},
                                                        "super_follow_eligible": true
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-42730080",
                                        "sortIndex": "1874476833932574719",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0MjczMDA4MA==",
                                                        "rest_id": "42730080",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue May 26 21:42:44 +0000 2009",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Slovenian political dissident. My values are limited government, free markets, private initiative, just peace, and right to bear arms. \uD83C\uDDF8\uD83C\uDDEE \uD83C\uDDEC\uD83C\uDDE7 \uD83C\uDDFA\uD83C\uDDE6",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "libertarec.blogspot.com",
                                                                            "expanded_url": "http://libertarec.blogspot.com",
                                                                            "url": "https://t.co/Dy9pmgFMpK",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 31080,
                                                            "followers_count": 25882,
                                                            "friends_count": 693,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 95,
                                                            "location": "Planet Earth",
                                                            "media_count": 6270,
                                                            "name": "Libertarec",
                                                            "normal_followers_count": 25882,
                                                            "pinned_tweet_ids_str": [
                                                                "1874362246333334004"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/42730080/1646894813",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/871457550486712320/mdZtLe5l_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Libertarec",
                                                            "statuses_count": 104098,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/Dy9pmgFMpK",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-79179041",
                                        "sortIndex": "1874476833932574718",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo3OTE3OTA0MQ==",
                                                        "rest_id": "79179041",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Fri Oct 02 13:22:21 +0000 2009",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Profesor, računalnikar, gradbenik, publicist.\nSkrajni resničar.\n\nFor opinions in \uD83C\uDDEC\uD83C\uDDE7 see @zigaTurkEU.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "zturk.com",
                                                                            "expanded_url": "http://www.zturk.com/",
                                                                            "url": "https://t.co/PEHAWN996u",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 29457,
                                                            "followers_count": 40059,
                                                            "friends_count": 525,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 185,
                                                            "location": "Republic of Slovenia",
                                                            "media_count": 4680,
                                                            "name": "Žiga Turk",
                                                            "normal_followers_count": 40059,
                                                            "pinned_tweet_ids_str": [
                                                                "1874211502787526764"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/79179041/1723108354",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1132867062060212227/qU3MibtF_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "ZigaTurk",
                                                            "statuses_count": 111713,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/PEHAWN996u",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1521821251643318274",
                                                            "professional_type": "Creator",
                                                            "category": [
                                                                {
                                                                    "id": 1083,
                                                                    "name": "Intellectual",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": true
                                                        }
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-60696848",
                                        "sortIndex": "1874476833932574717",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo2MDY5Njg0OA==",
                                                        "rest_id": "60696848",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Mon Jul 27 20:12:55 +0000 2009",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Perussuomalaisten Espoon kaupunginvaltuutettu, oikeistohippi ja homosatanisti. Kirjoitan politiikasta, hörhöilystä ja black metal -musiikista.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "keronen.wordpress.com/politiikka",
                                                                            "expanded_url": "http://keronen.wordpress.com/politiikka",
                                                                            "url": "https://t.co/wn3phAcfK8",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 89728,
                                                            "followers_count": 13382,
                                                            "friends_count": 549,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 43,
                                                            "location": "Finland",
                                                            "media_count": 1273,
                                                            "name": "Jiri Keronen",
                                                            "normal_followers_count": 13382,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/60696848/1665315814",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1865392217721880576/MU8yO5F7_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "keronen",
                                                            "statuses_count": 55094,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/wn3phAcfK8",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-2417209184",
                                        "sortIndex": "1874476833932574716",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyNDE3MjA5MTg0",
                                                        "rest_id": "2417209184",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Mar 29 09:32:33 +0000 2014",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Twiittailen enimmäkseen turvallisuuspolitiikasta. Maanpuolustus lähellä sydäntä. Conservative. Geopolitics, defence. Finnish and global military.\nTsuhna-Pirate.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 188971,
                                                            "followers_count": 21811,
                                                            "friends_count": 2596,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 40,
                                                            "location": "Helsinki",
                                                            "media_count": 1528,
                                                            "name": "ProFin\uD83C\uDDEB\uD83C\uDDEE\uD83C\uDDEA\uD83C\uDDFA\uD83C\uDDEA\uD83C\uDDEA\uD83C\uDDF8\uD83C\uDDEA\uD83C\uDDFA\uD83C\uDDE6",
                                                            "normal_followers_count": 21811,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/2417209184/1680603810",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/918365444012150784/tHloWWJU_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "RockOla69",
                                                            "statuses_count": 45552,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-887259665641418752",
                                        "sortIndex": "1874476833932574715",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo4ODcyNTk2NjU2NDE0MTg3NTI=",
                                                        "rest_id": "887259665641418752",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Jul 18 10:36:03 +0000 2017",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Professor at University of Ljubljana & Former Minister",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 22559,
                                                            "followers_count": 33582,
                                                            "friends_count": 353,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 38,
                                                            "location": "",
                                                            "media_count": 208,
                                                            "name": "Matej Lahovnik",
                                                            "normal_followers_count": 33582,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1451661507192836102/93pqKboD_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "LahovnikMatej",
                                                            "statuses_count": 12701,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1415754894997434373",
                                        "sortIndex": "1874476833932574714",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNDE1NzU0ODk0OTk3NDM0Mzcz",
                                                        "rest_id": "1415754894997434373",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Jul 15 19:28:00 +0000 2021",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "The heart of Europe, a bastion of Christianity, best of East and West.\n+-+-+-+⚔️\uD83D\uDC51\uD83D\uDEE1️+-+-+-+\nLegatum et hereditatum Piastorum+Iagiellonorum+Vasorum+Ioannis III",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 12200,
                                                            "followers_count": 17955,
                                                            "friends_count": 2346,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 67,
                                                            "location": "VarsoviaLeopolisVilnaCracovia",
                                                            "media_count": 1995,
                                                            "name": "Europe of the Jagiellonians",
                                                            "normal_followers_count": 17955,
                                                            "pinned_tweet_ids_str": [
                                                                "1725166837929775238"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1415754894997434373/1723908121",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1415758587108990977/lgmd3og1_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "EJagiellonians",
                                                            "statuses_count": 2396,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-2397890569",
                                        "sortIndex": "1874476833932574713",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyMzk3ODkwNTY5",
                                                        "rest_id": "2397890569",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Wed Mar 19 13:41:43 +0000 2014",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "https://t.co/nnFCOegbHh",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "buymeacoffee.com/alexandruc4",
                                                                            "expanded_url": "https://www.buymeacoffee.com/alexandruc4",
                                                                            "url": "https://t.co/nnFCOegbHh",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 10221,
                                                            "followers_count": 35243,
                                                            "friends_count": 11069,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 567,
                                                            "location": "",
                                                            "media_count": 17809,
                                                            "name": "AlexandruC4",
                                                            "normal_followers_count": 35243,
                                                            "pinned_tweet_ids_str": [
                                                                "1593007519865376768"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/678655355405119488/XWMUkrBV_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "AlexandruC4",
                                                            "statuses_count": 74726,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": true
                                                        }
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1300810545872896007",
                                        "sortIndex": "1874476833932574712",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMzAwODEwNTQ1ODcyODk2MDA3",
                                                        "rest_id": "1300810545872896007",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Sep 01 14:59:44 +0000 2020",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Conservative content from Poland. Born in France to Polish parents, now back to Poland. Support #V4 #3SI #IVRP  \uD83C\uDDF5\uD83C\uDDF1",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 14348,
                                                            "followers_count": 23368,
                                                            "friends_count": 1145,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 78,
                                                            "location": "Kraków, Poland",
                                                            "media_count": 3943,
                                                            "name": "Expat in Poland \uD83C\uDDF5\uD83C\uDDF1",
                                                            "normal_followers_count": 23368,
                                                            "pinned_tweet_ids_str": [
                                                                "1712562486996787612"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1300810545872896007/1604282059",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1305260313399484416/-hFkBC9A_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "BasedPoland2",
                                                            "statuses_count": 11847,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-4229156917",
                                        "sortIndex": "1874476833932574711",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0MjI5MTU2OTE3",
                                                        "rest_id": "4229156917",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Nov 14 00:51:38 +0000 2015",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "\uD83C\uDDF5\uD83C\uDDF1\uD83C\uDDEB\uD83C\uDDF7 Geoeconomist. My work: https://t.co/GTOmhhhM1W, Warsaw Express (Substack), https://t.co/jhlrjG0ZQg, & https://t.co/Sibn9dCkpd.\n⚠️☠️ Polish imperialist.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "ExcaliburInsight.com",
                                                                            "expanded_url": "http://ExcaliburInsight.com",
                                                                            "url": "https://t.co/GTOmhhhM1W",
                                                                            "indices": [
                                                                                28,
                                                                                51
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "CourrierDeVarsovie.com",
                                                                            "expanded_url": "http://CourrierDeVarsovie.com",
                                                                            "url": "https://t.co/jhlrjG0ZQg",
                                                                            "indices": [
                                                                                80,
                                                                                103
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "MaterialCivilization.com",
                                                                            "expanded_url": "http://MaterialCivilization.com",
                                                                            "url": "https://t.co/Sibn9dCkpd",
                                                                            "indices": [
                                                                                107,
                                                                                130
                                                                            ]
                                                                        }
                                                                    ]
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "open.substack.com/pub/danielfoub…",
                                                                            "expanded_url": "https://open.substack.com/pub/danielfoubert?r=1z78jf&utm_medium=ios",
                                                                            "url": "https://t.co/0w7xY4V6Ve",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 56516,
                                                            "followers_count": 28959,
                                                            "friends_count": 3396,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 280,
                                                            "location": "Warsaw (Intermarium)",
                                                            "media_count": 7931,
                                                            "name": "Daniel Foubert",
                                                            "normal_followers_count": 28959,
                                                            "pinned_tweet_ids_str": [
                                                                "1789353592882675960"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/4229156917/1691326542",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1675572628612956160/AeqaNoxM_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "d_foubert",
                                                            "statuses_count": 47770,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/0w7xY4V6Ve",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1552608730222559232",
                                                            "professional_type": "Creator",
                                                            "category": [
                                                                {
                                                                    "id": 957,
                                                                    "name": "Author",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-18896484",
                                        "sortIndex": "1874476833932574710",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODg5NjQ4NA==",
                                                        "rest_id": "18896484",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Mon Jan 12 09:36:16 +0000 2009",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "RT ni ❤️ --- https://t.co/pzsJilY2FJ",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "amazon.com/Confidence-sex…",
                                                                            "expanded_url": "http://amazon.com/Confidence-sexy-Self-confident-body-language-ebook/dp/B07R7HJ39C/",
                                                                            "url": "https://t.co/pzsJilY2FJ",
                                                                            "indices": [
                                                                                13,
                                                                                36
                                                                            ]
                                                                        }
                                                                    ]
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "edvardkadic.com",
                                                                            "expanded_url": "http://edvardkadic.com/",
                                                                            "url": "https://t.co/A3rlAbUZW3",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 81872,
                                                            "followers_count": 16361,
                                                            "friends_count": 4604,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 80,
                                                            "location": "",
                                                            "media_count": 8486,
                                                            "name": "Edvard Kadič",
                                                            "normal_followers_count": 16361,
                                                            "pinned_tweet_ids_str": [
                                                                "1873753305056461022"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/18896484/1722894692",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1522297815086555140/d2wrMlQo_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "EdvardKadic",
                                                            "statuses_count": 84661,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/A3rlAbUZW3",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1472969209076404228",
                                                            "professional_type": "Creator",
                                                            "category": [
                                                                {
                                                                    "id": 957,
                                                                    "name": "Author",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-2719115466",
                                        "sortIndex": "1874476833932574709",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyNzE5MTE1NDY2",
                                                        "rest_id": "2719115466",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Aug 09 11:21:06 +0000 2014",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Balkán, Afghánistán, a teď taky Ukrajina...\n\nHead  of PRT Logar, Afghanistan 2010/11\n\nCNN Prima News Foreign Desk Head",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 9078,
                                                            "followers_count": 23756,
                                                            "friends_count": 479,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 142,
                                                            "location": "",
                                                            "media_count": 2435,
                                                            "name": "Matyas Zrno",
                                                            "normal_followers_count": 23756,
                                                            "pinned_tweet_ids_str": [
                                                                "1825967998995239187"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/2719115466/1687812476",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1582065743436386308/W4BLQUy9_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "MatyasZrno",
                                                            "statuses_count": 10452,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-485079585",
                                        "sortIndex": "1874476833932574708",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0ODUwNzk1ODU=",
                                                        "rest_id": "485079585",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Mon Feb 06 20:19:48 +0000 2012",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Novinar: Pričevalci/Intervju -TVSLO\nZgodovinar: knjiga- Slovenski razkol, 8 ponatis\nHistorian Phd, Journalist, TV host,  director\nbook: The Slovenian razkol",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "jozemozina.si/o-avtorju/",
                                                                            "expanded_url": "https://jozemozina.si/o-avtorju/",
                                                                            "url": "https://t.co/nmO2YdaUt8",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 22761,
                                                            "followers_count": 18295,
                                                            "friends_count": 962,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 32,
                                                            "location": "..sin Primorske! ",
                                                            "media_count": 778,
                                                            "name": "Jože Možina",
                                                            "normal_followers_count": 18295,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/485079585/1412772895",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/2367957859/bakvljmqh7qg8nn98ot3_normal.jpeg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "JozeMozina",
                                                            "statuses_count": 2572,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/nmO2YdaUt8",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1460728296560287749",
                                                            "professional_type": "Creator",
                                                            "category": [
                                                                {
                                                                    "id": 955,
                                                                    "name": "Journalist",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-409170134",
                                        "sortIndex": "1874476833932574707",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0MDkxNzAxMzQ=",
                                                        "rest_id": "409170134",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Nov 10 11:14:35 +0000 2011",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Vesel različnih pogledov. Plačance brez sledilcev, ki le zmerjajo in žalijo, odpravim na hitro.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "spletnicasopis.eu",
                                                                            "expanded_url": "http://spletnicasopis.eu",
                                                                            "url": "https://t.co/jwF4aF6n6W",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 538,
                                                            "followers_count": 16288,
                                                            "friends_count": 1649,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 54,
                                                            "location": "Ljubljana",
                                                            "media_count": 2866,
                                                            "name": "peter jancic",
                                                            "normal_followers_count": 16288,
                                                            "pinned_tweet_ids_str": [
                                                                "1872557658043244715"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/409170134/1448229776",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/553156639134470146/gRMV_1F__normal.jpeg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "peterjancic",
                                                            "statuses_count": 42776,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/jwF4aF6n6W",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1598055263235440645",
                                                            "professional_type": "Creator",
                                                            "category": [
                                                                {
                                                                    "id": 580,
                                                                    "name": "Media & News Company",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-776093974188662784",
                                        "sortIndex": "1874476833932574706",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo3NzYwOTM5NzQxODg2NjI3ODQ=",
                                                        "rest_id": "776093974188662784",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Sep 14 16:23:17 +0000 2016",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Docteur ès lettres de l’Université Paris-Sorbonne.\nProf des Universités.\nMembre de l’Académie européenne des sciences et des arts (vice-doyen classe I)",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "en.wikipedia.org/wiki/Bo%C5%A1t…",
                                                                            "expanded_url": "https://en.wikipedia.org/wiki/Bo%C5%A1tjan_Marko_Turk",
                                                                            "url": "https://t.co/U6CdkovDCb",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 20731,
                                                            "followers_count": 13337,
                                                            "friends_count": 4288,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 24,
                                                            "location": "photo: Miha Tozon, nov. 24",
                                                            "media_count": 154,
                                                            "name": "Boštjan M. Turk",
                                                            "normal_followers_count": 13337,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1869081801332813824/P6Yqy_Mc_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "m_bostjan",
                                                            "statuses_count": 61576,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/U6CdkovDCb",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-135918040",
                                        "sortIndex": "1874476833932574705",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMzU5MTgwNDA=",
                                                        "rest_id": "135918040",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Apr 22 14:56:58 +0000 2010",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Ubi bene, ibi patria.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 142222,
                                                            "followers_count": 9569,
                                                            "friends_count": 715,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 26,
                                                            "location": "",
                                                            "media_count": 10894,
                                                            "name": "Peter Hrastelj",
                                                            "normal_followers_count": 9569,
                                                            "pinned_tweet_ids_str": [
                                                                "1861394977206505539"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/135918040/1697998907",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1804136623409446912/5qN-8xC4_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "hrastelj",
                                                            "statuses_count": 67857,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1473354685541146636",
                                                            "professional_type": "Creator",
                                                            "category": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-247247840",
                                        "sortIndex": "1874476833932574704",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyNDcyNDc4NDA=",
                                                        "rest_id": "247247840",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Feb 04 12:03:07 +0000 2011",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 8785,
                                                            "followers_count": 10544,
                                                            "friends_count": 1170,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 42,
                                                            "location": "",
                                                            "media_count": 238,
                                                            "name": "Uroš Urbanija",
                                                            "normal_followers_count": 10544,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/516985212933128192/IvtLzdx9_normal.jpeg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "UrosUrbanija",
                                                            "statuses_count": 42042,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-962475359668264960",
                                        "sortIndex": "1874476833932574703",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo5NjI0NzUzNTk2NjgyNjQ5NjA=",
                                                        "rest_id": "962475359668264960",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Feb 10 23:56:22 +0000 2018",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Dystocracy is cognitive dissonance to the clergy of the status quo and the lapdogs of the system who think they are revolutionaries.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "dystocracy.com",
                                                                            "expanded_url": "http://www.dystocracy.com",
                                                                            "url": "https://t.co/yxo7ew6Uu2",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 293714,
                                                            "followers_count": 7051,
                                                            "friends_count": 915,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 7,
                                                            "location": "Tampere, Finland",
                                                            "media_count": 11608,
                                                            "name": "Dystocracy",
                                                            "normal_followers_count": 7051,
                                                            "pinned_tweet_ids_str": [
                                                                "1743534972550406494"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/962475359668264960/1555480079",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/962477170168868864/bMPzcwzo_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Dystocracycrew",
                                                            "statuses_count": 147060,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/yxo7ew6Uu2",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-459532635",
                                        "sortIndex": "1874476833932574702",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0NTk1MzI2MzU=",
                                                        "rest_id": "459532635",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Mon Jan 09 19:38:58 +0000 2012",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 138792,
                                                            "followers_count": 13260,
                                                            "friends_count": 3110,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 40,
                                                            "location": "Slovenia",
                                                            "media_count": 9512,
                                                            "name": "vinko vasle novinar",
                                                            "normal_followers_count": 13260,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1694634699979653120/xZVRehjy_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "vinkovasle1",
                                                            "statuses_count": 325019,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-2965817333",
                                        "sortIndex": "1874476833932574701",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyOTY1ODE3MzMz",
                                                        "rest_id": "2965817333",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Wed Jan 07 11:01:17 +0000 2015",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Let's get rich together.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 19445,
                                                            "followers_count": 9779,
                                                            "friends_count": 805,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 35,
                                                            "location": "",
                                                            "media_count": 4869,
                                                            "name": "Mitja Irsic",
                                                            "normal_followers_count": 9779,
                                                            "pinned_tweet_ids_str": [
                                                                "923630423749615618"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/2965817333/1457629625",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/552782579254829056/oMVUs1VS_normal.jpeg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "MitjaIrsic",
                                                            "statuses_count": 45915,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874474007216467968",
                                        "sortIndex": "1874476833932574700",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDc0MDA3MjE2NDY3OTY4",
                                                        "rest_id": "1874474007216467968",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:13:50 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 17,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Mirla",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "mirlitapaz",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-913817634298368003",
                                        "sortIndex": "1874476833932574699",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo5MTM4MTc2MzQyOTgzNjgwMDM=",
                                                        "rest_id": "913817634298368003",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Sep 29 17:27:56 +0000 2017",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "whore's and sluts, don't follow me  #Block",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 251,
                                                            "followers_count": 17,
                                                            "friends_count": 110,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 6,
                                                            "name": "Veld Zwanen",
                                                            "normal_followers_count": 17,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/913817634298368003/1506707796",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1622935121748717569/P--K20yI_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "VeldZwanen",
                                                            "statuses_count": 169,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1859711609742000128",
                                        "sortIndex": "1874476833932574698",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODU5NzExNjA5NzQyMDAwMTI4",
                                                        "rest_id": "1859711609742000128",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Nov 21 21:33:17 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "FJB. TRUMP WON. Tesla owner and fan.\nIFBAP",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 156,
                                                            "followers_count": 33,
                                                            "friends_count": 75,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "New York, USA",
                                                            "media_count": 4,
                                                            "name": "We live in a Clown World \uD83E\uDD21\uD83C\uDF0F \uD83C\uDDFA\uD83C\uDDF8",
                                                            "normal_followers_count": 33,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1859711609742000128/1732225170",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1860441641418420224/o8OdLXVU_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "MAGADutchS516",
                                                            "statuses_count": 125,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1244337613860999172",
                                        "sortIndex": "1874476833932574697",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMjQ0MzM3NjEzODYwOTk5MTcy",
                                                        "rest_id": "1244337613860999172",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Sun Mar 29 18:56:29 +0000 2020",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "\uD83C\uDDFA\uD83C\uDDF8 US Army SAPPER (RET) | Former Federal Officer | Constitutional Conservative | Boxer | 8th Generation American | Soldier of God | Fellowcraft /G\\| \uD83C\uDDFA\uD83C\uDDF8",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 33929,
                                                            "followers_count": 3701,
                                                            "friends_count": 2153,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Florida",
                                                            "media_count": 642,
                                                            "name": "Nitestalker",
                                                            "normal_followers_count": 3701,
                                                            "pinned_tweet_ids_str": [
                                                                "1839122613416685662"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1244337613860999172/1730379586",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1839702722695540739/kHJUdgwf_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "NitestalkerX",
                                                            "statuses_count": 3385,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1100452901053444096",
                                        "sortIndex": "1874476833932574696",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMTAwNDUyOTAxMDUzNDQ0MDk2",
                                                        "rest_id": "1100452901053444096",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Feb 26 17:49:51 +0000 2019",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 11042,
                                                            "followers_count": 154,
                                                            "friends_count": 3593,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 1,
                                                            "location": "",
                                                            "media_count": 326,
                                                            "name": "Jay",
                                                            "normal_followers_count": 154,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1100452901053444096/1626762579",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1731483492544827392/i0c6h-lH_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Jay__ny1",
                                                            "statuses_count": 1667,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-15819609",
                                        "sortIndex": "1874476833932574695",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTgxOTYwOQ==",
                                                        "rest_id": "15819609",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Tue Aug 12 06:50:31 +0000 2008",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Entrepreneur. Views personal",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "ramprasad.info",
                                                                            "expanded_url": "http://www.ramprasad.info",
                                                                            "url": "https://t.co/zimnkNhvFw",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 53876,
                                                            "followers_count": 38408,
                                                            "friends_count": 1231,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 183,
                                                            "location": "Getting back to Airplanes. ",
                                                            "media_count": 1332,
                                                            "name": "Ram",
                                                            "normal_followers_count": 38408,
                                                            "pinned_tweet_ids_str": [
                                                                "1810334334592639163"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/15819609/1667925573",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1363832670216744960/67nZakX5_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "ramprasad_c",
                                                            "statuses_count": 48143,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/zimnkNhvFw",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1597004802231345152",
                                                            "professional_type": "Business",
                                                            "category": [
                                                                {
                                                                    "id": 958,
                                                                    "name": "Entrepreneur",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-836021948",
                                        "sortIndex": "1874476833932574694",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo4MzYwMjE5NDg=",
                                                        "rest_id": "836021948",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Sep 20 17:39:24 +0000 2012",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "crypto bitcoin /ethereum  traveler all over the world",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 3129,
                                                            "followers_count": 37,
                                                            "friends_count": 213,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 4,
                                                            "location": "",
                                                            "media_count": 62,
                                                            "name": "jacob janssen",
                                                            "normal_followers_count": 37,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1747117798118354944/7FUjQhL-_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "jacobjanssen2",
                                                            "statuses_count": 1282,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1608926951959396354",
                                        "sortIndex": "1874476833932574693",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNjA4OTI2OTUxOTU5Mzk2MzU0",
                                                        "rest_id": "1608926951959396354",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Dec 30 21:48:01 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "MUFC",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 68,
                                                            "followers_count": 45,
                                                            "friends_count": 421,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 8,
                                                            "name": "Grant Rostron",
                                                            "normal_followers_count": 45,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1608926951959396354/1672788131",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1804537481708097536/J-27RQX-_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "grant_rostron",
                                                            "statuses_count": 111,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1592715154331836416",
                                        "sortIndex": "1874476833932574692",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTkyNzE1MTU0MzMxODM2NDE2",
                                                        "rest_id": "1592715154331836416",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Nov 16 03:05:17 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Poet, Photographer, Provocateur: Waterman ...\nAbyssus Abyssum Invocat  ...  μολὼν λαβέ  ......\nChelsea FC Supporter... NO DMS!",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 60165,
                                                            "followers_count": 2223,
                                                            "friends_count": 1245,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 3,
                                                            "location": "San Diego",
                                                            "media_count": 1427,
                                                            "name": "Larry Kuechlin",
                                                            "normal_followers_count": 2223,
                                                            "pinned_tweet_ids_str": [
                                                                "1770121699553603769"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1592715154331836416/1728744845",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1622453811444076545/7xwEDyvc_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "KuechlinLarry",
                                                            "statuses_count": 16685,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1044570152824639488",
                                        "sortIndex": "1874476833932574691",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMDQ0NTcwMTUyODI0NjM5NDg4",
                                                        "rest_id": "1044570152824639488",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Sep 25 12:51:46 +0000 2018",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Kiryukai FOREVER",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 6061,
                                                            "followers_count": 6,
                                                            "friends_count": 103,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Texas, USA",
                                                            "media_count": 7,
                                                            "name": "FrontToEnemy",
                                                            "normal_followers_count": 6,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1049881433454641153/0L9T4rIq_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "FrontToEnemy",
                                                            "statuses_count": 72,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1074420281131851777",
                                        "sortIndex": "1874476833932574690",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMDc0NDIwMjgxMTMxODUxNzc3",
                                                        "rest_id": "1074420281131851777",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "protected": true,
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Sun Dec 16 21:45:31 +0000 2018",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "parody account. all tweet are meant for social study to reactions",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 24241,
                                                            "followers_count": 13,
                                                            "friends_count": 361,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Medicine lodge, kansas",
                                                            "media_count": 92,
                                                            "name": "stephanmurray",
                                                            "normal_followers_count": 13,
                                                            "pinned_tweet_ids_str": [
                                                                "1174813813469827073"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1074420281131851777/1600130627",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1638690350578425856/54jod3_3_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "imstevenmurray",
                                                            "statuses_count": 4977,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-988801767763857409",
                                        "sortIndex": "1874476833932574689",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo5ODg4MDE3Njc3NjM4NTc0MDk=",
                                                        "rest_id": "988801767763857409",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Apr 24 15:28:07 +0000 2018",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 62,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Josh Michael",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "joshuamichael98",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874475661613744128",
                                        "sortIndex": "1874476833932574688",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDc1NjYxNjEzNzQ0MTI4",
                                                        "rest_id": "1874475661613744128",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:20:09 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 4,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Jeannette Urena",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "JeannetteU38964",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-948581581433819136",
                                        "sortIndex": "1874476833932574687",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo5NDg1ODE1ODE0MzM4MTkxMzY=",
                                                        "rest_id": "948581581433819136",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 03 15:47:27 +0000 2018",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 1,
                                                            "followers_count": 1,
                                                            "friends_count": 1,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Assenede, België",
                                                            "media_count": 0,
                                                            "name": "Peter Audenaert",
                                                            "normal_followers_count": 1,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "p_audenaert",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874475485343944706",
                                        "sortIndex": "1874476833932574686",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDc1NDg1MzQzOTQ0NzA2",
                                                        "rest_id": "1874475485343944706",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:19:33 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 3,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Juan Matthews",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "JuanMatthe62875",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1760860414470955008",
                                        "sortIndex": "1874476833932574685",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzYwODYwNDE0NDcwOTU1MDA4",
                                                        "rest_id": "1760860414470955008",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Feb 23 02:53:56 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 699,
                                                            "followers_count": 742,
                                                            "friends_count": 1969,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Washington, DC",
                                                            "media_count": 42,
                                                            "name": "Reluctantant Fundamentalist",
                                                            "normal_followers_count": 742,
                                                            "pinned_tweet_ids_str": [
                                                                "1873799279040749916"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1873808221623693312/8MQ-Axtl_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "sartoAi",
                                                            "statuses_count": 4406,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": true
                                                        }
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1586505435707850753",
                                        "sortIndex": "1874476833932574684",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTg2NTA1NDM1NzA3ODUwNzUz",
                                                        "rest_id": "1586505435707850753",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Oct 29 23:49:57 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 7,
                                                            "followers_count": 9,
                                                            "friends_count": 182,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Publius",
                                                            "normal_followers_count": 9,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Publius_Z",
                                                            "statuses_count": 2,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1863159216715821056",
                                        "sortIndex": "1874476833932574683",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODYzMTU5MjE2NzE1ODIxMDU2",
                                                        "rest_id": "1863159216715821056",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sun Dec 01 09:52:45 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Christian✝️✝️ Army Veteran\uD83C\uDDEC\uD83C\uDDE7\nFollow and il follow back.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 138,
                                                            "followers_count": 98,
                                                            "friends_count": 537,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 6,
                                                            "name": "Ralph Garnett ✝️\uD83C\uDDEC\uD83C\uDDE7\uD83C\uDFF4\uDB40\uDC67\uDB40\uDC62\uDB40\uDC65\uDB40\uDC6E\uDB40\uDC67\uDB40\uDC7F",
                                                            "normal_followers_count": 98,
                                                            "pinned_tweet_ids_str": [
                                                                "1863251731548721462"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1863159216715821056/1733046947",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1863159761761337344/07gRsm-3_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "garnett75406",
                                                            "statuses_count": 115,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1346909520510017538",
                                        "sortIndex": "1874476833932574682",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMzQ2OTA5NTIwNTEwMDE3NTM4",
                                                        "rest_id": "1346909520510017538",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 06 20:00:45 +0000 2021",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 114,
                                                            "followers_count": 4,
                                                            "friends_count": 182,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 9,
                                                            "name": "Moss88",
                                                            "normal_followers_count": 4,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Moss8812",
                                                            "statuses_count": 61,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-18499767",
                                        "sortIndex": "1874476833932574681",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODQ5OTc2Nw==",
                                                        "rest_id": "18499767",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Dec 31 05:11:13 +0000 2008",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Retired, mom, grandma. Lover of sunshine.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 980,
                                                            "followers_count": 88,
                                                            "friends_count": 411,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 2,
                                                            "location": "Oro Valley, AZ",
                                                            "media_count": 5,
                                                            "name": "Terri Thomson",
                                                            "normal_followers_count": 88,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/18499767/1701902618",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1774493127661010944/WQzQBxQd_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "terriinhercave",
                                                            "statuses_count": 459,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1848547057788923904",
                                        "sortIndex": "1874476833932574680",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODQ4NTQ3MDU3Nzg4OTIzOTA0",
                                                        "rest_id": "1848547057788923904",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "protected": true,
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Tue Oct 22 02:09:58 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 742,
                                                            "followers_count": 0,
                                                            "friends_count": 338,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Danny",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Danny492175",
                                                            "statuses_count": 2,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1586488170555658240",
                                        "sortIndex": "1874476833932574679",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTg2NDg4MTcwNTU1NjU4MjQw",
                                                        "rest_id": "1586488170555658240",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Oct 29 22:40:59 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Trust Common Sense.  Believe rules are for the guidance of wise men and the obedience of fools.  Optimistic about the future.  Dad & Hubby. Skier. Paraglider.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 402,
                                                            "followers_count": 122,
                                                            "friends_count": 194,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 19,
                                                            "name": "Mike W",
                                                            "normal_followers_count": 122,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1586488170555658240/1733192910",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1863773132492013568/y-y147Dz_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "freeflyer65",
                                                            "statuses_count": 2013,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-860479564929224704",
                                        "sortIndex": "1874476833932574678",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo4NjA0Nzk1NjQ5MjkyMjQ3MDQ=",
                                                        "rest_id": "860479564929224704",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri May 05 13:01:30 +0000 2017",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "Country girl, love dogs, like Europe, hate the EU and governments that shackled us to it. Believer in our country and Brexit. NO DMs, please.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 298226,
                                                            "followers_count": 4497,
                                                            "friends_count": 5366,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 3,
                                                            "location": "South East, England",
                                                            "media_count": 176,
                                                            "name": "jane7",
                                                            "normal_followers_count": 4497,
                                                            "pinned_tweet_ids_str": [
                                                                "1090203072922570752"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "jane798m",
                                                            "statuses_count": 60868,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-2568591631",
                                        "sortIndex": "1874476833932574677",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyNTY4NTkxNjMx",
                                                        "rest_id": "2568591631",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sun Jun 15 07:41:47 +0000 2014",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 1,
                                                            "friends_count": 11,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "석지훈",
                                                            "normal_followers_count": 1,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "sukz7221",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1741156629687500800",
                                        "sortIndex": "1874476833932574676",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzQxMTU2NjI5Njg3NTAwODAw",
                                                        "rest_id": "1741156629687500800",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Dec 30 17:57:54 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Highlighting the criminally under appreciated Gulf South | Leasing & Investment Sales | New Orleans ⚜️",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 787,
                                                            "followers_count": 126,
                                                            "friends_count": 239,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 1,
                                                            "location": "",
                                                            "media_count": 43,
                                                            "name": "NolaCRE",
                                                            "normal_followers_count": 126,
                                                            "pinned_tweet_ids_str": [
                                                                "1788908685621657961"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1741156629687500800/1710725108",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1741162940873527296/TUWmaTA5_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "nola_cre",
                                                            "statuses_count": 355,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1741168795329151044",
                                                            "professional_type": "Business",
                                                            "category": [
                                                                {
                                                                    "id": 697,
                                                                    "name": "Real Estate",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874473488100286464",
                                        "sortIndex": "1874476833932574675",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDczNDg4MTAwMjg2NDY0",
                                                        "rest_id": "1874473488100286464",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:12:12 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 12,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Donna h",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1874473832670806016/CFfL39oo_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Donnah1060641",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1427491585646764032",
                                        "sortIndex": "1874476833932574674",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNDI3NDkxNTg1NjQ2NzY0MDMy",
                                                        "rest_id": "1427491585646764032",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Aug 17 04:45:07 +0000 2021",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "\uD83C\uDDFA\uD83C\uDDF2TRUMP2024\uD83C\uDDFA\uD83C\uDDF2 \nPissed off CA Patriot! \nGod~Family~Married~2A~FAFO~4 wheelin & Freedom\uD83C\uDDFA\uD83C\uDDF2 Vets \nB-4 Illegals\uD83C\uDDFA\uD83C\uDDF2IFBAP  \n\uD83D\uDEABNO DM'S\nI identify as #FAFO",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 227014,
                                                            "followers_count": 6271,
                                                            "friends_count": 6261,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Commiefornia ",
                                                            "media_count": 616,
                                                            "name": "\uD835\uDE3C\uD835\uDE41 \uD835\uDE48\uD835\uDE3C\uD835\uDE42\uD835\uDE3C \uD835\uDE3D\uD835\uDE67\uD835\uDE56\uD835\uDE69 \uD83E\uDD85\uD83C\uDDFA\uD83C\uDDF2",
                                                            "normal_followers_count": 6271,
                                                            "pinned_tweet_ids_str": [
                                                                "1494387113768341515"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1427491585646764032/1712001680",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1770197490681466880/aWQiO_Ln_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Hhowell1969",
                                                            "statuses_count": 16583,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1238532957138190343",
                                        "sortIndex": "1874476833932574673",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMjM4NTMyOTU3MTM4MTkwMzQz",
                                                        "rest_id": "1238532957138190343",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Fri Mar 13 18:30:55 +0000 2020",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "слишком дикий, чтобы жить. слишком редкий, чтобы сдохнуть.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 2480,
                                                            "followers_count": 17,
                                                            "friends_count": 98,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Saint Petersburg, Russia",
                                                            "media_count": 3,
                                                            "name": "Yan Solheim",
                                                            "normal_followers_count": 17,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1238532957138190343/1711898558",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1774457285172944896/RV5zuqfv_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Yan_Solheim",
                                                            "statuses_count": 34,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-4215151001",
                                        "sortIndex": "1874476833932574672",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0MjE1MTUxMDAx",
                                                        "rest_id": "4215151001",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Nov 18 04:09:47 +0000 2015",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 8,
                                                            "followers_count": 1,
                                                            "friends_count": 59,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Mikey Mike",
                                                            "normal_followers_count": 1,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "mrdude800",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1324721638282178560",
                                        "sortIndex": "1874476833932574671",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMzI0NzIxNjM4MjgyMTc4NTYw",
                                                        "rest_id": "1324721638282178560",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Nov 06 14:33:56 +0000 2020",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Catholic Tech Geek Deacon helping others come to Christ with new technologies",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 77,
                                                            "followers_count": 5,
                                                            "friends_count": 166,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Geeky_Deacon",
                                                            "normal_followers_count": 5,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1324721638282178560/1604693930",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1324808210939670528/YQY_n1fj_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Geeky_Deacon",
                                                            "statuses_count": 2,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874474241414086656",
                                        "sortIndex": "1874476833932574670",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDc0MjQxNDE0MDg2NjU2",
                                                        "rest_id": "1874474241414086656",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:14:29 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 5,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Joyce Buckmon",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "BuckmonJoy19604",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-899074157874262016",
                                        "sortIndex": "1874476833932574669",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo4OTkwNzQxNTc4NzQyNjIwMTY=",
                                                        "rest_id": "899074157874262016",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sun Aug 20 01:02:38 +0000 2017",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 2631,
                                                            "followers_count": 33,
                                                            "friends_count": 119,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Bennie",
                                                            "normal_followers_count": 33,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1839428073571381248/kT3R2DZK_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Bennie_Q_B",
                                                            "statuses_count": 197,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1372742190317895682",
                                        "sortIndex": "1874476833932574668",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMzcyNzQyMTkwMzE3ODk1Njgy",
                                                        "rest_id": "1372742190317895682",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Mar 19 02:50:24 +0000 2021",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 18,
                                                            "followers_count": 2,
                                                            "friends_count": 67,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Cow_goes_moooooo",
                                                            "normal_followers_count": 2,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "GoesMoooooo",
                                                            "statuses_count": 7,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1868253095547088896",
                                        "sortIndex": "1874476833932574667",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODY4MjUzMDk1NTQ3MDg4ODk2",
                                                        "rest_id": "1868253095547088896",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "protected": true,
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Sun Dec 15 11:14:16 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 3257,
                                                            "followers_count": 0,
                                                            "friends_count": 192,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Rams",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": true,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1868253238279163904/WJQypP2K_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "rams340034",
                                                            "statuses_count": 696,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874474129803591680",
                                        "sortIndex": "1874476833932574666",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDc0MTI5ODAzNTkxNjgw",
                                                        "rest_id": "1874474129803591680",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:14:06 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 4,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Sharon Malone",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "SharonMalo85176",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874473466079981568",
                                        "sortIndex": "1874476833932574665",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDczNDY2MDc5OTgxNTY4",
                                                        "rest_id": "1874473466079981568",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:11:27 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 98,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Adriana",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1874473700449239040/YKIBw3Yq_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Adriana52584022",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1874473669151670272",
                                        "sortIndex": "1874476833932574664",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODc0NDczNjY5MTUxNjcwMjcy",
                                                        "rest_id": "1874473669151670272",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Jan 01 15:12:21 +0000 2025",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 0,
                                                            "friends_count": 5,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Lindsay Hunt",
                                                            "normal_followers_count": 0,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "LindsayHun97480",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-707535824355778560",
                                        "sortIndex": "1874476833932574663",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo3MDc1MzU4MjQzNTU3Nzg1NjA=",
                                                        "rest_id": "707535824355778560",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Wed Mar 09 11:57:40 +0000 2016",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Für immer vegan \uD83D\uDC30\uD83E\uDD8A\uD83E\uDD9D\uD83D\uDC2E\uD83D\uDC37\uD83D\uDC39\uD83E\uDD84\uD83D\uDC38\uD83D\uDC17\uD83D\uDC36\uD83D\uDC9A\uD83C\uDF31\nkickboxing is my passion\uD83E\uDD4A\uD83E\uDD4A                                                        i stand with israel ✡️\uD83C\uDDEE\uD83C\uDDF1\uD83C\uDDE9\uD83C\uDDEA",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 20301,
                                                            "followers_count": 476,
                                                            "friends_count": 861,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 707,
                                                            "name": "Petry",
                                                            "normal_followers_count": 476,
                                                            "pinned_tweet_ids_str": [
                                                                "1836090802872947113"
                                                            ],
                                                            "possibly_sensitive": true,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/707535824355778560/1492023286",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1788152656407547904/Uk2c2pT2_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "PetryPp51",
                                                            "statuses_count": 3203,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1801710551531405314",
                                        "sortIndex": "1874476833932574662",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODAxNzEwNTUxNTMxNDA1MzE0",
                                                        "rest_id": "1801710551531405314",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Fri Jun 14 20:17:47 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 954,
                                                            "followers_count": 4,
                                                            "friends_count": 47,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 1,
                                                            "name": "hsder",
                                                            "normal_followers_count": 4,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1807788413783068672/Smyej8oD_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "gwKerr286382",
                                                            "statuses_count": 152,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1729441176002985984",
                                        "sortIndex": "1874476833932574661",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzI5NDQxMTc2MDAyOTg1OTg0",
                                                        "rest_id": "1729441176002985984",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Nov 28 10:04:57 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 2219,
                                                            "followers_count": 17,
                                                            "friends_count": 120,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 360,
                                                            "name": "Figas",
                                                            "normal_followers_count": 17,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1746839769949028352/gQAyj8ls_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "NunoFiguei38176",
                                                            "statuses_count": 1159,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1482287944845168640",
                                        "sortIndex": "1874476833932574660",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNDgyMjg3OTQ0ODQ1MTY4NjQw",
                                                        "rest_id": "1482287944845168640",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Jan 15 09:46:47 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 30,
                                                            "followers_count": 1,
                                                            "friends_count": 10,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Mojtaba Hosseini",
                                                            "normal_followers_count": 1,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1482288144775008261/gQD8Fwx4_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Hosseinii07",
                                                            "statuses_count": 9,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1457290890632962048",
                                        "sortIndex": "1874476833932574659",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNDU3MjkwODkwNjMyOTYyMDQ4",
                                                        "rest_id": "1457290890632962048",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sun Nov 07 10:17:04 +0000 2021",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Retired Law Enforcement, USAF Veteran- NO DMs I will not answer them. \n\uD83C\uDDFA\uD83C\uDDF8\uD83C\uDDEE\uD83C\uDDF1",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 107777,
                                                            "followers_count": 935,
                                                            "friends_count": 3874,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 2,
                                                            "location": "Dirty Jersey ",
                                                            "media_count": 419,
                                                            "name": "JW",
                                                            "normal_followers_count": 935,
                                                            "pinned_tweet_ids_str": [
                                                                "1595794578455801858"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1457290890632962048/1725932924",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1838363971956436992/iH4W6PPl_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "JasonWalk14",
                                                            "statuses_count": 22145,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-20495152",
                                        "sortIndex": "1874476833932574658",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyMDQ5NTE1Mg==",
                                                        "rest_id": "20495152",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Feb 10 06:17:15 +0000 2009",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "I love Beer, Bourbon, Bacon, Butter and My Family! #teamrealiy",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 34420,
                                                            "followers_count": 709,
                                                            "friends_count": 1750,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 5,
                                                            "location": "Deep in the Blue Country",
                                                            "media_count": 2997,
                                                            "name": "Not Tyler Durden",
                                                            "normal_followers_count": 709,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/20495152/1593840705",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1552747352187801601/03gIZPfL_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "mikediller",
                                                            "statuses_count": 28916,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1736122727562448896",
                                        "sortIndex": "1874476833932574657",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzM2MTIyNzI3NTYyNDQ4ODk2",
                                                        "rest_id": "1736122727562448896",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Dec 16 20:35:06 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 357,
                                                            "followers_count": 6,
                                                            "friends_count": 50,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Lorenz",
                                                            "normal_followers_count": 6,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Lorenz20005",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-322331585",
                                        "sortIndex": "1874476833932574656",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjozMjIzMzE1ODU=",
                                                        "rest_id": "322331585",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Jun 23 00:44:11 +0000 2011",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "treading lightly \uD83E\uDD97",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "last.fm/user/ezears",
                                                                            "expanded_url": "http://www.last.fm/user/ezears",
                                                                            "url": "https://t.co/N7KphqVTvl",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 50306,
                                                            "followers_count": 1752,
                                                            "friends_count": 1078,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 35,
                                                            "location": "CA DISGRACELAND",
                                                            "media_count": 1995,
                                                            "name": "Esther \uD83C\uDF43\uD83C\uDF42",
                                                            "normal_followers_count": 1752,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/322331585/1732232250",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/689624716127014912/5fMiYCFy_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "myrtleleaf",
                                                            "statuses_count": 47356,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/N7KphqVTvl",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1117684730",
                                        "sortIndex": "1874476833932574655",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMTE3Njg0NzMw",
                                                        "rest_id": "1117684730",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "protected": true,
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Thu Jan 24 20:30:14 +0000 2013",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 8,
                                                            "followers_count": 5,
                                                            "friends_count": 80,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "SpecialOperator",
                                                            "normal_followers_count": 5,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1117684730/1735128013",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1871886835376427008/iRKvTA9w_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "doctorkowalski",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-735121468867149827",
                                        "sortIndex": "1874476833932574654",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo3MzUxMjE0Njg4NjcxNDk4Mjc=",
                                                        "rest_id": "735121468867149827",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": false,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue May 24 14:53:11 +0000 2016",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 2,
                                                            "followers_count": 5,
                                                            "friends_count": 39,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "imagineyourshape",
                                                            "normal_followers_count": 5,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "imagineyoursha1",
                                                            "statuses_count": 0,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1501626359118016520",
                                        "sortIndex": "1874476833932574653",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTAxNjI2MzU5MTE4MDE2NTIw",
                                                        "rest_id": "1501626359118016520",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Mar 09 18:30:31 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Your mind is like a parachute, it does not work, if you don’t open it…           Es ist gefährlich, Recht zu haben, wenn die Regierung Unrecht hat. - Voltaire -",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 0,
                                                            "followers_count": 30,
                                                            "friends_count": 599,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "roli ka",
                                                            "normal_followers_count": 30,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": true,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1501626359118016520/1647423077",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1703309665705021440/D4r15NQ7_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "rolika0669",
                                                            "statuses_count": 1,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1729594893348089856",
                                        "sortIndex": "1874476833932574652",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzI5NTk0ODkzMzQ4MDg5ODU2",
                                                        "rest_id": "1729594893348089856",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Nov 28 20:15:47 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "I’m a human beeing, 36 y., german, psychologist. Posting mix of own thoughts, links & quotes in eng and ger\uD83C\uDDE9\uD83C\uDDEA\uD83C\uDDEC\uD83C\uDDE7 \uD83C\uDDEA\uD83C\uDDFA\uD83C\uDDFA\uD83C\uDDF3\n#peace #humanity \uD83D\uDD4A️✌\uD83C\uDFFB",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 2070,
                                                            "followers_count": 73,
                                                            "friends_count": 259,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Berlin, Deutschland",
                                                            "media_count": 137,
                                                            "name": "veritas",
                                                            "normal_followers_count": 73,
                                                            "pinned_tweet_ids_str": [
                                                                "1874437034120814600"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1729594893348089856/1716100872",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1729597802840236032/uBX_7m8E_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "veritas_psy",
                                                            "statuses_count": 561,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "cursor-bottom-1874476833932574651",
                                        "sortIndex": "1874476833932574651",
                                        "content": {
                                            "entryType": "TimelineTimelineCursor",
                                            "__typename": "TimelineTimelineCursor",
                                            "value": "1820059673811707107|1874476833932574649",
                                            "cursorType": "Bottom"
                                        }
                                    },
                                    {
                                        "entryId": "cursor-top-1874476833932574721",
                                        "sortIndex": "1874476833932574721",
                                        "content": {
                                            "entryType": "TimelineTimelineCursor",
                                            "__typename": "TimelineTimelineCursor",
                                            "value": "-1|1874476833932574721",
                                            "cursorType": "Top"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                }
            }
        }
    }
}
```

#### UserByRestIDs

###### Request

```javascript 
fetch("https://x.com/i/api/graphql/JnwU1UO8J1tWlOJktPZIzg/UsersByRestIds?variables=%7B%22userIds%22%3A%5B%2249676106%22%5D%7D&features=%7B%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "x2SVFww01APT5Bym1qJFt8TAbFhnFHcF8hXV9ddJi4cx51nDxw6/dSvMLbDmMFNhOns94sRf6Tnmwg5BrRfPY0PYq01jxA",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/visegrad24/following",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

###### Payload

```javascript
variables: {"userIds":["49676106"]}
features: {"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"responsive_web_graphql_timeline_navigation_enabled":true}
```

###### Response

``` javascript
{
    "data": {
        "users": [
            {
                "result": {
                    "__typename": "User",
                    "id": "VXNlcjo0OTY3NjEwNg==",
                    "rest_id": "49676106",
                    "affiliates_highlighted_label": {},
                    "has_graduated_access": true,
                    "is_blue_verified": true,
                    "profile_image_shape": "Circle",
                    "legacy": {
                        "following": false,
                        "can_dm": true,
                        "can_media_tag": false,
                        "created_at": "Mon Jun 22 16:17:33 +0000 2009",
                        "default_profile": false,
                        "default_profile_image": false,
                        "description": "automation absolutist • python shill • follow for updates on agents and prompt engineering",
                        "entities": {
                            "description": {
                                "urls": []
                            },
                            "url": {
                                "urls": [
                                    {
                                        "display_url": "github.com/alexfazio",
                                        "expanded_url": "http://github.com/alexfazio",
                                        "url": "https://t.co/n2pYx8ZFgs",
                                        "indices": [
                                            0,
                                            23
                                        ]
                                    }
                                ]
                            }
                        },
                        "fast_followers_count": 0,
                        "favourites_count": 12035,
                        "followers_count": 3719,
                        "friends_count": 824,
                        "has_custom_timelines": false,
                        "is_translator": false,
                        "listed_count": 67,
                        "location": "latent space",
                        "media_count": 919,
                        "name": "ale\uD835\uDD4F fazio",
                        "normal_followers_count": 3719,
                        "pinned_tweet_ids_str": [
                            "1873764113018708158"
                        ],
                        "possibly_sensitive": false,
                        "profile_banner_url": "https://pbs.twimg.com/profile_banners/49676106/1734108904",
                        "profile_image_url_https": "https://pbs.twimg.com/profile_images/1865367901366075393/JrM1ff_b_normal.jpg",
                        "profile_interstitial_type": "",
                        "screen_name": "alxfazio",
                        "statuses_count": 2676,
                        "translator_type": "none",
                        "url": "https://t.co/n2pYx8ZFgs",
                        "verified": false,
                        "want_retweets": false,
                        "withheld_in_countries": []
                    },
                    "tipjar_settings": {}
                }
            }
        ]
    }
}
```

#### BlueVerifiedFollowers

###### Response

```javascript

fetch("https://x.com/i/api/graphql/srYtCtUs5BuBPbYj7agW6A/BlueVerifiedFollowers?variables=%7B%22userId%22%3A%221222773302441148416%22%2C%22count%22%3A20%2C%22includePromotedContent%22%3Afalse%7D&features=%7B%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22premium_content_api_read_enabled%22%3Afalse%2C%22communities_web_enable_tweet_community_results_fetch%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_button_fetch_trends_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_post_followups_enabled%22%3Afalse%2C%22responsive_web_grok_share_attachment_enabled%22%3Atrue%2C%22articles_preview_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22creator_subscriptions_quote_tweet_preview_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "Oplo6vHJKf4uGeFbK1+4Sjk9kaWa6Yr4D+goCCq0dnrMGqQ+OvNCiNYx0E0bza6cx4DBHzndFKSDzswWwJcV6Fm2uF6tOQ",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/visegrad24/verified_followers",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

###### Payload


```javascript
variables: {"userId":"1222773302441148416","count":20,"includePromotedContent":false}
features: {"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"premium_content_api_read_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"responsive_web_grok_analyze_button_fetch_trends_enabled":true,"responsive_web_grok_analyze_post_followups_enabled":false,"responsive_web_grok_share_attachment_enabled":true,"articles_preview_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"rweb_video_timestamps_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"responsive_web_enhance_cards_enabled":false}
```


###### Response

```javascript
{
    "data": {
        "user": {
            "result": {
                "__typename": "User",
                "timeline": {
                    "timeline": {
                        "instructions": [
                            {
                                "type": "TimelineClearCache"
                            },
                            {
                                "type": "TimelineTerminateTimeline",
                                "direction": "Top"
                            },
                            {
                                "type": "TimelineTerminateTimeline",
                                "direction": "Bottom"
                            },
                            {
                                "type": "TimelineAddEntries",
                                "entries": [
                                    {
                                        "entryId": "user-1594889826159001600",
                                        "sortIndex": "1874477859959996416",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTk0ODg5ODI2MTU5MDAxNjAw",
                                                        "rest_id": "1594889826159001600",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Tue Nov 22 03:06:02 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Husband. Engineer. Curious.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 11149,
                                                            "followers_count": 788,
                                                            "friends_count": 1351,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 9,
                                                            "location": "Tulsa, OK",
                                                            "media_count": 13,
                                                            "name": "Mike Dearinger",
                                                            "normal_followers_count": 788,
                                                            "pinned_tweet_ids_str": [
                                                                "1778752421373583816"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1594889826159001600/1671228312",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1720309275686223872/xflZY7HS_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "AlienMiked",
                                                            "statuses_count": 5772,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-259679851",
                                        "sortIndex": "1874477859959996415",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyNTk2Nzk4NTE=",
                                                        "rest_id": "259679851",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Wed Mar 02 11:18:09 +0000 2011",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 354038,
                                                            "followers_count": 1474,
                                                            "friends_count": 1422,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 1,
                                                            "location": "Texas, USA",
                                                            "media_count": 48,
                                                            "name": "Anna M.",
                                                            "normal_followers_count": 1474,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/259679851/1700354624",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/558783198696452097/6FnKzVFq_normal.jpeg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "AnnaGilpin",
                                                            "statuses_count": 3377,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-929912597519925248",
                                        "sortIndex": "1874477859959996414",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo5Mjk5MTI1OTc1MTk5MjUyNDg=",
                                                        "rest_id": "929912597519925248",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Mon Nov 13 03:23:34 +0000 2017",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "❤️ #married 26 yr USMC brat. Semper Fi. I tweet stocks and politics, I am a conservative.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 212925,
                                                            "followers_count": 27417,
                                                            "friends_count": 20002,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 26,
                                                            "location": "Fly Over Country. ",
                                                            "media_count": 9897,
                                                            "name": "Midwest Iowa girl",
                                                            "normal_followers_count": 27417,
                                                            "pinned_tweet_ids_str": [
                                                                "1578141336070283264"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/929912597519925248/1688877817",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1862316304377446402/cMMkrBv1_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Maga4Justice",
                                                            "statuses_count": 252206,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1859711609742000128",
                                        "sortIndex": "1874477859959996413",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODU5NzExNjA5NzQyMDAwMTI4",
                                                        "rest_id": "1859711609742000128",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Nov 21 21:33:17 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "FJB. TRUMP WON. Tesla owner and fan.\nIFBAP",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 157,
                                                            "followers_count": 33,
                                                            "friends_count": 75,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "New York, USA",
                                                            "media_count": 4,
                                                            "name": "We live in a Clown World \uD83E\uDD21\uD83C\uDF0F \uD83C\uDDFA\uD83C\uDDF8",
                                                            "normal_followers_count": 33,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1859711609742000128/1732225170",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1860441641418420224/o8OdLXVU_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "MAGADutchS516",
                                                            "statuses_count": 125,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1244337613860999172",
                                        "sortIndex": "1874477859959996412",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMjQ0MzM3NjEzODYwOTk5MTcy",
                                                        "rest_id": "1244337613860999172",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Sun Mar 29 18:56:29 +0000 2020",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "\uD83C\uDDFA\uD83C\uDDF8 US Army SAPPER (RET) | Former Federal Officer | Constitutional Conservative | Boxer | 8th Generation American | Soldier of God | Fellowcraft /G\\| \uD83C\uDDFA\uD83C\uDDF8",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 33940,
                                                            "followers_count": 3701,
                                                            "friends_count": 2153,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Florida",
                                                            "media_count": 642,
                                                            "name": "Nitestalker",
                                                            "normal_followers_count": 3701,
                                                            "pinned_tweet_ids_str": [
                                                                "1839122613416685662"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1244337613860999172/1730379586",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1839702722695540739/kHJUdgwf_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "NitestalkerX",
                                                            "statuses_count": 3386,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-15819609",
                                        "sortIndex": "1874477859959996411",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTgxOTYwOQ==",
                                                        "rest_id": "15819609",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Tue Aug 12 06:50:31 +0000 2008",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "Entrepreneur. Views personal",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "ramprasad.info",
                                                                            "expanded_url": "http://www.ramprasad.info",
                                                                            "url": "https://t.co/zimnkNhvFw",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 53879,
                                                            "followers_count": 38408,
                                                            "friends_count": 1231,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 183,
                                                            "location": "Getting back to Airplanes. ",
                                                            "media_count": 1332,
                                                            "name": "Ram",
                                                            "normal_followers_count": 38408,
                                                            "pinned_tweet_ids_str": [
                                                                "1810334334592639163"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/15819609/1667925573",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1363832670216744960/67nZakX5_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "ramprasad_c",
                                                            "statuses_count": 48144,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/zimnkNhvFw",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1597004802231345152",
                                                            "professional_type": "Business",
                                                            "category": [
                                                                {
                                                                    "id": 958,
                                                                    "name": "Entrepreneur",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1592715154331836416",
                                        "sortIndex": "1874477859959996410",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTkyNzE1MTU0MzMxODM2NDE2",
                                                        "rest_id": "1592715154331836416",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Wed Nov 16 03:05:17 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Poet, Photographer, Provocateur: Waterman ...\nAbyssus Abyssum Invocat  ...  μολὼν λαβέ  ......\nChelsea FC Supporter... NO DMS!",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 60165,
                                                            "followers_count": 2223,
                                                            "friends_count": 1245,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 3,
                                                            "location": "San Diego",
                                                            "media_count": 1427,
                                                            "name": "Larry Kuechlin",
                                                            "normal_followers_count": 2223,
                                                            "pinned_tweet_ids_str": [
                                                                "1770121699553603769"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1592715154331836416/1728744845",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1622453811444076545/7xwEDyvc_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "KuechlinLarry",
                                                            "statuses_count": 16685,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1044570152824639488",
                                        "sortIndex": "1874477859959996409",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMDQ0NTcwMTUyODI0NjM5NDg4",
                                                        "rest_id": "1044570152824639488",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Sep 25 12:51:46 +0000 2018",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Kiryukai FOREVER",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 6061,
                                                            "followers_count": 6,
                                                            "friends_count": 103,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Texas, USA",
                                                            "media_count": 7,
                                                            "name": "FrontToEnemy",
                                                            "normal_followers_count": 6,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1049881433454641153/0L9T4rIq_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "FrontToEnemy",
                                                            "statuses_count": 72,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1760860414470955008",
                                        "sortIndex": "1874477859959996408",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzYwODYwNDE0NDcwOTU1MDA4",
                                                        "rest_id": "1760860414470955008",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Fri Feb 23 02:53:56 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 699,
                                                            "followers_count": 742,
                                                            "friends_count": 1969,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Washington, DC",
                                                            "media_count": 42,
                                                            "name": "Reluctantant Fundamentalist",
                                                            "normal_followers_count": 742,
                                                            "pinned_tweet_ids_str": [
                                                                "1873799279040749916"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1873808221623693312/8MQ-Axtl_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "sartoAi",
                                                            "statuses_count": 4406,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": true
                                                        }
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1586505435707850753",
                                        "sortIndex": "1874477859959996407",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTg2NTA1NDM1NzA3ODUwNzUz",
                                                        "rest_id": "1586505435707850753",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Oct 29 23:49:57 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": true,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 7,
                                                            "followers_count": 9,
                                                            "friends_count": 182,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 0,
                                                            "name": "Publius",
                                                            "normal_followers_count": 9,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "Publius_Z",
                                                            "statuses_count": 2,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1586488170555658240",
                                        "sortIndex": "1874477859959996406",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTg2NDg4MTcwNTU1NjU4MjQw",
                                                        "rest_id": "1586488170555658240",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Oct 29 22:40:59 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Trust Common Sense.  Believe rules are for the guidance of wise men and the obedience of fools.  Optimistic about the future.  Dad & Hubby. Skier. Paraglider.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 402,
                                                            "followers_count": 122,
                                                            "friends_count": 194,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 19,
                                                            "name": "Mike W",
                                                            "normal_followers_count": 122,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1586488170555658240/1733192910",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1863773132492013568/y-y147Dz_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "freeflyer65",
                                                            "statuses_count": 2013,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1734658526448906240",
                                        "sortIndex": "1874477859959996405",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzM0NjU4NTI2NDQ4OTA2MjQw",
                                                        "rest_id": "1734658526448906240",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Dec 12 19:37:12 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Grow your X account with the professor \uD83D\uDD25\n\nAmerica First \uD83C\uDDFA\uD83C\uDDF2 \nBorn again Christian \uD83D\uDE4F\nProfessional list hunter \uD83C\uDFAF\nTurn notifications on and grow with me \uD83D\uDC4A",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 89664,
                                                            "followers_count": 24353,
                                                            "friends_count": 22044,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 10,
                                                            "location": "The Backcountry",
                                                            "media_count": 7367,
                                                            "name": "Professor Squatch",
                                                            "normal_followers_count": 24353,
                                                            "pinned_tweet_ids_str": [
                                                                "1874466861762380091"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1734658526448906240/1702411259",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1824600273689440257/xAFFYYUB_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "ProSquatch",
                                                            "statuses_count": 61960,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1872094296716816384",
                                        "sortIndex": "1874477859959996404",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODcyMDk0Mjk2NzE2ODE2Mzg0",
                                                        "rest_id": "1872094296716816384",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Dec 26 01:37:24 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Анализируем мировую политику через республиканскую/консервативную призму. Боремся с антизападной и российской пропагандой. За западные ценности. @RepublicTakes",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 269,
                                                            "followers_count": 9,
                                                            "friends_count": 17,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 5,
                                                            "name": "Республиканские тейки",
                                                            "normal_followers_count": 9,
                                                            "pinned_tweet_ids_str": [
                                                                "1873147085136032188"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1872094296716816384/1735179273",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1872094445237055488/CyYj2J7y_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "RespublikaTeiki",
                                                            "statuses_count": 60,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-20495152",
                                        "sortIndex": "1874477859959996403",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoyMDQ5NTE1Mg==",
                                                        "rest_id": "20495152",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Tue Feb 10 06:17:15 +0000 2009",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "I love Beer, Bourbon, Bacon, Butter and My Family! #teamrealiy",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 34420,
                                                            "followers_count": 709,
                                                            "friends_count": 1750,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 5,
                                                            "location": "Deep in the Blue Country",
                                                            "media_count": 2997,
                                                            "name": "Not Tyler Durden",
                                                            "normal_followers_count": 709,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/20495152/1593840705",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1552747352187801601/03gIZPfL_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "mikediller",
                                                            "statuses_count": 28916,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1120000410",
                                        "sortIndex": "1874477859959996402",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMTIwMDAwNDEw",
                                                        "rest_id": "1120000410",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Fri Jan 25 18:47:35 +0000 2013",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "“Big Government means small citizens: it corrodes the integrity of a people, catastrophically.” ― Mark Steyn",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 60950,
                                                            "followers_count": 1690,
                                                            "friends_count": 3125,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 55,
                                                            "location": "Gaspé QC, Canada",
                                                            "media_count": 2056,
                                                            "name": "Dylan",
                                                            "normal_followers_count": 1690,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1120000410/1732365957",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1735316109279019008/7EKR1OW4_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "DylanofGaspe",
                                                            "statuses_count": 53869,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1815154023193858048",
                                        "sortIndex": "1874477859959996401",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODE1MTU0MDIzMTkzODU4MDQ4",
                                                        "rest_id": "1815154023193858048",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Sun Jul 21 22:37:06 +0000 2024",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 6116,
                                                            "followers_count": 291,
                                                            "friends_count": 601,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "",
                                                            "media_count": 196,
                                                            "name": "DonnaLyn67",
                                                            "normal_followers_count": 291,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1815154023193858048/1721602366",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1815157576847314944/ef-8WU6w_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "DonnaLyn67",
                                                            "statuses_count": 3780,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1678142148162584576",
                                        "sortIndex": "1874477859959996400",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNjc4MTQyMTQ4MTYyNTg0NTc2",
                                                        "rest_id": "1678142148162584576",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Sun Jul 09 20:41:25 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "I’m going to try something different! only post positive things! I’m only here to have fun! You couldn’t possibly offend me because of my sense of humor!",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 3172,
                                                            "followers_count": 348,
                                                            "friends_count": 245,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 3,
                                                            "location": "United States",
                                                            "media_count": 1404,
                                                            "name": "The Drunken Replier",
                                                            "normal_followers_count": 348,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1678142148162584576/1688936091",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1678145667561140225/BOqYLVqw_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "drunken126",
                                                            "statuses_count": 4759,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1214785486839066624",
                                        "sortIndex": "1874477859959996399",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMjE0Nzg1NDg2ODM5MDY2NjI0",
                                                        "rest_id": "1214785486839066624",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Wed Jan 08 05:47:15 +0000 2020",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Jesus follower, Mother, American, conservative, in that order.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 27178,
                                                            "followers_count": 1228,
                                                            "friends_count": 1606,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Houston, TX",
                                                            "media_count": 618,
                                                            "name": "MsTraLaLa",
                                                            "normal_followers_count": 1228,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1452356743300726785/uaptmRvu_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "MsTraLaLa1",
                                                            "statuses_count": 10064,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1711040777499856896",
                                        "sortIndex": "1874477859959996398",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNzExMDQwNzc3NDk5ODU2ODk2",
                                                        "rest_id": "1711040777499856896",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Sun Oct 08 15:28:24 +0000 2023",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Question everything! Be skeptical! Be strong! Be amazing! #UNITY",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 69346,
                                                            "followers_count": 331,
                                                            "friends_count": 372,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 0,
                                                            "location": "Michigan, USA",
                                                            "media_count": 116,
                                                            "name": "LA",
                                                            "normal_followers_count": 331,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1711040777499856896/1730769508",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1740803767820324864/xMcnKxsh_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "SkepticJaneGalt",
                                                            "statuses_count": 4123,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-4551763274",
                                        "sortIndex": "1874477859959996397",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0NTUxNzYzMjc0",
                                                        "rest_id": "4551763274",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": false,
                                                            "can_dm": false,
                                                            "can_media_tag": true,
                                                            "created_at": "Mon Dec 14 00:29:43 +0000 2015",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Coffee ☕️ Fitness, Muscle cars and America First \uD83C\uDDFA\uD83C\uDDF8Gen Xer",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 15546,
                                                            "followers_count": 772,
                                                            "friends_count": 1243,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 6,
                                                            "location": "New Jersey, USA",
                                                            "media_count": 2263,
                                                            "name": "MichaelJohn",
                                                            "normal_followers_count": 772,
                                                            "pinned_tweet_ids_str": [
                                                                "1753274491474510307"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/4551763274/1455379531",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1750271418934108160/8KLexIs6_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "MoparMike6T8",
                                                            "statuses_count": 21632,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": false,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "BlueVerifiedFollowersSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "cursor-bottom-1874477859959996396",
                                        "sortIndex": "1874477859959996396",
                                        "content": {
                                            "entryType": "TimelineTimelineCursor",
                                            "__typename": "TimelineTimelineCursor",
                                            "value": "1820058766362429878|1874477859959996394",
                                            "cursorType": "Bottom"
                                        }
                                    },
                                    {
                                        "entryId": "cursor-top-1874477859959996417",
                                        "sortIndex": "1874477859959996417",
                                        "content": {
                                            "entryType": "TimelineTimelineCursor",
                                            "__typename": "TimelineTimelineCursor",
                                            "value": "-1|1874477859959996417",
                                            "cursorType": "Top"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                }
            }
        }
    }
}
```

#### Followers Your Know

###### Request

```javascript 
fetch("https://x.com/i/api/graphql/qJuLtV192xrB8Wftv6eXFw/FollowersYouKnow?variables=%7B%22userId%22%3A%221222773302441148416%22%2C%22count%22%3A20%2C%22includePromotedContent%22%3Afalse%7D&features=%7B%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22premium_content_api_read_enabled%22%3Afalse%2C%22communities_web_enable_tweet_community_results_fetch%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_button_fetch_trends_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_post_followups_enabled%22%3Afalse%2C%22responsive_web_grok_share_attachment_enabled%22%3Atrue%2C%22articles_preview_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22creator_subscriptions_quote_tweet_preview_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "5Ua3NS4W9iHxxj6E9IBnlebiTnpFNlUn0Df31/VrqaUTxXvh5SydVwnuD5LEEnFDGPYZwOaGKHGDi9H5Hn8mM6OY2mRi5g",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/visegrad24/followers_you_follow",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

###### Payload

```javascript
variables: {"userId":"1222773302441148416","count":20,"includePromotedContent":false}
features: {"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"premium_content_api_read_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"responsive_web_grok_analyze_button_fetch_trends_enabled":true,"responsive_web_grok_analyze_post_followups_enabled":false,"responsive_web_grok_share_attachment_enabled":true,"articles_preview_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"rweb_video_timestamps_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"responsive_web_enhance_cards_enabled":false}
```

###### Response

```javascript
{
    "data": {
        "user": {
            "result": {
                "__typename": "User",
                "timeline": {
                    "timeline": {
                        "instructions": [
                            {
                                "type": "TimelineClearCache"
                            },
                            {
                                "type": "TimelineTerminateTimeline",
                                "direction": "Top"
                            },
                            {
                                "type": "TimelineTerminateTimeline",
                                "direction": "Bottom"
                            },
                            {
                                "type": "TimelineAddEntries",
                                "entries": [
                                    {
                                        "entryId": "user-1485689970593468416",
                                        "sortIndex": "1874478233513099264",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNDg1Njg5OTcwNTkzNDY4NDE2",
                                                        "rest_id": "1485689970593468416",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Mon Jan 24 19:05:23 +0000 2022",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Navigating the discourse",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "therabbithole84.substack.com",
                                                                            "expanded_url": "https://therabbithole84.substack.com/",
                                                                            "url": "https://t.co/Y5aSDpz7f9",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 72541,
                                                            "followers_count": 939815,
                                                            "friends_count": 1991,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 2676,
                                                            "location": "Mars, PA",
                                                            "media_count": 8653,
                                                            "name": "The Rabbit Hole",
                                                            "normal_followers_count": 939815,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1485689970593468416/1651711479",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1522017422550528001/6AceRKJQ_normal.png",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "TheRabbitHole84",
                                                            "statuses_count": 53323,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/Y5aSDpz7f9",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1519074240850395136",
                                                            "professional_type": "Creator",
                                                            "category": [
                                                                {
                                                                    "id": 1009,
                                                                    "name": "Community",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": false,
                                                            "cash_app_handle": "TheRabbitHole84"
                                                        },
                                                        "super_follow_eligible": true
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1111963348252397574",
                                        "sortIndex": "1874478233513099263",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMTExOTYzMzQ4MjUyMzk3NTc0",
                                                        "rest_id": "1111963348252397574",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Sat Mar 30 12:08:16 +0000 2019",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "cybersecurity/ ethical philosophy and more/ Likes means I've read\n\nhttps://t.co/ypk8kLJyFA",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "m.cmx.im/@machan",
                                                                            "expanded_url": "https://m.cmx.im/@machan",
                                                                            "url": "https://t.co/ypk8kLJyFA",
                                                                            "indices": [
                                                                                67,
                                                                                90
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 52534,
                                                            "followers_count": 2534,
                                                            "friends_count": 602,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 19,
                                                            "location": "127.0.0.1",
                                                            "media_count": 403,
                                                            "name": "甜味麻酱\uD83E\uDD88",
                                                            "normal_followers_count": 2534,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1111963348252397574/1688890580",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1503218322665619464/kP3337Ma_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "machan_47291",
                                                            "statuses_count": 8781,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-3474685637",
                                        "sortIndex": "1874478233513099262",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjozNDc0Njg1NjM3",
                                                        "rest_id": "3474685637",
                                                        "affiliates_highlighted_label": {
                                                            "label": {
                                                                "url": {
                                                                    "url": "https://twitter.com/Polymarket",
                                                                    "urlType": "DeepLink"
                                                                },
                                                                "badge": {
                                                                    "url": "https://pbs.twimg.com/profile_images/1789484330638561280/t4Js9UFO_bigger.jpg"
                                                                },
                                                                "description": "Polymarket",
                                                                "userLabelType": "BusinessLabel",
                                                                "userLabelDisplayType": "Badge"
                                                            }
                                                        },
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Sat Aug 29 00:04:15 +0000 2015",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "CEO @Polymarket. Ethereum since ’14. I love music and collect art.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "polymarket.com",
                                                                            "expanded_url": "http://polymarket.com",
                                                                            "url": "https://t.co/PG4WREYHgj",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 31943,
                                                            "followers_count": 85748,
                                                            "friends_count": 1118,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 558,
                                                            "location": "new york",
                                                            "media_count": 41,
                                                            "name": "Shayne Coplan \uD83E\uDD85",
                                                            "normal_followers_count": 85748,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/3474685637/1726884547",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1458132114105413633/youFcLMv_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "shayne_coplan",
                                                            "statuses_count": 1022,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/PG4WREYHgj",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1891490382",
                                        "sortIndex": "1874478233513099261",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxODkxNDkwMzgy",
                                                        "rest_id": "1891490382",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Sep 21 21:24:04 +0000 2013",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "The largest Eastern European media. To let the world know.\n\nOur website: https://t.co/4voPija7tJ\nBecome our patron: https://t.co/fOGkarNq2c",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "en.nexta.tv",
                                                                            "expanded_url": "https://en.nexta.tv",
                                                                            "url": "https://t.co/4voPija7tJ",
                                                                            "indices": [
                                                                                73,
                                                                                96
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "patreon.com/nexta_tv",
                                                                            "expanded_url": "https://patreon.com/nexta_tv",
                                                                            "url": "https://t.co/fOGkarNq2c",
                                                                            "indices": [
                                                                                116,
                                                                                139
                                                                            ]
                                                                        }
                                                                    ]
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "nexta.tv/support",
                                                                            "expanded_url": "https://nexta.tv/support",
                                                                            "url": "https://t.co/AfUDiu1lc0",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 9779,
                                                            "followers_count": 1240636,
                                                            "friends_count": 129,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 10730,
                                                            "location": "",
                                                            "media_count": 41802,
                                                            "name": "NEXTA",
                                                            "normal_followers_count": 1240636,
                                                            "pinned_tweet_ids_str": [
                                                                "1874138454516174871"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1891490382/1705316345",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1514923008707768321/hL4PN7Hf_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "nexta_tv",
                                                            "statuses_count": 47319,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/AfUDiu1lc0",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1459506202052866056",
                                                            "professional_type": "Business",
                                                            "category": [
                                                                {
                                                                    "id": 579,
                                                                    "name": "Media & News",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": true,
                                                            "bitcoin_handle": "1NFA7bjyp3UNyjeMYeXa9LmcsBFpbVQiUu",
                                                            "ethereum_handle": "",
                                                            "patreon_handle": "nexta_tv"
                                                        },
                                                        "super_follow_eligible": true
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1577241403",
                                        "sortIndex": "1874478233513099260",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxNTc3MjQxNDAz",
                                                        "rest_id": "1577241403",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": false,
                                                            "created_at": "Mon Jul 08 09:31:59 +0000 2013",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "\uD83C\uDDEA\uD83C\uDDFAhttps://t.co/NdorAWqJC3\n\uD83D\uDCBEhttps://t.co/T74ZwJ1F0C est 1994\n\uD83D\uDCF8https://t.co/lAyoqmSBRX $125K/m\n\uD83C\uDFE1https://t.co/1oqUgfD6CZ $31K/m\n\uD83D\uDEF0https://t.co/ZHSvI2wjyW $28K/m\n\uD83C\uDF0Dhttps://t.co/edCiRYU5yB $28K/m\n\uD83D\uDC59https://t.co/RyXpqGuFM3 $25K/m",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "euacc.com",
                                                                            "expanded_url": "http://euacc.com",
                                                                            "url": "https://t.co/NdorAWqJC3",
                                                                            "indices": [
                                                                                2,
                                                                                25
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "Pieter.com",
                                                                            "expanded_url": "http://Pieter.com",
                                                                            "url": "https://t.co/T74ZwJ1F0C",
                                                                            "indices": [
                                                                                27,
                                                                                50
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "PhotoAI.com",
                                                                            "expanded_url": "http://PhotoAI.com",
                                                                            "url": "https://t.co/lAyoqmSBRX",
                                                                            "indices": [
                                                                                61,
                                                                                84
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "InteriorAI.com",
                                                                            "expanded_url": "http://InteriorAI.com",
                                                                            "url": "https://t.co/1oqUgfD6CZ",
                                                                            "indices": [
                                                                                94,
                                                                                117
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "RemoteOK.com",
                                                                            "expanded_url": "http://RemoteOK.com",
                                                                            "url": "https://t.co/ZHSvI2wjyW",
                                                                            "indices": [
                                                                                126,
                                                                                149
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "nomads.com/@levelsio",
                                                                            "expanded_url": "http://nomads.com/@levelsio",
                                                                            "url": "https://t.co/edCiRYU5yB",
                                                                            "indices": [
                                                                                158,
                                                                                181
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "levelsio.com",
                                                                            "expanded_url": "http://levelsio.com",
                                                                            "url": "https://t.co/RyXpqGuFM3",
                                                                            "indices": [
                                                                                190,
                                                                                213
                                                                            ]
                                                                        }
                                                                    ]
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "readMAKE.com",
                                                                            "expanded_url": "http://readMAKE.com",
                                                                            "url": "https://t.co/IFRqZDK2Lw",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 122131,
                                                            "followers_count": 595524,
                                                            "friends_count": 1654,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 6505,
                                                            "location": "\uD83D\uDCD5 My book $20K/m \uD83D\uDC49",
                                                            "media_count": 15525,
                                                            "name": "@levelsio",
                                                            "normal_followers_count": 595524,
                                                            "pinned_tweet_ids_str": [
                                                                "1855359096444334315"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1577241403/1665964220",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1589756412078555136/YlXMBzhp_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "levelsio",
                                                            "statuses_count": 143871,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/IFRqZDK2Lw",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1800203633688027410",
                                                            "professional_type": "Creator",
                                                            "category": []
                                                        },
                                                        "tipjar_settings": {},
                                                        "super_follow_eligible": true
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-1180533698",
                                        "sortIndex": "1874478233513099259",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjoxMTgwNTMzNjk4",
                                                        "rest_id": "1180533698",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Thu Feb 14 22:56:27 +0000 2013",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Posting interesting science, gadgets, history, art, and more. Subscribe for in-depth posts. As an Amazon Associate I earn from qualifying purchases.",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "linktr.ee/fascinating1",
                                                                            "expanded_url": "https://linktr.ee/fascinating1",
                                                                            "url": "https://t.co/niO5oqBIRW",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 621,
                                                            "followers_count": 3281546,
                                                            "friends_count": 1824,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 9997,
                                                            "location": "United States",
                                                            "media_count": 12995,
                                                            "name": "Fascinating",
                                                            "normal_followers_count": 3281546,
                                                            "pinned_tweet_ids_str": [
                                                                "1870073777402356024"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/1180533698/1664386546",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1552059284086611968/l0vA92Sh_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "fasc1nate",
                                                            "statuses_count": 28268,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/niO5oqBIRW",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1599088121886277632",
                                                            "professional_type": "Business",
                                                            "category": []
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": true
                                                        },
                                                        "super_follow_eligible": true
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-630037008",
                                        "sortIndex": "1874478233513099258",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo2MzAwMzcwMDg=",
                                                        "rest_id": "630037008",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": false,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Sun Jul 08 06:16:17 +0000 2012",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "Father, husband, AI engineer and AI investor at Intel Capital.\nThe opinions expressed are mine alone.\n#IamIntel",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 823,
                                                            "followers_count": 306,
                                                            "friends_count": 1686,
                                                            "has_custom_timelines": false,
                                                            "is_translator": false,
                                                            "listed_count": 9,
                                                            "location": "Palo Alto, CA",
                                                            "media_count": 28,
                                                            "name": "Assaf Araki",
                                                            "normal_followers_count": 306,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/630037008/1666895658",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1405193322600157187/hjf8WirI_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "BigData_ML",
                                                            "statuses_count": 673,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-495854424",
                                        "sortIndex": "1874478233513099257",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0OTU4NTQ0MjQ=",
                                                        "rest_id": "495854424",
                                                        "affiliates_highlighted_label": {},
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": true,
                                                            "can_media_tag": true,
                                                            "created_at": "Sat Feb 18 11:24:22 +0000 2012",
                                                            "default_profile": true,
                                                            "default_profile_image": false,
                                                            "description": "加拿大持牌移民顾问｜微信 frankyuyong2｜油管频道 https://t.co/Fo56ETKbf6｜公司网站 https://t.co/ewRozHGlhP ｜咨询预约 https://t.co/sb8k4o5WDk｜推文不构成法律或专业意见",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "m.youtube.com/c/decisionmade",
                                                                            "expanded_url": "http://m.youtube.com/c/decisionmade",
                                                                            "url": "https://t.co/Fo56ETKbf6",
                                                                            "indices": [
                                                                                31,
                                                                                54
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "cedarhill.co",
                                                                            "expanded_url": "http://cedarhill.co",
                                                                            "url": "https://t.co/ewRozHGlhP",
                                                                            "indices": [
                                                                                60,
                                                                                83
                                                                            ]
                                                                        },
                                                                        {
                                                                            "display_url": "cedarhill.co/consulting/",
                                                                            "expanded_url": "http://cedarhill.co/consulting/",
                                                                            "url": "https://t.co/sb8k4o5WDk",
                                                                            "indices": [
                                                                                90,
                                                                                113
                                                                            ]
                                                                        }
                                                                    ]
                                                                },
                                                                "url": {
                                                                    "urls": [
                                                                        {
                                                                            "display_url": "cedarhill.io",
                                                                            "expanded_url": "https://www.cedarhill.io",
                                                                            "url": "https://t.co/zJtCjbEogE",
                                                                            "indices": [
                                                                                0,
                                                                                23
                                                                            ]
                                                                        }
                                                                    ]
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 5477,
                                                            "followers_count": 21306,
                                                            "friends_count": 642,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 131,
                                                            "location": "加拿大 渥太华",
                                                            "media_count": 1053,
                                                            "name": "yu胖",
                                                            "normal_followers_count": 21306,
                                                            "pinned_tweet_ids_str": [
                                                                "1325969351409807360"
                                                            ],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/495854424/1654079612",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1252568434224893953/WXvrfJRp_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "frankyuyong",
                                                            "statuses_count": 5449,
                                                            "translator_type": "none",
                                                            "url": "https://t.co/zJtCjbEogE",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1454564615216533508",
                                                            "professional_type": "Business",
                                                            "category": [
                                                                {
                                                                    "id": 477,
                                                                    "name": "Professional Services",
                                                                    "icon_name": "IconBriefcaseStroke"
                                                                }
                                                            ]
                                                        },
                                                        "tipjar_settings": {}
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "user-44196397",
                                        "sortIndex": "1874478233513099256",
                                        "content": {
                                            "entryType": "TimelineTimelineItem",
                                            "__typename": "TimelineTimelineItem",
                                            "itemContent": {
                                                "itemType": "TimelineUser",
                                                "__typename": "TimelineUser",
                                                "user_results": {
                                                    "result": {
                                                        "__typename": "User",
                                                        "id": "VXNlcjo0NDE5NjM5Nw==",
                                                        "rest_id": "44196397",
                                                        "affiliates_highlighted_label": {
                                                            "label": {
                                                                "url": {
                                                                    "url": "https://twitter.com/X",
                                                                    "urlType": "DeepLink"
                                                                },
                                                                "badge": {
                                                                    "url": "https://pbs.twimg.com/profile_images/1683899100922511378/5lY42eHs_bigger.jpg"
                                                                },
                                                                "description": "X",
                                                                "userLabelType": "BusinessLabel",
                                                                "userLabelDisplayType": "Badge"
                                                            }
                                                        },
                                                        "has_graduated_access": true,
                                                        "is_blue_verified": true,
                                                        "profile_image_shape": "Circle",
                                                        "legacy": {
                                                            "following": true,
                                                            "can_dm": false,
                                                            "can_media_tag": false,
                                                            "created_at": "Tue Jun 02 20:12:29 +0000 2009",
                                                            "default_profile": false,
                                                            "default_profile_image": false,
                                                            "description": "⚔️",
                                                            "entities": {
                                                                "description": {
                                                                    "urls": []
                                                                }
                                                            },
                                                            "fast_followers_count": 0,
                                                            "favourites_count": 107097,
                                                            "followers_count": 209978531,
                                                            "friends_count": 911,
                                                            "has_custom_timelines": true,
                                                            "is_translator": false,
                                                            "listed_count": 155989,
                                                            "location": "",
                                                            "media_count": 3023,
                                                            "name": "Kekius Maximus",
                                                            "normal_followers_count": 209978531,
                                                            "pinned_tweet_ids_str": [],
                                                            "possibly_sensitive": false,
                                                            "profile_banner_url": "https://pbs.twimg.com/profile_banners/44196397/1726163678",
                                                            "profile_image_url_https": "https://pbs.twimg.com/profile_images/1873974185766060032/gRZc3N-3_normal.jpg",
                                                            "profile_interstitial_type": "",
                                                            "screen_name": "elonmusk",
                                                            "statuses_count": 64699,
                                                            "translator_type": "none",
                                                            "verified": false,
                                                            "want_retweets": true,
                                                            "withheld_in_countries": []
                                                        },
                                                        "professional": {
                                                            "rest_id": "1679729435447275522",
                                                            "professional_type": "Creator",
                                                            "category": []
                                                        },
                                                        "tipjar_settings": {
                                                            "is_enabled": false
                                                        },
                                                        "super_follow_eligible": true
                                                    }
                                                },
                                                "userDisplayType": "User"
                                            },
                                            "clientEventInfo": {
                                                "component": "FriendsFollowingSgs",
                                                "element": "user"
                                            }
                                        }
                                    },
                                    {
                                        "entryId": "cursor-bottom-1874478233513099255",
                                        "sortIndex": "1874478233513099255",
                                        "content": {
                                            "entryType": "TimelineTimelineCursor",
                                            "__typename": "TimelineTimelineCursor",
                                            "value": "0|1874478233513099253",
                                            "cursorType": "Bottom"
                                        }
                                    },
                                    {
                                        "entryId": "cursor-top-1874478233513099265",
                                        "sortIndex": "1874478233513099265",
                                        "content": {
                                            "entryType": "TimelineTimelineCursor",
                                            "__typename": "TimelineTimelineCursor",
                                            "value": "-1|1874478233513099265",
                                            "cursorType": "Top"
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                }
            }
        }
    }
}
```

#### ProfileSpotlightsQuery

###### Request

```javascript
fetch("https://x.com/i/api/graphql/-0XdHI-mrHWBQd8-oLo1aA/ProfileSpotlightsQuery?variables=%7B%22screen_name%22%3A%22visegrad24%22%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "/URh+N+V+5tK21KjC3bBPcYkavDWJZXhY8barFFFoXM9FVEabl7xetQhhILUBwfoUf/n2/79Y35khwuF/0GLncN9BX9m/g",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/visegrad24",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

##### Playload

``` javascript
variables: {"screen_name":"visegrad24"}
```

###### Response

```javascript
{
    "data": {
        "user_result_by_screen_name": {
            "result": {
                "__typename": "User",
                "legacy": {
                    "blocking": false,
                    "blocked_by": false,
                    "protected": false,
                    "following": true,
                    "followed_by": false,
                    "name": "Visegrád 24",
                    "screen_name": "visegrad24"
                },
                "rest_id": "1222773302441148416",
                "is_verified_organization": false,
                "profilemodules": {
                    "v1": []
                },
                "id": "VXNlcjoxMjIyNzczMzAyNDQxMTQ4NDE2"
            },
            "id": "VXNlclJlc3VsdHM6MTIyMjc3MzMwMjQ0MTE0ODQxNg=="
        }
    }
}
```

#### TweetResultByRestID

###### Request

```javascript
fetch("https://x.com/i/api/graphql/zptD3jqLJ0arTQdM10uc3w/TweetResultByRestId?variables=%7B%22tweetId%22%3A%221626275168157851650%22%2C%22includePromotedContent%22%3Atrue%2C%22withBirdwatchNotes%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withCommunity%22%3Atrue%7D&features=%7B%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22premium_content_api_read_enabled%22%3Afalse%2C%22communities_web_enable_tweet_community_results_fetch%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_button_fetch_trends_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_post_followups_enabled%22%3Afalse%2C%22responsive_web_grok_share_attachment_enabled%22%3Atrue%2C%22articles_preview_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22creator_subscriptions_quote_tweet_preview_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D&fieldToggles=%7B%22withArticleRichContentState%22%3Atrue%2C%22withArticlePlainText%22%3Afalse%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "3ZqFN9AYrO+76UUDqHt7dOusUePQ8ztffaSR5WPIFOQi8E2a+7SWuIIlqrVig1K/17r/+94QYmBWCe2Rfo/OYqIm5Mts3g",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/justinwood_/status/1626275168157851650/media_tags",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

###### Payload

```javascript
variables: {"tweetId":"1626275168157851650","includePromotedContent":true,"withBirdwatchNotes":true,"withVoice":true,"withCommunity":true}
features: {"creator_subscriptions_tweet_preview_api_enabled":true,"premium_content_api_read_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"responsive_web_grok_analyze_button_fetch_trends_enabled":true,"responsive_web_grok_analyze_post_followups_enabled":false,"responsive_web_grok_share_attachment_enabled":true,"articles_preview_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"rweb_video_timestamps_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_enhance_cards_enabled":false}
fieldToggles: {"withArticleRichContentState":true,"withArticlePlainText":false}
```

###### Response
```javascript
{
    "data": {
        "tweetResult": {
            "result": {
                "__typename": "Tweet",
                "rest_id": "1626275168157851650",
                "has_birdwatch_notes": false,
                "core": {
                    "user_results": {
                        "result": {
                            "__typename": "User",
                            "id": "VXNlcjoxMTI3NzM2MTYzMTMxODk5OTA2",
                            "rest_id": "1127736163131899906",
                            "affiliates_highlighted_label": {},
                            "has_graduated_access": true,
                            "is_blue_verified": false,
                            "profile_image_shape": "Circle",
                            "legacy": {
                                "following": false,
                                "can_dm": true,
                                "can_media_tag": true,
                                "created_at": "Mon May 13 00:43:48 +0000 2019",
                                "default_profile": true,
                                "default_profile_image": false,
                                "description": "Sophomore LongSnapper at The University Of Houston #47 • 3.9 gpa • 2023-24 ACC Academic Honor Roll • Kornblue Snapping #fab50",
                                "entities": {
                                    "description": {
                                        "urls": []
                                    }
                                },
                                "fast_followers_count": 0,
                                "favourites_count": 516,
                                "followers_count": 518,
                                "friends_count": 490,
                                "has_custom_timelines": false,
                                "is_translator": false,
                                "listed_count": 3,
                                "location": "",
                                "media_count": 55,
                                "name": "jwood",
                                "normal_followers_count": 518,
                                "pinned_tweet_ids_str": [],
                                "possibly_sensitive": false,
                                "profile_banner_url": "https://pbs.twimg.com/profile_banners/1127736163131899906/1714424193",
                                "profile_image_url_https": "https://pbs.twimg.com/profile_images/1779498147883970560/GxYJ7ZMi_normal.jpg",
                                "profile_interstitial_type": "",
                                "screen_name": "justinwood_",
                                "statuses_count": 265,
                                "translator_type": "none",
                                "verified": false,
                                "want_retweets": false,
                                "withheld_in_countries": []
                            },
                            "tipjar_settings": {}
                        }
                    }
                },
                "unmention_data": {},
                "edit_control": {
                    "edit_tweet_ids": [
                        "1626275168157851650"
                    ],
                    "editable_until_msecs": "1676570989000",
                    "is_edit_eligible": true,
                    "edits_remaining": "5"
                },
                "is_translatable": false,
                "views": {
                    "count": "39525",
                    "state": "EnabledWithCount"
                },
                "source": "<a href=\"http://twitter.com/download/iphone\" rel=\"nofollow\">Twitter for iPhone</a>",
                "legacy": {
                    "bookmark_count": 2,
                    "bookmarked": false,
                    "created_at": "Thu Feb 16 17:39:49 +0000 2023",
                    "conversation_id_str": "1626275168157851650",
                    "display_text_range": [
                        0,
                        37
                    ],
                    "entities": {
                        "hashtags": [
                            {
                                "indices": [
                                    27,
                                    37
                                ],
                                "text": "committed"
                            }
                        ],
                        "media": [
                            {
                                "display_url": "pic.x.com/tW4I4w7P4a",
                                "expanded_url": "https://x.com/justinwood_/status/1626275168157851650/photo/1",
                                "id_str": "1626275160864043008",
                                "indices": [
                                    38,
                                    61
                                ],
                                "media_key": "3_1626275160864043008",
                                "media_url_https": "https://pbs.twimg.com/media/FpGwpHZXsAASRI8.jpg",
                                "type": "photo",
                                "url": "https://t.co/tW4I4w7P4a",
                                "ext_media_availability": {
                                    "status": "Available"
                                },
                                "features": {
                                    "all": {
                                        "tags": [
                                            {
                                                "user_id": "764347046",
                                                "name": "Miami Hurricanes Football",
                                                "screen_name": "CanesFootball",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "1504640117683761154",
                                                "name": "Marwan Maalouf",
                                                "screen_name": "CoachMaalouf",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "496368068",
                                                "name": "Danny Kalter",
                                                "screen_name": "CoachKalter",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "1375007264",
                                                "name": "Mario Cristobal",
                                                "screen_name": "coach_cristobal",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "1139663945860562944",
                                                "name": "Frank Tucker",
                                                "screen_name": "TheCribSouthFLA",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "3063335273",
                                                "name": "Geo Milian",
                                                "screen_name": "GeoMilian",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "4488634109",
                                                "name": "Ferras Isa \uD83C\uDDF1\uD83C\uDDFE",
                                                "screen_name": "CoachFerras",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "2204663483",
                                                "name": "Gaby Urrutia",
                                                "screen_name": "GabyUrrutia247",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "2381021022",
                                                "name": "Miami \uD83C\uDD7F️almetto Football",
                                                "screen_name": "PantherPride305",
                                                "type": "user"
                                            }
                                        ]
                                    },
                                    "large": {
                                        "faces": [
                                            {
                                                "x": 621,
                                                "y": 873,
                                                "h": 48,
                                                "w": 48
                                            },
                                            {
                                                "x": 420,
                                                "y": 580,
                                                "h": 59,
                                                "w": 59
                                            }
                                        ]
                                    },
                                    "medium": {
                                        "faces": [
                                            {
                                                "x": 600,
                                                "y": 843,
                                                "h": 46,
                                                "w": 46
                                            },
                                            {
                                                "x": 405,
                                                "y": 560,
                                                "h": 57,
                                                "w": 57
                                            }
                                        ]
                                    },
                                    "small": {
                                        "faces": [
                                            {
                                                "x": 339,
                                                "y": 477,
                                                "h": 26,
                                                "w": 26
                                            },
                                            {
                                                "x": 229,
                                                "y": 317,
                                                "h": 32,
                                                "w": 32
                                            }
                                        ]
                                    },
                                    "orig": {
                                        "faces": [
                                            {
                                                "x": 621,
                                                "y": 873,
                                                "h": 48,
                                                "w": 48
                                            },
                                            {
                                                "x": 420,
                                                "y": 580,
                                                "h": 59,
                                                "w": 59
                                            }
                                        ]
                                    }
                                },
                                "sizes": {
                                    "large": {
                                        "h": 1242,
                                        "w": 828,
                                        "resize": "fit"
                                    },
                                    "medium": {
                                        "h": 1200,
                                        "w": 800,
                                        "resize": "fit"
                                    },
                                    "small": {
                                        "h": 680,
                                        "w": 453,
                                        "resize": "fit"
                                    },
                                    "thumb": {
                                        "h": 150,
                                        "w": 150,
                                        "resize": "crop"
                                    }
                                },
                                "original_info": {
                                    "height": 1242,
                                    "width": 828,
                                    "focus_rects": [
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 464
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 828
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 944
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 621,
                                            "h": 1242
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 1242
                                        }
                                    ]
                                },
                                "media_results": {
                                    "result": {
                                        "media_key": "3_1626275160864043008"
                                    }
                                }
                            }
                        ],
                        "symbols": [],
                        "timestamps": [],
                        "urls": [],
                        "user_mentions": []
                    },
                    "extended_entities": {
                        "media": [
                            {
                                "display_url": "pic.x.com/tW4I4w7P4a",
                                "expanded_url": "https://x.com/justinwood_/status/1626275168157851650/photo/1",
                                "id_str": "1626275160864043008",
                                "indices": [
                                    38,
                                    61
                                ],
                                "media_key": "3_1626275160864043008",
                                "media_url_https": "https://pbs.twimg.com/media/FpGwpHZXsAASRI8.jpg",
                                "type": "photo",
                                "url": "https://t.co/tW4I4w7P4a",
                                "ext_media_availability": {
                                    "status": "Available"
                                },
                                "features": {
                                    "all": {
                                        "tags": [
                                            {
                                                "user_id": "764347046",
                                                "name": "Miami Hurricanes Football",
                                                "screen_name": "CanesFootball",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "1504640117683761154",
                                                "name": "Marwan Maalouf",
                                                "screen_name": "CoachMaalouf",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "496368068",
                                                "name": "Danny Kalter",
                                                "screen_name": "CoachKalter",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "1375007264",
                                                "name": "Mario Cristobal",
                                                "screen_name": "coach_cristobal",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "1139663945860562944",
                                                "name": "Frank Tucker",
                                                "screen_name": "TheCribSouthFLA",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "3063335273",
                                                "name": "Geo Milian",
                                                "screen_name": "GeoMilian",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "4488634109",
                                                "name": "Ferras Isa \uD83C\uDDF1\uD83C\uDDFE",
                                                "screen_name": "CoachFerras",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "2204663483",
                                                "name": "Gaby Urrutia",
                                                "screen_name": "GabyUrrutia247",
                                                "type": "user"
                                            },
                                            {
                                                "user_id": "2381021022",
                                                "name": "Miami \uD83C\uDD7F️almetto Football",
                                                "screen_name": "PantherPride305",
                                                "type": "user"
                                            }
                                        ]
                                    },
                                    "large": {
                                        "faces": [
                                            {
                                                "x": 621,
                                                "y": 873,
                                                "h": 48,
                                                "w": 48
                                            },
                                            {
                                                "x": 420,
                                                "y": 580,
                                                "h": 59,
                                                "w": 59
                                            }
                                        ]
                                    },
                                    "medium": {
                                        "faces": [
                                            {
                                                "x": 600,
                                                "y": 843,
                                                "h": 46,
                                                "w": 46
                                            },
                                            {
                                                "x": 405,
                                                "y": 560,
                                                "h": 57,
                                                "w": 57
                                            }
                                        ]
                                    },
                                    "small": {
                                        "faces": [
                                            {
                                                "x": 339,
                                                "y": 477,
                                                "h": 26,
                                                "w": 26
                                            },
                                            {
                                                "x": 229,
                                                "y": 317,
                                                "h": 32,
                                                "w": 32
                                            }
                                        ]
                                    },
                                    "orig": {
                                        "faces": [
                                            {
                                                "x": 621,
                                                "y": 873,
                                                "h": 48,
                                                "w": 48
                                            },
                                            {
                                                "x": 420,
                                                "y": 580,
                                                "h": 59,
                                                "w": 59
                                            }
                                        ]
                                    }
                                },
                                "sizes": {
                                    "large": {
                                        "h": 1242,
                                        "w": 828,
                                        "resize": "fit"
                                    },
                                    "medium": {
                                        "h": 1200,
                                        "w": 800,
                                        "resize": "fit"
                                    },
                                    "small": {
                                        "h": 680,
                                        "w": 453,
                                        "resize": "fit"
                                    },
                                    "thumb": {
                                        "h": 150,
                                        "w": 150,
                                        "resize": "crop"
                                    }
                                },
                                "original_info": {
                                    "height": 1242,
                                    "width": 828,
                                    "focus_rects": [
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 464
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 828
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 944
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 621,
                                            "h": 1242
                                        },
                                        {
                                            "x": 0,
                                            "y": 0,
                                            "w": 828,
                                            "h": 1242
                                        }
                                    ]
                                },
                                "media_results": {
                                    "result": {
                                        "media_key": "3_1626275160864043008"
                                    }
                                }
                            }
                        ]
                    },
                    "favorite_count": 820,
                    "favorited": false,
                    "full_text": "staying home, let’s work…\uD83C\uDF34 #committed https://t.co/tW4I4w7P4a",
                    "is_quote_status": false,
                    "lang": "en",
                    "possibly_sensitive": false,
                    "possibly_sensitive_editable": true,
                    "quote_count": 6,
                    "reply_count": 19,
                    "retweet_count": 72,
                    "retweeted": false,
                    "user_id_str": "1127736163131899906",
                    "id_str": "1626275168157851650"
                }
            }
        }
    }
}
```

#### UserByRestID

###### Request
```javascript
fetch("https://x.com/i/api/graphql/LWxkCeL8Hlx0-f24DmPAJw/UserByRestId?variables=%7B%22userId%22%3A%221589569521614000128%22%7D&features=%7B%22hidden_profile_subscriptions_enabled%22%3Atrue%2C%22profile_label_improvements_pcf_label_in_post_enabled%22%3Afalse%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22highlights_tweets_tab_ui_enabled%22%3Atrue%2C%22responsive_web_twitter_article_notes_tab_enabled%22%3Atrue%2C%22subscriptions_feature_can_gift_premium%22%3Atrue%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%7D", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
    "authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
    "content-type": "application/json",
    "priority": "u=1, i",
    "sec-ch-ua": "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "x-client-transaction-id": "h2l/RBTZ4AShnB+hPecPMjcITHy/phd0xdAMw0yFLO+YFHfDV5gwqgwkD8sylY/ngYugoYS7kkBKAsI29+zXXjGCYBg9hA",
    "x-csrf-token": "6fcfe6808df05fe3c499d1c61d58a4c804759a195eae476a5ac1e74963be9a5a89d39177a035f4e37c87f1a2898c6a27f76c15eac427d13172c145fb0a385a358c89c5579731151f4250a90e1ca5046b",
    "x-twitter-active-user": "yes",
    "x-twitter-auth-type": "OAuth2Session",
    "x-twitter-client-language": "en",
    "Referer": "https://x.com/_KingsWorld/status/1874519052582003145",
    "Referrer-Policy": "strict-origin-when-cross-origin"
  },
  "body": null,
  "method": "GET"
});
```

###### Payload

```javascript
variables: {"userId":"1589569521614000128"}
features: {"hidden_profile_subscriptions_enabled":true,"profile_label_improvements_pcf_label_in_post_enabled":false,"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"highlights_tweets_tab_ui_enabled":true,"responsive_web_twitter_article_notes_tab_enabled":true,"subscriptions_feature_can_gift_premium":true,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"responsive_web_graphql_timeline_navigation_enabled":true}
```

###### Response

```javascript
{
    "data": {
        "user": {
            "result": {
                "__typename": "User",
                "id": "VXNlcjoxNTg5NTY5NTIxNjE0MDAwMTI4",
                "rest_id": "1589569521614000128",
                "affiliates_highlighted_label": {},
                "has_graduated_access": true,
                "is_blue_verified": true,
                "profile_image_shape": "Square",
                "legacy": {
                    "following": false,
                    "can_dm": false,
                    "can_media_tag": true,
                    "created_at": "Mon Nov 07 10:44:48 +0000 2022",
                    "default_profile": true,
                    "default_profile_image": false,
                    "description": "\uD83D\uDC51 La liga de los reyes. Vive la Kings World Cup Nations del 1 al 12 de enero \uD83C\uDF0D",
                    "entities": {
                        "description": {
                            "urls": []
                        },
                        "url": {
                            "urls": [
                                {
                                    "display_url": "kingsleague.pro",
                                    "expanded_url": "https://kingsleague.pro",
                                    "url": "https://t.co/7OM2dbI2u7",
                                    "indices": [
                                        0,
                                        23
                                    ]
                                }
                            ]
                        }
                    },
                    "fast_followers_count": 0,
                    "favourites_count": 2895,
                    "followers_count": 847595,
                    "friends_count": 352,
                    "has_custom_timelines": true,
                    "is_translator": false,
                    "listed_count": 532,
                    "location": "",
                    "media_count": 7264,
                    "name": "Kings League InfoJobs Spain",
                    "normal_followers_count": 847595,
                    "pinned_tweet_ids_str": [
                        "1874029552651165846"
                    ],
                    "possibly_sensitive": false,
                    "profile_banner_url": "https://pbs.twimg.com/profile_banners/1589569521614000128/1735232713",
                    "profile_image_url_https": "https://pbs.twimg.com/profile_images/1720391010016989184/iXw5ZC5a_normal.jpg",
                    "profile_interstitial_type": "",
                    "screen_name": "KingsLeague",
                    "statuses_count": 9050,
                    "translator_type": "none",
                    "url": "https://t.co/7OM2dbI2u7",
                    "verified": false,
                    "verified_type": "Business",
                    "want_retweets": false,
                    "withheld_in_countries": []
                },
                "professional": {
                    "rest_id": "1590275897206231040",
                    "professional_type": "Business",
                    "category": [
                        {
                            "id": 839,
                            "name": "Sports League",
                            "icon_name": "IconBriefcaseStroke"
                        }
                    ]
                },
                "tipjar_settings": {},
                "business_account": {
                    "affiliates_count": 8
                },
                "highlights_info": {
                    "can_highlight_tweets": true,
                    "highlighted_tweets": "13"
                },
                "user_seed_tweet_count": 0,
                "premium_gifting_eligible": false,
                "creator_subscriptions_count": 0,
                "has_hidden_subscriptions_on_profile": false
            }
        }
    }
}
```
