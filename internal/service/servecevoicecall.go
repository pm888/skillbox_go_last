package service

import (
	"mymod/internal/data"
	"mymod/internal/method"
	"strings"
)

func FileVoice() ([]data.VoiceCallData, error) {
	var StorageDataVoiceCall = make([]data.VoiceCallData, 0)
	file, err := method.ReadFile(data.FileVoiceRead)
	if err != nil {
		return nil, err
	}
	stringsTemp := strings.Split(string(file), "\n")
	for i := 0; i < len(stringsTemp)-1; i++ {
		voice := strings.Split(stringsTemp[i], ";")
		if len(voice) == 8 {
			for key, _ := range method.Alfa2Data {
				if key == voice[0] && method.IsValidVoiceProvider(voice[3]) {
					connectionStability, err := method.StringIntoFloat32(voice[4])
					if err != nil {
						return nil, err
					}
					ttfb, err := method.StringIntoInt(voice[5])
					if err != nil {
						return nil, err
					}
					voicePurity, err := method.StringIntoInt(voice[6])
					if err != nil {
						return nil, err
					}
					medianOfCallsTime, err := method.StringIntoInt(voice[7])
					if err != nil {
						return nil, err
					}
					newPerson := data.VoiceCallData{
						Country:             voice[0],
						Bandwidth:           voice[1],
						ResponseTime:        voice[2],
						Provider:            voice[3],
						ConnectionStability: connectionStability,
						TTFB:                ttfb,
						VoicePurity:         voicePurity,
						MedianOfCallsTime:   medianOfCallsTime,
					}
					StorageDataVoiceCall = append(StorageDataVoiceCall, newPerson)
				}
			}
		}
	}

	return StorageDataVoiceCall, err
}
