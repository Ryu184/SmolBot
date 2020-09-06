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

type SearchResult []struct {
	Meta struct {
		ID        string   `json:"id"`
		UUID      string   `json:"uuid"`
		Sort      string   `json:"sort"`
		Src       string   `json:"src"`
		Section   string   `json:"section"`
		Stems     []string `json:"stems"`
		Offensive bool     `json:"offensive"`
	} `json:"meta"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Uros []struct {
		Ure string `json:"ure"`
		Fl  string `json:"fl"`
	} `json:"uros"`
	Et       [][]string `json:"et"`
	Shortdef []string   `json:"shortdef"`
}

const API_URL = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"

var (
	confApiKey  = config.RegisterOption("yagpdb.MerriamWebster", "Merriam-Webster API key", "")
)

func Query(searchTerm string) (*SearchResult, error) {
	resp, err := http.Get(API_URL + url.QueryEscape(searchTerm) + "?key=" + confApiKey.GetString())
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
	Description:  "Look up an Merriam-Webster definition",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "Topic", Type: dcmd.String},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {

		qResp, err := Query(data.Args[0].Str())
		if err != nil {
			return "Failed querying :(", err
		}

		qResp1 := *qResp
		if len(qResp1) < 1 {
			fmt.Println("No result :(", nil)
		}

		cmdResp := fmt.Sprintf("**%s**:\n", qResp1[0].Meta.ID)
		for i := 0 ; i < len(qResp1) ; i++{
			cmdResp += fmt.Sprintf("%d.[%s] %s: %s\n", i+1, qResp1[0].Hwi.Hw, qResp1[i].Fl, qResp1[i].Shortdef[0])
		}

		return cmdResp, nil
	},
}
