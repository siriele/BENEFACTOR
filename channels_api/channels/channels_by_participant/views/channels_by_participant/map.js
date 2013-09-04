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
		for(var u in doc.users){
			switch(status){
				case 0:
					if (doc.users[u].date_exited){
						emit([doc.channel_type, status, u, 0, doc.date_created], null);
						emit([doc.channel_type, status, u, 1, doc.date_updated], null);
						emit([doc.channel_type, status, u, 2, doc.date_closed-doc.date_created], null);
						emit([doc.channel_type, status, u, 3, doc.date_closed], null);					
					}
					break;
				case 1:
					if (!doc.users[u].date_exited){
						emit([doc.channel_type, status, u, 0, doc.date_created], null);
						emit([doc.channel_type, status, u, 1, doc.date_updated], null);			
					}				
					break;
			}
			emit([doc.channel_type, 2, u, 0, doc.date_created], null);
			emit([doc.channel_type, 2, u, 1, doc.date_updated], null);
			emit([doc.channel_type, 2, u, 2, doc.date_closed-doc.date_created], null);
			emit([doc.channel_type, 2, u, 3, doc.date_closed], null);
		}
	}
}