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
        //load up guys in data
        for (var i = 0; i < clients.clientids.length; i++) {
            //get everyones coordinates
            var ids = 'id' + clients.clientids[i];
            var boxes = $("." + ids);
            var bbox = boxes.offset();
            var bf1right = bbox.left + boxes.outerWidth();
            var bf1bottom = bbox.top + boxes.outerHeight();

            var guy = {
                left: bbox.left,
                top: bbox.top,
                right: bf1right,
                bottom: bf1bottom,
                clientid: clients.clientid[i]
            }

            //check if id is in array 
            if (getids.includes(clients.clientids[i]) == false) {
                getids.push(clients.clientids[i]);
                guys.push(guy);
                console.log(guys[i].clientids, "!=", clients.clientids[i])
            }

            //find the instance of this guy
            var className = $("id" + clients.clientid[i]).attr("class");
            var result = className.substring(className.lastIndexOf("id") + 2);

            $(".id" + result).css("position", "absolute").
                animate({
                    "top": guy.top,
                    "left": guy.left,
                    "right": guy.right,
                    "bottom": guy.bottom,
                });


        }
    }
}