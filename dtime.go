package o2o

/*date:2013-12-14
  created by :zhhigh

*/
import(
	"time"
	"strconv"
)

type CTime struct{
	CurrentTime time.Time
	Year string
	Month string
	Day   string
	Hour  string
	Minute string
	Second string
	NanoSecond    string
}
var Ctime CTime

func init(){
   Ctime.CurrentTime = time.Now().Add(time.Hour*8)
   currentTime := Ctime.CurrentTime
   Ctime.Year = strconv.Itoa(currentTime.Year())
   month := currentTime.Month()
   Ctime.Month = GetMonth(month)

	Ctime.Day = strconv.Itoa(currentTime.Day())
	Ctime.Hour= strconv.Itoa(currentTime.Hour())
	Ctime.Minute=strconv.Itoa(currentTime.Minute())
	Ctime.Second=strconv.Itoa(currentTime.Second())
	Ctime.NanoSecond=strconv.Itoa(currentTime.Nanosecond())
}

/*func getCurrentTime()(time.Time){
	now := time.Now().Add(time.Hour*8)
	return now
} */

func GetMonth(m time.Month)string{
	var monstr string
	switch m{
	case time.January:
		monstr ="01"
	case time.February:
		monstr ="02"
	case time.March:
		monstr ="03"
	case time.April:
		monstr = "04"
	case time.May:
		monstr ="05"
	case time.June:
		monstr ="06"
	case time.July:
		monstr ="07"
	case time.August:
		monstr ="08"
	case time.September:
		monstr ="09"
	case time.October:
		monstr ="10"
	case time.November:
		monstr ="11"
	case time.December:
		monstr = "12"
	default:
		monstr=m.String()
	}
	return monstr
}


