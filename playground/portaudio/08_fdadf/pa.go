package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"github.com/takuyaohashi/go-wav"
	"github.com/tetsuzawa/go-research/adflib/fdadf"
	"gonum.org/v1/gonum/floats"
	"log"
	"os"
	"time"
)

//numOutputChannels int, sampleRate float64, framesPerBuffer
const (
	RecordSeconds = 5

	NumInputChannels  = 1
	NumOutputChannels = 0
	SampleRate        = 48000
	FramesPerBuffer   = 1024

	FiltLen = 16
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	run()
}

var w1 *wav.Writer
var w2 *wav.Writer

func run() {
	f1, err := os.Create(`input.wav`)
	check(err)
	defer f1.Close()
	f2, err := os.Create(`filterd_input.wav`)
	check(err)
	defer f2.Close()

	err = portaudio.Initialize()
	if err != nil {
		fmt.Println(err)
		err = portaudio.Terminate()
		fmt.Println(err)
		check(err)
	}
	defer portaudio.Terminate()

	//////////////////////////////////
	h, err := portaudio.DefaultHostApi()
	check(err)
	paParam := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	paParam.Input.Channels = 1
	paParam.Output.Channels = 1

	fmt.Println("paParam.SampleRate", paParam.SampleRate)
	fmt.Println("paParam.FramesPerBuffer", paParam.FramesPerBuffer)
	paParam.FramesPerBuffer = FramesPerBuffer
	fmt.Println("paParam.FramesPerBuffer", paParam.FramesPerBuffer)

	fmt.Println("paParam.Input.Device.Name", paParam.Input.Device.Name)
	fmt.Println("paParam.Input.Device.MaxInputChannels", paParam.Input.Device.MaxInputChannels)
	fmt.Println("paParam.Input.Device.DefaultSampleRate", paParam.Input.Device.DefaultSampleRate)
	fmt.Println("paParam.Input.Device.HostApi", paParam.Input.Device.HostApi)
	fmt.Println("paParam.Input.Device.DefaultLowInputLatency", paParam.Input.Device.DefaultLowInputLatency)
	fmt.Println("paParam.Input.Device.DefaultHighInputLatency", paParam.Input.Device.DefaultHighInputLatency)

	fmt.Println("paParam.Output.Device.Name", paParam.Output.Device.Name)
	fmt.Println("paParam.Output.Device.MaxInputChannels", paParam.Output.Device.MaxInputChannels)
	fmt.Println("paParam.Output.Device.DefaultSampleRate", paParam.Output.Device.DefaultSampleRate)
	fmt.Println("paParam.Output.Device.HostApi", paParam.Output.Device.HostApi)
	fmt.Println("paParam.Output.Device.DefaultLowInputLatency", paParam.Output.Device.DefaultLowInputLatency)
	fmt.Println("paParam.Output.Device.DefaultHighInputLatency", paParam.Output.Device.DefaultHighInputLatency)

	//open stream
	stream, err := portaudio.OpenStream(paParam, callback)
	check(err)
	defer stream.Close()
	fmt.Println("info: ", stream.Info())

	//make input.wav
	p := wav.WriterParam{
		SampleRate:    48000,
		BitsPerSample: 16,
		NumChannels:   1,
		AudioFormat:   1,
	}
	w1, err = wav.NewWriter(f1, p)
	defer w1.Close()
	w2, err = wav.NewWriter(f2, p)
	defer w2.Close()

	fmt.Println("recording...")
	err = stream.Start()
	check(err)
	time.Sleep(5 * time.Second)
	err = stream.Stop()
	check(err)
	fmt.Println("end!!")
}

const (
	mu    = 0.000000000000000001
	order = 8
	eps   = 1e-5
)

//var afBuf = make([]float64, FiltLen)
var af, _ = fdadf.NewFiltFBLMS(FramesPerBuffer, mu, "zeros")

func callback(inBuf, outBuf []int16) {
	///*
	inBufF := Int16sToFloat64s(inBuf)
	outBufF := make([]float64, FramesPerBuffer)
	af.Adapt(inBufF, inBufF)
	predicted := af.Predict(inBufF)
	copy(outBufF, inBufF)
	floats.Sub(outBufF, predicted)
	copy(outBuf, Float64sToInt16s(outBufF))
	for i := range inBuf {
		fmt.Println(outBuf[i], inBuf[i], float64(inBuf[i]), predicted[i], 0.9*predicted[i], int16(predicted[i]))
	}
	//fmt.Println(afBuf)
	//*/

	//copy(outBuf, inBuf)
	w1.WriteSamples(inBuf)
	w2.WriteSamples(outBuf)
}
