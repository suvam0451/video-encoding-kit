# video-encoding-kit

ffmpeg encoding worker with file transfer support (GDrive, OneDrive). Build uses fedora but the generated go binaries should be usable from anywhere.

## Google drive

Registering an app for Google drive API should give you a `credentials.json` file as shown below. Choose Installed/Desktop app when opted.

```json
{
  "installed": {
    "client_id": "XYZ",
    "project_id": "quickstart-XYZ",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_secret": "XYZ",
    "redirect_uris": ["urn:ietf:wg:oauth:2.0:oob", "http://localhost"]
  }
}
```

Also, you would get the following `token.json` file when you provide your application access to your drive(s).

```json
{
  "access_token": "XYZ",
  "refresh_token": "XYZ",
  "scope": "https://www.googleapis.com/auth/drive",
  "token_type": "Bearer",
  "expiry_date": 1589726976058
}
```

## Testing locally

You can run this locally, if you have the golang tools. Run

```go
go run main.go client_gdrive.go
```

- You would need the path of the credential file as `GDRIVE_APP_CREDENTIAL_FILE_LOCATION` (environment variable).
- The path to token.json is given as first argument (assuming you wuld have more than 1 drive)
- When using **scanfolder** argument, the file **scanresult.json** will be generated with {name,id, fileextension}. You can use this later for the **getfile** argument.

Possible queries

```powershell
go run main.go gdrive -i "1qV-5YmODxtDVJGhXFq3RTC-E2NkLmL0b" -q "listdir"   # Lists files in folder. See below for how to get folder id
go run main.go gdrive -i "1qV-5YmODxtDVJGhXFq3RTC-E2NkLmL0b" -q "getfile"   # Lists files in folder. See below for how to get folder id
gdrive scanfolder id          # see below for how to get folder id
gdrive getfile id             # see below for how to get file id
gdrive getfile id name        # get the file with specific name
```
