package helpers

import (
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"mymod/internal/data"
)

var Alfa2Data = make(map[string]string)

func StringIntoInt(str string) (int, error) {
	chInt, err := strconv.Atoi(str)
	return chInt, err
}

func StringIntoFloat32(str string) (float32, error) {
	chFloat32, err := strconv.ParseFloat(str, 32)
	return float32(chFloat32), err
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

func ReadFileIntoMap() error {
	file, err := os.ReadFile(data.FileNameAlpha2)
	if err != nil {
		return err
	}
	sliceIntoFile := strings.Split(string(file), "\n")
	for i := 0; i < len(sliceIntoFile)-1; i++ {
		str := sliceIntoFile[i]
		sliceCod := strings.Split(str, ".")
		Alfa2Data[sliceCod[1]] = sliceCod[0]
	}
	return err

}

func GetBody(url string) ([]byte, error) {
	var body []byte
	var err error
	resp, err := http.Get(url)
	if resp.StatusCode == http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
	} else {
		log.Fatal(resp.StatusCode, err)
	}
	return body, err
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
