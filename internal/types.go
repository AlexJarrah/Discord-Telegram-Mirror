package internal

type Configuration struct {
	Credentials Credentials `json:"credentials"`
	Rules       Rules       `json:"rules"`
}

type Credentials struct {
	DiscordToken  string `json:"discordToken"`
	TelegramToken string `json:"telegramToken"`
}

type Rules struct {
	Guilds   []Rule `json:"guilds"`
	Channels []Rule `json:"channels"`
}

type Rule struct {
	ID     string `json:"id"`
	Output Output `json:"output"`
}

type Output struct {
	ChatID   string `json:"chat"`
	ThreadID string `json:"thread"`
}

type Message struct {
	Guild   string
	Channel string
	Message string
	Content string
	Profile Profile
	Embeds  []Embed
}

type Embed struct {
	Author Author
	Body   Body
	Fields []Field
	Image  Image
	Footer Footer
}

type Author struct {
	Name    string
	URL     string
	IconURL string
}

type Body struct {
	Title       string
	Description string
	URL         string
	Color       int
}

type Field struct {
	Name  string
	Value string
}

type Image struct {
	URL          string
	ThumbnailURL string
}

type Footer struct {
	Text          string
	Timestamp     string
	FooterIconURL string
}

type Profile struct {
	ID        string
	Name      string
	AvatarURL string
}
