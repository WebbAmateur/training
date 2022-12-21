package main

func main() {

	webbMap := make(map[string]Webb)
	webbMap["Todd"] = Webb{Index: 0, First: "James", Middle: "Todd", Last: "Webb", Day: 28, Month: 6, Year: 1960}
	webbMap["Alison"] = Webb{Index: 1, First: "Alison", Middle: "Pack", Last: "Webb", Day: 28, Month: 3, Year: 1962}
	webbMap["Alex"] = Webb{Index: 2, First: "Alexander", Middle: "Todd", Last: "Webb", Day: 17, Month: 6, Year: 1990}

	// webbs := []Webb{
	// 	{Index: 0, First: "James", Middle: "Todd", Last: "Webb", Day: 28, Month: 6, Year: 1960},
	// 	{Index: 1, First: "Alison", Middle: "Pack", Last: "Webb", Day: 28, Month: 3, Year: 1962},
	// 	{Index: 2, First: "Alexander", Middle: "Todd", Last: "Webb", Day: 17, Month: 6, Year: 1990},
	// }
	//todd := Webb{First: "James", Middle: "Todd", Last: "Webb", Day: 28, Month: 6, Year: 1960}

	// for index := range webbs {
	// 	println(webbs[index].String())
	// }

	for _, value := range webbMap {
		println(value.String())
	}

	println(webbMap["Todd"].String())

}
