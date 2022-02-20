# wason translate

## Usage

```
$ ./wason-translate
Using config file: $HOME/.wason-translate.yaml
Translate Power Point via Wason Translate API

Usage:
  wason-translate [flags]
  wason-translate [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Deletes a document.
  download    Gets the translated document associated with the given document ID.
  help        Help about any command
  list        Lists documents that have been submitted for translation.
  translate   Submit a document for translation.

Flags:
      --api_key string   Wason Translate API KEY (default "API KEY")
      --config string    config file (default is $HOME/.wason-translate.yaml)
  -h, --help             help for wason-translate
  -t, --toggle           Help message for toggle
      --url string       Wason Translate API URL (default "URL")
      --version string   Wason Translate API VERSION (default "VERSION")

Use "wason-translate [command] --help" for more information about a command.
```

1. Prepare `$HOME/.wason-translate.yaml`

    ```yaml
    api_key: "API KEY"
    url: "URL"
    version: 2018-05-01
    ```

2. Translate a file

    [![asciicast](https://asciinema.org/a/470291.svg)](https://asciinema.org/a/470291)

3. Download the translated file

    [![asciicast](https://asciinema.org/a/470292.svg)](https://asciinema.org/a/470292)

4. Delete the translated file

    [![asciicast](https://asciinema.org/a/470293.svg)](https://asciinema.org/a/470293)