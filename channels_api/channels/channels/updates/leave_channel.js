function (doc, req) {
	if (!doc){
		return [null, {code: 204}];
	}

	if (doc.date_closed){
		return [null, {code: 204}];
	}

	var game = req.query.gameId;
	var userId = req.query.breaktimeId;
	var channel_type = req.query.channel_type;
	if(!game || !userId || !channel_type){
		return [null, {code: 400, body: "breaktimeId, gameId, and channel_type are required"}];
	}

	if (doc.channel_type != channel_type){
		return [null, {code: 400, body: "channel_type missmatch"}];
	}

	if (doc.game != game){
		return [null, {code: 400, body: "game missmatch"}];
	}
	
	if(!doc.users[userId] || doc.users[userId].date_exited){
		return [null, {code: 204}];
	}

	//work on delta code
	var today = new Date().getTime();
	doc.users[userId].date_exited = today;
	return [doc, JSON.stringify(doc.users[userId])];
}