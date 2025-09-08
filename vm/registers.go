package vm

// RegisterType идентифицирует конкретный регистр процессора
type RegisterType uint8

const (
    REG_IP RegisterType = iota    // Instruction Pointer - указатель на следующую инструкцию
    REG_ACC                       // Accumulator - аккумулятор для арифметических операций
    REG_SP                        // Stack Pointer - указатель вершины стека
    REG_FLAGS                     // FLAGS - регистр флагов (ZERO, CARRY, etc.)
    REG_BP                        // Base Pointer - указатель на базовый адрес стекового фрейма
    REG_R1                        // R1 - регистр общего назначения 1
    REG_R2                        // R2 - регистр общего назначения 2
    REG_R3                        // R3 - регистр общего назначения 3
    REG_R4                        // R4 - регистр общего назначения 4
    REG_R5                        // R5 - регистр общего назначения 5
    REG_R6                        // R6 - регистр общего назначения 6
    REG_R7                        // R7 - регистр общего назначения 7
    REG_R8                        // R8 - регистр общего назначения 8
)