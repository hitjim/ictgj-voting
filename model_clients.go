package main

type Client struct {
	UUID string
	Auth bool
	Name string
	IP   string
}

func dbGetAllClients() []Client {
	var ret []Client
	var err error
	if err = db.OpenDB(); err != nil {
		return ret
	}
	defer db.CloseDB()

	var clientUids []string
	if clientUids, err = db.GetBucketList([]string{"clients"}); err != nil {
		return ret
	}
	for _, v := range clientUids {
		if cl := dbGetClient(v); cl != nil {
			ret = append(ret, *cl)
		}
	}
	return ret
}

func dbGetClient(id string) *Client {
	var err error
	if err = db.OpenDB(); err != nil {
		return nil
	}
	defer db.CloseDB()

	cl := new(Client)
	cl.UUID = id
	cl.Auth = dbClientIsAuth(id)
	cl.Name, _ = db.GetValue([]string{"clients", id}, "name")
	cl.IP, _ = db.GetValue([]string{"clients", id}, "ip")
	return cl
}

func dbGetClientByIp(ip string) *Client {
	var err error
	if err = db.OpenDB(); err != nil {
		return nil
	}
	defer db.CloseDB()

	allClients := dbGetAllClients()
	for i := range allClients {
		if allClients[i].IP == ip {
			return &allClients[i]
		}
	}
	return nil
}

func dbSetClientName(cid, name string) error {
	var err error
	if err = db.OpenDB(); err != nil {
		return nil
	}
	defer db.CloseDB()

	err = db.SetValue([]string{"clients", cid}, "name", name)
	return err
}

func dbGetClientName(cid string) string {
	if err := db.OpenDB(); err != nil {
		return ""
	}
	defer db.CloseDB()

	name, _ := db.GetValue([]string{"clients", cid}, "name")
	return name
}

func dbAddDeauthClient(cid, ip string) error {
	var err error
	if err = db.OpenDB(); err != nil {
		return err
	}
	defer db.CloseDB()

	err = db.SetBool([]string{"clients", cid}, "auth", false)
	if err != nil {
		return err
	}
	return db.SetValue([]string{"clients", cid}, "ip", ip)
}

func dbAuthClient(cid, ip string) error {
	var err error
	if err = db.OpenDB(); err != nil {
		return err
	}
	defer db.CloseDB()

	err = db.SetBool([]string{"clients", cid}, "auth", true)
	if err != nil {
		return err
	}
	return db.SetValue([]string{"clients", cid}, "ip", ip)
}

func dbDeAuthClient(cid string) error {
	var err error
	if err = db.OpenDB(); err != nil {
		return err
	}
	defer db.CloseDB()

	return db.SetBool([]string{"clients", cid}, "auth", false)
}

func dbClientIsAuth(cid string) bool {
	var err error
	if err = db.OpenDB(); err != nil {
		return false
	}
	defer db.CloseDB()

	var isAuth bool
	if isAuth, err = db.GetBool([]string{"clients", cid}, "auth"); err != nil {
		return false
	}
	return isAuth
}

func dbUpdateClientIP(cid, ip string) error {
	var err error
	if err = db.OpenDB(); err != nil {
		return err
	}
	defer db.CloseDB()

	return db.SetValue([]string{"clients", cid}, "ip", ip)
}
