function (doc) {
	if (doc.type == "channel" && doc.channel_type && doc.date_closed){
		//0-off,1-on,2-all
		//0 created, 1 updated, 2 duration, 3  exited
		//[game, type, status, role, id, sort_by] // maybe use strings here
		var status = 0;
		emit([doc.game, doc.channel_type, 2, doc.closer, 0, doc.date_created], null);
		emit([doc.game, doc.channel_type, status, doc.closer, 0, doc.date_created], null);
		emit([doc.game, doc.channel_type, 2, doc.closer, 1, doc.date_updated], null);
		emit([doc.game, doc.channel_type, status, doc.closer, 1, doc.date_updated], null);
		emit([doc.game, doc.channel_type, 2, doc.closer, 2, doc.date_closed-doc.date_created], null);
		emit([doc.game, doc.channel_type, status, doc.closer, 2, doc.date_closed-doc.date_created], null);	
		emit([doc.game, doc.channel_type, 2, doc.closer, 3, doc.date_closed], null);
		emit([doc.game, doc.channel_type, status, doc.closer, 3, doc.date_closed], null);				
	}
}