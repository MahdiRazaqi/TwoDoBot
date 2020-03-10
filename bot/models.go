package bot

// Params model
type Params map[string]interface{}

// TelegramResponse model
type TelegramResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update model
type Update struct {
	UpdateID          int     `json:"update_id"`
	Message           Message `json:"message"`
	EditedMessage     Message `json:"edited_message"`
	ChannelPost       Message `json:"channel_post"`
	EditedChannelPost Message `json:"edited_channel_post"`
	Poll              Poll    `json:"poll"`
	// InlineQuery       InlineQuery
	// ChosenInlineResult ChosenInlineResult
	// CallbackQuery      CallbackQuery
	// ShippingQuery      ShippingQuery
	// PreCheckoutQuery   PreCheckoutQuery
	// PollAnswer         PollAnswer
}

// Message model
type Message struct {
	MessageID             int                  `json:"message_id"`
	From                  User                 `json:"from"`
	Date                  int                  `json:"date"`
	Chat                  Chat                 `json:"chat"`
	ForwardFrom           User                 `json:"forward_from"`
	ForwardFromChat       Chat                 `json:"forward_from_chat"`
	ForwardFromMessageID  int                  `json:"forward_from_message_id"`
	ForwardSignature      string               `json:"forward_signature"`
	ForwardSenderName     string               `json:"forward_sender_name"`
	ForwardDate           int                  `json:"forward_date"`
	EditDate              int                  `json:"edit_date"`
	MediaGroupID          string               `json:"media_group_id"`
	AuthorSignature       string               `json:"author_signature"`
	Text                  string               `json:"text"`
	Entities              []MessageEntity      `json:"entities"`
	CaptionEntities       []MessageEntity      `json:"caption_entities"`
	Audio                 Audio                `json:"audio"`
	Document              Document             `json:"document"`
	Animation             Animation            `json:"animation"`
	Game                  Game                 `json:"game"`
	Photo                 []PhotoSize          `json:"photo"`
	Sticker               Sticker              `json:"sticker"`
	Video                 Video                `json:"video"`
	Voice                 Voice                `json:"voice"`
	VideoNote             VideoNote            `json:"video_note"`
	Caption               string               `json:"caption"`
	Contact               Contact              `json:"contact"`
	Location              Location             `json:"location"`
	Venue                 Venue                `json:"venue"`
	Poll                  Poll                 `json:"poll"`
	NewChatMembers        []User               `json:"new_chat_members"`
	LeftChatMember        User                 `json:"left_chat_member"`
	NewChatTitle          string               `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize          `json:"new_chat_photo"`
	DeleteChatPhoto       bool                 `json:"delete_chat_photo"`
	GroupChatCreated      bool                 `json:"group_chat_created"`
	SupergroupChatCreated bool                 `json:"supergroup_chat_created"`
	ChannelChatCreated    bool                 `json:"channel_chat_created"`
	MigrateToChatID       int                  `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                  `json:"migrate_from_chat_id"`
	Invoice               Invoice              `json:"invoice"`
	SuccessfulPayment     SuccessfulPayment    `json:"successful_payment"`
	ConnectedWebsite      string               `json:"connected_website"`
	PassportData          PassportData         `json:"passport_data"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`
	// ReplyToMessage        Message              `json:"reply_to_message"`
	// PinnedMessage         Message              `json:"pinned_message"`
}

// User model
type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// Chat model
type Chat struct {
	ID               int             `json:"id"`
	Type             string          `json:"type"`
	Title            string          `json:"title"`
	Username         string          `json:"username"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	Photo            ChatPhoto       `json:"photo"`
	Description      string          `json:"description"`
	InviteLink       string          `json:"invite_link"`
	Permissions      ChatPermissions `json:"permissions"`
	SlowModeDelay    int             `json:"slow_mode_delay"`
	StickerSetName   string          `json:"sticker_set_name"`
	CanSetStickerSet bool            `json:"can_set_sticker_set"`
	// PinnedMessage    Message         `json:"pinned_message"`
}

// ChatPhoto model
type ChatPhoto struct {
	SmallFileID       string
	SmallFileUniqueID string
	BigFileID         string
	BigFileUniqueID   string
}

// ChatPermissions model
type ChatPermissions struct {
	CanSendMessages       bool
	CanSendMediaMessages  bool
	CanSendPolls          bool
	CanSendOtherMessages  bool
	CanAddWebPagePreviews bool
	CanChangeInfo         bool
	CanInviteUsers        bool
	CanPinMessages        bool
}

// MessageEntity model
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`
	User     User   `json:"user"`
	Language string `json:"language"`
}

