package vm

// CPU интерфейс определяет основные операции центрального процессора
type CPU interface {
    // GetRegister возвращает текущее значение указанного регистра
    GetRegister(registerType RegisterType) (Word, error)

    // SetRegister устанавливает значение указанного регистра
    SetRegister(registerType RegisterType, value Word) error

    // Run начинает выполнение программного кода
    Run() error

    // Step выполняет одну следующую инструкцию
    Step() error

    // Reset сбрасывает состояние процессора в начальное положение
    Reset() error
}