function (doc) {
	if(doc.type == "channel"){
		emit([doc.channel_type,doc._id], null);
	}
}