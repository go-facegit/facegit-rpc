package db

import (
	"time"

	"gorm.io/gorm"
)

type UserType int

const (
	UserIndividual UserType = iota // Historic reason to make it starts at 0.
	UserOrganization
)

// User represents the object of individual and member of organization.
type User struct {
	ID        int64
	LowerName string `xorm:"UNIQUE NOT NULL" gorm:"UNIQUE"`
	Name      string `xorm:"UNIQUE NOT NULL" gorm:"NOT NULL"`
	FullName  string
	// Email is the primary email address (to be used for communication)
	Email       string `xorm:"NOT NULL" gorm:"NOT NULL"`
	Passwd      string `xorm:"NOT NULL" gorm:"NOT NULL"`
	LoginSource int64  `xorm:"NOT NULL DEFAULT 0" gorm:"NOT NULL;DEFAULT:0"`
	LoginName   string
	Type        UserType
	// OwnedOrgs   []*User       `xorm:"-" gorm:"-" json:"-"`
	// Orgs        []*User       `xorm:"-" gorm:"-" json:"-"`
	// Repos       []*Repository `xorm:"-" gorm:"-" json:"-"`
	Location string
	Website  string
	Rands    string `xorm:"VARCHAR(10)" gorm:"TYPE:VARCHAR(10)"`
	Salt     string `xorm:"VARCHAR(10)" gorm:"TYPE:VARCHAR(10)"`

	Created     time.Time `xorm:"-" gorm:"-" json:"-"`
	CreatedUnix int64
	Updated     time.Time `xorm:"-" gorm:"-" json:"-"`
	UpdatedUnix int64

	// Remember visibility choice for convenience, true for private
	LastRepoVisibility bool
	// Maximum repository creation limit, -1 means use global default
	MaxRepoCreation int `xorm:"NOT NULL DEFAULT -1" gorm:"NOT NULL;DEFAULT:-1"`

	// Permissions
	IsActive         bool // Activate primary email
	IsAdmin          bool
	AllowGitHook     bool
	AllowImportLocal bool // Allow migrate repository by local path
	ProhibitLogin    bool

	// Avatar
	Avatar          string `xorm:"VARCHAR(2048) NOT NULL" gorm:"TYPE:VARCHAR(2048);NOT NULL"`
	AvatarEmail     string `xorm:"NOT NULL" gorm:"NOT NULL"`
	UseCustomAvatar bool

	// Counters
	NumFollowers int
	NumFollowing int `xorm:"NOT NULL DEFAULT 0" gorm:"NOT NULL;DEFAULT:0"`
	NumStars     int
	NumRepos     int

	// For organization
	Description string
	NumTeams    int
	NumMembers  int
	// Teams       []*Team `xorm:"-" gorm:"-" json:"-"`
	// Members     []*User `xorm:"-" gorm:"-" json:"-"`
}

func (u *User) BeforeInsert(tx *gorm.DB) error {
	u.CreatedUnix = time.Now().Unix()
	u.UpdatedUnix = u.CreatedUnix
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.MaxRepoCreation < -1 {
		u.MaxRepoCreation = -1
	}
	u.UpdatedUnix = time.Now().Unix()
	return nil
}
