package models

type Questionnaire struct {
    ID         int    `json:"id"`
    QuestionID int    `json:"questionId"`
    Answer     string `json:"answer"`
    Date       string `json:"date"`
}

type Questionnaires []Questionnaire

func (questionnaire *Questionnaire) Create() error {
    if err := DB.Create(questionnaire).Error; err != nil {
        return err
    }
    return nil
}

func (questionnaire *Questionnaire) Delete(id int) error {
    if err := DB.Where("id = ?", id).Delete(questionnaire).Error; err != nil {
        return err
    }
    return nil
}

func (questionnaire *Questionnaire) Get(id int) error {
    if err := DB.Where("id = ?", id).Find(questionnaire).Error; err != nil {
        return err
    }
    return nil
}

func (questionnaires Questionnaires) GetAll() error {
    if err := DB.Find(questionnaires).Error; err != nil {
        return err
    }
    return nil
}
