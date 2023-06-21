

### Basic Usage

Build the [Plain Text](https://app.slack.com/block-kit-builder/#%7B%22blocks%22:%5B%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22This%20is%20a%20plain%20text%20section%20block.%22,%22emoji%22:true%7D%7D%5D%7D) example

```go
package main

func main() {
	text := slack.NewPlainText("This is a plain text section block.").EnableEmoji()
	section := slack.NewSection().SetText(text)
	basic := slack.NewMessage().AddBlock(section)

	fmt.Println(basic)
```
Outputs:

```json
{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "plain_text",
				"text": "This is a plain text section block.",
				"emoji": true
			}
		}
	]
}
```

This can be condensed

```go

package main

import (
	"fmt"
	"github.com/mresl/slack"
)

func main() {
	text := slack.NewPlainText("This is a plain text section block.").EnableEmoji()
	fmt.Println(slack.NewMessage().AddBlock(text.Section()))
}
```

### Real World Example

Composing the example message:

```go
package main

import (
	"fmt"
	slack "github.com/jeremyforan/go-flocks-of-blocks"
)

func main() {

	// create a new message
	msg := slack.NewMessage()

	// Add a header
	header := slack.NewHeader("Device Summary")
	msg = msg.AddBlock(header)

	// Add some info
	info := slack.NewSection().AddMrkdownField("*IP:* 192.168.0.1").AddMrkdownField("*Area:* basement")
	msg = msg.AddBlock(info)

	// Add some more info but in a single line
	msg = msg.AddBlock(slack.NewSection().AddMrkdownField("*Hardware:* Linksys WRT-54G").AddMrkdownField("*Uptime:* 7 Days, 3 Months"))

	// Add the info message to
	ssid := slack.NewSection().AddMrkdownField("*SSID:* FreshPrinceOfDonair")
	msg = msg.AddBlock(ssid)

	// make a "reset" button
	resetButton := slack.NewButton("Reboot Device", "actionId-0").SetValue("click_me_123")

	// Let's assume we want to add a note based on some arbitrary bool value
	rfIssue := true
	if rfIssue {
		note := slack.NewPlainText("*high levels of RF interference detected consider scan")
		msg = msg.AddBlock(note.Context())

		// We want to add the Danger styleing to the button due to the 'issue'
		resetButton = resetButton.MakeStyleDanger()
	}

	// Add the reset button to the message
	msg = msg.AddBlock(resetButton.Actions())

	// Generate a link that paste the body into the Slack interactive Block Kit Builder for validation
	fmt.Println(msg.GenerateKitBuilderUrl())
}
```