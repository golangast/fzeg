export function sendsocket(ws) {

    setInterval(function () {
        //your guy
        //you need two tabs so that an id can virtually be assigned to you.
        var first = $('.first');
        var className = $(first).attr("class");
        var guyid = className.substring(className.lastIndexOf("id") + 2);
        //get coordinates
        var guy = first.offset();
        var guyright = guy.left + first.outerWidth();
        var guybottom = guy.top + first.outerHeight();

        //load it
        var packagedguy = {
            left: guy.left,
            top: guy.top,
            right: guyright,
            bottom: guybottom,
            clientid: guyid
        }

        //send it up
        var gj = JSON.stringify(packagedguy);
        console.log("sending to server ", packagedguy);
        ws.send(gj);
    }, 1000); //end of repaint
}