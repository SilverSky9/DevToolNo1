package test

import (
	"testing"
)

func TestIntergation(t *testing.T) {

}

// func TestCreatePost(t *testing.T) {
// 	post_test := model.Post{
// 		PostId:        2,
// 		PinId:         99,
// 		ProductName:   "want thai ice tea",
// 		PostDate:      time.Now(),
// 		ProductOption: "buy",
// 		Price:         50,
// 		Amount:        1,
// 		TagId:         99,
// 	}
// 	user_test := model.User{
// 		UserId:      2,
// 		Name:        "tester naja",
// 		Address:     "localhost",
// 		Contact:     "test@gmail.com",
// 		PhoneNumber: "1212312121",
// 	}
// 	pin_test := model.Pin{
// 		PinId:  99,
// 		Pin:    "99999999",
// 		UserId: 2,
// 	}
// 	tag_test := model.Tag{
// 		TagId:   99,
// 		TagName: "food",
// 	}
// 	type args struct {
// 		post model.Post
// 		tag  model.Tag
// 		user model.User
// 		pin  model.Pin
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{name: "Create-post-test-1", args: args{post: post_test, tag: tag_test, user: user_test, pin: pin_test}, wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := services.CreatePost(tt.args.post, tt.args.tag, tt.args.user, tt.args.pin); (err != nil) != tt.wantErr {
// 				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
