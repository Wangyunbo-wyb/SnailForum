package logic

import (
	"SnailForum/model"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetGithubTrending 获取Github热榜项目
func GetGithubTrending(p *model.ParamGithubTrending) (data *model.GithubTrending, err error) {
	switch p.Language {
	case 0:
		data, err = GetGithubTrendingAll(p)
		//case 1:
		//	data, err = models.GetGithubTrendingGo(p.Since, p.Page, p.Size)
		//default:
		//	data, err = models.GetGithubTrendingAll(p.Since, p.Page, p.Size)
	}
	return
}

// GetGithubTrendingAll 获取Github全语言热榜项目
func GetGithubTrendingAll(p *model.ParamGithubTrending) (data *model.GithubTrending, err error) {
	query := url.QueryEscape("stars:>1")
	url := "https://api.github.com/search/repositories?q=" + query +
		"&sort=stars&order=desc&page=" + fmt.Sprintf("%d", p.Page) +
		"&per_page=" + fmt.Sprintf("%d", p.Size)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		zap.L().Error("http.NewRequest failed", zap.Error(err))
		return
	}
	res, err := client.Do(req)
	if err != nil {
		zap.L().Error("client.Do failed", zap.Error(err))
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		zap.L().Error("ioutil.ReadAll failed", zap.Error(err))
		return
	}
	fmt.Println(string(body))
	var githubTrendingAll model.GithubTrending
	err = json.Unmarshal(body, &githubTrendingAll)
	if err != nil {
		zap.L().Error("json.Unmarshal failed", zap.Error(err))
		return
	}
	return &githubTrendingAll, nil
}
