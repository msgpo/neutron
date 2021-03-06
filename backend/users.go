package backend

type UsersBackend interface {
	// Check if a username is available.
	IsUsernameAvailable(username string) (bool, error)
	// Get a user.
	GetUser(id string) (*User, error)
	// Check if the provided username and password are correct
	Auth(username, password string) (*User, error)
	// Insert a new user. Returns the newly created user.
	InsertUser(user *User, password string) (*User, error)
	// Update an existing user.
	UpdateUser(update *UserUpdate) error
	// Update a user's password.
	UpdateUserPassword(id, current, new string) error
	// Delete a user.
	//DeleteUser(id string) error
}

// A user.
type User struct {
	ID string
	Name string
	NotificationEmail string
	Signature string
	NumMessagePerPage int
	UsedSpace int
	Notify int
	AutoSaveContacts int
	Language string
	LogAuth int
	ComposerMode int
	MessageButtons int
	ShowImages int
	ShowEmbedded int
	ViewMode int
	ViewLayout int
	SwipeLeft int
	SwipeRight int
	Theme string
	Currency string
	Credit int
	InvoiceText string
	AlsoArchive int
	Hotkeys int
	PMSignature int
	TwoFactor int
	PasswordReset int
	PasswordMode int
	News int
	DisplayName string
	MaxSpace int
	MaxUpload int
	Role int
	Private int
	Subscribed int
	Deliquent int
	VPN interface{}
	Addresses []*Address
	Keys []*Keypair
}

func (u *User) GetMainAddress() *Address {
	for _, addr := range u.Addresses {
		if addr.Send == 1 { // 1 is main address, 2 is secondary address
			return addr
		}
	}
	return nil
}

const (
	RoleFreeUser int = iota
	RolePaidMember
	RolePaidAdmin
)

// A request to update a user.
// Fields set to true will be updated with values in User.
type UserUpdate struct {
	User *User
	DisplayName bool
	Signature bool
	AutoSaveContacts bool
	ShowImages bool
	ComposerMode bool
	ViewLayout bool
	MessageButtons bool
	Theme bool
}

// Apply this update on a user.
func (update *UserUpdate) Apply(user *User) {
	updated := update.User

	if update.DisplayName {
		user.DisplayName = updated.DisplayName
	}
	if update.Signature {
		user.Signature = updated.Signature
	}
	if update.AutoSaveContacts {
		user.AutoSaveContacts = updated.AutoSaveContacts
	}
	if update.ShowImages {
		user.ShowImages = updated.ShowImages
	}
	if update.ComposerMode {
		user.ComposerMode = updated.ComposerMode
	}
	if update.ViewLayout {
		user.ViewLayout = updated.ViewLayout
	}
	if update.MessageButtons {
		user.MessageButtons = updated.MessageButtons
	}
	if update.Theme {
		user.Theme = updated.Theme
	}
}
