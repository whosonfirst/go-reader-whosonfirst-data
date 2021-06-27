package reader

import (
	_ "github.com/whosonfirst/go-reader-github"
)

import (
	"context"
	"fmt"
	wof_reader "github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-findingaid"
	"github.com/whosonfirst/go-whosonfirst-findingaid/repo"
	"io"
	_ "log"
	"net/url"
	"sync"
	"time"
)

type WhosOnFirstDataReader struct {
	wof_reader.Reader
	throttle     <-chan time.Time
	provider     string
	organization string
	repo         string
	branch       string
	repos        *sync.Map
	readers      *sync.Map
	resolver     findingaid.Resolver
}

func init() {

	ctx := context.Background()
	err := wof_reader.RegisterReader(ctx, "whosonfirst-data", NewWhosOnFirstDataReader)

	if err != nil {
		panic(err)
	}
}

func NewWhosOnFirstDataReader(ctx context.Context, uri string) (wof_reader.Reader, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	provider := q.Get("provider")
	org := q.Get("organization")
	repo := q.Get("repo")
	branch := q.Get("branch")

	if provider == "" {
		provider = "github"
	}

	if org == "" {
		org = "whosonfirst-data"
	}

	fa_uri := q.Get("findingaid-uri")

	if fa_uri == "" {
		fa_uri = "repo-http"
	}

	resolver, err := findingaid.NewResolver(ctx, fa_uri)

	if err != nil {
		return nil, err
	}

	rate := time.Second / 3
	throttle := time.Tick(rate)

	repos := new(sync.Map)
	readers := new(sync.Map)

	r := &WhosOnFirstDataReader{
		throttle:     throttle,
		provider:     provider,
		organization: org,
		repo:         repo,
		branch:       branch,
		repos:        repos,
		readers:      readers,
		resolver:     resolver,
	}

	return r, nil
}

func (r *WhosOnFirstDataReader) Read(ctx context.Context, uri string) (io.ReadSeekCloser, error) {

	<-r.throttle

	select {
	case <-ctx.Done():
		return nil, nil
	default:
		// pass
	}

	repo := r.repo

	if repo == "" {

		this_repo, err := r.getRepo(ctx, uri)

		if err != nil {
			return nil, err
		}

		repo = this_repo
	}

	gh_r, err := r.getReader(ctx, repo)

	if err != nil {
		return nil, err
	}

	return gh_r.Read(ctx, uri)
}

func (r *WhosOnFirstDataReader) getReader(ctx context.Context, repo string) (wof_reader.Reader, error) {

	v, ok := r.readers.Load(repo)

	if ok {
		gh_r := v.(wof_reader.Reader)
		return gh_r, nil
	}

	gh_q := url.Values{}

	if r.branch != "" {
		gh_q.Set("branch", r.branch)
	}

	gh_uri := url.URL{}
	gh_uri.Scheme = r.provider
	gh_uri.Host = r.organization
	gh_uri.Path = repo
	gh_uri.RawQuery = gh_q.Encode()

	reader_uri := gh_uri.String()

	gh_r, err := wof_reader.NewReader(ctx, reader_uri)

	if err != nil {
		return nil, err
	}

	go func() {
		r.readers.Store(repo, gh_r)
	}()

	return gh_r, nil
}

func (r *WhosOnFirstDataReader) getRepo(ctx context.Context, uri string) (string, error) {

	v, ok := r.repos.Load(uri)

	if ok {
		repo_name := v.(string)
		return repo_name, nil
	}

	fa_rsp, err := r.resolver.ResolveURI(ctx, uri)

	if err != nil {
		return "", err
	}

	var repo_name string

	switch fa_rsp.(type) {
	case *repo.FindingAidResponse:

		rsp := fa_rsp.(*repo.FindingAidResponse)
		repo_name = rsp.Repo
	default:
		return "", fmt.Errorf("Unexpected response type from finding aid")
	}

	go func() {
		r.repos.Store(uri, repo_name)
	}()

	return repo_name, nil
}
