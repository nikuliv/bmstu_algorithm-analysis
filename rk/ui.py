import PySimpleGUI as sg
from typing import TypeVar
import matplotlib.pyplot as plt
import matplotlib.patches as pat
from matplotlib.ticker import NullFormatter
from matplotlib.backends.backend_tkagg import FigureCanvasTkAgg
import matplotlib
matplotlib.use('TkAgg')
from tkinter import messagebox as msg
from utils import SparseMatrix


def find_last_element(table):
    coefficient = 0
    while table[coefficient][0] != '   ':
        coefficient += 1
    return coefficient


def update_table(table_data, dots):
    if table_data[0][0] == '   ':
        table_data[0][0] = 1
        table_data[0][1] = dots[0]
        table_data[0][2] = dots[1]
        table_data[0][3] = dots[2]
    else:
        ind = find_last_element(table_data)
        table_data[ind][0] = ind + 1
        table_data[ind][1] = dots[0]
        table_data[ind][2] = dots[1]
        table_data[ind][3] = dots[2]


def change_point_value(table_data: list[list[str]], coefficient: int, coords: list[str]) -> None:
    check = find_last_element(table_data)
    if coefficient <= check and table_data[0][0] != '   ':
        table_data[coefficient - 1][1] = coords[0]
        table_data[coefficient - 1][2] = coords[1]
    else:
        msg.showinfo("Ошибка",
                     "Точка не найдена")


def del_point(table_data: list[list[str]], coefficient: int) -> None:
    check = find_last_element(table_data)
    if coefficient > check or table_data[0][0] == '   ':
        msg.showinfo("Ошибка",
                     "Точка не найдена")
    else:
        while table_data[coefficient][0] != '   ':
            # print(table_data[coefficient][0])
            table_data[coefficient][0] = table_data[coefficient + 1][0]
            table_data[coefficient][1] = table_data[coefficient + 1][1]
            table_data[coefficient][2] = table_data[coefficient + 1][2]
            try:
                table_data[coefficient][0] = str(int(table_data[coefficient][0]) - 1)
            except:
                pass
            coefficient += 1


def make_win_del():
    layout = [[sg.Text('Выберите матрицу: ', font=2), sg.Combo(['1', '2'], default_value='1', font=2)],
              [sg.Text('Введите номер элемента: ', font=2), sg.InputText('1', font=2, size=(5, 1))],
              [sg.Button('Удалить', font=2), sg.Button('Выход', font=2)]]
    return sg.Window('Удаление элемента', layout, finalize=True)


def make_win_ch():
    layout = [[sg.Text('Выберите множество: ', font=2), sg.Combo(['1', '2'], default_value='1', font=2)],
              [sg.Text('Введите номер точки: ', font=2), sg.InputText('1', font=2, size=(5, 1))],
              [sg.Text('Введите новые координаты: ', font=2), sg.InputText('1 1', font=2, size=(10, 1))],
              [sg.Button('Изменить', font=2), sg.Button('Выход', font=2)]]
    return sg.Window('Изменение координат точки', layout, finalize=True)


def make_win_res(data):
    if data is None:
        return
    sparse1: SparseMatrix = data[0]
    sparse2: SparseMatrix = data[1]
    res_sparse: SparseMatrix = data[2]
    layout = [[sg.Text('Матрица А:', font=2),
              sg.Text(sparse1.sparse_print(), font=2)],
              [sg.Text('Матрица B: ', font=2), sg.Text(sparse2.sparse_print(), font=2)],
              [sg.Text('Результат: ', font=2), sg.Text(res_sparse.sparse_print(), font=2)]
              ]
    return sg.Window('Результаты', layout, finalize=True, resizable=True, size=(320, 400))


def convert_table_data(set):
    dots = []
    i = 0
    while set[i][0] != '   ':
        dots.append(tuple([float(set[i][1]), float(set[i][2])]))
        i += 1
    return dots

