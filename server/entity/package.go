package entity

type Serializable interface {
	User | UidJSON | LoginJSON | RegisterJSON | Workspace | UserWorkspace
}
