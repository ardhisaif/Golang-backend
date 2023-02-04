package reservation

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo}
}

func (r *service)GetAll(rep *repository) *helper.Response{
	response, err := r.repo.FindAll
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error()).Send(w)
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response).Send(w)
}