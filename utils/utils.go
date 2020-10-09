
package utils
// change this to define a new struct called pages
// with arbitary json data afterwards
type WikiData struct {
	Batchcomplete string `json:"batchcomplete"`
	Query struct {
		Pages map[string](interface{}) `json:"pages"`
	} `json:"query"`
}


func findTerm(term string, client *http.Client) (string, error) {
	if term == "" {
		term = "Ease of movement"
	}

	// see https://stackoverflow.com/questions/7185288/how-to-get-wikipedia-content-using-wikipedias-api
	wikiUrl := "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1"
	req, err := http.NewRequest("GET", wikiUrl, nil)
    if err != nil {
        log.Print(err)
        return "",err
    }

    q := req.URL.Query()
		q.Add("titles", term )
    req.URL.RawQuery = q.Encode()

		// fmt.Println(req.URL.String())
		resp, err := client.Do(req)
		if err != nil {
			return "",err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		wikiData := WikiData{}
		err = json.Unmarshal(body, &wikiData)
		if err != nil {
			return "", err
		}
		// iterate across each page
		// cant seem to get more than one page atm
		// might be redirect property
		extract := ""
		for _, v := range wikiData.Query.Pages {
			// https://stackoverflow.com/questions/25214036/getting-invalid-operation-mymaptitle-type-interface-does-not-support-in
			// type cast insertion
			md, ok := v.(map[string]interface{})
			if ok == false {
				fmt.Println("FAILURE")
				return "", err
			}
			if _, ok := md["extract"]; ok {
				//do something here
				extract = md["extract"].(string)
			} else {
				fmt.Println("no good value")
			}
		}
		return extract, nil
}