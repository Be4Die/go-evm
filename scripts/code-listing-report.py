import os
import re
from reportlab.lib.pagesizes import A4
from reportlab.pdfgen import canvas
from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont

# -------- Настройки --------
PROJECT_ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
OUTPUT_FILE = os.path.join(PROJECT_ROOT, "out", "code-listing.pdf")

CODE_FONT_NAME = "CourierNew"
HEADER_FONT_NAME = "ArialBold"
CODE_FONT_PATH = "C:/Windows/Fonts/cour.ttf"
HEADER_FONT_PATH = "C:/Windows/Fonts/arialbd.ttf"

CODE_FONT_SIZE = 9
HEADER_FONT_SIZE = 9

PAGE_WIDTH, PAGE_HEIGHT = A4
MARGIN_LEFT = 40
MARGIN_RIGHT = 40
MARGIN_TOP = 40
MARGIN_BOTTOM = 40
COLUMN_GAP = 20

TAB_REPLACEMENT = " " * 4

# -------- Подготовка --------
pdfmetrics.registerFont(TTFont(CODE_FONT_NAME, CODE_FONT_PATH))
pdfmetrics.registerFont(TTFont(HEADER_FONT_NAME, HEADER_FONT_PATH))

# Ширина одной колонки
COLUMN_WIDTH = (PAGE_WIDTH - MARGIN_LEFT - MARGIN_RIGHT - COLUMN_GAP) / 2
LINE_HEIGHT = CODE_FONT_SIZE * 1.2


def collect_go_files():
    files = []
    for root, dirs, filenames in os.walk(PROJECT_ROOT):
        # исключаем ./scripts и ./out
        rel_root = os.path.relpath(root, PROJECT_ROOT)
        if rel_root.startswith("scripts") or rel_root.startswith("out"):
            continue
        for f in filenames:
            if f.endswith(".go"):
                files.append(os.path.join(root, f))
    return sorted(files)


def clean_line(line):
    line = line.replace("\t", TAB_REPLACEMENT)
    # убрать управляющие символы
    return re.sub(r"[^\x20-\x7E]", " ", line)


def wrap_line(line, num_str_width, c):
    """Перенос строки по ширине колонки"""
    max_chars = int((COLUMN_WIDTH - num_str_width) // pdfmetrics.stringWidth("M", CODE_FONT_NAME, CODE_FONT_SIZE))
    wrapped = []
    while line:
        wrapped.append(line[:max_chars])
        line = line[max_chars:]
    return wrapped


def measure_pages(files):
    """Первый проход: считаем, сколько страниц потребуется для каждого файла"""
    page_map = {}  # абсолютный_номер_страницы -> (файл, страница_внутри_файла)
    abs_page = 1
    for file in files:
        rel_path = os.path.relpath(file, PROJECT_ROOT)
        file_page = 1
        col = 0
        y = PAGE_HEIGHT - MARGIN_TOP - HEADER_FONT_SIZE * 2 - 4
        with open(file, "r", encoding="utf-8", errors="replace") as f:
            for i, line in enumerate(f, 1):
                line = clean_line(line.rstrip("\n"))
                num_str = f"{i:4d} "
                num_width = pdfmetrics.stringWidth(num_str, CODE_FONT_NAME, CODE_FONT_SIZE)
                wrapped = wrap_line(line, num_width, None)

                for j, seg in enumerate(wrapped):
                    if y - LINE_HEIGHT < MARGIN_BOTTOM:
                        # новая колонка или страница
                        if col == 0:
                            col = 1
                            y = PAGE_HEIGHT - MARGIN_TOP - HEADER_FONT_SIZE * 2 - 4
                        else:
                            col = 0
                            y = PAGE_HEIGHT - MARGIN_TOP - HEADER_FONT_SIZE * 2 - 4
                            abs_page += 1
                            file_page += 1
                            page_map[abs_page] = (rel_path, file_page)
                    y -= LINE_HEIGHT
        # после файла — новая страница
        abs_page += 1
        file_page += 1
        page_map[abs_page] = (rel_path, file_page)
    return page_map


def render_pdf(files, page_map):
    c = canvas.Canvas(OUTPUT_FILE, pagesize=A4)
    abs_page = 1
    for file in files:
        rel_path = os.path.relpath(file, PROJECT_ROOT)
        file_page = 1
        col = 0
        y = PAGE_HEIGHT - MARGIN_TOP - HEADER_FONT_SIZE * 2 - 4

        def draw_header():
            c.setFont(HEADER_FONT_NAME, HEADER_FONT_SIZE)
            left = f"{rel_path} — стр. {file_page}"
            right = f"стр. {abs_page}"
            c.drawString(MARGIN_LEFT, PAGE_HEIGHT - MARGIN_TOP, left)
            c.drawRightString(PAGE_WIDTH - MARGIN_RIGHT, PAGE_HEIGHT - MARGIN_TOP, right)
            c.line(MARGIN_LEFT, PAGE_HEIGHT - MARGIN_TOP - 2, PAGE_WIDTH - MARGIN_RIGHT, PAGE_HEIGHT - MARGIN_TOP - 2)
            c.setFont(CODE_FONT_NAME, CODE_FONT_SIZE)

        draw_header()
        c.setFont(CODE_FONT_NAME, CODE_FONT_SIZE)
        with open(file, "r", encoding="utf-8", errors="replace") as f:
            for i, line in enumerate(f, 1):
                line = clean_line(line.rstrip("\n"))
                num_str = f"{i:4d} "
                num_width = pdfmetrics.stringWidth(num_str, CODE_FONT_NAME, CODE_FONT_SIZE)
                wrapped = wrap_line(line, num_width, c)

                for j, seg in enumerate(wrapped):
                    if y - LINE_HEIGHT < MARGIN_BOTTOM:
                        # новая колонка или страница
                        if col == 0:
                            col = 1
                            y = PAGE_HEIGHT - MARGIN_TOP - HEADER_FONT_SIZE * 2 - 4
                        else:
                            col = 0
                            c.showPage()
                            abs_page += 1
                            file_page += 1
                            draw_header()
                            y = PAGE_HEIGHT - MARGIN_TOP - HEADER_FONT_SIZE * 2 - 4
                    x = MARGIN_LEFT + col * (COLUMN_WIDTH + COLUMN_GAP)
                    if j == 0:
                        c.drawString(x, y, num_str + seg)
                    else:
                        c.drawString(x + num_width, y, seg)
                    y -= LINE_HEIGHT
        # конец файла — новая страница
        c.showPage()
        abs_page += 1
        file_page += 1
    c.save()


def main():
    files = collect_go_files()
    page_map = measure_pages(files)
    render_pdf(files, page_map)
    print(f"Готово! PDF сохранён в {OUTPUT_FILE}")


if __name__ == "__main__":
    main()
