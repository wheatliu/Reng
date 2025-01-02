package twitter

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cast"
)

func ParseUserTweets(userTweetResp *UserTweetsResponse) (tweets *UserTweets, err error) {
	tweets = &UserTweets{}
	if userTweetResp.StatusCode == 429 {
		return nil, ErrTooManyRequests
	}
	if userTweetResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status code %d", userTweetResp.StatusCode)
	}

	if userTweetResp.Data.Data.User == nil || userTweetResp.Data.Data.User.Result == nil {
		return nil, ErrUserPostsProtected
	}

	if userTweetResp.Data.Data.User.Result.TypeName == "UserUnavailable" {
		return nil, ErrUserSuspended
	}

	timeLineV2 := userTweetResp.Data.Data.User.Result.TimelineV2
	if timeLineV2 == nil {
		return nil, ErrInvalidTimeLineData
	}
	timeLine := timeLineV2.Timeline
	if timeLine == nil {
		return nil, ErrInvalidTimeLineData
	}

	instructions := timeLine.Instructions
	if len(instructions) == 0 {
		return nil, ErrInvalidInstruction
	}
	entries := []*Entry{}
	for _, instruction := range instructions {
		if instruction.Type == "TimelineAddEntries" {
			entries = instruction.Entries
			break
		}
	}
	if len(entries) == 0 {
		return nil, ErrInvalidTimeLineEntries
	}

	for _, entry := range entries {
		tweets.Totals = tweets.Totals + 1
		entryID := entry.EntryID
		isNormalTweet := strings.HasPrefix(entryID, NormalTweetPrefix)
		isConversation := strings.HasPrefix(entryID, ConversationTweetPrefix)
		hasNextPage := strings.HasPrefix(entryID, CursorPrefix)
		if isNormalTweet || isConversation {
			itemContent := entry.Content.ItemContent
			if itemContent != nil {
				result := entry.Content.ItemContent.TweetResults.Result
				if result != nil && result.Legacy != nil {
					createdAt, err := time.Parse(time.RubyDate, result.Legacy.CreatedAt)
					if err != nil {
						continue
					}
					isRepost := result.Legacy.RetweetedStatusResult != nil
					isQuote := result.QuotedStatusResult != nil
					tweetType := TweetTypeNormal
					referenceRestID := ""
					if isRepost {
						tweetType = TweetTypeRetweet
						if result.Legacy.RetweetedStatusResult.Result == nil {
							continue
						}
						referenceRestID = result.Legacy.RetweetedStatusResult.Result.RestID
					} else if isQuote {
						tweetType = TweetTypeQuote
						if result.QuotedStatusResult.Result == nil {
							continue
						}
						referenceRestID = result.QuotedStatusResult.Result.RestID
					} else if isConversation {
						tweetType = TweetTypeReply
					}
					tweet := &Tweet{
						EntryID:         entryID,
						RestID:          result.RestID,
						FullText:        result.Legacy.FullText,
						ReplyCount:      result.Legacy.ReplyCount,
						RetweetCount:    result.Legacy.RetweetCount,
						FavoriteCount:   result.Legacy.FavoriteCount,
						QuoteCount:      result.Legacy.QuoteCount,
						ViewCount:       cast.ToInt64(result.Views.Count),
						CreatedAt:       createdAt,
						Type:            tweetType,
						ReferenceRestID: referenceRestID,
					}
					hashTags := result.Legacy.Entities.HashTags
					if len(hashTags) > 0 {
						for _, hashTag := range hashTags {
							tweet.HashTags = append(tweet.HashTags, hashTag.Text)
						}
					}
					tweets.Tweets = append(tweets.Tweets, tweet)
				}
			}
		}
		if hasNextPage {
			tweets.Cursor = entry.Content.Value
		}
	}
	return tweets, nil
}

func ParseUserFollowingList(followingList *FollowListResp) ([]string, error) {
	if followingList.StatusCode == 429 {
		return nil, ErrTooManyRequests
	}
	// if followingList.StatusCode != 200 {
	// 	return nil, fmt.Errorf("request failed with status code %d", followingList.StatusCode)
	// }
	if followingList == nil {
		return nil, ErrInvalidTweetEntry
	}
	if followingList.Data.User == nil || followingList.Data.User.Result == nil {
		return nil, ErrUserPostsProtected
	}

	timeLine := followingList.Data.User.Result.Timeline
	if timeLine == nil {
		return nil, ErrInvalidTimeLineData
	}
	subTimeLine := timeLine.Timeline
	if subTimeLine == nil {
		return nil, ErrInvalidTimeLineData
	}

	instructions := subTimeLine.Instructions
	if len(instructions) == 0 {
		return nil, ErrInvalidInstruction
	}

	if len(instructions) == 0 {
		return nil, ErrInvalidTweetEntry
	}
	var entries []*FollowListEntry
	for _, instruction := range instructions {
		if instruction.Type == "TimelineAddEntries" {
			entries = instruction.Entries
			break
		}
	}

	following := []string{}
	for _, entry := range entries {
		itemContent := entry.Content.ItemContent
		if itemContent != nil {
			result := itemContent.UserResults.Result
			if result != nil {
				following = append(following, result.Legacy.ScreenName)
			}
		}
	}

	return following, nil
}
