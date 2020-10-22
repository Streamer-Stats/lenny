package observers

import "fmt"

type CreatePlayerObserver struct {
	username string
}

func (c *CreatePlayerObserver) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.username, itemName)
}

func (c *CreatePlayerObserver) GetID() string {
	return c.username
}
