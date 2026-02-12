# VoiceLine App

## Context
The app uses:
* OpenAI to transcribe audio and extract structured data (generate the API Key)
* Google Sheets API to store the extracted structured data

## Setup
In order to get the app up & running you have to:
1. Create a .env file in the project root and add the required environement variables (See .env.example)
2. Configure the google sheet environement.

    a. In the Google Cloud console, create a new project and enable the Google Sheets API (APIs & Sevices > Library, search for Google Sheets API and enable it).
    
    b. Create a new Service Account and assign it the necessary permissions for read and write access. Then, generate the keys by selecting “Manage keys” and choosing the JSON format. Download the resulting .json file, rename it to `credentials.json`, and place it in the root directory of your project.

    c. Create the google sheet and share it with your Service Account email.

    d. Make sure to put the sheet Id in the .env file

3. Run `go build -o bin/voiceline ./cmd`
4. Run `./bin/voiceline`
5. Send an API request to `/upload` (Name the form-data field to `audio` and add the audio file).