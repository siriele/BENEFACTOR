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
  	var today = new Date().getTime();
  	var exclude = {
		state:true,
		users: true,
		type:true,
		game: true,

	};
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
	  		send(JSON.stringify({id: doc._id, info: info}));
  	}
  	send("]");
}