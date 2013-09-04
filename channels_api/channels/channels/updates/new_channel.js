function (doc, req) {
	if(doc){
		return [null, {code: 400, body: "This test aleady exists"}];
	}
	var game = req.query.gameId;
	var userId = req.query.breaktimeId;
	var channel_type = req.query.channel_type;
	var body = JSON.parse(req.body);
	var options = body.options;
	var state = body.state;
	if(!game || !userId || !channel_type){
		return [null, {code: 400, body: "breaktimeId, gameId, and channel_type are required"}];
	}
	var today = new Date().getTime();
	doc = {_id: req.uuid, type : "channel", users: {}, channel_type: channel_type};
	doc.state = state;
	doc.options = options;
	doc.creator = userId;
	//doc.updater = userId;
	doc.game = game;
	//doc.date_updated = today;
	doc.date_created = today;
	doc.total_updates = 0;
	return [doc, JSON.stringify({id: doc._id ,created: today})];
}