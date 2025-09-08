package vm

// Memory интерфейс определяет операции работы с памятью виртуальной машины
type Memory interface {
    // SetValue записывает значение по указанному адресу
    // Возвращает ошибку если адрес недопустим
    SetValue(addr Word, value Word) error
    
    // GetValue читает значение по указанному адресу
    // Возвращает ошибку если адрес недопустим
    GetValue(addr Word) (Word, error)
}