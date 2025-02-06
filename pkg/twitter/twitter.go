package twitter

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/go-resty/resty/v2"
)

type TweetType string

const (
	// Twitter graphql API URL
	TweetDetailURL      = "https://x.com/i/api/graphql/U0HTv-bAWTBYylwEMT7x5A/TweetDetail"
	UserTweetsURL       = "https://x.com/i/api/graphql/QWF3SzpHmykQHsQMixG0cg/UserTweets"
	UserFollowingURL    = "https://x.com/i/api/graphql/7oQrdmth4zE3EtD42ZxgOA/Following"
	UserInfoURL         = "https://x.com/i/api/graphql/QGIw94L0abhuohrr76cSbw/UserByScreenName"
	UserFollowersURL    = "https://x.com/i/api/graphql/r4fuEJKOqqzaYcvJU5ZWVA/Followers"
	UsersByRestIdsURL   = "https://x.com/i/api/graphql/JnwU1UO8J1tWlOJktPZIzg/UsersByRestIds"
	TweetResultByRestId = "https://x.com/i/api/graphql/zptD3jqLJ0arTQdM10uc3w/TweetResultByRestId"
	UserByRestIdURL     = "https://x.com/i/api/graphql/7oQrdmth4zE3EtD42ZxgOA/UserByRestId"

	// Timeline instruction typename
	TimelineAddEntries = "TimelineAddEntries"

	// Tweet entry type
	TimelineTimelineItem   string = "TimelineTimelineItem"
	TimelineTimelineModule string = "TimelineTimelineModule"
	TimelineTimelineCursor string = "TimelineTimelineCursor"

	ConversationTweetPrefix string = "profile-conversation-"
	CursorBottomPrefix      string = "cursor-bottom-"

	// Tweet created time format
	TimeFormat = time.RubyDate

	// Twitter account status
	UserUnavailable = "UserUnavailable"
	Protected       = "Protected"

	// tweet type
	Post   TweetType = "Post"
	Reply  TweetType = "Reply"
	Repost TweetType = "Repost"
	Quote  TweetType = "Quote"
)

var (
	CommonRequestHeaders = map[string]string{
		"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"content-type":              "application/json",
		"x-twitter-active-user":     "yes",
		"x-twitter-client-language": "en",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	}
)

type RequestParams interface {
	MarshalJsonWithDefaultValue() ([]byte, error)
}

type TwitterClient struct {
	httpCli *resty.Client
}

func NewTwitterClient(a *Account) (*TwitterClient, error) {
	cookiejar, _ := cookiejar.New(nil)
	httpCli := resty.New().SetCookieJar(cookiejar).SetHeaders(CommonRequestHeaders).SetTimeout(RequestTimeout).SetDebug(true)
	if !a.isAccountLogin() {
		err := a.Login(httpCli)
		if err != nil {
			return nil, errors.Wrap(err, "TwitterClient NewTwitterClient Login Error")
		}
	}
	headers := map[string]string{
		"x-csrf-token":        a.Cookies["ct0"].Value,
		"cookie":              a.GetRequestCookie(),
		"x-twitter-auth-type": "OAuth2Session",
	}
	httpCli.SetHeaders(headers)
	return &TwitterClient{
		httpCli: httpCli,
	}, nil
}

func (c *TwitterClient) SetCookies(cookies []*http.Cookie) {
	c.httpCli.SetCookies(cookies)
}

func (c *TwitterClient) SetHeaders(headers map[string]string) {
	c.httpCli.SetHeaders(headers)
}

func (c *TwitterClient) generateReqURL(baseURL string, variables, features, fieldToggles RequestParams) (string, error) {
	variablesJSON, err := variables.MarshalJsonWithDefaultValue()
	if err != nil {
		return "", errors.Wrapf(err, "Marshaling variables to JSON failed, baseURL: %s, variables: %v", baseURL, variables)
	}

	featuresJSON, err := features.MarshalJsonWithDefaultValue()
	if err != nil {
		return "", errors.Wrapf(err, "Marshaling features to JSON failed, baseURL: %s, features: %v", baseURL, features)
	}

	fieldTogglesJSON := []byte{}
	if fieldToggles != nil {
		fieldTogglesJSON, err = fieldToggles.MarshalJsonWithDefaultValue()
		if err != nil {
			return "", errors.Wrapf(err, "Marshaling fieldToggles to JSON failed, baseURL: %s, fieldToggles: %v", baseURL, fieldToggles)
		}
	}

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return "", errors.Wrapf(err, "Parse baseURL failed, baseURL: %s", baseURL)
	}

	query := reqURL.Query()
	query.Set("variables", string(variablesJSON))
	query.Set("features", string(featuresJSON))
	if string(fieldTogglesJSON) != "" {
		query.Set("fieldToggles", string(fieldTogglesJSON))
	}
	reqURL.RawQuery = query.Encode()
	return reqURL.String(), nil
}