// Audio model
type Audio struct {
	FileID       string
	FileUniqueID string
	Duration     int
	Performer    string
	Title        string
	MimeType     string
	FileSize     int
	Thumb        PhotoSize
}

// Document model
type Document struct {
	FileID       string
	FileUniqueID string
	Thumb        PhotoSize
	FileName     string
	MimeType     string
	FileSize     int
}

// PhotoSize model
type PhotoSize struct {
	FileID       string
	FileUniqueID string
	Width        int
	Height       int
	FileSize     int
}

// Animation model
type Animation struct {
	FileID       string
	FileUniqueID string
	Width        int
	Height       int
	Duration     int
	Thumb        PhotoSize
	FileName     string
	MimeType     string
	FileSize     int
}

// Game model
type Game struct {
	Title        string
	Description  string
	Photo        []PhotoSize
	Text         string
	TextEntities []MessageEntity
	Animation    Animation
}

// Sticker model
type Sticker struct {
	FileID       string
	FileUniqueID string
	Width        int
	Height       int
	IsAnimated   bool
	Thumb        PhotoSize
	Emoji        string
	SetName      string
	MaskPosition MaskPosition
	FileSize     int
}

// MaskPosition model
type MaskPosition struct {
	Point  string
	XShift float32
	YShift float32
	Scale  float32
}

// Video model
type Video struct {
	FileID       string
	FileUniqueID string
	Width        int
	Height       int
	Duration     int
	Thumb        PhotoSize
	MimeType     string
	FileSize     int
}

// Voice model
type Voice struct {
	FileID       string
	FileUniqueID string
	Duration     int
	MimeType     string
	FileSize     int
}

// VideoNote model
type VideoNote struct {
	FileID       string
	FileUniqueID string
	Length       int
	Duration     int
	Thumb        PhotoSize
	FileSize     int
}

// Contact model
type Contact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	UserID      int
	Vcard       string
}

// Location model
type Location struct {
	Longitude float32
	Latitude  float32
}

// Venue model
type Venue struct {
	Location       Location
	Title          string
	Address        string
	FoursquareID   string
	FoursquareType string
}

// Poll model
type Poll struct {
	ID                    string
	Question              string
	Options               []PollOption
	TotalVoterCount       int
	IsClosed              bool
	IsAnonymous           bool
	Type                  string
	AllowsMultipleAnswers bool
	CorrectOptionID       int
}

// PollOption model
type PollOption struct {
	Text       string
	VoterCount int
}

// Invoice model
type Invoice struct {
	Title          string
	Description    string
	StartParameter string
	Currency       string
	TotalAmount    int
}

// SuccessfulPayment model
type SuccessfulPayment struct {
	Currency                string
	TotalAmount             int
	InvoicePayload          string
	ShippingOptionID        string
	OrderInfo               OrderInfo
	TelegramPaymentChargeID string
	ProviderPaymentChargeID string
}

// OrderInfo model
type OrderInfo struct {
	Name            string
	PhoneNumber     string
	Email           string
	ShippingAddress ShippingAddress
}

// ShippingAddress model
type ShippingAddress struct {
	CountryCode string
	State       string
	City        string
	StreetLine1 string
	StreetLine2 string
	PostCode    string
}

// PassportData model
type PassportData struct {
	Data        []EncryptedPassportElement
	Credentials EncryptedCredentials
}

// EncryptedPassportElement model
type EncryptedPassportElement struct {
	Type        string
	Data        string
	PhoneNumber string
	Email       string
	Files       []PassportFile
	FrontSide   PassportFile
	ReverseSide PassportFile
	Selfie      PassportFile
	Translation []PassportFile
	Hash        string
}

// EncryptedCredentials model
type EncryptedCredentials struct {
	Data   string
	Hash   string
	Secret string
}

// PassportFile model
type PassportFile struct {
	FileID       string
	FileUniqueID string
	FileSize     int
	FileDate     int
}

// InlineKeyboardMarkup model
type InlineKeyboardMarkup struct {
	InlineKeyboard InlineKeyboardButton
}

// InlineKeyboardButton model
type InlineKeyboardButton struct {
	Text                         string
	URL                          string
	LoginURL                     LoginURL
	CallbackData                 string
	SwitchInlineQuery            string
	SwitchInlineQueryCurrentChat string
	pay                          bool
	// CallbackGame                 CallbackGame
}

// LoginURL model
type LoginURL struct {
	URL                string
	ForwardText        string
	BotUsername        string
	RequestWriteAccess bool
}
