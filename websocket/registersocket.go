/*
This is the meat of the websocket and
really the only place anything needs to change
It is split up into 3 parts
1. registering the user/joining per tab
2. unregistering the user/if they leave
3. sending the clients to the frontend
*/

package zegsocket

func (pool *Pool) Start() {

	for {
		select {
		//1. this is where a new tab registers
		case client := <-pool.Register: //join
			Registering(client, pool)
			break //end of join

			//2. this is when they leave
		case client := <-pool.Unregister: //leave
			Unregistering(client, pool)
			break //end of leave

			//3. this is where data from the clients gets updated and sent back out to the frontend
		case message := <-pool.Broadcast: //stay
			CycleClients(pool, message)

		} //end of Broadcast
	} //end of for loop
} //end of start
