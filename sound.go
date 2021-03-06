package main

import "github.com/pnegre/gogame"

const FREQUENCY = 22000

var sound_device *gogame.AudioDevice

func sound_init(psg *PSG) error {
	var err error
	sound_device, err = gogame.NewAudioDevice(FREQUENCY)
	if err != nil {
		return err
	}
	sound_device.SetCallback(func(data []int16) {
		for i := 0; i < len(data); i++ {
			data[i] = 0
		}
		psg.feedSamples(data)
		scc_feedSamples(data)

		// Limit maximum
		for i := 0; i < len(data); i++ {
			if data[i] > 32760 {
				data[i] = 32760
			}
			if data[i] < -32760 {
				data[i] = -32760
			}
		}
	})
	sound_device.Start()
	return nil
}

func sound_quit() {
	sound_device.Stop()
	sound_device.Close()
}
