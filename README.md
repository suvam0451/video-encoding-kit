# video-encoding-kit
ffmpeg encoding worker with file transfer support (GDrive, OneDrive). Build uses fedora but the generated go binaries should be usable from anywhere.

## Google drive

Registering an app for Google drive API should give you a `credentials.json` file as shown below. Choose Installed/Desktop app when opted.

```json
{
  "installed": {
    "client_id": "XYZ"
    "project_id": "quickstart-XYZ"
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
