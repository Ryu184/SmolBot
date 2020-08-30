package nekoslife

import (
	"errors"
	"fmt"
	"time"
	"encoding/json"
	"net/http"

	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/dstate"
	"github.com/jonas747/yagpdb/commands"
	"github.com/jonas747/yagpdb/common"
)

var (
	ErrConnection = errors.New("Couldn't contact the API...")
	logger = common.GetPluginLogger(&Plugin{})

	TickleText = "<@!%d> tickles <@!%d>!"
	FeedText = "<@!%d> feeds <@!%d>!"
	PokeText = "<@!%d> pokes <@!%d>!"
	SlapText = "<@!%d> slaps <@!%d>!"
	PatText = "<@!%d> pats <@!%d>!"
	KissText = "<@!%d> kisses <@!%d>!"
	CuddleText = "<@!%d> cuddles <@!%d>!"
	HugText = "<@!%d> hugs <@!%d>!"
	BakaText = "<@!%d> calls <@!%d> a BAKA!"
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
		NekoGifCommand,
		GECGCommand,
		KemonomimiCommand,
		AvatarCommand,
		HoloCommand,
		WaifuCommand,
		NekoCommand,
		FoxgirlCommand,
		SmugCommand,
		WoofCommand,

		TickleCommand,
		FeedCommand,
		PokeCommand,
		SlapCommand,
		PatCommand,
		KissCommand,
		CuddleCommand,
		HugCommand,
		BakaCommand,

		// lewd
		FeetCommand,
		YuriCommand,
		TrapCommand,
		FutaCommand,
		HololewdCommand,
		LewdKemonoCommand,
		SoloGifCommand,
		FeetGifCommand,
		CumGifCommand,
		EroKemonoCommand,
		LesbianGifCommand,
		LewdKitsuneCommand,
		LewdCommand,
		EroYuriCommand,
		EroNekoCommand,
		CumCommand,
		BlowjobGifCommand,
		NsfwNekoGifCommand,
		SoloCommand,
		NsfwAvatarCommand,
		AnalCommand,
		HentaiCommand,
		EroFeetCommand,
		KetaCommand,
		BlowjobCommand,
		PussyGifCommand,
		TitsCommand,
		HoloeroCommand,
		PussyCommand,
		PwankGifCommand,
		ClassicCommand,
		KuniCommand,
		FemdomCommand,
		SpankCommand,
		EroKitsuneCommand,
		BoobsCommand,
		RandomHentaiGifCommand,
		EroCommand,

		// misc
		CatCommand,
		WhyCommand,
		FactCommand,
	)
}

type ImageResult struct {
	Url string
}

type CatResult struct {
	Cat string
}

type WhyResult struct {
	Why string
}

type FactResult struct {
	Fact string
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
			URL: result.Url,
		},
	}

	return embed, nil
}

func getImageText(endpoint string, format string, data *dcmd.Data) (interface{}, error) {
	result := &ImageResult{}
	url := fmt.Sprintf("https://nekos.life/api/v2/img/%s", endpoint)
	err := getJson(url, result)
	
	var member *dstate.MemberState
	if data.Args[0].Value != nil {
		member = data.Args[0].Value.(*dstate.MemberState)
	} else {
		member = nil
	}

	if err != nil {
		return nil, err
	}

	var message string
	var embed interface{}

	if member == nil {
		message = fmt.Sprintf("Are you trying to %s the void...?", endpoint)
		embed = &discordgo.MessageEmbed{
			Description: message,
		}
	} else if member.ID == data.MS.ID {
		message = fmt.Sprintf("Sorry to see you alone, <@!%d> ;-;", member.ID)
		embed = &discordgo.MessageEmbed{
			Description: message,
		}
	} else {
		embed = &discordgo.MessageEmbed{
			Description: fmt.Sprintf(format, data.MS.ID, member.ID),
			Image: &discordgo.MessageEmbedImage{
				URL: result.Url,
			},
		}
	}

	return embed, nil
}

