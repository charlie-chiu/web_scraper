package scraper

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

const (
	rootURL      = "https://rent.591.com.tw/"
	itemsPerPage = 30
)

type Query struct {
	Region      int    `url:"region"`                // 地區 - 預設：`1`
	Section     string `url:"section,omitempty"`     // 鄉鎮 - 可選擇多個區域，例如：`section=7,4`
	Kind        int    `url:"kind"`                  // 租屋類型 - `0`：不限、`1`：整層住家、`2`：獨立套房、`3`：分租套房、`4`：雅房、`8`：車位，`24`：其他
	RentPrice   string `url:"rentprice,omitempty"`   // 租金 - `2`：5k - 10k、`3`：10k - 20k、`4`: 20k - 30k；或者可以輸入價格範圍，例如：`0,10000`
	Area        string `url:"area,omitempty"`        // 坪數格式 - `10,20`（10 到 20 坪）
	Order       string `url:"order"`                 // 貼文時間 - 預設使用刊登時間：`posttime`，或是使用價格排序：`money`
	OrderType   string `url:"orderType"`             // 排序方式 - `desc` 或 `asc`
	Sex         int    `url:"sex,omitempty"`         // 性別 - `0`：不限、`1`：男性、`2`：女性
	HasImg      string `url:"hasimg,omitempty"`      // 過濾是否有「房屋照片」 - ``：空值（不限）、`1`：是
	NotCover    string `url:"not_cover,omitempty"`   // 過濾是否為「頂樓加蓋」 - ``：空值（不限）、`1`：是
	Role        string `url:"role,omitempty"`        // 過濾是否為「屋主刊登」 - ``：空值（不限）、`1`：是
	Shape       string `url:"shape,omitempty"`       // 房屋類型 - `1`：公寓、`2`：電梯大樓、`3`：透天厝、`4`：別墅
	Pattern     string `url:"pattern,omitempty"`     // 格局單選 - `0`：不限、`1`：一房、`2``：兩房、`3`：三房、`4`：四房、`5`：五房以上
	PatternMore string `url:"patternMore,omitempty"` // 格局多選 - 參考「格局單選」，可以選多種格局，例如：`1,2,3,4,5`
	Floor       string `url:"floor,omitempty"`       // 樓層 - `0,0`：不限、`0,1`：一樓、`2,6`：二樓到六樓、`6,12`：六樓到十二樓、`12,`：十二樓以上
	Option      string `url:"option,omitempty"`      // 提供設備 - `tv`：電視、`cold`：冷氣、`icebox`：冰箱、`hotwater`：熱水器、`naturalgas`：天然瓦斯、`four`：第四台、`broadband`：網路、`washer`：洗衣機、`bed`：床、`wardrobe`：衣櫃、`sofa`：沙發。可選擇多個設備，例如：option=tv,cold
	Other       string `url:"other,omitempty"`       // 其他條件 - `cartplace`：有車位、`lift`：有電梯、`balcony_1`：有陽台、`cook`：可開伙、`pet`：可養寵物、`tragoods`：近捷運、`lease`：可短期租賃。可選擇多個條件，例如：other=cartplace,cook
	FirstRow    int    `url:"firstRow"`
}

func (q Query) URL() (string, error) {
	v, err := query.Values(q)
	if err != nil {
		return "", fmt.Errorf("query.Values error: %v", err)
	}

	return rootURL + "?" + v.Encode(), err
}

// NewQuery create a `Query` with default value.
func NewQuery() *Query {
	return &Query{
		Kind:      2,
		Region:    1,
		Section:   "0",
		RentPrice: "2",
		Order:     "posttime",
		OrderType: "desc",
		FirstRow:  0,
	}
}

//台中市小量試驗
var QueryMini = &Query{
	Region:  8,
	Section: "98,99,100",
	//Section:   "98,99,100,101",
	Kind:      0,
	RentPrice: "12000,15000",
	OrderType: "desc",
	Role:      "1",
	Sex:       0,
	FirstRow:  0,
}

// 台中市八區
var QueryTaiChung = &Query{
	Region:    8,
	Section:   "98,99,100,101,102,103,104,105",
	Kind:      0,
	Role:      "1",
	RentPrice: "0,100000",
	OrderType: "desc",
	Sex:       0,
	FirstRow:  0,
}
