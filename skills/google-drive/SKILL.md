---
name: google-drive
description: Access Google Drive via Google API using the google_api tool
homepage: https://developers.google.com/drive/api/v3/reference
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "📁"
---

# Google Drive

Manage files and folders using the `google_api` tool with Google Drive API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-drive",
  name="default",
  mode="oauth"
)
```

Or use a generic Google profile:

```bash
auth_profile_add(
  provider="google",
  name="work",
  mode="oauth"
)
```

## Common Operations

### List Files

```bash
google_api(
  service="drive",
  endpoint="/files",
  method="GET",
  params={
    "pageSize": "20",
    "fields": "files(id,name,mimeType,modifiedTime,size)"
  }
)
```

### Search Files

```bash
google_api(
  service="drive",
  endpoint="/files",
  method="GET",
  params={
    "q": "name contains 'report'",
    "pageSize": "20",
    "fields": "files(id,name,mimeType,modifiedTime)"
  }
)
```

### List Files in Folder

```bash
google_api(
  service="drive",
  endpoint="/files",
  method="GET",
  params={
    "q": "'FOLDER_ID' in parents",
    "pageSize": "50",
    "fields": "files(id,name,mimeType,modifiedTime,size)"
  }
)
```

### Get File Metadata

```bash
google_api(
  service="drive",
  endpoint":"/files/FILE_ID",
  method="GET",
  params={
    "fields": "id,name,mimeType,description,createdTime,modifiedTime,size,webViewLink"
  }
)
```

### Create Folder

```bash
google_api(
  service="drive",
  endpoint":"/files",
  method="POST",
  body={
    "name": "New Folder",
    "mimeType": "application/vnd.google-apps.folder",
    "parents": ["PARENT_FOLDER_ID"]
  }
)
```

### Update File

```bash
google_api(
  service="drive",
  endpoint":"/files/FILE_ID",
  method="PATCH",
  body={
    "name": "New File Name",
    "description": "Updated description"
  }
)
```

### Move File to Folder

```bash
google_api(
  service="drive",
  endpoint":"/files/FILE_ID",
  method="PATCH",
  body={
    "addParents": "NEW_PARENT_FOLDER_ID",
    "removeParents": "OLD_PARENT_FOLDER_ID"
  }
)
```

### Copy File

```bash
google_api(
  service="drive",
  endpoint":"/files/FILE_ID/copy",
  method="POST",
  body={
    "name": "Copy of File",
    "parents": ["PARENT_FOLDER_ID"]
  }
)
```

### Delete File

```bash
google_api(
  service="drive",
  endpoint":"/files/FILE_ID",
  method="DELETE"
)
```

### Empty Trash

```bash
google_api(
  service="drive",
  endpoint":"/files/trash:empty",
  method="DELETE"
)
```

### List Shared Drives

```bash
google_api(
  service="drive",
  endpoint="/drives",
  method="GET",
  params={
    "pageSize": "20"
  }
)
```

## Query Syntax (q parameter)

### Search by Name
```
name = 'exact name'
name contains 'partial name'
```

### Search by Type
```
mimeType = 'application/pdf'
mimeType = 'application/vnd.google-apps.folder'
mimeType = 'application/vnd.google-apps.spreadsheet'
mimeType = 'application/vnd.google-apps.document'
mimeType = 'application/vnd.google-apps.presentation'
```

### Search by Date
```
modifiedTime > '2024-01-01T00:00:00'
createdTime < '2023-12-31T23:59:59'
```

### Search in Folder
```
'FOLDER_ID' in parents
```

### Combined Queries
```
name contains 'report' and mimeType = 'application/pdf'
name contains 'budget' and modifiedTime > '2024-01-01'
mimeType = 'application/vnd.google-apps.folder' and trashed = false
```

## MIME Types

### Google Workspace Files
- `application/vnd.google-apps.folder` - Folder
- `application/vnd.google-apps.document` - Google Docs
- `application/vnd.google-apps.spreadsheet` - Google Sheets
- `application/vnd.google-apps.presentation` - Google Slides
- `application/vnd.google-apps.form` - Google Forms
- `application/vnd.google-apps.drawing` - Google Drawings

### Common File Types
- `application/pdf` - PDF
- `application/vnd.openxmlformats-officedocument.wordprocessingml.document` - Word
- `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet` - Excel
- `application/vnd.openxmlformats-officedocument.presentationml.presentation` - PowerPoint
- `image/jpeg` - JPEG image
- `image/png` - PNG image
- `text/plain` - Plain text
- `application/zip` - ZIP archive

## File Fields

Common fields to request:
- `id` - File ID
- `name` - File name
- `mimeType` - MIME type
- `description` - Description
- `createdTime` - Creation time
- `modifiedTime` - Last modified time
- `size` - File size (for binary files)
- `parents` - Parent folder IDs
- `webViewLink` - Web view URL
- `webContentLink` - Direct download URL
- `thumbnailLink` - Thumbnail URL
- `trashed` - Whether in trash
- `shared` - Whether shared

## Special Folders

- `root` - My Drive root
- `appDataFolder` - App data folder (hidden)

## Using Specific Profiles

```bash
google_api(
  service="drive",
  endpoint="/files",
  method="GET",
  profile="work"
)
```

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-drive", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## API Reference

Full API documentation: https://developers.google.com/drive/api/v3/reference

Note: This is a read-only guide. File upload/download operations may require additional steps.
