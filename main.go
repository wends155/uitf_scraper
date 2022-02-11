package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Fund struct {
	name    string
	fund_id int
}

func (f Fund) String() string {
	return fmt.Sprintf("name:%s/fund_id:%d", f.name, f.fund_id)
}

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("...Requesting...")
	})

	c.OnHTML("tbody", func(h *colly.HTMLElement) {
		funds := make([]Fund, 0)
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			td := make([]string, 0)
			el.ForEach("td", func(_ int, e *colly.HTMLElement) {

				if e.ChildAttr("a", "href") != "" {
					u, _ := url.Parse(e.ChildAttr("a", "href"))
					query, _ := url.ParseQuery(u.RawQuery)
					bank_id := query["bank_id"][0]
					fmt.Println(bank_id)
					td = append(td, query["fund_id"][0])
				}
				td = append(td, e.Text)
			})
			if strings.Contains(td[2], "* as of") {
				td[2] = td[2][:8]
			}
			fid, _ := strconv.Atoi(td[0])
			fmt.Println(td)

			f := Fund{name: td[1], fund_id: fid}
			funds = append(funds, f)

		})
		for _, v := range funds {
			fmt.Println(v)
		}

	})

	c.Visit("https://www.uitf.com.ph/daily_navpu.php?bank_id=31")

}
