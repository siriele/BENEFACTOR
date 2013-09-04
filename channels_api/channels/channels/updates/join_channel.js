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

	if(doc.users[userId]){
		return [null, {code: 204}];
	}

	//erase history of user
	// if (doc.users[userId].date_exited){
	// 	delete doc.users[userId];
	// }
	var today = new Date().getTime();
	doc.users[userId] = {date_joined: today ,total_updates: 0};
	return [doc, JSON.stringify(doc.users[userId])];
}