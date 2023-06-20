package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gsxhnd/garage/src/cmd"
	"github.com/inhies/go-bytesize"
)

var htmlSrc = `
<tr onmouseover="this.style.backgroundColor='#F4F9FD';this.style.cursor='pointer';" onmouseout="this.style.backgroundColor='#FFFFFF'" height="35px" style=" border-top:#DDDDDD solid 1px">
    <td width="70%" onclick="window.open('magnet:?xt=urn:btih:C0B1A4D395D6719B7C9442FA6E0DAF03B2B38469&dn=%40HALT029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:C0B1A4D395D6719B7C9442FA6E0DAF03B2B38469&dn=%40HALT029">@HALT029                 	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:C0B1A4D395D6719B7C9442FA6E0DAF03B2B38469&dn=%40HALT029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:C0B1A4D395D6719B7C9442FA6E0DAF03B2B38469&dn=%40HALT029">1.61GB                	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:C0B1A4D395D6719B7C9442FA6E0DAF03B2B38469&dn=%40HALT029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:C0B1A4D395D6719B7C9442FA6E0DAF03B2B38469&dn=%40HALT029">2023-06-18                	</a>
    </td>
</tr>
<tr onmouseover="this.style.backgroundColor='#F4F9FD';this.style.cursor='pointer';" onmouseout="this.style.backgroundColor='#FFFFFF'" height="35px" style=" border-top:#DDDDDD solid 1px">
    <td width="70%" onclick="window.open('magnet:?xt=urn:btih:9B158163E9B005D5FAD452D32D4F3001AF14519F&dn=HALT-029-C','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:9B158163E9B005D5FAD452D32D4F3001AF14519F&dn=HALT-029-C">
            HALT-029-C <a class="btn btn-mini-new btn-primary disabled" title="包含高清HD的磁力連結">高清</a>
            <a class="btn btn-mini-new btn-warning disabled" title="包含字幕的磁力連結">字幕</a>
        </a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:9B158163E9B005D5FAD452D32D4F3001AF14519F&dn=HALT-029-C','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:9B158163E9B005D5FAD452D32D4F3001AF14519F&dn=HALT-029-C">5.28GB                	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:9B158163E9B005D5FAD452D32D4F3001AF14519F&dn=HALT-029-C','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:9B158163E9B005D5FAD452D32D4F3001AF14519F&dn=HALT-029-C">2023-06-17                	</a>
    </td>
</tr>
<tr onmouseover="this.style.backgroundColor='#F4F9FD';this.style.cursor='pointer';" onmouseout="this.style.backgroundColor='#FFFFFF'" height="35px" style=" border-top:#DDDDDD solid 1px">
    <td width="70%" onclick="window.open('magnet:?xt=urn:btih:DF58F00F1DD171715439EE6FA63C035278465E42&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:DF58F00F1DD171715439EE6FA63C035278465E42&dn=halt-029">halt-029                 	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:DF58F00F1DD171715439EE6FA63C035278465E42&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:DF58F00F1DD171715439EE6FA63C035278465E42&dn=halt-029">1.41GB                	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:DF58F00F1DD171715439EE6FA63C035278465E42&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:DF58F00F1DD171715439EE6FA63C035278465E42&dn=halt-029">2023-06-17                	</a>
    </td>
</tr>
<tr onmouseover="this.style.backgroundColor='#F4F9FD';this.style.cursor='pointer';" onmouseout="this.style.backgroundColor='#FFFFFF'" height="35px" style=" border-top:#DDDDDD solid 1px">
    <td width="70%" onclick="window.open('magnet:?xt=urn:btih:68E445D5B2676B8C16A1CE72550D8C340E6013D2&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:68E445D5B2676B8C16A1CE72550D8C340E6013D2&dn=halt-029">
            halt-029 <a class="btn btn-mini-new btn-primary disabled" title="包含高清HD的磁力連結">高清</a>
        </a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:68E445D5B2676B8C16A1CE72550D8C340E6013D2&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:68E445D5B2676B8C16A1CE72550D8C340E6013D2&dn=halt-029">5.11GB                	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:68E445D5B2676B8C16A1CE72550D8C340E6013D2&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:68E445D5B2676B8C16A1CE72550D8C340E6013D2&dn=halt-029">2023-06-17                	</a>
    </td>
</tr>
<tr onmouseover="this.style.backgroundColor='#F4F9FD';this.style.cursor='pointer';" onmouseout="this.style.backgroundColor='#FFFFFF'" height="35px" style=" border-top:#DDDDDD solid 1px">
    <td width="70%" onclick="window.open('magnet:?xt=urn:btih:689A9037A5A27E31A05230404D8572CBDC8D4803&dn=HALT-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:689A9037A5A27E31A05230404D8572CBDC8D4803&dn=HALT-029">
            HALT-029 <a class="btn btn-mini-new btn-primary disabled" title="包含高清HD的磁力連結">高清</a>
        </a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:689A9037A5A27E31A05230404D8572CBDC8D4803&dn=HALT-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:689A9037A5A27E31A05230404D8572CBDC8D4803&dn=HALT-029">5.09GB                	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:689A9037A5A27E31A05230404D8572CBDC8D4803&dn=HALT-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:689A9037A5A27E31A05230404D8572CBDC8D4803&dn=HALT-029">2023-06-16                	</a>
    </td>
</tr>
<img src="asdads"/>
<tr onmouseover="this.style.backgroundColor='#F4F9FD';this.style.cursor='pointer';" onmouseout="this.style.backgroundColor='#FFFFFF'" height="35px" style=" border-top:#DDDDDD solid 1px">
    <td width="70%" onclick="window.open('magnet:?xt=urn:btih:A2DE179B414D270C6189FF330F0BB33C9E8AE12C&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:A2DE179B414D270C6189FF330F0BB33C9E8AE12C&dn=halt-029">
            halt-029 <a class="btn btn-mini-new btn-primary disabled" title="包含高清HD的磁力連結">高清</a>
        </a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:A2DE179B414D270C6189FF330F0BB33C9E8AE12C&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:A2DE179B414D270C6189FF330F0BB33C9E8AE12C&dn=halt-029">5.15GB                	</a>
    </td>
    <td style="text-align:center;white-space:nowrap" onclick="window.open('magnet:?xt=urn:btih:A2DE179B414D270C6189FF330F0BB33C9E8AE12C&dn=halt-029','_self')">
        <a style="color:#333" rel="nofollow" title="滑鼠右鍵點擊並選擇【複製連結網址】" href="magnet:?xt=urn:btih:A2DE179B414D270C6189FF330F0BB33C9E8AE12C&dn=halt-029">2023-06-16                	</a>
    </td>
</tr>`

