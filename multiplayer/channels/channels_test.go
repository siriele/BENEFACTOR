package channels

import (
	"net/http"
	"testing"
)

func deleteChannel(client *http.Client, id, game string, t *testing.T) {
	db := "http://btuser:43gorilla@ec2-23-22-125-97.compute-1.amazonaws.com:5985/"
	r1, _ := http.Head(db + game + "_channels/" + id)
	rev := r1.Header.Get("ETag")
	rev = rev[1 : len(rev)-1]
	req, _ := http.NewRequest("DELETE", db+game+"_channels/"+id, nil)
	req.Header.Set("If-Match", rev)
	r, _ := client.Do(req)
	if r.StatusCode != 200 {
		t.Fatalf("Failed to delete channel with id %s and rev %s", id, rev)
	}
}

func TestMakeChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	deleteChannel(c, r["id"].(string), g, tt)
	tt.Log("created:", v)
	//fmt.Println(v)
}

func TestJoinChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	deleteChannel(c, id, g, tt)
}

func TestLeaveChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)

	l := LeaveChannel(g, u, t, id, c)
	tt.Log("left", l)
	deleteChannel(c, id, g, tt)
}

func TestCloseChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	cl := CloseChannel(g, u, t, id, c)
	tt.Log("closed:", cl)
	deleteChannel(c, id, g, tt)
}

func TestUpdatePrependChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	building := map[string]interface{}{
		"name": "garage",
		"cost": 1200,
		"paid": 100,
	}
	building2 := map[string]interface{}{
		"name": "apartment",
		"cost": 2200,
		"paid": 0,
	}
	path := []interface{}{
		"projects",
		BEFORE,
	}
	deltas := []Delta{
		*NewAddDelta(path, building),
		*NewAddDelta(path, building2),
	}
	del := UpdateChannel(g, u, t, id, c, deltas)
	tt.Log("changed:", del)
	cl := CloseChannel(g, u, t, id, c)
	tt.Log("closed:", cl)
	deleteChannel(c, id, g, tt)
}

func TestUpdateAppendChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	building := map[string]interface{}{
		"name": "garage",
		"cost": 2200,
		"paid": 100,
	}
	building2 := map[string]interface{}{
		"name": "apartment",
		"cost": 2300,
		"paid": 0,
	}
	path := []interface{}{
		"projects",
		AFTER,
	}
	deltas := []Delta{
		*NewAddDelta(path, building),
		*NewAddDelta(path, building2),
	}
	del := UpdateChannel(g, u, t, id, c, deltas)
	tt.Log("changed:", del)
	cl := CloseChannel(g, u, t, id, c)
	tt.Log("closed:", cl)
	deleteChannel(c, id, g, tt)
}

func TestUpdateTailHeadChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	building := map[string]interface{}{
		"name": "garage",
		"cost": 2200,
		"paid": 100,
	}
	building2 := map[string]interface{}{
		"name": "apartment",
		"cost": 2300,
		"paid": 0,
	}
	path := []interface{}{
		"projects",
		BEGINNING,
	}

	path2 := []interface{}{
		"projects",
		END,
	}
	deltas := []Delta{
		*NewAddDelta(path, building),
		*NewAddDelta(path2, building2),
	}
	del := UpdateChannel(g, u, t, id, c, deltas)
	tt.Log("changed:", del)
	cl := CloseChannel(g, u, t, id, c)
	tt.Log("closed:", cl)
	deleteChannel(c, id, g, tt)
}

func TestUpdateAddIndexChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    200,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	building := map[string]interface{}{
		"name": "garage",
		"cost": -20,
		"paid": 100,
	}
	building2 := map[string]interface{}{
		"name": "apartment",
		"cost": -30,
		"paid": 0,
	}
	path := []interface{}{
		"projects",
		1,
	}

	path2 := []interface{}{
		"projects",
		3,
	}
	deltas := []Delta{
		*NewAddDelta(path, building),
		*NewAddDelta(path2, building2),
	}
	del := UpdateChannel(g, u, t, id, c, deltas)
	tt.Log("changed:", del)
	cl := CloseChannel(g, u, t, id, c)
	tt.Log("closed:", cl)
	deleteChannel(c, id, g, tt)
}

func TestRemoveMultiIndexChannel(tt *testing.T) {
	c := &http.Client{}
	s := map[string]interface{}{
		"projects": []map[string]interface{}{
			map[string]interface{}{
				"name": "garage",
				"cost": 200,
				"paid": 10,
				"workers": []string{
					"Jim",
					"Joe",
				},
			},
			map[string]interface{}{
				"name":    "garage",
				"cost":    20,
				"paid":    10,
				"workers": []string{},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 2000,
				"paid": 10,
				"workers": []string{
					"John",
				},
			},
			map[string]interface{}{
				"name": "garage",
				"cost": 1200,
				"paid": 10,
			},
		},
	}
	o := NewOptions()
	o.MaxParticpants = 7
	t := "projects"
	g := "city"
	u := "3000000"
	var v interface{}
	v = MakeChannel(g, u, t, c, o, s)
	r := v.(map[string]interface{})
	id := r["id"].(string)

	j := JoinChannel(g, u, t, id, c)
	tt.Log("joined:", j)
	// building := map[string]interface{}{
	// 	"name": "garage",
	// 	"cost": -20,
	// 	"paid": 100,
	// }
	// building2 := map[string]interface{}{
	// 	"name": "apartment",
	// 	"cost": -30,
	// 	"paid": 0,
	// }
	path := []interface{}{
		"projects",
		1,
	}
	deltas := []Delta{
		*NewRemoveDelta(path, 5),
	}
	del := UpdateChannel(g, u, t, id, c, deltas)
	tt.Log("changed:", del)
	cl := CloseChannel(g, u, t, id, c)
	tt.Log("closed:", cl)
	//deleteChannel(c, id, g, tt)
}
