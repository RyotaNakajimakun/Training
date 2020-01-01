package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	//空のclientを作成
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLはErrNoAvatarURLを返すべきです")
	}

	//空のclientに値を設定
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.GetAvatarURLはエラーを返すべきではありません")
	} else {
		if url != testUrl {
			t.Error("値が存在する場合、AuthAvatar.GetAvatarURLは正しいURLを返すべです")
		}
	}
}

func TestGravatarAvatar (t *testing.T) {
	var gravatarAvatar GravatarAvatar

	client := new(client)
	client.userData = map[string]interface{}{"userid": ""}


	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURLはエラーを返すべきではありません")
	}

	if url != "//www.garavatar.com/avatar/" {
		t.Errorf("GravatarAvatar.GetAvatarが%sという謝った値を返しました", url)
	}
}


func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.garavatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}