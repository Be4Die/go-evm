package vm

// интерфейс представляет исполняемую команду виртуальной машины
type Command interface {
    // Execute выполняет команду используя предоставленные CPU и Memory
    Execute(cpu CPU, memory Memory) error
}