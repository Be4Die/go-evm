; Вычисление факториала с вводом числа
entry main

section .data
N:       DW 0
RESULT:  DW 1
ONE:     DW 1
ZERO:    DW 0

section .code
main:    
    ; Ввод числа
    IN [N]
    
    ; Проверка на отрицательность
    PUSH [N]
    CMP_I [ZERO]
    JC POSITIVE      ; Если N >= 0, продолжаем
    JMP ERROR        ; Если отрицательное - ошибка

POSITIVE:
    ; Проверка на ноль
    PUSH [N]
    CMP_I [ZERO]
    JZ EXIT          ; Если ноль, результат = 1

    ; Основной цикл вычисления факториала
LOOP:
    PUSH [RESULT]
    MUL_I [N]        ; Умножаем текущий результат на N
    POP [RESULT]
    PUSH [N]
    SUB_I [ONE]      ; Уменьшаем N на 1
    POP [N]
    PUSH [N]
    CMP_I [ZERO]     ; Проверяем, достигли ли нуля
    JNZ LOOP         ; Если нет, продолжаем цикл
    JMP EXIT

ERROR:
    PUSH [ZERO]
    POP [RESULT]

EXIT:
    OUT [RESULT]     ; Вывод результата
    HALT             ; Остановка программы