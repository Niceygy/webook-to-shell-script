import WebSocket from 'ws';
const { spawn } = require('node:child_process')
const config = require("./config.json")


const ws = new WebSocket('ws://0.0.0.0:/neverGonnaGiveYouUp');

ws.on('open', function open() {
    ws.send('never gonna let you down');
});

ws.on('message', function message(data) {
    console.log('received: %s', data);
    // start the cmd
const command = spawn(config.PATH)

// the `data` event is fired every time data is
// output from the command
command.stdout.on('data', output => {
    // the output data is captured and printed in the callback
    console.log("Output: ", output.toString())
})
});


