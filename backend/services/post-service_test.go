package services

// func TestGetAllPost(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		want    model.Post
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		// {name: "User-case-Post-1",
// 		// 	want: {
// 		// 		"post_id":        1,
// 		// 		"product_name":   "kuy",
// 		// 		"post_date":      "2022-02-12T00:00:00Z",
// 		// 		"product_option": "buy",
// 		// 		"price":          555,
// 		// 		"amount":         8,
// 		// 		"pin_id":         1,
// 		// 		"tag_id":         1,
// 		// 	},
// 		// 	wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetAllPost()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetAllPost() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetAllPost() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetPostById(t *testing.T) {
// 	a, _ := json.Marshal(map[string]int{"post_id": 1, "pin_id": 1, "post_name": "Kuy", "post_date": "2022-02-12T00:00:00Z", "product_option": "buy", "price": 555, "amount": 8, "tag_id": 1})
// 	fmt.Println(string(a)) // {"bar":2,"baz":3,"foo":1}
// 	type args struct {
// 		id string
// 	}
// 	tests := []struct {
// 		name    string
// 		want    model.Post
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{name: "User-case-Post-1",
// 			want:    a,
// 			wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := GetPostById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetPostById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
