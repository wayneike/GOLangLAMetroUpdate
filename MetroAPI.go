//https://developer.metro.net/introduction/realtime-api-overview/realtime-api-returning-json/
//http://api.metro.net/agencies/lametro-rail/stops/80204/predictions/
//http://api.metro.net/agencies/lametro-rail/routes/802/runs/

//Also... https://www.metro.net/riding/nextrip/

package MetroAPI

import (
	"fmt"
	"time"
)

func main() {
	for {
		day := time.Now().Weekday()
		clock := time.Now()
		fmt.Printf("As of %s, %s\n\n", day, clock)

		var output1, output2 string
		finished1, finished2 := make(chan bool), make(chan bool)

		go getInfo(input{"80212", "Pershing Square"}, &output1, finished1)
		go getInfo(input{"80204", "Hollywood/Vine"}, &output2, finished2)

		<-finished1
		fmt.Println(output1)
		<-finished2
		fmt.Println(output2)

		fmt.Println("Press Enter to update info, Ctrl-c to exit...")
		fmt.Scanln() //wait for Enter key

		time.Sleep(time.Second / 2) //slow it down in case the user holds the key down
	}
}
