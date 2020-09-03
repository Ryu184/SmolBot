package actuallydefine

import (
	"encoding/json"
	"fmt"
	"time"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/jonas747/dcmd"
	"github.com/jonas747/yagpdb/commands"
	"github.com/jonas747/yagpdb/common/config"
)

type SearchResult struct {
	Type          string `json:"result_type"`
	Tags          []string
	Results       []Result `json:"definitions"`
	Word          string `json:"word"`
	Pronunciation string `json:"pronunciation"`
}

type Result struct {
	Type          string
	Definition    string
	Example       string
}

const API_URL = "https://owlbot.info/api/v4/dictionary/"

var (
	confApiKey  = config.RegisterOption("yagpdb.owlbotapikey", "OwlBot API key", "")
)

func Query(searchTerm string) (*SearchResult, error) {
	req, err := http.NewRequest("GET", API_URL + url.QueryEscape(searchTerm), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token " + confApiKey.GetString()) 
	
	client := &http.Client{Timeout: time.Second * 10}


	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Response was not a 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "Actually Define",
	Aliases:      []string{"acdf"},
	Description:  "Look up an owlbot definition",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "Topic", Type: dcmd.String},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {

		qResp, err := Query(data.Args[0].Str())
		if err != nil {
			return "Failed querying :(", err
		}

		if len(qResp.Results) < 1 {
			return "No result :(", nil
		}
		result := qResp.Results
		word := qResp.Word
		pronunciation := qResp.Pronunciation

		cmdResp := fmt.Sprintf("**%s**[%s]:\n 1. %s: %s\n*%s*\n", word, pronunciation, result[0].Type, result[0].Definition, result[0].Example)

		if len(qResp.Results) > 1 {
			for i := 1 ; i < len(qResp.Results) ; i++{
				cmdResp += fmt.Sprintf("%d. %s: %s\n%s\n", i+1, result[i].Type, result[i].Definition, result[i].Example)		
			}
		}
		
		cmdResp = strings.ReplaceAll(cmdResp, "<b>", "**")
		cmdResp = strings.ReplaceAll(cmdResp, "</b>", "**")
		cmdResp = strings.ReplaceAll(cmdResp, "<i>", "*")
		cmdResp = strings.ReplaceAll(cmdResp, "</i>", "*")

		return cmdResp, nil
	},
}
