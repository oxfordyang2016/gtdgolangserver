package models
import(
	"github.com/bradfitz/slice"
	"fmt"
	"time"

)

func  Sort_tasksbyday(tasks []Tasks) []Everyday {
	alldays:=make(map[string] []Tasks)
	for _,item :=range tasks{
	  // alldays[item.Plantime]=append(alldays[item.Plantime],item)
if item.Status!="unfinished"{
  if item.Status!="unfinish"{
  alldays[item.Finishtime]=append(alldays[item.Finishtime],item)}}
	}
	var alleverydays []Everyday
	var daybefore180119 Everyday
	var forgotten  Everyday
	for k,v := range alldays{
	 if k!="180119before"&&k!="forgotten"{
	   alleverydays =append(alleverydays,Everyday{k,v})
	}
	   if k=="180119before"{
		 daybefore180119 = Everyday{k,v}
	   }
	   if k=="forgotten"{
	   forgotten  = Everyday{k,v}
	   
}



}

//https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field

	slice.Sort(alleverydays, func(i, j int) bool {
return alleverydays[i].Name > alleverydays[j].Name
})
  if daybefore180119.Name !=""{
	alleverydays =append(alleverydays,daybefore180119)
  }
   if forgotten.Name !=""{
	alleverydays =append(alleverydays,forgotten)
   }


   return alleverydays
  }





  //get this days and days - 7 days
  //这是获取当前日期以及之前7天的
  func getweekday(datestring string) []string {	
	layout := "2006-01-02"
	//str := "2019-03-19"
	t, err := time.Parse(layout, datestring)
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	
	weekday := t.Weekday()
	fmt.Println(weekday)      // "Tuesday"
	fmt.Println(int(weekday)) 
	var count = int(weekday)-1
	if int(weekday) ==0{count=6}
	var daterange []string
	t = t.AddDate(0, 0, -count)	
	for i := 0; i < 7; i++ {
	daterange = append(daterange,t.AddDate(0, 0, i).Format("2006-01-02"))		
	}
	
	fmt.Println(daterange) 
	return daterange
	}
  


  //get this days and days - 7 days
  //这是获取当前这一周的数据的,if there only three days,it will return the three days
  func Getweekday() []string {
	//layout := "2019-01-02"
const layout = "2006/2/21"
loc, _ := time.LoadLocation("Asia/Shanghai")
	now :=  time.Now().In(loc)
	t,_ := time.Parse(layout,now.Format("2006/2/21"))
	fmt.Println(t)
	weekday := t.Weekday()
	fmt.Println(weekday)      // "Tuesday"
	fmt.Println(int(weekday))
	var count = int(weekday)-1
	if int(weekday) ==0{count = 7}
	var daterange []string
	// t = t.AddDate(0, 0, -count)
	for i := 0; i < count; i++ {
	daterange = append(daterange,t.AddDate(0, 0, -i).Format("060102"))
	}
	fmt.Println(daterange)
	return daterange
	}







//获取截止当前的时间
    // warn  :time format must same
	func Getmonthday() []string {
		layout := "2006/01/02"
		t, err := time.Parse(layout,time.Now().Format("2006/01/02"))
		fmt.Println(t)
		if err != nil {
				fmt.Println(err)
		}

		  currentYear, currentMonth, today := t.Date()
		  //currentLocation := t.Location()
		  loc, _ := time.LoadLocation("Asia/Shanghai")
		  firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
		  lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		  today_format := time.Date(currentYear, currentMonth, today, 0, 0, 0, 0, loc)


		  var daterange []string
		  daterange = append(daterange,firstOfMonth.Format("060102"))
		   day := firstOfMonth
				  for day !=today_format {
								 day =  day.AddDate(0, 0, 1)
								   daterange = append(daterange,day.Format("060102"))
				  }
		  fmt.Println(firstOfMonth)
		  fmt.Println(lastOfMonth)
  fmt.Println(daterange)
  return daterange
  }




  //获取截止当前的时间
    // warn  :time format must same
	func Getmonthallday() []string {
		layout := "2006/01/02"
		t, err := time.Parse(layout,time.Now().Format("2006/01/02"))
		fmt.Println(t)
		if err != nil {
				fmt.Println(err)
		}

		  currentYear, currentMonth, today := t.Date()
		  //currentLocation := t.Location()
		  loc, _ := time.LoadLocation("Asia/Shanghai")
		  firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
		  lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		  today_format := time.Date(currentYear, currentMonth, today, 0, 0, 0, 0, loc)
          fmt.Println(today_format)

		  var daterange []string
		  daterange = append(daterange,firstOfMonth.Format("060102"))
		   day := firstOfMonth
				  for day !=lastOfMonth {
								 day =  day.AddDate(0, 0, 1)
								   daterange = append(daterange,day.Format("060102"))
				  }
		  fmt.Println(firstOfMonth)
		  fmt.Println(lastOfMonth)
  fmt.Println(daterange)
  return daterange
  }