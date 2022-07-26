package usecases

type SubGroupManage struct {
	repo SubGroupRepository
	//webApi SubGroupWebApi
}

type SubGroupRepository interface {
	MakeSubGroup()
	UpdateSubGroup()
	DeleteGroup()
}

func NewSubGroup(sgr SubGroupRepository) *SubGroupManage {
	return &SubGroupManage{repo: sgr}
}

func (sgm *SubGroupManage) MakeSubGroup() {

}
func (sgm *SubGroupManage) UpdateSubGroup() {

}
func (sgm *SubGroupManage) DeleteSubGroup() {

}
