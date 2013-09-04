function  (doc, req) {
	if (!doc){
		return [null, {code: 204}];
	}

	if (doc.date_closed){
		return [null, {code: 204}];
	}

	var game = req.query.gameId;
	var userId = req.query.breaktimeId;
	var channel_type = req.query.channel_type;
	var delta = JSON.parse(req.body);
	if(!game || !userId || !channel_type){
		return [null, {code: 400, body: "breaktimeId, gameId, and channel_type are required"}];
	}

	if (doc.channel_type != channel_type){
		return [null, {code: 400, body: "channel_type missmatch"}];
	}

	if (doc.game != game){
		return [null, {code: 400, body: "game missmatch"}];
	}

	if(!doc.users[userId] || doc.users[userId].date_exited){
		return [null, {code: 204}];
	}

	if (doc.users[userId].date_exited){
		delete doc.users[userId];
	}
	var today = new Date().getTime();

	if (Object.prototype.toString.call(delta) != '[object Array]'){
		return [null, {code: 400, body: "delta set incorrectly"}];
	}

	//evaluate options here for things
	//then descend to apply delta

//refactor this with te 0, -1, -2,-3 change
  function setChange(path, s, value) {
   	//handle later
    if (path.length == 0){
      return value;
    }

    var head = path[0];
    var otype = Object.prototype.toString.call(head);
    
    //init if null
    if (!s[head]){
 		switch(otype){
			case '[object Number]':
					s[head] = [];
				break;
			case '[object String]':
					s[head] = {};
				break;
		}
    }

    var val = setChange(path.slice(1), s[head], value);
    //merge step;
 	switch(otype){
		case '[object Number]':
			switch(head){
				//before beginning
				case -1:
					s.splice(0,0,val);
					break;
				// end of array
				case -2:
					s[s.length-1]=val;
					break;
				//append
				case -3:
					s.push(val);
					break;
				default:
					var pos = Math.min(s.length-1,Math.abs(head));
					s[pos] = val;
			}
			break;
		case '[object String]':
			s[head] = val;
			break;
	}
    return s;
  }

  function addChange(path, s, value) {
   	//handle later
    if (path.length == 0){
   		var result = addDelta(s, value);
      	return result;
    }

    var head = path[0];
    var otype = Object.prototype.toString.call(head);
    
    //init if null
    if (!s[head]){
 		switch(otype){
			case '[object Number]':
					s[head] = [];
				break;
			case '[object String]':
					s[head] = {};
				break;
		}
    }

    var val = addChange(path.slice(1), s[head], value);
    //merge step;
 	switch(otype){
		case '[object Number]':
			switch(head){
				//before beginning
				case -1:
					s.splice(0,0,val);
					break;
				// end of array
				case -2:
					s[s.length-1]=val;
					break;
				//append
				case -3:
					s.push(val);
					break;
				default:
					var pos = Math.min(s.length-1,Math.abs(head));
					s[pos] = val;
			}
			break;
		case '[object String]':
				s[head] = val;
			break;
	}
    return s;
  }
  // add function for get position..and other crap
  function howMany (argument) {
  	var otype = Object.prototype.toString.call(argument)
  	switch(otype){
  		case '[object Number]': 
  			return Math.abs(argument);
  		default:
  			return 1;
  	}
  	return 1;
  }
  function removeChange(path, s, value) {
  	var head = path[0];

    var otype = Object.prototype.toString.call(head);
    if (path.length == 1){
	 	switch(otype){
			case '[object Number]':
				switch(head){
					case 0:
					case -1:
					var pos = 0;
					if (!s[pos] && s[pos] != 0){
				    	return s;
				    }				
					var count = howMany(value);
						count = Math.min(s.length-1, count);
						s.splice(pos,count);
						break;
					// end of array
					case -2:
					case -3:	
						var count = howMany(value);
						var pos = Math.max(s.length-count, 0);
						if (!s[pos] && s[pos] != 0){
					    	return s;
					    }
						s.splice(pos,count);
						break;
					default:
						var count = howMany(value);
						var pos = Math.abs(head);
						pos = Math.min(s.length-1,pos);
						if (!s[pos] && s[pos] != 0){
					    	return s;
					    }
						s.splice(pos,count);
				}
					
				break;
			case '[object String]':
					// I could expand this out to do more but not for now...bother the client guys for this
					if (!s[head] && s[head] != 0){
				    	return s;
				    }
					delete s[head];
				break;
		}

			 	switch(otype){
			case '[object Number]':
				switch(head){
					case 0:
					case -1:
					var pos = 0;
					if (!s[pos] && s[pos] != 0){
				    	return s;
				    }				
					var count = howMany(value);
						count = Math.min(s.length-1, count);
						break;
					// end of array
					case -2:
					case -3:	
						var count = howMany(value);
						var pos = Math.max(s.length-count, 0);
						if (!s[pos] && s[pos] != 0){
					    	return s;
					    }
						break;
					default:
						var count = howMany(value);
						var pos = Math.abs(head);
						pos = Math.min(s.length-1,pos);
						if (!s[pos] && s[pos] != 0){
					    	return s;
					    }
				}
					
				break;
			case '[object String]':
					// I could expand this out to do more but not for now...bother the client guys for this
					if (!s[head] && s[head] != 0){
				    	return s;
				    }
				break;
		}
      return s;
    }

    s[head] = removeChange(path.slice(1), s[head], value);
    return s;
  }
	function addDelta (s, toAdd) {
		//handles the inductive step of null values
								log(JSON.stringify(o));
		var otype = Object.prototype.toString.call(toAdd);
		if (s && s!=0 && !toAdd && toAdd != 0){
			//this case should NEVER HAPPEN!!!
			return null;
		}

		if (!s && s != 0){
			return toAdd;
		}

		if (!toAdd && toAdd != 0){
			return s;
		}

		switch(otype){
			case '[object Object]':
					for (var o in toAdd){
						s[o] = addDelta(s[o], toAdd[o]);
					}
					return s;
				break;
			case '[object Array]':
					for (var i = 0; i < toAdd.length; i++) {
						s[i] = addDelta(s[i], toAdd[i]);
					}
					return s;
				break;
			case '[object Number]':
					return (s+toAdd);
				break;
			case '[object String]':
					return toAdd;
				break;
			case '[object Boolean]':
					return toAdd;
				break;
		}
	}
	var changed = false;
	for (var d = 0; d < delta.length; d++) {
		var change = delta[d];
		//path/ value / action
		switch(change.action){
			case 'add':
				doc.state = addChange(change.path, doc.state,change.value);
				changed = true;
				break;
			case 'remove':
				doc.state = removeChange(change.path, doc.state,change.value);
				changed = true;
				break;
			case 'set':
				doc.state = setChange(change.path, doc.state,change.value);
				changed = true;
				break;
		}
	}

	if (!changed){
		return [null, {code: 204}];
	}
	doc.users[userId].date_updated = today 
	doc.users[userId].total_updates++;
	doc.total_updates++
	doc.date_updated= today;
	doc.updater= userId;
	//probably send back the entire game;
	return [doc, JSON.stringify(doc.state)];
}