type magnet struct {
	Name     string  `json:"name"`
	Link     string  `json:"link"`
	Size     float64 `json:"size"`
	Subtitle bool    `json:"subtitle"`
	HD       bool    `json:"hd"`
}

func main() {
	err := cmd.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}

	doc, _ := htmlquery.Parse(strings.NewReader("<table><tbody>" + htmlSrc + "</tbody></table>"))

	// list, _ := htmlquery.QueryAll(doc, "//tr/td/a[@href]")
	list, err := htmlquery.QueryAll(doc, "//tr")
	if err != nil {
		fmt.Println(err)
	}
	var mList = make([]magnet, 0)
	// fmt.Println("list: ", list)
	for _, n := range list {
		// fmt.Println(n)
		tdList, _ := htmlquery.QueryAll(n, "//td/a")
		var m = magnet{
			HD:       false,
			Subtitle: false,
		}
		for tdIndex, tdValue := range tdList {
			// fmt.Println(tdIndex)
			switch tdIndex {
			case 0:
				m.Link = htmlquery.SelectAttr(tdValue, "href")
				m.Name = htmlquery.InnerText(tdValue)
			default:
				if htmlquery.InnerText(tdValue) == "高清" {
					m.HD = true
				} else if htmlquery.InnerText(tdValue) == "字幕" {
					m.Subtitle = true
				} else {
					var sizeStr string = htmlquery.InnerText(tdValue)
					sizeStr = strings.Replace(sizeStr, " ", "", -1)
					sizeStr = strings.Replace(sizeStr, "\n", "", -1)
					sizeStr = strings.Replace(sizeStr, "\x09", "", -1)
					_, err := time.Parse("2006-01-02", sizeStr)
					if err != nil {
						b, err := bytesize.Parse(sizeStr)
						if err != nil {
							return
						}

						sizeStr = strings.Replace(b.Format("%.2f", "MB", false), "MB", "", -1)
						size, _ := strconv.ParseFloat(sizeStr, 64)
						m.Size = size
					}
				}
			}
		}
		mList = append(mList, m)
		fmt.Println(mList)
		// fmt.Println(htmlquery.SelectAttr(n, "href"))    // 打印磁力链接
		// fmt.Println(htmlquery.SelectAttr(n, "src"))     // 打印磁力链接
		// fmt.Println(htmlquery.InnerText(n.NextSibling)) // 打印文件大小
		// fmt.Println(htmlquery.SelectAttr(n, "href"))    // output @href value
	}
}
