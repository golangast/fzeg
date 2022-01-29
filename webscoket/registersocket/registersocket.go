/*
This is the meat of the websocket and
really the only place anything needs to change
It is split up into 3 parts
1. registering the user/joining per tab
2. unregistering the user/if they leave
3. sending the clients to the frontend
*/

package registersocket

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func (pool *Pool) Start() {

	for {
		select {
		//1. this is where a new tab registers
		case client := <-pool.Register: //join
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			pool.PoolID = "12"
			if pool.ClientID == "" {
				for client, _ := range pool.Clients {
					c := GetClientID()
					pool.ClientID = c
					var _ = client
				}
			}

			break //end of join

			//2. this is when they leave
		case client := <-pool.Unregister: //leave
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON("New User Joined...")
			}
			break //end of leave

			//3. this is where data from the clients gets
			//updated and sent back out to the frontend
		case message := <-pool.Broadcast: //stay

			spew.Dump(message.Body) //prints out what backend recieves

			//cycle through clients
			for client, _ := range pool.Clients {
				err := json.Unmarshal([]byte(message.Body), &client)
				if err != nil {
					panic(err)
				}

				fmt.Println(client)
				//add id
				var d = Data{}

				if !Contains(clientids, pool.ClientID) {
					clientids = append(clientids, pool.ClientID)

					d.ClientIDs = clientids
					d.ID = pool.ClientID
					d.Left = client.Left
					d.Top = client.Top
					d.Right = client.Right
					d.Bottom = client.Bottom
					d.PoolID = pool.PoolID

					// Das = append(Das, d)
					// fmt.Println(Das)

					//encode and send
					if err := client.Conn.WriteJSON(d); err != nil { //send it
						fmt.Println(err)
						return
					}

				} else {

					d.Left = client.Left
					d.Top = client.Top
					d.Right = client.Right
					d.Bottom = client.Bottom
					d.ID = pool.ClientID
					d.PoolID = pool.PoolID

					//encode and send back to frontend
					if err := client.Conn.WriteJSON(Das); err != nil { //send it
						fmt.Println(err)
						return
					}
				} //end of client range
			} //end of select
		} //end of Broadcast
	} //end of for loop
} //end of start
