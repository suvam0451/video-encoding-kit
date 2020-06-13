# video-encoding-kit

ffmpeg encoding worker with file transfer support (GDrive, OneDrive). Build uses fedora but the generated go binaries should be usable from anywhere.

### Overview

Simplified command-line access to drive resources. Built keeping containers and flexibility in mind.

Here are some example use cases :-

Uploading a folder to a folder on the cloud

```sh
drivekit gdrive upload -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./uplink
```

Downloading all files in a cloud folder to a local folder _(specified with -d)_

```sh
drivekit gdrive download -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./downlink
```

## Usage

- You can clone the repo and test the client locally from source.
- Download the linux binary and save in to /bin/ in your docker image and use it from there.

```sh
# For linux
wget https://github.com/suvam0451/video-encoding-kit/releases/latest/download/drivekit
```

## Guides for beginners

### Google drive users

<details>
<summary>Setting up the app !!!</summary>

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

</details>

<details>
<summary>Security and permissions 101 !!!</summary>
</details>

<details>
<summary>Additional directives !!!</summary>
</details>

## Testing locally

You can run this locally, if you have the golang tools. Run

```go
go run main.go client_gdrive.go
```

- You would need the path of the credential file as `GDRIVE_APP_CREDENTIAL_FILE_LOCATION` (environment variable).
- The path to token.json is given as first argument (assuming you wuld have more than 1 drive)
- When using **scanfolder** argument, the file **scanresult.json** will be generated with {name,id, fileextension}. You can use this later for the **getfile** argument.

## Anonymous access in containers 🕵🏿

On public servers, you would want to hide your auth codes. This is usually done by environment variables. For the **credentials.json**, you would need to setup the following environment variables.

- CLIEND_ID
- PROJECT_ID
- CLIENT_SECRET

If all these variables are found and are _non-empty_, then the **credentials.json** file would be generated as follows **in the same directory**

```json
// This file will be generated from environment variables, if present...
{
  "installed": {
    "client_id": "{CLIENT_ID}",
    "project_id": "{PROJECT_ID}",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_secret": "{CLIENT_SECRET}",
    "redirect_uris": ["urn:ietf:wg:oauth:2.0:oob", "http://localhost"]
  }
}
```

Here are the commands for the generators (consecutive tokens are replaced). They will

- always have the same name {credentials.json, token.json}
- be on the current directory.

```bash
drivekit gdrive generate credentials
drivekit gdrive generate token
```

For the token files, it's a little different. You would need to pass the name of the environment variables instead.
This way, you can have multiple pairs of key pairs in multiple variables.

```
drivekit gdrive generate token -
```

Feel free to open issues for additional feature requests/bugs. Thank you.

---

## Note to self 📝

```sh
cat ~/GH_TOKEN.txt | docker login docker.pkg.github.com -u suvam0451 --password-stdin
# build + push podman packages

# Copying binary to /bin/ -->
yes | sudo cp -rf drivekit /bin/
```
