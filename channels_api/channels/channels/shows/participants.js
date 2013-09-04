function (doc, req) {
	start({
	    headers: {
	      "Content-Type": "application/json"
	     }
  	});
  	
	if (!doc){
		return {
			code: 204
		};
	}

	var game = req.query.gameId;
	var channel_type = req.query.channel_type;
	var userId = req.query.breaktimeId;
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
	var users = {};
	for( var u in doc.users){
		if (doc.users[u].date_exited){
			users[u] = doc.users[u];
		}
	}

	return{
    	headers: {
      		"Content-Type": "application/json"
     	},
     	body: JSON.stringify(users)
  	};
}