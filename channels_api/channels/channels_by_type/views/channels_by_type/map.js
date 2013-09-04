function (doc) {
	if (doc.type == "channel" && doc.channel_type){
		//0-off,1-on,2-all
		//0 created, 1 updated, 2 duration, 3  exited
		//[game, type, status, role, id, sort_by] // maybe use strings here
		var status;
		if (doc.date_closed){
			status = 0;
		}else{
			status = 1;
		}
		emit([doc.channel_type, 2, 0, doc.date_created], null);
		emit([doc.channel_type, status, 0, doc.date_created], null);
		emit([doc.channel_type, 2, 1, doc.date_updated], null);
		emit([doc.channel_type, status, 1, doc.date_updated], null);
		if (doc.date_closed){
			emit([doc.channel_type, 2, 2, doc.date_closed-doc.date_created], null);
			emit([doc.channel_type, status, 2, doc.date_closed-doc.date_created], null);	
			emit([doc.channel_type, 2, 3, doc.date_closed], null);
			emit([doc.channel_type, status, 3, doc.date_closed], null);			
		}
	}
}