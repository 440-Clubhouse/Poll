package models

type Question struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Date  string `json:"date"`
}

type Questions []Question

func (question *Question) Create() error {
    if err := DB.Create(question).Error; err != nil {
        return err
    }
    return nil
}

func (question *Question) Delete(id int) error {
    if err := DB.Where("id = ?", id).Delete(question).Error; err != nil {
        return err
    }
    return nil
}

func (question *Question) Get(id int) error {
    if err := DB.Where("id = ?", id).Find(question).Error; err != nil {
        return err
    }
    return nil
}

func (questions *Questions) GetAll() error {
    if err := DB.Find(questions).Error; err != nil {
        return err
    }
    return nil
}
