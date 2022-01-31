import { createguys } from './guys/createguys.js';
var clients = [];
export function getsocket(ws) {
    ws.onopen = function () {
        console.log('Connected')
    }
    //entry point
    ws.onmessage = function (evt) {
        //backend sends all guys data not the other way per tab
        clients = jQuery.parseJSON(evt.data);
        //if the amount of guys grow then create another one
        createguys(clients);

        var guys = [];
        var getids = [];
        var counter = 0;
        //load up guys in data
        for (var i = 0; i < clients.length; i++) {
            counter++;
            //get everyones coordinates
            var ids = 'id' + clients[i].clientids[i];
            var boxes = $("." + ids);
            var bbox = boxes.offset();
            var bf1right = bbox.left + boxes.outerWidth();
            var bf1bottom = bbox.top + boxes.outerHeight();

            var guy = {
                left: clients[i].left,
                top: clients[i].top,
                right: bf1right,
                bottom: bf1bottom,
                clientid: clients[i].clientid[i]
            }

            //check if id is in array 
            if (getids.includes(clients[i].clientids[i]) == false) {
                getids.push(clients[i].clientids[i]);
                guys.push(guy);
                console.log(guys[i].clientids, "!=", clients[i].clientids[i])
            }
            //take all new guys
            $(".character").each(function () {
                //find the instance of this guy
                var className = $(this).attr("class");
                var result = className.substring(className.lastIndexOf("id") + 2);
                //per id of new guys
                if (clients[i].id == result) {
                    $(".id" + result).css("position", "absolute").
                        animate({
                            "top": clients[i].top,
                            "left": clients[i].left,
                            "right": clients[i].right,
                            "bottom": clients[i].bottom,
                        });
                }
            });
        }
    }
}