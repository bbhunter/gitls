package modules

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"

	model "github.com/hahwul/gitls/pkg/model"
	transport "github.com/hahwul/gitls/pkg/transport"
)

// GithubObject is json object of github api
type GithubObject struct {
	Html_URL string `"json:html_url"`
	Fork bool `"json:fork"`
}

// GetRepoListFromUser is gettting repo list from github
func GetRepoListFromUser(user,repoHost string, options model.Options){
	check := true
	for i:=1 ; check ; i++ {
		apiAddress := fmt.Sprintf("https://api."+repoHost+"/users/%v/repos?page=%v&per_page=100", user, i)
		req, err := http.NewRequest("GET",apiAddress,nil)
		transport := transport.GetTransport(options)
		client := &http.Client{
			Timeout:   5 * time.Second,
			Transport: transport,
		}

		resp, err := client.Do(req)
		if err	!= nil {

		}

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if string(data) == "[]" {
			check = false
		}
		var objects []GithubObject
		json.Unmarshal(data,&objects)
		for k, v := range objects {
			_ = k
			if !v.Fork {
				fmt.Println(v.Html_URL)
			}
		}
	}
}
