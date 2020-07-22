package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	scraper "web_scraper"
)

func main() {
	// get user input
	prompt := promptui.Prompt{
		Label:    "RegionCode",
		Validate: regionValidator,
	}

	region, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	//error already validate
	q, _ := getPresetQuery(region)
	regionName := presets[region].City
	fmt.Printf("Your choose %+v\n", regionName)

	startTime := time.Now()

	s := scraper.NewFiveN1()
	rentals := s.ScrapeRentals(q)
	s.ScrapeRentalsDetail(rentals)

	date := time.Now().Format("2006-01-02")
	filename := date + "-" + regionName
	rentals.ReplaceSection()
	rentals.Print()
	_ = rentals.SaveAsXLSX(filename + ".xlsx")

	log.Printf("execution time %s", time.Since(startTime))
}

func regionValidator(region string) error {
	_, err := getPresetQuery(region)

	return err
}

func getPresetQuery(regionCode string) (*scraper.Query, error) {
	preset, ok := presets[regionCode]
	if !ok {
		return nil, errors.New("region not found")
	}
	rCodeInt, err := strconv.Atoi(regionCode)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("strconv error: %v", err))
	}

	q := scraper.NewQuery()
	q.Kind = 0
	q.Role = "1"
	q.RentPrice = "0,100000"
	q.OrderType = "desc"
	q.Sex = 0
	q.Region = rCodeInt
	q.Section = preset.Sections

	return q, nil
}

type Preset struct {
	City     string
	Sections string
}

var presets = map[string]Preset{
	// with limit sections
	"1": {
		City:     "台北市",
		Sections: "1,2,3,4,5,6,7,8,9,10,11,12",
	},
	"8": {
		City:     "台中市",
		Sections: "98,99,100,101,102,103,104,105",
	},

	// other region
	"2": {
		City:     "基隆市",
		Sections: "13,14,15,16,17,18,19",
	},
	"3": {
		City:     "新北市",
		Sections: "20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52",
	},
	"4": {
		City:     "新竹市",
		Sections: "370,371,372",
	},
	"5": {
		City:     "新竹縣",
		Sections: "54,55,56,57,58,59,60,61,62,63,64,65,66",
	},
	"6": {
		City:     "桃園市",
		Sections: "67,68,69,70,71,72,73,74,75,76,77,78,79",
	},
	"7": {
		City:     "苗栗縣",
		Sections: "80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97",
	},
	"10": {
		City:     "彰化縣",
		Sections: "127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152",
	},
	"11": {
		City:     "南投縣",
		Sections: "127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152",
	},
	"12": {
		City:     "嘉義市",
		Sections: "373,374",
	},
	"13": {
		City:     "嘉義縣",
		Sections: "167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183",
	},
	"14": {
		City:     "雲林縣",
		Sections: "167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183",
	},
	"15": {
		City:     "台南市",
		Sections: "206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242",
	},
	"17": {
		City:     "高雄市",
		Sections: "243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,271,272,273,274,275,276,277,278,279,280,281,282",
	},
	"19": {
		City:     "屏東縣",
		Sections: "243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,271,272,273,274,275,276,277,278,279,280,281,282",
	},
	"21": {
		City:     "宜蘭縣",
		Sections: "328,329,330,331,332,333,334,335,336,337,338,339",
	},
	"22": {
		City:     "台東縣",
		Sections: "341,342,343,344,345,346,347,348,349,350,351,352,353,354,355,356",
	},
	"23": {
		City:     "花蓮縣",
		Sections: "357,358,359,360,361,362,363,364,365,366,367,368,369",
	},
	"24": {
		City:     "澎湖縣",
		Sections: "283,284,285,286,287,288",
	},
	"25": {
		City:     "金門縣",
		Sections: "289,290,291,292,293,294",
	},
	"26": {
		City:     "連江縣",
		Sections: "22,23,24,25,256,257",
	},
}
