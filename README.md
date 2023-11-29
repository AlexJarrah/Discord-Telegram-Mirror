# Discord-Telegram-Mirror

Discord-Telegram-Mirror is a tool designed to mirror messages from Discord to Telegram. It offers flexible configuration options, allowing you to monitor specific channels, entire servers, or even whole accounts. You can fine-tune each channel or server setting to direct your messages to different Telegram chats and threads based on their origin.

## Installation

To install the tool, run the below commands in the location you want to install it:

```
git clone https://github.com/quo0001/Discord-Telegram-Mirror.git
cd Discord-Telegram-Mirror
```

## Usage

1. Input your Telegram and Discord credentials in the data/config.json file. The Discord token can be sourced from either a user or bot account.
2. Set rules for which servers or channels to monitor, along with their corresponding output chat and thread IDs. Note: If you want to mirror an entire account, you can use a wildcard ("\*") as the ID for either a guild or channel rule. All servers/channels without a specific rule will automatically be forwarded to the wildcard’s output chat.
3. Run the cmd/Discord-Telegram-Monitor/main.go file and watch your Telegram chats fill up with incoming messages!

## Disclaimer

Using a user account as a bot is against the Discord Terms of Service. By using this tool, you acknowledge and understand that the developer is not responsible for any consequences or actions taken against your account due to violation of the Discord Terms of Service.

This tool is not intended to publicize any private, exclusive, or sensitive content, and such behavior is not endorsed by the developer. Please use responsibly.
