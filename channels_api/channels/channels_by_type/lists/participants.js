function (doc, req) {
	//repect options like max duration
	//and other options that are placed o the fly
	start({
    headers: {
      "Content-Type": "application/json"
     }
  	});

  	var row;
  	var comma = false;
  	send("[");
  	while (row = getRow()){
  		var doc = row.doc;
  		var emiter = row.id;
  		var key = row.key;
  		if(comma){
  			send(",")
  		}else{
  			comma = true;
  		}
  		var users = {};
		for( var u in doc.users){
			if (doc.users[u].date_exited){
				users[u] = doc.users[u];
			}
		}
  		send(JSON.stringify({id: doc._id, participants: users}));
  	}
  	send("]");
}