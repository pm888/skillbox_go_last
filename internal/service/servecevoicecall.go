package service

import (
	"mymod/internal/data"
	"mymod/internal/method"
	"strings"
)

var StorageDataVoiceCall = make([]data.VoiceCallData, 0)
var counter int

func FileVoice() []data.VoiceCallData {
	file := method.ReadFile(data.FileVoiceRead)
	stringsTemp := strings.Split(string(file), "\n")
	for i := 0; i < len(stringsTemp)-1; i++ {
		voice := strings.Split(stringsTemp[i], ";")
		if len(voice) == 8 {
			for key, _ := range method.Alfa2Data {
				if key == voice[0] && method.IsValidVoiceProvider(voice[3]) {
					newPerson := data.VoiceCallData{
						Country:             voice[0],
						Bandwidth:           voice[1],
						ResponseTime:        voice[2],
						Provider:            voice[3],
						ConnectionStability: method.StringIntoFloat32(voice[4]),
						TTFB:                method.StringIntoInt(voice[5]),
						VoicePurity:         method.StringIntoInt(voice[6]),
						MedianOfCallsTime:   method.StringIntoInt(voice[7]),
					}
					StorageDataVoiceCall = append(StorageDataVoiceCall, newPerson)
					counter++
				}
			}
		}
	}

	return StorageDataVoiceCall
}
