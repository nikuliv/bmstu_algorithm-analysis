
from ui import *
from utils import *
import time

sg.theme('LightGreen6')  # Add a touch of color
if __name__ == "__main__":
    #
    # a = [[0, 0, 0], [0, 1, 1], [0, 0, 0]]
    # b = [[0, 0, 0], [0, 0, 1], [0, 0, 0]]
    # sp1 = transform_matrix_to_sparse(a, 3, 3)
    # sp2 = transform_matrix_to_sparse(b, 3, 3)
    # print(sp1.matrix_print())
    # print(sp1.sparse_print())
    # print(sp2.matrix_print())
    # print(sp2.sparse_print())
    # c = sparse_matrix_addition(sp1, sp2)
    # print(c.matrix_print())
    # print(c.sparse_print())




    # Example of table data
    set_1 = [[' ' * 3, ' ' * 4, ' ' * 4, ' ' * 6] for col in range(100)]
    set_2 = [[' ' * 3, ' ' * 4, ' ' * 4, ' ' * 6] for col in range(100)]
    n, m = None, None
    mtx1 = None
    mtx2 = None

    # All the stuff inside the window.
    layout = [[sg.Text('N       M', font=5)],
              [sg.Input('3', size=(3, 2), font=4),
               sg.Input('3', size=(3, 2), font=4),
               sg.Button('Инициализировать', size=(15, 1), font=2)],
              [sg.Text('I          J       Value', font=5)],
               [sg.Input('0', size=(3, 2), font=4),
               sg.Input('0', size=(3, 2), font=4),
               sg.Input('0', size=(3, 2), font=4)],
              [sg.Button('Добавить значение в 1-ую матрицу', size=(22, 2), font=2),
               sg.Button('Добавить значение во 2-ую матрицу', size=(22, 2), font=2)],
              [sg.Table(set_1, headings=["№", "i", "j", "Value"], num_rows=10, key='set1', font=4),
               sg.Table(set_2, headings=["№", "i", "j", "Value"], num_rows=10, key='set2', font=4)],
              [sg.Button('Изменить', size=(14, 2), font=2), sg.Button('Удалить', size=(14, 2), font=2),
               sg.Button('Сложить', size=(14, 2), font=2)]]


    # Create the Window
    window = sg.Window('Рубежный контроль №1 Никуленко И.В.', layout)
    figure = None
    # Event Loop to process "events" and get the "values" of the inputs
    while True:
        event, values = window.read()
        if event == sg.WIN_CLOSED:  # if user closes window or clicks cancel
            break
        if event == 'Инициализировать':
            try:
                n, m = int(values[0]), int(values[1])
                mtx1 = [[0 for i in range(n)] for j in range(m)]
                mtx2 = [[0 for i in range(n)] for j in range(m)]
            except:
                msg.showinfo("Ошибка чтения",
                             "Некорректный ввод")
        if event == 'Удалить':
            delete_wind = make_win_del()
            while True:
                d_event, d_values = delete_wind.read()
                if d_event == sg.WIN_CLOSED:
                    break
                if d_event == 'Удалить':
                    if d_values[0] == "1":
                        del_point(set_1, int(d_values[1]) - 1)
                    else:
                        del_point(set_2, int(d_values[1]) - 1)
                elif d_event == 'Выход':
                    delete_wind.close()
                    break
                window['set1'].update(set_1)
                window['set2'].update(set_2)
        elif event == 'Добавить значение в 1-ую матрицу':
            try:
                if n and m:
                    elems = [int(values[2]), int(values[3]), int(values[4])]
                    if elems[0] > 0 and elems[1] > 0:
                        update_table(set_1, elems)
                        mtx1[elems[0]][elems[1]] = elems[2]
                    else:
                        raise Exception
                else:
                    raise Exception
            except:
                msg.showinfo("Ошибка чтения",
                             "Некорректный ввод")
        elif event == 'Добавить значение во 2-ую матрицу':
            try:
                if n and m:
                    elems = [int(values[2]), int(values[3]), int(values[4])]
                    if elems[0] > 0 and elems[1] > 0:
                        update_table(set_2, elems)
                        mtx2[elems[0]][elems[1]] = elems[2]
                    else:
                        raise Exception
                else:
                    raise Exception
            except:
                msg.showinfo("Ошибка чтения",
                             "Некорректный ввод")
        elif event == 'Изменить':
            change_wind = make_win_ch()
            while True:
                c_event, e_values = change_wind.read()
                if c_event == sg.WIN_CLOSED:
                    break
                if c_event == 'Изменить':
                    if e_values[0] == "1":
                        change_point_value(set_1, int(e_values[1]), e_values[2].split(" "))
                    else:
                        change_point_value(set_2, int(e_values[1]), e_values[2].split(" "))
                    change_wind.close()
                if c_event == 'Выход':
                    change_wind.close()
                    break
        elif event == 'Сложить':
            if n and m:
                sparse1 = transform_matrix_to_rm(mtx1, n, m)
                sparse2 = transform_matrix_to_rm(mtx2, n, m)
                start_time = time.time()
                res = sum_sparse(sparse1, sparse2)
                end_time = time.time()
                data = (sparse1, sparse2, res)
                res_wind = make_win_res(data)

                print("--- %s seconds ---" % (end_time - start_time))
            else:
                msg.showinfo("Ошибка чтения",
                             "Некорректный ввод")

        window['set1'].update(set_1)
        window['set2'].update(set_2)

    window.close()