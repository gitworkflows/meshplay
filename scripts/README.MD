1. Component Generation, along with spreadsheet update:
Points to consider when executing the script:

- The script references entries listed in `../../server/meshmodel/component_models.yaml` and generates components. At the time of generation, entries in the spreadsheet are updated.

- The mappings `NameToIndex` inside the script are sensitive and the order resembles the cols inside the spreadsheet.

- The script requires Google Spreadsheet API token as environment variable.
To set the environment variable in your system, run the below command.
`export CRED=($(cat <path to JSON API token> | base64 ))`

Usage:
To run the script execute following command 
`CRED=$CRED go run main.go https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub\?gid\=0\&single\=true\&output\=csv --system meshplay ../../server/meshmodel `

2. Component Updater
Uses a spreadsheet, which contains centralized information about MeshModel components and their metadata like color, icon, and so on. The script is used to update components metadata (svgs, icons etc) for Meshplay, Websites (khulnasoft.com, Meshplay.io), and Remote Provider.

Supported flags are:
   - --system
        defined type of system to update. Can be one of "meshplay", "docs", or "remote-provider".

   - --only-published
        Only handle components that have a value of "true" under the "Published?" column in spreadsheet.

To run the script execute following command 
1. To update docs on khulnasoft.com website:
`./main https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub\?gid\=0\&single\=true\&output\=csv --system docs khulnasoft/src/collections/integrations meshplay.khulnasoft.com/integrations docs/ --published-only`
2. To update component metadata for meshmodels:
`./main https://docs.google.com/spreadsheets/d/e/2PACX-1vSgOXuiqbhUgtC9oNbJlz9PYpOEaFVoGNUFMIk4NZciFfQv1ewZg8ahdrWHKI79GkKK9TbmnZx8CqIe/pub\?gid\=0\&single\=true\&output\=csv --system meshplay ../../server/meshmodel`