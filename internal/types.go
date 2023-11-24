package internal

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
	Name      string
	AvatarURL string
}
