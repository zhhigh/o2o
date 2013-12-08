package o2o

import(
	"github.com/PuerkitoBio/goquery"
)

type PM25 struct{

}

func NewPM25()*PM25{
   return &PM25{}
}

func (p *PM25)ReadFromPM25WebPage()string{
 	var doc *goquery.Document
	var e error
	if doc, e = goquery.NewDocument(TARGET_URL); e != nil {
		panic(e.Error())
	}

	sel2 := doc.Find(`script`).Text()
	matches := regexp.MustCompile(`var\slocalCityData\s\=\s{.*};`).FindAllString(sel2, -1)

	return matches[0]
	/*fileName := "pm.json"
	dstFile,_ := os.Create(fileName)

	defer dstFile.Close()

	dstFile.WriteString(matches[0])*/
}

