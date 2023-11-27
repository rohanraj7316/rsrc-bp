package repository

import (
	"errors"
	"fmt"
	"net/url"
	"sort"

	"github.com/zhengxiaowai/shortuuid"
)

type ShortUrls interface {
	Get(shortId string) (originalUrl string, err error)
	Top3ShortedDomain() (list map[string]int, err error)
	Create(originalUrl string) (shortId string, err error)
}

type shortUrls struct {
	shortIdToOriginalUrls map[string]string
}

func NewShortUrls() ShortUrls {
	return &shortUrls{shortIdToOriginalUrls: map[string]string{}}
}

func createShortId() string {
	su := shortuuid.NewShortUUID()
	return su.Random(5)
}

func (s *shortUrls) Get(srtId string) (orgUrl string, err error) {
	if originalUrl, ok := s.shortIdToOriginalUrls[srtId]; ok {
		return originalUrl, nil
	}

	return "", errors.New("not found")
}

func (s *shortUrls) Top3ShortedDomain() (map[string]int, error) {
	numberOfTimesDomainShorted := map[string]int{}
	for _, originalUrl := range s.shortIdToOriginalUrls {
		parsedUrl, err := url.Parse(originalUrl)
		if err != nil {
			return nil, err
		}

		if currentCount, exists := numberOfTimesDomainShorted[parsedUrl.Host]; exists {
			numberOfTimesDomainShorted[parsedUrl.Host] = currentCount + 1
		} else {
			numberOfTimesDomainShorted[parsedUrl.Host] = 1
		}

	}

	type kv struct {
		key   string
		value int
	}

	var ss []kv
	for k, v := range numberOfTimesDomainShorted {
		ss = append(ss, kv{k, v})
	}

	if len(ss) < 3 {
		return nil, fmt.Errorf("domain count: %d which is less than 3", len(ss))
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})

	ss = ss[:3]

	top3ShortedDomain := map[string]int{}
	for _, s := range ss {
		top3ShortedDomain[s.key] = s.value
	}

	return top3ShortedDomain, nil
}

func (s *shortUrls) Create(originalUrl string) (srtId string, err error) {
	if len(originalUrl) == 0 {
		return "", errors.New("empty original url")
	}

	// validation for original url duplication
	for shortId, haveOriginalUrl := range s.shortIdToOriginalUrls {
		if originalUrl == haveOriginalUrl {
			return shortId, nil
		}
	}

	// create shortId
	shortId := createShortId()

	s.shortIdToOriginalUrls[shortId] = originalUrl

	return shortId, nil
}
