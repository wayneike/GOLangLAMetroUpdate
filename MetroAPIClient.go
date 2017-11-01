package MetroAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

func constructInfo(buf *bytes.Buffer, minutes float32) {
	buf.WriteString(fmt.Sprintf("%.f minutes\n", minutes))
}

func reqHTTPGet(url string) []byte {
	apiClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

func getInfo(info input, output *string, finished chan bool) {
	url := (fmt.Sprintf("http://api.metro.net/agencies/lametro-rail/stops/%s/predictions/", info.StopID))

	body := reqHTTPGet(url)

	pred := items{}

	jsonErr := json.Unmarshal(body, &pred)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}

	var nBuffer, sBuffer, wBuffer bytes.Buffer

	sort.Sort(byMinutes(pred.Items))

	for i := 0; i < len(pred.Items); i++ {

		switch pred.Items[i].RunID {
		case "802_0_var0", "805_0_var0":
			constructInfo(&sBuffer, pred.Items[i].Minutes)
		case "802_1_var0":
			constructInfo(&nBuffer, pred.Items[i].Minutes)
		case "805_1_var0":
			constructInfo(&wBuffer, pred.Items[i].Minutes)
		}
	}

	const str = " bound trains arriving in\n"
	if sBuffer.Len() > 0 {
		*output += fmt.Sprintf("Union Station%s%s\n", str, sBuffer.String())
	}
	if nBuffer.Len() > 0 {
		*output += fmt.Sprintf("NoHo%s%s\n", str, nBuffer.String())
	}
	if wBuffer.Len() > 0 {
		*output += fmt.Sprintf("Wil/Western%s%s\n", str, wBuffer.String())
	}

	*output = fmt.Sprintf("--- %s Station ---\n%s", info.StationName, *output)

	finished <- true
}
