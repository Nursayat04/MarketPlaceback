package clients

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type ReviewClient struct {
	client  *resty.Client
	baseURL string
}

func NewReviewClient() *ReviewClient {
	client := resty.New()

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		log.Printf("[ReviewClient] --> %s %s", req.Method, req.URL)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		log.Printf("[ReviewClient] <-- %d %s", resp.StatusCode(), resp.Request.URL)
		return nil
	})

	return &ReviewClient{
		client:  client,
		baseURL: "http://review-app:8081",
	}
}

func (rc *ReviewClient) GetReviews(productID string) (string, error) {
	resp, err := rc.client.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/reviews/%s", rc.baseURL, productID))
	if err != nil {
		return "", fmt.Errorf("failed to get reviews: %w", err)
	}
	return resp.String(), nil
}

func (rc *ReviewClient) AddReview(body map[string]interface{}) (string, error) {
	resp, err := rc.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(fmt.Sprintf("%s/reviews", rc.baseURL))
	if err != nil {
		return "", fmt.Errorf("failed to add review: %w", err)
	}
	return resp.String(), nil
}

func (rc *ReviewClient) DeleteReview(id string) (string, error) {
	resp, err := rc.client.R().
		Delete(fmt.Sprintf("%s/reviews/%s", rc.baseURL, id))
	if err != nil {
		return "", fmt.Errorf("failed to delete review: %w", err)
	}
	return resp.String(), nil
}
