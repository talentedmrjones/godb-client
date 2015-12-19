var
	//JSONStream = require('JSONStream'),
	net = require('net'),
	uuid = require('uuid')
;


var socket = net.connect('6000')

// socket.pipe(JSONStream.parse()).pipe(function() {
// 	console.log(arguments);
// });

socket.on('data', function (buffer) {
	console.log(buffer.toString());
});

function getCommand(data) {
	return {
		ID:uuid.v4(),
		Action:"c",
		Db:"tmj",
		Table:"orders",
		Data:data
	}
}

function getEvent(time) {
	var events = [
		// {domain:"talentedmrjones.com", path:"/"},
		// {domain:"talentedmrjones.com", path:"/about"},
		// {domain:"talentedmrjones.com", path:"/services"},
		// {domain:"talentedmrjones.com", path:"/cloud"},
		// {domain:"talentedmrjones.com", path:"/mobile"},
		// {domain:"talentedmrjones.com", path:"/contact"},
		{action:"placed", total:0},
		{action:"placed", total:0},
		{action:"placed", total:0},
		{action:"placed", total:0},
		{action:"placed", total:0},
		{action:"canceled", total:0},
		{action:"timeout", total:0}
	];
	var event = events[Math.floor((Math.random()*(events.length)))];
	event.time = time;
	if (event.action=="placed") {
		event.total = Math.floor(Math.random()*100)
	}
	return event;
}

setInterval(function(){
	var
		event = getEvent(Date.now()),
		command = getCommand(event)
	;

	socket.write(JSON.stringify(command));
},10);





//socket.end();
