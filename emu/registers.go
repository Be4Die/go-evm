package emu

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
)