export function load() {

    //need a first guy
    if ($("guy").length == 0) {
        $("#pane").append(`<guy id="box" border='0' class="first character"><img src=""></guy>`);
    }
    var pane = $('#pane'),
        first = $('.first'),
        firstimg = $('.first img'),
        wh = pane.width() - first.width(),
        wv = pane.height() - first.height(),
        d = {},
        x = 5;
    first.css("top", "0%");
    first.css("left", "0%");
    //ensure limits
    function newh(v, a, b) {
        var n = parseInt(v, 10) - (d[a] ? x : 0) + (d[b] ? x : 0);
        return n < 0 ? 0 : n > wh ? wh : n;
    }

    function newv(v, a, b) {
        var n = parseInt(v, 10) - (d[a] ? x : 0) + (d[b] ? x : 0);
        return n < 0 ? 0 : n > wv ? wv : n;
    }
    //keys
    $(window).keydown(function (e) {
        switch (e.which) {
            case 39: // right
                firstimg.attr('src', "/static/img/walk.gif");
                break;
            case 38: // up
                firstimg.attr('src', "/static/img/up.gif");
                break;
            case 37: // left
                firstimg.attr('src', "/static/img/back.gif");
                break;
            case 40: // down
                firstimg.attr('src', "/static/img/down.gif");
                break;
            default:
                firstimg.attr('src', "/static/img/glogo.gif");
        }
        d[e.which] = true;
    });
    $(window).keyup(function (e) {
        d[e.which] = false;
    });

    //to ensure user stays in pane
    setInterval(function () {
        first.css({
            left: function (i, v) {
                return newh(v, 37, 39);
            },
            top: function (i, v) {
                return newv(v, 38, 40);
            }
        });
        //sizes
        wh = pane.width() - first.width();
        wv = pane.height() - first.height();
    }, 90); //end of repaint
    var loc = window.location;
    var uri = 'ws:';
    if (loc.protocol === 'https:') {
        uri = 'wss:';
    }
    uri += '//' + loc.host + "/ws";
    uri += loc.pathname;
    var ws = new WebSocket(uri);
    return ws
}