package twitter

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cast"

	"github.com/go-resty/resty/v2"
)

const (
	LoggerPrefix = "[TwitterClient]"
)

type TwitterClient struct {
	RestyClient *resty.Client
}

func NewTwitterClient(headers map[string]string) *TwitterClient {
	return &TwitterClient{
		RestyClient: resty.New().SetTimeout(10 * time.Second).SetHeaders(headers),
	}
}

func (c *TwitterClient) generateReqURL(baseURL string, variables, features, fieldToggles twitterAPIParams) (string, error) {
	variablesJSON, err := variables.MarshalJsonWithDefaultValue()
	if err != nil {
		return "", errors.Wrap(err, "TwitterClient generateReqURL variables.MarshalJsonWithDefaultValue error")
	}

	featuresJSON, err := features.MarshalJsonWithDefaultValue()
	if err != nil {
		return "", errors.Wrap(err, "TwitterClient generateReqURL features.MarshalJsonWithDefaultValue error")
	}

	fieldTogglesJSON := []byte{}
	if fieldToggles != nil {
		fieldTogglesJSON, err = fieldToggles.MarshalJsonWithDefaultValue()
		if err != nil {
			return "", errors.Wrap(err, "TwitterClient generateReqURL fieldToggles.MarshalJsonWithDefaultValue error")
		}
	}

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return "", errors.Wrap(err, "TwitterClient generateReqURL url.Parse error")
	}

	query := reqURL.Query()
	query.Set("variables", string(variablesJSON))
	query.Set("features", string(featuresJSON))
	if string(fieldTogglesJSON) != "" {
		query.Set("fieldToggles", string(fieldTogglesJSON))
	}
	// log.Info(query)
	reqURL.RawQuery = query.Encode()
	return reqURL.String(), nil
}

func (c *TwitterClient) GetUserTweets(query *TimeLineReqVariables) (resp *UserTweetsResponse, err error) {
	resp = &UserTweetsResponse{}
	features := &CommonReqFeatures{}
	reqURL, err := c.generateReqURL(UserTweetsGraphqlURL, query, features, nil)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient generateReqURL Error")
	}

	userTweetsResp := &TimeLineResponse{}
	req, err := c.RestyClient.R().SetResult(userTweetsResp).Get(reqURL)
	if err != nil {
		fmt.Println("Fetched User Tweets Failed, TwitterID: %s, Response Status: %d, resp: %s",
			query.UserID, req.StatusCode(), req.String())
		return nil, errors.Wrap(err, "TwitterClient Request GetUserTweets error")
	}
	// fmt.Println("req status: %d, resp: %s", req.StatusCode(), req.String())

	headers := req.Header()

	fmt.Println(
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
	reqURL, err := c.generateReqURL(TweetDetailGraphqlURL, &variables, features, fieldToggles)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient generateReqURL error")
	}

	tweetDetailResp := &TweetDetailResponse{}
	req, err := c.RestyClient.R().SetResult(&tweetDetailResp).Get(reqURL)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient Request GetUserTweets error")
	}
	headers := req.Header()
	fmt.Println(
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

func (c *TwitterClient) GetUserFollowing(query *FollowingReqVariables) (resp *FollowingResp, err error) {
	resp = &FollowingResp{}
	features := &FollowingReqFeatures{}
	reqURL, err := c.generateReqURL(UserFollowGraphqlURL, query, features, nil)
	if err != nil {
		return nil, errors.Wrap(err, "TwitterClient generateReqURL Error")
	}

	followListResp := &FollowListResp{}
	req, err := c.RestyClient.R().SetResult(followListResp).Get(reqURL)
	if err != nil {
		fmt.Println("Fetched User Following Failed, TwitterID: %s, Response Status: %d, resp: %s",
			query.UserID, req.StatusCode(), req.String())
		return nil, errors.Wrap(err, "TwitterClient Request GetUserTweets error")
	}
	headers := req.Header()
	fmt.Println(
		"Fetched User Following Successful, TwitterID: %s, Response Status: %d, Header: limit: %s, reset: %s, time: %s, remaining: %s",
		query.UserID, req.StatusCode(),
		headers.Get("X-Rate-Limit-Limit"), headers.Get("X-Rate-Limit-Reset"),
		headers.Get("X-Response-Time"), headers.Get("X-Rate-Limit-Remaining"),
	)

	// fmt.Println("req status: %d, resp: %s", req.StatusCode(), req.String())
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
