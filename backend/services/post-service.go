package services

import (
	model "daeng-market/models"
	"daeng-market/queries"
	"fmt"
	"strconv"
)

func CreatePost(post model.Post, tag model.Tag, user model.User, pin model.Pin) error {
	resp_tag, err := queries.GetTagByNameQueries(tag.TagName)
	if err != nil {
		//no duplicate tag
		//create new tag
		tag_id, err := CreateTag(tag.TagName)
		if err != nil {
			return err
		}
		post.TagId = tag_id
	} else {
		post.TagId = resp_tag.TagId
	}

	//create post owner user
	user_id, err := CreateUser(user)
	if err != nil {
		return err
	}

	//create owner pin
	pin_id, err := CreatePin(pin.Pin, user_id)
	if err != nil {
		return err
	}
	post.PinId = pin_id

	//create post
	err = queries.CreatePostQueries(post)
	if err != nil {
		return err
	}

	return nil
}
func GetAllPost() ([]model.Post, error) {
	resp, err := queries.GetAllPostQueries()
	if err != nil {
		return []model.Post{}, err
	}
	return resp, nil
	// result, error = queries.GetAllPostQueries()
	// return result, error
}
func GetPostByTag(tag1, tag2 string) ([]model.Post, error) {
	resp, err := queries.GetPostByTagQueries(tag1, tag2)
	if err != nil {
		return []model.Post{}, err
	}

	return resp, nil
}

func GetPostById(id string) (model.Post, error) {
	// pase id datatype string to int
	post_id, _ := strconv.Atoi(id)

	resp, err := queries.GetPostByIdQueries(post_id)
	if err != nil {
		return model.Post{}, err
	}
	return resp, nil
}

func GetPostByName(post_name string) ([]model.Post, error) {
	// pase id datatype string to int

	resp, err := queries.GetPostByNameQueries(post_name)
	if err != nil {
		return []model.Post{}, err
	}
	return resp, nil
}

func GetPostByMultiTag(tags string) ([]model.Post, error) {
	fmt.Println(tags)
	tag_list := MultiTagToArray(tags)

	resp, err := queries.GetPostByMultiTagQueries(tag_list)
	if err != nil {
		return []model.Post{}, nil
	}

	return resp, nil
}

func MultiTagToArray(tags string) []int {
	// var x string = "1-20-32-402"
	var list []int
	temp := ""
	for i := 0; i < len(tags); i++ {
		num := string(tags[i])
		if num != "," {
			temp += num
		}
		if num == "," {
			to_num, _ := strconv.Atoi(temp)
			list = append(list, to_num)

			temp = ""
		}
	}
	return list
}
