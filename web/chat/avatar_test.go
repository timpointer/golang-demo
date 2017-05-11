// Package main provides ...
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	//set a value
	testURL := "http://url-to-gravatar/"
	testUser = &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return(testURL, nil)
	testChatUser.User = testUser
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return should return no error when value present")
	}

	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	testChatUser := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	tables := []struct {
		filename string
		path     string
	}{
		{"abc.jpg", "/avatars/abc.jpg"},
		{"abc.png", "/avatars/abc.png"},
	}
	for _, v := range tables {
		t.Log(v.filename)
		filename := filepath.Join("avatars", v.filename)
		ioutil.WriteFile(filename, []byte{}, 0777)
		var fileSystemAvatar FileSystemAvatar
		testChatUser := &chatUser{uniqueID: "abc"}
		url, err := fileSystemAvatar.GetAvatarURL(testChatUser)
		if err != nil {
			t.Error("GravatarAvatar.GetAvatarURL should not return an error")
		}
		if url != v.path {
			t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
		}
		os.Remove(filename)
	}

}
