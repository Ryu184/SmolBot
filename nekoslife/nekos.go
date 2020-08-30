package nekoslife

import (
	"errors"
	"fmt"
	"strings"
	"encoding/json"

	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/yagpdb/commands"
)

var (
	ErrConnection = errors.New("Couldn't contact the API...")
	logger = common.GetPluginLogger(&Plugin{})
)

type Plugin struct {}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Nekos.Life",
		SysName:  "nekos_life",
		Category: common.PluginCategoryMisc,
	}
}

func RegisterPlugin() {
	common.RegisterPlugin(&Plugin{})
}

var _ commands.CommandProvider = (*Plugin)(nil)

func (p *Plugin) AddCommands() {
	commands.AddRootCommands(p,
		// non-lewd
		WallpaperCommand,
		// NekoGifCommand,
		// TickleCommand,
		// FeedCommand,
		// GECGCommand,
		// KemonomimiCommand,
		// PokeCommand,
		// SlapCommand,
		// AvatarCommand,
		// HoloCommand,
		// WaifuCommand,
		// PatCommand,
		// KissCommand,
		// NekoCommand,
		// CuddleCommand,
		// FoxgirlCommand,
		// HugCommand,
		// SmugCommand,
		// BakaCommand,
		// WoofCommand,

		// // lewd
		// FeetCommand,
		// YuriCommand,
		// TrapCommand,
		// FutaCommand,
		// HololewdCommand,
		// LewdKemonoCommand,
		// SoloGifCommand,
		// FeetGifCommand,
		// CumGifCommand,
		// EroKemonoCommand,
		// LesbianGifCommand,
		// LewdKitsuneCommand,
		// LewdCommand,
		// FeedCommand,
		// EroYuriCommand,
		// EroNekoCommand,
		// CumCommand,
		// BlowjobGifCommand,
		// NsfwNekoGifCommand,
		// SoloCommand,
		// NsfwAvatarCommand,
		// AnalCommand,
		// HentaiCommand,
		// EroFeetCommand,
		// KetaCommand,
		// BlowjobCommand,
		// PussyGifCommand,
		// TitsCommand,
		// HoloeroCommand,
		// PussyCommand,
		// PwankGifCommand,
		// ClassicCommand,
		// KuniCommand,
		// FemdomCommand,
		// SpankCommand,
		// EroKitsuneCommand,
		// BoobsCommand,
		// RandomHentaiGifCommand,
		// SmallBoobsCommand,
		// EroCommand,

		// // misc
		// OwoifyCommand,
		// CatCommand,
		// WhyCommand,
		// FactCommand,
	)
}

type ImageResult struct {
	Url string
}

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 3 * time.Second}
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func getImage(endpoint string, data *dcmd.Data) (interface{}, error) {
	result := &ImageResult{}
	url := fmt.Sprintf("https://nekos.life/api/v2/img/%s", endpoint)
	err := getJson(url, result)

	if err != nil {
		return nil, err
	}

	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: result.Url
		},
	}

	return embed, nil
}

var Command = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Wallpaper",
	Description: "Grabs a wallpaper from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("wallpaper", data) },
}
