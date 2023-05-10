package repository

type Repo struct {
	U UsersRepo
}

func NewRepo(ud UsersRepo) *Repo {
	return &Repo{
		U: ud,
	}
}
