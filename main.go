package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const (
	misso = "<@734253752877580288>"
	lais  = "<@267454529652523013>"
	chau  = "<@510247392193544197>"
)

type MarketEntry struct {
	Data []ItemSellingInfo `json:"data"`
}

type ItemSellingInfo struct {
	OrderResponse OrderResponse `json:"orders"`
	ItemResponse  ItemResponse  `json:"items"`
}

type OrderResponse struct {
	ItemPrice    int64  `json:"price"`
	ItemRefine   int64  `json:"refine"`
	ItemLocation string `json:"location"`
}

type ItemResponse struct {
	ItemProperty string `json:"property"`
}

type ItemMatch struct {
	ItemId       string
	Name         string
	Refine       int64
	MinPrice     int64
	Property     []string
	Owners       []string
	DisableUntil time.Time
}

func PopulateItemMatch() []ItemMatch {
	return []ItemMatch{
		{
			ItemId: "1631", Name: "Holy Stick", Refine: 16, MinPrice: 1000000000, Owners: []string{misso},
		},

		{
			ItemId: "33463", Name: "Watchwork Butterfly", Refine: 0, MinPrice: 400000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32688", Name: "Bunbun Scarf", Refine: 0, MinPrice: 300000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "19976", Name: "Cat Santa Hat", Refine: 0, MinPrice: 50000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "19678", Name: "Bell Ribbon", Refine: 0, MinPrice: 100000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "29704", Name: "Costume Glowing Halo", Refine: 0, MinPrice: 150000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "33233", Name: "Evil Druid Cross", Refine: 0, MinPrice: 150000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "33219", Name: "Master of Darkness", Refine: 0, MinPrice: 150000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "33220", Name: "Evil's Scythe", Refine: 0, MinPrice: 400000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "29678", Name: "Poring Soap Pipe", Refine: 0, MinPrice: 100000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32482", Name: "Sleepy Sheep", Refine: 0, MinPrice: 120000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32476", Name: "Master of Wind", Refine: 0, MinPrice: 200000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "29629", Name: "Costume Happy Summer Ribbon", Refine: 0, MinPrice: 50000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "18663", Name: "Sunglasses Baseball Hat", Refine: 0, MinPrice: 50000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "31607", Name: "Clock Tower Hat", Refine: 0, MinPrice: 80000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32362", Name: "Summer Petals", Refine: 0, MinPrice: 143000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32272", Name: "Master of Flowe", Refine: 0, MinPrice: 130000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "33647", Name: "Fairy Of Eden", Refine: 0, MinPrice: 350000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32862", Name: "Occult Magic", Refine: 0, MinPrice: 300000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32802", Name: "Wing of Heart", Refine: 0, MinPrice: 300000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32922", Name: "Starry Star", Refine: 0, MinPrice: 60000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "31448", Name: "Pink Silk Ribbon", Refine: 0, MinPrice: 60000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "33564", Name: "Knigh's Cloak", Refine: 0, MinPrice: 400000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "5979", Name: "Fluttering Angel", Refine: 0, MinPrice: 100000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32790", Name: "Cat Ear Ribbon", Refine: 0, MinPrice: 200000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32566", Name: "Loli Ruri Moon", Refine: 0, MinPrice: 450000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "9455", Name: "Hillslion Egg", Refine: 0, MinPrice: 20000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "9580", Name: "Aries Egg", Refine: 0, MinPrice: 20000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32084", Name: "Rose Scent", Refine: 0, MinPrice: 143000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32255", Name: "Queen's Anne", Refine: 0, MinPrice: 350000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "32406", Name: "Nova Point", Refine: 0, MinPrice: 4500000, Owners: []string{misso, lais},
		},

		{
			ItemId: "672", Name: "Gold Coin", Refine: 0, MinPrice: 30000, Owners: []string{misso, lais, chau},
		},

		{
			ItemId: "30043", Name: "Aquastone Ingot", Refine: 0, MinPrice: 1400000, Owners: []string{misso, lais},
		},

		{
			ItemId: "30042", Name: "Flamestone Ingot", Refine: 0, MinPrice: 1400000, Owners: []string{misso, lais},
		},

		{
			ItemId: "30044", Name: "Dragonstone Ingot", Refine: 0, MinPrice: 1400000, Owners: []string{misso, lais},
		},

		{
			ItemId: "25786", Name: "Biological Research Documents", Refine: 0, MinPrice: 100000, Owners: []string{misso, lais},
		},

		{
			ItemId: "6635", Name: "Blacksmith's Blessing", Refine: 0, MinPrice: 9000000, Owners: []string{misso, lais},
		},

		{
			ItemId: "12103", Name: "Bloody Branch", Refine: 0, MinPrice: 3000000, Owners: []string{misso, lais},
		},
	}

}

func GetBestMatch(itemmatch ItemMatch, marketentry MarketEntry) *ItemSellingInfo {

	var result *ItemSellingInfo

	if itemmatch.DisableUntil.After(time.Now()) {
		return nil
	}

	for i := range marketentry.Data {
		if marketentry.Data[i].OrderResponse.ItemRefine >= itemmatch.Refine && marketentry.Data[i].OrderResponse.ItemPrice <= itemmatch.MinPrice {
			if result == nil {
				result = &marketentry.Data[i]
			} else if result.OrderResponse.ItemPrice > marketentry.Data[i].OrderResponse.ItemPrice {
				result = &marketentry.Data[i]
			}

		}
	}
	return result

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_KEY"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	client := &http.Client{}
	itemmatch := PopulateItemMatch()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	// Cleanly close down the Discord session.
	defer dg.Close()

	for true {
		for i := range itemmatch {
			url := fmt.Sprintf("https://www.novaragnarok.com/data/cache/ajax/item_%s.json", itemmatch[i].ItemId)

			authcookieid := "fluxSessionData=" + os.Getenv("FLUX_SESSION_DATA")

			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Add("cookie", authcookieid)
			resp, _ := client.Do(req)

			itemsumary := MarketEntry{}

			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			_ = json.Unmarshal(body, &itemsumary)

			bestmatch := GetBestMatch(itemmatch[i], itemsumary)

			if bestmatch != nil {
				msg := fmt.Sprintf(
					"%s Item **%s** por **%d** no market. Localizacao **%s**",
					strings.Join(itemmatch[i].Owners, ", "),
					itemmatch[i].Name,
					bestmatch.OrderResponse.ItemPrice,
					bestmatch.OrderResponse.ItemLocation,
				)

				dg.ChannelMessageSend("841069698305359892", msg)
				itemmatch[i].DisableUntil = time.Now().Add(3 * time.Hour)
			}

		}
		time.Sleep(5 * time.Minute)
	}
}
