function tick(bc) {
	var ball = bc._queue.shift();

	if (bc._1minTrack.length < 4)
		bc._1minTrack.push(ball);
	else {
		emptyTrack(bc, bc._1minTrack);
		if (bc._5minTrack.length < 11)
			bc._5minTrack.push(ball);
		else {
			emptyTrack(bc, bc._5minTrack);
			if (bc._1hourTrack.length < 11)
				bc._1hourTrack.push(ball);
			else {
				emptyTrack(bc, bc._1hourTrack);
				bc._queue.push(ball);

				bc.days += .5;
				bc.areBallsInSeq = checkSeqOrder(bc);
			}
		}
	}
}

function emptyTrack(bc, track) {
	while (track.length > 0) {
		bc._queue.push(track.pop());
	}
}

function checkSeqOrder(bc) {
	for(var b = 1, i = 0; i < bc._queue.length; i++, b++) {
		if(b != bc._queue[i])
			return false;
	}
	return true;
}

function newBallClock(numBalls) {
	var bc = {
		numBalls: numBalls,
		days: 0,
		areBallsInSeq: false,
		_1minTrack: [],
		_5minTrack: [],
		_1hourTrack: [],
		_queue: []
	}
	
	for (var i = 1; i <= bc.numBalls; i++) { 
		bc._queue.push(i);
	}
	return bc;
}

function runClock(input) {
	var result = "";
	var list = input.split("\n");
	
	for(var i = 0; i < list.length; i++) {
		var numBalls = parseInt(list[i])
		if(isNaN(numBalls) || numBalls < 27 || numBalls > 127) {
			result += "'" + list[i] + "' is not a valid number of balls!<br>";
			continue;
		}
		
		var clock = newBallClock(numBalls);
		var start = new Date();
		while (!clock.areBallsInSeq) {
			tick(clock);
		}
		
		result += clock.numBalls + " balls cycle after " + clock.days + " days. time=" + (new Date() - start) + "ms<br>";
	}
	
	return result;
}
