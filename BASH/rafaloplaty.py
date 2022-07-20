#!/usr/bin/python3


def zaplacil_rok(dict_zapla, rok):
    Rafal_zaplacil = sum(dict_zapla.values())
    print("Rafał zapłacił za prąd, wodę i ogrzewanie w roku",
          "{}: {} PLN".format(rok, Rafal_zaplacil))


def zaplacil_kaucji_fun(dict_zapla):
    Rafal_zaplacil_kaucji = sum(dict_zapla.values())
    print("Rafał wpłacił już: {} PLN".format(Rafal_zaplacil_kaucji))
    print("Musi jeszcze zapłacić: {}".format(700 - Rafal_zaplacil_kaucji))


##############################
zaplacil_wynajem_2019 = {}


zaplacil_wynajem_2019["17 września 2019"] = 700
zaplacil_wynajem_2019["4 listopada 2019"] = 700
zaplacil_wynajem_2019["3 grudnia 2019"] = 700


##############################
zaplacil_wynajem_2020 = {}

zaplacil_wynajem_2020["7 stycznia 2020"] = 700
zaplacil_wynajem_2020["1 luty 2020"] = 700
zaplacil_wynajem_2020["28 luty 2020"] = 700


##############################
# Centralne ogrzewanie, ciepła woda, zimna woda, wywóz śmieci, prąd
zaplacil_prad_etc_2019 = {}


zaplacil_prad_etc_2019["koło 10 października 2019"] = 140
zaplacil_prad_etc_2019["4 listopada 2019"] = 140
zaplacil_prad_etc_2019["3 grudnia 2019"] = 140


##############################
# Centralne ogrzewanie, ciepła woda, zimna woda, wywóz śmieci, prąd
zaplacil_prad_etc_2020 = {}

zaplacil_prad_etc_2020["7 stycznia 2020"] = 50
zaplacil_prad_etc_2020["8 stycznia 2020"] = 90
zaplacil_prad_etc_2020["1 luty 2020"] = 100
zaplacil_prad_etc_2020["koło 5 lutego 2020"] = 40
zaplacil_prad_etc_2020["28 luty 2020"] = 140


##############################
# Ile ja zapłaciłem za prąd etc., w okresie od października 2019 roku
# do lutego 2020
suma_oplat_za_okres_2019_2020 = 0


##########
# Październik 2019

# c.o. - opłata stała
suma_oplat_za_okres_2019_2020 += 64.87
# c.o. - opłata zmienna
suma_oplat_za_okres_2019_2020 += 82.54
# ciepła woda
suma_oplat_za_okres_2019_2020 += 40.33
# zimna woda
suma_oplat_za_okres_2019_2020 += 32.78
# wywóz śmieci
suma_oplat_za_okres_2019_2020 += 16.96


# prąd
suma_oplat_za_okres_2019_2020 += 58.34


##########
# Listopad 2019


# c.o. - opłata stała
suma_oplat_za_okres_2019_2020 += 64.87
# c.o. - opłata zmienna
suma_oplat_za_okres_2019_2020 += 82.54
# ciepła woda
suma_oplat_za_okres_2019_2020 += 40.33
# zimna woda
suma_oplat_za_okres_2019_2020 += 32.78
# wywóz śmieci
suma_oplat_za_okres_2019_2020 += 16.96


# prąd
suma_oplat_za_okres_2019_2020 += 58.34


##########
# Grudzień 2019


# c.o. - opłata stała
suma_oplat_za_okres_2019_2020 += 64.87
# c.o. - opłata zmienna

# ciepła woda
suma_oplat_za_okres_2019_2020 += 40.33
# zimna woda
suma_oplat_za_okres_2019_2020 += 32.78
# wywóz śmieci
suma_oplat_za_okres_2019_2020 += 16.96


# prąd
suma_oplat_za_okres_2019_2020 += 58.34


##########
# Styczeń 2020


# c.o. - opłata stała
suma_oplat_za_okres_2019_2020 += 64.87
# c.o. - opłata zmienna

# ciepła woda
suma_oplat_za_okres_2019_2020 += 40.33
# zimna woda
suma_oplat_za_okres_2019_2020 += 32.78
# wywóz śmieci
suma_oplat_za_okres_2019_2020 += 16.96


# prąd
suma_oplat_za_okres_2019_2020 += 58.34


##########
# Luty 2020


# c.o. - opłata stała
suma_oplat_za_okres_2019_2020 += 64.87
# c.o. - opłata zmienna

# ciepła woda
suma_oplat_za_okres_2019_2020 += 40.33
# zimna woda
suma_oplat_za_okres_2019_2020 += 32.78
# wywóz śmieci
suma_oplat_za_okres_2019_2020 += 16.96


# prąd
suma_oplat_za_okres_2019_2020 += 58.34


##################################################
# Ceny mediów zmieniły się od 1 marca 2020
suma_oplat_za_okres_2020_pierwsza_polowa = 0

##########
# Marzec 2020
# Za marzec zapłacił 28 lutego 2020


# c.o. - opłata stała

# c.o. - opłata zmienna

# ciepła woda

# zimna woda

# wywóz śmieci


# prąd


##############################
kalcja_wplacil = {}


kalcja_wplacil["10 październiku 2019"] = 100
kalcja_wplacil["4 listopada 2019"] = 100
kalcja_wplacil["2 grudnia 2019"] = 100
kalcja_wplacil["8 stycznia 2020"] = 100
kalcja_wplacil["koło 5 lutego 2020"] = 100
kalcja_wplacil["28 luty 2020"] = 100
