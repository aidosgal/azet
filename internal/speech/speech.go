package speech

import (
  "context"
  "log"
  "fmt"
  "bytes"
  "github.com/faiface/beep"
  "github.com/faiface/beep/mp3"
  "github.com/faiface/beep/speaker"
  texttospeech "cloud.google.com/go/texttospeech/apiv1"
  texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
  speech "cloud.google.com/go/speech/apiv1"
  speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
  "github.com/gordonklaus/portaudio"
  "time"
)

type readCloser struct {
    *bytes.Reader
}

func (rc *readCloser) Close() error {
    return nil
}

func TextToSpeech(text string) {
    ctx := context.Background()
    client, err := texttospeech.NewClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    defer client.Close()

    req := &texttospeechpb.SynthesizeSpeechRequest{
        Input: &texttospeechpb.SynthesisInput{
            InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
        },
        Voice: &texttospeechpb.VoiceSelectionParams{
            LanguageCode: "en-US",
            SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
        },
        AudioConfig: &texttospeechpb.AudioConfig{
            AudioEncoding: texttospeechpb.AudioEncoding_MP3,
        },
    }

    resp, err := client.SynthesizeSpeech(ctx, req)
    if err != nil {
        log.Fatalf("Failed to synthesize speech: %v", err)
    }

    // Wrap the bytes.Reader with readCloser
    rc := &readCloser{bytes.NewReader(resp.AudioContent)}

    // Play the audio content directly
    streamer, format, err := mp3.Decode(rc)
    if err != nil {
        log.Fatalf("Failed to decode MP3: %v", err)
    }
    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
    speaker.Play(beep.Seq(streamer, beep.Callback(func() {
        fmt.Println("Playback finished")
    }))) 
}

func init() {
    portaudio.Initialize()
}

func int16ToByte(data []int16) []byte {
    buf := new(bytes.Buffer)
    for _, v := range data {
        buf.WriteByte(byte(v))
        buf.WriteByte(byte(v >> 8))
    }
    return buf.Bytes()
}

func SpeechToText() string {
    ctx := context.Background()
    client, err := speech.NewClient(ctx)
    if err != nil {
        log.Printf("Failed to create client: %v", err)
        return "Error: Failed to create client"
    }
    defer client.Close()

    stream, err := client.StreamingRecognize(ctx)
    if err != nil {
        log.Printf("Failed to create streaming client: %v", err)
        return "Error: Failed to create streaming client"
    }

    if err := stream.Send(&speechpb.StreamingRecognizeRequest{
        StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
            StreamingConfig: &speechpb.StreamingRecognitionConfig{
                Config: &speechpb.RecognitionConfig{
                    Encoding:        speechpb.RecognitionConfig_LINEAR16,
                    SampleRateHertz: 16000,
                    LanguageCode:    "en-US",
                },
                InterimResults: true,
            },
        },
    }); err != nil {
        log.Printf("Failed to send streaming config: %v", err)
        return "Error: Failed to send streaming config"
    }

    audioStream, err := portaudio.OpenDefaultStream(1, 0, 16000, 1024, func(in []int16) {
        if err := stream.Send(&speechpb.StreamingRecognizeRequest{
            StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
                AudioContent: int16ToByte(in),
            },
        }); err != nil {
            log.Printf("Failed to send audio: %v", err)
        }
    })
    if err != nil {
        log.Printf("Failed to open PortAudio stream: %v", err)
        return "Error: Failed to open PortAudio stream"
    }
    defer audioStream.Close()

    if err := audioStream.Start(); err != nil {
        log.Printf("Failed to start PortAudio stream: %v", err)
        return "Error: Failed to start PortAudio stream"
    }
    defer audioStream.Stop()

    var resultText string
    done := make(chan struct{})

    go func() {
        for {
            resp, err := stream.Recv()
            if err != nil {
                log.Printf("Failed to receive audio: %v", err)
                close(done)
                return
            }
            for _, result := range resp.Results {
                for _, alt := range result.Alternatives {
                    resultText = alt.Transcript
                }
            }
        }
    }()

    time.Sleep(10 * time.Second) // Adjust the duration as needed
    close(done)
    return resultText 
}
