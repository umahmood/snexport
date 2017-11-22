# snexport

snexport exports your notes in plaintext from the [Simplenote](https://simplenote.com/) 
desktop app. You can use this tool to automatically backup your notes from the 
terminal after an app sync. Each note is exported in plaintext format.

**Note**: snexport requires you to have the [Simplenote](https://simplenote.com/) 
desktop app installed, as it reads its database.

**Note**: You can manually export your notes from the Simplenote web UI.

# Installation

> $ go get github.com/umahmood/snexport


# Usage

Inorder to export your notes, we need to locate the Simplenote app database 
file **' Simplenote.storedata'**. On macOS, this file is located at:

- /Users/USER-NAME/Library/Containers/com.automattic.SimplenoteMac/Data/Library/Simplenote

> $ snexport /path/to/Simplenote.storedata /backups/simplenotes

Output:
```
Exported 224 files
Location /Users/username/backups/simplenotes
```

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
