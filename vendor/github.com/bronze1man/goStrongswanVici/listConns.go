package goStrongswanVici

import (
	"fmt"
)

func (c *ClientConn) ListConns(ike string) ([]map[string]IKEConf, error) {
	conns := []map[string]IKEConf{}
	var eventErr error
	var err error

	err = c.RegisterEvent("list-conn", func(response map[string]interface{}) {
		fmt.Printf("Entered the response function \n")
		conn := &map[string]IKEConf{}
		err = ConvertFromGeneral(response, conn)
		if err != nil {
			fmt.Printf("error from convert from general\n")
			eventErr = fmt.Errorf("list-conn event error: %v", err)
			return
		}
		fmt.Printf("Converted from general \n")
		conns = append(conns, *conn)
		fmt.Printf("Appended to conn\n")
	})

	if err != nil {
		return nil, fmt.Errorf("error registering list-conn event: %v", err)
	}

	if eventErr != nil {
		return nil, eventErr
	}

	reqMap := map[string]interface{}{}

	if ike != "" {
		reqMap["ike"] = ike
	}

	_, err = c.Request("list-conns", reqMap)
	if err != nil {
		return nil, fmt.Errorf("error requesting list-conns: %v", err)
	}

	err = c.UnregisterEvent("list-conn")
	if err != nil {
		return nil, fmt.Errorf("error unregistering list-conns event: %v", err)
	}

	return conns, nil
}
