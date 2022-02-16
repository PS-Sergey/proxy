- Кэширующий http-прокси (обратный прокси).
- Утилита конфигурируема. В конфигурационном файле config/api.yaml задаются значения для следующих параметров:
    1) адрес бэкенда (сайт, который будем проксировать);
    2) локальный порт, на котором утилита слушает запросы;
    3) максимальный размер кэша, в количестве записей.


- Утилита работает следующим образом:

    1) принмает HTTP GET запрос;
    2) если url есть в кэше, передает в ответ клиенту файл из кэша;
    3) если url нет в кэше:

        - пересылает запрос на бэкенд;
        - читает ответ от бэкенда;
        - сохраняет запись в кэш, если достигнут максимальный размер кэша перезаписывает самую старую запись;
        ответ передает клиенту.