package bot

import (
	"github.com/enescakir/emoji"
)

const (
	creatorChatID int64 = 5785150199
)

var (
	isBotRunning bool
)

var (
	GreenHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
	CactusEmoji     = emoji.Sprintf("%v", emoji.Cactus)
	IdkEmoji        = emoji.Sprintf("%v", emoji.OpenHands)
	StopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	InfinityEmoji   = emoji.Sprintf("%v", emoji.Infinity)
	InfoEmoji       = emoji.Sprintf("%v", emoji.Information)
	OkEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
)

var (
	IdkMsg            = "[ ? ] I don't know what you mean by that\n[ ? ] Please, use /help command to see the list of available commands"
	BugReportMsg      = "[ ! ] Bug report from user @%s:\n[ ! ] %s"
	ThxMsg            = GreenHeartEmoji + " Thanks for your bug report!"
	PlsDescribeMsg    = CactusEmoji + " Please describe the problem:"
	StartMsg          = "Hello, I'm Kanojo!\nI can generate images from your text prompts.\n\nTo see how to use me, use /help"
	StopMsg           = StopEmoji + " Kanojo has been stopped"
	AlreadyRunningMsg = IdkEmoji + " Kanojo is already running"
	AlreadyStoppedMsg = IdkEmoji + " Kanojo is already stopped"
	YourPromptMsg     = InfinityEmoji + " Your prompt:"
	HelpMsg           = InfoEmoji + " Available commands:\n\n/help - show this message\n/generator - generate image from your prompt\n/bug_report - report a bug\n/stop - stop kanojo"
)
