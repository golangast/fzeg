


export function createguys(client) {
    console.log("what you are getting from server", client);
    if (client.clientids.length > document.querySelectorAll("guy").length) {
        //add first id to first guy but gotta remember every tab has their first guy
        if ($(".first").not("id" + client.clientids[0])) {
            $(".first").addClass("id" + client.clientids[0]);
        }
        var count = client.clientids.length - document.querySelectorAll("guy").length
        //create guys
        for (let i = 0; i < count; i++) {
            if ($("guy").not("id" + client.clientids[i])) {
                $("#pane").append(`<guy border='0' class="character newguy id` + client.clientids[i] + `"><img src=""></guy>`);
            }
        }
    }
}


