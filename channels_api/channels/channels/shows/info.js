function (doc, req) {
  	
	if (!doc){
		return {
			code: 204
		};
	}

	var game = req.query.gameId;
	var channel_type = req.query.channel_type;
	var userId = req.query.breaktimeId;
	var today = new Date().getTime();
	if(!game || !userId || !channel_type){
		return {
			code: 400,
			body :"breaktimeId, gameId, and channel_type are required"
		};
	}

	if (doc.channel_type != channel_type){
		return {
			code: 400,
			body: "channel type missmatch"
		};
	}

	if (doc.game != game){
		return {
			code: 400,
			body: "game type missmatch"
		};
	}
	var exclude = {
		state:true,
		users: true,
		type:true,
		game: true,

	};
	var info = {};
	var users = {};
	info.id = doc._id;
	for( var u in doc.users){
		if (doc.users[u].date_exited){
			users[u] = doc.users[u];
		}
	}

	for ( var k in doc ){
		if (k.indexOf("_") ==0){
			continue;
		}
		if (exclude[k]){
			continue;
		}
		info[k] = doc[k];
	}

	if (doc.date_closed){
		info.duration = doc.date_closed - doc.date_created;
	}else{
		info.duration = today - doc.date_created;
	}

	info.users = Object.keys(users).length;
	info.timestamp = today;
	return {
    	headers: {
      		"Content-Type": "application/json"
     	},
     	body: JSON.stringify(info)
  	};
}