package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	AuthURL     = "https://auth.docker.io/token?service=registry.docker.io&scope=repository:ratelimitpreview/test:pull"
	RegistryURL = "https://registry-1.docker.io/v2/ratelimitpreview/test/manifests/latest"
)

type RateLimit struct {
	Limit     int
	Remaining int
}

type Response struct {
	Token       string    `json:"token"`
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	IssuedAt    time.Time `json:"issued_at"`
}

type Options struct {
	username string
	password string
	token    string
}

type Option func(*Options)

func withAuth(username, password string) Option {
	return func(args *Options) {
		args.username = username
		args.password = password
	}
}

func withToken(token string) Option {
	return func(args *Options) {
		args.token = token
	}
}

func newRequestWithContext(ctx context.Context, method string, url string, options ...Option) (*http.Response, error) {
	opt := Options{}
	for _, o := range options {
		o(&opt)
	}

	c := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	if opt.username != "" && opt.password != "" {
		req.SetBasicAuth(opt.username, opt.password)
	}

	if opt.token != "" {
		req.Header.Add("Authorization", "Bearer "+opt.token)
	}

	return c.Do(req)
}

func getAuthToken(ctx context.Context, username, password string) (string, error) {
	resp, err := newRequestWithContext(ctx, http.MethodGet, AuthURL, withAuth(username, password))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authenticate is failed. status code is %d", resp.StatusCode)
	}

	var res Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

func checkLimit(ctx context.Context, username, password string) (*RateLimit, error) {
	token, err := getAuthToken(ctx, username, password)
	if err != nil {
		return nil, err
	}

	resp, err := newRequestWithContext(ctx, http.MethodHead, RegistryURL, withToken(token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authenticate is failed. status code is %d", resp.StatusCode)
	}

	limit, err := parseHeader(resp.Header.Get("RateLimit-Limit"))
	if err != nil {
		return nil, err
	}
	remaining, err := parseHeader(resp.Header.Get("RateLimit-Remaining"))
	if err != nil {
		return nil, err
	}

	res := &RateLimit{
		Limit:     limit,
		Remaining: remaining,
	}
	return res, nil
}

func parseHeader(s string) (int, error) {
	return strconv.Atoi(strings.Split(s, ";")[0])
}