// Simple image commands (SFW)
var WallpaperCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Wallpaper",
	Description: "Grabs a wallpaper from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("wallpaper", data) },
}

var NekoGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "NekoGif",
	Description: "Grabs a NekoGif from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("ngif", data) },
}

var GECGCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "GECG",
	Description: "Grabs a GeneticallyEngineeredCatGirl meme from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("gecg", data) },
}

var KemonomimiCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Kemonomimi",
	Description: "Grabs a Kemonomimi from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("kemonomimi", data) },
}

var AvatarCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "AnimeAvatar",
	Description: "Grabs a Avatar from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("avatar", data) },
}

var HoloCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Holo",
	Description: "Grabs a Holo from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("holo", data) },
}

var WaifuCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Waifu",
	Description: "Grabs a Waifu from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("waifu", data) },
}

var NekoCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Neko",
	Description: "Grabs a Neko from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("neko", data) },
}

var FoxgirlCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Foxgirl",
	Description: "Grabs a Foxgirl from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("fox_girl", data) },
}

var SmugCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Smug",
	Description: "Grabs a Smug from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("smug", data) },
}

var WoofCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Woof",
	Description: "Grabs a Woof from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("woof", data) },
}


// Commands to kiss/poke/pat/etc someone.
var TickleCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Tickle",
	Description: "Tickle someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("tickle", TickleText, data) },
}

var FeedCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Feed",
	Description: "Feed someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("feed", FeedText, data) },
}

var PokeCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Poke",
	Description: "Poke someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("poke", PokeText, data) },
}

var SlapCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Slap",
	Description: "Slap someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("slap", SlapText, data) },
}

var PatCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Pat",
	Description: "Pat someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("pat", PatText, data) },
}

var KissCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Kiss",
	Description: "Kiss someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("kiss", KissText, data) },
}

var CuddleCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Cuddle",
	Description: "Cuddle someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("cuddle", CuddleText, data) },
}

var HugCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Hug",
	Description: "Hug someone with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("hug", HugText, data) },
}

var BakaCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Baka",
	Description: "Call someone a BAKA with a GIF from nekos.life.",
	Arguments: []*dcmd.ArgDef{
		{Name: "User", Type: &commands.MemberArg{}},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImageText("baka", BakaText, data) },
}


// Simple image commands (NSFW)
var FeetCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Feet",
	Description: "Grabs a NSFW Feet image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("feet", data) },
}

var YuriCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Yuri",
	Description: "Grabs a NSFW Yuri image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("yuri", data) },
}

var TrapCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Trap",
	Description: "Grabs a NSFW Trap image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("trap", data) },
}

var FutaCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Futa",
	Description: "Grabs a NSFW Futa image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("futanari", data) },
}

var HololewdCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Hololewd",
	Description: "Grabs a NSFW Holo image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("hololewd", data) },
}

var LewdKemonoCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "LewdKemono",
	Description: "Grabs a NSFW Kemonomimi image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("lewdkemo", data) },
}

var SoloGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "SoloGif",
	Description: "Grabs a NSFW Solo GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("solog", data) },
}

var FeetGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "FeetGif",
	Description: "Grabs a NSFW Feet GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("feetg", data) },
}

var CumGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "CumGif",
	Description: "Grabs a NSFW Cum GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("cum", data) },
}

var EroKemonoCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "EroKemono",
	Description: "Grabs a NSFW Kemono image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("erokemo", data) },
}

var LesbianGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "LesbianGif",
	Description: "Grabs a NSFW Lesbian GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("les", data) },
}

var LewdKitsuneCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "LewdKitsune",
	Description: "Grabs a NSFW Kitsune image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("lewdk", data) },
}

var LewdCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Lewd",
	Description: "Grabs a NSFW Lewd image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("lewd", data) },
}

var EroYuriCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "EroYuri",
	Description: "Grabs a NSFW Yuri image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("eroyuri", data) },
}

var EroNekoCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "EroNeko",
	Description: "Grabs a NSFW Neko image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("eron", data) },
}

var CumCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Cum",
	Description: "Grabs a NSFW Cum image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("cum_jpg", data) },
}

var BlowjobGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "BlowjobGif",
	Description: "Grabs a NSFW Blowjob GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("bj", data) },
}

var NsfwNekoGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "NsfwNekoGif",
	Description: "Grabs a NSFW Neko GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("nsfw_neko_gif", data) },
}

var SoloCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd,
	Name:        "Solo",
	Description: "Grabs a NSFW Solo image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("solo", data) },
}

var NsfwAvatarCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "NsfwAnimeAvatar",
	Description: "Grabs a NSFW Anime Avatar image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("nsfw_avatar", data) },
}

var AnalCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Anal",
	Description: "Grabs a NSFW Anal image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("anal", data) },
}

var HentaiCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Hentai",
	Description: "Grabs a NSFW Hentai image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("hentai", data) },
}

var EroFeetCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "EroFeet",
	Description: "Grabs a NSFW Feet image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("erofeet", data) },
}

var KetaCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Keta",
	Description: "Grabs a NSFW Ke-Ta image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("keta", data) },
}

var BlowjobCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Blowjob",
	Description: "Grabs a NSFW Blowjob image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("blowjob", data) },
}

var PussyGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "PussyGif",
	Description: "Grabs a NSFW Pussy GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("pussy", data) },
}

var TitsCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Tits",
	Description: "Grabs a NSFW Tits image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("tits", data) },
}

var HoloeroCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Holoero",
	Description: "Grabs a NSFW Holo image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("holoero", data) },
}

var PussyCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Pussy",
	Description: "Grabs a NSFW Pussy image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("pussy_jpg", data) },
}

var PwankGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "PwankGif",
	Description: "Grabs a NSFW Pwank GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("pwankg", data) },
}

var ClassicCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Classic",
	Description: "Grabs a NSFW Classic image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("classic", data) },
}

var KuniCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Kuni",
	Description: "Grabs a NSFW Kuni image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("kuni", data) },
}

var FemdomCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Femdom",
	Description: "Grabs a NSFW Femdom image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("femdom", data) },
}

var SpankCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Spank",
	Description: "Grabs a NSFW Spank image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("spank", data) },
}

var EroKitsuneCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "EroKitsune",
	Description: "Grabs a NSFW Kitsune image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("erok", data) },
}

var BoobsCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Boobs",
	Description: "Grabs a NSFW Boobs image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("boobs", data) },
}

var RandomHentaiGifCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "RandomHentaiGif",
	Description: "Grabs a NSFW Random Hentai GIF from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("Random_hentai_gif", data) },
}

var EroCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnimeLewd2,
	Name:        "Ero",
	Description: "Grabs a NSFW Ero image from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) { return getImage("ero", data) },
}


// Misc text-only commands
var CatCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Neko",
	Description: "Grabs a random cat text emoji from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		result := &CatResult{}
		url := fmt.Sprintf("https://nekos.life/api/v2/cat")
		err := getJson(url, result)

		if err != nil {
			return nil, err
		}

		return result.Cat, nil
	},
}

var WhyCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Why",
	Description: "Grabs a random why question from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		result := &WhyResult{}
		url := fmt.Sprintf("https://nekos.life/api/v2/why")
		err := getJson(url, result)

		if err != nil {
			return nil, err
		}

		return result.Why, nil
	},
}

var FactCommand = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryAnime,
	Name:        "Fact",
	Description: "Grabs a random fact from nekos.life.",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		result := &FactResult{}
		url := fmt.Sprintf("https://nekos.life/api/v2/fact")
		err := getJson(url, result)

		if err != nil {
			return nil, err
		}

		return result.Fact, nil
	},
}
