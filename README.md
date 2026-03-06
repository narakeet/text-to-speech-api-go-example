# Narakeet Text to Speech Streaming API example in Go

This repository provides a quick example demonstrating how to access the Narakeet [streaming Text to Speech API](https://www.narakeet.com/docs/automating/text-to-speech-api/) from Go.

The example sends a request to generate audio from text and saves it to output.mp3 in the local directory.

Note that Narakeet also has a text to speech API to generate long content, suitable for larger conversion tasks.

## Prerequisites

This example works with Go 1.22 and later. You can run it inside Docker (then it does not require a local Go installation), or on a system with Go 1.22 or later.

## Running the example

1. set and export a local environment variable called `NARAKEET_API_KEY`, containing your API key, or alternatively edit [tts.go](tts.go) and add your API key on line 14.
2. Optionally modify the voice and text variables in `tts.go`, which control the text to speech synthesis voice and the text sent to the API for synthesis.
2. To run inside docker, execute `make run`
3. Or to run outside docker, on a system with `go` command line, execute `go run tts.go`

## More information

Check out <https://www.narakeet.com/docs/automating/text-to-speech-api/> for more information on the Narakeet Text to Speech API
