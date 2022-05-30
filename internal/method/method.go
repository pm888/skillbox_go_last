package method

import (
	"io/ioutil"
	"log"
	"math"
	"mymod/internal/data"
	"net/http"
	"strconv"
	"strings"
)

var Alfa2Data = make(map[string]string)

func ReadFile(nameFile string) []byte {
	file, err := ioutil.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func StringIntoInt(str string) (chInt int) {
	chInt, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Bad string", err)
	}
	return
}

func StringIntoFloat32(str string) float32 {
	chFloat32, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Fatal("Bad string", err)
	}

	return float32(chFloat32)
}

func IsValidSmsMmsProvider(name string) bool {
	_, ok := data.ValidProviders[name]
	return ok
}

func IsValidVoiceProvider(name string) bool {
	_, ok := data.ValidProvidersVoice[name]
	return ok
}

func IsValidEmailProvider(name string) bool {
	_, ok := data.ValidProvidersEmail[name]
	return ok
}

func FileIntoMap() {
	file := ReadFile(data.FileNameAlpha2)
	sliceIntoFile := strings.Split(string(file), "\n")
	for i := 0; i < len(sliceIntoFile)-1; i++ {
		str := sliceIntoFile[i]
		sliceCod := strings.Split(str, ".")
		Alfa2Data[sliceCod[1]] = sliceCod[0]

	}

}

func GetBody(url string) []byte {
	var body []byte
	var err error
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	} else {
		log.Fatal(resp.StatusCode, err)
	}
	return body
}

func Interpretation(sliceFile []byte) float64 {
	total := 0.0
	for i := len(sliceFile) - 1; i >= 0; i-- {
		if sliceFile[i] != 48 {
			ch := 2.0
			position := (len(sliceFile) - 1) - i
			n := math.Pow(ch, float64(position))
			total = total + n
		} else {
			continue
		}

	}
	return total

}
