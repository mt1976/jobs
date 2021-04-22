package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/mt1976/common"
	tools "github.com/mt1976/mwtgostringtools"
)

const constRateLen = 8

var funcName = ""

type fxRate struct {
	ccyPair   string
	bidRate   float64
	bidString string
	askRate   float64
	askString string
}

type fxRateCard struct {
	fxRates []fxRate
}

func getFXrate(inCCYpair string) fxRate {
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	inString := string(html)

	searchString := "\"bidPrice\":\""
	searchString2 := "\",\"askPrice\":\""
	searchString3 := "\",\"bidSize\":\""

	bidPriceStart := tools.FindStartPos(inString, searchString)
	bidPriceStop := tools.FindPos(inString, searchString2)
	askPriceStart := tools.FindStartPos(inString, searchString2)
	askPriceStop := tools.FindPos(inString, searchString3)
	if askPriceStop == -1 {
		askPriceStop = askPriceStart + 7
	}

	thisRate.bidRate, _ = strconv.ParseFloat(inString[bidPriceStart:bidPriceStop], 64)
	thisRate.bidString = tools.TruncateString(inString[bidPriceStart:bidPriceStop], constRateLen)
	//fmt.Println("bidPrice=", bidPrice)

	thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)
	thisRate.askString = tools.TruncateString(inString[askPriceStart:askPriceStop], constRateLen)
	//fmt.Println("askPrice=", askPrice)
	log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)

	return thisRate
}

var wg sync.WaitGroup

func getASYCFXrate(inCCYpair string, rateChan chan fxRate) {
	defer wg.Done()
	common.Snooze()
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	inString := string(html)

	searchString := "\"bidPrice\":\""
	searchString2 := "\",\"askPrice\":\""
	searchString3 := "\",\"bidSize\":\""

	bidPriceStart := tools.FindStartPos(inString, searchString)
	bidPriceStop := tools.FindPos(inString, searchString2)
	askPriceStart := tools.FindStartPos(inString, searchString2)
	askPriceStop := tools.FindPos(inString, searchString3)
	if askPriceStop == -1 {
		askPriceStop = askPriceStart + 7
	}

	thisRate.bidRate, _ = strconv.ParseFloat(inString[bidPriceStart:bidPriceStop], 64)
	thisRate.bidString = tools.TruncateString(inString[bidPriceStart:bidPriceStop], constRateLen)
	//fmt.Println("bidPrice=", bidPrice)

	thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)
	thisRate.askString = tools.TruncateString(inString[askPriceStart:askPriceStop], constRateLen)
	//fmt.Println("askPrice=", askPrice)

	rateChan <- thisRate

	log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)

}

func buildRateCard() fxRateCard {
	var rateCard fxRateCard
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("AUDUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURGBP"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURJPY"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("GBPUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("NZDUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDCAD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDCHF"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDHKD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDJPY"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDSGD"))
	return rateCard
}

func buldFXRVRates(rateCard fxRateCard) string {
	outputString := ""
	for _, s := range rateCard.fxRates {
		//fmt.Println(i, s)
		//fmt.Println(i, s)
		//fmt.Sprint(s)
		abc := fmt.Sprintf("s%ssptD%11sD%11s", s.ccyPair, s.bidString, s.askString)
		//fmt.Println(abc)
		outputString = outputString + abc + "\n"
	}
	log.Println("\n", outputString)
	return outputString
}

func deliverRVData(name string, record string) {

	f, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(record)

	if err2 != nil {
		log.Fatal(err2)
	}

}

func RunJobFXSPOT(actionType string) {
	funcName = "RunJobFXSPOT"
	logit(funcName, actionType, "*** REFRESH RATES ***")
	rateCard := buildRateCard()
	//log.Println(rateCard, len(rateCard.fxRates))
	//fmt.Println(rateCard, len(rateCard))
	logit(funcName, actionType, "*** BUILD RV RATES ***")
	outputString := buldFXRVRates(rateCard)
	logit(funcName, actionType, "*** DELIVER RATES ***")
	deliverRVData("RVMARKET", outputString)
	logit(funcName, actionType, "*** DONE ***")
}
