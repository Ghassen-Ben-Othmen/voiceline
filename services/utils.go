package services

import "fmt"

func getPromptForDataExtraction(transcript string) string {
	return fmt.Sprintf(`
Extract structured sales information from this transcript.
Return a flat JSON object. Do not invent data.

Transcript:
%s
`, transcript)
}
