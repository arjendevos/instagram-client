package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/arjendevos/instagram-client/models/responses"
)

func (c *Client) GetRecommendedAccounts(userID string) (*responses.RecommendedAccountsResponse, error) {
	res, err := c.get(fmt.Sprintf("/fbsearch/accounts_recs?target_user_id=%v&include_friendship_status=true", userID))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	//Create a variable of the same type as our model
	var recommendedAccountsResponse responses.RecommendedAccountsResponse
	//Decode the data
	err = json.NewDecoder(res.Body).Decode(&recommendedAccountsResponse)
	if err != nil {
		bodyB, _ := ioutil.ReadAll(res.Body)
		bodyStr := string(bytes.Replace(bodyB, []byte("\r"), []byte("\r\n"), -1))
		fmt.Println(bodyStr)
		return nil, err
	}
	//Invoke the text output function & return it with nil as the error value
	return &recommendedAccountsResponse, nil

}

func (c *Client) GetProfile(username string) (*responses.Profile, error) {
	res, err := c.get(fmt.Sprintf("/users/web_profile_info/?username=%v", username))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var profile responses.Profile
	err = json.NewDecoder(res.Body).Decode(&profile)
	if err != nil {
		bodyB, _ := ioutil.ReadAll(res.Body)
		bodyStr := string(bytes.Replace(bodyB, []byte("\r"), []byte("\r\n"), -1))
		fmt.Println(bodyStr)
		return nil, err
	}
	return &profile, nil

}
