; Вычисление факториала
ORG 0x0200

START:
    PUSH [N]        ; Загружаем n в стек
    CMP_I [ONE]     ; Сравниваем с 1
    JZ EXIT         ; Если равно 1, выходим
    JNC EXIT        ; Если меньше 1, выходим
    PUSH [RESULT]   ; Загружаем текущий результат
    MUL_I [N]       ; Умножаем на n
    POP [RESULT]    ; Сохраняем результат
    PUSH [N]        ; Загружаем n
    SUB_I [ONE]     ; Вычитаем 1
    POP [N]         ; Сохраняем новое n
    JMP START       ; Повторяем цикл

EXIT:
    OUT [RESULT]    ; Выводим результат
    HALT            ; Останавливаем программу

; Секция данных
DS
N:      DW 5
RESULT: DW 1
ONE:    DW 1
DE