#!/usr/bin/python3


##############################
Ania_zapl_dict = {149.0}        # Nie pamiętam za co.

Kamil_zapl_list = []

Kamil_zapl_list.append(100.0)  # Za książkę ,,Art of biological''.
# Mniej więcej tyle kosztowała, ale dokładną kwotę musiałbym sprawdzić
# najpierw na Amazonie, potem znaleźć kurs walut i przeliczyć, a tego
# mi się strasznie nie chce robić.
Kamil_zapl_list.append(30.99)  # 12 XII 2019 r.
# Górski Marek, Koczewski Andrzej


##############################
Ania_zapl = sum(Ania_zapl_dict.keys())
Kamil_zapl = sum(Kamil_zapl_list.keys())

roznica_kwot = Ania_zapl - Kamil_zapl


if roznica_kwot > 0.0:
    print("Należy oddać Anni: {} PLN".format(roznica_kwot))
else:
    print("Należy dać Kamilowi: {} PLN".format(roznica_kwot))
    print("Ależ może jej nie przypomnisz ;)?")


##############################
data_aktualizacji = "3 XII 2019"


print("Ostatnia aktualizacja pliku", data_aktualizacji)
