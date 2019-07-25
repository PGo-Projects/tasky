package user

import (
	"fmt"
	"regexp"
	"strconv"
)

type User struct {
	Username      string
	UnusedIndices string `bson:"unusedindices,omitempty"`
}

func (u *User) InsertUnusedIndex(index int64) {
	u.UnusedIndices = fmt.Sprintf("%d:%s", index, u.UnusedIndices)
}

func (u *User) PopNextUnusedIndex() int64 {
	re := regexp.MustCompile(`(\d+):?(.*)`)
	matches := re.FindStringSubmatch(u.UnusedIndices)

	// Ignoring the error in conversion as the regexp has already ensure
	// that it is an integer
	index, _ := strconv.ParseInt(matches[1], 10, 64)
	u.UnusedIndices = matches[2]
	if u.UnusedIndices == "" {
		u.UnusedIndices = strconv.FormatInt(index+1, 10)
	}

	return index
}
