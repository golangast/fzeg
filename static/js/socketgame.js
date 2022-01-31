// https://javascript.info/import-export
import { load } from './mod/loadsocket.js';
import { getsocket } from './mod/getsocket.js';
import { sendsocket } from './mod/sendsocket.js';

function main() {
    var ws = load();
    getsocket(ws);
    sendsocket(ws);
}
main();





