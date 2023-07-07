package main

// func TestAPI(t *testing.T) {
// 	APIKey := os.Getenv("APIKEY")
// 	client := NewDUEAPIClient(APIKey)
// 	// state := NewStateFromJsonOrPanic([]byte(`{"current_step": "clients", "next_page_number": 45, "remaining_steps": ["answer_sets"], "all_steps": ["clients", "answer_sets"]}`))
// 	state := NewStateFromJsonOrPanic([]byte(`{"current_step": "answer_sets", "next_page_number": 810, "remaining_steps": [], "steps": ["clients", "answer_sets"]}`))
// 	fmt.Println(state)
// 	data, nextState, hasMore, err := client.CallAPI(state)
// 	if err != nil {
// 		t.Log(err.Error())
// 		t.Fail()
// 	}
// 	fmt.Println(data)
// 	fmt.Println(nextState)
// 	fmt.Println(hasMore)
// }
